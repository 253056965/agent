package example

import (
	"reflect"
	"time"

	"github.com/agent/net/core"
	"github.com/agent/util"
)

type TelnetHandler struct {
	log *util.Logger
}

func NewTelnetHandler() *TelnetHandler {

	return &TelnetHandler{util.NewLogger()}
}

func (t *TelnetHandler) SessionOpened(session *core.IOSession) {
	t.log.Infoln("session 被打开了:", session)
	// 设置读写超时设置
	session.SetIdleTime(time.Second * 30)
	session.Write("我第一次用程序给你发代码\r\n")
}

func (t *TelnetHandler) SessionCreated(session *core.IOSession) {
	t.log.Infof("socket 远程地址是:%s", session.GetRemoteAddr())
}

func (t *TelnetHandler) SessionClosed(session *core.IOSession) {
	t.log.Infoln("socket 关闭了")
}

func (t *TelnetHandler) SessionIdle(session *core.IOSession, idle *core.IdleStatus) {
	t.log.Errorln("超时被触发了", idle)
	session.Close()
}

func (t *TelnetHandler) MessageSent(session *core.IOSession, message interface{}) {

}
func (t *TelnetHandler) MessageReceived(session *core.IOSession, message interface{}) {

	str := reflect.ValueOf(message).String()
	t.log.Infof("客户端发送的消息是:%s", str)
	//session.Write("我收到你的消息了:" + str + "\r\n")
}
