package genericcomparable

// <T extends Comparable>

type Number interface {
	int64 | float64
}

// comparable关键字就可以限制map的Key的类型是否为可比较的
// comparable底层实现也就是基于一个interface的，只是里面有一个行为方法而已
// Go内置的可比较数据类型在更新之后会默认实现comparable
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func whoIsMin[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}
