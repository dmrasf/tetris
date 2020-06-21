package main

import (
	"strconv"

	"github.com/jroimartin/gocui"
)

// 根据形状类型、旋转反向、中心点，返回形状数组（右下角）
func getShapeArr(shapeType int, direction int, centerPos [2]int) [][2]int {
	posX, posY := centerPos[0], centerPos[1]
	switch shapeType {
	case L:
		switch direction {
		case UP:
			return append(make([][2]int, 0), centerPos, [2]int{posX, posY - 1},
				[2]int{posX, posY - 2}, [2]int{posX + 2, posY})
		case LEFT:
			return append(make([][2]int, 0), [2]int{posX, posY + 1}, [2]int{posX + 2, posY},
				[2]int{posX - 2, posY + 1}, [2]int{posX + 2, posY + 1})
		case DOWN:
			return append(make([][2]int, 0), [2]int{posX, posY}, [2]int{posX + 2, posY},
				[2]int{posX + 2, posY + 1}, [2]int{posX + 2, posY + 2})
		case RIGHT:
			return append(make([][2]int, 0), [2]int{posX, posY}, [2]int{posX + 2, posY},
				[2]int{posX + 4, posY}, [2]int{posX, posY + 1})
		}
	case RL:
		switch direction {
		case UP:
			return append(make([][2]int, 0), [2]int{posX + 2, posY - 1}, [2]int{posX, posY + 1},
				[2]int{posX + 2, posY}, [2]int{posX + 2, posY + 1})
		case LEFT:
			return append(make([][2]int, 0), [2]int{posX, posY}, [2]int{posX + 2, posY},
				[2]int{posX - 2, posY}, [2]int{posX + 2, posY + 1})
		case DOWN:
			return append(make([][2]int, 0), [2]int{posX, posY}, [2]int{posX + 2, posY},
				[2]int{posX, posY + 1}, [2]int{posX, posY + 2})
		case RIGHT:
			return append(make([][2]int, 0), [2]int{posX, posY}, [2]int{posX, posY + 1},
				[2]int{posX + 4, posY + 1}, [2]int{posX + 2, posY + 1})
		}
	case T:
		switch direction {
		case UP:
			return append(make([][2]int, 0), centerPos, [2]int{posX, posY - 1},
				[2]int{posX - 2, posY}, [2]int{posX + 2, posY})
		case LEFT:
			return append(make([][2]int, 0), [2]int{posX, posY}, [2]int{posX, posY - 1},
				[2]int{posX - 2, posY}, [2]int{posX, posY + 1})
		case DOWN:
			return append(make([][2]int, 0), [2]int{posX, posY}, [2]int{posX + 2, posY},
				[2]int{posX - 2, posY}, [2]int{posX, posY + 1})
		case RIGHT:
			return append(make([][2]int, 0), [2]int{posX, posY}, [2]int{posX + 2, posY},
				[2]int{posX, posY - 1}, [2]int{posX, posY + 1})
		}
	case S:
		if direction == DOWN || direction == UP {
			return append(make([][2]int, 0), centerPos, [2]int{posX, posY - 1},
				[2]int{posX + 2, posY}, [2]int{posX + 2, posY + 1})
		} else if direction == LEFT || direction == RIGHT {
			return append(make([][2]int, 0), [2]int{posX, posY}, [2]int{posX + 2, posY},
				[2]int{posX - 2, posY + 1}, [2]int{posX, posY + 1})
		}
	case RS:
		if direction == DOWN || direction == UP {
			return append(make([][2]int, 0), centerPos, [2]int{posX, posY + 1},
				[2]int{posX + 2, posY}, [2]int{posX + 2, posY - 1})
		} else if direction == LEFT || direction == RIGHT {
			return append(make([][2]int, 0), [2]int{posX, posY}, [2]int{posX + 2, posY + 1},
				[2]int{posX - 2, posY}, [2]int{posX, posY + 1})
		}
	case SQUARE:
		return append(make([][2]int, 0), centerPos, [2]int{posX + 2, posY},
			[2]int{posX, posY + 1}, [2]int{posX + 2, posY + 1})
	case I:
		if direction == DOWN || direction == UP {
			return append(make([][2]int, 0), [2]int{posX, posY - 1}, centerPos,
				[2]int{posX, posY + 1}, [2]int{posX, posY + 2})
		} else if direction == LEFT || direction == RIGHT {
			return append(make([][2]int, 0), [2]int{posX, posY}, [2]int{posX + 2, posY},
				[2]int{posX - 2, posY}, [2]int{posX + 4, posY})
		}
	}
	return make([][2]int, 0)
}

func updateShape(g *gocui.Gui) {
	if len(ShapeArr) != 0 {
		deleteShape(g, &ShapeArr)
	}
	ShapeArr = getShapeArr(I, CurrentPos, pos)
	drawShape(g, &ShapeArr)
}

// 根据形状数组绘制view
func drawShape(g *gocui.Gui, shapeArr *[][2]int) error {
	for _, item := range *shapeArr {
		_, err := g.SetView(getViewName(item),
			item[0]-2, item[1]-1, item[0], item[1])
		if err != nil && err != gocui.ErrUnknownView {
			return err
		}
	}
	return nil
}

// 根据形状数组删除view
func deleteShape(g *gocui.Gui, shapeArr *[][2]int) error {
	for _, item := range *shapeArr {
		if err := g.DeleteView(getViewName(item)); err != nil && err != gocui.ErrUnknownView {
			return err
		}
	}
	return nil
}

// 获取view名
func getViewName(pos [2]int) string {
	return strconv.Itoa(pos[0]) + "," + strconv.Itoa(pos[1])
}

// 根据形状判断是否接触到左边界或已固定的view
func isTouchLeftBorder(g *gocui.Gui, shapeArr [][2]int) bool {
	for _, item := range shapeArr {
		if item[0]-2 <= 0 {
			return false
		}
	}
	//for _, item := range *shapeArr {
	//if _, err := g.ViewByPosition(item[0]-4, item[1]); err != gocui.ErrUnknownView {
	//_, ok := itemName[[2]int{item[0] - 2, item[1]}]
	//if !ok {
	//return false
	//}
	//}
	//}
	return true
}

func isTouchRightBorder(g *gocui.Gui, shapeArr [][2]int) bool {
	v_move, _ := g.View("move")
	maxX, _ := v_move.Size()
	for _, item := range shapeArr {
		if item[0] > maxX {
			return false
		}
	}
	return true
}

func isTouchDownBorder(g *gocui.Gui, shapeArr [][2]int) bool {
	v_move, _ := g.View("move")
	_, maxY := v_move.Size()
	for _, item := range shapeArr {
		if item[1] > maxY {
			return false
		}
	}
	return true
}
