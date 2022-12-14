package main

import "testing"

func Test_hasOnlyUniqueCharacters(t *testing.T) {
	type args struct {
		seq string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "abc",
			args: args{
				seq: "abc",
			},
			want: true,
		},
		{
			name: "cbc",
			args: args{
				seq: "cbc",
			},
			want: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasOnlyUniqueCharacters(tt.args.seq); got != tt.want {
				t.Errorf("hasOnlyUniqueCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
