package main

import "testing"

func Test_isOverlapping(t *testing.T) {
	type args struct {
		elf1Min int
		elf1Max int
		elf2Min int
		elf2Max int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "2-4,6-8",
			args: args{
				elf1Min: 2,
				elf1Max: 4,
				elf2Min: 6,
				elf2Max: 8,
			},
			want: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOverlapping(tt.args.elf1Min, tt.args.elf1Max, tt.args.elf2Min, tt.args.elf2Max); got != tt.want {
				t.Errorf("isOverlapping() = %v, want %v", got, tt.want)
			}
		})
	}
}
