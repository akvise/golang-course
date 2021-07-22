package main

import "testing"

func TestCircle_Area(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{name: "return err if negative value",
			fields: fields{radius: -10},
			want: 0,
			wantErr: true,
		},
		{name: "return 201.06192982974676",
			fields: fields{radius: 8},
			want: 201.06192982974676,
			wantErr: false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				radius: tt.fields.radius,
			}
			got, err := c.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Area() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Perimeter(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name: "negative value",
			fields: fields{radius: -10},
			want: 0,
			wantErr: true,
		},
		{
			name: "want to 50.26548245743669 value",
			fields: fields{radius: 8},
			want: 50.26548245743669,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				radius: tt.fields.radius,
			}
			got, err := c.Perimeter()
			if (err != nil) != tt.wantErr {
				t.Errorf("Perimeter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Perimeter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Area(t *testing.T) {
	type fields struct {
		height float64
		width  float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name: "want to 27.0 value",
			fields: fields{width: 3, height: 9},
			want: 27.0,
			wantErr: false,
		},
		{
			name: "err negative width",
			fields: fields{width: -3, height: 9},
			want: 0,
			wantErr: true,
		},
		{
			name: "err negative height",
			fields: fields{width: 3, height: -9},
			want: 0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				height: tt.fields.height,
				width:  tt.fields.width,
			}
			got, err := r.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Area() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Perimeter(t *testing.T) {
	type fields struct {
		height float64
		width  float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name: "want to 27.0 value",
			fields: fields{width: 3, height: 9},
			want: 24.0,
			wantErr: false,
		},
		{
			name: "err negative width",
			fields: fields{width: -3, height: 9},
			want: 0,
			wantErr: true,
		},
		{
			name: "err negative height",
			fields: fields{width: 3, height: -9},
			want: 0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				height: tt.fields.height,
				width:  tt.fields.width,
			}
			got, err := r.Perimeter()
			if (err != nil) != tt.wantErr {
				t.Errorf("Perimeter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Perimeter() got = %v, want %v", got, tt.want)
			}
		})
	}
}