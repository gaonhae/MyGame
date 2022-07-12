package client

import (
	"fmt"
	"log"
	"net"
)

func main(){
	connection, error := net.Dial("tcp", "14.37.169.160:10000")
	word := ""
	
	defer connection.Close()
	
	
	for{
		fmt.Scanln(&word)
		if nil != error {
			log.Printf("접속 실패")
		} else {
			_, error := connection.Write([]byte(word)) 
			if nil == error {
				log.Printf("전송 성공")
			}
		}
	}
}