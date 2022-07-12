package main

import (
	"fmt"
	"time"

	MyGameFuncs "github.com/gaonhae/Go/MyGame/MyGameFuncs"
)

func main() {
	fmt.Println("이 프로그램에는 266357개의 단어가 내장되어 있습니다.")
	MyGameFuncs.ReadDB()
	fmt.Println("작업 종료")
	time.Sleep(time.Second * 30)
}