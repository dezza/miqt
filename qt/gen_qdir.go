package qt

/*

#include "gen_qdir.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QDir__Filter int

const (
	QDir__Dirs           QDir__Filter = 1
	QDir__Files          QDir__Filter = 2
	QDir__Drives         QDir__Filter = 4
	QDir__NoSymLinks     QDir__Filter = 8
	QDir__AllEntries     QDir__Filter = 7
	QDir__TypeMask       QDir__Filter = 15
	QDir__Readable       QDir__Filter = 16
	QDir__Writable       QDir__Filter = 32
	QDir__Executable     QDir__Filter = 64
	QDir__PermissionMask QDir__Filter = 112
	QDir__Modified       QDir__Filter = 128
	QDir__Hidden         QDir__Filter = 256
	QDir__System         QDir__Filter = 512
	QDir__AccessMask     QDir__Filter = 1008
	QDir__AllDirs        QDir__Filter = 1024
	QDir__CaseSensitive  QDir__Filter = 2048
	QDir__NoDot          QDir__Filter = 8192
	QDir__NoDotDot       QDir__Filter = 16384
	QDir__NoDotAndDotDot QDir__Filter = 24576
	QDir__NoFilter       QDir__Filter = -1
)

type QDir__SortFlag int

const (
	QDir__Name        QDir__SortFlag = 0
	QDir__Time        QDir__SortFlag = 1
	QDir__Size        QDir__SortFlag = 2
	QDir__Unsorted    QDir__SortFlag = 3
	QDir__SortByMask  QDir__SortFlag = 3
	QDir__DirsFirst   QDir__SortFlag = 4
	QDir__Reversed    QDir__SortFlag = 8
	QDir__IgnoreCase  QDir__SortFlag = 16
	QDir__DirsLast    QDir__SortFlag = 32
	QDir__LocaleAware QDir__SortFlag = 64
	QDir__Type        QDir__SortFlag = 128
	QDir__NoSort      QDir__SortFlag = -1
)

type QDir struct {
	h *C.QDir
}

func (this *QDir) cPointer() *C.QDir {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QDir) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQDir constructs the type using only CGO pointers.
func newQDir(h *C.QDir) *QDir {
	if h == nil {
		return nil
	}

	return &QDir{h: h}
}

// UnsafeNewQDir constructs the type using only unsafe pointers.
func UnsafeNewQDir(h unsafe.Pointer) *QDir {
	return newQDir((*C.QDir)(h))
}

// NewQDir constructs a new QDir object.
func NewQDir(param1 *QDir) *QDir {

	return newQDir(C.QDir_new(param1.cPointer()))
}

// NewQDir2 constructs a new QDir object.
func NewQDir2() *QDir {

	return newQDir(C.QDir_new2())
}

// NewQDir3 constructs a new QDir object.
func NewQDir3(path string, nameFilter string) *QDir {
	path_ms := C.struct_miqt_string{}
	path_ms.data = C.CString(path)
	path_ms.len = C.size_t(len(path))
	defer C.free(unsafe.Pointer(path_ms.data))
	nameFilter_ms := C.struct_miqt_string{}
	nameFilter_ms.data = C.CString(nameFilter)
	nameFilter_ms.len = C.size_t(len(nameFilter))
	defer C.free(unsafe.Pointer(nameFilter_ms.data))

	return newQDir(C.QDir_new3(path_ms, nameFilter_ms))
}

// NewQDir4 constructs a new QDir object.
func NewQDir4(path string) *QDir {
	path_ms := C.struct_miqt_string{}
	path_ms.data = C.CString(path)
	path_ms.len = C.size_t(len(path))
	defer C.free(unsafe.Pointer(path_ms.data))

	return newQDir(C.QDir_new4(path_ms))
}

// NewQDir5 constructs a new QDir object.
func NewQDir5(path string, nameFilter string, sort QDir__SortFlag) *QDir {
	path_ms := C.struct_miqt_string{}
	path_ms.data = C.CString(path)
	path_ms.len = C.size_t(len(path))
	defer C.free(unsafe.Pointer(path_ms.data))
	nameFilter_ms := C.struct_miqt_string{}
	nameFilter_ms.data = C.CString(nameFilter)
	nameFilter_ms.len = C.size_t(len(nameFilter))
	defer C.free(unsafe.Pointer(nameFilter_ms.data))

	return newQDir(C.QDir_new5(path_ms, nameFilter_ms, (C.int)(sort)))
}

// NewQDir6 constructs a new QDir object.
func NewQDir6(path string, nameFilter string, sort QDir__SortFlag, filter QDir__Filter) *QDir {
	path_ms := C.struct_miqt_string{}
	path_ms.data = C.CString(path)
	path_ms.len = C.size_t(len(path))
	defer C.free(unsafe.Pointer(path_ms.data))
	nameFilter_ms := C.struct_miqt_string{}
	nameFilter_ms.data = C.CString(nameFilter)
	nameFilter_ms.len = C.size_t(len(nameFilter))
	defer C.free(unsafe.Pointer(nameFilter_ms.data))

	return newQDir(C.QDir_new6(path_ms, nameFilter_ms, (C.int)(sort), (C.int)(filter)))
}

func (this *QDir) OperatorAssign(param1 *QDir) {
	C.QDir_operatorAssign(this.h, param1.cPointer())
}

func (this *QDir) OperatorAssignWithPath(path string) {
	path_ms := C.struct_miqt_string{}
	path_ms.data = C.CString(path)
	path_ms.len = C.size_t(len(path))
	defer C.free(unsafe.Pointer(path_ms.data))
	C.QDir_operatorAssignWithPath(this.h, path_ms)
}

func (this *QDir) Swap(other *QDir) {
	C.QDir_swap(this.h, other.cPointer())
}

func (this *QDir) SetPath(path string) {
	path_ms := C.struct_miqt_string{}
	path_ms.data = C.CString(path)
	path_ms.len = C.size_t(len(path))
	defer C.free(unsafe.Pointer(path_ms.data))
	C.QDir_setPath(this.h, path_ms)
}

func (this *QDir) Path() string {
	var _ms C.struct_miqt_string = C.QDir_path(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDir) AbsolutePath() string {
	var _ms C.struct_miqt_string = C.QDir_absolutePath(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDir) CanonicalPath() string {
	var _ms C.struct_miqt_string = C.QDir_canonicalPath(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDir_AddResourceSearchPath(path string) {
	path_ms := C.struct_miqt_string{}
	path_ms.data = C.CString(path)
	path_ms.len = C.size_t(len(path))
	defer C.free(unsafe.Pointer(path_ms.data))
	C.QDir_addResourceSearchPath(path_ms)
}

func QDir_SetSearchPaths(prefix string, searchPaths []string) {
	prefix_ms := C.struct_miqt_string{}
	prefix_ms.data = C.CString(prefix)
	prefix_ms.len = C.size_t(len(prefix))
	defer C.free(unsafe.Pointer(prefix_ms.data))
	searchPaths_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(searchPaths))))
	defer C.free(unsafe.Pointer(searchPaths_CArray))
	for i := range searchPaths {
		searchPaths_i_ms := C.struct_miqt_string{}
		searchPaths_i_ms.data = C.CString(searchPaths[i])
		searchPaths_i_ms.len = C.size_t(len(searchPaths[i]))
		defer C.free(unsafe.Pointer(searchPaths_i_ms.data))
		searchPaths_CArray[i] = searchPaths_i_ms
	}
	searchPaths_ma := C.struct_miqt_array{len: C.size_t(len(searchPaths)), data: unsafe.Pointer(searchPaths_CArray)}
	C.QDir_setSearchPaths(prefix_ms, searchPaths_ma)
}

func QDir_AddSearchPath(prefix string, path string) {
	prefix_ms := C.struct_miqt_string{}
	prefix_ms.data = C.CString(prefix)
	prefix_ms.len = C.size_t(len(prefix))
	defer C.free(unsafe.Pointer(prefix_ms.data))
	path_ms := C.struct_miqt_string{}
	path_ms.data = C.CString(path)
	path_ms.len = C.size_t(len(path))
	defer C.free(unsafe.Pointer(path_ms.data))
	C.QDir_addSearchPath(prefix_ms, path_ms)
}

func QDir_SearchPaths(prefix string) []string {
	prefix_ms := C.struct_miqt_string{}
	prefix_ms.data = C.CString(prefix)
	prefix_ms.len = C.size_t(len(prefix))
	defer C.free(unsafe.Pointer(prefix_ms.data))
	var _ma C.struct_miqt_array = C.QDir_searchPaths(prefix_ms)
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

func (this *QDir) DirName() string {
	var _ms C.struct_miqt_string = C.QDir_dirName(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDir) FilePath(fileName string) string {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	var _ms C.struct_miqt_string = C.QDir_filePath(this.h, fileName_ms)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDir) AbsoluteFilePath(fileName string) string {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	var _ms C.struct_miqt_string = C.QDir_absoluteFilePath(this.h, fileName_ms)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDir) RelativeFilePath(fileName string) string {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	var _ms C.struct_miqt_string = C.QDir_relativeFilePath(this.h, fileName_ms)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDir_ToNativeSeparators(pathName string) string {
	pathName_ms := C.struct_miqt_string{}
	pathName_ms.data = C.CString(pathName)
	pathName_ms.len = C.size_t(len(pathName))
	defer C.free(unsafe.Pointer(pathName_ms.data))
	var _ms C.struct_miqt_string = C.QDir_toNativeSeparators(pathName_ms)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDir_FromNativeSeparators(pathName string) string {
	pathName_ms := C.struct_miqt_string{}
	pathName_ms.data = C.CString(pathName)
	pathName_ms.len = C.size_t(len(pathName))
	defer C.free(unsafe.Pointer(pathName_ms.data))
	var _ms C.struct_miqt_string = C.QDir_fromNativeSeparators(pathName_ms)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDir) Cd(dirName string) bool {
	dirName_ms := C.struct_miqt_string{}
	dirName_ms.data = C.CString(dirName)
	dirName_ms.len = C.size_t(len(dirName))
	defer C.free(unsafe.Pointer(dirName_ms.data))
	return (bool)(C.QDir_cd(this.h, dirName_ms))
}

func (this *QDir) CdUp() bool {
	return (bool)(C.QDir_cdUp(this.h))
}

func (this *QDir) NameFilters() []string {
	var _ma C.struct_miqt_array = C.QDir_nameFilters(this.h)
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

func (this *QDir) SetNameFilters(nameFilters []string) {
	nameFilters_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(nameFilters))))
	defer C.free(unsafe.Pointer(nameFilters_CArray))
	for i := range nameFilters {
		nameFilters_i_ms := C.struct_miqt_string{}
		nameFilters_i_ms.data = C.CString(nameFilters[i])
		nameFilters_i_ms.len = C.size_t(len(nameFilters[i]))
		defer C.free(unsafe.Pointer(nameFilters_i_ms.data))
		nameFilters_CArray[i] = nameFilters_i_ms
	}
	nameFilters_ma := C.struct_miqt_array{len: C.size_t(len(nameFilters)), data: unsafe.Pointer(nameFilters_CArray)}
	C.QDir_setNameFilters(this.h, nameFilters_ma)
}

func (this *QDir) Filter() QDir__Filter {
	return (QDir__Filter)(C.QDir_filter(this.h))
}

func (this *QDir) SetFilter(filter QDir__Filter) {
	C.QDir_setFilter(this.h, (C.int)(filter))
}

func (this *QDir) Sorting() QDir__SortFlag {
	return (QDir__SortFlag)(C.QDir_sorting(this.h))
}

func (this *QDir) SetSorting(sort QDir__SortFlag) {
	C.QDir_setSorting(this.h, (C.int)(sort))
}

func (this *QDir) Count() uint {
	return (uint)(C.QDir_count(this.h))
}

func (this *QDir) IsEmpty() bool {
	return (bool)(C.QDir_isEmpty(this.h))
}

func (this *QDir) OperatorSubscript(param1 int) string {
	var _ms C.struct_miqt_string = C.QDir_operatorSubscript(this.h, (C.int)(param1))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDir_NameFiltersFromString(nameFilter string) []string {
	nameFilter_ms := C.struct_miqt_string{}
	nameFilter_ms.data = C.CString(nameFilter)
	nameFilter_ms.len = C.size_t(len(nameFilter))
	defer C.free(unsafe.Pointer(nameFilter_ms.data))
	var _ma C.struct_miqt_array = C.QDir_nameFiltersFromString(nameFilter_ms)
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

func (this *QDir) EntryList() []string {
	var _ma C.struct_miqt_array = C.QDir_entryList(this.h)
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

func (this *QDir) EntryListWithNameFilters(nameFilters []string) []string {
	nameFilters_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(nameFilters))))
	defer C.free(unsafe.Pointer(nameFilters_CArray))
	for i := range nameFilters {
		nameFilters_i_ms := C.struct_miqt_string{}
		nameFilters_i_ms.data = C.CString(nameFilters[i])
		nameFilters_i_ms.len = C.size_t(len(nameFilters[i]))
		defer C.free(unsafe.Pointer(nameFilters_i_ms.data))
		nameFilters_CArray[i] = nameFilters_i_ms
	}
	nameFilters_ma := C.struct_miqt_array{len: C.size_t(len(nameFilters)), data: unsafe.Pointer(nameFilters_CArray)}
	var _ma C.struct_miqt_array = C.QDir_entryListWithNameFilters(this.h, nameFilters_ma)
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

func (this *QDir) EntryInfoList() []QFileInfo {
	var _ma C.struct_miqt_array = C.QDir_entryInfoList(this.h)
	_ret := make([]QFileInfo, int(_ma.len))
	_outCast := (*[0xffff]*C.QFileInfo)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQFileInfo(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QDir) EntryInfoListWithNameFilters(nameFilters []string) []QFileInfo {
	nameFilters_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(nameFilters))))
	defer C.free(unsafe.Pointer(nameFilters_CArray))
	for i := range nameFilters {
		nameFilters_i_ms := C.struct_miqt_string{}
		nameFilters_i_ms.data = C.CString(nameFilters[i])
		nameFilters_i_ms.len = C.size_t(len(nameFilters[i]))
		defer C.free(unsafe.Pointer(nameFilters_i_ms.data))
		nameFilters_CArray[i] = nameFilters_i_ms
	}
	nameFilters_ma := C.struct_miqt_array{len: C.size_t(len(nameFilters)), data: unsafe.Pointer(nameFilters_CArray)}
	var _ma C.struct_miqt_array = C.QDir_entryInfoListWithNameFilters(this.h, nameFilters_ma)
	_ret := make([]QFileInfo, int(_ma.len))
	_outCast := (*[0xffff]*C.QFileInfo)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQFileInfo(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QDir) Mkdir(dirName string) bool {
	dirName_ms := C.struct_miqt_string{}
	dirName_ms.data = C.CString(dirName)
	dirName_ms.len = C.size_t(len(dirName))
	defer C.free(unsafe.Pointer(dirName_ms.data))
	return (bool)(C.QDir_mkdir(this.h, dirName_ms))
}

func (this *QDir) Rmdir(dirName string) bool {
	dirName_ms := C.struct_miqt_string{}
	dirName_ms.data = C.CString(dirName)
	dirName_ms.len = C.size_t(len(dirName))
	defer C.free(unsafe.Pointer(dirName_ms.data))
	return (bool)(C.QDir_rmdir(this.h, dirName_ms))
}

func (this *QDir) Mkpath(dirPath string) bool {
	dirPath_ms := C.struct_miqt_string{}
	dirPath_ms.data = C.CString(dirPath)
	dirPath_ms.len = C.size_t(len(dirPath))
	defer C.free(unsafe.Pointer(dirPath_ms.data))
	return (bool)(C.QDir_mkpath(this.h, dirPath_ms))
}

func (this *QDir) Rmpath(dirPath string) bool {
	dirPath_ms := C.struct_miqt_string{}
	dirPath_ms.data = C.CString(dirPath)
	dirPath_ms.len = C.size_t(len(dirPath))
	defer C.free(unsafe.Pointer(dirPath_ms.data))
	return (bool)(C.QDir_rmpath(this.h, dirPath_ms))
}

func (this *QDir) RemoveRecursively() bool {
	return (bool)(C.QDir_removeRecursively(this.h))
}

func (this *QDir) IsReadable() bool {
	return (bool)(C.QDir_isReadable(this.h))
}

func (this *QDir) Exists() bool {
	return (bool)(C.QDir_exists(this.h))
}

func (this *QDir) IsRoot() bool {
	return (bool)(C.QDir_isRoot(this.h))
}

func QDir_IsRelativePath(path string) bool {
	path_ms := C.struct_miqt_string{}
	path_ms.data = C.CString(path)
	path_ms.len = C.size_t(len(path))
	defer C.free(unsafe.Pointer(path_ms.data))
	return (bool)(C.QDir_isRelativePath(path_ms))
}

func QDir_IsAbsolutePath(path string) bool {
	path_ms := C.struct_miqt_string{}
	path_ms.data = C.CString(path)
	path_ms.len = C.size_t(len(path))
	defer C.free(unsafe.Pointer(path_ms.data))
	return (bool)(C.QDir_isAbsolutePath(path_ms))
}

func (this *QDir) IsRelative() bool {
	return (bool)(C.QDir_isRelative(this.h))
}

func (this *QDir) IsAbsolute() bool {
	return (bool)(C.QDir_isAbsolute(this.h))
}

func (this *QDir) MakeAbsolute() bool {
	return (bool)(C.QDir_makeAbsolute(this.h))
}

func (this *QDir) OperatorEqual(dir *QDir) bool {
	return (bool)(C.QDir_operatorEqual(this.h, dir.cPointer()))
}

func (this *QDir) OperatorNotEqual(dir *QDir) bool {
	return (bool)(C.QDir_operatorNotEqual(this.h, dir.cPointer()))
}

func (this *QDir) Remove(fileName string) bool {
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	return (bool)(C.QDir_remove(this.h, fileName_ms))
}

func (this *QDir) Rename(oldName string, newName string) bool {
	oldName_ms := C.struct_miqt_string{}
	oldName_ms.data = C.CString(oldName)
	oldName_ms.len = C.size_t(len(oldName))
	defer C.free(unsafe.Pointer(oldName_ms.data))
	newName_ms := C.struct_miqt_string{}
	newName_ms.data = C.CString(newName)
	newName_ms.len = C.size_t(len(newName))
	defer C.free(unsafe.Pointer(newName_ms.data))
	return (bool)(C.QDir_rename(this.h, oldName_ms, newName_ms))
}

func (this *QDir) ExistsWithName(name string) bool {
	name_ms := C.struct_miqt_string{}
	name_ms.data = C.CString(name)
	name_ms.len = C.size_t(len(name))
	defer C.free(unsafe.Pointer(name_ms.data))
	return (bool)(C.QDir_existsWithName(this.h, name_ms))
}

func QDir_Drives() []QFileInfo {
	var _ma C.struct_miqt_array = C.QDir_drives()
	_ret := make([]QFileInfo, int(_ma.len))
	_outCast := (*[0xffff]*C.QFileInfo)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQFileInfo(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QDir_ListSeparator() *QChar {
	_goptr := newQChar(C.QDir_listSeparator())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDir_Separator() *QChar {
	_goptr := newQChar(C.QDir_separator())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDir_SetCurrent(path string) bool {
	path_ms := C.struct_miqt_string{}
	path_ms.data = C.CString(path)
	path_ms.len = C.size_t(len(path))
	defer C.free(unsafe.Pointer(path_ms.data))
	return (bool)(C.QDir_setCurrent(path_ms))
}

func QDir_Current() *QDir {
	_goptr := newQDir(C.QDir_current())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDir_CurrentPath() string {
	var _ms C.struct_miqt_string = C.QDir_currentPath()
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDir_Home() *QDir {
	_goptr := newQDir(C.QDir_home())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDir_HomePath() string {
	var _ms C.struct_miqt_string = C.QDir_homePath()
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDir_Root() *QDir {
	_goptr := newQDir(C.QDir_root())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDir_RootPath() string {
	var _ms C.struct_miqt_string = C.QDir_rootPath()
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDir_Temp() *QDir {
	_goptr := newQDir(C.QDir_temp())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDir_TempPath() string {
	var _ms C.struct_miqt_string = C.QDir_tempPath()
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDir_Match(filters []string, fileName string) bool {
	filters_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(filters))))
	defer C.free(unsafe.Pointer(filters_CArray))
	for i := range filters {
		filters_i_ms := C.struct_miqt_string{}
		filters_i_ms.data = C.CString(filters[i])
		filters_i_ms.len = C.size_t(len(filters[i]))
		defer C.free(unsafe.Pointer(filters_i_ms.data))
		filters_CArray[i] = filters_i_ms
	}
	filters_ma := C.struct_miqt_array{len: C.size_t(len(filters)), data: unsafe.Pointer(filters_CArray)}
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	return (bool)(C.QDir_match(filters_ma, fileName_ms))
}

func QDir_Match2(filter string, fileName string) bool {
	filter_ms := C.struct_miqt_string{}
	filter_ms.data = C.CString(filter)
	filter_ms.len = C.size_t(len(filter))
	defer C.free(unsafe.Pointer(filter_ms.data))
	fileName_ms := C.struct_miqt_string{}
	fileName_ms.data = C.CString(fileName)
	fileName_ms.len = C.size_t(len(fileName))
	defer C.free(unsafe.Pointer(fileName_ms.data))
	return (bool)(C.QDir_match2(filter_ms, fileName_ms))
}

func QDir_CleanPath(path string) string {
	path_ms := C.struct_miqt_string{}
	path_ms.data = C.CString(path)
	path_ms.len = C.size_t(len(path))
	defer C.free(unsafe.Pointer(path_ms.data))
	var _ms C.struct_miqt_string = C.QDir_cleanPath(path_ms)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDir) Refresh() {
	C.QDir_refresh(this.h)
}

func (this *QDir) IsEmptyWithFilters(filters QDir__Filter) bool {
	return (bool)(C.QDir_isEmptyWithFilters(this.h, (C.int)(filters)))
}

func (this *QDir) EntryListWithFilters(filters QDir__Filter) []string {
	var _ma C.struct_miqt_array = C.QDir_entryListWithFilters(this.h, (C.int)(filters))
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

func (this *QDir) EntryList2(filters QDir__Filter, sort QDir__SortFlag) []string {
	var _ma C.struct_miqt_array = C.QDir_entryList2(this.h, (C.int)(filters), (C.int)(sort))
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

func (this *QDir) EntryList3(nameFilters []string, filters QDir__Filter) []string {
	nameFilters_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(nameFilters))))
	defer C.free(unsafe.Pointer(nameFilters_CArray))
	for i := range nameFilters {
		nameFilters_i_ms := C.struct_miqt_string{}
		nameFilters_i_ms.data = C.CString(nameFilters[i])
		nameFilters_i_ms.len = C.size_t(len(nameFilters[i]))
		defer C.free(unsafe.Pointer(nameFilters_i_ms.data))
		nameFilters_CArray[i] = nameFilters_i_ms
	}
	nameFilters_ma := C.struct_miqt_array{len: C.size_t(len(nameFilters)), data: unsafe.Pointer(nameFilters_CArray)}
	var _ma C.struct_miqt_array = C.QDir_entryList3(this.h, nameFilters_ma, (C.int)(filters))
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

func (this *QDir) EntryList4(nameFilters []string, filters QDir__Filter, sort QDir__SortFlag) []string {
	nameFilters_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(nameFilters))))
	defer C.free(unsafe.Pointer(nameFilters_CArray))
	for i := range nameFilters {
		nameFilters_i_ms := C.struct_miqt_string{}
		nameFilters_i_ms.data = C.CString(nameFilters[i])
		nameFilters_i_ms.len = C.size_t(len(nameFilters[i]))
		defer C.free(unsafe.Pointer(nameFilters_i_ms.data))
		nameFilters_CArray[i] = nameFilters_i_ms
	}
	nameFilters_ma := C.struct_miqt_array{len: C.size_t(len(nameFilters)), data: unsafe.Pointer(nameFilters_CArray)}
	var _ma C.struct_miqt_array = C.QDir_entryList4(this.h, nameFilters_ma, (C.int)(filters), (C.int)(sort))
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

func (this *QDir) EntryInfoListWithFilters(filters QDir__Filter) []QFileInfo {
	var _ma C.struct_miqt_array = C.QDir_entryInfoListWithFilters(this.h, (C.int)(filters))
	_ret := make([]QFileInfo, int(_ma.len))
	_outCast := (*[0xffff]*C.QFileInfo)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQFileInfo(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QDir) EntryInfoList2(filters QDir__Filter, sort QDir__SortFlag) []QFileInfo {
	var _ma C.struct_miqt_array = C.QDir_entryInfoList2(this.h, (C.int)(filters), (C.int)(sort))
	_ret := make([]QFileInfo, int(_ma.len))
	_outCast := (*[0xffff]*C.QFileInfo)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQFileInfo(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QDir) EntryInfoList3(nameFilters []string, filters QDir__Filter) []QFileInfo {
	nameFilters_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(nameFilters))))
	defer C.free(unsafe.Pointer(nameFilters_CArray))
	for i := range nameFilters {
		nameFilters_i_ms := C.struct_miqt_string{}
		nameFilters_i_ms.data = C.CString(nameFilters[i])
		nameFilters_i_ms.len = C.size_t(len(nameFilters[i]))
		defer C.free(unsafe.Pointer(nameFilters_i_ms.data))
		nameFilters_CArray[i] = nameFilters_i_ms
	}
	nameFilters_ma := C.struct_miqt_array{len: C.size_t(len(nameFilters)), data: unsafe.Pointer(nameFilters_CArray)}
	var _ma C.struct_miqt_array = C.QDir_entryInfoList3(this.h, nameFilters_ma, (C.int)(filters))
	_ret := make([]QFileInfo, int(_ma.len))
	_outCast := (*[0xffff]*C.QFileInfo)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQFileInfo(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QDir) EntryInfoList4(nameFilters []string, filters QDir__Filter, sort QDir__SortFlag) []QFileInfo {
	nameFilters_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(nameFilters))))
	defer C.free(unsafe.Pointer(nameFilters_CArray))
	for i := range nameFilters {
		nameFilters_i_ms := C.struct_miqt_string{}
		nameFilters_i_ms.data = C.CString(nameFilters[i])
		nameFilters_i_ms.len = C.size_t(len(nameFilters[i]))
		defer C.free(unsafe.Pointer(nameFilters_i_ms.data))
		nameFilters_CArray[i] = nameFilters_i_ms
	}
	nameFilters_ma := C.struct_miqt_array{len: C.size_t(len(nameFilters)), data: unsafe.Pointer(nameFilters_CArray)}
	var _ma C.struct_miqt_array = C.QDir_entryInfoList4(this.h, nameFilters_ma, (C.int)(filters), (C.int)(sort))
	_ret := make([]QFileInfo, int(_ma.len))
	_outCast := (*[0xffff]*C.QFileInfo)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQFileInfo(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

// Delete this object from C++ memory.
func (this *QDir) Delete() {
	C.QDir_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QDir) GoGC() {
	runtime.SetFinalizer(this, func(this *QDir) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
