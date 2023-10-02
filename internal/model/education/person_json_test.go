package education

import (
	"reflect"
	"testing"
	"time"
)

func TestPerson_MarshalJSON(t *testing.T) {
	type fields struct {
		ID         uint64
		FirstName  string
		MiddleName string
		LastName   string
		Birthday   time.Time
		Sex        Sex
		Education  Education
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			"zero middle name",
			fields{
				ID:        314,
				FirstName: "Вася",
				LastName:  "Пупкин",
				Birthday:  time.Date(2001, 02, 03, 0, 0, 0, 0, time.UTC),
				Sex:       Male,
				Education: BasicGeneral,
			},
			[]byte(`{"id":314,"first_name":"Вася","last_name":"Пупкин","birthday":"2001-02-03","sex":"Male","education":"BasicGeneral"}`),
			false,
		},
		{
			"zero birthday",
			fields{
				ID:        314,
				FirstName: "Вася",
				LastName:  "Пупкин",
				Sex:       Male,
				Education: BasicGeneral,
			},
			[]byte(`{"id":314,"first_name":"Вася","last_name":"Пупкин","sex":"Male","education":"BasicGeneral"}`),
			false,
		},
		{
			"zero sex",
			fields{
				ID:        314,
				FirstName: "Вася",
				LastName:  "Пупкин",
				Education: BasicGeneral,
			},
			[]byte(`{"id":314,"first_name":"Вася","last_name":"Пупкин","education":"BasicGeneral"}`),
			false,
		},
		{
			"zero education",
			fields{
				ID:        314,
				FirstName: "Вася",
				LastName:  "Пупкин",
			},
			[]byte(`{"id":314,"first_name":"Вася","last_name":"Пупкин"}`),
			false,
		},
		{
			"unknown education",
			fields{
				ID:        314,
				FirstName: "Вася",
				LastName:  "Пупкин",
				Education: 100500,
			},
			nil,
			true,
		},
		{
			"unknown sex",
			fields{
				ID:        314,
				FirstName: "Вася",
				LastName:  "Пупкин",
				Sex:       3,
			},
			nil,
			true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Person{
				ID:         tt.fields.ID,
				FirstName:  tt.fields.FirstName,
				MiddleName: tt.fields.MiddleName,
				LastName:   tt.fields.LastName,
				Birthday:   tt.fields.Birthday,
				Sex:        tt.fields.Sex,
				Education:  tt.fields.Education,
			}
			got, err := p.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				//return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %q,\n want %q", got, tt.want)
			}
		})
	}
}
