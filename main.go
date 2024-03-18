package main

import (
	"fmt"
	"gocaster/cmd/rtmpserver"
	"net"
	"time"
)

func main() {
	println(uint32(time.Now().Unix()))
	// TCP 서버 시작
	listener, err := net.Listen("tcp", ":1935") // RTMP 기본 포트
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on port 1935")

	for {
		// 클라이언트 연결 대기
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("Client connected:", conn.RemoteAddr())

		// RTMP 핸드쉐이크 수행
		if err := rtmpserver.PerformRTMPHandshake(conn); err != nil {
			fmt.Println("Error performing RTMP handshake:", err)
			conn.Close()
			continue
		}

		// 핸드셰이킹이 완료되었음을 출력
		fmt.Println("RTMP handshake completed successfully.")

		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("Error reading:", err)
				break
			}
			fmt.Println("Read:", string(buf[:n]))
		}
	}
}
