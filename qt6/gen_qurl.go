package qt6

/*

#include "gen_qurl.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QUrl__ParsingMode int

const (
	QUrl__TolerantMode QUrl__ParsingMode = 0
	QUrl__StrictMode   QUrl__ParsingMode = 1
	QUrl__DecodedMode  QUrl__ParsingMode = 2
)

type QUrl__UrlFormattingOption uint

const (
	QUrl__None                  QUrl__UrlFormattingOption = 0
	QUrl__RemoveScheme          QUrl__UrlFormattingOption = 1
	QUrl__RemovePassword        QUrl__UrlFormattingOption = 2
	QUrl__RemoveUserInfo        QUrl__UrlFormattingOption = 6
	QUrl__RemovePort            QUrl__UrlFormattingOption = 8
	QUrl__RemoveAuthority       QUrl__UrlFormattingOption = 30
	QUrl__RemovePath            QUrl__UrlFormattingOption = 32
	QUrl__RemoveQuery           QUrl__UrlFormattingOption = 64
	QUrl__RemoveFragment        QUrl__UrlFormattingOption = 128
	QUrl__PreferLocalFile       QUrl__UrlFormattingOption = 512
	QUrl__StripTrailingSlash    QUrl__UrlFormattingOption = 1024
	QUrl__RemoveFilename        QUrl__UrlFormattingOption = 2048
	QUrl__NormalizePathSegments QUrl__UrlFormattingOption = 4096
)

type QUrl__ComponentFormattingOption uint

const (
	QUrl__PrettyDecoded    QUrl__ComponentFormattingOption = 0
	QUrl__EncodeSpaces     QUrl__ComponentFormattingOption = 1048576
	QUrl__EncodeUnicode    QUrl__ComponentFormattingOption = 2097152
	QUrl__EncodeDelimiters QUrl__ComponentFormattingOption = 12582912
	QUrl__EncodeReserved   QUrl__ComponentFormattingOption = 16777216
	QUrl__DecodeReserved   QUrl__ComponentFormattingOption = 33554432
	QUrl__FullyEncoded     QUrl__ComponentFormattingOption = 32505856
	QUrl__FullyDecoded     QUrl__ComponentFormattingOption = 133169152
)

type QUrl__UserInputResolutionOption int

const (
	QUrl__DefaultResolution QUrl__UserInputResolutionOption = 0
	QUrl__AssumeLocalFile   QUrl__UserInputResolutionOption = 1
)

type QUrl__AceProcessingOption uint

const (
	QUrl__IgnoreIDNWhitelist        QUrl__AceProcessingOption = 1
	QUrl__AceTransitionalProcessing QUrl__AceProcessingOption = 2
)

type QUrl struct {
	h *C.QUrl
}

func (this *QUrl) cPointer() *C.QUrl {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QUrl) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQUrl constructs the type using only CGO pointers.
func newQUrl(h *C.QUrl) *QUrl {
	if h == nil {
		return nil
	}

	return &QUrl{h: h}
}

// UnsafeNewQUrl constructs the type using only unsafe pointers.
func UnsafeNewQUrl(h unsafe.Pointer) *QUrl {
	return newQUrl((*C.QUrl)(h))
}

// NewQUrl constructs a new QUrl object.
func NewQUrl() *QUrl {

	return newQUrl(C.QUrl_new())
}

// NewQUrl2 constructs a new QUrl object.
func NewQUrl2(copyVal *QUrl) *QUrl {

	return newQUrl(C.QUrl_new2(copyVal.cPointer()))
}

// NewQUrl3 constructs a new QUrl object.
func NewQUrl3(url string) *QUrl {
	url_ms := C.struct_miqt_string{}
	url_ms.data = C.CString(url)
	url_ms.len = C.size_t(len(url))
	defer C.free(unsafe.Pointer(url_ms.data))

	return newQUrl(C.QUrl_new3(url_ms))
}

// NewQUrl4 constructs a new QUrl object.
func NewQUrl4(url string, mode QUrl__ParsingMode) *QUrl {
	url_ms := C.struct_miqt_string{}
	url_ms.data = C.CString(url)
	url_ms.len = C.size_t(len(url))
	defer C.free(unsafe.Pointer(url_ms.data))

	return newQUrl(C.QUrl_new4(url_ms, (C.int)(mode)))
}

func (this *QUrl) OperatorAssign(copyVal *QUrl) {
	C.QUrl_operatorAssign(this.h, copyVal.cPointer())
}

func (this *QUrl) OperatorAssignWithUrl(url string) {
	url_ms := C.struct_miqt_string{}
	url_ms.data = C.CString(url)
	url_ms.len = C.size_t(len(url))
	defer C.free(unsafe.Pointer(url_ms.data))
	C.QUrl_operatorAssignWithUrl(this.h, url_ms)
}

func (this *QUrl) Swap(other *QUrl) {
	C.QUrl_swap(this.h, other.cPointer())
}

func (this *QUrl) SetUrl(url string) {
	url_ms := C.struct_miqt_string{}
	url_ms.data = C.CString(url)
	url_ms.len = C.size_t(len(url))
	defer C.free(unsafe.Pointer(url_ms.data))
	C.QUrl_setUrl(this.h, url_ms)
}

func (this *QUrl) Url() string {
	var _ms C.struct_miqt_string = C.QUrl_url(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) ToString() string {
	var _ms C.struct_miqt_string = C.QUrl_toString(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) ToDisplayString() string {
	var _ms C.struct_miqt_string = C.QUrl_toDisplayString(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) ToEncoded() []byte {
	var _bytearray C.struct_miqt_string = C.QUrl_toEncoded(this.h)
	_ret := C.GoBytes(unsafe.Pointer(_bytearray.data), C.int(int64(_bytearray.len)))
	C.free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QUrl_FromEncoded(url []byte) *QUrl {
	url_alias := C.struct_miqt_string{}
	if len(url) > 0 {
		url_alias.data = (*C.char)(unsafe.Pointer(&url[0]))
	} else {
		url_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	url_alias.len = C.size_t(len(url))
	_goptr := newQUrl(C.QUrl_fromEncoded(url_alias))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QUrl_FromUserInput(userInput string) *QUrl {
	userInput_ms := C.struct_miqt_string{}
	userInput_ms.data = C.CString(userInput)
	userInput_ms.len = C.size_t(len(userInput))
	defer C.free(unsafe.Pointer(userInput_ms.data))
	_goptr := newQUrl(C.QUrl_fromUserInput(userInput_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUrl) IsValid() bool {
	return (bool)(C.QUrl_isValid(this.h))
}

func (this *QUrl) ErrorString() string {
	var _ms C.struct_miqt_string = C.QUrl_errorString(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) IsEmpty() bool {
	return (bool)(C.QUrl_isEmpty(this.h))
}

func (this *QUrl) Clear() {
	C.QUrl_clear(this.h)
}

func (this *QUrl) SetScheme(scheme string) {
	scheme_ms := C.struct_miqt_string{}
	scheme_ms.data = C.CString(scheme)
	scheme_ms.len = C.size_t(len(scheme))
	defer C.free(unsafe.Pointer(scheme_ms.data))
	C.QUrl_setScheme(this.h, scheme_ms)
}

func (this *QUrl) Scheme() string {
	var _ms C.struct_miqt_string = C.QUrl_scheme(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetAuthority(authority string) {
	authority_ms := C.struct_miqt_string{}
	authority_ms.data = C.CString(authority)
	authority_ms.len = C.size_t(len(authority))
	defer C.free(unsafe.Pointer(authority_ms.data))
	C.QUrl_setAuthority(this.h, authority_ms)
}

func (this *QUrl) Authority() string {
	var _ms C.struct_miqt_string = C.QUrl_authority(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetUserInfo(userInfo string) {
	userInfo_ms := C.struct_miqt_string{}
	userInfo_ms.data = C.CString(userInfo)
	userInfo_ms.len = C.size_t(len(userInfo))
	defer C.free(unsafe.Pointer(userInfo_ms.data))
	C.QUrl_setUserInfo(this.h, userInfo_ms)
}

func (this *QUrl) UserInfo() string {
	var _ms C.struct_miqt_string = C.QUrl_userInfo(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetUserName(userName string) {
	userName_ms := C.struct_miqt_string{}
	userName_ms.data = C.CString(userName)
	userName_ms.len = C.size_t(len(userName))
	defer C.free(unsafe.Pointer(userName_ms.data))
	C.QUrl_setUserName(this.h, userName_ms)
}

func (this *QUrl) UserName() string {
	var _ms C.struct_miqt_string = C.QUrl_userName(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetPassword(password string) {
	password_ms := C.struct_miqt_string{}
	password_ms.data = C.CString(password)
	password_ms.len = C.size_t(len(password))
	defer C.free(unsafe.Pointer(password_ms.data))
	C.QUrl_setPassword(this.h, password_ms)
}

func (this *QUrl) Password() string {
	var _ms C.struct_miqt_string = C.QUrl_password(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetHost(host string) {
	host_ms := C.struct_miqt_string{}
	host_ms.data = C.CString(host)
	host_ms.len = C.size_t(len(host))
	defer C.free(unsafe.Pointer(host_ms.data))
	C.QUrl_setHost(this.h, host_ms)
}

func (this *QUrl) Host() string {
	var _ms C.struct_miqt_string = C.QUrl_host(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetPort(port int) {
	C.QUrl_setPort(this.h, (C.int)(port))
}

func (this *QUrl) Port() int {
	return (int)(C.QUrl_port(this.h))
}

func (this *QUrl) SetPath(path string) {
	path_ms := C.struct_miqt_string{}
	path_ms.data = C.CString(path)
	path_ms.len = C.size_t(len(path))
	defer C.free(unsafe.Pointer(path_ms.data))
	C.QUrl_setPath(this.h, path_ms)
}

func (this *QUrl) Path() string {
	var _ms C.struct_miqt_string = C.QUrl_path(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) FileName() string {
	var _ms C.struct_miqt_string = C.QUrl_fileName(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) HasQuery() bool {
	return (bool)(C.QUrl_hasQuery(this.h))
}

func (this *QUrl) SetQuery(query string) {
	query_ms := C.struct_miqt_string{}
	query_ms.data = C.CString(query)
	query_ms.len = C.size_t(len(query))
	defer C.free(unsafe.Pointer(query_ms.data))
	C.QUrl_setQuery(this.h, query_ms)
}

func (this *QUrl) SetQueryWithQuery(query *QUrlQuery) {
	C.QUrl_setQueryWithQuery(this.h, query.cPointer())
}

func (this *QUrl) Query() string {
	var _ms C.struct_miqt_string = C.QUrl_query(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) HasFragment() bool {
	return (bool)(C.QUrl_hasFragment(this.h))
}

func (this *QUrl) Fragment() string {
	var _ms C.struct_miqt_string = C.QUrl_fragment(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetFragment(fragment string) {
	fragment_ms := C.struct_miqt_string{}
	fragment_ms.data = C.CString(fragment)
	fragment_ms.len = C.size_t(len(fragment))
	defer C.free(unsafe.Pointer(fragment_ms.data))
	C.QUrl_setFragment(this.h, fragment_ms)
}

func (this *QUrl) Resolved(relative *QUrl) *QUrl {
	_goptr := newQUrl(C.QUrl_resolved(this.h, relative.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUrl) IsRelative() bool {
	return (bool)(C.QUrl_isRelative(this.h))
}

func (this *QUrl) IsParentOf(url *QUrl) bool {
	return (bool)(C.QUrl_isParentOf(this.h, url.cPointer()))
}

func (this *QUrl) IsLocalFile() bool {
	return (bool)(C.QUrl_isLocalFile(this.h))
}

func QUrl_FromLocalFile(localfile string) *QUrl {
	localfile_ms := C.struct_miqt_string{}
	localfile_ms.data = C.CString(localfile)
	localfile_ms.len = C.size_t(len(localfile))
	defer C.free(unsafe.Pointer(localfile_ms.data))
	_goptr := newQUrl(C.QUrl_fromLocalFile(localfile_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUrl) ToLocalFile() string {
	var _ms C.struct_miqt_string = C.QUrl_toLocalFile(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) Detach() {
	C.QUrl_detach(this.h)
}

func (this *QUrl) IsDetached() bool {
	return (bool)(C.QUrl_isDetached(this.h))
}

func (this *QUrl) OperatorLesser(url *QUrl) bool {
	return (bool)(C.QUrl_operatorLesser(this.h, url.cPointer()))
}

func (this *QUrl) OperatorEqual(url *QUrl) bool {
	return (bool)(C.QUrl_operatorEqual(this.h, url.cPointer()))
}

func (this *QUrl) OperatorNotEqual(url *QUrl) bool {
	return (bool)(C.QUrl_operatorNotEqual(this.h, url.cPointer()))
}

func QUrl_FromPercentEncoding(param1 []byte) string {
	param1_alias := C.struct_miqt_string{}
	if len(param1) > 0 {
		param1_alias.data = (*C.char)(unsafe.Pointer(&param1[0]))
	} else {
		param1_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	param1_alias.len = C.size_t(len(param1))
	var _ms C.struct_miqt_string = C.QUrl_fromPercentEncoding(param1_alias)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QUrl_ToPercentEncoding(param1 string) []byte {
	param1_ms := C.struct_miqt_string{}
	param1_ms.data = C.CString(param1)
	param1_ms.len = C.size_t(len(param1))
	defer C.free(unsafe.Pointer(param1_ms.data))
	var _bytearray C.struct_miqt_string = C.QUrl_toPercentEncoding(param1_ms)
	_ret := C.GoBytes(unsafe.Pointer(_bytearray.data), C.int(int64(_bytearray.len)))
	C.free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QUrl_FromAce(domain []byte) string {
	domain_alias := C.struct_miqt_string{}
	if len(domain) > 0 {
		domain_alias.data = (*C.char)(unsafe.Pointer(&domain[0]))
	} else {
		domain_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	domain_alias.len = C.size_t(len(domain))
	var _ms C.struct_miqt_string = C.QUrl_fromAce(domain_alias)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QUrl_ToAce(domain string) []byte {
	domain_ms := C.struct_miqt_string{}
	domain_ms.data = C.CString(domain)
	domain_ms.len = C.size_t(len(domain))
	defer C.free(unsafe.Pointer(domain_ms.data))
	var _bytearray C.struct_miqt_string = C.QUrl_toAce(domain_ms)
	_ret := C.GoBytes(unsafe.Pointer(_bytearray.data), C.int(int64(_bytearray.len)))
	C.free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QUrl_IdnWhitelist() []string {
	var _ma C.struct_miqt_array = C.QUrl_idnWhitelist()
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

func QUrl_ToStringList(uris []QUrl) []string {
	uris_CArray := (*[0xffff]*C.QUrl)(C.malloc(C.size_t(8 * len(uris))))
	defer C.free(unsafe.Pointer(uris_CArray))
	for i := range uris {
		uris_CArray[i] = uris[i].cPointer()
	}
	uris_ma := C.struct_miqt_array{len: C.size_t(len(uris)), data: unsafe.Pointer(uris_CArray)}
	var _ma C.struct_miqt_array = C.QUrl_toStringList(uris_ma)
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

func QUrl_FromStringList(uris []string) []QUrl {
	uris_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(uris))))
	defer C.free(unsafe.Pointer(uris_CArray))
	for i := range uris {
		uris_i_ms := C.struct_miqt_string{}
		uris_i_ms.data = C.CString(uris[i])
		uris_i_ms.len = C.size_t(len(uris[i]))
		defer C.free(unsafe.Pointer(uris_i_ms.data))
		uris_CArray[i] = uris_i_ms
	}
	uris_ma := C.struct_miqt_array{len: C.size_t(len(uris)), data: unsafe.Pointer(uris_CArray)}
	var _ma C.struct_miqt_array = C.QUrl_fromStringList(uris_ma)
	_ret := make([]QUrl, int(_ma.len))
	_outCast := (*[0xffff]*C.QUrl)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQUrl(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QUrl_SetIdnWhitelist(idnWhitelist []string) {
	idnWhitelist_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(idnWhitelist))))
	defer C.free(unsafe.Pointer(idnWhitelist_CArray))
	for i := range idnWhitelist {
		idnWhitelist_i_ms := C.struct_miqt_string{}
		idnWhitelist_i_ms.data = C.CString(idnWhitelist[i])
		idnWhitelist_i_ms.len = C.size_t(len(idnWhitelist[i]))
		defer C.free(unsafe.Pointer(idnWhitelist_i_ms.data))
		idnWhitelist_CArray[i] = idnWhitelist_i_ms
	}
	idnWhitelist_ma := C.struct_miqt_array{len: C.size_t(len(idnWhitelist)), data: unsafe.Pointer(idnWhitelist_CArray)}
	C.QUrl_setIdnWhitelist(idnWhitelist_ma)
}

func (this *QUrl) SetUrl2(url string, mode QUrl__ParsingMode) {
	url_ms := C.struct_miqt_string{}
	url_ms.data = C.CString(url)
	url_ms.len = C.size_t(len(url))
	defer C.free(unsafe.Pointer(url_ms.data))
	C.QUrl_setUrl2(this.h, url_ms, (C.int)(mode))
}

func QUrl_FromEncoded2(url []byte, mode QUrl__ParsingMode) *QUrl {
	url_alias := C.struct_miqt_string{}
	if len(url) > 0 {
		url_alias.data = (*C.char)(unsafe.Pointer(&url[0]))
	} else {
		url_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	url_alias.len = C.size_t(len(url))
	_goptr := newQUrl(C.QUrl_fromEncoded2(url_alias, (C.int)(mode)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QUrl_FromUserInput2(userInput string, workingDirectory string) *QUrl {
	userInput_ms := C.struct_miqt_string{}
	userInput_ms.data = C.CString(userInput)
	userInput_ms.len = C.size_t(len(userInput))
	defer C.free(unsafe.Pointer(userInput_ms.data))
	workingDirectory_ms := C.struct_miqt_string{}
	workingDirectory_ms.data = C.CString(workingDirectory)
	workingDirectory_ms.len = C.size_t(len(workingDirectory))
	defer C.free(unsafe.Pointer(workingDirectory_ms.data))
	_goptr := newQUrl(C.QUrl_fromUserInput2(userInput_ms, workingDirectory_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QUrl_FromUserInput3(userInput string, workingDirectory string, options QUrl__UserInputResolutionOption) *QUrl {
	userInput_ms := C.struct_miqt_string{}
	userInput_ms.data = C.CString(userInput)
	userInput_ms.len = C.size_t(len(userInput))
	defer C.free(unsafe.Pointer(userInput_ms.data))
	workingDirectory_ms := C.struct_miqt_string{}
	workingDirectory_ms.data = C.CString(workingDirectory)
	workingDirectory_ms.len = C.size_t(len(workingDirectory))
	defer C.free(unsafe.Pointer(workingDirectory_ms.data))
	_goptr := newQUrl(C.QUrl_fromUserInput3(userInput_ms, workingDirectory_ms, (C.int)(options)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUrl) SetAuthority2(authority string, mode QUrl__ParsingMode) {
	authority_ms := C.struct_miqt_string{}
	authority_ms.data = C.CString(authority)
	authority_ms.len = C.size_t(len(authority))
	defer C.free(unsafe.Pointer(authority_ms.data))
	C.QUrl_setAuthority2(this.h, authority_ms, (C.int)(mode))
}

func (this *QUrl) AuthorityWithOptions(options QUrl__ComponentFormattingOption) string {
	var _ms C.struct_miqt_string = C.QUrl_authorityWithOptions(this.h, (C.uint)(options))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetUserInfo2(userInfo string, mode QUrl__ParsingMode) {
	userInfo_ms := C.struct_miqt_string{}
	userInfo_ms.data = C.CString(userInfo)
	userInfo_ms.len = C.size_t(len(userInfo))
	defer C.free(unsafe.Pointer(userInfo_ms.data))
	C.QUrl_setUserInfo2(this.h, userInfo_ms, (C.int)(mode))
}

func (this *QUrl) UserInfoWithOptions(options QUrl__ComponentFormattingOption) string {
	var _ms C.struct_miqt_string = C.QUrl_userInfoWithOptions(this.h, (C.uint)(options))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetUserName2(userName string, mode QUrl__ParsingMode) {
	userName_ms := C.struct_miqt_string{}
	userName_ms.data = C.CString(userName)
	userName_ms.len = C.size_t(len(userName))
	defer C.free(unsafe.Pointer(userName_ms.data))
	C.QUrl_setUserName2(this.h, userName_ms, (C.int)(mode))
}

func (this *QUrl) UserNameWithOptions(options QUrl__ComponentFormattingOption) string {
	var _ms C.struct_miqt_string = C.QUrl_userNameWithOptions(this.h, (C.uint)(options))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetPassword2(password string, mode QUrl__ParsingMode) {
	password_ms := C.struct_miqt_string{}
	password_ms.data = C.CString(password)
	password_ms.len = C.size_t(len(password))
	defer C.free(unsafe.Pointer(password_ms.data))
	C.QUrl_setPassword2(this.h, password_ms, (C.int)(mode))
}

func (this *QUrl) PasswordWithQUrlComponentFormattingOptions(param1 QUrl__ComponentFormattingOption) string {
	var _ms C.struct_miqt_string = C.QUrl_passwordWithQUrlComponentFormattingOptions(this.h, (C.uint)(param1))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetHost2(host string, mode QUrl__ParsingMode) {
	host_ms := C.struct_miqt_string{}
	host_ms.data = C.CString(host)
	host_ms.len = C.size_t(len(host))
	defer C.free(unsafe.Pointer(host_ms.data))
	C.QUrl_setHost2(this.h, host_ms, (C.int)(mode))
}

func (this *QUrl) HostWithQUrlComponentFormattingOptions(param1 QUrl__ComponentFormattingOption) string {
	var _ms C.struct_miqt_string = C.QUrl_hostWithQUrlComponentFormattingOptions(this.h, (C.uint)(param1))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) PortWithDefaultPort(defaultPort int) int {
	return (int)(C.QUrl_portWithDefaultPort(this.h, (C.int)(defaultPort)))
}

func (this *QUrl) SetPath2(path string, mode QUrl__ParsingMode) {
	path_ms := C.struct_miqt_string{}
	path_ms.data = C.CString(path)
	path_ms.len = C.size_t(len(path))
	defer C.free(unsafe.Pointer(path_ms.data))
	C.QUrl_setPath2(this.h, path_ms, (C.int)(mode))
}

func (this *QUrl) PathWithOptions(options QUrl__ComponentFormattingOption) string {
	var _ms C.struct_miqt_string = C.QUrl_pathWithOptions(this.h, (C.uint)(options))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) FileNameWithOptions(options QUrl__ComponentFormattingOption) string {
	var _ms C.struct_miqt_string = C.QUrl_fileNameWithOptions(this.h, (C.uint)(options))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetQuery2(query string, mode QUrl__ParsingMode) {
	query_ms := C.struct_miqt_string{}
	query_ms.data = C.CString(query)
	query_ms.len = C.size_t(len(query))
	defer C.free(unsafe.Pointer(query_ms.data))
	C.QUrl_setQuery2(this.h, query_ms, (C.int)(mode))
}

func (this *QUrl) QueryWithQUrlComponentFormattingOptions(param1 QUrl__ComponentFormattingOption) string {
	var _ms C.struct_miqt_string = C.QUrl_queryWithQUrlComponentFormattingOptions(this.h, (C.uint)(param1))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) FragmentWithOptions(options QUrl__ComponentFormattingOption) string {
	var _ms C.struct_miqt_string = C.QUrl_fragmentWithOptions(this.h, (C.uint)(options))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUrl) SetFragment2(fragment string, mode QUrl__ParsingMode) {
	fragment_ms := C.struct_miqt_string{}
	fragment_ms.data = C.CString(fragment)
	fragment_ms.len = C.size_t(len(fragment))
	defer C.free(unsafe.Pointer(fragment_ms.data))
	C.QUrl_setFragment2(this.h, fragment_ms, (C.int)(mode))
}

func QUrl_ToPercentEncoding2(param1 string, exclude []byte) []byte {
	param1_ms := C.struct_miqt_string{}
	param1_ms.data = C.CString(param1)
	param1_ms.len = C.size_t(len(param1))
	defer C.free(unsafe.Pointer(param1_ms.data))
	exclude_alias := C.struct_miqt_string{}
	if len(exclude) > 0 {
		exclude_alias.data = (*C.char)(unsafe.Pointer(&exclude[0]))
	} else {
		exclude_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	exclude_alias.len = C.size_t(len(exclude))
	var _bytearray C.struct_miqt_string = C.QUrl_toPercentEncoding2(param1_ms, exclude_alias)
	_ret := C.GoBytes(unsafe.Pointer(_bytearray.data), C.int(int64(_bytearray.len)))
	C.free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QUrl_ToPercentEncoding3(param1 string, exclude []byte, include []byte) []byte {
	param1_ms := C.struct_miqt_string{}
	param1_ms.data = C.CString(param1)
	param1_ms.len = C.size_t(len(param1))
	defer C.free(unsafe.Pointer(param1_ms.data))
	exclude_alias := C.struct_miqt_string{}
	if len(exclude) > 0 {
		exclude_alias.data = (*C.char)(unsafe.Pointer(&exclude[0]))
	} else {
		exclude_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	exclude_alias.len = C.size_t(len(exclude))
	include_alias := C.struct_miqt_string{}
	if len(include) > 0 {
		include_alias.data = (*C.char)(unsafe.Pointer(&include[0]))
	} else {
		include_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	include_alias.len = C.size_t(len(include))
	var _bytearray C.struct_miqt_string = C.QUrl_toPercentEncoding3(param1_ms, exclude_alias, include_alias)
	_ret := C.GoBytes(unsafe.Pointer(_bytearray.data), C.int(int64(_bytearray.len)))
	C.free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QUrl_FromAce2(domain []byte, options QUrl__AceProcessingOption) string {
	domain_alias := C.struct_miqt_string{}
	if len(domain) > 0 {
		domain_alias.data = (*C.char)(unsafe.Pointer(&domain[0]))
	} else {
		domain_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	domain_alias.len = C.size_t(len(domain))
	var _ms C.struct_miqt_string = C.QUrl_fromAce2(domain_alias, (C.uint)(options))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QUrl_ToAce2(domain string, options QUrl__AceProcessingOption) []byte {
	domain_ms := C.struct_miqt_string{}
	domain_ms.data = C.CString(domain)
	domain_ms.len = C.size_t(len(domain))
	defer C.free(unsafe.Pointer(domain_ms.data))
	var _bytearray C.struct_miqt_string = C.QUrl_toAce2(domain_ms, (C.uint)(options))
	_ret := C.GoBytes(unsafe.Pointer(_bytearray.data), C.int(int64(_bytearray.len)))
	C.free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QUrl_FromStringList2(uris []string, mode QUrl__ParsingMode) []QUrl {
	uris_CArray := (*[0xffff]C.struct_miqt_string)(C.malloc(C.size_t(int(unsafe.Sizeof(C.struct_miqt_string{})) * len(uris))))
	defer C.free(unsafe.Pointer(uris_CArray))
	for i := range uris {
		uris_i_ms := C.struct_miqt_string{}
		uris_i_ms.data = C.CString(uris[i])
		uris_i_ms.len = C.size_t(len(uris[i]))
		defer C.free(unsafe.Pointer(uris_i_ms.data))
		uris_CArray[i] = uris_i_ms
	}
	uris_ma := C.struct_miqt_array{len: C.size_t(len(uris)), data: unsafe.Pointer(uris_CArray)}
	var _ma C.struct_miqt_array = C.QUrl_fromStringList2(uris_ma, (C.int)(mode))
	_ret := make([]QUrl, int(_ma.len))
	_outCast := (*[0xffff]*C.QUrl)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQUrl(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

// Delete this object from C++ memory.
func (this *QUrl) Delete() {
	C.QUrl_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QUrl) GoGC() {
	runtime.SetFinalizer(this, func(this *QUrl) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
