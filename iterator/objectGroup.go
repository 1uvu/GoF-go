package iterator

import "fmt"

// ===
// 定义实现了 Aggregate 接口的测试 ObjectGroup
// ===

type ObjectGroup struct {
	objects  []*Object
	capacity int
}

func NewObjectGroup(capacity int) *ObjectGroup {
	return &ObjectGroup{
		objects:  make([]*Object, 0, capacity),
		capacity: capacity,
	}
}

// 实现 Aggregate 接口
func (og *ObjectGroup) Iterator() *ObjectIterator {
	return newObjectIterator(og, 0)
}

func (og *ObjectGroup) Size() int {
	return len(og.objects)
}

func (og *ObjectGroup) Capacity() int {
	return og.capacity
}

func (og *ObjectGroup) GetObject(index int) *Object {
	return og.objects[index]
}

func (og *ObjectGroup) AddObject(o *Object) error {
	if og.Size() == og.Capacity() {
		return fmt.Errorf("object group's capacity is equal to its size: %d.\n", og.Capacity())
	}
	og.objects = append(og.objects, o)
	return nil
}
