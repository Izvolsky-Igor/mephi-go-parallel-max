package main

import "testing"

// TestGenerateRandomElements проверяет длину слайса и положительность чисел.
// Нулевой и отрицательный размеры должны давать пустой слайс.
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{name: "negative", size: -1, want: 0},
		{name: "zero", size: 0, want: 0},
		{name: "one", size: 1, want: 1},
		{name: "many", size: 100, want: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)
			if len(data) != tt.want {
				t.Fatalf("length = %d, want %d", len(data), tt.want)
			}
			for i, value := range data {
				if value <= 0 {
					t.Errorf("data[%d] = %d, want a positive value", i, value)
				}
			}
		})
	}
}

// TestMaximum проверяет последовательный поиск для nil, пустого слайса,
// одного элемента, повторяющегося максимума, отрицательных чисел и нуля.
func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil", want: 0},
		{name: "empty", data: []int{}, want: 0},
		{name: "one", data: []int{7}, want: 7},
		{name: "duplicates", data: []int{2, 9, 9, 4}, want: 9},
		{name: "negative", data: []int{-8, -2, -4}, want: -2},
		{name: "zero", data: []int{-3, 0, -1}, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

// TestMaxChunks проверяет пустой и короткий слайсы, а также разбиение
// на восемь частей с остатком и без него.
func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "empty", want: 0},
		{name: "short", data: []int{-3, -1, -2}, want: -1},
		{name: "even chunks", data: []int{1, 2, 3, 9, 4, 5, 6, 7}, want: 9},
		// Максимум находится в элементе, который попадает в остаток.
		{name: "remainder", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, want: 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}
