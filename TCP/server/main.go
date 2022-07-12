package server

import (
	"io"
	"log"
	"net"
)

func main() {
	//통신 방식과 포트값을 전달해 리스너 객체 생성
	listener, error := net.Listen("tcp", ":10000")

	//예외처리
	if nil != error {
		log.Fatalf("fail to bind address to 5032; err: %v", error)
	}
	defer listener.Close()

	log.Printf("## 프로그램 시작")
	//메인 루프
	for {
		//연결 대기
		connection, error := listener.Accept()
		//연결 실패
		if nil != error {
			log.Printf("연결 실패: %v", error)
			continue
		} else {
			log.Printf("클라이언트 연결: %v", connection.RemoteAddr())
		}

		//각 연결에 대한 처리를 고루틴으로 실행
		go func() {
			buffer := make([]byte, 1000) //버퍼

			//다 받을때까지 반복하며 읽음
			for {
				//입력
				count, error := connection.Read(buffer)
				if nil != error {
					//입력이 종료되면 중료
					if io.EOF == error {
						log.Printf("연결 종료: %v", connection.RemoteAddr().String())
					} else {
						log.Printf("수신 실패: %v", error)
					}
					return
				}
				if 0 < count {
					//받아온 길이만큼 슬라이스를 잘라서 출력
					data := buffer[:count]
					log.Println(string(data))
				}
			}
		}()
	}
}