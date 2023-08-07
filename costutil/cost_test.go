package utils

import (
	"fmt"
	"testing"
)

func TestAddCostNode(t *testing.T) {
	cost := NewCostAnalyse("start")
	cost.Add("1")
	fmt.Println(cost.Done())
}
