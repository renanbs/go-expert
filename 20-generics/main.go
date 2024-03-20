package main

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

type MyNumber int

type Number interface {
	~int | float64
}

func Soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Soma2[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Renan": 1000, "Juca": 2000}
	m2 := map[string]float64{"Renan": 1000.5, "Juca": 2000.5}
	m3 := map[string]MyNumber{"Renan": 1000, "Juca": 2000}
	println(SomaInteiro(m))
	println(SomaFloat(m2))

	println(Soma(m))
	println(Soma(m2))

	println(Soma2(m))
	println(Soma2(m2))
	println(Soma2(m3))

	println(Compara(2, 2))
	println(Compara(3, 2))
}
