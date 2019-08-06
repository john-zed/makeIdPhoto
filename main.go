package main

import (
	. "../staff_remote/database"
)

func main() {

	defer SqlDB.Close()
	router := initRouter() //同一个包内可引用
	router.Run(":8102")
}
