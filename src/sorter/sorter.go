package sorter

// SortedSlice defines the interface for storing sorted data
type SortedSlice interface {
	Insert(interface{})
	String() string
}
