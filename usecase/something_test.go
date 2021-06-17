package usecase

import (
	"reflect"
	"testing"
)

func TestNewSomething(t *testing.T) {
	type args struct {
		s SomeService
	}
	tests := []struct {
		name string
		args args
		want *Something
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSomething(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSomething() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomething_DoSomething(t *testing.T) {
	type fields struct {
		SomeService SomeService
	}
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &Something{
				SomeService: tt.fields.SomeService,
			}
			got, err := uc.DoSomething(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Something.DoSomething() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Something.DoSomething() = %v, want %v", got, tt.want)
			}
		})
	}
}
