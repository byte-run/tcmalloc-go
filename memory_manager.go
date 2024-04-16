package tcmallocgo

import "unsafe"

type MemoryManager struct {
}

// AcquireMemory 从mem-pool中申请合适的内存块
func (mem MemoryManager) AcquireMemory(numbytes uintptr) unsafe.Pointer {
	// TODO Unimplement
	return unsafe.Pointer(uintptr(0))
}

// ReleaseMemory 将使用完的内存块释放回mem-pool中
func (mem MemoryManager) ReleaseMemory(offset unsafe.Pointer, numbytes uintptr) {

}

// 单例控制
var memoryManager *MemoryManager

func NewMemoryManager() *MemoryManager {
	// TODO Unimplement

	if memoryManager != nil {
		return memoryManager
	}

	manager := new(MemoryManager)
	memoryManager = manager
	return manager
}
