package dto

type MessageOut struct {
	Entity any
	User   string
	Action string
}

func NewMessageOut(entity any, user, action string) *MessageOut {
	return &MessageOut{
		Entity: entity,
		User:   user,
		Action: action,
	}
}
