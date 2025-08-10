package leetcode_go

/*
Given the root of a binary tree, return the bottom-up level order traversal of its nodes' values. (i.e., from left to right, level by level from leaf to root).
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		var a []int
		size := len(queue)
		for i := 0; i < size; i++ {
			top := queue[i]

			if top.Left != nil {
				queue = append(queue, top.Left)
			}

			if top.Right != nil {
				queue = append(queue, top.Right)
			}

			a = append(a, top.Val)
		}
		res = append([][]int{a}, res...)
		queue = queue[size:]
	}
	return res
}
