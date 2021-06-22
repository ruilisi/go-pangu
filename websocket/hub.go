package websocket

import (
	"sync"

	jsoniter "github.com/json-iterator/go"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	rooms *sync.Map

	// Inbound messages from the clients.
	broadcast chan *message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func newHub() *Hub {
	h := &Hub{
		broadcast:  make(chan *message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		rooms:      &sync.Map{},
	}
	go h.run()
	return h
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			conns, _ := h.rooms.LoadOrStore(client.room, &sync.Map{})
			conns.(*sync.Map).Store(client, true)
		case client := <-h.unregister:
			_conns, ok := h.rooms.Load(client.room)
			if ok {
				conns := _conns.(*sync.Map)
				if _, ok := conns.Load(client); ok {
					conns.Delete(client)
					close(client.send)
					if mapLen(conns) == 0 {
						h.rooms.Delete(client.room)
					}
				}
			}
		case message := <-h.broadcast:
			h.Broadcast(message.sender.room, message)
		}
	}
}

// Broadcast send broadcast message to
func (h *Hub) Broadcast(room string, msg *message) {
	if _conns, ok := h.rooms.Load(room); ok {
		conns := _conns.(*sync.Map)
		conns.Range(func(k, _ interface{}) bool {
			client := k.(*Client)
			select {
			case client.send <- msg:
			default:
				close(client.send)
				conns.Delete(client)
				if mapLen(conns) == 0 {
					h.rooms.Delete(room)
				}
			}
			return true
		})
	}
}

// SysBroadcastJSON broadcast JSON message without a user
func (h *Hub) SysBroadcastJSON(room string, msg interface{}) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	if data, err := json.Marshal(msg); err == nil {
		h.Broadcast(room, &message{nil, data})
	}
}

func mapLen(m *sync.Map) int {
	length := 0
	m.Range(func(_, _ interface{}) bool {
		length++
		return true
	})
	return length
}

// GetHub returns the instance of hub
func GetHub() *Hub {
	return hub
}
