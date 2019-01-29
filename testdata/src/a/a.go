package a

import "bar"

type foo struct {
	A int
	B int
	C int
}

func main() {
	_ = &foo{
		A: 1,
		B: 2,
		C: 3,
	}
	_ = foo{}
	_ = foo{ // want "find missing props: B"
		A: 1,
		C: 3,
	}
	_ = bar.Boo{ // want "find missing props: E"
		D: 1,
		F: 3,
	}
}
