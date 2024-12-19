package wordsearch

import (
	"errors"
	"fmt"
)

type Point struct {
	row, col int
}

var directions = []Point{
	{0, 1},   // Left to Right
	{0, -1},  // Right to Left
	{1, 0},   // Top to Bottom
	{-1, 0},  // Bottom to Top
	{1, 1},   // Top-left to Bottom-right diagonal
	{-1, -1}, // Bottom-right to Top-left diagonal
	{1, -1},  // Top-right to Bottom-left diagonal
	{-1, 1},  // Bottom-left to Top-right diagonal
}

// searchFromPosition tries to find a word starting from a specific position in the puzzle
func searchFromPosition(puzzle []string, word string, startRow, startCol int) (bool, [2][2]int) {
	rows := len(puzzle)
	cols := len(puzzle[0])
	wordLen := len(word)

	for _, direction := range directions {
		endRow := startRow + direction.row*(wordLen-1)
		endCol := startCol + direction.col*(wordLen-1)

		// Check if the end position is within bounds
		if endRow >= 0 && endRow < rows && endCol >= 0 && endCol < cols {
			match := true
			for i := 0; i < wordLen; i++ {
				row := startRow + i*direction.row
				col := startCol + i*direction.col
				if puzzle[row][col] != word[i] {
					match = false
					break
				}
			}
			if match {
				return true, [2][2]int{{startCol, startRow}, {endCol, endRow}}
			}
		}
	}
	return false, [2][2]int{}
}

// Solve finds the given words in the puzzle and returns their coordinates
func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	if len(puzzle) == 0 || len(words) == 0 {
		return nil, errors.New("puzzle or word list is empty")
	}

	rows := len(puzzle)
	cols := len(puzzle[0])

	for _, row := range puzzle {
		if len(row) != cols {
			return nil, errors.New("puzzle rows must have the same length")
		}
	}

	foundWords := make(map[string][2][2]int)

	for _, word := range words {
		wordFound := false
		for r := 0; r < rows && !wordFound; r++ {
			for c := 0; c < cols && !wordFound; c++ {
				if puzzle[r][c] == word[0] {
					found, coords := searchFromPosition(puzzle, word, r, c)
					if found {
						foundWords[word] = coords
						wordFound = true
					}
				}
			}
		}
		if !wordFound {
			return nil, fmt.Errorf("word %s not found in the puzzle", word)
		}
	}

	return foundWords, nil
}
