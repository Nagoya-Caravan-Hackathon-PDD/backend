package types

type MessageType string

const (
	FlagStart  MessageType = "start"
	FlagTurn   MessageType = "turn"
	FlagResult MessageType = "result"
	FlagEnd    MessageType = "end"
)

type WSRequest struct {
	MessageType MessageType `json:"message_type"`
	Time        int         `json:"time"`
}
