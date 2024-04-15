package better

// Open Close Principle (OCP)
// 應該要能夠在不更改現有代碼(關閉)的情況下擴展(開放)其行為

// Advantages
// 可維護性, 因為在進行更改或添加新功能時不太可能引入錯誤或破壞現有功能
// 靈活性, 輕鬆適應新的要求或更改, 而不影響現有功能
// 可測試性, 每個組件可以單獨測試, 更容易識別和修復問題

// Disadvantages
// 複雜性增加, 因為導入額外的抽象(interface)以及組件數量增加
// 過早抽象, 存在創建不必要的抽象的風險, 可能導致過度工程化或設計過於複雜

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func CalculateArea(shape Shape) float64 {
	return shape.Area()
}
