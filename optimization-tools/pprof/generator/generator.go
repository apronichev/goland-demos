package generator

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/goland-demos/optimization-tools/pprof/model"
)

var firstNames = []string{
	"John", "Mary", "Robert", "Patricia", "Michael", "Linda", "William", "Barbara",
	"David", "Elizabeth", "Richard", "Susan", "Joseph", "Jessica", "Thomas", "Sarah",
	"Charles", "Karen", "Christopher", "Nancy", "Daniel", "Lisa", "Matthew", "Margaret",
	"Anthony", "Betty", "Mark", "Sandra", "Donald", "Ashley", "Steven", "Kimberly",
	"Paul", "Emily", "Andrew", "Donna", "Joshua", "Michelle", "Kenneth", "Carol",
	"Anna", "Sofia", "Liam", "Olivia", "Noah", "Emma", "Lucas", "Mia",
	"Ethan", "Amelia", "Aiden", "Isabella", "Hiroshi", "Yuki", "Wei", "Mei",
	"Carlos", "Maria", "Diego", "Lucia", "Ahmed", "Fatima", "Omar", "Aisha",
}

var lastNames = []string{
	"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller", "Davis",
	"Rodriguez", "Martinez", "Hernandez", "Lopez", "Gonzalez", "Wilson", "Anderson",
	"Thomas", "Taylor", "Moore", "Jackson", "Martin", "Lee", "Perez", "Thompson",
	"White", "Harris", "Sanchez", "Clark", "Ramirez", "Lewis", "Robinson", "Walker",
	"Young", "Allen", "King", "Wright", "Scott", "Torres", "Nguyen", "Hill",
	"Flores", "Green", "Adams", "Nakamura", "Tanaka", "Suzuki", "Sato", "Chen",
	"Wang", "Liu", "Khan", "Patel", "Singh", "Kumar", "Müller", "Schmidt",
}

var countries = []string{
	"United States", "Germany", "France", "United Kingdom", "Japan", "Brazil",
	"India", "China", "Canada", "Australia", "Spain", "Italy", "Mexico", "Russia",
	"Netherlands", "Sweden", "Norway", "Finland", "Denmark", "Poland", "Portugal",
	"Greece", "Turkey", "Egypt", "Argentina", "Chile", "South Korea", "Vietnam",
	"Thailand", "Indonesia", "South Africa", "Nigeria", "Kenya", "Israel",
	"Switzerland", "Austria", "Belgium", "Ireland", "New Zealand", "Singapore",
}

var emailDomains = []string{
	"example.com", "mail.com", "demo.org", "sample.net", "test.io",
	"corp.example", "users.example", "company.example", "service.example",
}

var interests = []string{
	"databases", "distributed systems", "programming languages", "compilers",
	"operating systems", "networking", "machine learning", "artificial intelligence",
	"cloud computing", "kubernetes", "docker", "containers", "microservices",
	"event sourcing", "functional programming", "type theory", "category theory",
	"cryptography", "security", "blockchain", "web development", "mobile development",
	"game development", "graphics programming", "computer vision", "robotics",
	"embedded systems", "performance optimization", "concurrent programming",
	"go", "rust", "python", "java", "kotlin", "scala", "erlang", "elixir",
	"javascript", "typescript", "haskell", "ocaml", "lisp", "ruby", "c++", "c",
	"linux", "open source software", "database tuning", "query optimization",
	"observability", "monitoring", "tracing", "profiling", "benchmarking",
	"data engineering", "stream processing", "batch processing", "data warehousing",
	"reading", "hiking", "running", "cycling", "photography", "music", "cooking",
	"traveling", "writing", "painting", "chess", "board games", "podcasts",
}

var bioTemplates = []string{
	"Interested in %s and %s.",
	"Passionate about %s, %s and %s.",
	"Working on %s. Enjoys %s in spare time.",
	"Software engineer focused on %s and %s.",
	"Builder. Tinkerer. Loves %s, %s and %s.",
	"Researcher exploring %s and %s.",
	"Currently learning %s. Background in %s.",
	"Lifelong student of %s, %s, and %s.",
}

// Generate produces n deterministic synthetic users.
func Generate(n int, seed uint64) []model.User {
	r := rand.New(rand.NewPCG(seed, seed^0x9E3779B97F4A7C15))
	users := make([]model.User, n)
	for i := 0; i < n; i++ {
		first := firstNames[r.IntN(len(firstNames))]
		last := lastNames[r.IntN(len(lastNames))]
		country := countries[r.IntN(len(countries))]
		domain := emailDomains[r.IntN(len(emailDomains))]

		name := first + " " + last
		email := fmt.Sprintf("%s.%s%d@%s",
			strings.ToLower(first),
			strings.ToLower(last),
			r.IntN(10000),
			domain,
		)

		bio := makeBio(r)

		users[i] = model.User{
			ID:      int64(i + 1),
			Name:    name,
			Email:   email,
			Country: country,
			Bio:     bio,
		}
	}
	return users
}

func makeBio(r *rand.Rand) string {
	tmpl := bioTemplates[r.IntN(len(bioTemplates))]
	slots := strings.Count(tmpl, "%s")
	args := make([]any, slots)
	for i := 0; i < slots; i++ {
		args[i] = interests[r.IntN(len(interests))]
	}
	return fmt.Sprintf(tmpl, args...)
}
