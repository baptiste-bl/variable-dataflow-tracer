// Fonctions générales qui ne sont pas liées directement à la manipulation des nœuds ou des langages, mais qui fournissent une utilité générale.

package utilityService

// Helper function to check if a slice contains a string
func ContainsString(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

type OrderedSet struct {
	items []string
	set   map[string]struct{}
}

func NewOrderedSet() *OrderedSet {
	return &OrderedSet{
		items: make([]string, 0),
		set:   make(map[string]struct{}),
	}
}

func (s *OrderedSet) Add(item string) {
	if _, exists := s.set[item]; !exists {
		s.items = append(s.items, item)
		s.set[item] = struct{}{}
	}
}

func (s *OrderedSet) Remove(item string) {
	if _, exists := s.set[item]; exists {
		delete(s.set, item)
		for i, v := range s.items {
			if v == item {
				s.items = append(s.items[:i], s.items[i+1:]...)
				break
			}
		}
	}
}

func (s *OrderedSet) Contains(item string) bool {
	_, exists := s.set[item]
	return exists
}

// Helper function to check if a slice contains a string
func ContainString(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func (s *OrderedSet) Items() []string {
	return s.items
}
