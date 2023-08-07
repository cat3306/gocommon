package table

import (
	"encoding/json"
	"fmt"
	"testing"
)

type House struct {
	Name  string
	Sigil string
	Motto string
}
type SurvivalStatisticsData struct {
	Day             string   `json:"day"`
	DaySurvivalRate []string `json:"day_survival_rate"`
}

func Test1(t *testing.T) {
	by := `[
  {
    "day": "2022-09-30",
    "day_survival_rate": [
      "19.44%",
      "7.42%",
      "4.39%",
      "3.24%",
      "2.59%",
      "2.38%",
      "2.81%",
      "",
      ""
    ]
  },
  {
    "day": "2022-10-01",
    "day_survival_rate": [
      "16.23%",
      "6.55%",
      "4.08%",
      "3.01%",
      "2.63%",
      "2.29%",
      "1.03%",
      "",
      ""
    ]
  },
  {
    "day": "2022-10-02",
    "day_survival_rate": [
      "14.47%",
      "5.55%",
      "3.88%",
      "3.23%",
      "2.43%",
      "1.31%",
      "",
      "",
      ""
    ]
  },
  {
    "day": "2022-10-03",
    "day_survival_rate": [
      "14.31%",
      "6.24%",
      "4.22%",
      "3.29%",
      "1.74%",
      "",
      "",
      "",
      ""
    ]
  },
  {
    "day": "2022-10-04",
    "day_survival_rate": [
      "15.65%",
      "6.54%",
      "4.29%",
      "2.89%",
      "",
      "",
      "",
      "",
      ""
    ]
  },
  {
    "day": "2022-10-05",
    "day_survival_rate": [
      "12.96%",
      "5.39%",
      "2.70%",
      "",
      "",
      "",
      "",
      "",
      ""
    ]
  },
  {
    "day": "2022-10-06",
    "day_survival_rate": [
      "12.56%",
      "4.14%",
      "",
      "",
      "",
      "",
      "",
      "",
      ""
    ]
  },
  {
    "day": "2022-10-07",
    "day_survival_rate": [
      "11.09%",
      "",
      "",
      "",
      "",
      "",
      "",
      "",
      ""
    ]
  }
]`
	list := make([]SurvivalStatisticsData, 0)
	err := json.Unmarshal([]byte(by), &list)
	if err != nil {
		fmt.Println(err)
	}
	//s:=House{
	//	Name:  "中国ker",
	//	Sigil: "s是s",
	//	Motto: "sa、就d",
	//}
	tt := Table(list)

	fmt.Println(tt)
}
