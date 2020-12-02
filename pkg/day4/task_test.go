package main

import "testing"

var valid = Passport{
	Byr: 1980,
	Iyr: 2012,
	Eyr: 2030,
	Hgt: "74in",
	Hcl: "#623a2f",
	Ecl: "grn",
	Pid: "087499704",
}

func TestPassportManager_validate(t *testing.T) {
	ValidateSetUp()

	type fields struct {
		passports []Passport
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "fails on invalid pid",
			fields: fields{
				passports: []Passport{valid.WithPid("0215724")},
			},
			want: 0,
		},
		{
			name: "fails on invalid expiration year",
			fields: fields{
				passports: []Passport{valid.WithEyr(1972)},
			},
			want: 0,
		},
		{
			name: "fails on invalid hair color",
			fields: fields{
				[]Passport{valid.WithHcl("dab227")},
			},
			want: 0,
		},
		{
			name: "fails on invalid eye color",
			fields: fields{
				passports: []Passport{valid.WithEcl("horse")},
			},
			want: 0,
		},
		{
			name: "fails on invalid birth year (min)",
			fields: fields{
				passports: []Passport{valid.WithByr(1912)},
			},
			want: 0,
		},
		{
			name: "fails on invalid birth year (max)",
			fields: fields{
				passports: []Passport{valid.WithByr(2022)},
			},
			want: 0,
		},
		{
			name: "fails on invalid hgt (max cm)",
			fields: fields{
				passports: []Passport{valid.WithHgt("194cm")},
			},
			want: 0,
		},
		{
			name: "fails on invalid hgt (min cm)",
			fields: fields{
				passports: []Passport{valid.WithHgt("149cm")},
			},
			want: 0,
		},
		{
			name: "fails on invalid hgt (max in)",
			fields: fields{
				passports: []Passport{valid.WithHgt("77in")},
			},
			want: 0,
		},
		{
			name: "fails on invalid hgt (min in)",
			fields: fields{
				passports: []Passport{valid.WithHgt("58in")},
			},
			want: 0,
		},
		{
			name: "passes on valid",
			fields: fields{
				passports: []Passport{valid},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := &PassportManager{
				passports: tt.fields.passports,
			}
			if got := receiver.validate(); got != tt.want {
				t.Errorf("validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
