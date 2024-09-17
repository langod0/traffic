package main

import (
	. "main/binary"
	. "main/request"
)

func main() {
	ConfigInit()
	LogInit()
	if Setting.UseRedis {
		cmd := RedisInit()
		defer func() {
			if err := StopRedis(cmd); err != nil {
				DebugLog.Printf("Error stopping Redis: %v\n", err)
			}
		}()
	}
	DBInit()
	ServeInit()
	//defer F.Close()
	func() {
		dr := Account{Name: "林禄创", StaffId: "R209980"}
		Db.Model(&Account{}).First(&dr)

		tr := Train{ID: "1"}
		Db.Model(&tr).Association("Drivers").Append(&dr)
	}()
	err := R.Run(":7234")
	if err != nil {
		println("Error: ", err)
		return
	}
	println("YES")
}
