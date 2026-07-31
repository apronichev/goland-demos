package store

// Document is an item held by the Store.
type Document struct {
	ID      string
	Content string
}

// Store is an in-memory collection of documents keyed by their identifier.
type Store struct {
	docs map[string]Document
}

// NewStore returns an empty Store.
func NewStore() *Store {
	return &Store{docs: make(map[string]Document)}
}

// Add stores content as a new document, assigns it a unique identifier, and
// returns that identifier.
//
// TODO: implement.
func (s *Store) Add(content string) string {
	return ""
}

// Get returns the document with the given id, and whether it was found.
func (s *Store) Get(id string) (Document, bool) {
	d, ok := s.docs[id]
	return d, ok
}
