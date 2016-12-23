package tunnel

import (
	"testing"
)

func TestMagicWell(t *testing.T) {
	var tests = []struct {
		a   int
		b   int
		c   int
		exp int
	}{
		{1, 2, 2, 8},
		{1, 1, 1, 1},
		{6, 5, 3, 128},
	}

	for _, test := range tests {
		if actual := magicalWell(test.a, test.b, test.c); actual != test.exp {
			t.Errorf("magicWell(%d, %d, %d) = %v, expected %v", test.a, test.b,
				test.c, actual, test.exp)
		}
	}
}

func BenchmarkMagicWell(b *testing.B) {
	for i := 0; i < b.N; i++ {
		magicalWell(1, 2, 1000)
	}
}

func Test_countSumOfTwoRepresentations2(t *testing.T) {
	type args struct {
		n int
		l int
		r int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{6, 2, 4}, 2},
		{"2", args{6, 3, 3}, 1},
		{"3", args{10, 9, 11}, 0},
		{"4", args{24, 8, 16}, 5},
		{"5", args{24, 12, 12}, 1},
		{"6", args{93, 24, 58}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSumOfTwoRepresentations2(tt.args.n, tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("countSumOfTwoRepresentations3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_countSumOfTwoRepresentations2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countSumOfTwoRepresentations2(100, 2, 120)
	}
}

func Test_leastFactorial(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{17}, 24},
		{"", args{1}, 1},
		{"", args{5}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leastFactorial(tt.args.n); got != tt.want {
				t.Errorf("leastFactorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
