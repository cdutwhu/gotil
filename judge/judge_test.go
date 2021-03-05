package judge

import (
	"math"
	"testing"
)

func TestIsNumeric(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{s: ".234"},
			want: true,
		},
		{
			name: "OK",
			args: args{s: "1.234"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumeric(tt.args.s); got != tt.want {
				t.Errorf("IsNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsContInts(t *testing.T) {
	type args struct {
		ints []int
	}
	tests := []struct {
		name    string
		args    args
		wantOk  bool
		wantMin int
		wantMax int
	}{
		// TODO: Add test cases.
		{
			name:    "OK",
			args:    args{ints: []int{1, 2, 3, 4, 5}},
			wantOk:  true,
			wantMin: 1,
			wantMax: 5,
		},
		{
			name:    "OK",
			args:    args{ints: []int{5, 4, 3, 2, 1, 0, -1}},
			wantOk:  true,
			wantMin: -1,
			wantMax: 5,
		},
		{
			name:    "OK",
			args:    args{ints: []int{5, 3, 2, 1, 0, -1}},
			wantOk:  false,
			wantMin: -1,
			wantMax: 5,
		},
		{
			name:    "OK",
			args:    args{ints: []int{}},
			wantOk:  false,
			wantMin: math.MinInt32,
			wantMax: math.MaxInt32,
		},
		{
			name:    "OK",
			args:    args{ints: []int{1}},
			wantOk:  true,
			wantMin: 1,
			wantMax: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOk, gotMin, gotMax := IsContInts(tt.args.ints)
			if gotOk != tt.wantOk {
				t.Errorf("IsContInts() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
			if gotMin != tt.wantMin {
				t.Errorf("IsContInts() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
			if gotMax != tt.wantMax {
				t.Errorf("IsContInts() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
