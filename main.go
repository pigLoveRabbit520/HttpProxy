package main

import (
	"fmt"
	"log"
	"net"
	salamanderHttp "github.com/SalamanderHttpProxy/http"
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
		go salamanderHttp.HandleClientRequest(client)
	}
}