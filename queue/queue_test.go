package queue

import (
	"log"
	"testing"
)

// 测试队列
// @author staitc7 <static7@qq.com>
func TestConstructor(t *testing.T) {
	// Your MyCircularQueue object will be instantiated and called as such:
	obj := Constructor(6)
	for i := 1; i <= 6; i++ {
		obj.EnQueue(i)
		log.Println(obj.Data)
	}
	log.Println(obj.Rear())
	log.Println(obj.IsFull())
	log.Println(obj.DeQueue())
	log.Println(obj.EnQueue(8))
	log.Println(obj.Rear())

}
