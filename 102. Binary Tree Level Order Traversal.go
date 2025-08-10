package leetcode_go

/*
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).



Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
Example 2:

Input: root = [1]
Output: [[1]]
Example 3:

Input: root = []
Output: []
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
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
		res = append(res, a)
		queue = queue[size:]
	}
	return res
}
