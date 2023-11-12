package cache

import (
	"context"

	"github.com/graph-gophers/dataloader/v7"
	lru "github.com/hashicorp/golang-lru"
)

type Cache[K comparable, V any] struct {
	archCache *lru.ARCCache
}

func NewCache[K comparable, V any](size int) (*Cache[K, V], error) {
	arcCache, err := lru.NewARC(size)
	if err != nil {
		return nil, err
	}

	return &Cache[K, V]{archCache: arcCache}, nil
}

func (s *Cache[K, V]) Get(_ context.Context, key K) (dataloader.Thunk[V], bool) {
	value, ok := s.archCache.Get(key)
	if ok {
		return value.(dataloader.Thunk[V]), ok
	}

	return nil, ok
}

func (s *Cache[K, V]) Set(_ context.Context, key K, value dataloader.Thunk[V]) {
	s.archCache.Add(key, value)
}

func (s *Cache[K, V]) Delete(_ context.Context, key K) bool {
	if s.archCache.Contains(key) {
		s.archCache.Remove(key)

		return true
	}

	return false
}

func (s *Cache[K, V]) Clear() {
	s.archCache.Purge()
}
