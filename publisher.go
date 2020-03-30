package eventchannel

import (
	"fmt"
)

// Channels - каналы
type Channels map[string]*Channel

// Publisher - структура публикации
type Publisher struct {
	channels Channels
}

// NewPublisher - Новая публикация
func NewPublisher() *Publisher {
	return &Publisher{
		channels: Channels{},
	}
}

// AddChannel - создаем канал
func (p *Publisher) AddChannel(name string, channel *Channel) {
	p.channels[name] = channel
}

// DeleteChannel - удаление канала по имени
func (p *Publisher) DeleteChannel(name string, сhannels ...string) error {
	if _, ok := p.channels[name]; ok {
		delete(p.channels, name)
	}
	return nil
}

// SendAllChannels - передаем привет всем существующим каналам
func (p *Publisher) SendAllChannels(msg string) {
	if len(p.channels) != 0 {
		for key := range p.channels {
			ch := p.channels[key]
			ch.Send(msg)
		}
	} else {
		fmt.Printf("%s", "Нет каналов")
	}
}

// Send - передаем публикацию в каналы
func (p *Publisher) Send(msg string, сhannels ...string) error {
	for _, ch := range сhannels {
		channel, ok := p.channels[ch]
		if !ok {
			return fmt.Errorf("channel %s cant't be found", ch)
		}
		channel.Send(msg)
	}
	return nil
}
