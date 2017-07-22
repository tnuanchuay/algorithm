package main

import (
	"fmt"
	"time"
)

const(
	MAX_DIMENSION =	4
)

type Tree struct {
	Row	int
	Col 	int
	Board  [MAX_DIMENSION][MAX_DIMENSION]int
	Branch [8]*Tree
}

func main() {
	var start_row, start_col int

	fmt.Scanf("%d,%d", &start_row, &start_col)

	var tree *Tree = new(Tree)
	tree.Row = start_row
	tree.Col = start_col
	tree.Board[start_row][start_col] = 1
	start_time := time.Now()
	CreateEntireTree(tree, 0)
	high := findHigh(tree) - 1
	duration := time.Since(start_time)
	fmt.Println(high, "times horse can walk")
	fmt.Println("time for calculation", duration.Seconds(), "seconds")

}

func findHigh(tree *Tree) int{
	if isEndLeaf(tree) {
		return 1;
	}else {
		var max int
		for i := 0; i < len(tree.Branch) ; i++{
			if tree.Branch[i] != nil{
				h := findHigh(tree.Branch[i])
				if max < h{
					max = h
				}
			}
		}

		return max + 1
	}
}

func isEndLeaf(tree *Tree) bool{
	for i := 0; i < len(tree.Branch); i++{
		if tree.Branch[i] != nil{
			return false
		}
	}
	return true
}


func CreateEntireTree(tree *Tree, layer int){
	CreateBranch(tree)
	for i := 0; i < len(tree.Branch); i++{
		if tree.Branch[i] != nil{
			CreateEntireTree(tree.Branch[i], layer+1)
		}
	}
}

func CopyBoard(dst, src *Tree){
	for i := 0; i < len(dst.Board); i++{
		for j := 0; j < len(dst.Board[i]); j++{
			dst.Board[i][j] = src.Board[i][j]
		}
	}
}

func CreateBranch(tree *Tree){
	var branch [8]*Tree

	for i := 0; i < len(branch); i++{
		branch[i] = new(Tree)
		CopyBoard(branch[i], tree)
	}

	if InTheTable(tree.Row-2, tree.Col-1) &&  IsAvailable(tree, tree.Row-2, tree.Col-1){
		branch[0].Board[tree.Row-2][tree.Col-1] = 1
		branch[0].Row = tree.Row-2
		branch[0].Col = tree.Col-1
	}else{
		branch[0] = nil
	}

	if InTheTable(tree.Row-2, tree.Col+1) &&  IsAvailable(tree, tree.Row-2, tree.Col+1){
		branch[1].Board[tree.Row-2][tree.Col+1] = 1
		branch[1].Row = tree.Row-2
		branch[1].Col = tree.Col+1
	}else{
		branch[1] = nil
	}

	if InTheTable(tree.Row-1, tree.Col+2) &&  IsAvailable(tree, tree.Row-1, tree.Col+2){
		branch[2].Board[tree.Row-1][tree.Col+2] = 1
		branch[2].Row = tree.Row-1
		branch[2].Col = tree.Col+2
	}else{
		branch[2] = nil
	}

	if InTheTable(tree.Row+1, tree.Col+2) &&  IsAvailable(tree, tree.Row+1, tree.Col+2){
		branch[3].Board[tree.Row+1][tree.Col+2] = 1
		branch[3].Row = tree.Row+1
		branch[3].Col = tree.Col+2
	}else{
		branch[3] = nil
	}

	if InTheTable(tree.Row+2, tree.Col+1) &&  IsAvailable(tree, tree.Row+2, tree.Col+1){
		branch[4].Board[tree.Row+2][tree.Col+1] = 1
		branch[4].Row = tree.Row+2
		branch[4].Col = tree.Col+1
	}else{
		branch[4] = nil
	}

	if InTheTable(tree.Row+2, tree.Col-1) &&  IsAvailable(tree, tree.Row+2, tree.Col-1){
		branch[5].Board[tree.Row+2][tree.Col-1] = 1
		branch[5].Row = tree.Row+2
		branch[5].Col = tree.Col-1
	}else{
		branch[5] = nil
	}

	if InTheTable(tree.Row+1, tree.Col-2) &&  IsAvailable(tree, tree.Row+1, tree.Col-2){
		branch[6].Board[tree.Row+1][tree.Col-2] = 1
		branch[6].Row = tree.Row+1
		branch[6].Col = tree.Col-2
	}else{
		branch[6] = nil
	}

	if InTheTable(tree.Row-1, tree.Col-2) &&  IsAvailable(tree, tree.Row-1, tree.Col-2){
		branch[7].Board[tree.Row-1][tree.Col-2] = 1
		branch[7].Row = tree.Row-1
		branch[7].Col = tree.Col-2
	}else{
		branch[7] = nil
	}

	tree.Branch = branch
}

func IsAvailable(tree *Tree, row, col int) bool{
	return tree.Board[row][col] == 0
}

func InTheTable(row, col int) bool{
	return row >= 0 && row < MAX_DIMENSION && col >= 0 && col < MAX_DIMENSION
}



