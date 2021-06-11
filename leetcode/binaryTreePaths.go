package leetcode

import "strconv"

//给定一个二叉树，返回所有从根节点到叶子节点的路径。
//
//说明: 叶子节点是指没有子节点的节点。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var strs []string

func binaryTreePaths(root *TreeNode) []string {
	strs = []string{}
	if root == nil {
		return strs
	}
	getTreeVal(root, "")
	return strs
}

func getTreeVal(root *TreeNode, path string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		strs = append(strs, path+strconv.Itoa(root.Val))
		return
	}

	getTreeVal(root.Left, path+strconv.Itoa(root.Val)+"->")
	getTreeVal(root.Right, path+strconv.Itoa(root.Val)+"->")
	return

}
