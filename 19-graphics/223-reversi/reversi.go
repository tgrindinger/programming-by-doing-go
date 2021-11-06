package main

import "fmt"

type gameStatus int

const (
	STATUS_NONE gameStatus = iota
	STATUS_BLACK
	STATUS_WHITE
)

type Reversi struct {
	gameOver   bool
	started    bool
	showLabels bool
	showLegal  bool
	player     gameStatus
	other      gameStatus
	boxes      [][]*Rectangle
	pieces     [][]gameStatus
	scoreA     int
	scoreB     int
	legalMoves int
}

func NewReversi() *Reversi {
	size := 8
	boxes := make([][]*Rectangle, size)
	pieces := make([][]gameStatus, size)
	for r := 0; r < size; r++ {
		boxes[r] = make([]*Rectangle, size)
		pieces[r] = make([]gameStatus, size)
		for c := 0; c < size; c++ {
			boxes[r][c] = NewRectangleAsRect(float64(30 + c * 90), float64(100 + r * 90), 90.0, 90.0)
		}
	}
	return &Reversi{false, false, true, true, STATUS_BLACK, STATUS_WHITE, boxes, pieces, 0, 0, 0}
}

func (rev *Reversi) ResetBoard() {
	for r := 0; r < len(rev.pieces); r++ {
		for c := 0; c < len(rev.pieces[r]); c++ {
			rev.pieces[r][c] = STATUS_NONE
		}
	}
	rev.pieces[3][3] = STATUS_WHITE
	rev.pieces[3][4] = STATUS_BLACK
	rev.pieces[4][3] = STATUS_BLACK
	rev.pieces[4][4] = STATUS_WHITE
}

func (rev *Reversi) paint(g IGraphics) {
	g.setColor(COLOR_GREEN)
	g.fillRect(0, 0, 800, 900)
	g.setColor(COLOR_BLACK)
	g.setFont("Arial", SLANT_NORMAL, WEIGHT_BOLD, 48)
	if !rev.started {
		g.drawString("REVERSI", 305, 300)
		g.drawString("click to play", 275, 400)
	} else {
		rev.drawGameInProgress(g)
	}
}

func (rev *Reversi) drawGameInProgress(g IGraphics) {
	rev.drawBoard(g)
	rev.drawPieces(g)
	rev.drawLegalMoves(g)
	if rev.legalMoves == 0 {
		rev.swapPlayer()
		rev.drawLegalMoves(g)
		if rev.legalMoves == 0 {
			rev.gameOver = true
		}
	}
	rev.drawLabels(g)
}

func (rev *Reversi) drawLabels(g IGraphics) {
	if rev.showLabels {
		g.setFont("Arial", SLANT_NORMAL, WEIGHT_NORMAL, 24)
		g.setColor(COLOR_GRAY)
		for r := 0; r < len(rev.boxes); r++ {
			for c := 0; c < len(rev.boxes[0]); c++ {
				g.drawString(fmt.Sprintf("%d,%d", r, c), rev.boxes[r][c].x + 30, rev.boxes[r][c].y + 75)
			}
		}
	}
	g.setColor(COLOR_BLACK)
	g.setFont("Arial", SLANT_NORMAL, WEIGHT_BOLD, 48)
	turn := "Black's turn"
	if rev.player == STATUS_WHITE {
		turn = "White's turn"
	}
	g.drawString(turn, 50, 70)
	g.drawString(fmt.Sprintf("%d : %d", rev.scoreA, rev.scoreB), 620, 70)
}

func (rev *Reversi) drawLegalMoves(g IGraphics) {
	rev.legalMoves = 0
	g.setColor(COLOR_RED)
	for r := 0; r < len(rev.pieces); r++ {
		for c := 0; c < len(rev.pieces[0]); c++ {
			if rev.pieces[r][c] == STATUS_NONE && rev.isLegalMove(r, c) {
				rev.legalMoves++
				if rev.showLegal {
					box := rev.boxes[r][c]
					g.drawOval(box.x+5, box.y+5, box.width-10, box.height-10)
				}
			}
		}
	}
}

func (rev *Reversi) drawBoard(g IGraphics) {
	for r := 0; r < len(rev.boxes); r++ {
		for c := 0; c < len(rev.boxes[0]); c++ {
			g.draw(rev.boxes[r][c])
		}
	}
}

func (rev *Reversi) drawPieces(g IGraphics) {
	rev.scoreA = 0
	rev.scoreB = 0
	for r := 0; r < len(rev.pieces); r++ {
		for c := 0; c < len(rev.pieces[0]); c++ {
			if rev.pieces[r][c] > 0 {
				if rev.pieces[r][c] == STATUS_BLACK {
					g.setColor(COLOR_BLACK)
					rev.scoreA++
				} else if rev.pieces[r][c] == STATUS_WHITE {
					g.setColor(COLOR_WHITE)
					rev.scoreB++
				}
				box := rev.boxes[r][c]
				g.fillOval(box.x+5, box.y+5, box.width-10, box.height-10)
			}
		}
	}
}

func (rev *Reversi) mouseReleased(x, y float64) {
	if rev.gameOver {
		rev.started = false
		rev.gameOver = false
	} else if !rev.started {
		rev.started = true
		rev.ResetBoard()
	} else {
		for r := 0; r < len(rev.boxes); r++ {
			for c := 0; c < len(rev.boxes[0]); c++ {
				if rev.boxes[r][c].contains(x, y) {
					if rev.pieces[r][c] == STATUS_NONE && rev.isLegalMove(r, c) {
						rev.pieces[r][c] = rev.player
						rev.flipOthers(r, c)
						rev.swapPlayer()
					}
					return
				}
			}
		}
	}
}

func (rev *Reversi) swapPlayer() {
	rev.other = rev.player
	if rev.player == STATUS_BLACK {
		rev.player = STATUS_WHITE
	} else {
		rev.player = STATUS_BLACK
	}
}

func (rev *Reversi) isLegalMove(r, c int) bool {
	for _, dir := range compass {
		first := rev.firstMatch(r, c, dir)
		adj := NewLocation(r, c).getAdjacentLocation(dir)
		if first != nil && adj != nil && rev.otherStreak(r, c, dir) > 0 {
			return true
		}
	}
	return false
}

func (rev *Reversi) firstMatch(r, c int, dir *Direction) *Location {
	cur := NewLocation(r, c)
	for {
		cur = cur.getAdjacentLocation(dir)
		if !rev.isValidLoc(cur) {
			return nil
		}
		if rev.pieces[cur.r][cur.c] == rev.player {
			return cur
		}
	}
}

func (rev *Reversi) flipOthers(r, c int) {
	for _, dir := range compass {
		rev.flipOthersDir(r, c, dir)
	}
}

func (rev *Reversi) flipOthersDir(r, c int, dir *Direction) {
	cur := NewLocation(r, c)
	streak := rev.otherStreak(r, c, dir)
	for i := 0; i < streak; i++ {
		cur = cur.getAdjacentLocation(dir)
		rev.pieces[cur.r][cur.c] = rev.player
	}
}

func (rev *Reversi) otherStreak(r, c int, dir *Direction) int {
	count := 0
	cur := NewLocation(r, c)
	for {
		cur = cur.getAdjacentLocation(dir)
		if !rev.isValidLoc(cur) || rev.pieces[cur.r][cur.c] == STATUS_NONE {
			count = 0
			break
		}
		if rev.pieces[cur.r][cur.c] == rev.player {
			break
		}
		count++
	}
	return count
}

func (rev *Reversi) isValidLoc(loc *Location) bool {
	return 0 <= loc.r && loc.r < len(rev.pieces) && 0 <= loc.c && loc.c < len(rev.pieces)
}

func (rev *Reversi) isValid(r, c int) bool {
	return 0 <= r && r < len(rev.pieces) && 0 <= c && c < len(rev.pieces)
}
