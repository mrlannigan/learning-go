package main

import "fmt"

const (
	Up = iota
	Down
	Left
	Right
)

type matrix [][]int

type position struct {
	x           int
	y           int
	matrix      *matrix
	connections []pathConnection
}

type pathConnection struct {
	to        *position
	from      *position
	direction int
}

type pathConnectionContainer struct {
	connections []*[]pathConnection
}

type movementModifier struct {
	x int
	y int
}

func main() {
	test := matrix{[]int{9, 9, 4}, []int{6, 6, 8}, []int{2, 1, 1}}
	fmt.Printf("Longest path length of matrix %+v is %d\n", test, longestIncreasingPath(test))
}

func longestIncreasingPath(m matrix) int {
	positions := createPositionsFromMatrix(m)
	allPaths := pathConnectionContainer{}

	for _, position := range positions {
		flattenConnections(position, []pathConnection{}, &allPaths)
	}

	longestPathLen, _ := allPaths.findLongestPath()

	return longestPathLen
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

func (pcc pathConnectionContainer) findLongestPath() (int, []pathConnection) {
	longestPathLen := 0
	var longestPath []pathConnection

	for _, pathConnections := range pcc.connections {
		currentLen := len(*pathConnections)

		if currentLen > longestPathLen {
			longestPath = *pathConnections
			longestPathLen = currentLen
		}
	}

	return longestPathLen + 1, longestPath
}

func createPositionsFromMatrix(m matrix) []*position {
	var queue []*position
	cacheMap := map[string]*position{}

	for i, j := range m {
		for k := range j {
			position := position{
				x:      k,
				y:      i,
				matrix: &m,
			}
			cacheMap[position.getCacheKey()] = &position
			queue = append(queue, &position)
		}
	}

	for _, p := range queue {
		for _, d := range []int{Up, Down, Left, Right} {
			if p.canMove(d) {
				mod := newMovementModifier(d)
				p.connections = append(p.connections, pathConnection{
					from:      p,
					to:        cacheMap[fmt.Sprintf("%d-%d", p.x+mod.x, p.y+mod.y)],
					direction: d,
				})
			}
		}
	}

	return queue
}

func flattenConnections(position *position, currentPath []pathConnection, queue *pathConnectionContainer) {
	if len(position.connections) == 0 {
		return
	}

	for _, connection := range position.connections {
		nextConnection := append(currentPath, connection)
		queue.connections = append(queue.connections, &nextConnection)
		flattenConnections(connection.to, nextConnection, queue)
	}
}
