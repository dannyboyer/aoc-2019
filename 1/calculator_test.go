package main

import (
	"testing"
)

func Test_calcFuelReqWithoutMass(t *testing.T) {
	type args struct {
		mass float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"basic test - 01", args{12}, 2},
		{"basic test - 01", args{14}, 2},
		{"basic test - 01", args{1969}, 654},
		{"basic test - 01", args{100756}, 33583},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcFuelReqWithoutMass(tt.args.mass); got != tt.want {
				t.Errorf("calcFuelReqWithoutMass() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcFuelReq(t *testing.T) {
	type args struct {
		mass float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"basic test - 01", args{12}, 2},
		{"basic test - 01", args{14}, 2},
		{"basic test - 01", args{1969}, 966},
		{"basic test - 01", args{100756}, 50346},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcFuelReq(tt.args.mass); got != tt.want {
				t.Errorf("calcFuelReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAnswer2(t *testing.T) {
	type args struct {
		massList []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1 ele list", args{[]float64{12}}, 2},
		{"2 ele list", args{[]float64{12, 14}}, 4},
		{"3 ele list", args{[]float64{12, 14, 1969}}, 970},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAnswer2(tt.args.massList); got != tt.want {
				t.Errorf("getAnswer2() = %v, want %v", got, tt.want)
			}
		})
	}
}
