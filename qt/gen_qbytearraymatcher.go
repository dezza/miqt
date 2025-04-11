package qt

/*

#include "gen_qbytearraymatcher.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QByteArrayMatcher struct {
	h *C.QByteArrayMatcher
}

func (this *QByteArrayMatcher) cPointer() *C.QByteArrayMatcher {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QByteArrayMatcher) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQByteArrayMatcher constructs the type using only CGO pointers.
func newQByteArrayMatcher(h *C.QByteArrayMatcher) *QByteArrayMatcher {
	if h == nil {
		return nil
	}

	return &QByteArrayMatcher{h: h}
}

// UnsafeNewQByteArrayMatcher constructs the type using only unsafe pointers.
func UnsafeNewQByteArrayMatcher(h unsafe.Pointer) *QByteArrayMatcher {
	return newQByteArrayMatcher((*C.QByteArrayMatcher)(h))
}

// NewQByteArrayMatcher constructs a new QByteArrayMatcher object.
func NewQByteArrayMatcher() *QByteArrayMatcher {

	return newQByteArrayMatcher(C.QByteArrayMatcher_new())
}

// NewQByteArrayMatcher2 constructs a new QByteArrayMatcher object.
func NewQByteArrayMatcher2(pattern []byte) *QByteArrayMatcher {
	pattern_alias := C.struct_miqt_string{}
	if len(pattern) > 0 {
		pattern_alias.data = (*C.char)(unsafe.Pointer(&pattern[0]))
	} else {
		pattern_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	pattern_alias.len = C.size_t(len(pattern))

	return newQByteArrayMatcher(C.QByteArrayMatcher_new2(pattern_alias))
}

// NewQByteArrayMatcher3 constructs a new QByteArrayMatcher object.
func NewQByteArrayMatcher3(pattern string, length int) *QByteArrayMatcher {
	pattern_Cstring := C.CString(pattern)
	defer C.free(unsafe.Pointer(pattern_Cstring))

	return newQByteArrayMatcher(C.QByteArrayMatcher_new3(pattern_Cstring, (C.int)(length)))
}

// NewQByteArrayMatcher4 constructs a new QByteArrayMatcher object.
func NewQByteArrayMatcher4(other *QByteArrayMatcher) *QByteArrayMatcher {

	return newQByteArrayMatcher(C.QByteArrayMatcher_new4(other.cPointer()))
}

func (this *QByteArrayMatcher) OperatorAssign(other *QByteArrayMatcher) {
	C.QByteArrayMatcher_operatorAssign(this.h, other.cPointer())
}

func (this *QByteArrayMatcher) SetPattern(pattern []byte) {
	pattern_alias := C.struct_miqt_string{}
	if len(pattern) > 0 {
		pattern_alias.data = (*C.char)(unsafe.Pointer(&pattern[0]))
	} else {
		pattern_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	pattern_alias.len = C.size_t(len(pattern))
	C.QByteArrayMatcher_setPattern(this.h, pattern_alias)
}

func (this *QByteArrayMatcher) IndexIn(ba []byte) int {
	ba_alias := C.struct_miqt_string{}
	if len(ba) > 0 {
		ba_alias.data = (*C.char)(unsafe.Pointer(&ba[0]))
	} else {
		ba_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	ba_alias.len = C.size_t(len(ba))
	return (int)(C.QByteArrayMatcher_indexIn(this.h, ba_alias))
}

func (this *QByteArrayMatcher) IndexIn2(str string, lenVal int) int {
	str_Cstring := C.CString(str)
	defer C.free(unsafe.Pointer(str_Cstring))
	return (int)(C.QByteArrayMatcher_indexIn2(this.h, str_Cstring, (C.int)(lenVal)))
}

func (this *QByteArrayMatcher) Pattern() []byte {
	var _bytearray C.struct_miqt_string = C.QByteArrayMatcher_pattern(this.h)
	_ret := C.GoBytes(unsafe.Pointer(_bytearray.data), C.int(int64(_bytearray.len)))
	C.free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QByteArrayMatcher) IndexIn3(ba []byte, from int) int {
	ba_alias := C.struct_miqt_string{}
	if len(ba) > 0 {
		ba_alias.data = (*C.char)(unsafe.Pointer(&ba[0]))
	} else {
		ba_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	ba_alias.len = C.size_t(len(ba))
	return (int)(C.QByteArrayMatcher_indexIn3(this.h, ba_alias, (C.int)(from)))
}

func (this *QByteArrayMatcher) IndexIn4(str string, lenVal int, from int) int {
	str_Cstring := C.CString(str)
	defer C.free(unsafe.Pointer(str_Cstring))
	return (int)(C.QByteArrayMatcher_indexIn4(this.h, str_Cstring, (C.int)(lenVal), (C.int)(from)))
}

// Delete this object from C++ memory.
func (this *QByteArrayMatcher) Delete() {
	C.QByteArrayMatcher_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QByteArrayMatcher) GoGC() {
	runtime.SetFinalizer(this, func(this *QByteArrayMatcher) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}

type QStaticByteArrayMatcherBase struct {
	h *C.QStaticByteArrayMatcherBase
}

func (this *QStaticByteArrayMatcherBase) cPointer() *C.QStaticByteArrayMatcherBase {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QStaticByteArrayMatcherBase) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQStaticByteArrayMatcherBase constructs the type using only CGO pointers.
func newQStaticByteArrayMatcherBase(h *C.QStaticByteArrayMatcherBase) *QStaticByteArrayMatcherBase {
	if h == nil {
		return nil
	}

	return &QStaticByteArrayMatcherBase{h: h}
}

// UnsafeNewQStaticByteArrayMatcherBase constructs the type using only unsafe pointers.
func UnsafeNewQStaticByteArrayMatcherBase(h unsafe.Pointer) *QStaticByteArrayMatcherBase {
	return newQStaticByteArrayMatcherBase((*C.QStaticByteArrayMatcherBase)(h))
}

// Delete this object from C++ memory.
func (this *QStaticByteArrayMatcherBase) Delete() {
	C.QStaticByteArrayMatcherBase_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QStaticByteArrayMatcherBase) GoGC() {
	runtime.SetFinalizer(this, func(this *QStaticByteArrayMatcherBase) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
