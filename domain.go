package main

type Status int

var (
	TodoState  Status = 1
	DoingState Status = 2
	DoneState  Status = 3
)

type Todo struct {
	ID     uint
	Title  string
	Status Status
}
