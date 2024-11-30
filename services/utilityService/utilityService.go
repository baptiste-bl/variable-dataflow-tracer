// General functions that are not directly related to the manipulation of nodes or languages, but provide general utility.

package utilityService

// -----------------------------------------------------------------------------
// ContainsString - Checks if a slice contains a specific string
// -----------------------------------------------------------------------------
//
// Parameters:
//   - slice ([]string): The slice to check.
//   - item (string): The string to look for in the slice.
//
// Returns:
//   - (bool): True if the slice contains the string, otherwise false.
//
// -----------------------------------------------------------------------------
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

// -----------------------------------------------------------------------------
// NewOrderedSet - Creates a new instance of OrderedSet
// -----------------------------------------------------------------------------
//
// Parameters:
//   - None
//
// Returns:
//   - (*OrderedSet): A pointer to the newly created OrderedSet.
//
// -----------------------------------------------------------------------------
func NewOrderedSet() *OrderedSet {
	return &OrderedSet{
		items: make([]string, 0),
		set:   make(map[string]struct{}),
	}
}

// -----------------------------------------------------------------------------
// Add - Adds a string to the OrderedSet if it is not already present
// -----------------------------------------------------------------------------
//
// Parameters:
//   - item (string): The string to add to the OrderedSet.
//
// Returns:
//   - None
//
// -----------------------------------------------------------------------------
func (s *OrderedSet) Add(item string) {
	if _, exists := s.set[item]; !exists {
		s.items = append(s.items, item)
		s.set[item] = struct{}{}
	}
}

// -----------------------------------------------------------------------------
// Remove - Removes a string from the OrderedSet if it is present
// -----------------------------------------------------------------------------
//
// Parameters:
//   - item (string): The string to remove from the OrderedSet.
//
// Returns:
//   - None
//
// -----------------------------------------------------------------------------
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

// -----------------------------------------------------------------------------
// Contains - Checks if the OrderedSet contains a specific string
// -----------------------------------------------------------------------------
//
// Parameters:
//   - item (string): The string to check for in the OrderedSet.
//
// Returns:
//   - (bool): True if the OrderedSet contains the string, otherwise false.
//
// -----------------------------------------------------------------------------
func (s *OrderedSet) Contains(item string) bool {
	_, exists := s.set[item]
	return exists
}

// -----------------------------------------------------------------------------
// ContainString - Checks if a slice contains a specific string
// -----------------------------------------------------------------------------
//
// Parameters:
//   - slice ([]string): The slice to check.
//   - item (string): The string to look for in the slice.
//
// Returns:
//   - (bool): True if the slice contains the string, otherwise false.
//
// -----------------------------------------------------------------------------
func ContainString(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// -----------------------------------------------------------------------------
// Items - Returns all items in the OrderedSet in the order they were added
// -----------------------------------------------------------------------------
//
// Parameters:
//   - None
//
// Returns:
//   - ([]string): A slice of strings containing all items in the OrderedSet.
//
// -----------------------------------------------------------------------------
func (s *OrderedSet) Items() []string {
	return s.items
}

// -----------------------------------------------------------------------------
// Max - Calculates the maximum of two integers
// -----------------------------------------------------------------------------
//
// Parameters:
//   - x (int): The first integer.
//   - y (int): The second integer.
//
// Returns:
//   - (int): The maximum of the two integers.
//
// -----------------------------------------------------------------------------
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// -----------------------------------------------------------------------------
// Min - Calculates the minimum of two integers
// -----------------------------------------------------------------------------
//
// Parameters:
//   - x (int): The first integer.
//   - y (int): The second integer.
//
// Returns:
//   - (int): The minimum of the two integers.
//
// -----------------------------------------------------------------------------
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
