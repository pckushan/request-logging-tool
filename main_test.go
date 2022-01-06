package main

import (
	"reflect"
	"testing"
)

func Test_reAdjustURLs(t *testing.T) {
	type args struct {
		urls []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "simple_urls_test",
			args: args{
				urls: []string{"abc.com", "http://pqrst.xxx", "ddd.com"},
			},
			want: []string{"http://abc.com", "http://pqrst.xxx", "http://ddd.com"},
		},
		{
			name: "simple_urls_test",
			args: args{
				urls: []string{"abc.com", "http://pqrst.xxx", "https://ddd.com"},
			},
			want: []string{"http://abc.com", "http://pqrst.xxx", "https://ddd.com"},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			if got := reAdjustURLs(tt.args.urls); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readjustURLs() = %v, want %v", got, tt.want)
			}
		})
	}
}
