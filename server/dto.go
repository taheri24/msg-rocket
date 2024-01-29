package server

import "encoding/json"

type MsgKind string

const (
	mkSendMessage MsgKind = "send"
	mkSyncMessage MsgKind = "sync"
	mkPullMessage MsgKind = "pull"
	mkPullReady   MsgKind = "r:pull"
	mkConfrimPull MsgKind = "ok:pull"
)

type MsgDTO struct {
	ID         int64
	Kind       MsgKind
	From, To   string
	TimeStamp  int64
	BodyLength int32
	Body       json.RawMessage
}
