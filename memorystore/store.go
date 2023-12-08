package memorystore

import (
	"sync"

	"github.com/truthatt11/go-tus"
)

// MemoryStore implements an in-memory Store.
type MemoryStore struct {
	m map[string]string

	mutex *sync.Mutex
}

// NewMemoryStore creates a new MemoryStore.
func NewMemoryStore() (tus.Store, error) {
	return &MemoryStore{
		m:     make(map[string]string),
		mutex: &sync.Mutex{},
	}, nil
}

func (s *MemoryStore) Get(fingerprint string) (string, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	url, ok := s.m[fingerprint]
	return url, ok
}

func (s *MemoryStore) Set(fingerprint, url string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.m[fingerprint] = url
}

func (s *MemoryStore) Delete(fingerprint string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.m, fingerprint)
}

func (s *MemoryStore) Close() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for k := range s.m {
		delete(s.m, k)
	}
}
