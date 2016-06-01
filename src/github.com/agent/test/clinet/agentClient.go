package main

import (
	"fmt"
	"time"

	"github.com/agent/net/core"
	"github.com/agent/net/example"
)

func main() {

	core.StartSessionJob()
	iohandler := example.NewTelnetHandler()
	protobufDecoder := example.NewTelnetProtobufDecoder()
	protobufEncoder := example.NewTelnetProtobufEncoder()
	protocolCodecFactory := example.NewTelnetProtocolCodecFactory(protobufDecoder, protobufEncoder)
	client := core.NewClient2("172.16.10.65", 8889, iohandler, protocolCodecFactory)
	err := client.Connection()
	if err != nil {
		fmt.Println("程序异常即将推出")
	}
	session := client.GetSession()
	if session != nil {
		session.Write("ddddddd\r\n")
	}
	for {
		time.Sleep(time.Second * 30)
	}

}
