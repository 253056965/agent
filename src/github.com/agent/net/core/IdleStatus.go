package core

//IdleStatus 超时状态
type IdleStatus struct{
    stateStr string;
}

func  (idl *IdleStatus) String()string{
    return idl.stateStr+"";
}

func newIdleStatus( statestr string)  *IdleStatus{
    
    return &IdleStatus{statestr};
}
//READER_IDLE 读超时
 var READER_IDLE *IdleStatus = newIdleStatus("reader idle");
//WRITER_IDLE 写超时
var  WRITER_IDLE *IdleStatus= newIdleStatus("writer idle");
//BOTH_IDLE 两者都超时
var  BOTH_IDLE *IdleStatus= newIdleStatus("both idle")

