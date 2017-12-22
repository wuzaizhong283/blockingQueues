package blockingQueues

import (
	"container/list"
	"sync"
)

type LinkedListStore struct {
	store *list.List
}

func NewLinkedListStore() *LinkedListStore {
	return &LinkedListStore{
		store: list.New(),
	}
}

func (s *LinkedListStore) Set(value interface{}, pos uint64) {
	s.store.PushBack(value)
}

func (s *LinkedListStore) Get(pos uint64) interface{} {
	return s.store.Front()
}

func (s *LinkedListStore) Remove(pos uint64) interface{} {
	var item = s.store.Remove(s.store.Front())
	return item
}

func (s LinkedListStore) Size() uint64 {
	return uint64(s.store.Len())
}

// Creates an BlockingQueue backed by an LinkedList with the given (fixed) capacity
// returns an error if the capacity is less than 1
func NewLinkedBlockingQueue(capacity uint64) (*BlockingQueue, error) {
	if capacity < 1 {
		return nil, ErrorCapacity
	}

	lock := new(sync.RWMutex)

	return &BlockingQueue{
		lock:     lock,
		notEmpty: sync.NewCond(lock.RLocker()),
		notFull:  sync.NewCond(lock.RLocker()),
		count:    uint64(0),
		store:    NewLinkedListStore(),
	}, nil
}