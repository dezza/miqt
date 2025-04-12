package script

/*

#include "gen_qscriptstring.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QScriptString struct {
	h *C.QScriptString
}

func (this *QScriptString) cPointer() *C.QScriptString {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QScriptString) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQScriptString constructs the type using only CGO pointers.
func newQScriptString(h *C.QScriptString) *QScriptString {
	if h == nil {
		return nil
	}

	return &QScriptString{h: h}
}

// UnsafeNewQScriptString constructs the type using only unsafe pointers.
func UnsafeNewQScriptString(h unsafe.Pointer) *QScriptString {
	return newQScriptString((*C.QScriptString)(h))
}

// NewQScriptString constructs a new QScriptString object.
func NewQScriptString() *QScriptString {

	return newQScriptString(C.QScriptString_new())
}

// NewQScriptString2 constructs a new QScriptString object.
func NewQScriptString2(other *QScriptString) *QScriptString {

	return newQScriptString(C.QScriptString_new2(other.cPointer()))
}

func (this *QScriptString) OperatorAssign(other *QScriptString) {
	C.QScriptString_operatorAssign(this.h, other.cPointer())
}

func (this *QScriptString) IsValid() bool {
	return (bool)(C.QScriptString_isValid(this.h))
}

func (this *QScriptString) OperatorEqual(other *QScriptString) bool {
	return (bool)(C.QScriptString_operatorEqual(this.h, other.cPointer()))
}

func (this *QScriptString) OperatorNotEqual(other *QScriptString) bool {
	return (bool)(C.QScriptString_operatorNotEqual(this.h, other.cPointer()))
}

func (this *QScriptString) ToArrayIndex() uint {
	return (uint)(C.QScriptString_toArrayIndex(this.h))
}

func (this *QScriptString) ToString() string {
	var _ms C.struct_miqt_string = C.QScriptString_toString(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QScriptString) ToQString() string {
	var _ms C.struct_miqt_string = C.QScriptString_ToQString(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QScriptString) ToArrayIndexWithOk(ok *bool) uint {
	return (uint)(C.QScriptString_toArrayIndexWithOk(this.h, (*C.bool)(unsafe.Pointer(ok))))
}

// Delete this object from C++ memory.
func (this *QScriptString) Delete() {
	C.QScriptString_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QScriptString) GoGC() {
	runtime.SetFinalizer(this, func(this *QScriptString) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
