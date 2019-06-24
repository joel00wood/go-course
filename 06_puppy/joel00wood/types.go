package main

type Puppy struct {
	ID     uint64  `json:"id"`
	Breed  string  `json:"breed"`
	Colour string  `json:"colour"`
	Value  float32 `json:"value"`
}

type Storer interface {
	CreatePuppy(p *Puppy) error
	ReadPuppy(ID uint64) (*Puppy, error)
	UpdatePuppy(ID uint64, p *Puppy) error
	DeletePuppy(ID uint64) (bool, error)
}

type MapStore map[uint64](*Puppy)

type SyncStore struct {
	sync.Mutex
	sync.Map
}
