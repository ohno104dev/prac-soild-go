package better

// Interface Segregation Principle (ISP)
// 不應該去依賴你根本不會用到的東西
// 應該使用多個小的, 集中的interface, 而不是一個涵蓋所有可能的大型單體interface

// Advantages
// 靈活性, 更小,更集中的interface, 方便創建更靈活的系統
// 可維護性, 更小,更集中的interface更易於理解和維護
// 鬆耦合, 減少模塊之間的耦合

// Disadvantages
// 複雜性增加, 將interface分解為更小的部分, 導致interface數量增加
// 組織難度增加, interface數量的增加, 有效地組織代碼庫可能變得具有挑戰性

type Walker interface {
	Walk() string
}

type Flyer interface {
	Fly() string
}

type Swimmer interface {
	Swim() string
}

// Now, we can create specific animal types that implement only the actions they can perform
type Penguin struct{}

func (p *Penguin) Walk() string {
	return "penguin can walk"
}

func (p *Penguin) Swim() string {
	return "penguin can swim"
}
