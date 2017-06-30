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
}

func main() {
	PrintAuthorInfo()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	for {
		_, err := listener.Accept()
		if err != nil {
			log.Panic(err)
		}

	}
}
