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

func init() {
	rand.Seed(time.Now().UnixNano())
	NextDirection = rand.Intn(4)
	NextShape = rand.Intn(7)
	CenterPos = [2]int{30, -1}
	NextShapeArr = getShapeArr(NextShape, NextDirection, CenterPos)
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	go changeShape(g)
	go move(g)
	go updateScore(g)
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
	if v, err := g.SetView("score", 61, 0, 83, 12); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Score"
	}
	if v, err := g.SetView("nextShape", 61, 13, 83, 20); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "NextShape"
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
		if FixShape || len(ShapeArr) == 0 {
			IsCanMoveToDown = true
			IsCanMoveToLeft = true
			IsCanMoveToRight = true
			FixShape = false

			ShapeArr = NextShapeArr
			CurrentDirection = NextDirection
			CurrentShape = NextShape
			if err := drawShape(g, ShapeArr); err != nil {
				log.Panic(err)
			}

			rand.Seed(time.Now().UnixNano())
			NextDirection = rand.Intn(4)
			NextShape = rand.Intn(7)
			CenterPos = [2]int{30, 0}

			deleteShape(g, getPreviewShapePos([2]int{30, 0}))
			NextShapeArr = getShapeArr(NextShape, NextDirection, [2]int{30, 0})
			drawShape(g, getPreviewShapePos([2]int{30, 0}))
		}
	}
}

func move(g *gocui.Gui) {
	for {
		time.Sleep(300 * time.Millisecond)
		shapeToDown(g, nil)

		g.Update(func(g *gocui.Gui) error {
			return nil
		})
	}
}
