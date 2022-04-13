package main

import "fmt"

const (
	Up = iota
	Down
	Left
	Right
)

type matrix [][]int

type positionDirection struct {
	position
	direction int
}

type position struct {
	x int
	y int
	matrix
	validDirections     canMoveDetermination
	validDirectionCache positionValidDirectionCache
}

type positionValidDirectionCache map[string]canMoveDetermination

type movementModifier struct {
	x int
	y int
}

type canMoveDetermination struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool
}

type directionPositions struct {
	Up    []position
	Down  []position
	Left  []position
	Right []position
}

func main() {
	test := matrix{[]int{9, 9, 4}, []int{6, 6, 8}, []int{2, 1, 1}}
	fmt.Println(longestIncreasingPath(test))

	startpos := test.getStartPosition()
	nextpos := startpos.movePosition(1, 0)
	nextpos2 := nextpos.movePosition(2, 0)

	fmt.Printf("%+v\n", startpos)

	nextpos.determineValidDirections()
	nextpos2.determineValidDirections()

	fmt.Printf("%+v\n", nextpos)
	fmt.Printf("%+v\n", nextpos2)
}

func longestIncreasingPath(m matrix) int {
	return m.getValue(0, 4)
}

func newMovementModifier(direction int) movementModifier {
	if direction == Up {
		return movementModifier{x: 0, y: -1}
	} else if direction == Down {
		return movementModifier{x: 0, y: 1}
	} else if direction == Left {
		return movementModifier{x: -1, y: 0}
	} else if direction == Right {
		return movementModifier{x: 1, y: 0}
	}

	return movementModifier{}
}

func (p position) canMove(direction int) bool {

	movement := newMovementModifier(direction)

	// if invalid direction
	if (movement == movementModifier{}) {
		return false
	}

	curValue := p.currentValue()
	nextValue := p.matrix.getMovementValue(p, movement)

	return curValue < nextValue
}

func (p position) movePosition(x int, y int) position {
	return position{
		x:                   x,
		y:                   y,
		matrix:              p.matrix,
		validDirectionCache: p.validDirectionCache,
	}
}

// func (p position) findPaths(positionValues directionPositions) (directionPositions, bool) {
// 	canMove := p.determineValidDirections()

// 	if (canMove == canMoveDetermination{}) {
// 		// no valid paths, break
// 		return positionValues, false
// 	}

// 	if canMove.Up {
// 		upPath := p.movePosition(Up).findPaths(positionValues)

// 		if upPath != positionValues {

// 			positionValues = append(positionValues, directionPositions{})
// 		}
// 	}
// }

func (p position) determineValidDirections() canMoveDetermination {
	if (p.validDirections == canMoveDetermination{}) {
		if (p.validDirectionCache[p.getCacheKey()] != canMoveDetermination{}) {
			p.validDirections = p.validDirectionCache[p.getCacheKey()]
		} else {
			p.validDirections = canMoveDetermination{
				Up:    p.canMove(Up),
				Down:  p.canMove(Down),
				Left:  p.canMove(Left),
				Right: p.canMove(Right),
			}
			p.validDirectionCache[p.getCacheKey()] = p.validDirections
		}
	}

	return p.validDirections
}

func (p position) currentValue() int {
	return p.matrix.getValue(p.x, p.y)
}

func (p position) getCacheKey() string {
	return fmt.Sprintf("%d-%d", p.x, p.y)
}

func (m matrix) getValue(x int, y int) int {
	if m.cellExists(x, y) == false {
		return 0
	}
	return m[y][x]
}

func (m matrix) cellExists(x int, y int) bool {
	if y >= 0 && y < len(m) {
		if x >= 0 && x < len(m[y]) {
			return true
		}
	}

	return false
}

func (m matrix) getMovementValue(p position, mm movementModifier) int {
	return m.getValue(p.x+mm.x, p.y+mm.y)
}

func (m matrix) getStartPosition() position {
	position := position{
		x:                   0,
		y:                   0,
		matrix:              m,
		validDirectionCache: make(positionValidDirectionCache),
	}

	position.determineValidDirections()
	return position
}

// func createPositionsFromMatrix(m matrix) []position {
// 	var queue []position

// 	for i, j := range m {
// 		for k := range j {
// 			position := position{
// 				x:      i,
// 				y:      k,
// 				matrix: m,
// 			}
// 			position.determineValidDirections()
// 			queue = append(queue, position)
// 		}
// 	}

// 	return queue
// }
