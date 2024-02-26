package repo

import (
	"context"
	"sync"
	"time"
)

type InMemoryCache struct {
	data  map[string]time.Time
	mutex sync.Mutex
}

// NewInMemoryCache creates a new InMemoryCache.
func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{data: make(map[string]time.Time)}
}

// GetLastSentTime retrieves the last sent time for the given key from the in-memory map.
func (s *InMemoryCache) GetLastSentTime(ctx context.Context, key string) (time.Time, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	value, ok := s.data[key]
	if !ok {
		return time.Time{}, nil // Key not found
	}

	return value, nil
}

// SetLastSentTime stores the last sent time for the given key in the in-memory map.
func (s *InMemoryCache) SetLastSentTime(ctx context.Context, key string, value time.Time) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.data[key] = value
	return nil
}
