package qt

/*

#include "gen_qcompleter.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QCompleter__CompletionMode int

const (
	QCompleter__PopupCompletion           QCompleter__CompletionMode = 0
	QCompleter__UnfilteredPopupCompletion QCompleter__CompletionMode = 1
	QCompleter__InlineCompletion          QCompleter__CompletionMode = 2
)

type QCompleter__ModelSorting int

const (
	QCompleter__UnsortedModel                QCompleter__ModelSorting = 0
	QCompleter__CaseSensitivelySortedModel   QCompleter__ModelSorting = 1
	QCompleter__CaseInsensitivelySortedModel QCompleter__ModelSorting = 2
)

type QCompleter struct {
	h *C.QCompleter
	*QObject
}

func (this *QCompleter) cPointer() *C.QCompleter {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QCompleter) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQCompleter constructs the type using only CGO pointers.
func newQCompleter(h *C.QCompleter) *QCompleter {
	if h == nil {
		return nil
	}
	var outptr_QObject *C.QObject = nil
	C.QCompleter_virtbase(h, &outptr_QObject)

	return &QCompleter{h: h,
		QObject: newQObject(outptr_QObject)}
}

// UnsafeNewQCompleter constructs the type using only unsafe pointers.
func UnsafeNewQCompleter(h unsafe.Pointer) *QCompleter {
	return newQCompleter((*C.QCompleter)(h))
}

// NewQCompleter constructs a new QCompleter object.
func NewQCompleter() *QCompleter {

	return newQCompleter(C.QCompleter_new())
}

// NewQCompleter2 constructs a new QCompleter object.
func NewQCompleter2(model *QAbstractItemModel) *QCompleter {

	return newQCompleter(C.QCompleter_new2(model.cPointer()))
}

// NewQCompleter3 constructs a new QCompleter object.
func NewQCompleter3(completions []string) *QCompleter {
	completions_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(completions))))
	defer C.free(unsafe.Pointer(completions_CArray))
	for i := range completions {
		completions_i_ms := C.struct_miqt_string{}
		completions_i_ms.data = C.CString(completions[i])
		completions_i_ms.len = C.size_t(len(completions[i]))
		defer C.free(unsafe.Pointer(completions_i_ms.data))
		completions_CArray[i] = completions_i_ms
	}
	completions_ma := C.struct_miqt_array{len: C.size_t(len(completions)), data: unsafe.Pointer(completions_CArray)}

	return newQCompleter(C.QCompleter_new3(completions_ma))
}

// NewQCompleter4 constructs a new QCompleter object.
func NewQCompleter4(parent *QObject) *QCompleter {

	return newQCompleter(C.QCompleter_new4(parent.cPointer()))
}

// NewQCompleter5 constructs a new QCompleter object.
func NewQCompleter5(model *QAbstractItemModel, parent *QObject) *QCompleter {

	return newQCompleter(C.QCompleter_new5(model.cPointer(), parent.cPointer()))
}

// NewQCompleter6 constructs a new QCompleter object.
func NewQCompleter6(completions []string, parent *QObject) *QCompleter {
	completions_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(completions))))
	defer C.free(unsafe.Pointer(completions_CArray))
	for i := range completions {
		completions_i_ms := C.struct_miqt_string{}
		completions_i_ms.data = C.CString(completions[i])
		completions_i_ms.len = C.size_t(len(completions[i]))
		defer C.free(unsafe.Pointer(completions_i_ms.data))
		completions_CArray[i] = completions_i_ms
	}
	completions_ma := C.struct_miqt_array{len: C.size_t(len(completions)), data: unsafe.Pointer(completions_CArray)}

	return newQCompleter(C.QCompleter_new6(completions_ma, parent.cPointer()))
}

func (this *QCompleter) MetaObject() *QMetaObject {
	return newQMetaObject(C.QCompleter_metaObject(this.h))
}

func (this *QCompleter) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QCompleter_metacast(this.h, param1_Cstring))
}

func QCompleter_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QCompleter_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCompleter_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QCompleter_trUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCompleter) SetWidget(widget *QWidget) {
	C.QCompleter_setWidget(this.h, widget.cPointer())
}

func (this *QCompleter) Widget() *QWidget {
	return newQWidget(C.QCompleter_widget(this.h))
}

func (this *QCompleter) SetModel(c *QAbstractItemModel) {
	C.QCompleter_setModel(this.h, c.cPointer())
}

func (this *QCompleter) Model() *QAbstractItemModel {
	return newQAbstractItemModel(C.QCompleter_model(this.h))
}

func (this *QCompleter) SetCompletionMode(mode QCompleter__CompletionMode) {
	C.QCompleter_setCompletionMode(this.h, (C.int)(mode))
}

func (this *QCompleter) CompletionMode() QCompleter__CompletionMode {
	return (QCompleter__CompletionMode)(C.QCompleter_completionMode(this.h))
}

func (this *QCompleter) SetFilterMode(filterMode MatchFlag) {
	C.QCompleter_setFilterMode(this.h, (C.int)(filterMode))
}

func (this *QCompleter) FilterMode() MatchFlag {
	return (MatchFlag)(C.QCompleter_filterMode(this.h))
}

func (this *QCompleter) Popup() *QAbstractItemView {
	return newQAbstractItemView(C.QCompleter_popup(this.h))
}

func (this *QCompleter) SetPopup(popup *QAbstractItemView) {
	C.QCompleter_setPopup(this.h, popup.cPointer())
}

func (this *QCompleter) SetCaseSensitivity(caseSensitivity CaseSensitivity) {
	C.QCompleter_setCaseSensitivity(this.h, (C.int)(caseSensitivity))
}

func (this *QCompleter) CaseSensitivity() CaseSensitivity {
	return (CaseSensitivity)(C.QCompleter_caseSensitivity(this.h))
}

func (this *QCompleter) SetModelSorting(sorting QCompleter__ModelSorting) {
	C.QCompleter_setModelSorting(this.h, (C.int)(sorting))
}

func (this *QCompleter) ModelSorting() QCompleter__ModelSorting {
	return (QCompleter__ModelSorting)(C.QCompleter_modelSorting(this.h))
}

func (this *QCompleter) SetCompletionColumn(column int) {
	C.QCompleter_setCompletionColumn(this.h, (C.int)(column))
}

func (this *QCompleter) CompletionColumn() int {
	return (int)(C.QCompleter_completionColumn(this.h))
}

func (this *QCompleter) SetCompletionRole(role int) {
	C.QCompleter_setCompletionRole(this.h, (C.int)(role))
}

func (this *QCompleter) CompletionRole() int {
	return (int)(C.QCompleter_completionRole(this.h))
}

func (this *QCompleter) WrapAround() bool {
	return (bool)(C.QCompleter_wrapAround(this.h))
}

func (this *QCompleter) MaxVisibleItems() int {
	return (int)(C.QCompleter_maxVisibleItems(this.h))
}

func (this *QCompleter) SetMaxVisibleItems(maxItems int) {
	C.QCompleter_setMaxVisibleItems(this.h, (C.int)(maxItems))
}

func (this *QCompleter) CompletionCount() int {
	return (int)(C.QCompleter_completionCount(this.h))
}

func (this *QCompleter) SetCurrentRow(row int) bool {
	return (bool)(C.QCompleter_setCurrentRow(this.h, (C.int)(row)))
}

func (this *QCompleter) CurrentRow() int {
	return (int)(C.QCompleter_currentRow(this.h))
}

func (this *QCompleter) CurrentIndex() *QModelIndex {
	_goptr := newQModelIndex(C.QCompleter_currentIndex(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCompleter) CurrentCompletion() string {
	var _ms C.struct_miqt_string = C.QCompleter_currentCompletion(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCompleter) CompletionModel() *QAbstractItemModel {
	return newQAbstractItemModel(C.QCompleter_completionModel(this.h))
}

func (this *QCompleter) CompletionPrefix() string {
	var _ms C.struct_miqt_string = C.QCompleter_completionPrefix(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCompleter) SetCompletionPrefix(prefix string) {
	prefix_ms := C.struct_miqt_string{}
	prefix_ms.data = C.CString(prefix)
	prefix_ms.len = C.size_t(len(prefix))
	defer C.free(unsafe.Pointer(prefix_ms.data))
	C.QCompleter_setCompletionPrefix(this.h, prefix_ms)
}

func (this *QCompleter) Complete() {
	C.QCompleter_complete(this.h)
}

func (this *QCompleter) SetWrapAround(wrap bool) {
	C.QCompleter_setWrapAround(this.h, (C.bool)(wrap))
}

func (this *QCompleter) PathFromIndex(index *QModelIndex) string {
	var _ms C.struct_miqt_string = C.QCompleter_pathFromIndex(this.h, index.cPointer())
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCompleter) SplitPath(path string) []string {
	path_ms := C.struct_miqt_string{}
	path_ms.data = C.CString(path)
	path_ms.len = C.size_t(len(path))
	defer C.free(unsafe.Pointer(path_ms.data))
	var _ma C.struct_miqt_array = C.QCompleter_splitPath(this.h, path_ms)
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

func (this *QCompleter) Activated(text string) {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	C.QCompleter_activated(this.h, text_ms)
}
func (this *QCompleter) OnActivated(slot func(text string)) {
	C.QCompleter_connect_activated(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCompleter_activated
func miqt_exec_callback_QCompleter_activated(cb C.intptr_t, text C.struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(text string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var text_ms C.struct_miqt_string = text
	text_ret := C.GoStringN(text_ms.data, C.int(int64(text_ms.len)))
	C.free(unsafe.Pointer(text_ms.data))
	slotval1 := text_ret

	gofunc(slotval1)
}

func (this *QCompleter) ActivatedWithIndex(index *QModelIndex) {
	C.QCompleter_activatedWithIndex(this.h, index.cPointer())
}
func (this *QCompleter) OnActivatedWithIndex(slot func(index *QModelIndex)) {
	C.QCompleter_connect_activatedWithIndex(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCompleter_activatedWithIndex
func miqt_exec_callback_QCompleter_activatedWithIndex(cb C.intptr_t, index *C.QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	gofunc(slotval1)
}

func (this *QCompleter) Highlighted(text string) {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	C.QCompleter_highlighted(this.h, text_ms)
}
func (this *QCompleter) OnHighlighted(slot func(text string)) {
	C.QCompleter_connect_highlighted(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCompleter_highlighted
func miqt_exec_callback_QCompleter_highlighted(cb C.intptr_t, text C.struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(text string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var text_ms C.struct_miqt_string = text
	text_ret := C.GoStringN(text_ms.data, C.int(int64(text_ms.len)))
	C.free(unsafe.Pointer(text_ms.data))
	slotval1 := text_ret

	gofunc(slotval1)
}

func (this *QCompleter) HighlightedWithIndex(index *QModelIndex) {
	C.QCompleter_highlightedWithIndex(this.h, index.cPointer())
}
func (this *QCompleter) OnHighlightedWithIndex(slot func(index *QModelIndex)) {
	C.QCompleter_connect_highlightedWithIndex(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCompleter_highlightedWithIndex
func miqt_exec_callback_QCompleter_highlightedWithIndex(cb C.intptr_t, index *C.QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	gofunc(slotval1)
}

func QCompleter_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QCompleter_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCompleter_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QCompleter_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCompleter_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QCompleter_trUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCompleter_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QCompleter_trUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCompleter) CompleteWithRect(rect *QRect) {
	C.QCompleter_completeWithRect(this.h, rect.cPointer())
}

// Sender can only be called from a QCompleter that was directly constructed.
func (this *QCompleter) Sender() *QObject {

	var _dynamic_cast_ok C.bool = false
	_method_ret := newQObject(C.QCompleter_protectedbase_sender(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// SenderSignalIndex can only be called from a QCompleter that was directly constructed.
func (this *QCompleter) SenderSignalIndex() int {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QCompleter_protectedbase_senderSignalIndex(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// Receivers can only be called from a QCompleter that was directly constructed.
func (this *QCompleter) Receivers(signal string) int {
	signal_Cstring := C.CString(signal)
	defer C.free(unsafe.Pointer(signal_Cstring))

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QCompleter_protectedbase_receivers(&_dynamic_cast_ok, unsafe.Pointer(this.h), signal_Cstring))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// IsSignalConnected can only be called from a QCompleter that was directly constructed.
func (this *QCompleter) IsSignalConnected(signal *QMetaMethod) bool {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QCompleter_protectedbase_isSignalConnected(&_dynamic_cast_ok, unsafe.Pointer(this.h), signal.cPointer()))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

func (this *QCompleter) callVirtualBase_PathFromIndex(index *QModelIndex) string {

	var _ms C.struct_miqt_string = C.QCompleter_virtualbase_pathFromIndex(unsafe.Pointer(this.h), index.cPointer())
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QCompleter) OnPathFromIndex(slot func(super func(index *QModelIndex) string, index *QModelIndex) string) {
	ok := C.QCompleter_override_virtual_pathFromIndex(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QCompleter_pathFromIndex
func miqt_exec_callback_QCompleter_pathFromIndex(self *C.QCompleter, cb C.intptr_t, index *C.QModelIndex) C.struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) string, index *QModelIndex) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QCompleter{h: self}).callVirtualBase_PathFromIndex, slotval1)
	virtualReturn_ms := C.struct_miqt_string{}
	virtualReturn_ms.data = C.CString(virtualReturn)
	virtualReturn_ms.len = C.size_t(len(virtualReturn))
	defer C.free(unsafe.Pointer(virtualReturn_ms.data))

	return virtualReturn_ms

}

func (this *QCompleter) callVirtualBase_SplitPath(path string) []string {
	path_ms := C.struct_miqt_string{}
	path_ms.data = C.CString(path)
	path_ms.len = C.size_t(len(path))
	defer C.free(unsafe.Pointer(path_ms.data))

	var _ma C.struct_miqt_array = C.QCompleter_virtualbase_splitPath(unsafe.Pointer(this.h), path_ms)
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
func (this *QCompleter) OnSplitPath(slot func(super func(path string) []string, path string) []string) {
	ok := C.QCompleter_override_virtual_splitPath(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QCompleter_splitPath
func miqt_exec_callback_QCompleter_splitPath(self *C.QCompleter, cb C.intptr_t, path C.struct_miqt_string) C.struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(path string) []string, path string) []string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var path_ms C.struct_miqt_string = path
	path_ret := C.GoStringN(path_ms.data, C.int(int64(path_ms.len)))
	C.free(unsafe.Pointer(path_ms.data))
	slotval1 := path_ret

	virtualReturn := gofunc((&QCompleter{h: self}).callVirtualBase_SplitPath, slotval1)
	virtualReturn_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(virtualReturn))))
	defer C.free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_i_ms := C.struct_miqt_string{}
		virtualReturn_i_ms.data = C.CString(virtualReturn[i])
		virtualReturn_i_ms.len = C.size_t(len(virtualReturn[i]))
		defer C.free(unsafe.Pointer(virtualReturn_i_ms.data))
		virtualReturn_CArray[i] = virtualReturn_i_ms
	}
	virtualReturn_ma := C.struct_miqt_array{len: C.size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QCompleter) callVirtualBase_EventFilter(o *QObject, e *QEvent) bool {

	return (bool)(C.QCompleter_virtualbase_eventFilter(unsafe.Pointer(this.h), o.cPointer(), e.cPointer()))

}
func (this *QCompleter) OnEventFilter(slot func(super func(o *QObject, e *QEvent) bool, o *QObject, e *QEvent) bool) {
	ok := C.QCompleter_override_virtual_eventFilter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QCompleter_eventFilter
func miqt_exec_callback_QCompleter_eventFilter(self *C.QCompleter, cb C.intptr_t, o *C.QObject, e *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(o *QObject, e *QEvent) bool, o *QObject, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(o)

	slotval2 := newQEvent(e)

	virtualReturn := gofunc((&QCompleter{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QCompleter) callVirtualBase_Event(param1 *QEvent) bool {

	return (bool)(C.QCompleter_virtualbase_event(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QCompleter) OnEvent(slot func(super func(param1 *QEvent) bool, param1 *QEvent) bool) {
	ok := C.QCompleter_override_virtual_event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QCompleter_event
func miqt_exec_callback_QCompleter_event(self *C.QCompleter, cb C.intptr_t, param1 *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent) bool, param1 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	virtualReturn := gofunc((&QCompleter{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QCompleter) callVirtualBase_TimerEvent(event *QTimerEvent) {

	C.QCompleter_virtualbase_timerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCompleter) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	ok := C.QCompleter_override_virtual_timerEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QCompleter_timerEvent
func miqt_exec_callback_QCompleter_timerEvent(self *C.QCompleter, cb C.intptr_t, event *C.QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QCompleter{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QCompleter) callVirtualBase_ChildEvent(event *QChildEvent) {

	C.QCompleter_virtualbase_childEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCompleter) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	ok := C.QCompleter_override_virtual_childEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QCompleter_childEvent
func miqt_exec_callback_QCompleter_childEvent(self *C.QCompleter, cb C.intptr_t, event *C.QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QCompleter{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QCompleter) callVirtualBase_CustomEvent(event *QEvent) {

	C.QCompleter_virtualbase_customEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCompleter) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	ok := C.QCompleter_override_virtual_customEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QCompleter_customEvent
func miqt_exec_callback_QCompleter_customEvent(self *C.QCompleter, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QCompleter{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QCompleter) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	C.QCompleter_virtualbase_connectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QCompleter) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	ok := C.QCompleter_override_virtual_connectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QCompleter_connectNotify
func miqt_exec_callback_QCompleter_connectNotify(self *C.QCompleter, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QCompleter{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QCompleter) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	C.QCompleter_virtualbase_disconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QCompleter) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	ok := C.QCompleter_override_virtual_disconnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QCompleter_disconnectNotify
func miqt_exec_callback_QCompleter_disconnectNotify(self *C.QCompleter, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QCompleter{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

// Delete this object from C++ memory.
func (this *QCompleter) Delete() {
	C.QCompleter_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QCompleter) GoGC() {
	runtime.SetFinalizer(this, func(this *QCompleter) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
