package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

type Connection struct {
	Subject string
	Host    string
	Port    int
	nc      *nats.Conn
	Sub     *nats.Subscription
	Ch      chan string
}

func NewConnection(host string, port int, subject string) *Connection {
	return &Connection{
		Subject: subject,
		Host:    host,
		Port:    port,
	}
}

func (c *Connection) Connect() error {
	ch := make(chan string)
	c.Ch = ch

	nc, err := nats.Connect(fmt.Sprintf("nats://%s:%s", c.Host, c.Port))
	if err != nil {
		return err
	}
	c.nc = nc

	sub, err := nc.Subscribe(c.Subject, func(msg *nats.Msg) {
		c.Ch <- string(msg.Data)
	})

	if err != nil {
		return err
	}
	c.Sub = sub

	return nil
}

func (c *Connection) MustConnect() {
	if err := c.Connect(); err != nil {
		panic(err)
	}
}

func (c *Connection) Close() {
	_ = c.Sub.Unsubscribe()
	c.nc.Close()
}
