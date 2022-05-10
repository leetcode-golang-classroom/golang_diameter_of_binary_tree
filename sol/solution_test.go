package sol

import (
	"testing"
)

func CreateBinaryTree(input *[]int) *TreeNode {
	var tree, cur *TreeNode
	hashMap := make(map[int]*TreeNode)
	for idx, val := range *input {
		if idx == 0 {
			tree = &TreeNode{Val: val}
			hashMap[val] = tree
		}
		if node, ok := hashMap[val]; ok {
			cur = node
		} else {
			cur = &TreeNode{Val: val}
			hashMap[val] = cur
		}
		if 2*idx+1 < len(*input) {
			// cur.Left
			left_val := (*input)[2*idx+1]
			if left, exist := hashMap[left_val]; exist {
				cur.Left = left
			} else {
				left := &TreeNode{Val: left_val}
				hashMap[left_val] = left
				cur.Left = left
			}

		}
		if 2*idx+2 < len(*input) {
			right_val := (*input)[2*idx+2]
			if right, exist := hashMap[right_val]; exist {
				cur.Right = right
			} else {
				right := &TreeNode{Val: right_val}
				hashMap[right_val] = right
				cur.Right = right
			}

		}
	}
	return tree
}
func Test_diameterOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "root = [1,2,3,4,5]",
			args: args{root: CreateBinaryTree(&[]int{1, 2, 3, 4, 5})},
			want: 3,
		},
		{
			name: "root = [1,2]",
			args: args{root: CreateBinaryTree(&[]int{1, 2})},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
