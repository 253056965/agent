package core



type IoHandler interface {
	SessionCreated(session *IOSession)
	SessionOpened(session *IOSession)
	SessionClosed(session *IOSession)
	SessionIdle(session *IOSession, status *IdleStatus)
	//func   ExceptionCaught(session IoSession, Throwable cause) ;
     MessageReceived(session *IOSession ,  message interface{}) ;
     MessageSent(session *IOSession , message interface{}) ;
}
