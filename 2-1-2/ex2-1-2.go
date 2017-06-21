package main

import "fmt"

// Ex は
// 与えられた 0, 1 の二次元配列から、
// 8 近傍で隣接している塊の数を返します。
func Ex(list [][]int) (resolve int) {
	// 与えられた問題を出力。
	ListPrintln(list)

	k := 0
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list[i]); j++ {
			// 最初の 1 を見つけたら、
			if list[i][j] == 1 {
				// カウントアップして、
				k++
				// そこから 8 近傍でつながっている箇所を
				// 探索して 0 にリセット。
				ListSearchAndRewrite(list, i, j)
			}
		}
	}

	return k
}

// ListPrintln は
// 与えられた 0, 1 の二次元配列を整形して標準出力します。
func ListPrintln(list [][]int) {
	fmt.Printf("N = %v, M = %v\n\n", len(list), len(list[0]))

	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list[i]); j++ {
			fmt.Print(list[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// ListSearchAndRewrite は
// 与えられた 0, 1 の二次元配列と計算の起点となる座標から、
// その全ての 8 近傍をたどって 1 を 0 に書き換えます。
func ListSearchAndRewrite(list [][]int, x int, y int) {

	// 探索対象が 0 の場合はその場で終了。
	if list[x][y] == 0 {
		return
	}

	// 現在の座標を 0 にリセット。
	list[x][y] = 0

	// 探索範囲の 8 近傍の座標の初期値と max 値を設定。
	minX := x - 1
	if minX < 0 {
		minX = 0
	}
	minY := y - 1
	if minY < 0 {
		minY = 0
	}
	maxX := x + 1
	if maxX >= len(list) {
		maxX = len(list) - 1
	}
	maxY := y + 1
	if maxY >= len(list[maxX]) {
		maxY = len(list[maxX]) - 1
	}

	// 現在の座標を起点にして 8 近傍を探索して、
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			// 値が 1 なら、
			if list[i][j] == 1 {
				// その点を起点に 8 近傍を探索して 0 にリセット。
				ListSearchAndRewrite(list, i, j)
			}
		}
	}

	return
}
