package storage

type IdentityStore interface {
	GetIdentity(id string) (string, error)
	SaveIdentity(id string, data string) error
}

type MemoryStore struct {
	data map[string]string
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{data: make(map[string]string)}
}

func (m *MemoryStore) GetIdentity(id string) (string, error) {
	return m.data[id], nil
}

func (m *MemoryStore) SaveIdentity(id string, data string) error {
	m.data[id] = data
	return nil
}
