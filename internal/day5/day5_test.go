package day5

import (
	"reflect"
	"testing"
)

func TestGetId(t *testing.T) {
	type args struct {
		rowCol string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"works",
			args{
				rowCol: "BFFFBBFRRR",
			},
			567,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetId(tt.args.rowCol); got != tt.want {
				t.Errorf("GetId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseRowCols(t *testing.T) {
	type args struct {
		rowcols string
	}
	tests := []struct {
		name     string
		args     args
		wantRow  string
		wantCols string
	}{
		{
			name: "works",
			args: args{
				rowcols: "BFFFBBFRRR",
			},
			wantRow:  "BFFFBBF",
			wantCols: "RRR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRow, gotCols := ParseRowCols(tt.args.rowcols)
			if gotRow != tt.wantRow {
				t.Errorf("ParseRowCols() gotRow = %v, want %v", gotRow, tt.wantRow)
			}
			if gotCols != tt.wantCols {
				t.Errorf("ParseRowCols() gotCols = %v, want %v", gotCols, tt.wantCols)
			}
		})
	}
}

func TestConvertToBinary(t *testing.T) {
	type args struct {
		z   string
		o   string
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "works",
			args: args{
				z:   "F",
				o:   "B",
				str: "BFFFBBF",
			},
			want: "1000110",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToBinary(tt.args.z, tt.args.o)(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
