package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/jroimartin/gocui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	go move(g)
	go changeShape(g)

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

func changeShape(g *gocui.Gui) {
	for {
		if !IsCanMoveToDown {
			IsCanMoveToDown = true
			IsCanMoveToLeft = true
			IsCanMoveToRight = true

			rand.Seed(time.Now().UnixNano())
			CurrentDirection = rand.Intn(4)
			CurrentShape = rand.Intn(7)
			pos = [2]int{30, -1}

			ShapeArr = getShapeArr(CurrentShape, CurrentDirection, pos)
			if err := drawShape(g, &ShapeArr); err != nil {
				log.Panic(err)
			}
		}
	}
}

func move(g *gocui.Gui) {
	for {
		time.Sleep(500 * time.Millisecond)
		if IsCanMoveToDown {
			pos[1]++
			updateShape(g)
		}
		g.Update(func(g *gocui.Gui) error {
			return nil
		})
	}
}
