package service

import (
	"reflect"
	"testing"
)

func TestNewSomeService(t *testing.T) {
	type args struct {
		param string
	}
	tests := []struct {
		name string
		args args
		want *SomeService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSomeService(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSomeService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomeService_GetParam(t *testing.T) {
	type fields struct {
		Param string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SomeService{
				Param: tt.fields.Param,
			}
			if got := s.GetParam(); got != tt.want {
				t.Errorf("SomeService.GetParam() = %v, want %v", got, tt.want)
			}
		})
	}
}
