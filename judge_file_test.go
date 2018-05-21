package judge

import (
	"testing"
)

func Test_checkExtension(t *testing.T) {
	type args struct {
		fileName  string
		extension string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test check extension",
			args: args{"abcd.py", ".py"},
			want: true,
		},
		{
			name: "Test check extension - 2",
			args: args{"abcd.XDV", ".XDV"},
			want: true,
		},
		{
			name: "Test check extension - false",
			args: args{"abcd.py", ".cpp"},
			want: false,
		},
		{
			name: "Test check extension - false 2",
			args: args{"ab.cd.123", ".py"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkExtension(tt.args.fileName, tt.args.extension); got != tt.want {
				t.Errorf("checkExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJudge_AddSourceFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		judge   Judge
		args    args
		wantErr bool
	}{
		{
			name:    "Test AddSourceFile C++",
			judge:   GetCPPJudge(Config{}),
			args:    args{file: "ab.cpp"},
			wantErr: false,
		},
		{
			name:    "Test AddSourceFile C++ - error",
			judge:   GetCPPJudge(Config{}),
			args:    args{file: "ab.py"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			judge := tt.judge
			if err := judge.AddSourceFile(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("Judge.AddSourceFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJudge_AddInputFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		judge   Judge
		args    args
		wantErr bool
	}{
		{
			name:    "Test AddInputFile C++",
			judge:   GetCPPJudge(Config{}),
			args:    args{file: "ab.txt"},
			wantErr: false,
		},
		{
			name:    "Test AddInputFile C++ - error",
			judge:   GetCPPJudge(Config{}),
			args:    args{file: "ab.py"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			judge := tt.judge
			if err := judge.AddInputFile(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("Judge.AddInputFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
