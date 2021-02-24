package alipay

import (
	"fmt"
	"github.com/ooncn/common/obj"
	"github.com/ooncn/common/util"
	"reflect"
	"testing"
	"time"
)

func TestGetConfigOptions(t *testing.T) {
	fmt.Println(util.TimeUtil.DateToyMdHmsSepTo(time.Now().Add(30 * time.Minute)))
}

func TestNewAopClient(t *testing.T) {
	tests := []struct {
		name string
		want *AopClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAopClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAopClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAopClientFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want *AopClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAopClientFile(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAopClientFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRequest(t *testing.T) {
	tests := []struct {
		name string
		want *Request
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRequestClient(t *testing.T) {
	type args struct {
		client AopClient
	}
	tests := []struct {
		name string
		args args
		want *Request
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRequestClient(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRequestClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequest_CheckSign(t *testing.T) {
	type fields struct {
		Map obj.Map
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Request{
				Map: tt.fields.Map,
			}
			if got := r.CheckSign(); got != tt.want {
				t.Errorf("CheckSign() = %v, want %v", got, tt.want)
			}
		})
	}
}
