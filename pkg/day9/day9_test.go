package main

import "testing"

func TestFind(t *testing.T) {
	type args struct {
		numbers   []int
		blockSize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "works for default",
			args: args{
				numbers: []int{
					35,
					20,
					15,
					25,
					47,
					40,
					62,
					55,
					65,
					95,
					102,
					117,
					150,
					182,
					127,
					219,
					299,
					277,
					309,
					576,
				},
				blockSize: 5,
			},
			want: 127,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.numbers, tt.args.blockSize); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindWeakness(t *testing.T) {
	type args struct {
		search  int
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "works for default - part 2",
			args: args{
				search: 127,
				numbers: []int{
					35,
					20,
					15,
					25,
					47,
					40,
					62,
					55,
					65,
					95,
					102,
					117,
					150,
					182,
					127,
					219,
					299,
					277,
					309,
					576,
				},
			},
			want: 62,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindWeakness(tt.args.search, tt.args.numbers); got != tt.want {
				t.Errorf("FindWeakness() = %v, want %v", got, tt.want)
			}
		})
	}
}
