package code

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"1", args{&ListNode{1, &ListNode{2, &ListNode{3, nil}}}}, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
