package main

import (
	"sync"
)

func NewSyncStore() *SyncStore {
	return &SyncStore{}
}

func (s *SyncStore) CreatePuppy(p *Puppy) {
	s.Lock()
	defer s.Unlock()
	s.Store(p.ID, *Puppy)
}

func (s *SyncStore) ReadPuppy(pID uint64) *Puppy {
	puppyData := s.Load(pID)
	puppy := puppyData.(Puppy)
	return &puppy
}

func (s *SyncStore) UpdatePuppy(pID uint64, p *Puppy) {
	s.store(pID, *p)
}

func (s *SyncStore) DeletePuppy(pID uint64) {
	s.Lock()
	defer s.Unlock()
	s.Delete(pID)
}
