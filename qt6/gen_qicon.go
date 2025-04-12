package qt6

/*

#include "gen_qicon.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QIcon__Mode int

const (
	QIcon__Normal   QIcon__Mode = 0
	QIcon__Disabled QIcon__Mode = 1
	QIcon__Active   QIcon__Mode = 2
	QIcon__Selected QIcon__Mode = 3
)

type QIcon__State int

const (
	QIcon__On  QIcon__State = 0
	QIcon__Off QIcon__State = 1
)

type QIcon struct {
	h *C.QIcon
}

func (this *QIcon) cPointer() *C.QIcon {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QIcon) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQIcon constructs the type using only CGO pointers.
func newQIcon(h *C.QIcon) *QIcon {
	if h == nil {
		return nil
	}

	return &QIcon{h: h}
}

// UnsafeNewQIcon constructs the type using only unsafe pointers.
func UnsafeNewQIcon(h unsafe.Pointer) *QIcon {
	return newQIcon((*C.QIcon)(h))
}

// NewQIcon constructs a new QIcon object.
func NewQIcon() *QIcon {

	return newQIcon(C.QIcon_new())
}

// NewQIcon2 constructs a new QIcon object.
func NewQIcon2(pixmap *QPixmap) *QIcon {

	return newQIcon(C.QIcon_new2(pixmap.cPointer()))
}

// NewQIcon3 constructs a new QIcon object.
func NewQIcon3(other *QIcon) *QIcon {

	return newQIcon(C.QIcon_new3(other.cPointer()))
}

// NewQIcon4 constructs a new QIcon object.
func NewQIcon4(fileName string) *QIcon {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))

	return newQIcon(C.QIcon_new4(fileName_ms))
}

// NewQIcon5 constructs a new QIcon object.
func NewQIcon5(engine *QIconEngine) *QIcon {

	return newQIcon(C.QIcon_new5(engine.cPointer()))
}

func (this *QIcon) OperatorAssign(other *QIcon) {
	C.QIcon_operatorAssign(this.h, other.cPointer())
}

func (this *QIcon) Swap(other *QIcon) {
	C.QIcon_swap(this.h, other.cPointer())
}

func (this *QIcon) ToQVariant() *QVariant {
	_goptr := newQVariant(C.QIcon_ToQVariant(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap(size *QSize) *QPixmap {
	_goptr := newQPixmap(C.QIcon_pixmap(this.h, size.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap2(w int, h int) *QPixmap {
	_goptr := newQPixmap(C.QIcon_pixmap2(this.h, (C.int)(w), (C.int)(h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) PixmapWithExtent(extent int) *QPixmap {
	_goptr := newQPixmap(C.QIcon_pixmapWithExtent(this.h, (C.int)(extent)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap3(size *QSize, devicePixelRatio float64) *QPixmap {
	_goptr := newQPixmap(C.QIcon_pixmap3(this.h, size.cPointer(), (C.double)(devicePixelRatio)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap4(window *QWindow, size *QSize) *QPixmap {
	_goptr := newQPixmap(C.QIcon_pixmap4(this.h, window.cPointer(), size.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) ActualSize(size *QSize) *QSize {
	_goptr := newQSize(C.QIcon_actualSize(this.h, size.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) ActualSize2(window *QWindow, size *QSize) *QSize {
	_goptr := newQSize(C.QIcon_actualSize2(this.h, window.cPointer(), size.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Name() string {
	var _ms C.struct_miqt_string = C.QIcon_name(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QIcon) Paint(painter *QPainter, rect *QRect) {
	C.QIcon_paint(this.h, painter.cPointer(), rect.cPointer())
}

func (this *QIcon) Paint2(painter *QPainter, x int, y int, w int, h int) {
	C.QIcon_paint2(this.h, painter.cPointer(), (C.int)(x), (C.int)(y), (C.int)(w), (C.int)(h))
}

func (this *QIcon) IsNull() bool {
	return (bool)(C.QIcon_isNull(this.h))
}

func (this *QIcon) IsDetached() bool {
	return (bool)(C.QIcon_isDetached(this.h))
}

func (this *QIcon) Detach() {
	C.QIcon_detach(this.h)
}

func (this *QIcon) CacheKey() int64 {
	return (int64)(C.QIcon_cacheKey(this.h))
}

func (this *QIcon) AddPixmap(pixmap *QPixmap) {
	C.QIcon_addPixmap(this.h, pixmap.cPointer())
}

func (this *QIcon) AddFile(fileName string) {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	C.QIcon_addFile(this.h, fileName_ms)
}

func (this *QIcon) AvailableSizes() []QSize {
	var _ma C.struct_miqt_array = C.QIcon_availableSizes(this.h)
	_ret := make([]QSize, int(_ma.len))
	_outCast := (*[0xffff]*C.QSize)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSize(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QIcon) SetIsMask(isMask bool) {
	C.QIcon_setIsMask(this.h, (C.bool)(isMask))
}

func (this *QIcon) IsMask() bool {
	return (bool)(C.QIcon_isMask(this.h))
}

func QIcon_FromTheme(name string) *QIcon {
	name_ms := C.struct_miqt_string{}
	name_ms.data = C.CString(name)
	name_ms.len = C.size_t(len(name))
	defer C.free(unsafe.Pointer(name_ms.data))
	_goptr := newQIcon(C.QIcon_fromTheme(name_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QIcon_FromTheme2(name string, fallback *QIcon) *QIcon {
	name_ms := C.struct_miqt_string{}
	name_ms.data = C.CString(name)
	name_ms.len = C.size_t(len(name))
	defer C.free(unsafe.Pointer(name_ms.data))
	_goptr := newQIcon(C.QIcon_fromTheme2(name_ms, fallback.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QIcon_HasThemeIcon(name string) bool {
	name_ms := C.struct_miqt_string{}
	name_ms.data = C.CString(name)
	name_ms.len = C.size_t(len(name))
	defer C.free(unsafe.Pointer(name_ms.data))
	return (bool)(C.QIcon_hasThemeIcon(name_ms))
}

func QIcon_ThemeSearchPaths() []string {
	var _ma C.struct_miqt_array = C.QIcon_themeSearchPaths()
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

func QIcon_SetThemeSearchPaths(searchpath []string) {
	searchpath_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(searchpath))))
	defer C.free(unsafe.Pointer(searchpath_CArray))
	for i := range searchpath {
		searchpath_i_ms := C.struct_miqt_string{}
		searchpath_i_ms.data = C.CString(searchpath[i])
		searchpath_i_ms.len = C.size_t(len(searchpath[i]))
		defer C.free(unsafe.Pointer(searchpath_i_ms.data))
		searchpath_CArray[i] = searchpath_i_ms
	}
	searchpath_ma := C.struct_miqt_array{len: C.size_t(len(searchpath)), data: unsafe.Pointer(searchpath_CArray)}
	C.QIcon_setThemeSearchPaths(searchpath_ma)
}

func QIcon_FallbackSearchPaths() []string {
	var _ma C.struct_miqt_array = C.QIcon_fallbackSearchPaths()
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

func QIcon_SetFallbackSearchPaths(paths []string) {
	paths_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(paths))))
	defer C.free(unsafe.Pointer(paths_CArray))
	for i := range paths {
		paths_i_ms := C.struct_miqt_string{}
		paths_i_ms.data = C.CString(paths[i])
		paths_i_ms.len = C.size_t(len(paths[i]))
		defer C.free(unsafe.Pointer(paths_i_ms.data))
		paths_CArray[i] = paths_i_ms
	}
	paths_ma := C.struct_miqt_array{len: C.size_t(len(paths)), data: unsafe.Pointer(paths_CArray)}
	C.QIcon_setFallbackSearchPaths(paths_ma)
}

func QIcon_ThemeName() string {
	var _ms C.struct_miqt_string = C.QIcon_themeName()
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QIcon_SetThemeName(path string) {
	path_ms := C.struct_miqt_string{}
	path_ms.data = C.CString(path)
	path_ms.len = C.size_t(len(path))
	defer C.free(unsafe.Pointer(path_ms.data))
	C.QIcon_setThemeName(path_ms)
}

func QIcon_FallbackThemeName() string {
	var _ms C.struct_miqt_string = C.QIcon_fallbackThemeName()
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QIcon_SetFallbackThemeName(name string) {
	name_ms := C.struct_miqt_string{}
	name_ms.data = C.CString(name)
	name_ms.len = C.size_t(len(name))
	defer C.free(unsafe.Pointer(name_ms.data))
	C.QIcon_setFallbackThemeName(name_ms)
}

func (this *QIcon) Pixmap5(size *QSize, mode QIcon__Mode) *QPixmap {
	_goptr := newQPixmap(C.QIcon_pixmap5(this.h, size.cPointer(), (C.int)(mode)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap6(size *QSize, mode QIcon__Mode, state QIcon__State) *QPixmap {
	_goptr := newQPixmap(C.QIcon_pixmap6(this.h, size.cPointer(), (C.int)(mode), (C.int)(state)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap7(w int, h int, mode QIcon__Mode) *QPixmap {
	_goptr := newQPixmap(C.QIcon_pixmap7(this.h, (C.int)(w), (C.int)(h), (C.int)(mode)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap8(w int, h int, mode QIcon__Mode, state QIcon__State) *QPixmap {
	_goptr := newQPixmap(C.QIcon_pixmap8(this.h, (C.int)(w), (C.int)(h), (C.int)(mode), (C.int)(state)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap9(extent int, mode QIcon__Mode) *QPixmap {
	_goptr := newQPixmap(C.QIcon_pixmap9(this.h, (C.int)(extent), (C.int)(mode)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap10(extent int, mode QIcon__Mode, state QIcon__State) *QPixmap {
	_goptr := newQPixmap(C.QIcon_pixmap10(this.h, (C.int)(extent), (C.int)(mode), (C.int)(state)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap11(size *QSize, devicePixelRatio float64, mode QIcon__Mode) *QPixmap {
	_goptr := newQPixmap(C.QIcon_pixmap11(this.h, size.cPointer(), (C.double)(devicePixelRatio), (C.int)(mode)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap12(size *QSize, devicePixelRatio float64, mode QIcon__Mode, state QIcon__State) *QPixmap {
	_goptr := newQPixmap(C.QIcon_pixmap12(this.h, size.cPointer(), (C.double)(devicePixelRatio), (C.int)(mode), (C.int)(state)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap13(window *QWindow, size *QSize, mode QIcon__Mode) *QPixmap {
	_goptr := newQPixmap(C.QIcon_pixmap13(this.h, window.cPointer(), size.cPointer(), (C.int)(mode)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Pixmap14(window *QWindow, size *QSize, mode QIcon__Mode, state QIcon__State) *QPixmap {
	_goptr := newQPixmap(C.QIcon_pixmap14(this.h, window.cPointer(), size.cPointer(), (C.int)(mode), (C.int)(state)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) ActualSize3(size *QSize, mode QIcon__Mode) *QSize {
	_goptr := newQSize(C.QIcon_actualSize3(this.h, size.cPointer(), (C.int)(mode)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) ActualSize4(size *QSize, mode QIcon__Mode, state QIcon__State) *QSize {
	_goptr := newQSize(C.QIcon_actualSize4(this.h, size.cPointer(), (C.int)(mode), (C.int)(state)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) ActualSize5(window *QWindow, size *QSize, mode QIcon__Mode) *QSize {
	_goptr := newQSize(C.QIcon_actualSize5(this.h, window.cPointer(), size.cPointer(), (C.int)(mode)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) ActualSize6(window *QWindow, size *QSize, mode QIcon__Mode, state QIcon__State) *QSize {
	_goptr := newQSize(C.QIcon_actualSize6(this.h, window.cPointer(), size.cPointer(), (C.int)(mode), (C.int)(state)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIcon) Paint3(painter *QPainter, rect *QRect, alignment AlignmentFlag) {
	C.QIcon_paint3(this.h, painter.cPointer(), rect.cPointer(), (C.int)(alignment))
}

func (this *QIcon) Paint4(painter *QPainter, rect *QRect, alignment AlignmentFlag, mode QIcon__Mode) {
	C.QIcon_paint4(this.h, painter.cPointer(), rect.cPointer(), (C.int)(alignment), (C.int)(mode))
}

func (this *QIcon) Paint5(painter *QPainter, rect *QRect, alignment AlignmentFlag, mode QIcon__Mode, state QIcon__State) {
	C.QIcon_paint5(this.h, painter.cPointer(), rect.cPointer(), (C.int)(alignment), (C.int)(mode), (C.int)(state))
}

func (this *QIcon) Paint6(painter *QPainter, x int, y int, w int, h int, alignment AlignmentFlag) {
	C.QIcon_paint6(this.h, painter.cPointer(), (C.int)(x), (C.int)(y), (C.int)(w), (C.int)(h), (C.int)(alignment))
}

func (this *QIcon) Paint7(painter *QPainter, x int, y int, w int, h int, alignment AlignmentFlag, mode QIcon__Mode) {
	C.QIcon_paint7(this.h, painter.cPointer(), (C.int)(x), (C.int)(y), (C.int)(w), (C.int)(h), (C.int)(alignment), (C.int)(mode))
}

func (this *QIcon) Paint8(painter *QPainter, x int, y int, w int, h int, alignment AlignmentFlag, mode QIcon__Mode, state QIcon__State) {
	C.QIcon_paint8(this.h, painter.cPointer(), (C.int)(x), (C.int)(y), (C.int)(w), (C.int)(h), (C.int)(alignment), (C.int)(mode), (C.int)(state))
}

func (this *QIcon) AddPixmap2(pixmap *QPixmap, mode QIcon__Mode) {
	C.QIcon_addPixmap2(this.h, pixmap.cPointer(), (C.int)(mode))
}

func (this *QIcon) AddPixmap3(pixmap *QPixmap, mode QIcon__Mode, state QIcon__State) {
	C.QIcon_addPixmap3(this.h, pixmap.cPointer(), (C.int)(mode), (C.int)(state))
}

func (this *QIcon) AddFile2(fileName string, size *QSize) {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	C.QIcon_addFile2(this.h, fileName_ms, size.cPointer())
}

func (this *QIcon) AddFile3(fileName string, size *QSize, mode QIcon__Mode) {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	C.QIcon_addFile3(this.h, fileName_ms, size.cPointer(), (C.int)(mode))
}

func (this *QIcon) AddFile4(fileName string, size *QSize, mode QIcon__Mode, state QIcon__State) {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	C.QIcon_addFile4(this.h, fileName_ms, size.cPointer(), (C.int)(mode), (C.int)(state))
}

func (this *QIcon) AvailableSizesWithMode(mode QIcon__Mode) []QSize {
	var _ma C.struct_miqt_array = C.QIcon_availableSizesWithMode(this.h, (C.int)(mode))
	_ret := make([]QSize, int(_ma.len))
	_outCast := (*[0xffff]*C.QSize)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSize(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QIcon) AvailableSizes2(mode QIcon__Mode, state QIcon__State) []QSize {
	var _ma C.struct_miqt_array = C.QIcon_availableSizes2(this.h, (C.int)(mode), (C.int)(state))
	_ret := make([]QSize, int(_ma.len))
	_outCast := (*[0xffff]*C.QSize)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSize(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

// Delete this object from C++ memory.
func (this *QIcon) Delete() {
	C.QIcon_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QIcon) GoGC() {
	runtime.SetFinalizer(this, func(this *QIcon) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
