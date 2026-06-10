package inspections

func _() {
	conf := LoadConfig()

	if conf == nil {
		conf, err := LoadConfigFromPath() // `conf` shadows outer `conf` variable
		if err != nil || conf == nil {
			return
		}
	}
	process(conf.Root)
}

func _() {
	user := NewUser("123")
	println(user.Age, user.Name)
}

func _() {
	user, err := LoadUser()
	if err != nil {
		return
	}
	println(user.Age, user.Name)
}
