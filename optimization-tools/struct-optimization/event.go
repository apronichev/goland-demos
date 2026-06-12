package main

type Event struct {
	OK   bool
	ID   int64
	Done bool
	Time int64
}

type OptimizedEvent struct {
	ID   int64
	Time int64
	OK   bool
	Done bool
}
