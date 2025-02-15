package utils

import "sync"

type Singleton[T any] struct {
	once     sync.Once
	instance *T
}

func (s *Singleton[T]) GetInstance() *T {
	s.once.Do(func() {
		s.instance = new(T)
	})
	return s.instance
}

func NewSingleton[T any]() *Singleton[T] {
	return &Singleton[T]{}
}
