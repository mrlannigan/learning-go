package main

import "testing"

type testMatrices struct {
	matrix
	expectedPathLen int
}

var expectations = []testMatrices{
	{matrix: matrix{[]int{9, 9, 4}, []int{6, 6, 8}, []int{2, 1, 1}}, expectedPathLen: 4},
	{matrix: matrix{[]int{3, 4, 5}, []int{3, 2, 6}, []int{2, 2, 1}}, expectedPathLen: 4},
	{matrix: matrix{[]int{1}}, expectedPathLen: 1},
}

func TestLongestIncreasingPath(t *testing.T) {
	for caseIndex, testCase := range expectations {
		longestPath := longestIncreasingPath(testCase.matrix)
		if longestPath != testCase.expectedPathLen {
			t.Errorf("[Test Matrix %d] Expected a longest path length of %d, but received %d.", caseIndex, testCase.expectedPathLen, longestPath)
		}
	}
}

func TestGeneratedIncreasingPath(t *testing.T) {
	generatedMatrix := matrix{}

	// make a big L
	for y := 0; y < 500; y++ {
		generatedMatrix = append(generatedMatrix, make([]int, 0))
		for x := 0; x < 500; x++ {
			if x == 0 && y != 499 {
				generatedMatrix[y] = append(generatedMatrix[y], y+x+1)
			} else if y == 499 {
				generatedMatrix[y] = append(generatedMatrix[y], y+x+1)
			} else {
				generatedMatrix[y] = append(generatedMatrix[y], 0)
			}
		}
	}

	longestPath := longestIncreasingPath(generatedMatrix)
	if longestPath != 1000 {
		t.Errorf("Expected a longest path length of %d, but received %d.", 1000, longestPath)
	}
}
