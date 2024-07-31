package paasio

import (
	"io"
	"sync"
)

// readCounter 结构体
type readCounter struct {
	reader    io.Reader
	readBytes int64
	readOps   int
	mu        sync.Mutex
}

// writeCounter 结构体
type writeCounter struct {
	writer       io.Writer
	writtenBytes int64
	writeOps     int
	mu           sync.Mutex
}

// readWriteCounter 结构体
type readWriteCounter struct {
	readCounter
	writeCounter
}

// NewReadCounter 返回一个新的 readCounter 实例
func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader}
}

// NewWriteCounter 返回一个新的 writeCounter 实例
func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer}
}

// NewReadWriteCounter 返回一个新的 readWriteCounter 实例
func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		readCounter:  readCounter{reader: readwriter},
		writeCounter: writeCounter{writer: readwriter},
	}
}

// Read 实现了 io.Reader 接口，并记录读取的字节数和读取操作次数
func (rc *readCounter) Read(p []byte) (int, error) {
	n, err := rc.reader.Read(p)
	rc.mu.Lock()
	rc.readBytes += int64(n)
	rc.readOps++
	rc.mu.Unlock()
	return n, err
}

// ReadCount 返回读取的总字节数和读取操作的次数
func (rc *readCounter) ReadCount() (int64, int) {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	return rc.readBytes, rc.readOps
}

// Write 实现了 io.Writer 接口，并记录写入的字节数和写入操作次数
func (wc *writeCounter) Write(p []byte) (int, error) {
	n, err := wc.writer.Write(p)
	wc.mu.Lock()
	wc.writtenBytes += int64(n)
	wc.writeOps++
	wc.mu.Unlock()
	return n, err
}

// WriteCount 返回写入的总字节数和写入操作的次数
func (wc *writeCounter) WriteCount() (int64, int) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	return wc.writtenBytes, wc.writeOps
}
