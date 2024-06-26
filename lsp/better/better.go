package better

// Liskov Substitution Principle (LSP)
// 只要不是subType關係(滿足LSP 7 rule), 就不該使用繼承
// subType: 如果S是T的subType 那在所有T出現過的地方都可以用S取代, 架構更有彈性
// subClass: A extend B (不符合7項條件) 目的是code reuse

// Advantages
// 可維護性, 確保實現interface的所有類型都符合預期行為
// 可讀性, 行為更具可預測性和一致性, 從而減少混淆和引入錯誤的可能
// 鬆耦合, 促進interface的使用, 助於在組件之間創建鬆耦合

// Disadvantages
// 複雜性增加, 可能需要創建其他interface或重構現有組件以確保他們遵循原則
// 過早抽象, 存在創建不必要的interface或抽象的風險, 而這些interface不會提供任何實際價值

// https://www.jyt0532.com/2020/03/22/lsp/

import "fmt"

type Bird interface {
	Walk()
}

type FlyingBird interface {
	Bird
	Fly()
}

type Pigeon struct{}

func (p *Pigeon) Walk() {
	fmt.Println("pigeon is walking")
}

func (p *Pigeon) Fly() {
	fmt.Println("pigeon is flying")
}

type Penguin struct{}

func (p *Penguin) Walk() {
	fmt.Println("penguin is walking")
}

func LetItFly(fb FlyingBird) {
	fb.Fly()
}

func LetItWalk(b Bird) {
	b.Walk()
}
