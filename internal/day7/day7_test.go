package day7

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want BagWrapper
	}{
		{
			name: "works",
			args: args{
				line: "shiny cyan bags contain 4 plaid green bags, 4 dim coral bags, 4 dull indigo bags.\n",
			},
			want: BagWrapper{
				children: []Bag{
					{
						count: 4,
						name:  "plaid-green",
					}, {
						count: 4,
						name:  "dim-coral",
					},
					{
						count: 4,
						name:  "dull-indigo",
					},
				},
				name: "shiny-cyan",
			},
		},
		{
			name: "works",
			args: args{
				line: "mirrored gold bags contain no other bags.\n",
			},
			want: BagWrapper{
				children: make([]Bag, 0, 10),
				name:     "mirrored-gold",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseLine(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBagMachine_Contains(t *testing.T) {
	type fields struct {
		bags map[string]BagWrapper
	}
	type args struct {
		bagWrapper BagWrapper
		slug       string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "does contain slug - directly",
			fields: fields{
				bags: map[string]BagWrapper{
					"plaid-green": {
						children: make([]Bag, 0, 10),
						name:     "plaid-green",
					},
				},
			},
			args: args{
				bagWrapper: BagWrapper{
					children: []Bag{
						{
							count: 4,
							name:  "plaid-green",
						},
					},
					name: "shiny-cyan",
				},
				slug: "plaid-green",
			},
			want: true,
		},
		{
			name: "does contain slug - indirectly",
			fields: fields{
				bags: map[string]BagWrapper{
					"plaid-green": {
						children: make([]Bag, 0, 10),
						name:     "plaid-green",
					},
					"dull-indigo": {
						children: []Bag{
							{
								count: 4,
								name:  "plaid-green",
							},
						},
						name: "shiny-cyan",
					},
				},
			},
			args: args{
				bagWrapper: BagWrapper{
					children: []Bag{
						{
							count: 4,
							name:  "dull-indigo",
						},
					},
					name: "shiny-cyan",
				},
				slug: "plaid-green",
			},
			want: true,
		},
		{
			name: "does contain slug - indirectly (2 levels deep)",
			fields: fields{
				bags: map[string]BagWrapper{
					"plaid-green": {
						children: make([]Bag, 0, 10),
						name:     "plaid-green",
					},
					"magic-red": {
						children: []Bag{
							{
								count: 4,
								name:  "dull-indigo",
							},
						},
						name: "magic-red",
					},
					"dull-indigo": {
						children: []Bag{
							{
								count: 4,
								name:  "plaid-green",
							},
						},
						name: "shiny-cyan",
					},
					"magic-blue": {
						children: make([]Bag, 0, 10),
						name:     "magic-blue",
					},
				},
			},

			args: args{
				bagWrapper: BagWrapper{
					children: []Bag{
						{
							count: 4,
							name:  "magic-blue",
						},
						{
							count: 4,
							name:  "magic-red",
						},
					},
					name: "shiny-cyan",
				},
				slug: "plaid-green",
			},
			want: true,
		},

		{
			name: "does not contain slug",
			fields: fields{
				bags: map[string]BagWrapper{
					"plaid-green": {
						children: make([]Bag, 0, 10),
						name:     "plaid-green",
					},
				},
			},
			args: args{
				bagWrapper: BagWrapper{
					children: []Bag{
						{
							count: 4,
							name:  "plaid-green",
						},
					},
					name: "shiny-cyan",
				},
				slug: "does-not-exist",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			machine := &BagMachine{
				bags: tt.fields.bags,
			}
			if got := machine.Contains(tt.args.bagWrapper, tt.args.slug); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBagMachine_Count(t *testing.T) {
	type fields struct {
		bags map[string]BagWrapper
	}
	type args struct {
		slug string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "counts the number of bags",
			fields: fields{
				bags: map[string]BagWrapper{
					"plaid-green": {
						children: make([]Bag, 0, 10),
						name:     "plaid-green",
					},
					"magic-red": {
						children: []Bag{
							{
								count: 4,
								name:  "dull-indigo",
							},
						},
						name: "magic-red",
					},
					"dull-indigo": {
						children: []Bag{
							{
								count: 4,
								name:  "plaid-green",
							},
						},
						name: "dull-indigo",
					},
				},
			},
			args: args{slug: "magic-red"},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			machine := &BagMachine{
				bags: tt.fields.bags,
			}
			if got := machine.Count(tt.args.slug); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
