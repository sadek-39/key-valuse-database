package main

import (
	"fmt"
	"github.com/sadek-39/key-value-database/handlers"
	"github.com/sadek-39/key-value-database/storage"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":6379")

	if err != nil {
		fmt.Println(err)
		return
	}

	con, err := l.Accept()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer con.Close()

	for {
		buf := make([]byte, 1024)

		storage.LoadDataFromFile()
		i, err := con.Read(buf)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error from reading : ", err.Error())
			os.Exit(1)
		}
		command := strings.TrimSpace(string(buf[:i]))
		fmt.Println(command)

		for k, handlers := range handlers.Handlers {
			if strings.HasPrefix(command, k) {
				args := strings.TrimPrefix(command, k)
				handlers(args)
			}
		}

		//if handler, exists := handlers.Handlers[command]; exists {
		//	handler(command)
		//} else {
		//	fmt.Println("unknown command")
		//	os.Exit(1)
		//}

		//con.Write([]byte("+OK\r\n"))
	}
}
