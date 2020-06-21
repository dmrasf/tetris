package main

const (
	L      int = 0
	RL     int = 1
	SQUARE int = 2
	T      int = 3
	S      int = 4
	RS     int = 5
	I      int = 6

	UP    int = 0
	LEFT  int = 1
	DOWN  int = 2
	RIGHT int = 3
)

var (
	ShapeArr         = [][2]int{}
	CurrentDirection = UP
	CurrentShape     = I
	pos              = [2]int{30, -1}
	IsCanMoveToLeft  = true
	IsCanMoveToRight = true
	IsCanMoveToDown  = true
	Score            = 0
)

type SizeError struct {
	errMsg string
}
