package main

import (
	"reflect"
	"testing"
)

func TestParseInstruction(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want Instruction
	}{
		{
			name: "passes nop",
			args: args{
				str: "nop +0",
			},
			want: Instruction{
				command: "nop",
				unit:    0,
				symbol:  "+",
			},
		},
		{
			name: "passes acc",
			args: args{
				str: "acc +1",
			},
			want: Instruction{
				command: "acc",
				unit:    1,
				symbol:  "+",
			},
		},
		{
			name: "passes jmp",
			args: args{
				str: "jmp -4",
			},
			want: Instruction{
				command: "jmp",
				unit:    4,
				symbol:  "-",
			},
		},
		{
			name: "passes jmp (longer number)",
			args: args{
				str: "jmp -487",
			},
			want: Instruction{
				command: "jmp",
				unit:    487,
				symbol:  "-",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseInstruction(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseInstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadInstructions(t *testing.T) {
	type args struct {
		instructions []Instruction
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name: "works with example input",
			args: args{
				instructions: []Instruction{
					{command: "nop", unit: 0, symbol: "+"},
					{command: "acc", unit: 1, symbol: "+"},
					{command: "jmp", unit: 4, symbol: "+"},
					{command: "acc", unit: 3, symbol: "+"},
					{command: "jmp", unit: 3, symbol: "-"},
					{command: "acc", unit: 99, symbol: "-"},
					{command: "acc", unit: 1, symbol: "+"},
					{command: "jmp", unit: 4, symbol: "-"},
					{command: "acc", unit: 6, symbol: "+"},
				}},
			want:  5,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindPath(tt.args.instructions)
			if got != tt.want {
				t.Errorf("FindPath() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindPath() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFindPath(t *testing.T) {
	type args struct {
		instructions []Instruction
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name: "works with example input",
			args: args{
				instructions: []Instruction{
					{command: "nop", unit: 0, symbol: "+"},
					{command: "acc", unit: 1, symbol: "+"},
					{command: "jmp", unit: 4, symbol: "+"},
					{command: "acc", unit: 3, symbol: "+"},
					{command: "jmp", unit: 3, symbol: "-"},
					{command: "acc", unit: 99, symbol: "-"},
					{command: "acc", unit: 1, symbol: "+"},
					{command: "jmp", unit: 4, symbol: "-"},
					{command: "acc", unit: 6, symbol: "+"},
				}},
			want:  8,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindPath(tt.args.instructions)
			if got != tt.want {
				t.Errorf("FindPath() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindPath() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
