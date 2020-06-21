package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/jroimartin/gocui"
)

func (e *SizeError) Error() string {
	return e.errMsg
}

func getGuiSize(g *gocui.Gui) (int, int, error) {
	maxX, maxY := g.Size()
	if maxX < 50 || maxY < 20 {
		return maxX, maxY, &SizeError{"too small"}
	}
	return maxX, maxY, nil
}

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
	_, _, err := getGuiSize(g)
	if err != nil {
		return err
	}
	v_move, err := g.SetView("move", 0, 0, 60, 36)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v_move.Title = "Move"
	if v, err := g.SetView("score", 61, 0, 83, 20); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Score"
	}
	if v, err := g.SetView("record", 61, 21, 83, 36); err != nil {
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
