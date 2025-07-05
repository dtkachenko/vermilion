package memory

import (
	"sync"

	"github.com/dtkachenko/vermilion/internal/storage"
)

type MemoryStorage struct {
	mu   sync.Mutex
	data []storage.PodInfo
}

func New() *MemoryStorage {
	return &MemoryStorage{}
}

func (m *MemoryStorage) Save(pod storage.PodInfo) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data = append(m.data, pod)
	return nil
}

func (m *MemoryStorage) GetAll() ([]storage.PodInfo, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]storage.PodInfo(nil), m.data...), nil
}
