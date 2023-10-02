package education

import (
	"reflect"
	"strconv"
	"testing"
)

func TestEducation_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		e       Education
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
			Education(0),
			[]byte(`null`),
			false,
		},
		{
			"Preschool",
			Preschool,
			[]byte(`"Preschool"`),
			false,
		},
		{
			"PrimaryGeneral",
			PrimaryGeneral,
			[]byte(`"PrimaryGeneral"`),
			false,
		},
		{
			"BasicGeneral",
			BasicGeneral,
			[]byte(`"BasicGeneral"`),
			false,
		},
		{
			"SecondaryGeneral",
			SecondaryGeneral,
			[]byte(`"SecondaryGeneral"`),
			false,
		},
		{
			"SecondaryVocational",
			SecondaryVocational,
			[]byte(`"SecondaryVocational"`),
			false,
		},
		{
			"Higher1",
			Higher1,
			[]byte(`"Higher1"`),
			false,
		},
		{
			"Higher2",
			Higher2,
			[]byte(`"Higher2"`),
			false,
		},
		{
			"Higher3",
			Higher3,
			[]byte(`"Higher3"`),
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.e.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestEducation_UnmarshalJSON(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name      string
		e         Education
		args      args
		wantValue Education
		wantErr   bool
	}{
		{
			"unknown text",
			2,
			args{[]byte(`"unknown"`)},
			2,
			true,
		},
		{
			"unknown number",
			2,
			args{[]byte(`100500`)},
			2,
			true,
		},
		{
			"unknown negative number",
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
			"Preschool (case-insensitive)",
			0,
			args{[]byte(`"pReScHoOl"`)},
			Preschool,
			false,
		},
		{
			"Preschool (number)",
			0,
			args{[]byte(strconv.Itoa(int(Preschool)))},
			Preschool,
			false,
		},
		//{
		//	"Preschool (unquoted)", // BROKEN! now `Preschool` -> Preschool
		//	0,
		//	args{[]byte(`Preschool`)},
		//	0,
		//	true,
		//},
		//{
		//	"Preschool (quoted number)", // BROKEN! now `"1"` -> Preschool
		//	0,
		//	args{[]byte(strconv.Quote(strconv.Itoa(int(Preschool))))},
		//	0,
		//	true,
		//},
		{
			"PrimaryGeneral",
			0,
			args{[]byte(`"PrimaryGeneral"`)},
			PrimaryGeneral,
			false,
		},
		{
			"BasicGeneral",
			0,
			args{[]byte(`"BasicGeneral"`)},
			BasicGeneral,
			false,
		},
		{
			"SecondaryGeneral",
			0,
			args{[]byte(`"SecondaryGeneral"`)},
			SecondaryGeneral,
			false,
		},
		{
			"SecondaryVocational",
			0,
			args{[]byte(`"SecondaryVocational"`)},
			SecondaryVocational,
			false,
		},
		{
			"Higher1",
			0,
			args{[]byte(`"Higher1"`)},
			Higher1,
			false,
		},
		{
			"Higher2",
			0,
			args{[]byte(`"Higher2"`)},
			Higher2,
			false,
		},
		{
			"Higher3",
			0,
			args{[]byte(`"Higher3"`)},
			Higher3,
			false,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.e.UnmarshalJSON(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				//t.Logf("UnmarshalJSON() expected error = %v", err)
			}
			if tt.e != tt.wantValue {
				t.Errorf("value = %q, want value %q", tt.e, tt.wantValue)
			}
		})
	}
}
