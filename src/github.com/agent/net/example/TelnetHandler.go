package example

import (
	"reflect"
	"github.com/agent/net/core"
	"github.com/agent/util"
)

type TelnetHandler struct {
	log * util.Logger
}

func NewTelnetHandler() *TelnetHandler {
	
	return &TelnetHandler{util.NewLogger()}
}

func (t *TelnetHandler) SessionOpened(session *core.IOSession) {

}

func (t *TelnetHandler) SessionCreated(session *core.IOSession) {
	t.log.Infof("socket 远程地址是:%s", session.GetRemoteAddr());
}

func (t *TelnetHandler) SessionClosed(session *core.IOSession) {
	t.log.Infoln("socket 关闭了");
}

func (t *TelnetHandler) SessionIdle(session *core.IOSession, idle *core.IdleStatus) {

}

func (t *TelnetHandler) MessageSent(session *core.IOSession, message interface{}) {

}
func (t *TelnetHandler) MessageReceived(session *core.IOSession, message interface{}) {

	str := reflect.ValueOf(message).String();
	t.log.Infof("客户端发送的消息是:%s",str);
	session.Write("我收到你的消息了:" + str+"\r\n")
}
