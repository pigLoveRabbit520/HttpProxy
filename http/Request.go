package http

import (
	"log"
	"net"
//	"strings"
	"fmt"
)

const (
	BufferSize = 1024
	LF         = "\n"
	CRLF       = "\r\n"
	TWOCRLF    = "\r\n\r\n" // 两个CRLF
	TWOLF      = "\n"
)

// 请求类
type Request struct {
	URL      string
	Method   string
	Protocal string
	Headers
}

func HandleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()
	//allBytesNum := 0
	//allBytes := []byte{}
	request := Request{}
	request.Headers = make(map[string]string)

	for {
		var buf [BufferSize]byte
		readBytesNum, err := client.Read(buf[:])
		if err != nil {
			log.Println("读取字节失败:" + err.Error())
			return
		}
		str := string(buf[0:readBytesNum])
		fmt.Println(str)
		break
		//allBytesNum += readBytesNum
		//allBytes = append(allBytes, buf[:]...)
		//tmpStr := string(allBytes)
		//if strings.Contains(TWOCRLF, tmpStr) {
		//	tmpArr := strings.Split(tmpStr, TWOCRLF)
		//	if(!request.ExtractHeaders(tmpArr[0], CRLF)) {
		//		return
		//	}
		//	break
		//} else if strings.Contains(TWOLF, tmpStr) {
		//	tmpArr := strings.Split(tmpStr, TWOLF)
		//	if(!request.ExtractHeaders(tmpArr[0], LF)) {
		//		return
		//	}
		//	break
		//}
	}

	//fmt.Println(request)

}
