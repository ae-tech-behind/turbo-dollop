package controller

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/ae-tech-behind/turbo-dollop/controller/mocks"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
)

func TestNewSomething(t *testing.T) {
	type args struct {
		uc SomethingUseCase
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	SomethingC := &Something{
		UseCase: mocks.NewMockSomethingUseCase(mockCtrl),
	}
	tests := []struct {
		name string
		args args
		want *Something
	}{
		{
			name: "Success NewSomething",
			args: args{
				uc: SomethingC.UseCase,
			},
			want: SomethingC,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSomething(tt.args.uc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSomething() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSomething_HandlerSomething(t *testing.T) {
	type fields struct {
		UseCase *mocks.MockSomethingUseCase
	}
	type args struct {
		req *http.Request
		rec *httptest.ResponseRecorder
	}
	type argsUC struct {
		params   string
		response map[string]int
		err      error
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	e := echo.New()

	tests := []struct {
		name     string
		fields   fields
		args     args
		argsUC   argsUC
		wantResp string
		wantCode int
		wantErr  bool
	}{
		{
			name: "Success HandlerSomething",
			fields: fields{
				UseCase: mocks.NewMockSomethingUseCase(mockCtrl),
			},
			args: args{
				req: httptest.NewRequest(
					http.MethodPost,
					"/",
					strings.NewReader(`{"info":"test"}`),
				),
				rec: httptest.NewRecorder(),
			},
			argsUC: argsUC{
				params:   "test",
				response: map[string]int{"test": 1},
				err:      nil,
			},
			wantResp: "{\"test\":1}\n",
			wantCode: http.StatusOK,
			wantErr:  false,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			tt.fields.UseCase.EXPECT().DoSomething(tt.argsUC.params).
				Return(tt.argsUC.response, tt.argsUC.err)

			c := &Something{
				UseCase: tt.fields.UseCase,
			}

			eCtx := e.NewContext(tt.args.req, tt.args.rec)
			if err := c.HandlerSomething(eCtx); (err != nil) != tt.wantErr {
				t.Errorf("Something.HandlerSomething() error = %v, wantErr %v", err, tt.wantErr)
			}

			respBody := tt.args.rec.Body
			if respBody.String() != tt.wantResp {
				t.Errorf("Something.HandlerSomething() expected = %v, got %v", tt.wantResp, respBody.String())
			}
			if tt.args.rec.Code != tt.wantCode {
				t.Errorf("Something.HandlerSomething() expected = %v, got %v", tt.wantCode, tt.args.rec.Code)
			}
		})
	}
}
