package main

import "fmt"

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

// Save simply calls Save on the interface type. The Poem object neither
// knows nor cares about which actual storage object receives this method
// call.
func (p *Poem) Save(name string) {
	p.storage.Save(name, p.content)
}

// Load also invokes the injected storage object object without knowing it.
func (p *Poem) Load(name string) {
	p.content = p.storage.Load(name)
}

// String makes Poem a Sringer, allowing us to drop it anywhere string
// would be expected.
func (p *Poem) String() string {
	return string(p.content)
}

// A Notebook is the classic storage device of a poet
type Notebook struct {
	poems map[string][]byte
}

func NewNoteBook() *Notebook {
	return &Notebook{
		poems: map[string][]byte{},
	}
}

func (n *Notebook) Save(name string, contents []byte) {
	n.poems[name] = contents
}

func (n *Notebook) Load(name string) []byte {
	return n.poems[name]
}

func (n *Notebook) Type() string {
	return "Notebook"
}

// A Napkin is the emergecy storage device of a poet. It can storage only
// one poem.
type Napkin struct {
	poem []byte
}

func NewNapkin() *Napkin {
	return &Napkin{
		poem: []byte{},
	}
}

func (n *Napkin) Save(name string, contents []byte) {
	n.poem = contents
}

func (n *Napkin) Load(name string) []byte {
	return n.poem
}

func (n *Napkin) Type() string {
	return "Napkip"
}

func main() {
	notebook := NewNoteBook()
	napkip := NewNapkin()

	poem := NewPoem(notebook)
	poem.Save("My first poem")

	poem = NewPoem(notebook)
	poem.Load("My first poem")
	fmt.Println(poem)

	poem = NewPoem(napkip)

	poem.Save("My second poem")
	poem = NewPoem(napkip)
	poem.Load("My second poem")
	fmt.Println(poem)
}
