package request

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"main/binary"
	"os"
	"time"
)

var (
	Db *gorm.DB
)

const (
	dsn = `sqlserver://%s:%s@localhost:%s?database=%s`
)

type ScheduleJson struct {
	Schedule  []Shift    `json:"schedule"`
	Drivers   [][]string `json:"drivers"`
	Trains    []string   `json:"trains"`
	Type      []string   `json:"type"`
	StartTime string     `json:"startTime,omitempty" binding:"required"`
	EndTime   string     `json:"endTime,omitempty" binding:"required"`
}
type WorkingSchedule struct {
	ID       uint   `gorm:"type:int;not null;primary key;unique;" json:"id" binding:"required"`
	Name     string `gorm:"type:string" json:"name" binding:"required"`
	UseIn    uint   `gorm:"type:int;default:0" json:"useIn" binding:"required"`
	Filename string `gorm:"type:string" json:"filename" binding:"required"`
}
type Account struct {
	Name       string `gorm:"type:varchar(20);not null;comment:'姓名'" json:"name" binding:"required"`
	Password   string `gorm:"size:255;not null" json:"-" binding:"required"`
	Email      string `gorm:"type:varchar(60);not null;comment:'邮箱'" json:"email" binding:"required"`
	Phone      string `gorm:"type:varchar(20);comment:'手机号'" json:"phone" binding:"required"`
	IsAdmin    bool   `gorm:"type:boolean ;comment:'是否为管理员'" binding:"required"`
	StaffId    string `gorm:"type:varchar(20);unique;primaryKey;comment:'工号'" json:"staff_id" binding:"required"`
	Post       string `gorm:"type:varchar(20);comment:'职位'" json:"post" binding:"required"`
	Sex        string `gorm:"type:varchar(6);default:'male';comment:'性别'" json:"sex" binding:"required"`
	ScheduleID uint   `gorm:"type:int;" json:"schedule_id" binding:"required"`
	TrainId    string `gorm:"type:varchar(30); comment:'所属列车编号'" json:"train_id" binding:"required"`
	gorm.Model `json:"-"`
}
type Train struct {
	ID       string    `gorm:"type:varchar(30);not null;primaryKey;unique;comment:'列车编号'" json:"id" binding:"required"`
	Drivers  []Account `json:"Drivers,omitempty"`
	LineId   uint      `gorm:"type:int" json:"line_id" binding:"required"`
	Capacity uint      `gorm:"type:int" json:"capacity" binding:"required"`
}

type SubwayLine struct {
	LineId uint    `gorm:"type:int;not null;unique;primary key;index;" json:"line_id" binding:"required"`
	Name   string  `gorm:"type:varchar(20);not null;" json:"name" binding:"required"`
	Trains []Train `gorm:"foreignKey:LineId;references:LineId" json:"Trains,omitempty"`

	SubwayStations []SubwayStation `gorm:"many2many:subway_station_subwayline;	foreignKey:LineId;references:ID;joinForeignKey:SubwayLineId;joinReferences:SubwayStationId;" json:"subway_stations,omitempty"`
}
type SubwayStation struct {
	ID          uint         `gorm:"type:integer;primary key;index" json:"id" binding:"required"`
	Lon         string       `gorm:"type:varchar(40);not null;" json:"lon" binding:"required"`
	Lat         string       `gorm:"type:varchar(40);not null;" json:"lat" binding:"required"`
	Name        string       `gorm:"type:varchar(20);not null;" json:"name" binding:"required"`
	Subwaylines []SubwayLine `gorm:"many2many:subway_station_subwayline;foreignKey:ID;references:LineId;joinForeignKey:SubwayStationId;joinReferences:SubwayLineId;" json:"subwaylines,omitempty" binding:"required"`
}
type SubwayStationSubwayline struct {
	SubwayLineId    uint `gorm:"primaryKey"` // 外键：SubwayLine 的 LineId
	SubwayStationId uint `gorm:"primaryKey"` // 外键：SubwayStation 的 ID

	Up   uint `gorm:"type:int;comment:'上行' default:0;" json:"up" binding:"required"`
	Down uint `gorm:"type:int;comment:'下行' default:0 " json:"down" binding:"required"`
}

func (a *WorkingSchedule) TableName() string {
	return "working_schedule"
}
func (a *SubwayStationSubwayline) TableName() string {
	return "subway_station_subwayline"
}
func (a *SubwayLine) TableName() string {
	return "subwayline"
}
func (a *SubwayStation) TableName() string {
	return "subwaystation"
}
func (a *Train) String() string {
	return "train"
}
func (a *Account) TableName() string {
	return "account"
}

func DBInit() {
	var err error
	//var DSN = fmt.Sprintf(dsn, binary.Setting.Sqlusername, binary.Setting.Sqlpassword, binary.Setting.Sqlport, binary.Setting.Sqlbase)
	//binary.InfoLog.Println(DSN)
	Db, err = gorm.Open(sqlite.Open("Subway.db"), &gorm.Config{})
	if err != nil {
		os.Exit(15)
		binary.DebugLog.Println(err)
	}
	sqlDB, _ := Db.DB()
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	//  设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	//  设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 10秒钟
	Db.AutoMigrate(&Account{})
	Db.AutoMigrate(&Train{})
	Db.AutoMigrate(&SubwayStation{})
	Db.AutoMigrate(&SubwayLine{})
	Db.AutoMigrate(&SubwayStationSubwayline{})
	Db.AutoMigrate(&WorkingSchedule{})
	if !Db.Migrator().HasTable(&Account{}) {
		os.Exit(17)
	}
	if !Db.Migrator().HasTable(&SubwayStation{}) {
		os.Exit(18)
	}
	if !Db.Migrator().HasTable(&SubwayLine{}) {
		os.Exit(19)
	}
	if !Db.Migrator().HasTable(&Train{}) {
		os.Exit(20)
	}
	if !Db.Migrator().HasTable(&SubwayStationSubwayline{}) {
		os.Exit(21)
	}
}
