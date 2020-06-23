package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func updateScore(g *gocui.Gui) {
	v_move, _ := g.View("move")
	sizeX, sizeY := v_move.Size()
	v_score, _ := g.View("score")
	fmt.Fprintf(v_score, "score = %d\n", Score)
	for {
		// 从最底部到shape最底部 判断得分情况
		shapePosY := ShapeArr[0][1]
		for y := sizeY + 1; y > shapePosY; y-- {
			x := 2
			for ; x <= sizeX+1; x += 2 {
				if _, err := g.View(getViewName([2]int{x, y})); err == gocui.ErrUnknownView || err == nil {
					break
				}
			}
			if x == sizeX+3 {
				Score++
				v_score.Clear()
				fmt.Fprintf(v_score, "score = %d\n", Score)
			}
		}
	}
}
