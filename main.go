package main

import (
	"log"
	"time"

	"github.com/jroimartin/gocui"
)

var pos = [2]int{4, 4}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	go move(g)

	keyBindings(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v_move, err := g.SetView("move", 0, 0, maxX-30, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v_move.Title = "Move"
	if v, err := g.SetView("score", maxX-29, 0, maxX-1, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Score"
	}
	if v, err := g.SetView("record", maxX-29, maxY/2+3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Record"
	}
	return nil
}

func move(g *gocui.Gui) {
	for {
		time.Sleep(1 * time.Second)
		arr := getShapeArr(SQUARE, UP, pos)
		deleteShape(g, &arr)
		pos[1]++
		arr1 := getShapeArr(SQUARE, UP, pos)
		drawShape(g, &arr1)

		g.Update(func(g *gocui.Gui) error {
			return nil
		})
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
