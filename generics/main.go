package main

import "fmt"

// Number は、int64 または float64 を許容する型制約です。
type Number interface {
	int64 | float64
}

func main() {
	// int64 の値を持つマップを初期化
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// float64 の値を持つマップを初期化
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}

// SumInts は、マップ m の int64 値を合計して返します。
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats は、マップ m の float64 値を合計して返します。
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats は、マップ m の値を合計します。
// この関数は、マップの値が int64 または float64 の場合に利用できます。
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumbers は、マップ m の値を合計します。整数と浮動小数点数の両方に対応します。
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
