package board

import "fmt"

type Matrix88 [8][8]int

type Coordinates struct {
	x, y int // 0 to 7
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
	IsFinal() bool
	PlayAt(color int, position Coordinates) (Board, error)
}

func (m *Matrix88) String() string {
	s := ""
	for x := 0 ; x < 8 ; x++ {
		for y := 0 ; y < 8 ; y++ {
			switch m[x][y] {
			case EMPTY:
				s += " ."
			case BLACK:
				s += " X"
			case WHITE:
				s += " O"
			}
		}
		s += "\n"
	}
    // TODO show where players can play and number of possibilities by player
	black, white := m.CountByPlayer()
	s += fmt.Sprintf("Black: %02d - White: %02d\n", black, white)
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

func (m *Matrix88) CanPlayerPlayAt(color int, position Coordinates) bool {
	if m[position.x][position.y] != EMPTY {
		return false
	}
	// TODO short-circuit - find at least one
	return true
}

func (m *Matrix88) IsFinal() bool {
	black, white := m.PossibilitiesByPlayer()
	if black == 0 && white == 0 {
		return true
	}
	return false
}

func (m *Matrix88) PlayAt(color int, position Coordinates) (Board, error) {
	// TODO
	return m, nil
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
