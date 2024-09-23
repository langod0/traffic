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
	defer F.Close()
	err := R.Run(":7234")
	if err != nil {
		println("Error: ", err)
		return
	}
	println("YES")
}
