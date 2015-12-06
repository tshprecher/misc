package common


const EmptyField = 0 // TODO: are we using this?

/*
  represent the general n-doku puzzle state
  where size == 3 => standard sudoku 9x9.
 */
type Ndoku struct{
	Size uint8
	Values [][]int
}

func MakeSudoku(values [][]int) (*Ndoku, bool) {
	return MakeNdoku(values, 3)
}

func MakeNdoku(values [][]int, size uint8) (*Ndoku, bool) {
	dim := int(size) * int(size)

	// verify dimensions of input and range of the values
	if len(values) != dim {
		return nil, false
	}
	for _, row := range(values) {
		if len(row) != dim {
			return nil, false
		}
		for _, v := range(row) {
			if v < 0 || v > dim {
				return nil, false
			}

		}
	}

	// copy the input size
	valuesCopy := make([][]int, len(values))
	for r, row := range(values) {
		rowCopy := make([]int, len(row))
		copy(rowCopy, row)
		valuesCopy[r] = rowCopy
	}
	return &Ndoku{size, valuesCopy}, true
}

func validate(ndoku *Ndoku, startRow int, startCol int, next func (row int, col int, size uint8) (int, int, bool)) bool {
	seen := make(map[int]bool)
	dim := int(ndoku.Size)*int(ndoku.Size)
	done := false

	for row, col := startRow, startCol; !done; row, col, done = next(row, col, ndoku.Size) {
		if row < 0 || row >= dim || col < 0 || col >= dim {
			return false
		}

		value := ndoku.Values[row][col]
		if value != EmptyField {
			if value < 1 || value > dim {
				return false
			}
			if seen[value] {
				return false
			}
			seen[value] = true
		}
	}
	return true
}

func IsValidRow(ndoku *Ndoku, row int) bool {
	return validate(ndoku, row, 0, func (r int, c int, s uint8) (int, int, bool) {
		if c+1 >= int(s)*int(s) {
			return 0, 0, true
		}
		return r, c+1, false
	})
}

func IsValidColumn(ndoku *Ndoku, col int) bool {
	return validate(ndoku, 0, col, func (r int, c int, s uint8) (int, int, bool) {
		if r+1 >= int(s)*int(s) {
			return 0, 0, true
		}
		return r+1, c, false
	})
}

func IsValidBlock(ndoku *Ndoku, row int, col int) bool {
	startRow := row / int(ndoku.Size)
	startCol := col / int(ndoku.Size)

	return validate(ndoku, startRow, startCol, func (r int, c int, s uint8) (int, int, bool) {
		size := int(s)
		newColumn, newRow := (c+1)/size, (r+1)/size

		if newRow > r/size && newColumn > c/size {
			return 0, 0, true
		}

		if newColumn == c/size {
			return r, c+1, false
		} else {
			return r+1, 0, false
		}
	})
}

func IsValid(ndoku *Ndoku) bool {
	// TODO: add method to Ndoku to compute the dim
	dim := int(ndoku.Size) * int(ndoku.Size)

	for d := 0; d < dim; d++ {
		ok := IsValidRow(ndoku, d)
		if !ok {
			return false
		}

		ok = IsValidColumn(ndoku, d)
		if !ok {
			return false
		}
	}

	for br := 0; br < int(ndoku.Size); br++ {
		for bc := 0; bc < int(ndoku.Size); bc++ {
			ok := IsValidBlock(ndoku, br * int(ndoku.Size), bc * int(ndoku.Size))
			if !ok {
				return false
			}
		}
	}
	return true
}
