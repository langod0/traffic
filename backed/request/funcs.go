package request

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"main/binary"
	"math"
	"net/http"
	"os"
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
	ld := 0
	for i := startDate; !i.Equal(endDate); i = i.AddDate(0, 0, 1) {
		var tt Shift
		tt.Date = i.Format("2006-01-02")
		for i, j := 0, ld; i < 4; i++ {
			ld = (ld + 1) % tot
			tt.Allot = append(tt.Allot, struct {
				Shift string `json:"shift"`
				Class string `json:"class"`
			}{Class: classes[i], Shift: cycle[j]})

			j = (j + 1) % tot
		}
		ld = (ld + 1) % tot
		shifts = append(shifts, tt)
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
	Name      string   `json:"name" binding:"required"`
	StartTime string   `json:"startTime,omitempty" binding:"required"`
	EndTime   string   `json:"endTime,omitempty" binding:"required"`
	Drivers   []string `json:"drivers,omitempty" binding:"required"`
	Trains    []string `json:"trains,omitempty" binding:"required"`
	Type      string   `json:"type,omitempty" binding:"required"`
}
type QueryLine struct {
	Line_id         uint   `json:"line_id,omitempty"`
	Station_id      uint   `json:"station_id,omitempty"`
	Up_station_id   uint   `json:"up_station_id"`
	Down_station_id uint   `json:"down_station_id"`
	Lontitude       string `json:"lontitude,omitempty"`
	Latitude        string `json:"latitude,omitempty"`
	Station         string `json:"station,omitempty"`
	Up_station      string `json:"up_station,omitempty"`
	Down_station    string `json:"down_station,omitempty"`
}

// 查找线路
func FindStationLine(c *gin.Context) {
	ques := c.Query("query")
	res := SubwayLine{}
	err := Db.Where("name = ? ", ques).First(&res).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"error": "地铁线不存在",
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
		c.JSON(200, gin.H{
			"code":  0,
			"error": "查找失败",
		})
	}

	//var result []SubwayLine
	//Db.Model(&SubwayLine{}).Preload("SubwayStations.Subwaylines").Where("name = ?", ques).Find(&res)
	//jsonr, _ := json.Marshal(&result)
	c.JSON(http.StatusOK, gin.H{
		"code":     1,
		"num":      len(result),
		"Stations": result,
	})
}

// 查找站点
func FindStation(c *gin.Context) {
	ques := c.Query("query")
	res := SubwayStation{}
	err := Db.Where("name = ? ", ques).First(&res).Error
	if err != nil || res.ID == 0 {
		c.JSON(200, gin.H{
			"code":  0,
			"error": "地铁站不存在",
		})
		return
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
			"code":  0,
			"error": "查找失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     1,
		"num":      len(result),
		"Stations": result,
	})
}

// 计算排班表
func CreateSchedule(c *gin.Context) {
	val, has := c.Get("isadmin")
	val2, has2 := c.Get("staff_id")
	if !has || !has2 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  0,
			"error": "Invalid token",
		})
		c.Abort()
		if binary.Setting.Debug {
			binary.DebugLog.Println(has, has2)
		}
		return
	}
	can := val.(bool)
	if !can {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  0,
			"error": "no right to create",
		})
		c.Abort()
		return
	}
	var quest Quest
	err := c.BindJSON(&quest)
	if err != nil {
		c.JSON(417, gin.H{
			"code":  0,
			"error": "failed to bind JSON",
		})
		c.Abort()
		return
	}
	if quest.Type != "1" && quest.Type != "2" && quest.Type != "3" {
		c.JSON(417, gin.H{
			"code":  0,
			"error": "bad request",
		})
		c.Abort()
		return
	}
	var jsondata ScheduleJson
	tot := len(quest.Drivers)
	number := len(quest.Trains)
	if tot%number != 0 {
		c.JSON(417, gin.H{
			"code":  0,
			"error": "bad request",
		})
		c.Abort()
		return
	}
	jsondata.Drivers = make([][]string, tot/number)

	if quest.Type == "2" {
		if len(quest.Drivers)%3 != 0 {
			c.JSON(417, gin.H{
				"code":  0,
				"error": "bad request",
			})
			c.Abort()
			return
		}
		jsondata.Type = ThreeDayCycle
	} else {
		if len(quest.Drivers)%4 != 0 {
			c.JSON(417, gin.H{
				"code":  0,
				"error": "bad request",
			})
			c.Abort()
			return
		}
		if quest.Type == "1" {
			jsondata.Type = FourDayCycle
		} else {
			jsondata.Type = FourDayCycle2
		}
	}
	var res Schedule
	user := make([]string, 0)
	mm := make(map[string]bool)
	var num int64
	for idx, v := range quest.Drivers {
		if Db.Model(&Account{}).Where("staff_id = ?", v).Count(&num); num == 0 {
			user = append(user, v)
		}
		jsondata.Drivers[idx/number] = append(jsondata.Drivers[idx/number], quest.Drivers[idx])
		mm[v] = true
	}
	if 0 != len(user) {
		c.JSON(417, gin.H{
			"code":  0,
			"error": "users not found",
			"users": user,
		})
		c.Abort()
		return
	}
	if len(mm) != tot {
		c.JSON(417, gin.H{
			"code":  0,
			"error": "driver repeated",
		})
		c.Abort()
		return
	}
	user = make([]string, 0)
	mm = make(map[string]bool)
	for _, v := range quest.Trains {
		if Db.Model(&Train{}).Where("id = ?", v).Count(&num); num == 0 {
			user = append(user, v)
		}
		jsondata.Trains = append(jsondata.Trains, v)
		mm[v] = true
	}
	if len(mm) != number {

		c.JSON(417, gin.H{
			"code":  0,
			"error": "train repeated",
		})
		c.Abort()
		return
	}
	if 0 != len(user) {
		c.JSON(417, gin.H{
			"code":   0,
			"error":  "trains not found",
			"trains": user,
		})
		c.Abort()
		return
	}
	var st, fi time.Time
	st, err = time.Parse("2006-01-02", quest.StartTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  0,
			"error": "time error",
		})
		c.Abort()
		return
	}
	fi, err = time.Parse("2006-01-02", quest.EndTime)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"code":  0,
			"error": "time error",
		})
		c.Abort()
		return
	}
	res = GenerateSchedule(st.Truncate(time.Hour*24), fi.Truncate(time.Hour*24), quest.Type)
	jsondata.Schedule = res
	jsondata.StartTime = quest.StartTime
	jsondata.EndTime = quest.EndTime
	//jsondata.StartTime = quest.StartTime.Format("2006-01-02")
	//jsondata.EndTime = quest.EndTime.Format("2006-01-02")
	var schedule WorkingSchedule
	schedule.Name = quest.Name
	schedule.Filename = time.Now().Format("2006-01-02_03-04_05") + val2.(string) + ".json"
	file, err := os.Create("./data/" + schedule.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  0,
			"error": "system error",
		})
		c.Abort()
		return
	}
	W := bufio.NewWriter(file)
	data, _ := json.Marshal(jsondata)
	W.Write(data)
	W.Flush()
	defer file.Close()
	var tmp WorkingSchedule
	if err := Db.Model(&WorkingSchedule{}).Order("id DESC").First(&tmp).Error; err != nil {
		tmp.ID = 2
	} else {
		tmp.ID++
	}
	tmp.Name = quest.Name
	tmp.Filename = schedule.Filename
	Db.Create(&tmp)
	// TODO: sql save
	//Db.Where("id ")
	c.JSON(200, gin.H{
		"code":     1,
		"schedule": res,
	})
}
func GetSchedules(c *gin.Context) {
	val, has := c.Get("isadmin")
	if !has {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  0,
			"error": "Invalid token",
		})
		c.Abort()
		return
	}
	can := val.(bool)
	if !can {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  0,
			"error": "no right to create",
		})
		c.Abort()
		return
	}
	var Qs []WorkingSchedule
	Db.Model(&WorkingSchedule{}).Where("id != 0").Find(&Qs)
	var result []ScheduleJson
	for _, v := range Qs {
		file, err := os.Open("./data/" + v.Filename)
		if err != nil {
			binary.DebugLog.Println(err)
			c.JSON(200, gin.H{
				"code":  0,
				"error": "file not found ",
			})
			c.Abort()
			return
		}
		rd := json.NewDecoder(file)
		var schedule ScheduleJson
		err = rd.Decode(&schedule)
		if err != nil {
			binary.DebugLog.Println(err)
			c.JSON(200, gin.H{
				"code":  0,
				"error": "system error ",
			})
			c.Abort()
			return
		}
		result = append(result, schedule)
		file.Close()
	}
	c.JSON(200, gin.H{
		"code":      1,
		"num":       len(result),
		"schedules": result,
	})
}

// 查询用户
func GetUsers(c *gin.Context) {
	var result []Account
	db := Db.Model(&Account{})
	if c.Query("post") != "" {
		db = db.Where("Post = ?", c.Query("post"))
	} else {
		//db = db.Where("Post = ?", "司机")
	}
	if c.Query("name") != "" {
		db = db.Where(fmt.Sprintf("name like %q", "%"+c.Query("name")))
	}
	if c.Query("staff") != "" {
		db = db.Where("staff = ?", c.Query("staff"))
	}
	err := db.Find(&result).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 0,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"num":     len(result),
		"drivers": result,
	})
}

func GetRelations(c *gin.Context) {
	var lines []SubwayLine
	Db.Where("line_id != 0").Find(&lines)
	var result []SubwayStationSubwayline
	Db.Find(&result)
	c.JSON(200, gin.H{
		"code":   1,
		"result": result,
	})

}

// 获取个人信息
func GetInfo(c *gin.Context) {
	staff_id, is := c.Get("staff_id")
	if !is {
		c.JSON(404, gin.H{
			"error":  "invalid token",
			"status": "0",
		})
		c.Abort()
		return
	}
	acc := Account{
		StaffId: staff_id.(string),
	}
	val, has := c.Get("isadmin")
	if !has {
		c.JSON(404, gin.H{
			"code":  0,
			"error": "invalid token",
		})
		c.Abort()
		return
	}
	isadmin := val.(bool)
	Db.Where("staff_id = ?", acc.StaffId).First(&acc)

	var stations []SubwayStation
	var lines []SubwayLine
	var trains []Train
	var subs []Submission
	Db.Where("id != 0").Find(&stations)
	Db.Where("line_id != 0").Order("line_id ASC").Find(&lines)
	Db.Where("id != '无'").Find(&trains)
	if isadmin {
		Db.Find(&subs)
	} else {
		Db.Where("user_id = ?", acc.StaffId).Find(&subs)
	}
	c.JSON(200, gin.H{
		"code":     1,
		"user":     acc,
		"lines":    lines,
		"stations": stations,
		"trains":   trains,
		"subs":     subs,
		//"schedule": schedule,
	})
}
