package person

import (
	"reflect"
	"testing"
)

func Test_parseArguments(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"1 22 333 4444",
			args{"1 22 333 4444"},
			[]string{"1", "22", "333", "4444"},
			false,
		},
		{" 1 2  3   4  ",
			args{" 1 2 \t 3   \n4  "},
			[]string{"1", "2", "3", "4"},
			false,
		},
		{"empty string",
			args{""},
			nil,
			false,
		},
		{"   ",
			args{"   "},
			nil,
			false,
		},
		{"1",
			args{"1"},
			[]string{"1"},
			false,
		},
		{"один два три",
			args{"один два три"},
			[]string{"один", "два", "три"},
			false,
		},
		{"один два три\n",
			args{"один два три\n"},
			[]string{"один", "два", "три"},
			false,
		},
		{"'один два' три\n",
			args{"'один два' три\n"},
			[]string{"один два", "три"},
			false,
		},
		{`один "два' три"`,
			args{`один "два' три"`},
			[]string{"один", "два' три"},
			false,
		},
		{`один "два' три'`,
			args{`один "два' три'`},
			nil,
			true,
		},
		{`"Вася Пупкин2" "Федя Ножкин`,
			args{`"Вася Пупкин2" "Федя Ножкин`},
			nil,
			true,
		},
		//{`"some text,"" other text"`,
		//	args{`"some text,"" other text"`},
		//	[]string{"some text, other text"},
		//	false,
		//},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseArguments(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitIntoWords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitIntoWords() got = %q, want %q", got, tt.want)
			}
		})
	}
}
