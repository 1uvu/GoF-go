package iterator

// ===
// 定义 Aggregate 接口
// ===

type Aggregate interface {
	// Aggergate 与 Iterator 是一一对应的
	Iterator() Iterator // 定义 Iterator(), 返回对应的实现了 Iterator 接口的实例
	// 定义生成 Iterator 的方法, 可能会用到另一种设计模式: Factory Method 模式
}
