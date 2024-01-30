package queue

type Producer interface {
	Send(queue string, obj []byte) error
}
