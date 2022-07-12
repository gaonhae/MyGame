package main

import (
	"fmt"
	"time"

	MyGameFuncs "github.com/gaonhae/Go/MyGame/MyGameFuncs"
)

func main() {
	MyGameFuncs.CleanDB()
	fmt.Println("작업 종료")
	time.Sleep(time.Second * 30)
}