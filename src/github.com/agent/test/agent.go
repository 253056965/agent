package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/agent/net/core"
	"github.com/agent/net/example"
)

func main() {
	//jobrunner.Start()

	//cront := cron.New()
	// log := util.NewLogger()

	// mainCron := cron.New()
	// mainCron.Start()
	// mainCron.Schedule(cron.Every(time.Second*10), newJobs(func() {
	// 	log.Infoln("ddddddddddd")
	// }))
	// jobrunner.Schedule("@every 5s", ReminderEmails{})
	// jobrunner.Schedule("@every 5s", ReminderEmails{})
	// jobrunner.Schedule("@every 5s", ReminderEmails{})

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
