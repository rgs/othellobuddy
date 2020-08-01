package board

import "fmt"

type Matrix88 [8][8]int

type Coordinates struct {
	X, Y int // 0 to 7
}

const (
	EMPTY = 0
	BLACK = 1
	WHITE = 2
)

type Board interface {
	CountByPlayer() (black, white int)
	PossibilitiesByPlayer() (black, white int)
	CanPlayerPlayAt(color int, position Coordinates) bool
	IsStable(position Coordinates) bool
	StableDisksByPlayer() (black, white int)
	IsFinal() bool
	PlayAt(color int, position Coordinates) (Board, error)
}

func (m *Matrix88) String() string {
	s := "  a b c d e f g h\n"
	possibilitiesBlack, possibilitiesWhite := 0, 0
	for x := 0 ; x < 8 ; x++ {
		s += fmt.Sprintf("%d", x + 1)
		for y := 0 ; y < 8 ; y++ {
			switch m[x][y] {
			case EMPTY:
				canBlack := m.CanPlayerPlayAt(BLACK, Coordinates{x,y})
				canWhite := m.CanPlayerPlayAt(WHITE, Coordinates{x,y})
				if (canBlack) {
					possibilitiesBlack++
				}
				if (canWhite) {
					possibilitiesWhite++
				}
				if (canBlack && canWhite) {
					s += " ;"
				} else if (canBlack) {
					s += " ."
				} else if (canWhite) {
					s += " ,"
				} else {
					s += " _"
				}
			case BLACK:
				s += " \u25CF" // BLACK CIRCLE
			case WHITE:
				s += " \u25CB" // WHITE CIRCLE
			}
		}
		s += "\n"
	}
	black, white := m.CountByPlayer()
	s += fmt.Sprintf("Black \u25CF %02d (%d) - White \u25CB %02d (%d)", black, possibilitiesBlack, white, possibilitiesWhite)
	return s
}

func (m *Matrix88) CountByPlayer() (black, white int) {
	for x := 0 ; x < 8 ; x++ {
		for y := 0 ; y < 8 ; y++ {
			switch m[x][y] {
			case BLACK:
				black++
			case WHITE:
				white++
			}
		}
	}
	return
}

func (m *Matrix88) PossibilitiesByPlayer() (black, white int) {
	for x := 0 ; x < 8 ; x++ {
		for y := 0 ; y < 8 ; y++ {
			if m.CanPlayerPlayAt(BLACK, Coordinates{x, y}) {
				black++
			}
			if m.CanPlayerPlayAt(WHITE, Coordinates{x, y}) {
				white++
			}
		}
	}
	return
}

func opponent(color int) int {
	if color == BLACK {
		return WHITE
	} else if color == WHITE {
		return BLACK
	} else {
		return EMPTY
	}
}

func (m *Matrix88) disksToFlipInOneDirection(player int, start Coordinates, dx, dy int) []Coordinates {
	// we already verified that start.X,start.Y is empty
	x, y := start.X + dx, start.Y + dy
	var flips []Coordinates
	enclosed := false
	opponent := opponent(player)
	for ; x >= 0 && y >= 0 && x < 8 && y < 8 ; x, y = x + dx, y + dy {
		if m[x][y] == opponent {
			flips = append(flips, Coordinates{x, y})
		} else if m[x][y] == player {
			enclosed = true
			break
		} else {
			break
		}
	}
	if enclosed {
		return flips
	}
	return nil
}

func (m *Matrix88) CanPlayerPlayAt(color int, position Coordinates) bool {
	if m[position.X][position.Y] != EMPTY {
		return false
	}
	// find at least one direction that works (short-circuit)
	var flips []Coordinates
	flips = m.disksToFlipInOneDirection(color, position, -1, -1)
	if len(flips) > 0 {
		return true
	}
	flips = m.disksToFlipInOneDirection(color, position, -1, 0)
	if len(flips) > 0 {
		return true
	}
	flips = m.disksToFlipInOneDirection(color, position, -1, 1)
	if len(flips) > 0 {
		return true
	}
	flips = m.disksToFlipInOneDirection(color, position, 0, -1)
	if len(flips) > 0 {
		return true
	}
	flips = m.disksToFlipInOneDirection(color, position, 0, 1)
	if len(flips) > 0 {
		return true
	}
	flips = m.disksToFlipInOneDirection(color, position, 1, -1)
	if len(flips) > 0 {
		return true
	}
	flips = m.disksToFlipInOneDirection(color, position, 1, 0)
	if len(flips) > 0 {
		return true
	}
	flips = m.disksToFlipInOneDirection(color, position, 1, 1)
	if len(flips) > 0 {
		return true
	}
	return false
}

func (m *Matrix88) IsStable(position Coordinates) bool {
	if m[position.X][position.Y] == EMPTY {
		return false
	}
	// TODO
	return true
}

func (m *Matrix88) StableDisksByPlayer() (black, white int) {
	for x := 0 ; x < 8 ; x++ {
		for y := 0 ; y < 8 ; y++ {
			if m.IsStable(Coordinates{x, y}) {
				if m[x][y] == BLACK {
					black++
				} else {
					white++
				}
			}
		}
	}
	return
}

func (m *Matrix88) IsFinal() bool {
	black, white := m.PossibilitiesByPlayer()
	if black == 0 && white == 0 {
		return true
	}
	return false
}

func (m *Matrix88) PlayAt(color int, position Coordinates) (Board, error) {
	if m[position.X][position.Y] != EMPTY {
		return nil, &BoardError{"Position already occupied", position}
	}
	var flips [][]Coordinates = make([][]Coordinates, 8)
	flips[0] = m.disksToFlipInOneDirection(color, position, -1, -1)
	flips[1] = m.disksToFlipInOneDirection(color, position, -1, 0)
	flips[2] = m.disksToFlipInOneDirection(color, position, -1, 1)
	flips[3] = m.disksToFlipInOneDirection(color, position, 0, -1)
	flips[4] = m.disksToFlipInOneDirection(color, position, 0, 1)
	flips[5] = m.disksToFlipInOneDirection(color, position, 1, -1)
	flips[6] = m.disksToFlipInOneDirection(color, position, 1, 0)
	flips[7] = m.disksToFlipInOneDirection(color, position, 1, 1)
	flipped := false
	for i := 0 ; i < 8 ; i++ {
		for _, c := range flips[i] {
			m[c.X][c.Y] = color
			flipped = true
		}
	}
	if flipped {
		m[position.X][position.Y] = color
		return m, nil
	}
	return nil, &BoardError{"Cannot play here", position}
}

func InitialBoard() Board {
	return &Matrix88{
		[8]int{0,0,0,0,0,0,0,0},
		[8]int{0,0,0,0,0,0,0,0},
		[8]int{0,0,0,0,0,0,0,0},
		[8]int{0,0,0,1,2,0,0,0},
		[8]int{0,0,0,2,1,0,0,0},
		[8]int{0,0,0,0,0,0,0,0},
		[8]int{0,0,0,0,0,0,0,0},
		[8]int{0,0,0,0,0,0,0,0},
	}
}
