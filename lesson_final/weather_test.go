package main

import "testing"


func TestReadToken(t *testing.T) {
	type args struct {
		Path      string
		TokenType string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "err",
			args: args{Path: "./token.json", TokenType: "123"},
			want: "",
			wantErr: true,
		},
		{name: "err",
			args: args{Path: "123", TokenType: "telegram"},
			want: "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadToken(tt.args.Path, tt.args.TokenType)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}