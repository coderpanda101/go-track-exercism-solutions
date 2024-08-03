package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File [8]bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	occupiedSquares := 0
	for _, rank := range cb[file] {
		if rank {
			occupiedSquares++
		}
	}
	return occupiedSquares
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	occupiedSquares := 0
	for _, file := range cb {
		if file[rank-1] {
			occupiedSquares++
		}
	}
	return occupiedSquares
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	totalSquares := 0
	for range cb {
		for i := 0; i < 8; i++ {
			totalSquares++
		}
	}
	return totalSquares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	totalOccupied := 0
	for _, file := range cb {
		for _, rank := range file {
			if rank {
				totalOccupied++
			}
		}
	}
	return totalOccupied
}
