package main

func NewMapStore() *MapStore {
	return &MapStore{}
}

func (m *MapStore) CreatePuppy(p *Puppy) {
	append(m, p)
}

func (m *MapStore) ReadPuppy(pID uint64) *Puppy {
	return m[pID]
}

func (m *MapStore) UpdatePuppy(pID uint64, p *Puppy) {
	m[pID] = p
}

func (m *MapStore) DeletePuppy(pID uint64) {
	delete(m, pID)
}
