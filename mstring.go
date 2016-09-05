package metrics

import "sync"

// plains hold an int64 value that can be set arbitrarily.
type Mstring interface {
	// Update the mstring's value.
	Update(value string)

	// Return the mstring's current value.
	Value() string
}

// The standard implementation of a mstring uses the sync/atomic package
// to manage a single int64 value.
type mstring struct {
	// just from Concurrent operation
	mutex sync.Mutex
	value string
}

// Create a new mstring.
func NewMString() Mstring {
	return &mstring{value: ""}
}

func (m *mstring) Update(s string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.value = s
}

func (m *mstring) Value() string {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.value
}
