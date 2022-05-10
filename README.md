# golang_diameter_of_binary_tree

Given the `root` of a binary tree, return *the length of the **diameter** of the tree*.

The **diameter** of a binary tree is the **length** of the longest path between any two nodes in a tree. This path may or may not pass through the `root`.

The **length** of a path between two nodes is represented by the number of edges between them.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2021/03/06/diamtree.jpg](https://assets.leetcode.com/uploads/2021/03/06/diamtree.jpg)

```
Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].

```

**Example 2:**

```
Input: root = [1,2]
Output: 1

```

**Constraints:**

- The number of nodes in the tree is in the range `[1,$10^4$]`.
- `100 <= Node.val <= 100`

## 解析

定義二元樹的 diameter 為樹內兩點距離最遠的點所形成的距離

給定一個二元樹，希望能夠找出這顆樹的Diameter

首先看下圖

![](https://i.imgur.com/fAxCMMK.png)

每個結點 node 的 diameter 有兩種可能 

1  通過 node 

    如果要通過 node 則代表需要可以目前 Max(左子結點樹高, 右子結點樹高) + 2 (左右子結點到 node 的距離)

2  不通過 node

    目前累計最大的距離

而要注意如果是 nil 的結點，定義其 level 是 -1

這樣每次都只要走訪一次該結點就可以知道目前累積的最大距離

## 程式碼
```go
package sol

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0
	FindDiameter(root, &max)
	return max
}
func FindDiameter(root *TreeNode, max *int) int {
	if root == nil {
		return -1 // -1 height
	}
	leftLevel := FindDiameter(root.Left, max)
	rightLevel := FindDiameter(root.Right, max)
	curMax := leftLevel + rightLevel + 2
	*max = Max(*max, curMax)
	return Max(leftLevel, rightLevel) + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
## 困難點

1. 理解 二元樹 diameter 定義
2. 看出 level 對應 diameter的關係 

## Solve Point

- [x]  Understand what problem need to be solve
- [x]  Analysis Complexity