package eventchannel

import (
	"fmt"
)

// Subscribers - подписчики
type Subscribers map[string]Subscriber

// Channel - Каналы с подписчиками
type Channel struct {
	subscribers Subscribers
}

// NewChannel - Builder можно задать ограничение на кол-во пользователей, уровни прав и прочее
func NewChannel() *Channel {
	return &Channel{
		subscribers: Subscribers{},
	}
}

// Send - отсылаем сообщения
func (ch *Channel) Send(msg string) {
	for _, sub := range ch.subscribers {
		sub.OnReceive(msg)
	}
}

// Subscribe - подписываем на канал по id
func (ch *Channel) Subscribe(sub Subscriber) {
	id := sub.GetID()
	ch.subscribers[id] = sub
}

// UnSubscribe - отписываем от канала по id
func (ch *Channel) UnSubscribe(sub Subscriber) error {
	id := sub.GetID()
	if _, ok := ch.subscribers[id]; ok {
		delete(ch.subscribers, id)
		return nil
	}
	return fmt.Errorf("can't find user %s", id)
}

// UnSubscribeAll - отписываем от канала всех пользователей
func (ch *Channel) UnSubscribeAll() error {
	for _, sub := range ch.subscribers {
		return ch.UnSubscribe(sub)
	}
	return nil
}
