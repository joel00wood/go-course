package main

import (
	"sync"
)

type Puppy struct {
	ID     uint64
	Breed  string
	Colour string
	Value  float64
}

type Storer interface {
	CreatePuppy(Puppy)
	ReadPuppy(ID uint64) Puppy
	UpdatePuppy(ID uint64, p Puppy)
	DeletePuppy(ID uint64)
}

type MapStore map[uint64]Puppy

func (m MapStore) CreatePuppy(p Puppy) {
	append(m, p)
}

func (m MapStore) ReadPuppy(pID uint64) Puppy {
	return m[pID]
}

func (m MapStore) UpdatePuppy(pID uint64, p Puppy) {
	m[pID] = p
}

func (m MapStore) DeletePuppy(pID uint64) {
	delete(m, pID)
}

type SyncStore struct {
	sync.Mutex
	sync.Map
}

func (s SyncStore) CreatePuppy(p Puppy) {
	s.Lock()
	defer s.Unlock()
	s.Store(p.ID, Puppy)
}

func (s SyncStore) ReadPuppy(pID uint64) Puppy {
	puppyData := s.Load(pID)
	puppy := puppyData.(Puppy)
	return puppy
}

func (s SyncStore) UpdatePuppy(pID uint64, p Puppy) {
	s.store(pID, p)
}

func (s SyncStore) DeletePuppy(pID uint64) {
	s.Lock()
	defer s.Unlock()
	s.Delete(pID)
}
