package pongs

import (
	"context"
	"sync"
	"time"

	"github.com/DavidLee0620/GoIM/chat/app/sdk/chat"
	"github.com/DavidLee0620/GoIM/chat/foundation/logger"
)

type Pongs struct {
	log   *logger.Logger
	users map[string]time.Time
	mu    sync.RWMutex
}

func New(log *logger.Logger) *Pongs {
	p := Pongs{
		log:   log,
		users: make(map[string]time.Time),
	}
	return &p
}

func (p *Pongs) Add(ctx context.Context, usr chat.User) {

	h := func(appData string) error {
		p.mu.Lock()
		defer p.mu.Unlock()
		p.users[usr.ID] = time.Now()
		return nil
	}
	usr.Conn.SetPongHandler(h)
}

func (p *Pongs) Vaildata(usr chat.User, maxWait time.Duration) bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	t := p.users[usr.ID]
	return time.Since(t) <= maxWait
}
