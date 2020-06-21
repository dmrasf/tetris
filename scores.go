package main

import "github.com/jroimartin/gocui"

func getScore(g *gocui.Gui) {
	v_move, _ := g.View("move")
	sizeX, sizeY := v_move.Size()
	for y := sizeY; y > 0; y-- {
		for x := 2; x <= sizeX; x += 2 {

		}
	}
}
