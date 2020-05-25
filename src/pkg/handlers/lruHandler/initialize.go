package lruHandler

import(
	"cache-routine/src/pkg/algorithms/lru"
	"fmt"
)

var Cache *lru.LRU

func Initialize(){
	Cache = lru.New(1);
  Cache.Set("empty","empty")
	fmt.Printf("LRU cache has been created \n")

}

func SetNewSize(size int){
	Cache = lru.New(size);
  Cache.Set(" size changed ","empty")
}
