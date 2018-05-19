package judge

import (
	"reflect"
	"testing"
)

func TestGetJudge(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    Judge
		wantErr bool
	}{
		{
			name:    "Test CPP",
			args:    args{"cpp"},
			want:    GetCPPRunner(),
			wantErr: false,
		},
		{
			name:    "Test Python2",
			args:    args{"python2"},
			want:    GetPython2Runner(),
			wantErr: false,
		},
		{
			name:    "Test Python3",
			args:    args{"python3"},
			want:    GetPython3Runner(),
			wantErr: false,
		},
		{
			name:    "Test language",
			args:    args{"language"},
			want:    Judge{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetJudge(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJudge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJudge() = %v, want %v", got, tt.want)
			}
		})
	}
}
