package code

import (
	"reflect"
	"testing"
)

func Test_mergeTrees(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{"case 1", args{&TreeNode{1, &TreeNode{3, &TreeNode{5, nil, nil}, nil}, &TreeNode{2, nil, nil}}, &TreeNode{2, &TreeNode{1, nil, &TreeNode{4, nil, nil}}, &TreeNode{3, nil, &TreeNode{7, nil, nil}}}}, &TreeNode{3, &TreeNode{4, &TreeNode{5, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{5, nil, &TreeNode{7, nil, nil}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTrees(tt.args.root1, tt.args.root2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
