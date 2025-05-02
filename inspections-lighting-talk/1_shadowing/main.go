package main

func main() {
	config := LoadConfig()
	if config == nil {
		config, err := LoadConfigFromFile("config.txt")
		if err != nil || config == nil {
			println("Error loading config.txt")
			return
		}
	}
	println(config.id)
}
