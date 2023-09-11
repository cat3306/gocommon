package confutil

import (
	"context"
	"errors"
	"fmt"
	etcd "go.etcd.io/etcd/client/v3"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

const (
	fileMod = 0644
)

type IConf interface {
	Save(string, interface{}) error
	Load(string, dst interface{}) error
}
type Config struct {
	etcClient *etcd.Client
	lock      sync.RWMutex
}

func (d *Config) EtcdMode(clint *etcd.Client) *Config {
	d.etcClient = clint
	return d
}
func (d *Config) Save(filename string, data interface{}) error {

	if data == nil {
		return errors.New("data must not nil")
	}
	d.lock.Lock()
	defer d.lock.Unlock()

	if d.etcClient != nil {
		return d.saveFileConfigEtcd(filename, data)
	}
	// Save data.
	return d.saveFileConfig(filename, data)
}
func (d *Config) saveFileConfigEtcd(fileName string, v interface{}) error {
	raw, err := marshal(fileName, v)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err = d.etcClient.Put(ctx, fileName, string(raw))
	if err == context.DeadlineExceeded {
		return fmt.Errorf("etcd setup is unreachable, please check your endpoints %s", d.etcClient.Endpoints())
	} else if err != nil {
		return fmt.Errorf("unexpected error %w returned by etcd setup, please check your endpoints %s", err, d.etcClient.Endpoints())
	}
	return nil
}
func (d *Config) saveFileConfig(fileName string, v interface{}) error {

	confRaw, err := marshal(fileName, v)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, fileMod)
	if err != nil {
		return err
	}
	defer f.Close()
	n, err := f.Write(confRaw)
	if err != nil {
		return err
	}
	if n != len(confRaw) {
		return fmt.Errorf("write incomplete n:%d,raw.len:%d", n, len(confRaw))
	}
	return f.Sync()
}

func (d *Config) Load(fileName string, dst interface{}) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.etcClient != nil {
		return d.loadConfigFromEtcd(fileName, dst)
	}

	//TODO check
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	return unMarshal(fileName, raw, dst)
}

func (d *Config) loadConfigFromEtcd(fileName string, v interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	rsp, err := d.etcClient.Get(ctx, fileName)
	if err != nil {
		if err == context.DeadlineExceeded {
			return fmt.Errorf("etcd setup is unreachable, please check your endpoints %s", d.etcClient.Endpoints())
		}
		return fmt.Errorf("unexpected error %w returned by etcd setup, please check your endpoints %s", err, d.etcClient.Endpoints())
	}
	if rsp.Count == 0 {
		return os.ErrNotExist
	}
	for _, ev := range rsp.Kvs {
		if string(ev.Key) == fileName {
			fileData := ev.Value
			// Unmarshal file's content
			return unMarshal(fileName, fileData, v)
		}
	}
	return os.ErrNotExist
}
