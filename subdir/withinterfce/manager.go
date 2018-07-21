// Package withinterfce provides a test on interface type for checking
// godoc documentation.
package withinterfce

type manager struct {
	name string
}

// const checking
const (
	// A enum checking
	A = iota
	B
	C
)

type Manager interface {
	GetName() string
}

func New(name string) Manager {
	return &manager{name: name}
}

func (m *manager) GetName() string {
	return m.name
}
