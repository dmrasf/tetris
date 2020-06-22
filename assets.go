package main

import "sync"

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
	CurrentDirection int
	CurrentShape     int
	CenterPos        = [2]int{30, -1}
	NextShapeArr     = [][2]int{}
	NextDirection    int
	NextShape        int

	IsCanMoveToLeft  = true
	IsCanMoveToRight = true
	IsCanMoveToDown  = true
	FixShape         = false

	Score = 0

	WriteCenterPosMutex sync.RWMutex
	WriteStateMutex     sync.RWMutex
)

type SizeError struct {
	errMsg string
}
