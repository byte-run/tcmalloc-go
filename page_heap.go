package tcmallocgo

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type pageHeap struct {
	lock sync.Mutex

	offset unsafe.Pointer // heap begin address offset
	size   uintptr        // heap size

	pageInUse atomic.Uint64
}
