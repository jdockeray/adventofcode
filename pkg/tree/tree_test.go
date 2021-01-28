package tree

import (
	"reflect"
	"testing"
)

func TestBuildTree(t *testing.T) {
	type args struct {
		numbs []int
	}
	tests := []struct {
		name string
		args args
		want *Tree
	}{
		{
			name: "builds tree",
			args: args{numbs: []int{1}},
			want: &Tree{
				Value:    1,
				children: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildTree(tt.args.numbs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountLeaves(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "count leaves",
			args: args{
				t: BuildTree([]int{1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, 22}),
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountLeaves(tt.args.t); got != tt.want {
				t.Errorf("CountLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
}
