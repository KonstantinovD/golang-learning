package message

const (
	Undefined Level = iota
	Firing
	Pending
	Clear
)

type severity struct {
	Level Level
}

type Message struct {
	Value  string
	Sevrt  severity
	hidden string
}

type Level int64

func OpenMessage() Message {
	return Message{
		Value:  "Open msg",
		Sevrt:  severity{Level: Clear},
		hidden: "012"}
}

func ClosedMessage() Message {
	return Message{
		Value:  "Closed msg",
		Sevrt:  severity{Level: 12}, // but no restrictions for 'enums'
		hidden: "054"}
}
