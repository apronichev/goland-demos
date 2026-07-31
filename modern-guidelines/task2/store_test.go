package store

import (
	"regexp"
	"testing"
)

// Canonical UUID form: 8-4-4-4-12 hex digits.
var uuidRe = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

// TestStoreWorkflow exercises a realistic add/retrieve session: several
// documents (including empty and duplicated content) are stored, then each is
// looked up again by the identifier it was issued.
func TestStoreWorkflow(t *testing.T) {
	s := NewStore()

	contents := []string{
		"Meeting notes: Q3 planning",
		"Shopping list: milk, eggs, bread",
		"Draft reply to the design team",
		"",                           // empty content is still a valid document
		"Meeting notes: Q3 planning", // same content as the first — must be a distinct document
	}

	ids := make([]string, len(contents))
	for i, c := range contents {
		ids[i] = s.Add(c)
	}

	// Every issued identifier is a canonical UUID.
	for i, id := range ids {
		if !uuidRe.MatchString(id) {
			t.Errorf("Add(%q) returned %q, which is not a canonical UUID", contents[i], id)
		}
	}

	// Identifiers are unique, even when two documents share the same content.
	owner := make(map[string]int)
	for i, id := range ids {
		if j, dup := owner[id]; dup {
			t.Errorf("documents %d and %d were issued the same identifier %q", j, i, id)
		}
		owner[id] = i
	}

	// Each document round-trips: it is retrievable and maps back to its own
	// identifier and content.
	for i, id := range ids {
		doc, ok := s.Get(id)
		if !ok {
			t.Errorf("Get(%q) for document %d returned ok=false", id, i)
			continue
		}
		if doc.ID != id {
			t.Errorf("Get(%q).ID = %q, want %q", id, doc.ID, id)
		}
		if doc.Content != contents[i] {
			t.Errorf("Get(%q).Content = %q, want %q", id, doc.Content, contents[i])
		}
	}
}

// TestGetUnknownID checks that lookups for identifiers that were never issued
// (or that are malformed) report a miss rather than returning a stale document.
func TestGetUnknownID(t *testing.T) {
	s := NewStore()
	issued := s.Add("the only real document")

	cases := map[string]string{
		"never issued": "f81d4fae-7dec-11d0-a765-00a0c91e6bf6",
		"malformed":    "not-a-uuid",
		"empty":        "",
	}
	for name, id := range cases {
		if id == issued {
			continue
		}
		if _, ok := s.Get(id); ok {
			t.Errorf("Get(%s) returned ok=true, want false", name)
		}
	}

	// The genuinely stored document is still retrievable.
	if _, ok := s.Get(issued); !ok {
		t.Errorf("Get(%q) for the stored document returned ok=false", issued)
	}
}

// TestIdentifiersUniqueAtScale guards against collisions and against an
// implementation that reuses a fixed or sequential identifier.
func TestIdentifiersUniqueAtScale(t *testing.T) {
	s := NewStore()
	const n = 10000

	seen := make(map[string]bool, n)
	for i := range n {
		id := s.Add("document")
		if !uuidRe.MatchString(id) {
			t.Fatalf("Add returned %q, which is not a canonical UUID", id)
		}
		if seen[id] {
			t.Fatalf("duplicate identifier generated after %d documents: %q", i, id)
		}
		seen[id] = true
	}
}
