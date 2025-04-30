package inspections

type Person struct {
	Name string
	Age  int
}

func process(a any) {}

type Config struct {
	Root int
}

func LoadConfig() *Config {
	return nil
}

func LoadConfigFromPath() (*Config, error) {
	return &Config{}, nil
}

type User struct {
	Name string
	Age  int
}

func isValid(name string) bool {
	return name != ""
}

func NewUser(name string) *User {
	if !isValid(name) {
		return nil
	}
	return &User{Name: name}
}

func LoadUser() (*User, error) {
	if !isValid("") {
		return nil, nil
	}
	return &User{Name: ""}, nil
}
