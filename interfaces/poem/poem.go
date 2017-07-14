package main

// A Poem contains some poetry and an abstract storage reference.
type Poem struct {
	content []byte
	storage PoemStorage
}

// PoemStorage is just an interface that defines the behavior of a poem
// storage. This is all that Poem (and needs to know) about storing
// and retrieving poems. Nothing from the "outer ring" appers here.
type PoemStorage interface {
	Type() string
	Load(string) []byte
	Save(string, []byte)
}

// NewPoem constructs a Poem object. We use this constructor to inject an
// object that satisfies the PoemStorage interface
func NewPoem(ps PoemStorage) *Poem {
	return &Poem{
		content: []byte("I am a poem from a " + ps.Type() + "."),
		storage: ps,
	}
}
