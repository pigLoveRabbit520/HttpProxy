package http

import (
	"fmt"
	"log"
	"net"
	"regexp"
	"strconv"
	"strings"
)

const (
	BufferSize               = 1024
	LF                       = "\n"
	CRLF                     = "\r\n"
	TWOCRLF                  = "\r\n\r\n" // 两个CRLF
	TWOLF                    = "\n\n"     // 两个LF
	HEADER_CONTENT_LENGTH    = "content-length"
	HEADER_TRANSFER_ENCODING = "transfer-encoding"
	EMPTY_CHUNK              = "0\r\n\r\n"
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

	var buf [BufferSize]byte
	request := Request{}
	request.Headers = make(map[string]string)
	readBuffer := make([]byte, 0)
	bufferStr := ""
	var headerEndPos int
	bodyStartPos := 0

	for {
		readBytesNum, err := client.Read(buf[:])
		if err != nil {
			log.Println("读取字节失败:" + err.Error())
			return
		}
		readBuffer = append(readBuffer, buf[:readBytesNum]...)
		bufferStr = string(readBuffer)
		if headerEndPos = strings.Index(bufferStr, TWOCRLF); headerEndPos >= 0 {
			bodyStartPos = headerEndPos + len(TWOCRLF)
			break
		} else if headerEndPos = strings.Index(bufferStr, TWOLF); headerEndPos >= 0 { // 兼容性
			bodyStartPos = headerEndPos + len(TWOLF)
			break
		}
	}
	startLineAndHeadersStr := bufferStr[:headerEndPos+1]
	lines := regexp.MustCompile("\r\n|\n").Split(startLineAndHeadersStr, -1)
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
			fmt.Println("header is illegal: " + line)
			if len(twoParts) != 2 {
				return
			}
			headerName := strings.ToLower(strings.TrimSpace(twoParts[0]))
			headerValue := strings.TrimSpace(twoParts[1])
			request.Headers[headerName] = headerValue
		}
	}

	if host, ok := request.Headers["host"]; !ok {
		fmt.Println("without host header")
		return
	} else {
		request.Host = host
	}
	// transfer-encoding和content-length互斥
	if transferEncoding, ok := request.Headers[HEADER_TRANSFER_ENCODING]; ok {
		if transferEncoding == "chunked" {
			bodyBuffer := make([]byte, 0)
			// 可能还没读到body部分
			if len(readBuffer)-1 > bodyStartPos {
				bodyBuffer = readBuffer[bodyStartPos:]
			}
			// 判断最后5个字节
			for len(bodyBuffer) < len(EMPTY_CHUNK) || string(bodyBuffer[len(bodyBuffer)-len(EMPTY_CHUNK):]) != EMPTY_CHUNK {
				readBytesNum, err := client.Read(buf[:])
				if err != nil {
					log.Println("读取字节失败:" + err.Error())
					return
				}
				bodyBuffer = append(bodyBuffer, buf[:readBytesNum]...)
			}
			readBuffer = append(readBuffer[:bodyStartPos], bodyBuffer...)
		} else {
			readBuffer = readBuffer[:bodyStartPos]
		}
	} else if contentLengthStr, ok := request.Headers[HEADER_CONTENT_LENGTH]; ok {
		contentLength, err := strconv.Atoi(contentLengthStr)
		if err != nil {
			fmt.Println("content-length is illegal: " + contentLengthStr)
			return
		}
		if contentLength > 0 {
			bodyBuffer := make([]byte, 0)
			// 可能还没读到body部分
			if len(readBuffer)-1 > bodyStartPos {
				bodyBuffer = readBuffer[bodyStartPos:]
			}
			for len(bodyBuffer) < contentLength {
				readBytesNum, err := client.Read(buf[:])
				if err != nil {
					log.Println("读取字节失败:" + err.Error())
					return
				}
				bodyBuffer = append(bodyBuffer, buf[:readBytesNum]...)
			}
			readBuffer = append(readBuffer[:bodyStartPos], bodyBuffer...)
		} else {
			readBuffer = readBuffer[:bodyStartPos]
		}
	} else {
		readBuffer = readBuffer[:bodyStartPos]
	}
}

func requestServer(host string, clientData []byte) ([]byte, error) {
	serverSocket, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	defer serverSocket.Close()

}
