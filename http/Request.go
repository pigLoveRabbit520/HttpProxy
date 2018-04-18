package http

import (
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"
)

const (
	BufferSize = 1024
	CRLF       = "\r\n"
	TWOCRLF    = "\r\n\r\n" // 两个CRLF
)

// 请求类
type Request struct {
	Method   string
	Path     string
	Host     string
	Protocol string // http协议版本
	Headers
}

func HandleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()

	request := Request{}
	request.Headers = make(map[string]string)
	var bufferStr = ""
	var headerEndPos int

	for {
		var buf [BufferSize]byte
		readBytesNum, err := client.Read(buf[:])
		if err != nil {
			log.Println("读取字节失败:" + err.Error())
			return
		}
		bufferStr += string(buf[0:readBytesNum])
		if headerEndPos = strings.Index(bufferStr, TWOCRLF); headerEndPos >= 0 {
			break
		}
	}
	headersStr := bufferStr[:headerEndPos+1]
	fmt.Println("recevie header: " + headersStr)
	lines := strings.Split(headersStr, CRLF)
	lineCount := len(lines)
	for i := 0; i < lineCount; i++ {
		var line = lines[i]
		if i == 0 {
			// 第一行为请求行
			strs := regexp.MustCompile("\\s").Split(line, -1)
			request.Method = strs[0]
			request.Path = strs[1]
			request.Protocol = strs[2]
		} else {
			twoParts := strings.Split(line, ":")
			headerName := strings.TrimSpace(twoParts[0])
			headerValue := strings.TrimSpace(twoParts[1])
			request.Headers[headerName] = headerValue
		}
	}
	fmt.Printf("your request is %+v\n", request)
}
