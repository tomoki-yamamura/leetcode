package main

import (

	"io"
	"net"

	"os"
)

// コネクションを裁く
// func handlerConn(c net.Conn) {
// 	defer c.Close()
// 	for {
// 		// リクエストを受け付けたらサーバー側に「response from server」を返す
// 		_, err := io.WriteString(c, "response from server\n")
// 		if err != nil {
// 			return
// 		}
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	// tcp, かつ、localhost:8080でリクエストを受け取る
// 	listner, err := net.Listen("tcp", "localhost:8080")
// 	fmt.Println("Server Start")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//　ループで回して複数リクエストを受け取れるようにする
// 	for {
// 		// 接続
// 		conn, err := listner.Accept()
// 		fmt.Println("Connect")

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		go handlerConn(conn)
// 	}
// }

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("hogehoge"))
	io.Copy(os.Stdout, conn)
}
