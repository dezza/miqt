#pragma once
#ifndef MIQT_QT_GEN_QSCROLLBAR_H
#define MIQT_QT_GEN_QSCROLLBAR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractSlider;
class QActionEvent;
class QChildEvent;
class QCloseEvent;
class QContextMenuEvent;
class QDragEnterEvent;
class QDragLeaveEvent;
class QDragMoveEvent;
class QDropEvent;
class QEvent;
class QFocusEvent;
class QHideEvent;
class QInputMethodEvent;
class QKeyEvent;
class QMetaMethod;
class QMetaObject;
class QMouseEvent;
class QMoveEvent;
class QObject;
class QPaintDevice;
class QPaintEngine;
class QPaintEvent;
class QPainter;
class QPoint;
class QResizeEvent;
class QScrollBar;
class QShowEvent;
class QSize;
class QStyleOptionSlider;
class QTabletEvent;
class QTimerEvent;
class QVariant;
class QWheelEvent;
class QWidget;
#else
typedef struct QAbstractSlider QAbstractSlider;
typedef struct QActionEvent QActionEvent;
typedef struct QChildEvent QChildEvent;
typedef struct QCloseEvent QCloseEvent;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QDragEnterEvent QDragEnterEvent;
typedef struct QDragLeaveEvent QDragLeaveEvent;
typedef struct QDragMoveEvent QDragMoveEvent;
typedef struct QDropEvent QDropEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QHideEvent QHideEvent;
typedef struct QInputMethodEvent QInputMethodEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QMoveEvent QMoveEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEngine QPaintEngine;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPainter QPainter;
typedef struct QPoint QPoint;
typedef struct QResizeEvent QResizeEvent;
typedef struct QScrollBar QScrollBar;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QStyleOptionSlider QStyleOptionSlider;
typedef struct QTabletEvent QTabletEvent;
typedef struct QTimerEvent QTimerEvent;
typedef struct QVariant QVariant;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;
#endif

QScrollBar* QScrollBar_new(QWidget* parent);
QScrollBar* QScrollBar_new2();
QScrollBar* QScrollBar_new3(int param1);
QScrollBar* QScrollBar_new4(int param1, QWidget* parent);
void QScrollBar_virtbase(QScrollBar* src, QAbstractSlider** outptr_QAbstractSlider);
QMetaObject* QScrollBar_metaObject(const QScrollBar* self);
void* QScrollBar_metacast(QScrollBar* self, const char* param1);
struct miqt_string QScrollBar_tr(const char* s);
struct miqt_string QScrollBar_trUtf8(const char* s);
QSize* QScrollBar_sizeHint(const QScrollBar* self);
bool QScrollBar_event(QScrollBar* self, QEvent* event);
void QScrollBar_wheelEvent(QScrollBar* self, QWheelEvent* param1);
void QScrollBar_paintEvent(QScrollBar* self, QPaintEvent* param1);
void QScrollBar_mousePressEvent(QScrollBar* self, QMouseEvent* param1);
void QScrollBar_mouseReleaseEvent(QScrollBar* self, QMouseEvent* param1);
void QScrollBar_mouseMoveEvent(QScrollBar* self, QMouseEvent* param1);
void QScrollBar_hideEvent(QScrollBar* self, QHideEvent* param1);
void QScrollBar_sliderChange(QScrollBar* self, int change);
void QScrollBar_contextMenuEvent(QScrollBar* self, QContextMenuEvent* param1);
struct miqt_string QScrollBar_tr2(const char* s, const char* c);
struct miqt_string QScrollBar_tr3(const char* s, const char* c, int n);
struct miqt_string QScrollBar_trUtf82(const char* s, const char* c);
struct miqt_string QScrollBar_trUtf83(const char* s, const char* c, int n);

bool QScrollBar_override_virtual_sizeHint(void* self, intptr_t slot);
QSize* QScrollBar_virtualbase_sizeHint(const void* self);
bool QScrollBar_override_virtual_event(void* self, intptr_t slot);
bool QScrollBar_virtualbase_event(void* self, QEvent* event);
bool QScrollBar_override_virtual_wheelEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_wheelEvent(void* self, QWheelEvent* param1);
bool QScrollBar_override_virtual_paintEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_paintEvent(void* self, QPaintEvent* param1);
bool QScrollBar_override_virtual_mousePressEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_mousePressEvent(void* self, QMouseEvent* param1);
bool QScrollBar_override_virtual_mouseReleaseEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_mouseReleaseEvent(void* self, QMouseEvent* param1);
bool QScrollBar_override_virtual_mouseMoveEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_mouseMoveEvent(void* self, QMouseEvent* param1);
bool QScrollBar_override_virtual_hideEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_hideEvent(void* self, QHideEvent* param1);
bool QScrollBar_override_virtual_sliderChange(void* self, intptr_t slot);
void QScrollBar_virtualbase_sliderChange(void* self, int change);
bool QScrollBar_override_virtual_contextMenuEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_contextMenuEvent(void* self, QContextMenuEvent* param1);
bool QScrollBar_override_virtual_keyPressEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_keyPressEvent(void* self, QKeyEvent* ev);
bool QScrollBar_override_virtual_timerEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_timerEvent(void* self, QTimerEvent* param1);
bool QScrollBar_override_virtual_changeEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_changeEvent(void* self, QEvent* e);
bool QScrollBar_override_virtual_devType(void* self, intptr_t slot);
int QScrollBar_virtualbase_devType(const void* self);
bool QScrollBar_override_virtual_setVisible(void* self, intptr_t slot);
void QScrollBar_virtualbase_setVisible(void* self, bool visible);
bool QScrollBar_override_virtual_minimumSizeHint(void* self, intptr_t slot);
QSize* QScrollBar_virtualbase_minimumSizeHint(const void* self);
bool QScrollBar_override_virtual_heightForWidth(void* self, intptr_t slot);
int QScrollBar_virtualbase_heightForWidth(const void* self, int param1);
bool QScrollBar_override_virtual_hasHeightForWidth(void* self, intptr_t slot);
bool QScrollBar_virtualbase_hasHeightForWidth(const void* self);
bool QScrollBar_override_virtual_paintEngine(void* self, intptr_t slot);
QPaintEngine* QScrollBar_virtualbase_paintEngine(const void* self);
bool QScrollBar_override_virtual_mouseDoubleClickEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_mouseDoubleClickEvent(void* self, QMouseEvent* event);
bool QScrollBar_override_virtual_keyReleaseEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_keyReleaseEvent(void* self, QKeyEvent* event);
bool QScrollBar_override_virtual_focusInEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_focusInEvent(void* self, QFocusEvent* event);
bool QScrollBar_override_virtual_focusOutEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_focusOutEvent(void* self, QFocusEvent* event);
bool QScrollBar_override_virtual_enterEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_enterEvent(void* self, QEvent* event);
bool QScrollBar_override_virtual_leaveEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_leaveEvent(void* self, QEvent* event);
bool QScrollBar_override_virtual_moveEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_moveEvent(void* self, QMoveEvent* event);
bool QScrollBar_override_virtual_resizeEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_resizeEvent(void* self, QResizeEvent* event);
bool QScrollBar_override_virtual_closeEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_closeEvent(void* self, QCloseEvent* event);
bool QScrollBar_override_virtual_tabletEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_tabletEvent(void* self, QTabletEvent* event);
bool QScrollBar_override_virtual_actionEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_actionEvent(void* self, QActionEvent* event);
bool QScrollBar_override_virtual_dragEnterEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_dragEnterEvent(void* self, QDragEnterEvent* event);
bool QScrollBar_override_virtual_dragMoveEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_dragMoveEvent(void* self, QDragMoveEvent* event);
bool QScrollBar_override_virtual_dragLeaveEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_dragLeaveEvent(void* self, QDragLeaveEvent* event);
bool QScrollBar_override_virtual_dropEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_dropEvent(void* self, QDropEvent* event);
bool QScrollBar_override_virtual_showEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_showEvent(void* self, QShowEvent* event);
bool QScrollBar_override_virtual_nativeEvent(void* self, intptr_t slot);
bool QScrollBar_virtualbase_nativeEvent(void* self, struct miqt_string eventType, void* message, long* result);
bool QScrollBar_override_virtual_metric(void* self, intptr_t slot);
int QScrollBar_virtualbase_metric(const void* self, int param1);
bool QScrollBar_override_virtual_initPainter(void* self, intptr_t slot);
void QScrollBar_virtualbase_initPainter(const void* self, QPainter* painter);
bool QScrollBar_override_virtual_redirected(void* self, intptr_t slot);
QPaintDevice* QScrollBar_virtualbase_redirected(const void* self, QPoint* offset);
bool QScrollBar_override_virtual_sharedPainter(void* self, intptr_t slot);
QPainter* QScrollBar_virtualbase_sharedPainter(const void* self);
bool QScrollBar_override_virtual_inputMethodEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_inputMethodEvent(void* self, QInputMethodEvent* param1);
bool QScrollBar_override_virtual_inputMethodQuery(void* self, intptr_t slot);
QVariant* QScrollBar_virtualbase_inputMethodQuery(const void* self, int param1);
bool QScrollBar_override_virtual_focusNextPrevChild(void* self, intptr_t slot);
bool QScrollBar_virtualbase_focusNextPrevChild(void* self, bool next);
bool QScrollBar_override_virtual_eventFilter(void* self, intptr_t slot);
bool QScrollBar_virtualbase_eventFilter(void* self, QObject* watched, QEvent* event);
bool QScrollBar_override_virtual_childEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_childEvent(void* self, QChildEvent* event);
bool QScrollBar_override_virtual_customEvent(void* self, intptr_t slot);
void QScrollBar_virtualbase_customEvent(void* self, QEvent* event);
bool QScrollBar_override_virtual_connectNotify(void* self, intptr_t slot);
void QScrollBar_virtualbase_connectNotify(void* self, QMetaMethod* signal);
bool QScrollBar_override_virtual_disconnectNotify(void* self, intptr_t slot);
void QScrollBar_virtualbase_disconnectNotify(void* self, QMetaMethod* signal);

void QScrollBar_protectedbase_initStyleOption(bool* _dynamic_cast_ok, const void* self, QStyleOptionSlider* option);
void QScrollBar_protectedbase_setRepeatAction(bool* _dynamic_cast_ok, void* self, int action);
int QScrollBar_protectedbase_repeatAction(bool* _dynamic_cast_ok, const void* self);
void QScrollBar_protectedbase_updateMicroFocus(bool* _dynamic_cast_ok, void* self);
void QScrollBar_protectedbase_create(bool* _dynamic_cast_ok, void* self);
void QScrollBar_protectedbase_destroy(bool* _dynamic_cast_ok, void* self);
bool QScrollBar_protectedbase_focusNextChild(bool* _dynamic_cast_ok, void* self);
bool QScrollBar_protectedbase_focusPreviousChild(bool* _dynamic_cast_ok, void* self);
QObject* QScrollBar_protectedbase_sender(bool* _dynamic_cast_ok, const void* self);
int QScrollBar_protectedbase_senderSignalIndex(bool* _dynamic_cast_ok, const void* self);
int QScrollBar_protectedbase_receivers(bool* _dynamic_cast_ok, const void* self, const char* signal);
bool QScrollBar_protectedbase_isSignalConnected(bool* _dynamic_cast_ok, const void* self, QMetaMethod* signal);

void QScrollBar_delete(QScrollBar* self);

#ifdef __cplusplus
} /* extern C */
#endif

#endif
