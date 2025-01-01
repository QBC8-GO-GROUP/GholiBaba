package broadcast

import (
	"context"
)

type Server interface {
	Subscribe() <-chan string
	CancelSubscribe(<-chan string)
}

type broadcastServer struct {
	source         <-chan string
	listeners      []chan string
	addListener    chan chan string
	removeListener chan (<-chan string)
}

func (b *broadcastServer) Subscribe() <-chan string {
	newListener := make(chan string)
	b.addListener <- newListener
	return newListener
}

func (b *broadcastServer) CancelSubscribe(ch <-chan string) {
	b.removeListener <- ch
}

func NewBroadcastServer(ctx context.Context, source <-chan string) Server {
	service := &broadcastServer{
		source:         source,
		listeners:      make([]chan string, 0),
		addListener:    make(chan chan string),
		removeListener: make(chan (<-chan string)),
	}
	go service.serve(ctx)
	return service
}

func (b *broadcastServer) serve(ctx context.Context) {
	defer func() {
		for _, listener := range b.listeners {
			if listener != nil {
				close(listener)
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case newListener := <-b.addListener:
			b.listeners = append(b.listeners, newListener)
		case listenerToRemove := <-b.removeListener:
			for i, ch := range b.listeners {
				if ch == listenerToRemove {
					b.listeners[i] = b.listeners[len(b.listeners)-1]
					b.listeners = b.listeners[:len(b.listeners)-1]
					close(ch)
					break
				}
			}
		case val, ok := <-b.source:
			if !ok {
				return
			}
			for _, listener := range b.listeners {
				if listener != nil {
					select {
					case listener <- val:
					case <-ctx.Done():
						return
					}

				}
			}
		}
	}
}
