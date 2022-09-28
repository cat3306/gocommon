package utils

import (
	"fmt"
	"time"
)

//简单性能分析工具
type costAnalyseNode struct {
	Name string
	At   time.Time
}
type CostAnalyse []costAnalyseNode

func NewCostAnalyse(node string) CostAnalyse {
	c := CostAnalyse(make([]costAnalyseNode, 0))
	c.Add(node)
	return c
}
func (c *CostAnalyse) Add(node string) {
	*c = append(*c, costAnalyseNode{
		Name: node,
		At:   time.Now(),
	})
}

func (c CostAnalyse) Done() string {

	if len(c) < 2 {
		return ""
	}

	var str string
	str = fmt.Sprintf("[total:%d ms] ", c[len(c)-1].At.Sub(c[0].At)/1000000)
	for i, costNode := range c {
		if i == len(c)-1 {
			break
		}
		nextNode := c[i+1]
		costTime := nextNode.At.Sub(costNode.At).Nanoseconds() / 1000000
		str += fmt.Sprintf("[%s~%s:%d ms]", costNode.Name, nextNode.Name, costTime)
	}
	return str
}
