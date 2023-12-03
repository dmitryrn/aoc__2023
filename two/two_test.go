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
		want game
	}{
        {
            name: "1",
            args: args{
                s: "Game 10: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
            },
            want: game{
                id: 10,
                possible: true,
            },
        },
        {
            name: "2",
            args: args{
                s: "Game 10: 33 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
            },
            want: game{
                id: 10,
                possible: false,
            },
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseGame(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
