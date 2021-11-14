package strings

import "testing"

func TestSubString(t *testing.T) {
	s := "ストレージの空123456き容量が足りないため"
	type args struct {
		str       string
		startIdx  uint
		maxLength uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"start over string length",
			args{
				str:       s,
				startIdx:  255,
				maxLength: 10,
			},
			"",
		},
		{
			"from start to middle",
			args{
				str:       s,
				startIdx:  0,
				maxLength: 10,
			},
			"ストレージの空123",
		},
		{
			"from start to end",
			args{
				str:       s,
				startIdx:  0,
				maxLength: 23,
			},
			"ストレージの空123456き容量が足りないため",
		},
		{
			"from start to over length",
			args{
				str:       s,
				startIdx:  0,
				maxLength: 255,
			},
			"ストレージの空123456き容量が足りないため",
		},
		{
			"from middle to end",
			args{
				str:       s,
				startIdx:  5,
				maxLength: 23,
			},
			"の空123456き容量が足りないため",
		},
		{
			"from middle to over length",
			args{
				str:       s,
				startIdx:  5,
				maxLength: 255,
			},
			"の空123456き容量が足りないため",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubString(tt.args.str, tt.args.startIdx, tt.args.maxLength); got != tt.want {
				t.Errorf("SubString() = %v, want %v", got, tt.want)
			}
		})
	}
}
