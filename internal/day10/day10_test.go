package day10

import "testing"

func TestCalculate(t *testing.T) {
	type args struct {
		adapters []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "works with sample",
			args: args{
				adapters: []int{
					16,
					10,
					15,
					5,
					1,
					11,
					7,
					19,
					6,
					12,
					4,
				},
			},
			want: 35,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate(tt.args.adapters); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculate2(t *testing.T) {
	type args struct {
		adapters []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "works with sample",
			args: args{
				adapters: []int{
					16,
					10,
					15,
					5,
					1,
					11,
					7,
					19,
					6,
					12,
					4,
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Calculate2(tt.args.adapters); got != tt.want {
				t.Errorf("Calculate2() = %v, want %v", got, tt.want)
			}
		})
	}
}
