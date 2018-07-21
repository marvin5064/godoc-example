package fruit

import "fmt"

func ExampleNewApple() {
	apple := NewApple()
	fmt.Println(apple.GetExpireTime())
	fmt.Println(apple.GetTaste())
	// Output: 720h0m0s
	// [sweet juicy]
}

func ExampleNewOrange_expireate() {
	orange := NewOrange()
	fmt.Println(orange.GetExpireTime())
	// Output: 1440h0m0s
}

func ExampleNewOrange_taste() {
	orange := NewOrange()
	fmt.Println(orange.GetTaste())
	// Output: [sour juicy]
}
