package controller

import (
	"reflect"
	"testing"

	"github.com/labstack/echo"
)

func TestNewStatus(t *testing.T) {
	type args struct {
		uc StatusUseCase
	}
	tests := []struct {
		name string
		args args
		want *Status
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStatus(tt.args.uc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_HandlerStatusz(t *testing.T) {
	type fields struct {
		UseCase StatusUseCase
	}
	type args struct {
		eCtx echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Status{
				UseCase: tt.fields.UseCase,
			}
			if err := c.HandlerStatusz(tt.args.eCtx); (err != nil) != tt.wantErr {
				t.Errorf("Status.HandlerStatusz() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStatus_HandlerHealthz(t *testing.T) {
	type fields struct {
		UseCase StatusUseCase
	}
	type args struct {
		eCtx echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Status{
				UseCase: tt.fields.UseCase,
			}
			if err := c.HandlerHealthz(tt.args.eCtx); (err != nil) != tt.wantErr {
				t.Errorf("Status.HandlerHealthz() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
