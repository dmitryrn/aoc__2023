package main

import "testing"

func Test_isAdjacentToSymbol(t *testing.T) {
	type args struct {
		matrix [][]rune
		startX int
		startY int
		length int
	}

	m1 := inputToMatrix(`
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
                `)

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				matrix: m1,
				startX: 0,
				startY: 0,
				length: 2,
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				matrix: m1,
				startX: 5,
				startY: 0,
				length: 2,
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				matrix: m1,
				startX: 2,
				startY: 2,
				length: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, got := isAdjacentToSymbol(tt.args.matrix, tt.args.startX, tt.args.startY, tt.args.length); got != tt.want {
				t.Errorf("isAdjacentToSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}
