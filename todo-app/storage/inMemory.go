package storage

import "sync"

func NewInMemoryStorage[T Item]() *InMemoryStorage[T] {
	return &InMemoryStorage[T]{
		items: map[int]T{},
		mutex: sync.RWMutex{},
	}
}

type InMemoryStorage[T Item] struct {
	items map[int]T
	mutex sync.RWMutex
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
	prev := s.items[item.Identity()]
	s.items[item.Identity()] = item
	return prev
}

func (s *InMemoryStorage[T]) Remove(id int) T {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	removedItem := s.items[id]
	delete(s.items, id)
	return removedItem
}
