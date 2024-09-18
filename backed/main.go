package main

import (
	"fmt"
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
		var Ac []Account
		Db.Model(&Account{}).Limit(10).Find(&Ac)
		var Tr Train
		Db.Preload("Drivers").Where("id = ?", "111020").First(&Tr)
		fmt.Println(Tr.ID)
		for _, v := range Tr.Drivers {
			fmt.Println(v.StaffId)
		}
		//Db.Model(&Train{}).Limit(3).Find(&Tr)
		//len := len(Tr)
		//for i := 0; i < 10; i++ {
		//	Ac[i].TrainId = Tr[rand.Intn(len)].ID
		//	fmt.Println(Ac[i].TrainId)
		//	Db.Save(&Ac[i])
		//}

		//Db.Model(&Account{}).Save(&Ac)
	}()
	return
	err := R.Run(":7234")
	if err != nil {
		println("Error: ", err)
		return
	}
	println("YES")
}
