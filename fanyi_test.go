package baidufanyiapi

import (
	"testing"
)

func TestTranslate(t *testing.T) {
	type args struct {
		appid  string
		appkey string
		fr     string
		to     string
		query  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"translate", args{"yourappid", "yourappkey", "en", "zh", "Hello"}, "你好", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Translate(tt.args.appid, tt.args.appkey, tt.args.fr, tt.args.to, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Translate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}
