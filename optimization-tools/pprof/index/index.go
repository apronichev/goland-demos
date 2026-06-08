package index

import (
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/goland-demos/optimization-tools/pprof/model"
)

// Index is an in-memory inverted index over users.
type Index struct {
	mu       sync.Mutex
	postings map[string][]int64
	users    map[int64]model.User

	lastBuild time.Duration
}

// New returns an empty Index.
func New() *Index {
	return &Index{
		postings: make(map[string][]int64),
		users:    make(map[int64]model.User),
	}
}

// Build builds the index from a list of users.
func (idx *Index) Build(users []model.User) {
	start := time.Now()

	idx.mu.Lock()
	idx.postings = make(map[string][]int64)
	idx.users = make(map[int64]model.User, len(users))
	for _, u := range users {
		idx.users[u.ID] = u
	}
	idx.mu.Unlock()

	jobs := make(chan model.User, 1024)
	var wg sync.WaitGroup
	workers := runtime.NumCPU()
	for range workers {
		wg.Go(func() {
			for u := range jobs {
				idx.indexUser(u)
			}
		})
	}
	for _, u := range users {
		jobs <- u
	}
	close(jobs)
	wg.Wait()

	idx.mu.Lock()
	idx.lastBuild = time.Since(start)
	idx.mu.Unlock()
}

func (idx *Index) indexUser(u model.User) {
	if !validateEmail(u.Email) {
		return
	}

	nameTokens := tokenize(strings.ToLower(u.Name))
	emailTokens := tokenize(strings.ToLower(u.Email))
	countryTokens := tokenize(strings.ToLower(u.Country))
	bioTokens := tokenize(strings.ToLower(u.Bio))

	idx.mu.Lock()
	defer idx.mu.Unlock()
	for _, t := range nameTokens {
		idx.postings[t] = append(idx.postings[t], u.ID)
	}
	for _, t := range emailTokens {
		idx.postings[t] = append(idx.postings[t], u.ID)
	}
	for _, t := range countryTokens {
		idx.postings[t] = append(idx.postings[t], u.ID)
	}
	for _, t := range bioTokens {
		idx.postings[t] = append(idx.postings[t], u.ID)
	}
}

// Search returns every user that contains every token in the query.
func (idx *Index) Search(query string) []model.User {
	tokens := tokenize(strings.ToLower(query))
	if len(tokens) == 0 {
		return nil
	}

	idx.mu.Lock()
	defer idx.mu.Unlock()

	var matched map[int64]struct{}
	for i, t := range tokens {
		postings := idx.postings[t]
		set := make(map[int64]struct{}, len(postings))
		for _, id := range postings {
			set[id] = struct{}{}
		}
		if i == 0 {
			matched = set
			continue
		}
		for id := range matched {
			if _, ok := set[id]; !ok {
				delete(matched, id)
			}
		}
	}

	out := make([]model.User, 0, len(matched))
	for id := range matched {
		out = append(out, idx.users[id])
	}
	return out
}

// Get returns a user by ID and whether it was found.
func (idx *Index) Get(id int64) (model.User, bool) {
	idx.mu.Lock()
	defer idx.mu.Unlock()
	u, ok := idx.users[id]
	return u, ok
}

// Stats describes the current state of the index.
type Stats struct {
	UserCount         int           `json:"user_count"`
	TokenCount        int           `json:"token_count"`
	IndexingDuration  time.Duration `json:"indexing_duration_ns"`
	IndexingDurationS string        `json:"indexing_duration"`
}

// Stats reports current index statistics.
func (idx *Index) Stats() Stats {
	idx.mu.Lock()
	defer idx.mu.Unlock()
	return Stats{
		UserCount:         len(idx.users),
		TokenCount:        len(idx.postings),
		IndexingDuration:  idx.lastBuild,
		IndexingDurationS: idx.lastBuild.String(),
	}
}

// Users returns a copy of the indexed users — used by /reindex.
func (idx *Index) Users() []model.User {
	idx.mu.Lock()
	defer idx.mu.Unlock()
	out := make([]model.User, 0, len(idx.users))
	for _, u := range idx.users {
		out = append(out, u)
	}
	return out
}

// validateEmail compiles its regular expression on every call.
// This is the intended CPU hotspot.
func validateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func tokenize(s string) []string {
	return strings.Fields(s)
}
