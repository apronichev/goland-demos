package storage

type Item interface {
	Identity() int
	SetIdentity(identity int)
}

type Storage[T Item] interface {
	Init()
	Get(id int) T
	GetAll() []T
	Put(item T) T // todo rename method
	Remove(id int) T
}
