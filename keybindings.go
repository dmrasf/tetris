package main

import (
	"github.com/jroimartin/gocui"
)

func keyBindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeySpace, gocui.ModNone, changeShapeDir); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowLeft, gocui.ModNone, shapeToLeft); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowRight, gocui.ModNone, shapeToRight); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, shapeToDown); err != nil {
		return err
	}
	return nil
}

func shapeToLeft(g *gocui.Gui, v *gocui.View) error {
	if IsCanMoveToLeft {
		pos[0] -= 2
		updateShape(g)
	}
	return nil
}
func shapeToRight(g *gocui.Gui, v *gocui.View) error {
	if IsCanMoveToRight {
		pos[0] += 2
		updateShape(g)
	}
	return nil
}
func shapeToDown(g *gocui.Gui, v *gocui.View) error {
	if IsCanMoveToDown {
		pos[1]++
		updateShape(g)
	}
	return nil
}

// 顺时针方向旋转
func changeShapeDir(g *gocui.Gui, v *gocui.View) error {
	switch CurrentDirection {
	case UP:
		CurrentDirection = RIGHT
	case LEFT:
		CurrentDirection = UP
	case DOWN:
		CurrentDirection = LEFT
	case RIGHT:
		CurrentDirection = DOWN
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
