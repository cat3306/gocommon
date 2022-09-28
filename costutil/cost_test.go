package utils

import (
	"fmt"
	"testing"
)

func TestAddCostNode(t *testing.T) {
	cost := NewCostAnalyse("start")
	cost.Add("1")
	cost.Add("2")
	cost.Add("3")
	cost.Add("4")
	fmt.Println(cost.Done())
}
