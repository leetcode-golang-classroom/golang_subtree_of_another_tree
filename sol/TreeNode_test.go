package sol

import (
	"strconv"
	"testing"
)

func CreateBinaryTree(input *[]string) *TreeNode {
	var tree, cur *TreeNode
	hashMap := make(map[int]*TreeNode)
	for idx, val := range *input {
		num := 0
		if val != "null" {
			num, _ = strconv.Atoi(val)
		}
		if idx == 0 {
			tree = &TreeNode{Val: num}
			hashMap[num] = tree
		}
		if node, ok := hashMap[num]; ok {
			cur = node
		} else {
			cur = &TreeNode{Val: num}
			hashMap[num] = cur
		}
		if 2*idx+1 < len(*input) {
			// cur.Left
			if (*input)[2*idx+1] != "null" {
				left_val, _ := strconv.Atoi((*input)[2*idx+1])
				if left, exist := hashMap[left_val]; exist {
					cur.Left = left
				} else {
					left := &TreeNode{Val: left_val}
					hashMap[left_val] = left
					cur.Left = left
				}
			}
		}
		if 2*idx+2 < len(*input) {
			if (*input)[2*idx+2] != "null" {
				right_val, _ := strconv.Atoi((*input)[2*idx+2])
				if right, exist := hashMap[right_val]; exist {
					cur.Right = right
				} else {
					right := &TreeNode{Val: right_val}
					hashMap[right_val] = right
					cur.Right = right
				}
			}
		}
	}
	return tree
}
func Test_isSubtree(t *testing.T) {
	type args struct {
		root    *TreeNode
		subRoot *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "root = [3,4,5,1,2], subRoot = [4,1,2]",
			args: args{root: CreateBinaryTree(&[]string{"3", "4", "5", "1", "2"}),
				subRoot: CreateBinaryTree(&[]string{"4", "1", "2"}),
			},
			want: true,
		},
		{
			name: "root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]",
			args: args{root: CreateBinaryTree(&[]string{"3", "4", "5", "1", "2", "null", "null", "null", "null", "0"}),
				subRoot: CreateBinaryTree(&[]string{"4", "1", "2"}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.root, tt.args.subRoot); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
