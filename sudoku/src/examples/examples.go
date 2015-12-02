// example puzzles
package examples


// completed puzzle
var ex1 [9][9]int = [9][9]int{
	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	{4, 5, 6, 7, 8, 9, 1, 2, 3},
	{7, 8, 9, 1, 2, 3, 4, 5, 6},
	{2, 3, 4, 5, 6, 7, 8, 9, 1},
	{5, 6, 7, 8, 9, 1, 2, 3, 4},
	{8, 9, 1, 2, 3, 4, 5, 6, 7},
	{3, 4, 5, 6, 7, 8, 9, 1, 2},
	{6, 7, 8, 9, 1, 2, 3, 4, 5},
	{9, 1, 2, 3, 4, 5, 6, 7, 8}}


// easy puzzle 4,859,614,836 from websudoku.com
var ex2 [9][9]int = [9][9]int{
	{0, 7, 0, 0, 4, 5, 9, 3, 2},
	{4, 0, 8, 0, 3, 0, 0, 0, 1},
	{0, 3, 0, 0, 2, 0, 0, 0, 0},
	{0, 0, 7, 0, 1, 0, 0, 4, 5},
	{0, 8, 9, 0, 0, 0, 1, 6, 0},
	{5, 4, 0, 0, 7, 0, 8, 0, 0},
	{0, 0, 0, 0, 8, 0, 0, 1, 0},
	{6, 0, 0, 0, 9, 0, 3, 0, 4},
	{8, 5, 4, 1, 6, 0, 0, 2, 0}}

// slices of examples
var Ex1 [][]int = toSlice(ex1)
var Ex2 [][]int = toSlice(ex2)


func toSlice(ex [9][9]int) [][]int {
	sl := make([][]int, 9)
	for i, _ := range(ex) {
		sl[i] = ex[i][:]
	}
	return sl
}
