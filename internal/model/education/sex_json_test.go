package education

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSex_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		sx      Sex
		want    []byte
		wantErr bool
	}{
		{
			"unknown",
			100500,
			nil,
			true,
		},
		{
			"unknown (negative)",
			-100500,
			nil,
			true,
		},
		{
			"zero",
			0,
			[]byte(`null`),
			false,
		},
		{
			"Female",
			Female,
			[]byte(`"Female"`),
			false,
		},
		{
			"Male",
			Male,
			[]byte(`"Male"`),
			false,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.sx.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				//t.Logf("UnmarshalJSON() expected error = %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSex_UnmarshalJSON(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name      string
		sx        Sex
		args      args
		wantValue Sex
		wantErr   bool
	}{
		{
			"unknown (text)",
			2,
			args{[]byte(`"unknown"`)},
			2,
			true,
		},
		{
			"unknown (number)",
			2,
			args{[]byte(`100500`)},
			2,
			true,
		},
		{
			"unknown (negative number)",
			2,
			args{[]byte(`-100500`)},
			2,
			true,
		},
		{
			"zero (number)",
			2,
			args{[]byte(`0`)},
			0,
			false,
		},
		{
			"zero (quoted number)",
			2,
			args{[]byte(`"0"`)},
			2,
			true,
		},
		{
			"zero (empty string)",
			2,
			args{[]byte(`""`)},
			0,
			false,
		},
		{
			"null",
			2,
			args{[]byte(`null`)},
			2,
			false,
		},
		{
			"Female (case-insensitive)",
			0,
			args{[]byte(`"fEmAle"`)},
			Female,
			false,
		},
		{
			"Female (number)",
			0,
			args{[]byte(strconv.Itoa(int(Female)))},
			Female,
			false,
		},
		//{
		//	"Female (unquoted)", // BROKEN! now `Female` -> Female
		//	0,
		//	args{[]byte(`Female`)},
		//	0,
		//	true,
		//},
		//{
		//	"Female (quoted number)", // BROKEN! now `"1"` -> Female
		//	0,
		//	args{[]byte(strconv.Quote(strconv.Itoa(int(Female))))},
		//	0,
		//	true,
		//},
		{
			"Female (bad unquoted)",
			0,
			args{[]byte(`"Female`)},
			0,
			true,
		},
		{
			"Male",
			0,
			args{[]byte(`"Male"`)},
			Male,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.sx.UnmarshalJSON(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				//t.Logf("UnmarshalJSON() expected error = %v", err)
			}
			if tt.sx != tt.wantValue {
				t.Errorf("value = %v, want value %v", tt.sx, tt.wantValue)
			}
		})
	}
}
