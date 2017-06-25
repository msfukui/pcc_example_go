package main

import "fmt"

// Ex は
// 与えられた 0, 1, 2, 3 の迷路を示す二次元配列
// 0 : 通路
// 1 : 壁
// 2 : スタート
// 3 : ゴール
// から、
// スタートからゴールまで移動するのに最小のターン数を返します。
func Ex(list [][]int) (resolve int) {
	// 与えられた問題を出力。
	ListPrintln(list)

	// 仮実装
	k := 22

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
