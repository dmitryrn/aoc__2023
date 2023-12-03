package main

import (
	"reflect"
	"testing"
)

func Test_parseGame(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			},
			want: 48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPower(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
