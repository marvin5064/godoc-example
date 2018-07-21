package withinterfce

type manager struct {
	name string
}
type Manager interface {
	GetName() string
}

func New(name string) Manager {
	return &manager{name: name}
}

func (m *manager) GetName() string {
	return m.name
}
