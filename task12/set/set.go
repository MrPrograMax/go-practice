package set

type MySet struct {
	contains map[string]bool
	items    []string
}

func NewSet() *MySet {
	return &MySet{
		contains: make(map[string]bool),
		items:    make([]string, 0),
	}
}

func (set *MySet) GetItems() []string {
	return set.items
}

func (set *MySet) Add(item string) {
	if status := set.IsContains(item); !status {
		set.items = append(set.items, item)
		set.contains[item] = true
	}
}

func (set *MySet) AddSlice(item []string) {
	for _, value := range item {
		set.Add(value)
	}
}

func (set *MySet) IsContains(item string) bool {
	_, status := set.contains[item]
	return status
}
