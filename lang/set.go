package lang

type _void struct{}

var member _void

type set struct {
	items map[interface{}]_void
}

func NewSet() *set {
	return &set{
		items: make(map[interface{}]_void),
	}
}

func (s *set) Add(item interface{}) {
	s.items[item] = member
}

func (s *set) Remove(item interface{}) {
	delete(s.items, item)
}

func (s *set) Contains(item interface{}) bool {
	_, exists := s.items[item]
	return exists
}

func (s *set) Size() int {
	return len(s.items)
}

// Items 返回 set 中的所有元素
func (s *set) Items() []interface{} {
	keys := make([]interface{}, 0, s.Size())
	for k := range s.items {
		keys = append(keys, k)
	}
	return keys
}

func (s *set) Iter() chan interface{} {
	c := make(chan interface{})
	go func() {
		for item := range s.items {
			c <- item
		}
		close(c)
	}()
	return c
}

func (s *set) Loop(b func(interface{}) bool) {
	for k := range s.items {
		if !b(k) {
			break
		}
	}
}
