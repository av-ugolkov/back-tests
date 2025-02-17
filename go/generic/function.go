package generic

type Number interface {
	int | uint | float64
}

func Sum[T Number](a, b T) T {
	return a + b
}
