package storage

import "sync"

const NO_ID = 0

type InMemoryStorage[T Item] struct {
	sequence int
	items    map[int]T
	mutex    sync.RWMutex
}

func (s *InMemoryStorage[T]) Init() {
	// nothing to init
}

func NewInMemoryStorage[T Item]() *InMemoryStorage[T] {
	return &InMemoryStorage[T]{
		sequence: NO_ID,
		items:    map[int]T{},
		mutex:    sync.RWMutex{},
	}
}

func (s *InMemoryStorage[T]) Get(id int) T {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.items[id]
}

func (s *InMemoryStorage[T]) GetAll() []T {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	result := make([]T, 0, len(s.items))
	for _, item := range s.items {
		result = append(result, item)
	}
	return result
}

func (s *InMemoryStorage[T]) Put(item T) T {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if item.Identity() == NO_ID {
		s.sequence++
		item.SetIdentity(s.sequence)
	}
	s.items[item.Identity()] = item
	return item
}

func (s *InMemoryStorage[T]) Remove(id int) T {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	removedItem := s.items[id]
	delete(s.items, id)
	return removedItem
}
