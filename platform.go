package tcmallocgo

import "unsafe"

/**
通过使用go:linkname编译指令，链接到runtime包中申请内存的私有方法
*/
//go:linkname sysAllocOS runtime.sysAllocOS
func sysAllocOS(n uintptr) unsafe.Pointer

//go:linkname sysFreeOS runtime.sysFreeOS
func sysFreeOS(v unsafe.Pointer, n uintptr)
