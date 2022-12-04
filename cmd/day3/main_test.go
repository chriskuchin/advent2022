package main

import "testing"

func Test_getItemPriority(t *testing.T) {
	type args struct {
		item rune
	}
	tests := []struct {
		name         string
		args         args
		wantPriority int
	}{
		{
			name: "p",
			args: args{
				item: 'p',
			},
			wantPriority: 16,
		},
		{
			name: "L",
			args: args{
				item: 'L',
			},
			wantPriority: 38,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPriority := getItemPriority(tt.args.item); gotPriority != tt.wantPriority {
				t.Errorf("getItemPriority() = %v, want %v", gotPriority, tt.wantPriority)
			}
		})
	}
}
