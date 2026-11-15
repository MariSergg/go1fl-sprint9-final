package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name       string
		size       int
		wantLength int
	}{
		{"отрицательный размер", -5, 0},
		{"нулевой размер", 0, 0},
		{"один элемент", 1, 1},
		{"десять элементов", 10, 10},
		{"тысяча элементов", 1000, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)

			if len(data) != tt.wantLength {
				t.Fatalf(
					"размер %d: получили длину %d, ожидали %d",
					tt.size,
					len(data),
					tt.wantLength,
				)
			}

			for i, value := range data {
				if value < 1 || value > 1_000_000 {
					t.Fatalf(
						"элемент %d: получили %d, ожидали положительное число от 1 до 1000000",
						i,
						value,
					)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil слайс", nil, 0},
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{42}, 42},
		{"максимум в начале", []int{99, 3, 7, 1}, 99},
		{"максимум в середине", []int{1, 3, 99, 7}, 99},
		{"максимум в конце", []int{1, 3, 7, 99}, 99},
		{"равные значения", []int{5, 5, 5}, 5},
		{"нулевые значения", []int{0, 0, 0}, 0},
		{"отрицательные значения", []int{-9, -2, -7}, -2},
		{"смешанные значения", []int{-9, 0, 12, -2}, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)

			if got != tt.want {
				t.Errorf(
					"maximum(%v) = %d, ожидали %d",
					tt.data,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			"nil слайс",
			nil,
			0,
		},
		{
			"пустой слайс",
			[]int{},
			0,
		},
		{
			"один элемент",
			[]int{42},
			42,
		},
		{
			"меньше восьми",
			[]int{1, 7, 3, 4},
			7,
		},
		{
			"ровно восемь",
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			8,
		},
		{
			"девять элементов: максимум в остатке",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 100},
			100,
		},
		{
			"шестнадцать элементов",
			[]int{30, 2, 3, 4, 5, 6, 7, 8,
				9, 10, 11, 12, 13, 14, 15, 16},
			30,
		},
		{
			"семнадцать элементов: максимум в остатке",
			[]int{1, 2, 3, 4, 5, 6, 7, 8,
				9, 10, 11, 12, 13, 14, 15, 16, 100},
			100,
		},
		{
			"отрицательные значения",
			[]int{-9, -8, -7, -6, -5, -4, -3, -2, -1},
			-1,
		},
		{
			"одинаковые элементы",
			[]int{5, 5, 5, 5, 5, 5, 5, 5},
			5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)

			if got != tt.want {
				t.Errorf(
					"maxChunks(%v) = %d, ожидали %d",
					tt.data,
					got,
					tt.want,
				)
			}
		})
	}
}
