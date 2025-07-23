package circular

import (
	"errors"
	"sync"
)

type Buffer struct {
	data     []byte
	capacity int
	size     int
	readIdx  int
	writeIdx int
	mutex    sync.Mutex
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		data:     make([]byte, size),
		capacity: size,
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.size == 0 {
		return 0, errors.New("empty")
	}
	val := b.data[b.readIdx]
	b.readIdx = (b.readIdx + 1) % b.capacity
	b.size--
	return val, nil
}

func (b *Buffer) WriteByte(c byte) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.size == b.capacity {
		return errors.New("full")
	}
	b.data[b.writeIdx] = c
	b.writeIdx = (b.writeIdx + 1) % b.capacity
	b.size++
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.size == b.capacity {
		// buffer full, overwrite oldest
		b.data[b.writeIdx] = c
		b.writeIdx = (b.writeIdx + 1) % b.capacity
		b.readIdx = (b.readIdx + 1) % b.capacity
	} else {
		// 不能调用WriteByte，重复加锁会死锁
		//b.WriteByte(c)
		b.data[b.writeIdx] = c
		b.writeIdx = (b.writeIdx + 1) % b.capacity
		b.size++
	}
}

func (b *Buffer) Reset() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	// 虽然 只重置 size 即可， 原有数据读不出来，下次写也会覆盖
	// 但是可以进行 防御式清理

	// 方式1：创建了一个全新的零初始化的切片，旧的切片 gc 了
	// b.data = make([]byte, b.capacity)
	// b.readIdx = 0
	// b.writeIdx = 0

	// 方式2：不会触发额外的内存分配，原数组地址不变
	for i := range b.data {
		b.data[i] = 0
	}

	b.size = 0
}
