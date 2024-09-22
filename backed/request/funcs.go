package request

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
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
func GetInfo(c *gin.Context) {
	staff_id, is := c.Get("staff_id")
	if !is {
		c.JSON(404, gin.H{
			"error":  "invalid token",
			"status": "0",
		})
	}
	acc := Account{
		StaffId: staff_id.(string),
	}
	Db.First(&acc)
	c.JSON(200, gin.H{
		"status": "1",
		"user":   acc,
	})
}

type Quest struct {
	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
	Drivers   []string  `json:"drivers,omitempty"`
	Type      string    `json:"type,omitempty"`
}
type QueryLine struct {
	Line_id         uint   `json:"line_id,omitempty"`
	Station_id      uint   `json:"station_id,omitempty"`
	Up_station_id   uint   `json:"up_station_id,omitempty"`
	Down_station_id uint   `json:"down_station_id,omitempty"`
	Lontitude       string `json:"lontitude,omitempty"`
	Latitude        string `json:"latitude,omitempty"`
	Station         string `json:"station,omitempty"`
	Up_station      string `json:"up_station,omitempty"`
	Down_station    string `json:"down_station,omitempty"`
}

func FindStationLine(c *gin.Context) {
	ques := c.Query("query")
	res := SubwayLine{}
	err := Db.Where("name = ? ", ques).First(&res).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "地铁线不存在",
		})
	}
	query := `select subway_station_subwayline.subway_line_id as line_id,subway_station_subwayline.subway_station_id as station_id ,subway_station_subwayline.up as up_station_id,subway_station_subwayline.down as down_station_id,c.lon as lontitude,c.lat as latitude,c.name as station,a.name as up_station,b.name as down_station
	from subwaystation as c ,subway_station_subwayline
								 join subwaystation as a
									  on subway_station_subwayline.up == a.id
								 join subwaystation as b
									  on subway_station_subwayline.down == b.id
	where subway_station_subwayline.subway_line_id == ? and subway_station_subwayline.subway_station_id  == c.id
	`
	var result []QueryLine
	if err = Db.Raw(query, res.LineId).Scan(&result).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    0,
			"message": "查找失败",
		})
	}

	//var result []SubwayLine
	//Db.Model(&SubwayLine{}).Preload("SubwayStations.Subwaylines").Where("name = ?", ques).Find(&res)
	//jsonr, _ := json.Marshal(&result)
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"num":      len(result),
		"Stations": result,
	})
}
func FindStation(c *gin.Context) {
	ques := c.Query("query")
	res := SubwayStation{}
	err := Db.Where("name = ? ", ques).First(&res).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    0,
			"message": "地铁站不存在",
		})
	}
	query := `
	select subway_station_subwayline.subway_line_id as line_id,subway_station_subwayline.subway_station_id as station_id ,subway_station_subwayline.up as up_station_id,subway_station_subwayline.down as down_station_id,c.lon as lontitude,c.lat as latitude,c.name as station,a.name as up_station,b.name as down_station
	from subwaystation as c ,subway_station_subwayline
								 join subwaystation as a
									  on subway_station_subwayline.up == a.id
								 join subwaystation as b
									  on subway_station_subwayline.down == b.id
	where subway_station_subwayline.subway_station_id == ? and subway_station_subwayline.subway_station_id  == c.id
	`
	var result []QueryLine
	if err = Db.Raw(query, res.ID).Scan(&result).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    0,
			"message": "查找失败",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"num":      len(result),
		"Stations": result,
	})
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
