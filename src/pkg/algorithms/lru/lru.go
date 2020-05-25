package lru

import (
	"net/http"
	"container/list"

)


type T interface{}
// Element has key and value
type Element struct {
	ID T
	val T
}
type Map map[T]*list.Element

type LRU struct {
	cache Map
	link  *list.List
	size  int
}

// New returns new LRU cache of the given size
func New(size int) *LRU {
	return &LRU{
		cache: make(Map),
		link:  list.New(),
		size:  size,
	}

}

// Set sets the given key,value pair in the cache.
// Returns true if new value is set in the cache.
func (l *LRU) Set(ID T, val T) {
	el := Element{ID, val}
//	defer l.dump()

	if e, ok := l.cache[ID]; ok {
		e.Value = el
		l.link.MoveToFront(e)
		return
	}
	if len(l.cache) < l.size {
		l.cache[ID] = l.link.PushFront(el)
		return

	}
	e := l.link.Back()
	delete(l.cache, e.Value.(Element).ID)
	l.link.Remove(e)
	l.cache[ID] = l.link.PushFront(el)
	return
}

// Get gets a value from the cache
func (l *LRU) Get(ID T,w http.ResponseWriter, r *http.Request) (T, bool) {
//	defer l.dump(w,r)
	val, ok := l.cache[ID]
	if !ok {
		return nil, ok
	}
	l.link.MoveToFront(val)
	el := val.Value.(Element)
	return el.val, true
}

// dump dumps cache content for debugging
func (l *LRU) Dump() ([]T) {
	e := l.link.Back()
	var state []T

	for i := 0; i < l.size; i++ {
		var ID T
		if e != nil {
			ID = e.Value.(Element).ID
			e = e.Prev()
		}
	if ID != nil{
state = append(state,ID);
	}
}
return state
}
