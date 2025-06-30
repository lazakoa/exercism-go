package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	row, exists := cb[file]
	if !exists {
		return 0
	}
	total := 0
	for i := range row {
		if row[i] {
			total += 1
		}
	}
	return total
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	total := 0
	for _, v := range cb {
		if pos := v[rank-1]; pos {
			total += 1
		}
	}
	return total
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	total := 0
	for _, v := range cb {
		total += len(v)
	}
	return total
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	total := 0
	for _, v := range cb {
		for _, sq := range v {
			if sq {
				total += 1
			}
		}
	}
	return total
}
