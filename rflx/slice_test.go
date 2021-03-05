package rflx

import (
	"reflect"
	"testing"
)

func TestSliceAttach(t *testing.T) {
	type args struct {
		s1  interface{}
		s2  interface{}
		pos int
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "OK",
			args: args{
				s1:  []string{},
				s2:  []string{"A"},
				pos: 10,
			},
			want: []string{"A"},
		},
		{
			name: "OK",
			args: args{
				s1:  []string{"a", "b", "c"},
				s2:  []string{"B"},
				pos: 1,
			},
			want: []string{"a", "B", "c"},
		},
		{
			name: "OK",
			args: args{
				s1:  []string{"a", "b", "c"},
				s2:  []string{"D"},
				pos: 3,
			},
			want: []string{"a", "b", "c", "D"},
		},
		{
			name: "OK",
			args: args{
				s1:  []string{"a", "b", "c"},
				s2:  []rune{'E'},
				pos: 5,
			},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceAttach(tt.args.s1, tt.args.s2, tt.args.pos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceAttach() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceCover(t *testing.T) {
	type args struct {
		ss []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{ss: []interface{}{[]string{"a", "b", "c"}, []string{"A", "B", "C"}, []string{"1", "2", "3"}}},
			want: []string{"1", "2", "3"},
		},
		{
			name: "OK",
			args: args{ss: []interface{}{[]string{"a", "b", "c"}, []string{"A", "B"}, []string{"1"}}},
			want: []string{"1", "B", "c"},
		},
		{
			name: "OK",
			args: args{ss: []interface{}{[]string{"a", "b", "c"}, []string{"A", "B"}, []string{"1", "2", "3", "4"}}},
			want: []string{"1", "2", "3", "4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceCover(tt.args.ss...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceCover() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToGSlc(t *testing.T) {
	type args struct {
		slc interface{}
	}
	tests := []struct {
		name     string
		args     args
		wantGSlc []interface{}
	}{
		// TODO: Add test cases.
		{
			name:     "OK",
			args:     args{slc: []string{"abc", "def"}},
			wantGSlc: []interface{}{"abc", "def"},
		},
		{
			name:     "OK",
			args:     args{slc: []int{1}},
			wantGSlc: []interface{}{1},
		},
		{
			name:     "OK",
			args:     args{slc: []int{}},
			wantGSlc: []interface{}{},
		},
		{
			name:     "OK",
			args:     args{slc: nil},
			wantGSlc: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotGSlc := ToGSlc(tt.args.slc); !reflect.DeepEqual(gotGSlc, tt.wantGSlc) {
				t.Errorf("ToGSlc() = %v, want %v", gotGSlc, tt.wantGSlc)
			}
		})
	}
}

func TestToTSlc(t *testing.T) {
	type args struct {
		gSlc []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{gSlc: []interface{}{"a", "b"}},
			want: []string{"a", "b"},
		},
		{
			name: "OK",
			args: args{gSlc: []interface{}{1, 2}},
			want: []int{1, 2},
		},
		{
			name: "OK",
			args: args{gSlc: []interface{}{1, 2.01}},
			want: []float64{1, 2.01},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToTSlc(tt.args.gSlc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GSlc2Slc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSet(t *testing.T) {
	type args struct {
		slc interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{[]interface{}{1, "1", 2, "1", 3, 1, 2, 3}},
			want: []interface{}{1, "1", 2, 3},
		},
		{
			name: "OK",
			args: args{[]int{1, 2, 3, 1}},
			want: []int{1, 2, 3},
		},
		{
			name: "OK",
			args: args{[]int{}},
			want: []int{},
		},
		{
			name: "OK",
			args: args{nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSet(tt.args.slc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKind(t *testing.T) {

	// obj := 1
	// obj := "abc"
	// obj := []int{1, 2}
	// obj := []interface{}{1, 2}
	// obj := struct{ a, b int }{1, 2}

	type P struct {
		a int
		b int
	}
	obj := P{1, 2}

	fPln(reflect.ValueOf(obj))
	fPln(reflect.ValueOf(obj).Kind())
	fPln(reflect.ValueOf(obj).Type().Kind())
	fPln(reflect.TypeOf(obj))
	fPln(reflect.TypeOf(obj).Kind())
}
