package subscriber

import (
	"sync"

	"github.com/google/uuid"

	"github.com/redhat-cne/sdk-go/pkg/subscriber"
)

// Store  defines subscribers connection store struct
type Store struct {
	sync.RWMutex
	// Store stores subscribers in a map
	Store map[uuid.UUID]*subscriber.Subscriber
}

// Set is a wrapper for setting the value of a key in the underlying map
func (ss *Store) Set(clientID uuid.UUID, val subscriber.Subscriber) {
	ss.Lock()
	defer ss.Unlock()
	ss.Store[clientID] = &val
}

// Delete ... delete from store
func (ss *Store) Delete(clientID uuid.UUID) {
	ss.Lock()
	defer ss.Unlock()
	delete(ss.Store, clientID)
}
