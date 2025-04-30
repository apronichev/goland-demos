package quick_doc

type Ctx struct {
	DebugEnabled bool
}

func (user *User) Validate() bool {
	return user.Id >= 0 // `user` is dereferenced at `user.Id`
}
