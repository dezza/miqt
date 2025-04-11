package qt6

/*

#include "gen_qinputdevice.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QInputDevice__DeviceType int

const (
	QInputDevice__Unknown     QInputDevice__DeviceType = 0
	QInputDevice__Mouse       QInputDevice__DeviceType = 1
	QInputDevice__TouchScreen QInputDevice__DeviceType = 2
	QInputDevice__TouchPad    QInputDevice__DeviceType = 4
	QInputDevice__Puck        QInputDevice__DeviceType = 8
	QInputDevice__Stylus      QInputDevice__DeviceType = 16
	QInputDevice__Airbrush    QInputDevice__DeviceType = 32
	QInputDevice__Keyboard    QInputDevice__DeviceType = 4096
	QInputDevice__AllDevices  QInputDevice__DeviceType = 2147483647
)

type QInputDevice__Capability int

const (
	QInputDevice__None               QInputDevice__Capability = 0
	QInputDevice__Position           QInputDevice__Capability = 1
	QInputDevice__Area               QInputDevice__Capability = 2
	QInputDevice__Pressure           QInputDevice__Capability = 4
	QInputDevice__Velocity           QInputDevice__Capability = 8
	QInputDevice__NormalizedPosition QInputDevice__Capability = 32
	QInputDevice__MouseEmulation     QInputDevice__Capability = 64
	QInputDevice__PixelScroll        QInputDevice__Capability = 128
	QInputDevice__Scroll             QInputDevice__Capability = 256
	QInputDevice__Hover              QInputDevice__Capability = 512
	QInputDevice__Rotation           QInputDevice__Capability = 1024
	QInputDevice__XTilt              QInputDevice__Capability = 2048
	QInputDevice__YTilt              QInputDevice__Capability = 4096
	QInputDevice__TangentialPressure QInputDevice__Capability = 8192
	QInputDevice__ZPosition          QInputDevice__Capability = 16384
	QInputDevice__All                QInputDevice__Capability = 2147483647
)

type QInputDevice struct {
	h *C.QInputDevice
	*QObject
}

func (this *QInputDevice) cPointer() *C.QInputDevice {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QInputDevice) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQInputDevice constructs the type using only CGO pointers.
func newQInputDevice(h *C.QInputDevice) *QInputDevice {
	if h == nil {
		return nil
	}
	var outptr_QObject *C.QObject = nil
	C.QInputDevice_virtbase(h, &outptr_QObject)

	return &QInputDevice{h: h,
		QObject: newQObject(outptr_QObject)}
}

// UnsafeNewQInputDevice constructs the type using only unsafe pointers.
func UnsafeNewQInputDevice(h unsafe.Pointer) *QInputDevice {
	return newQInputDevice((*C.QInputDevice)(h))
}

// NewQInputDevice constructs a new QInputDevice object.
func NewQInputDevice() *QInputDevice {

	return newQInputDevice(C.QInputDevice_new())
}

// NewQInputDevice2 constructs a new QInputDevice object.
func NewQInputDevice2(name string, systemId int64, typeVal QInputDevice__DeviceType) *QInputDevice {
	name_ms := C.struct_miqt_string{}
	name_ms.data = C.CString(name)
	name_ms.len = C.size_t(len(name))
	defer C.free(unsafe.Pointer(name_ms.data))

	return newQInputDevice(C.QInputDevice_new2(name_ms, (C.longlong)(systemId), (C.int)(typeVal)))
}

// NewQInputDevice3 constructs a new QInputDevice object.
func NewQInputDevice3(parent *QObject) *QInputDevice {

	return newQInputDevice(C.QInputDevice_new3(parent.cPointer()))
}

// NewQInputDevice4 constructs a new QInputDevice object.
func NewQInputDevice4(name string, systemId int64, typeVal QInputDevice__DeviceType, seatName string) *QInputDevice {
	name_ms := C.struct_miqt_string{}
	name_ms.data = C.CString(name)
	name_ms.len = C.size_t(len(name))
	defer C.free(unsafe.Pointer(name_ms.data))
	seatName_ms := C.struct_miqt_string{}
	seatName_ms.data = C.CString(seatName)
	seatName_ms.len = C.size_t(len(seatName))
	defer C.free(unsafe.Pointer(seatName_ms.data))

	return newQInputDevice(C.QInputDevice_new4(name_ms, (C.longlong)(systemId), (C.int)(typeVal), seatName_ms))
}

// NewQInputDevice5 constructs a new QInputDevice object.
func NewQInputDevice5(name string, systemId int64, typeVal QInputDevice__DeviceType, seatName string, parent *QObject) *QInputDevice {
	name_ms := C.struct_miqt_string{}
	name_ms.data = C.CString(name)
	name_ms.len = C.size_t(len(name))
	defer C.free(unsafe.Pointer(name_ms.data))
	seatName_ms := C.struct_miqt_string{}
	seatName_ms.data = C.CString(seatName)
	seatName_ms.len = C.size_t(len(seatName))
	defer C.free(unsafe.Pointer(seatName_ms.data))

	return newQInputDevice(C.QInputDevice_new5(name_ms, (C.longlong)(systemId), (C.int)(typeVal), seatName_ms, parent.cPointer()))
}

func (this *QInputDevice) MetaObject() *QMetaObject {
	return newQMetaObject(C.QInputDevice_metaObject(this.h))
}

func (this *QInputDevice) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QInputDevice_metacast(this.h, param1_Cstring))
}

func QInputDevice_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QInputDevice_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QInputDevice) Name() string {
	var _ms C.struct_miqt_string = C.QInputDevice_name(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QInputDevice) Type() QInputDevice__DeviceType {
	return (QInputDevice__DeviceType)(C.QInputDevice_type(this.h))
}

func (this *QInputDevice) Capabilities() QInputDevice__Capability {
	return (QInputDevice__Capability)(C.QInputDevice_capabilities(this.h))
}

func (this *QInputDevice) HasCapability(cap QInputDevice__Capability) bool {
	return (bool)(C.QInputDevice_hasCapability(this.h, (C.int)(cap)))
}

func (this *QInputDevice) SystemId() int64 {
	return (int64)(C.QInputDevice_systemId(this.h))
}

func (this *QInputDevice) SeatName() string {
	var _ms C.struct_miqt_string = C.QInputDevice_seatName(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QInputDevice) AvailableVirtualGeometry() *QRect {
	_goptr := newQRect(C.QInputDevice_availableVirtualGeometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QInputDevice_SeatNames() []string {
	var _ma C.struct_miqt_array = C.QInputDevice_seatNames()
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]C.struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms C.struct_miqt_string = _outCast[i]
		_lv_ret := C.GoStringN(_lv_ms.data, C.int(int64(_lv_ms.len)))
		C.free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QInputDevice_Devices() []*QInputDevice {
	var _ma C.struct_miqt_array = C.QInputDevice_devices()
	_ret := make([]*QInputDevice, int(_ma.len))
	_outCast := (*[0xffff]*C.QInputDevice)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQInputDevice(_outCast[i])
	}
	return _ret
}

func QInputDevice_PrimaryKeyboard() *QInputDevice {
	return newQInputDevice(C.QInputDevice_primaryKeyboard())
}

func (this *QInputDevice) OperatorEqual(other *QInputDevice) bool {
	return (bool)(C.QInputDevice_operatorEqual(this.h, other.cPointer()))
}

func (this *QInputDevice) AvailableVirtualGeometryChanged(area QRect) {
	C.QInputDevice_availableVirtualGeometryChanged(this.h, area.cPointer())
}
func (this *QInputDevice) OnAvailableVirtualGeometryChanged(slot func(area QRect)) {
	C.QInputDevice_connect_availableVirtualGeometryChanged(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputDevice_availableVirtualGeometryChanged
func miqt_exec_callback_QInputDevice_availableVirtualGeometryChanged(cb C.intptr_t, area *C.QRect) {
	gofunc, ok := cgo.Handle(cb).Value().(func(area QRect))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	area_goptr := newQRect(area)
	area_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval1 := *area_goptr

	gofunc(slotval1)
}

func QInputDevice_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QInputDevice_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDevice_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QInputDevice_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDevice_PrimaryKeyboardWithSeatName(seatName string) *QInputDevice {
	seatName_ms := C.struct_miqt_string{}
	seatName_ms.data = C.CString(seatName)
	seatName_ms.len = C.size_t(len(seatName))
	defer C.free(unsafe.Pointer(seatName_ms.data))
	return newQInputDevice(C.QInputDevice_primaryKeyboardWithSeatName(seatName_ms))
}

// Sender can only be called from a QInputDevice that was directly constructed.
func (this *QInputDevice) Sender() *QObject {

	var _dynamic_cast_ok C.bool = false
	_method_ret := newQObject(C.QInputDevice_protectedbase_sender(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// SenderSignalIndex can only be called from a QInputDevice that was directly constructed.
func (this *QInputDevice) SenderSignalIndex() int {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QInputDevice_protectedbase_senderSignalIndex(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// Receivers can only be called from a QInputDevice that was directly constructed.
func (this *QInputDevice) Receivers(signal string) int {
	signal_Cstring := C.CString(signal)
	defer C.free(unsafe.Pointer(signal_Cstring))

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QInputDevice_protectedbase_receivers(&_dynamic_cast_ok, unsafe.Pointer(this.h), signal_Cstring))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// IsSignalConnected can only be called from a QInputDevice that was directly constructed.
func (this *QInputDevice) IsSignalConnected(signal *QMetaMethod) bool {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QInputDevice_protectedbase_isSignalConnected(&_dynamic_cast_ok, unsafe.Pointer(this.h), signal.cPointer()))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

func (this *QInputDevice) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(C.QInputDevice_virtualbase_event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QInputDevice) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	ok := C.QInputDevice_override_virtual_event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QInputDevice_event
func miqt_exec_callback_QInputDevice_event(self *C.QInputDevice, cb C.intptr_t, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QInputDevice{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QInputDevice) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(C.QInputDevice_virtualbase_eventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QInputDevice) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	ok := C.QInputDevice_override_virtual_eventFilter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QInputDevice_eventFilter
func miqt_exec_callback_QInputDevice_eventFilter(self *C.QInputDevice, cb C.intptr_t, watched *C.QObject, event *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QInputDevice{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QInputDevice) callVirtualBase_TimerEvent(event *QTimerEvent) {

	C.QInputDevice_virtualbase_timerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QInputDevice) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	ok := C.QInputDevice_override_virtual_timerEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QInputDevice_timerEvent
func miqt_exec_callback_QInputDevice_timerEvent(self *C.QInputDevice, cb C.intptr_t, event *C.QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QInputDevice{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QInputDevice) callVirtualBase_ChildEvent(event *QChildEvent) {

	C.QInputDevice_virtualbase_childEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QInputDevice) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	ok := C.QInputDevice_override_virtual_childEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QInputDevice_childEvent
func miqt_exec_callback_QInputDevice_childEvent(self *C.QInputDevice, cb C.intptr_t, event *C.QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QInputDevice{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QInputDevice) callVirtualBase_CustomEvent(event *QEvent) {

	C.QInputDevice_virtualbase_customEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QInputDevice) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	ok := C.QInputDevice_override_virtual_customEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QInputDevice_customEvent
func miqt_exec_callback_QInputDevice_customEvent(self *C.QInputDevice, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QInputDevice{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QInputDevice) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	C.QInputDevice_virtualbase_connectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QInputDevice) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	ok := C.QInputDevice_override_virtual_connectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QInputDevice_connectNotify
func miqt_exec_callback_QInputDevice_connectNotify(self *C.QInputDevice, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QInputDevice{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QInputDevice) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	C.QInputDevice_virtualbase_disconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QInputDevice) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	ok := C.QInputDevice_override_virtual_disconnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QInputDevice_disconnectNotify
func miqt_exec_callback_QInputDevice_disconnectNotify(self *C.QInputDevice, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QInputDevice{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

// Delete this object from C++ memory.
func (this *QInputDevice) Delete() {
	C.QInputDevice_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QInputDevice) GoGC() {
	runtime.SetFinalizer(this, func(this *QInputDevice) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
