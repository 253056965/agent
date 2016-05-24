package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/agent/net/core"
	"github.com/agent/net/example"
)

func main() {
	go func() {
		ipprot := "127.0.0.1:6060"
		fmt.Printf("debug访问:http://%s/debug/pprof/", ipprot)
		http.ListenAndServe(ipprot, nil)

	}()
	iohandler := example.NewTelnetHandler()
	protobufDecoder := example.NewTelnetProtobufDecoder()
	protobufEncoder := example.NewTelnetProtobufEncoder()
	protocolCodecFactory := example.NewTelnetProtocolCodecFactory(protobufDecoder, protobufEncoder)
	server := core.NewServer2("0.0.0.0", 8889, iohandler, protocolCodecFactory)
	err := server.StartServer()
	if err != nil {
		fmt.Println("程序异常即将推出")
	}
}
