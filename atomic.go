package atomic
// Package atomic implements sync/atomic in a slower, incorrect way. This
// package should not be used by anyone for any reason ever. If you see this
// package being used, you should immediately stop what you are doing and
// replace its import with "sync/atomic". You have been warned.
//
// This package is not correct, faster, stronger, better, or any positive adjective
// than the official sync/atomic package.

import (
	"sync"
	"unsafe"
)

var u sync.Mutex

func AddInt32(addr *int32, delta int32) (new int32) {
	u.Lock()
	defer u.Unlock()
	*addr += delta
	return *addr
}

func AddInt64(addr *int64, delta int64) (new int64) {
	u.Lock()
	defer u.Unlock()
	*addr += delta
	return *addr
}
func AddUint32(addr *uint32, delta uint32) (new uint32) {
	u.Lock()
	defer u.Unlock()
	*addr += delta
	return *addr
}
func AddUint64(addr *uint64, delta uint64) (new uint64) {
	u.Lock()
	defer u.Unlock()
	*addr += delta
	return *addr
}
func AddUintptr(addr *uintptr, delta uintptr) (new uintptr) {
	u.Lock()
	defer u.Unlock()
	*addr += delta
	return *addr
}
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool) {
	u.Lock()
	defer u.Unlock()
	if *addr == old {
		*addr = new
		return true
	}
	return false
}
func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool) {
	u.Lock()
	defer u.Unlock()
	if *addr == old {
		*addr = new
		return true
	}
	return false
}
func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool) {
	u.Lock()
	defer u.Unlock()
	if *addr == old {
		*addr = new
		return true
	}
	return false
}
func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool) {
	u.Lock()
	defer u.Unlock()
	if *addr == old {
		*addr = new
		return true
	}
	return false
}
func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool) {
	u.Lock()
	defer u.Unlock()
	if *addr == old {
		*addr = new
		return true
	}
	return false
}
func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool) {
	u.Lock()
	defer u.Unlock()
	if *addr == old {
		*addr = new
		return true
	}
	return false
}
func LoadInt32(addr *int32) (val int32) {
	u.Lock()
	defer u.Unlock()
	return *addr
}
func LoadInt64(addr *int64) (val int64) {
	u.Lock()
	defer u.Unlock()
	return *addr
}
func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer) {
	u.Lock()
	defer u.Unlock()
	return *addr
}
func LoadUint32(addr *uint32) (val uint32) {
	u.Lock()
	defer u.Unlock()
	return *addr
}
func LoadUint64(addr *uint64) (val uint64) {
	u.Lock()
	defer u.Unlock()
	return *addr
}
func LoadUintptr(addr *uintptr) (val uintptr) {
	u.Lock()
	defer u.Unlock()
	return *addr
}
func StoreInt32(addr *int32, val int32) {
	u.Lock()
	defer u.Unlock()
	*addr = val
}
func StoreInt64(addr *int64, val int64) {
	u.Lock()
	defer u.Unlock()
	*addr = val
}
func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer) {
	u.Lock()
	defer u.Unlock()
	*addr = val
}
func StoreUint32(addr *uint32, val uint32) {
	u.Lock()
	defer u.Unlock()
	*addr = val
}
func StoreUint64(addr *uint64, val uint64) {
	u.Lock()
	defer u.Unlock()
	*addr = val
}
func StoreUintptr(addr *uintptr, val uintptr) {
	u.Lock()
	defer u.Unlock()
	*addr = val
}
func SwapInt32(addr *int32, new int32) (old int32) {
	u.Lock()
	defer u.Unlock()
	old = *addr
	*addr = new
	return old
}
func SwapInt64(addr *int64, new int64) (old int64) {
	u.Lock()
	defer u.Unlock()
	old = *addr
	*addr = new
	return old
}
func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer) {
	u.Lock()
	defer u.Unlock()
	old = *addr
	*addr = new
	return old
}
func SwapUint32(addr *uint32, new uint32) (old uint32) {
	u.Lock()
	defer u.Unlock()
	old = *addr
	*addr = new
	return old
}
func SwapUint64(addr *uint64, new uint64) (old uint64) {
	u.Lock()
	defer u.Unlock()
	old = *addr
	*addr = new
	return old
}
func SwapUintptr(addr *uintptr, new uintptr) (old uintptr) {
	u.Lock()
	defer u.Unlock()
	old = *addr
	*addr = new
	return old
}

type Value struct {
	val interface{}
}

func (v *Value) Load() (x interface{}) {
	u.Lock()
	defer u.Unlock()
	return v.val
}

func (v *Value) Store(x interface{}) {
	u.Lock()
	defer u.Unlock()
	v.val = x
}
