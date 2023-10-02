package education

import (
	"reflect"
	"testing"
	"time"
)

func TestDate_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		d        Date
		wantText []byte
		wantErr  bool
	}{
		{
			"zero",
			Date(time.Time{}),
			[]byte(`null`),
			false,
		},
		{
			"2001-02-03",
			Date(time.Date(2001, 2, 3, 0, 0, 0, 0, time.UTC)),
			[]byte(`"2001-02-03"`),
			false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotText, err := tt.d.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotText, tt.wantText) {
				t.Errorf("MarshalJSON() gotText = %q, want %q", gotText, tt.wantText)
			}
		})
	}
}

func TestDate_UnmarshalJSON(t *testing.T) {
	type args struct {
		text []byte
	}
	tests := []struct {
		name      string
		d         Date
		args      args
		wantValue Date
		wantErr   bool
	}{
		{
			"zero (empty string)",
			Date(time.Date(2001, 2, 3, 0, 0, 0, 0, time.UTC)),
			args{[]byte(`""`)},
			Date{},
			false,
		},
		{
			"null",
			Date(time.Date(2001, 2, 3, 0, 0, 0, 0, time.UTC)),
			args{[]byte(`null`)},
			Date(time.Date(2001, 2, 3, 0, 0, 0, 0, time.UTC)),
			false,
		},
		{
			"2001-02-03",
			Date{},
			args{[]byte(`"2001-02-03"`)},
			Date(time.Date(2001, 2, 3, 0, 0, 0, 0, time.UTC)),
			false,
		},
		{
			"2001-13-30 (month out of range)",
			Date(time.Date(2001, 2, 3, 0, 0, 0, 0, time.UTC)),
			args{[]byte(`"2001-13-30"`)},
			Date(time.Date(2001, 2, 3, 0, 0, 0, 0, time.UTC)),
			true,
		},
		{
			"2001-02-30 (day out of range)",
			Date(time.Date(2001, 2, 3, 0, 0, 0, 0, time.UTC)),
			args{[]byte(`"2001-02-30"`)},
			Date(time.Date(2001, 2, 3, 0, 0, 0, 0, time.UTC)),
			true,
		},
		{
			"not date text",
			Date(time.Date(2001, 2, 3, 0, 0, 0, 0, time.UTC)),
			args{[]byte(`"not date text"`)},
			Date(time.Date(2001, 2, 3, 0, 0, 0, 0, time.UTC)),
			true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.UnmarshalJSON(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				//t.Logf("UnmarshalJSON() expected error = %v", err)
			}
			if tt.d != tt.wantValue {
				t.Errorf("UnmarshalJSON() value = %v, wantValue %v", tt.d, tt.wantValue)
			}
		})
	}
}
