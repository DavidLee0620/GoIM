package chat

import (
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type inMessage struct {
	ToID string `json:"toID"`
	Msg  string `json:"msg"`
}
type outMessage struct {
	From User   `json:"from"`
	Msg  string `json:"msg"`
}

type User struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	LastPing time.Time       `json:"lastping"`
	LastPong time.Time       `json:"lastpong"`
	Conn     *websocket.Conn `json:"-"`
}

type Connection struct {
	Conn     *websocket.Conn
	LastPing time.Time
	LastPong time.Time
}
type busMessage struct {
	CapID    uuid.UUID `json:"capID"`
	FromID   string    `json:"from"`
	FromName string    `json:"fromName"`
	ToID     string    `json:"toID"`
	Msg      string    `json:"msg"`
}
