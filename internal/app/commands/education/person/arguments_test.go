package person

import (
	"reflect"
	"testing"
)

func Test_splitIntoArguments(t *testing.T) {
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
			[]string{"один", "два' три'"},
			true,
		},
		{`"Вася Пупкин2" "Федя Ножкин`,
			args{`"Вася Пупкин2" "Федя Ножкин`},
			[]string{"Вася Пупкин2", "Федя Ножкин"},
			true,
		},
		{`"some text,"" other text"`,
			args{`"some text,"" other text"`},
			[]string{"some text, other text"},
			false,
		},
		{`Вася" "Иванович" "Крузенштерн`,
			args{`Вася" "Иванович" "Крузенштерн`},
			[]string{"Вася Иванович Крузенштерн"},
			false,
		},
		{`var1="some text" var2='other text'`,
			args{`var1="some text" var2='other text'`},
			[]string{"var1=some text", "var2=other text"},
			false,
		},
		{`var1="some text" var2='other text  `,
			args{`var1="some text" var2='other text  `},
			[]string{"var1=some text", "var2=other text  "},
			true,
		},
		{`var1="some text1" var2='some text2  var3="some text3"`,
			args{`var1="some text1" var2='some text2  var3="some text3"`},
			[]string{"var1=some text1", `var2=some text2  var3="some text3"`},
			true,
		},
		{`var1="some text1" var2='some text2  var3='some text3''`,
			args{`var1="some text1" var2='some text2  var3='some text3'`},
			[]string{"var1=some text1", "var2=some text2  var3=some", "text3"},
			true,
		},

		// TODO: AddField test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := splitIntoArguments(tt.args.text)
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
