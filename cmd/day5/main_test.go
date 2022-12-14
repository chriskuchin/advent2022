package main

import (
	"reflect"
	"testing"
)

func Test_pop(t *testing.T) {
	type args struct {
		src []string
	}
	tests := []struct {
		name     string
		args     args
		wantVal  string
		wantList []string
	}{
		{
			name: "pop",
			args: args{
				src: []string{"a", "b", "c"},
			},
			wantVal:  "c",
			wantList: []string{"a", "b"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVal, gotList := pop(tt.args.src)
			if gotVal != tt.wantVal {
				t.Errorf("pop() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
			if !reflect.DeepEqual(gotList, tt.wantList) {
				t.Errorf("pop() gotList = %v, want %v", gotList, tt.wantList)
			}
		})
	}
}

func Test_popMultiple(t *testing.T) {
	type args struct {
		src   []string
		count int
	}
	tests := []struct {
		name     string
		args     args
		wantVal  []string
		wantList []string
	}{
		{
			name: "pop",
			args: args{
				src:   []string{"a", "b", "c", "d"},
				count: 2,
			},
			wantVal:  []string{"c", "d"},
			wantList: []string{"a", "b"},
		},
		{
			name: "pop",
			args: args{
				src:   []string{"a", "b", "c", "d"},
				count: 1,
			},
			wantVal:  []string{"d"},
			wantList: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVal, gotList := popMultiple(tt.args.src, tt.args.count)
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("popMultiple() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
			if !reflect.DeepEqual(gotList, tt.wantList) {
				t.Errorf("popMultiple() gotList = %v, want %v", gotList, tt.wantList)
			}
		})
	}
}
