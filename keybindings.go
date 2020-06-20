package main

import "github.com/jroimartin/gocui"

func keyBindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeySpace, gocui.ModNone, changeDir); err != nil {
		return err
	}
	return nil
}

func changeDir(g *gocui.Gui, v *gocui.View) error {
	switch CurrentPos {
	case UP:
		CurrentPos = RIGHT
	case LEFT:
		CurrentPos = UP
	case DOWN:
		CurrentPos = LEFT
	case RIGHT:
		CurrentPos = DOWN
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
