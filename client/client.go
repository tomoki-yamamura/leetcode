package main
 
import (
	"io"
	"log"
	"net"
	"os"
)
 
func response(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
 
func main() {
	// tcp, localhost:8080でリクエストを送る
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
 
	// connからレスポンスを標準出力にだす
	response(os.Stdout, conn)
	defer conn.Close()
}
 