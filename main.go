package main

import (
	"fmt"
	"log"
	"net"
)

// 打印作者信息
func PrintAuthorInfo() {
	logo := `
   _____       _                                 _           
  / ____|     | |                               | |          
 | (___   __ _| | __ _ _ __ ___   __ _ _ __   __| | ___ _ __` + "\n" +
		"  \\___ \\ / _` | |/ _` | '_ ` _ \\ / _` | '_ \\ / _` |/ _ \\ '__|" + "\n" +
		`  ____) | (_| | | (_| | | | | | | (_| | | | | (_| |  __/ |   
 |_____/ \__,_|_|\__,_|_| |_| |_|\__,_|_| |_|\__,_|\___|_|   
` + "\n"
	fmt.Print(logo)
	fmt.Println("Project Address: https://github.com/salamander-mh/SalamanderHttpProxy")
}

func main() {
	PrintAuthorInfo()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	for {
		client, err := listener.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleClientRequest(client)
	}
}

func handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()

	var buf [1024]byte
	n, err := client.Read(buf[:])
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("recv msg:", string(buf[0:n]))
}
