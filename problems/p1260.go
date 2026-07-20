// Package p1260 solves LeetCode 1260: Shifting the Grid

package p1260

import "fmt"

// three approaches: 
// flatten: each destination in ans calculates its source from grid
// reversal: uses triple reversal technique to mutate grid in place
// scatter: each source calculates its destination in ans


// first approach: flatten, then rebuild (gather)
// build a fresh grid then for each destination index 
// compute which source index gives the value
// time O(m*n), space O(m*n)

func shiftGridFlatten(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	total := m*n
	k %= total 

	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	for idx := 0; idx < total; idx++ {
		// Go's % operator keeps the dividend's sign, so (idx-k)%total
		// can be negative. Adding total doesn't change the value modulo
		// total, but guarantees non-negative.
		srcIdx := idx - k
		if srcIdx < 0 {
			srcIdx += total
		}
		ans[idx/n][idx%n] = grid[srcIdx/n][srcIdx%n]
	}

	return ans
}

// second approach: triple reversal (transform in place)

// right rotating an array is equivalent to
// 1 - reverse the whole array
// 2 - reverse the first k elements
// 3 - reverse the remaining total-k elements

// in this code we use `get` and `set` to simulate flat indices for grid coords

// Time: O(m*n)
// Space: O(1)

func shiftGridReversal(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	total := m*n
	k %= total

	if k == 0 {
		return grid
	}

	get := func(idx int) int {return grid[idx/n][idx%n] }
	set := func(idx, val int) { grid[idx/n][idx%n] = val }

	reverse := func(lo, hi int) {
		for lo < hi {
			a, b := get(lo), get(hi)
			set(lo, b)
			set(hi, a)
			lo++
			hi--
		}
	}

	reverse(0, total-1)
	reverse(0, k-1)
	reverse(k, total-1)

	return grid
}

// third approach: scatter with position index

func shiftGridScatter(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	// we don't need to calculate total here because we are only looping 
	// `total` times below
	// we also don't mind large `ks` because we are wrapping mod `m`
	// so large k still has the correct destination, so we don't need to
	// k%total

	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	// starting position for our `scatter` is the offset `k`
	pos := k
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			row := (pos / n) % m
			col := pos % n
			ans[row][col] = grid[i][j]
			pos++
		}
	}

	return ans
}

// sanity testing: run all three on the same input

func main() {
	base := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},	
	}
	k := 2

	fmt.Println("flatten: ", shiftGridFlatten(base, k))
	fmt.Println("scatter: ", shiftGridScatter(base, k))
	// reverse mutates in place, so needs to be last:
	fmt.Println("reverse: ", shiftGridReversal(base, k))
}


