package storage

type IdentityStore interface {
	GetIdentity(id string) (string, error)
	SaveIdentity(id string, data string) error
	RevokeIdentity(id string) error
	IsRevoked(id string) bool
}

type MemoryStore struct {
	data    map[string]string
	revoked map[string]bool
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data:    make(map[string]string),
		revoked: make(map[string]bool),
	}
}

func (m *MemoryStore) GetIdentity(id string) (string, error) {
	return m.data[id], nil
}

func (m *MemoryStore) SaveIdentity(id string, data string) error {
	m.data[id] = data
	return nil
}

func (m *MemoryStore) RevokeIdentity(id string) error {
	m.revoked[id] = true
	return nil
}

func (m *MemoryStore) IsRevoked(id string) bool {
	return m.revoked[id]
}
