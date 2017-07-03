package http

import (
	"log"
	"net"
//	"strings"
	"fmt"
	"strings"
	"io"
	"bytes"
	"net/url"
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

	var buf [BufferSize]byte
	readBytesNum, err := client.Read(buf[:])
	if err != nil {
		log.Println("读取字节失败:" + err.Error())
		return
	}
	str := string(buf[0:readBytesNum])
	fmt.Println("recevie msg: " + str)
	var method, host, address string
	fmt.Sscanf(string(buf[:bytes.IndexByte(buf[:], '\n')]), "%s%s", &method, &host)
	hostPortURL, err := url.Parse(host)
	if err != nil {
		log.Println(err)
		return
	}

	if hostPortURL.Opaque == "443" { // https访问
		address = hostPortURL.Scheme + ":443"
	} else { //http访问
		if strings.Index(hostPortURL.Host, ":") == -1 { // host不带端口， 默认80
			address = hostPortURL.Host + ":80"
		} else {
			address = hostPortURL.Host
		}
	}

	//获得了请求的host和port，就开始拨号吧
	server, err := net.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		return
	}
	if method == "CONNECT" {
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n")
	} else {
		server.Write(buf[:readBytesNum])
	}
	//进行转发
	go io.Copy(server, client)
	io.Copy(client, server)



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

	//fmt.Println(request)

}
