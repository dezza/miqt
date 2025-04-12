package qt

/*

#include "gen_qmessagebox.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QMessageBox__Icon int

const (
	QMessageBox__NoIcon      QMessageBox__Icon = 0
	QMessageBox__Information QMessageBox__Icon = 1
	QMessageBox__Warning     QMessageBox__Icon = 2
	QMessageBox__Critical    QMessageBox__Icon = 3
	QMessageBox__Question    QMessageBox__Icon = 4
)

type QMessageBox__ButtonRole int

const (
	QMessageBox__InvalidRole     QMessageBox__ButtonRole = -1
	QMessageBox__AcceptRole      QMessageBox__ButtonRole = 0
	QMessageBox__RejectRole      QMessageBox__ButtonRole = 1
	QMessageBox__DestructiveRole QMessageBox__ButtonRole = 2
	QMessageBox__ActionRole      QMessageBox__ButtonRole = 3
	QMessageBox__HelpRole        QMessageBox__ButtonRole = 4
	QMessageBox__YesRole         QMessageBox__ButtonRole = 5
	QMessageBox__NoRole          QMessageBox__ButtonRole = 6
	QMessageBox__ResetRole       QMessageBox__ButtonRole = 7
	QMessageBox__ApplyRole       QMessageBox__ButtonRole = 8
	QMessageBox__NRoles          QMessageBox__ButtonRole = 9
)

type QMessageBox__StandardButton int

const (
	QMessageBox__NoButton        QMessageBox__StandardButton = 0
	QMessageBox__Ok              QMessageBox__StandardButton = 1024
	QMessageBox__Save            QMessageBox__StandardButton = 2048
	QMessageBox__SaveAll         QMessageBox__StandardButton = 4096
	QMessageBox__Open            QMessageBox__StandardButton = 8192
	QMessageBox__Yes             QMessageBox__StandardButton = 16384
	QMessageBox__YesToAll        QMessageBox__StandardButton = 32768
	QMessageBox__No              QMessageBox__StandardButton = 65536
	QMessageBox__NoToAll         QMessageBox__StandardButton = 131072
	QMessageBox__Abort           QMessageBox__StandardButton = 262144
	QMessageBox__Retry           QMessageBox__StandardButton = 524288
	QMessageBox__Ignore          QMessageBox__StandardButton = 1048576
	QMessageBox__Close           QMessageBox__StandardButton = 2097152
	QMessageBox__Cancel          QMessageBox__StandardButton = 4194304
	QMessageBox__Discard         QMessageBox__StandardButton = 8388608
	QMessageBox__Help            QMessageBox__StandardButton = 16777216
	QMessageBox__Apply           QMessageBox__StandardButton = 33554432
	QMessageBox__Reset           QMessageBox__StandardButton = 67108864
	QMessageBox__RestoreDefaults QMessageBox__StandardButton = 134217728
	QMessageBox__FirstButton     QMessageBox__StandardButton = 1024
	QMessageBox__LastButton      QMessageBox__StandardButton = 134217728
	QMessageBox__YesAll          QMessageBox__StandardButton = 32768
	QMessageBox__NoAll           QMessageBox__StandardButton = 131072
	QMessageBox__Default         QMessageBox__StandardButton = 256
	QMessageBox__Escape          QMessageBox__StandardButton = 512
	QMessageBox__FlagMask        QMessageBox__StandardButton = 768
	QMessageBox__ButtonMask      QMessageBox__StandardButton = -769
)

type QMessageBox struct {
	h *C.QMessageBox
	*QDialog
}

func (this *QMessageBox) cPointer() *C.QMessageBox {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QMessageBox) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQMessageBox constructs the type using only CGO pointers.
func newQMessageBox(h *C.QMessageBox) *QMessageBox {
	if h == nil {
		return nil
	}
	var outptr_QDialog *C.QDialog = nil
	C.QMessageBox_virtbase(h, &outptr_QDialog)

	return &QMessageBox{h: h,
		QDialog: newQDialog(outptr_QDialog)}
}

// UnsafeNewQMessageBox constructs the type using only unsafe pointers.
func UnsafeNewQMessageBox(h unsafe.Pointer) *QMessageBox {
	return newQMessageBox((*C.QMessageBox)(h))
}

// NewQMessageBox constructs a new QMessageBox object.
func NewQMessageBox(parent *QWidget) *QMessageBox {

	return newQMessageBox(C.QMessageBox_new(parent.cPointer()))
}

// NewQMessageBox2 constructs a new QMessageBox object.
func NewQMessageBox2() *QMessageBox {

	return newQMessageBox(C.QMessageBox_new2())
}

// NewQMessageBox3 constructs a new QMessageBox object.
func NewQMessageBox3(icon QMessageBox__Icon, title string, text string) *QMessageBox {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))

	return newQMessageBox(C.QMessageBox_new3((C.int)(icon), title_ms, text_ms))
}

// NewQMessageBox4 constructs a new QMessageBox object.
func NewQMessageBox4(title string, text string, icon QMessageBox__Icon, button0 int, button1 int, button2 int) *QMessageBox {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))

	return newQMessageBox(C.QMessageBox_new4(title_ms, text_ms, (C.int)(icon), (C.int)(button0), (C.int)(button1), (C.int)(button2)))
}

// NewQMessageBox5 constructs a new QMessageBox object.
func NewQMessageBox5(icon QMessageBox__Icon, title string, text string, buttons QMessageBox__StandardButton) *QMessageBox {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))

	return newQMessageBox(C.QMessageBox_new5((C.int)(icon), title_ms, text_ms, (C.int)(buttons)))
}

// NewQMessageBox6 constructs a new QMessageBox object.
func NewQMessageBox6(icon QMessageBox__Icon, title string, text string, buttons QMessageBox__StandardButton, parent *QWidget) *QMessageBox {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))

	return newQMessageBox(C.QMessageBox_new6((C.int)(icon), title_ms, text_ms, (C.int)(buttons), parent.cPointer()))
}

// NewQMessageBox7 constructs a new QMessageBox object.
func NewQMessageBox7(icon QMessageBox__Icon, title string, text string, buttons QMessageBox__StandardButton, parent *QWidget, flags WindowType) *QMessageBox {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))

	return newQMessageBox(C.QMessageBox_new7((C.int)(icon), title_ms, text_ms, (C.int)(buttons), parent.cPointer(), (C.int)(flags)))
}

// NewQMessageBox8 constructs a new QMessageBox object.
func NewQMessageBox8(title string, text string, icon QMessageBox__Icon, button0 int, button1 int, button2 int, parent *QWidget) *QMessageBox {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))

	return newQMessageBox(C.QMessageBox_new8(title_ms, text_ms, (C.int)(icon), (C.int)(button0), (C.int)(button1), (C.int)(button2), parent.cPointer()))
}

// NewQMessageBox9 constructs a new QMessageBox object.
func NewQMessageBox9(title string, text string, icon QMessageBox__Icon, button0 int, button1 int, button2 int, parent *QWidget, f WindowType) *QMessageBox {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))

	return newQMessageBox(C.QMessageBox_new9(title_ms, text_ms, (C.int)(icon), (C.int)(button0), (C.int)(button1), (C.int)(button2), parent.cPointer(), (C.int)(f)))
}

func (this *QMessageBox) MetaObject() *QMetaObject {
	return newQMetaObject(C.QMessageBox_metaObject(this.h))
}

func (this *QMessageBox) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(C.QMessageBox_metacast(this.h, param1_Cstring))
}

func QMessageBox_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QMessageBox_tr(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMessageBox_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms C.struct_miqt_string = C.QMessageBox_trUtf8(s_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMessageBox) AddButton(button *QAbstractButton, role QMessageBox__ButtonRole) {
	C.QMessageBox_addButton(this.h, button.cPointer(), (C.int)(role))
}

func (this *QMessageBox) AddButton2(text string, role QMessageBox__ButtonRole) *QPushButton {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return newQPushButton(C.QMessageBox_addButton2(this.h, text_ms, (C.int)(role)))
}

func (this *QMessageBox) AddButtonWithButton(button QMessageBox__StandardButton) *QPushButton {
	return newQPushButton(C.QMessageBox_addButtonWithButton(this.h, (C.int)(button)))
}

func (this *QMessageBox) RemoveButton(button *QAbstractButton) {
	C.QMessageBox_removeButton(this.h, button.cPointer())
}

func (this *QMessageBox) Buttons() []*QAbstractButton {
	var _ma C.struct_miqt_array = C.QMessageBox_buttons(this.h)
	_ret := make([]*QAbstractButton, int(_ma.len))
	_outCast := (*[0xffff]*C.QAbstractButton)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQAbstractButton(_outCast[i])
	}
	return _ret
}

func (this *QMessageBox) ButtonRole(button *QAbstractButton) QMessageBox__ButtonRole {
	return (QMessageBox__ButtonRole)(C.QMessageBox_buttonRole(this.h, button.cPointer()))
}

func (this *QMessageBox) SetStandardButtons(buttons QMessageBox__StandardButton) {
	C.QMessageBox_setStandardButtons(this.h, (C.int)(buttons))
}

func (this *QMessageBox) StandardButtons() QMessageBox__StandardButton {
	return (QMessageBox__StandardButton)(C.QMessageBox_standardButtons(this.h))
}

func (this *QMessageBox) StandardButton(button *QAbstractButton) QMessageBox__StandardButton {
	return (QMessageBox__StandardButton)(C.QMessageBox_standardButton(this.h, button.cPointer()))
}

func (this *QMessageBox) Button(which QMessageBox__StandardButton) *QAbstractButton {
	return newQAbstractButton(C.QMessageBox_button(this.h, (C.int)(which)))
}

func (this *QMessageBox) DefaultButton() *QPushButton {
	return newQPushButton(C.QMessageBox_defaultButton(this.h))
}

func (this *QMessageBox) SetDefaultButton(button *QPushButton) {
	C.QMessageBox_setDefaultButton(this.h, button.cPointer())
}

func (this *QMessageBox) SetDefaultButtonWithButton(button QMessageBox__StandardButton) {
	C.QMessageBox_setDefaultButtonWithButton(this.h, (C.int)(button))
}

func (this *QMessageBox) EscapeButton() *QAbstractButton {
	return newQAbstractButton(C.QMessageBox_escapeButton(this.h))
}

func (this *QMessageBox) SetEscapeButton(button *QAbstractButton) {
	C.QMessageBox_setEscapeButton(this.h, button.cPointer())
}

func (this *QMessageBox) SetEscapeButtonWithButton(button QMessageBox__StandardButton) {
	C.QMessageBox_setEscapeButtonWithButton(this.h, (C.int)(button))
}

func (this *QMessageBox) ClickedButton() *QAbstractButton {
	return newQAbstractButton(C.QMessageBox_clickedButton(this.h))
}

func (this *QMessageBox) Text() string {
	var _ms C.struct_miqt_string = C.QMessageBox_text(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMessageBox) SetText(text string) {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	C.QMessageBox_setText(this.h, text_ms)
}

func (this *QMessageBox) Icon() QMessageBox__Icon {
	return (QMessageBox__Icon)(C.QMessageBox_icon(this.h))
}

func (this *QMessageBox) SetIcon(icon QMessageBox__Icon) {
	C.QMessageBox_setIcon(this.h, (C.int)(icon))
}

func (this *QMessageBox) IconPixmap() *QPixmap {
	_goptr := newQPixmap(C.QMessageBox_iconPixmap(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMessageBox) SetIconPixmap(pixmap *QPixmap) {
	C.QMessageBox_setIconPixmap(this.h, pixmap.cPointer())
}

func (this *QMessageBox) TextFormat() TextFormat {
	return (TextFormat)(C.QMessageBox_textFormat(this.h))
}

func (this *QMessageBox) SetTextFormat(format TextFormat) {
	C.QMessageBox_setTextFormat(this.h, (C.int)(format))
}

func (this *QMessageBox) SetTextInteractionFlags(flags TextInteractionFlag) {
	C.QMessageBox_setTextInteractionFlags(this.h, (C.int)(flags))
}

func (this *QMessageBox) TextInteractionFlags() TextInteractionFlag {
	return (TextInteractionFlag)(C.QMessageBox_textInteractionFlags(this.h))
}

func (this *QMessageBox) SetCheckBox(cb *QCheckBox) {
	C.QMessageBox_setCheckBox(this.h, cb.cPointer())
}

func (this *QMessageBox) CheckBox() *QCheckBox {
	return newQCheckBox(C.QMessageBox_checkBox(this.h))
}

func QMessageBox_Information(parent *QWidget, title string, text string) QMessageBox__StandardButton {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (QMessageBox__StandardButton)(C.QMessageBox_information(parent.cPointer(), title_ms, text_ms))
}

func QMessageBox_Question(parent *QWidget, title string, text string) QMessageBox__StandardButton {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (QMessageBox__StandardButton)(C.QMessageBox_question(parent.cPointer(), title_ms, text_ms))
}

func QMessageBox_Warning(parent *QWidget, title string, text string) QMessageBox__StandardButton {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (QMessageBox__StandardButton)(C.QMessageBox_warning(parent.cPointer(), title_ms, text_ms))
}

func QMessageBox_Critical(parent *QWidget, title string, text string) QMessageBox__StandardButton {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (QMessageBox__StandardButton)(C.QMessageBox_critical(parent.cPointer(), title_ms, text_ms))
}

func QMessageBox_About(parent *QWidget, title string, text string) {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	C.QMessageBox_about(parent.cPointer(), title_ms, text_ms)
}

func QMessageBox_AboutQt(parent *QWidget) {
	C.QMessageBox_aboutQt(parent.cPointer())
}

func QMessageBox_Information2(parent *QWidget, title string, text string, button0 int) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (int)(C.QMessageBox_information2(parent.cPointer(), title_ms, text_ms, (C.int)(button0)))
}

func QMessageBox_Information3(parent *QWidget, title string, text string, button0Text string) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	button0Text_ms := C.struct_miqt_string{}
	button0Text_ms.data = C.CString(button0Text)
	button0Text_ms.len = C.size_t(len(button0Text))
	defer C.free(unsafe.Pointer(button0Text_ms.data))
	return (int)(C.QMessageBox_information3(parent.cPointer(), title_ms, text_ms, button0Text_ms))
}

func QMessageBox_Information4(parent *QWidget, title string, text string, button0 QMessageBox__StandardButton) QMessageBox__StandardButton {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (QMessageBox__StandardButton)(C.QMessageBox_information4(parent.cPointer(), title_ms, text_ms, (C.int)(button0)))
}

func QMessageBox_Question2(parent *QWidget, title string, text string, button0 int) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (int)(C.QMessageBox_question2(parent.cPointer(), title_ms, text_ms, (C.int)(button0)))
}

func QMessageBox_Question3(parent *QWidget, title string, text string, button0Text string) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	button0Text_ms := C.struct_miqt_string{}
	button0Text_ms.data = C.CString(button0Text)
	button0Text_ms.len = C.size_t(len(button0Text))
	defer C.free(unsafe.Pointer(button0Text_ms.data))
	return (int)(C.QMessageBox_question3(parent.cPointer(), title_ms, text_ms, button0Text_ms))
}

func QMessageBox_Question4(parent *QWidget, title string, text string, button0 QMessageBox__StandardButton, button1 QMessageBox__StandardButton) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (int)(C.QMessageBox_question4(parent.cPointer(), title_ms, text_ms, (C.int)(button0), (C.int)(button1)))
}

func QMessageBox_Warning2(parent *QWidget, title string, text string, button0 int, button1 int) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (int)(C.QMessageBox_warning2(parent.cPointer(), title_ms, text_ms, (C.int)(button0), (C.int)(button1)))
}

func QMessageBox_Warning3(parent *QWidget, title string, text string, button0Text string) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	button0Text_ms := C.struct_miqt_string{}
	button0Text_ms.data = C.CString(button0Text)
	button0Text_ms.len = C.size_t(len(button0Text))
	defer C.free(unsafe.Pointer(button0Text_ms.data))
	return (int)(C.QMessageBox_warning3(parent.cPointer(), title_ms, text_ms, button0Text_ms))
}

func QMessageBox_Warning4(parent *QWidget, title string, text string, button0 QMessageBox__StandardButton, button1 QMessageBox__StandardButton) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (int)(C.QMessageBox_warning4(parent.cPointer(), title_ms, text_ms, (C.int)(button0), (C.int)(button1)))
}

func QMessageBox_Critical2(parent *QWidget, title string, text string, button0 int, button1 int) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (int)(C.QMessageBox_critical2(parent.cPointer(), title_ms, text_ms, (C.int)(button0), (C.int)(button1)))
}

func QMessageBox_Critical3(parent *QWidget, title string, text string, button0Text string) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	button0Text_ms := C.struct_miqt_string{}
	button0Text_ms.data = C.CString(button0Text)
	button0Text_ms.len = C.size_t(len(button0Text))
	defer C.free(unsafe.Pointer(button0Text_ms.data))
	return (int)(C.QMessageBox_critical3(parent.cPointer(), title_ms, text_ms, button0Text_ms))
}

func QMessageBox_Critical4(parent *QWidget, title string, text string, button0 QMessageBox__StandardButton, button1 QMessageBox__StandardButton) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (int)(C.QMessageBox_critical4(parent.cPointer(), title_ms, text_ms, (C.int)(button0), (C.int)(button1)))
}

func (this *QMessageBox) ButtonText(button int) string {
	var _ms C.struct_miqt_string = C.QMessageBox_buttonText(this.h, (C.int)(button))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMessageBox) SetButtonText(button int, text string) {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	C.QMessageBox_setButtonText(this.h, (C.int)(button), text_ms)
}

func (this *QMessageBox) InformativeText() string {
	var _ms C.struct_miqt_string = C.QMessageBox_informativeText(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMessageBox) SetInformativeText(text string) {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	C.QMessageBox_setInformativeText(this.h, text_ms)
}

func (this *QMessageBox) DetailedText() string {
	var _ms C.struct_miqt_string = C.QMessageBox_detailedText(this.h)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMessageBox) SetDetailedText(text string) {
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	C.QMessageBox_setDetailedText(this.h, text_ms)
}

func (this *QMessageBox) SetWindowTitle(title string) {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	C.QMessageBox_setWindowTitle(this.h, title_ms)
}

func (this *QMessageBox) SetWindowModality(windowModality WindowModality) {
	C.QMessageBox_setWindowModality(this.h, (C.int)(windowModality))
}

func QMessageBox_StandardIcon(icon QMessageBox__Icon) *QPixmap {
	_goptr := newQPixmap(C.QMessageBox_standardIcon((C.int)(icon)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMessageBox) ButtonClicked(button *QAbstractButton) {
	C.QMessageBox_buttonClicked(this.h, button.cPointer())
}
func (this *QMessageBox) OnButtonClicked(slot func(button *QAbstractButton)) {
	C.QMessageBox_connect_buttonClicked(this.h, C.intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMessageBox_buttonClicked
func miqt_exec_callback_QMessageBox_buttonClicked(cb C.intptr_t, button *C.QAbstractButton) {
	gofunc, ok := cgo.Handle(cb).Value().(func(button *QAbstractButton))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAbstractButton(button)

	gofunc(slotval1)
}

func QMessageBox_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMessageBox_tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMessageBox_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMessageBox_tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMessageBox_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMessageBox_trUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMessageBox_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms C.struct_miqt_string = C.QMessageBox_trUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMessageBox_Information5(parent *QWidget, title string, text string, buttons QMessageBox__StandardButton) QMessageBox__StandardButton {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (QMessageBox__StandardButton)(C.QMessageBox_information5(parent.cPointer(), title_ms, text_ms, (C.int)(buttons)))
}

func QMessageBox_Information6(parent *QWidget, title string, text string, buttons QMessageBox__StandardButton, defaultButton QMessageBox__StandardButton) QMessageBox__StandardButton {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (QMessageBox__StandardButton)(C.QMessageBox_information6(parent.cPointer(), title_ms, text_ms, (C.int)(buttons), (C.int)(defaultButton)))
}

func QMessageBox_Question5(parent *QWidget, title string, text string, buttons QMessageBox__StandardButton) QMessageBox__StandardButton {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (QMessageBox__StandardButton)(C.QMessageBox_question5(parent.cPointer(), title_ms, text_ms, (C.int)(buttons)))
}

func QMessageBox_Question6(parent *QWidget, title string, text string, buttons QMessageBox__StandardButton, defaultButton QMessageBox__StandardButton) QMessageBox__StandardButton {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (QMessageBox__StandardButton)(C.QMessageBox_question6(parent.cPointer(), title_ms, text_ms, (C.int)(buttons), (C.int)(defaultButton)))
}

func QMessageBox_Warning5(parent *QWidget, title string, text string, buttons QMessageBox__StandardButton) QMessageBox__StandardButton {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (QMessageBox__StandardButton)(C.QMessageBox_warning5(parent.cPointer(), title_ms, text_ms, (C.int)(buttons)))
}

func QMessageBox_Warning6(parent *QWidget, title string, text string, buttons QMessageBox__StandardButton, defaultButton QMessageBox__StandardButton) QMessageBox__StandardButton {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (QMessageBox__StandardButton)(C.QMessageBox_warning6(parent.cPointer(), title_ms, text_ms, (C.int)(buttons), (C.int)(defaultButton)))
}

func QMessageBox_Critical5(parent *QWidget, title string, text string, buttons QMessageBox__StandardButton) QMessageBox__StandardButton {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (QMessageBox__StandardButton)(C.QMessageBox_critical5(parent.cPointer(), title_ms, text_ms, (C.int)(buttons)))
}

func QMessageBox_Critical6(parent *QWidget, title string, text string, buttons QMessageBox__StandardButton, defaultButton QMessageBox__StandardButton) QMessageBox__StandardButton {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (QMessageBox__StandardButton)(C.QMessageBox_critical6(parent.cPointer(), title_ms, text_ms, (C.int)(buttons), (C.int)(defaultButton)))
}

func QMessageBox_AboutQt2(parent *QWidget, title string) {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	C.QMessageBox_aboutQt2(parent.cPointer(), title_ms)
}

func QMessageBox_Information7(parent *QWidget, title string, text string, button0 int, button1 int) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (int)(C.QMessageBox_information7(parent.cPointer(), title_ms, text_ms, (C.int)(button0), (C.int)(button1)))
}

func QMessageBox_Information8(parent *QWidget, title string, text string, button0 int, button1 int, button2 int) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (int)(C.QMessageBox_information8(parent.cPointer(), title_ms, text_ms, (C.int)(button0), (C.int)(button1), (C.int)(button2)))
}

func QMessageBox_Information9(parent *QWidget, title string, text string, button0Text string, button1Text string) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	button0Text_ms := C.struct_miqt_string{}
	button0Text_ms.data = C.CString(button0Text)
	button0Text_ms.len = C.size_t(len(button0Text))
	defer C.free(unsafe.Pointer(button0Text_ms.data))
	button1Text_ms := C.struct_miqt_string{}
	button1Text_ms.data = C.CString(button1Text)
	button1Text_ms.len = C.size_t(len(button1Text))
	defer C.free(unsafe.Pointer(button1Text_ms.data))
	return (int)(C.QMessageBox_information9(parent.cPointer(), title_ms, text_ms, button0Text_ms, button1Text_ms))
}

func QMessageBox_Information10(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	button0Text_ms := C.struct_miqt_string{}
	button0Text_ms.data = C.CString(button0Text)
	button0Text_ms.len = C.size_t(len(button0Text))
	defer C.free(unsafe.Pointer(button0Text_ms.data))
	button1Text_ms := C.struct_miqt_string{}
	button1Text_ms.data = C.CString(button1Text)
	button1Text_ms.len = C.size_t(len(button1Text))
	defer C.free(unsafe.Pointer(button1Text_ms.data))
	button2Text_ms := C.struct_miqt_string{}
	button2Text_ms.data = C.CString(button2Text)
	button2Text_ms.len = C.size_t(len(button2Text))
	defer C.free(unsafe.Pointer(button2Text_ms.data))
	return (int)(C.QMessageBox_information10(parent.cPointer(), title_ms, text_ms, button0Text_ms, button1Text_ms, button2Text_ms))
}

func QMessageBox_Information11(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	button0Text_ms := C.struct_miqt_string{}
	button0Text_ms.data = C.CString(button0Text)
	button0Text_ms.len = C.size_t(len(button0Text))
	defer C.free(unsafe.Pointer(button0Text_ms.data))
	button1Text_ms := C.struct_miqt_string{}
	button1Text_ms.data = C.CString(button1Text)
	button1Text_ms.len = C.size_t(len(button1Text))
	defer C.free(unsafe.Pointer(button1Text_ms.data))
	button2Text_ms := C.struct_miqt_string{}
	button2Text_ms.data = C.CString(button2Text)
	button2Text_ms.len = C.size_t(len(button2Text))
	defer C.free(unsafe.Pointer(button2Text_ms.data))
	return (int)(C.QMessageBox_information11(parent.cPointer(), title_ms, text_ms, button0Text_ms, button1Text_ms, button2Text_ms, (C.int)(defaultButtonNumber)))
}

func QMessageBox_Information12(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int, escapeButtonNumber int) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	button0Text_ms := C.struct_miqt_string{}
	button0Text_ms.data = C.CString(button0Text)
	button0Text_ms.len = C.size_t(len(button0Text))
	defer C.free(unsafe.Pointer(button0Text_ms.data))
	button1Text_ms := C.struct_miqt_string{}
	button1Text_ms.data = C.CString(button1Text)
	button1Text_ms.len = C.size_t(len(button1Text))
	defer C.free(unsafe.Pointer(button1Text_ms.data))
	button2Text_ms := C.struct_miqt_string{}
	button2Text_ms.data = C.CString(button2Text)
	button2Text_ms.len = C.size_t(len(button2Text))
	defer C.free(unsafe.Pointer(button2Text_ms.data))
	return (int)(C.QMessageBox_information12(parent.cPointer(), title_ms, text_ms, button0Text_ms, button1Text_ms, button2Text_ms, (C.int)(defaultButtonNumber), (C.int)(escapeButtonNumber)))
}

func QMessageBox_Information13(parent *QWidget, title string, text string, button0 QMessageBox__StandardButton, button1 QMessageBox__StandardButton) QMessageBox__StandardButton {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (QMessageBox__StandardButton)(C.QMessageBox_information13(parent.cPointer(), title_ms, text_ms, (C.int)(button0), (C.int)(button1)))
}

func QMessageBox_Question7(parent *QWidget, title string, text string, button0 int, button1 int) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (int)(C.QMessageBox_question7(parent.cPointer(), title_ms, text_ms, (C.int)(button0), (C.int)(button1)))
}

func QMessageBox_Question8(parent *QWidget, title string, text string, button0 int, button1 int, button2 int) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (int)(C.QMessageBox_question8(parent.cPointer(), title_ms, text_ms, (C.int)(button0), (C.int)(button1), (C.int)(button2)))
}

func QMessageBox_Question9(parent *QWidget, title string, text string, button0Text string, button1Text string) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	button0Text_ms := C.struct_miqt_string{}
	button0Text_ms.data = C.CString(button0Text)
	button0Text_ms.len = C.size_t(len(button0Text))
	defer C.free(unsafe.Pointer(button0Text_ms.data))
	button1Text_ms := C.struct_miqt_string{}
	button1Text_ms.data = C.CString(button1Text)
	button1Text_ms.len = C.size_t(len(button1Text))
	defer C.free(unsafe.Pointer(button1Text_ms.data))
	return (int)(C.QMessageBox_question9(parent.cPointer(), title_ms, text_ms, button0Text_ms, button1Text_ms))
}

func QMessageBox_Question10(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	button0Text_ms := C.struct_miqt_string{}
	button0Text_ms.data = C.CString(button0Text)
	button0Text_ms.len = C.size_t(len(button0Text))
	defer C.free(unsafe.Pointer(button0Text_ms.data))
	button1Text_ms := C.struct_miqt_string{}
	button1Text_ms.data = C.CString(button1Text)
	button1Text_ms.len = C.size_t(len(button1Text))
	defer C.free(unsafe.Pointer(button1Text_ms.data))
	button2Text_ms := C.struct_miqt_string{}
	button2Text_ms.data = C.CString(button2Text)
	button2Text_ms.len = C.size_t(len(button2Text))
	defer C.free(unsafe.Pointer(button2Text_ms.data))
	return (int)(C.QMessageBox_question10(parent.cPointer(), title_ms, text_ms, button0Text_ms, button1Text_ms, button2Text_ms))
}

func QMessageBox_Question11(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	button0Text_ms := C.struct_miqt_string{}
	button0Text_ms.data = C.CString(button0Text)
	button0Text_ms.len = C.size_t(len(button0Text))
	defer C.free(unsafe.Pointer(button0Text_ms.data))
	button1Text_ms := C.struct_miqt_string{}
	button1Text_ms.data = C.CString(button1Text)
	button1Text_ms.len = C.size_t(len(button1Text))
	defer C.free(unsafe.Pointer(button1Text_ms.data))
	button2Text_ms := C.struct_miqt_string{}
	button2Text_ms.data = C.CString(button2Text)
	button2Text_ms.len = C.size_t(len(button2Text))
	defer C.free(unsafe.Pointer(button2Text_ms.data))
	return (int)(C.QMessageBox_question11(parent.cPointer(), title_ms, text_ms, button0Text_ms, button1Text_ms, button2Text_ms, (C.int)(defaultButtonNumber)))
}

func QMessageBox_Question12(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int, escapeButtonNumber int) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	button0Text_ms := C.struct_miqt_string{}
	button0Text_ms.data = C.CString(button0Text)
	button0Text_ms.len = C.size_t(len(button0Text))
	defer C.free(unsafe.Pointer(button0Text_ms.data))
	button1Text_ms := C.struct_miqt_string{}
	button1Text_ms.data = C.CString(button1Text)
	button1Text_ms.len = C.size_t(len(button1Text))
	defer C.free(unsafe.Pointer(button1Text_ms.data))
	button2Text_ms := C.struct_miqt_string{}
	button2Text_ms.data = C.CString(button2Text)
	button2Text_ms.len = C.size_t(len(button2Text))
	defer C.free(unsafe.Pointer(button2Text_ms.data))
	return (int)(C.QMessageBox_question12(parent.cPointer(), title_ms, text_ms, button0Text_ms, button1Text_ms, button2Text_ms, (C.int)(defaultButtonNumber), (C.int)(escapeButtonNumber)))
}

func QMessageBox_Warning7(parent *QWidget, title string, text string, button0 int, button1 int, button2 int) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (int)(C.QMessageBox_warning7(parent.cPointer(), title_ms, text_ms, (C.int)(button0), (C.int)(button1), (C.int)(button2)))
}

func QMessageBox_Warning8(parent *QWidget, title string, text string, button0Text string, button1Text string) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	button0Text_ms := C.struct_miqt_string{}
	button0Text_ms.data = C.CString(button0Text)
	button0Text_ms.len = C.size_t(len(button0Text))
	defer C.free(unsafe.Pointer(button0Text_ms.data))
	button1Text_ms := C.struct_miqt_string{}
	button1Text_ms.data = C.CString(button1Text)
	button1Text_ms.len = C.size_t(len(button1Text))
	defer C.free(unsafe.Pointer(button1Text_ms.data))
	return (int)(C.QMessageBox_warning8(parent.cPointer(), title_ms, text_ms, button0Text_ms, button1Text_ms))
}

func QMessageBox_Warning9(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	button0Text_ms := C.struct_miqt_string{}
	button0Text_ms.data = C.CString(button0Text)
	button0Text_ms.len = C.size_t(len(button0Text))
	defer C.free(unsafe.Pointer(button0Text_ms.data))
	button1Text_ms := C.struct_miqt_string{}
	button1Text_ms.data = C.CString(button1Text)
	button1Text_ms.len = C.size_t(len(button1Text))
	defer C.free(unsafe.Pointer(button1Text_ms.data))
	button2Text_ms := C.struct_miqt_string{}
	button2Text_ms.data = C.CString(button2Text)
	button2Text_ms.len = C.size_t(len(button2Text))
	defer C.free(unsafe.Pointer(button2Text_ms.data))
	return (int)(C.QMessageBox_warning9(parent.cPointer(), title_ms, text_ms, button0Text_ms, button1Text_ms, button2Text_ms))
}

func QMessageBox_Warning10(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	button0Text_ms := C.struct_miqt_string{}
	button0Text_ms.data = C.CString(button0Text)
	button0Text_ms.len = C.size_t(len(button0Text))
	defer C.free(unsafe.Pointer(button0Text_ms.data))
	button1Text_ms := C.struct_miqt_string{}
	button1Text_ms.data = C.CString(button1Text)
	button1Text_ms.len = C.size_t(len(button1Text))
	defer C.free(unsafe.Pointer(button1Text_ms.data))
	button2Text_ms := C.struct_miqt_string{}
	button2Text_ms.data = C.CString(button2Text)
	button2Text_ms.len = C.size_t(len(button2Text))
	defer C.free(unsafe.Pointer(button2Text_ms.data))
	return (int)(C.QMessageBox_warning10(parent.cPointer(), title_ms, text_ms, button0Text_ms, button1Text_ms, button2Text_ms, (C.int)(defaultButtonNumber)))
}

func QMessageBox_Warning11(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int, escapeButtonNumber int) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	button0Text_ms := C.struct_miqt_string{}
	button0Text_ms.data = C.CString(button0Text)
	button0Text_ms.len = C.size_t(len(button0Text))
	defer C.free(unsafe.Pointer(button0Text_ms.data))
	button1Text_ms := C.struct_miqt_string{}
	button1Text_ms.data = C.CString(button1Text)
	button1Text_ms.len = C.size_t(len(button1Text))
	defer C.free(unsafe.Pointer(button1Text_ms.data))
	button2Text_ms := C.struct_miqt_string{}
	button2Text_ms.data = C.CString(button2Text)
	button2Text_ms.len = C.size_t(len(button2Text))
	defer C.free(unsafe.Pointer(button2Text_ms.data))
	return (int)(C.QMessageBox_warning11(parent.cPointer(), title_ms, text_ms, button0Text_ms, button1Text_ms, button2Text_ms, (C.int)(defaultButtonNumber), (C.int)(escapeButtonNumber)))
}

func QMessageBox_Critical7(parent *QWidget, title string, text string, button0 int, button1 int, button2 int) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	return (int)(C.QMessageBox_critical7(parent.cPointer(), title_ms, text_ms, (C.int)(button0), (C.int)(button1), (C.int)(button2)))
}

func QMessageBox_Critical8(parent *QWidget, title string, text string, button0Text string, button1Text string) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	button0Text_ms := C.struct_miqt_string{}
	button0Text_ms.data = C.CString(button0Text)
	button0Text_ms.len = C.size_t(len(button0Text))
	defer C.free(unsafe.Pointer(button0Text_ms.data))
	button1Text_ms := C.struct_miqt_string{}
	button1Text_ms.data = C.CString(button1Text)
	button1Text_ms.len = C.size_t(len(button1Text))
	defer C.free(unsafe.Pointer(button1Text_ms.data))
	return (int)(C.QMessageBox_critical8(parent.cPointer(), title_ms, text_ms, button0Text_ms, button1Text_ms))
}

func QMessageBox_Critical9(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	button0Text_ms := C.struct_miqt_string{}
	button0Text_ms.data = C.CString(button0Text)
	button0Text_ms.len = C.size_t(len(button0Text))
	defer C.free(unsafe.Pointer(button0Text_ms.data))
	button1Text_ms := C.struct_miqt_string{}
	button1Text_ms.data = C.CString(button1Text)
	button1Text_ms.len = C.size_t(len(button1Text))
	defer C.free(unsafe.Pointer(button1Text_ms.data))
	button2Text_ms := C.struct_miqt_string{}
	button2Text_ms.data = C.CString(button2Text)
	button2Text_ms.len = C.size_t(len(button2Text))
	defer C.free(unsafe.Pointer(button2Text_ms.data))
	return (int)(C.QMessageBox_critical9(parent.cPointer(), title_ms, text_ms, button0Text_ms, button1Text_ms, button2Text_ms))
}

func QMessageBox_Critical10(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	button0Text_ms := C.struct_miqt_string{}
	button0Text_ms.data = C.CString(button0Text)
	button0Text_ms.len = C.size_t(len(button0Text))
	defer C.free(unsafe.Pointer(button0Text_ms.data))
	button1Text_ms := C.struct_miqt_string{}
	button1Text_ms.data = C.CString(button1Text)
	button1Text_ms.len = C.size_t(len(button1Text))
	defer C.free(unsafe.Pointer(button1Text_ms.data))
	button2Text_ms := C.struct_miqt_string{}
	button2Text_ms.data = C.CString(button2Text)
	button2Text_ms.len = C.size_t(len(button2Text))
	defer C.free(unsafe.Pointer(button2Text_ms.data))
	return (int)(C.QMessageBox_critical10(parent.cPointer(), title_ms, text_ms, button0Text_ms, button1Text_ms, button2Text_ms, (C.int)(defaultButtonNumber)))
}

func QMessageBox_Critical11(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int, escapeButtonNumber int) int {
	title_ms := C.struct_miqt_string{}
	title_ms.data = C.CString(title)
	title_ms.len = C.size_t(len(title))
	defer C.free(unsafe.Pointer(title_ms.data))
	text_ms := C.struct_miqt_string{}
	text_ms.data = C.CString(text)
	text_ms.len = C.size_t(len(text))
	defer C.free(unsafe.Pointer(text_ms.data))
	button0Text_ms := C.struct_miqt_string{}
	button0Text_ms.data = C.CString(button0Text)
	button0Text_ms.len = C.size_t(len(button0Text))
	defer C.free(unsafe.Pointer(button0Text_ms.data))
	button1Text_ms := C.struct_miqt_string{}
	button1Text_ms.data = C.CString(button1Text)
	button1Text_ms.len = C.size_t(len(button1Text))
	defer C.free(unsafe.Pointer(button1Text_ms.data))
	button2Text_ms := C.struct_miqt_string{}
	button2Text_ms.data = C.CString(button2Text)
	button2Text_ms.len = C.size_t(len(button2Text))
	defer C.free(unsafe.Pointer(button2Text_ms.data))
	return (int)(C.QMessageBox_critical11(parent.cPointer(), title_ms, text_ms, button0Text_ms, button1Text_ms, button2Text_ms, (C.int)(defaultButtonNumber), (C.int)(escapeButtonNumber)))
}

// AdjustPosition can only be called from a QMessageBox that was directly constructed.
func (this *QMessageBox) AdjustPosition(param1 *QWidget) {

	var _dynamic_cast_ok C.bool = false
	C.QMessageBox_protectedbase_adjustPosition(&_dynamic_cast_ok, unsafe.Pointer(this.h), param1.cPointer())

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

}

// UpdateMicroFocus can only be called from a QMessageBox that was directly constructed.
func (this *QMessageBox) UpdateMicroFocus() {

	var _dynamic_cast_ok C.bool = false
	C.QMessageBox_protectedbase_updateMicroFocus(&_dynamic_cast_ok, unsafe.Pointer(this.h))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

}

// Create can only be called from a QMessageBox that was directly constructed.
func (this *QMessageBox) Create() {

	var _dynamic_cast_ok C.bool = false
	C.QMessageBox_protectedbase_create(&_dynamic_cast_ok, unsafe.Pointer(this.h))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

}

// Destroy can only be called from a QMessageBox that was directly constructed.
func (this *QMessageBox) Destroy() {

	var _dynamic_cast_ok C.bool = false
	C.QMessageBox_protectedbase_destroy(&_dynamic_cast_ok, unsafe.Pointer(this.h))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

}

// FocusNextChild can only be called from a QMessageBox that was directly constructed.
func (this *QMessageBox) FocusNextChild() bool {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QMessageBox_protectedbase_focusNextChild(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// FocusPreviousChild can only be called from a QMessageBox that was directly constructed.
func (this *QMessageBox) FocusPreviousChild() bool {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QMessageBox_protectedbase_focusPreviousChild(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// Sender can only be called from a QMessageBox that was directly constructed.
func (this *QMessageBox) Sender() *QObject {

	var _dynamic_cast_ok C.bool = false
	_method_ret := newQObject(C.QMessageBox_protectedbase_sender(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// SenderSignalIndex can only be called from a QMessageBox that was directly constructed.
func (this *QMessageBox) SenderSignalIndex() int {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QMessageBox_protectedbase_senderSignalIndex(&_dynamic_cast_ok, unsafe.Pointer(this.h)))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// Receivers can only be called from a QMessageBox that was directly constructed.
func (this *QMessageBox) Receivers(signal string) int {
	signal_Cstring := C.CString(signal)
	defer C.free(unsafe.Pointer(signal_Cstring))

	var _dynamic_cast_ok C.bool = false
	_method_ret := (int)(C.QMessageBox_protectedbase_receivers(&_dynamic_cast_ok, unsafe.Pointer(this.h), signal_Cstring))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

// IsSignalConnected can only be called from a QMessageBox that was directly constructed.
func (this *QMessageBox) IsSignalConnected(signal *QMetaMethod) bool {

	var _dynamic_cast_ok C.bool = false
	_method_ret := (bool)(C.QMessageBox_protectedbase_isSignalConnected(&_dynamic_cast_ok, unsafe.Pointer(this.h), signal.cPointer()))

	if !_dynamic_cast_ok {
		panic("miqt: can only call protected methods for directly constructed types")
	}

	return _method_ret

}

func (this *QMessageBox) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(C.QMessageBox_virtualbase_event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QMessageBox) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	ok := C.QMessageBox_override_virtual_event(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_event
func miqt_exec_callback_QMessageBox_event(self *C.QMessageBox, cb C.intptr_t, e *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QMessageBox{h: self}).callVirtualBase_Event, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QMessageBox) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	C.QMessageBox_virtualbase_resizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	ok := C.QMessageBox_override_virtual_resizeEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_resizeEvent
func miqt_exec_callback_QMessageBox_resizeEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_ShowEvent(event *QShowEvent) {

	C.QMessageBox_virtualbase_showEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	ok := C.QMessageBox_override_virtual_showEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_showEvent
func miqt_exec_callback_QMessageBox_showEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_CloseEvent(event *QCloseEvent) {

	C.QMessageBox_virtualbase_closeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	ok := C.QMessageBox_override_virtual_closeEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_closeEvent
func miqt_exec_callback_QMessageBox_closeEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	C.QMessageBox_virtualbase_keyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	ok := C.QMessageBox_override_virtual_keyPressEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_keyPressEvent
func miqt_exec_callback_QMessageBox_keyPressEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_ChangeEvent(event *QEvent) {

	C.QMessageBox_virtualbase_changeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnChangeEvent(slot func(super func(event *QEvent), event *QEvent)) {
	ok := C.QMessageBox_override_virtual_changeEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_changeEvent
func miqt_exec_callback_QMessageBox_changeEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_SetVisible(visible bool) {

	C.QMessageBox_virtualbase_setVisible(unsafe.Pointer(this.h), (C.bool)(visible))

}
func (this *QMessageBox) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	ok := C.QMessageBox_override_virtual_setVisible(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_setVisible
func miqt_exec_callback_QMessageBox_setVisible(self *C.QMessageBox, cb C.intptr_t, visible C.bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QMessageBox{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QMessageBox) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(C.QMessageBox_virtualbase_sizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QMessageBox) OnSizeHint(slot func(super func() *QSize) *QSize) {
	ok := C.QMessageBox_override_virtual_sizeHint(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_sizeHint
func miqt_exec_callback_QMessageBox_sizeHint(self *C.QMessageBox, cb C.intptr_t) *C.QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMessageBox{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QMessageBox) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(C.QMessageBox_virtualbase_minimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QMessageBox) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	ok := C.QMessageBox_override_virtual_minimumSizeHint(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_minimumSizeHint
func miqt_exec_callback_QMessageBox_minimumSizeHint(self *C.QMessageBox, cb C.intptr_t) *C.QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMessageBox{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QMessageBox) callVirtualBase_Open() {

	C.QMessageBox_virtualbase_open(unsafe.Pointer(this.h))

}
func (this *QMessageBox) OnOpen(slot func(super func())) {
	ok := C.QMessageBox_override_virtual_open(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_open
func miqt_exec_callback_QMessageBox_open(self *C.QMessageBox, cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QMessageBox{h: self}).callVirtualBase_Open)

}

func (this *QMessageBox) callVirtualBase_Exec() int {

	return (int)(C.QMessageBox_virtualbase_exec(unsafe.Pointer(this.h)))

}
func (this *QMessageBox) OnExec(slot func(super func() int) int) {
	ok := C.QMessageBox_override_virtual_exec(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_exec
func miqt_exec_callback_QMessageBox_exec(self *C.QMessageBox, cb C.intptr_t) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMessageBox{h: self}).callVirtualBase_Exec)

	return (C.int)(virtualReturn)

}

func (this *QMessageBox) callVirtualBase_Done(param1 int) {

	C.QMessageBox_virtualbase_done(unsafe.Pointer(this.h), (C.int)(param1))

}
func (this *QMessageBox) OnDone(slot func(super func(param1 int), param1 int)) {
	ok := C.QMessageBox_override_virtual_done(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_done
func miqt_exec_callback_QMessageBox_done(self *C.QMessageBox, cb C.intptr_t, param1 C.int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int), param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc((&QMessageBox{h: self}).callVirtualBase_Done, slotval1)

}

func (this *QMessageBox) callVirtualBase_Accept() {

	C.QMessageBox_virtualbase_accept(unsafe.Pointer(this.h))

}
func (this *QMessageBox) OnAccept(slot func(super func())) {
	ok := C.QMessageBox_override_virtual_accept(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_accept
func miqt_exec_callback_QMessageBox_accept(self *C.QMessageBox, cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QMessageBox{h: self}).callVirtualBase_Accept)

}

func (this *QMessageBox) callVirtualBase_Reject() {

	C.QMessageBox_virtualbase_reject(unsafe.Pointer(this.h))

}
func (this *QMessageBox) OnReject(slot func(super func())) {
	ok := C.QMessageBox_override_virtual_reject(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_reject
func miqt_exec_callback_QMessageBox_reject(self *C.QMessageBox, cb C.intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QMessageBox{h: self}).callVirtualBase_Reject)

}

func (this *QMessageBox) callVirtualBase_ContextMenuEvent(param1 *QContextMenuEvent) {

	C.QMessageBox_virtualbase_contextMenuEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QMessageBox) OnContextMenuEvent(slot func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent)) {
	ok := C.QMessageBox_override_virtual_contextMenuEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_contextMenuEvent
func miqt_exec_callback_QMessageBox_contextMenuEvent(self *C.QMessageBox, cb C.intptr_t, param1 *C.QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(param1)

	gofunc((&QMessageBox{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_EventFilter(param1 *QObject, param2 *QEvent) bool {

	return (bool)(C.QMessageBox_virtualbase_eventFilter(unsafe.Pointer(this.h), param1.cPointer(), param2.cPointer()))

}
func (this *QMessageBox) OnEventFilter(slot func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool) {
	ok := C.QMessageBox_override_virtual_eventFilter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_eventFilter
func miqt_exec_callback_QMessageBox_eventFilter(self *C.QMessageBox, cb C.intptr_t, param1 *C.QObject, param2 *C.QEvent) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(param1)

	slotval2 := newQEvent(param2)

	virtualReturn := gofunc((&QMessageBox{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (C.bool)(virtualReturn)

}

func (this *QMessageBox) callVirtualBase_DevType() int {

	return (int)(C.QMessageBox_virtualbase_devType(unsafe.Pointer(this.h)))

}
func (this *QMessageBox) OnDevType(slot func(super func() int) int) {
	ok := C.QMessageBox_override_virtual_devType(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_devType
func miqt_exec_callback_QMessageBox_devType(self *C.QMessageBox, cb C.intptr_t) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMessageBox{h: self}).callVirtualBase_DevType)

	return (C.int)(virtualReturn)

}

func (this *QMessageBox) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(C.QMessageBox_virtualbase_heightForWidth(unsafe.Pointer(this.h), (C.int)(param1)))

}
func (this *QMessageBox) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	ok := C.QMessageBox_override_virtual_heightForWidth(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_heightForWidth
func miqt_exec_callback_QMessageBox_heightForWidth(self *C.QMessageBox, cb C.intptr_t, param1 C.int) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QMessageBox{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (C.int)(virtualReturn)

}

func (this *QMessageBox) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(C.QMessageBox_virtualbase_hasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QMessageBox) OnHasHeightForWidth(slot func(super func() bool) bool) {
	ok := C.QMessageBox_override_virtual_hasHeightForWidth(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_hasHeightForWidth
func miqt_exec_callback_QMessageBox_hasHeightForWidth(self *C.QMessageBox, cb C.intptr_t) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMessageBox{h: self}).callVirtualBase_HasHeightForWidth)

	return (C.bool)(virtualReturn)

}

func (this *QMessageBox) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(C.QMessageBox_virtualbase_paintEngine(unsafe.Pointer(this.h)))

}
func (this *QMessageBox) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	ok := C.QMessageBox_override_virtual_paintEngine(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_paintEngine
func miqt_exec_callback_QMessageBox_paintEngine(self *C.QMessageBox, cb C.intptr_t) *C.QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMessageBox{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QMessageBox) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	C.QMessageBox_virtualbase_mousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	ok := C.QMessageBox_override_virtual_mousePressEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_mousePressEvent
func miqt_exec_callback_QMessageBox_mousePressEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	C.QMessageBox_virtualbase_mouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	ok := C.QMessageBox_override_virtual_mouseReleaseEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_mouseReleaseEvent
func miqt_exec_callback_QMessageBox_mouseReleaseEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	C.QMessageBox_virtualbase_mouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	ok := C.QMessageBox_override_virtual_mouseDoubleClickEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_mouseDoubleClickEvent
func miqt_exec_callback_QMessageBox_mouseDoubleClickEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	C.QMessageBox_virtualbase_mouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	ok := C.QMessageBox_override_virtual_mouseMoveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_mouseMoveEvent
func miqt_exec_callback_QMessageBox_mouseMoveEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_WheelEvent(event *QWheelEvent) {

	C.QMessageBox_virtualbase_wheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	ok := C.QMessageBox_override_virtual_wheelEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_wheelEvent
func miqt_exec_callback_QMessageBox_wheelEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	C.QMessageBox_virtualbase_keyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	ok := C.QMessageBox_override_virtual_keyReleaseEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_keyReleaseEvent
func miqt_exec_callback_QMessageBox_keyReleaseEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	C.QMessageBox_virtualbase_focusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	ok := C.QMessageBox_override_virtual_focusInEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_focusInEvent
func miqt_exec_callback_QMessageBox_focusInEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	C.QMessageBox_virtualbase_focusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	ok := C.QMessageBox_override_virtual_focusOutEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_focusOutEvent
func miqt_exec_callback_QMessageBox_focusOutEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_EnterEvent(event *QEvent) {

	C.QMessageBox_virtualbase_enterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnEnterEvent(slot func(super func(event *QEvent), event *QEvent)) {
	ok := C.QMessageBox_override_virtual_enterEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_enterEvent
func miqt_exec_callback_QMessageBox_enterEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_LeaveEvent(event *QEvent) {

	C.QMessageBox_virtualbase_leaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	ok := C.QMessageBox_override_virtual_leaveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_leaveEvent
func miqt_exec_callback_QMessageBox_leaveEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_PaintEvent(event *QPaintEvent) {

	C.QMessageBox_virtualbase_paintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	ok := C.QMessageBox_override_virtual_paintEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_paintEvent
func miqt_exec_callback_QMessageBox_paintEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_MoveEvent(event *QMoveEvent) {

	C.QMessageBox_virtualbase_moveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	ok := C.QMessageBox_override_virtual_moveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_moveEvent
func miqt_exec_callback_QMessageBox_moveEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_TabletEvent(event *QTabletEvent) {

	C.QMessageBox_virtualbase_tabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	ok := C.QMessageBox_override_virtual_tabletEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_tabletEvent
func miqt_exec_callback_QMessageBox_tabletEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_ActionEvent(event *QActionEvent) {

	C.QMessageBox_virtualbase_actionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	ok := C.QMessageBox_override_virtual_actionEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_actionEvent
func miqt_exec_callback_QMessageBox_actionEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	C.QMessageBox_virtualbase_dragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	ok := C.QMessageBox_override_virtual_dragEnterEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_dragEnterEvent
func miqt_exec_callback_QMessageBox_dragEnterEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	C.QMessageBox_virtualbase_dragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	ok := C.QMessageBox_override_virtual_dragMoveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_dragMoveEvent
func miqt_exec_callback_QMessageBox_dragMoveEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	C.QMessageBox_virtualbase_dragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	ok := C.QMessageBox_override_virtual_dragLeaveEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_dragLeaveEvent
func miqt_exec_callback_QMessageBox_dragLeaveEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_DropEvent(event *QDropEvent) {

	C.QMessageBox_virtualbase_dropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	ok := C.QMessageBox_override_virtual_dropEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_dropEvent
func miqt_exec_callback_QMessageBox_dropEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_HideEvent(event *QHideEvent) {

	C.QMessageBox_virtualbase_hideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	ok := C.QMessageBox_override_virtual_hideEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_hideEvent
func miqt_exec_callback_QMessageBox_hideEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *int64) bool {
	eventType_alias := C.struct_miqt_string{}
	if len(eventType) > 0 {
		eventType_alias.data = (*C.char)(unsafe.Pointer(&eventType[0]))
	} else {
		eventType_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	eventType_alias.len = C.size_t(len(eventType))

	return (bool)(C.QMessageBox_virtualbase_nativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*C.long)(unsafe.Pointer(result))))

}
func (this *QMessageBox) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *int64) bool, eventType []byte, message unsafe.Pointer, result *int64) bool) {
	ok := C.QMessageBox_override_virtual_nativeEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_nativeEvent
func miqt_exec_callback_QMessageBox_nativeEvent(self *C.QMessageBox, cb C.intptr_t, eventType C.struct_miqt_string, message unsafe.Pointer, result *C.long) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(eventType []byte, message unsafe.Pointer, result *int64) bool, eventType []byte, message unsafe.Pointer, result *int64) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var eventType_bytearray C.struct_miqt_string = eventType
	eventType_ret := C.GoBytes(unsafe.Pointer(eventType_bytearray.data), C.int(int64(eventType_bytearray.len)))
	C.free(unsafe.Pointer(eventType_bytearray.data))
	slotval1 := eventType_ret
	slotval2 := (unsafe.Pointer)(message)

	slotval3 := (*int64)(unsafe.Pointer(result))

	virtualReturn := gofunc((&QMessageBox{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (C.bool)(virtualReturn)

}

func (this *QMessageBox) callVirtualBase_Metric(param1 QPaintDevice__PaintDeviceMetric) int {

	return (int)(C.QMessageBox_virtualbase_metric(unsafe.Pointer(this.h), (C.int)(param1)))

}
func (this *QMessageBox) OnMetric(slot func(super func(param1 QPaintDevice__PaintDeviceMetric) int, param1 QPaintDevice__PaintDeviceMetric) int) {
	ok := C.QMessageBox_override_virtual_metric(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_metric
func miqt_exec_callback_QMessageBox_metric(self *C.QMessageBox, cb C.intptr_t, param1 C.int) C.int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 QPaintDevice__PaintDeviceMetric) int, param1 QPaintDevice__PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QPaintDevice__PaintDeviceMetric)(param1)

	virtualReturn := gofunc((&QMessageBox{h: self}).callVirtualBase_Metric, slotval1)

	return (C.int)(virtualReturn)

}

func (this *QMessageBox) callVirtualBase_InitPainter(painter *QPainter) {

	C.QMessageBox_virtualbase_initPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QMessageBox) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	ok := C.QMessageBox_override_virtual_initPainter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_initPainter
func miqt_exec_callback_QMessageBox_initPainter(self *C.QMessageBox, cb C.intptr_t, painter *C.QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QMessageBox{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QMessageBox) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(C.QMessageBox_virtualbase_redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QMessageBox) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	ok := C.QMessageBox_override_virtual_redirected(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_redirected
func miqt_exec_callback_QMessageBox_redirected(self *C.QMessageBox, cb C.intptr_t, offset *C.QPoint) *C.QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QMessageBox{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QMessageBox) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(C.QMessageBox_virtualbase_sharedPainter(unsafe.Pointer(this.h)))

}
func (this *QMessageBox) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	ok := C.QMessageBox_override_virtual_sharedPainter(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_sharedPainter
func miqt_exec_callback_QMessageBox_sharedPainter(self *C.QMessageBox, cb C.intptr_t) *C.QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMessageBox{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QMessageBox) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	C.QMessageBox_virtualbase_inputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QMessageBox) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	ok := C.QMessageBox_override_virtual_inputMethodEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_inputMethodEvent
func miqt_exec_callback_QMessageBox_inputMethodEvent(self *C.QMessageBox, cb C.intptr_t, param1 *C.QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QMessageBox{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(C.QMessageBox_virtualbase_inputMethodQuery(unsafe.Pointer(this.h), (C.int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QMessageBox) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	ok := C.QMessageBox_override_virtual_inputMethodQuery(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_inputMethodQuery
func miqt_exec_callback_QMessageBox_inputMethodQuery(self *C.QMessageBox, cb C.intptr_t, param1 C.int) *C.QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QMessageBox{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QMessageBox) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(C.QMessageBox_virtualbase_focusNextPrevChild(unsafe.Pointer(this.h), (C.bool)(next)))

}
func (this *QMessageBox) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	ok := C.QMessageBox_override_virtual_focusNextPrevChild(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_focusNextPrevChild
func miqt_exec_callback_QMessageBox_focusNextPrevChild(self *C.QMessageBox, cb C.intptr_t, next C.bool) C.bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QMessageBox{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (C.bool)(virtualReturn)

}

func (this *QMessageBox) callVirtualBase_TimerEvent(event *QTimerEvent) {

	C.QMessageBox_virtualbase_timerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	ok := C.QMessageBox_override_virtual_timerEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_timerEvent
func miqt_exec_callback_QMessageBox_timerEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_ChildEvent(event *QChildEvent) {

	C.QMessageBox_virtualbase_childEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	ok := C.QMessageBox_override_virtual_childEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_childEvent
func miqt_exec_callback_QMessageBox_childEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_CustomEvent(event *QEvent) {

	C.QMessageBox_virtualbase_customEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QMessageBox) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	ok := C.QMessageBox_override_virtual_customEvent(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_customEvent
func miqt_exec_callback_QMessageBox_customEvent(self *C.QMessageBox, cb C.intptr_t, event *C.QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QMessageBox{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QMessageBox) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	C.QMessageBox_virtualbase_connectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QMessageBox) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	ok := C.QMessageBox_override_virtual_connectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_connectNotify
func miqt_exec_callback_QMessageBox_connectNotify(self *C.QMessageBox, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QMessageBox{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QMessageBox) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	C.QMessageBox_virtualbase_disconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QMessageBox) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	ok := C.QMessageBox_override_virtual_disconnectNotify(unsafe.Pointer(this.h), C.intptr_t(cgo.NewHandle(slot)))
	if !ok {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
}

//export miqt_exec_callback_QMessageBox_disconnectNotify
func miqt_exec_callback_QMessageBox_disconnectNotify(self *C.QMessageBox, cb C.intptr_t, signal *C.QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QMessageBox{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

// Delete this object from C++ memory.
func (this *QMessageBox) Delete() {
	C.QMessageBox_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QMessageBox) GoGC() {
	runtime.SetFinalizer(this, func(this *QMessageBox) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
