package eventchannel

// Subscriber - подписчик
type Subscriber interface {
	OnReceive(msg string)
	GetID() string
}

// SubscriberDefault -  вспомогательная структура
type SubscriberDefault struct{}

// OnReceive -
func (SubscriberDefault) OnReceive(string) {
	panic("not implemented")
}

// GetID -
func (SubscriberDefault) GetID(string) {
	panic("not implemented")
}
