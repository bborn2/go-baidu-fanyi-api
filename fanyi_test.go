package baidufanyiapi

import "testing"

func Test_translate(t *testing.T) {
	type args struct {
		appid  string
		appkey string
		fr     string
		to     string
		query  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"translate", args{"yourappid", "yourappkey", "en", "zh", "Hello"}, "你好"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translate(tt.args.appid, tt.args.appkey, tt.args.fr, tt.args.to, tt.args.query); got != tt.want {
				t.Errorf("translate() = %v, want %v", got, tt.want)
			}
		})
	}
}
