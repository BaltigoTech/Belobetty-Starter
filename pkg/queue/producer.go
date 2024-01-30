package queue

type Producer interface {
	SendMessage(obj []byte) error
}
