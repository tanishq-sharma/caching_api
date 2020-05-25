package datastructures
import(
	"container/list"
)
type T interface{}

type Element struct {
	ID T
	Value T
}

type Object struct {
Value T
}


type Map map[T]*list.Element

type Lru struct {
	cache Map
	link  *list.List
	size  int
}

type Node_json struct {
	ID T `json:"ID"`
	Value T `json:"Value"`
}

type State_json struct {

	State []T `json: "State"`
}

type Size_json struct {

	Size T `json: "Size"`
}
