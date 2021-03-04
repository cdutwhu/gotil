package gotil

import "testing"

func TestIndexOf(t *testing.T) {
	type args struct {
		e   interface{}
		set []interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{1, []interface{}{5, 4, 3, 2, 1}},
			want: 4,
		},
		{
			name: "OK",
			args: args{"1", []interface{}{5, 4, 3, 2, 1}},
			want: -1,
		},
		{
			name: "OK",
			args: args{"1", []interface{}{"1", "5", "4", "3", "2", "1"}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexOf(tt.args.e, tt.args.set...); got != tt.want {
				t.Errorf("IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastIndexOf(t *testing.T) {
	type args struct {
		e   interface{}
		set []interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{1, []interface{}{5, 4, 3, 2, 1}},
			want: 4,
		},
		{
			name: "OK",
			args: args{"1", []interface{}{5, 4, 3, 2, 1}},
			want: -1,
		},
		{
			name: "OK",
			args: args{"1", []interface{}{"1", "5", "4", "3", "2", "1"}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastIndexOf(tt.args.e, tt.args.set...); got != tt.want {
				t.Errorf("LastIndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
