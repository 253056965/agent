package main

import (
	"fmt"
	"reflect"
	"github.com/agent/net/core"
	"github.com/agent/net/example"
)

func main() {
	//test("ddddddddddddddd");

	iohandler := example.NewTelnetHandler()
	protobufDecoder := example.NewTelnetProtobufDecoder()
	protobufEncoder := example.NewTelnetProtobufEncoder()
	protocolCodecFactory := example.NewTelnetProtocolCodecFactory(protobufDecoder, protobufEncoder)
	server := core.NewServer2("0.0.0.0",8889,  iohandler,protocolCodecFactory)
	err := server.StartServer()
	if err != nil {
		fmt.Println("程序异常即将推出")
	}
	message := []byte(fmt.Sprintf("%v", "obj"))
	test(message)
}

func test(mes interface{}) {
	str := reflect.ValueOf(mes).Bytes()
	fmt.Println(string(str))
}
