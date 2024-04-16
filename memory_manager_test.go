package tcmallocgo

import (
	"testing"
	"unsafe"
)

func TestSingleInstance_MemoryManager(t *testing.T) {
	mm := NewMemoryManager()
	// mm addr
	mmAddr := unsafe.Pointer(mm)
	t.Logf("first mm addr: %v\n", mmAddr)

	// 再次创建
	mm2 := NewMemoryManager()
	mmAddr2 := unsafe.Pointer(mm2)
	t.Logf("second mm addr: %v\n", mmAddr2)
	t.Logf("twice addr equal: %v\n", mm == mm2) // true
}
