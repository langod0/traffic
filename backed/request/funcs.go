package request

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"time"
)

const (
	WhiteShift     = "白班"
	NightShift     = "夜班"
	Rest           = "休息"
	MorningShift   = "早班"
	AfternoonShift = "中班"
)

var (
	FourDayCycle = []string{
		WhiteShift, NightShift, Rest, Rest,
	}
	ThreeDayCycle = []string{
		WhiteShift, NightShift, Rest,
	}
	FourDayCycle2 = []string{
		MorningShift, AfternoonShift, NightShift, Rest,
	}
)

type Shift struct {
	Date  string `json:"date"`
	Allot []struct {
		Shift string `json:"shift"`
		Class string `json:"class"`
	} `json:"allot"`
}
type Schedule []Shift

func (s Schedule) Print() {
	for _, v := range s {
		fmt.Println(v.Date)
		for _, e := range v.Allot {
			fmt.Printf("班次: %s班, 班次类型: %s\n", e.Class, e.Shift)
		}
	}
}
func GenerateSchedule(startDate, endDate time.Time, Type string) Schedule {
	shifts := make(Schedule, 0)
	var classes []string
	var cycle []string
	if Type == "1" {
		classes = []string{"A", "B", "C", "D"}
		cycle = FourDayCycle
	} else if Type == "2" {
		classes = []string{"A", "B", "C"}
		cycle = ThreeDayCycle
	} else if Type == "3" {
		classes = []string{"A", "B", "C", "D"} // 四个班次
		cycle = FourDayCycle2
	}
	tot := len(cycle)
	ld, rd := 0, tot-1
	for i := startDate; !i.Equal(endDate); i = i.AddDate(0, 0, 1) {
		var tt Shift
		tt.Date = i.Format("2006-01-02")
		for ld != rd {
			tt.Allot = append(tt.Allot, struct {
				Shift string `json:"shift"`
				Class string `json:"class"`
			}{Class: classes[ld], Shift: cycle[ld]})
			ld = (ld + 1) % tot
		}
		shifts = append(shifts, tt)
		rd = (rd + tot - 1) % tot
	}
	return shifts
}

type Re struct {
	Maxflow       int64   `json:"maxflow"`        // max flow
	RunTime       float64 `json:"runtime"`        // 全周转时长
	RatedCapacity int64   `json:"rated_capacity"` // 列车定员数
	Departure     float64 `json:"departure"`      // 发车间隔
	Type          string  `json:"type"`           // 计算类型
}
type Recommend struct {
	Trains    int64 `json:"trains"`     //
	AllTrains int64 `json:"all_trains"` //
}

func CalcSubway(c *gin.Context) {

	var rec Re
	err := c.BindJSON(&rec)
	if err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}
	var res Recommend
	if rec.Type == "1" {
		res.Trains = int64(math.Round(rec.RunTime / rec.Departure))
		res.Trains += res.Trains % 2
	} else if rec.Type == "2" {
		res.Trains = int64(math.Ceil(float64(rec.Maxflow)/float64(rec.RatedCapacity)*rec.RunTime)) * 2
	}
	res.AllTrains = int64(math.Round(float64(res.Trains) * 1.2))

}

type Quest struct {
	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
	Drivers   []string  `json:"drivers,omitempty"`
	Type      string    `json:"type,omitempty"`
}

func CalcSchedule(c *gin.Context) {
	var quest Quest
	err := c.BindJSON(&quest)
	if err != nil {
		c.JSON(502, gin.H{
			"error": err.Error(),
		})
		return
	}
	if quest.Type != "1" && quest.Type != "2" && quest.Type != "3" {
		c.JSON(502, gin.H{
			"error": "Type Error",
		})
	}
	var res Schedule
	res = GenerateSchedule(quest.StartTime.Truncate(time.Hour*24), quest.EndTime.Truncate(time.Hour*24), quest.Type)
	c.JSON(200, gin.H{
		"schedule": res,
	})
}
