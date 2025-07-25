#pragma once
#ifndef MIQT_QT6_GEN_QFILEDIALOG_H
#define MIQT_QT6_GEN_QFILEDIALOG_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAbstractFileIconProvider;
class QAbstractItemDelegate;
class QAbstractProxyModel;
class QActionEvent;
class QChildEvent;
class QCloseEvent;
class QContextMenuEvent;
class QDialog;
class QDir;
class QDragEnterEvent;
class QDragLeaveEvent;
class QDragMoveEvent;
class QDropEvent;
class QEnterEvent;
class QEvent;
class QFileDialog;
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
class QShowEvent;
class QSize;
class QTabletEvent;
class QTimerEvent;
class QUrl;
class QVariant;
class QWheelEvent;
class QWidget;
#else
typedef struct QAbstractFileIconProvider QAbstractFileIconProvider;
typedef struct QAbstractItemDelegate QAbstractItemDelegate;
typedef struct QAbstractProxyModel QAbstractProxyModel;
typedef struct QActionEvent QActionEvent;
typedef struct QChildEvent QChildEvent;
typedef struct QCloseEvent QCloseEvent;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QDialog QDialog;
typedef struct QDir QDir;
typedef struct QDragEnterEvent QDragEnterEvent;
typedef struct QDragLeaveEvent QDragLeaveEvent;
typedef struct QDragMoveEvent QDragMoveEvent;
typedef struct QDropEvent QDropEvent;
typedef struct QEnterEvent QEnterEvent;
typedef struct QEvent QEvent;
typedef struct QFileDialog QFileDialog;
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
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QTabletEvent QTabletEvent;
typedef struct QTimerEvent QTimerEvent;
typedef struct QUrl QUrl;
typedef struct QVariant QVariant;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;
#endif

QFileDialog* QFileDialog_new(QWidget* parent);
QFileDialog* QFileDialog_new2(QWidget* parent, int f);
QFileDialog* QFileDialog_new3();
QFileDialog* QFileDialog_new4(QWidget* parent, struct miqt_string caption);
QFileDialog* QFileDialog_new5(QWidget* parent, struct miqt_string caption, struct miqt_string directory);
QFileDialog* QFileDialog_new6(QWidget* parent, struct miqt_string caption, struct miqt_string directory, struct miqt_string filter);
void QFileDialog_virtbase(QFileDialog* src, QDialog** outptr_QDialog);
QMetaObject* QFileDialog_metaObject(const QFileDialog* self);
void* QFileDialog_metacast(QFileDialog* self, const char* param1);
struct miqt_string QFileDialog_tr(const char* s);
void QFileDialog_setDirectory(QFileDialog* self, struct miqt_string directory);
void QFileDialog_setDirectoryWithDirectory(QFileDialog* self, QDir* directory);
QDir* QFileDialog_directory(const QFileDialog* self);
void QFileDialog_setDirectoryUrl(QFileDialog* self, QUrl* directory);
QUrl* QFileDialog_directoryUrl(const QFileDialog* self);
void QFileDialog_selectFile(QFileDialog* self, struct miqt_string filename);
struct miqt_array /* of struct miqt_string */  QFileDialog_selectedFiles(const QFileDialog* self);
void QFileDialog_selectUrl(QFileDialog* self, QUrl* url);
struct miqt_array /* of QUrl* */  QFileDialog_selectedUrls(const QFileDialog* self);
void QFileDialog_setNameFilter(QFileDialog* self, struct miqt_string filter);
void QFileDialog_setNameFilters(QFileDialog* self, struct miqt_array /* of struct miqt_string */  filters);
struct miqt_array /* of struct miqt_string */  QFileDialog_nameFilters(const QFileDialog* self);
void QFileDialog_selectNameFilter(QFileDialog* self, struct miqt_string filter);
struct miqt_string QFileDialog_selectedMimeTypeFilter(const QFileDialog* self);
struct miqt_string QFileDialog_selectedNameFilter(const QFileDialog* self);
void QFileDialog_setMimeTypeFilters(QFileDialog* self, struct miqt_array /* of struct miqt_string */  filters);
struct miqt_array /* of struct miqt_string */  QFileDialog_mimeTypeFilters(const QFileDialog* self);
void QFileDialog_selectMimeTypeFilter(QFileDialog* self, struct miqt_string filter);
int QFileDialog_filter(const QFileDialog* self);
void QFileDialog_setFilter(QFileDialog* self, int filters);
void QFileDialog_setViewMode(QFileDialog* self, int mode);
int QFileDialog_viewMode(const QFileDialog* self);
void QFileDialog_setFileMode(QFileDialog* self, int mode);
int QFileDialog_fileMode(const QFileDialog* self);
void QFileDialog_setAcceptMode(QFileDialog* self, int mode);
int QFileDialog_acceptMode(const QFileDialog* self);
void QFileDialog_setSidebarUrls(QFileDialog* self, struct miqt_array /* of QUrl* */  urls);
struct miqt_array /* of QUrl* */  QFileDialog_sidebarUrls(const QFileDialog* self);
struct miqt_string QFileDialog_saveState(const QFileDialog* self);
bool QFileDialog_restoreState(QFileDialog* self, struct miqt_string state);
void QFileDialog_setDefaultSuffix(QFileDialog* self, struct miqt_string suffix);
struct miqt_string QFileDialog_defaultSuffix(const QFileDialog* self);
void QFileDialog_setHistory(QFileDialog* self, struct miqt_array /* of struct miqt_string */  paths);
struct miqt_array /* of struct miqt_string */  QFileDialog_history(const QFileDialog* self);
void QFileDialog_setItemDelegate(QFileDialog* self, QAbstractItemDelegate* delegate);
QAbstractItemDelegate* QFileDialog_itemDelegate(const QFileDialog* self);
void QFileDialog_setIconProvider(QFileDialog* self, QAbstractFileIconProvider* provider);
QAbstractFileIconProvider* QFileDialog_iconProvider(const QFileDialog* self);
void QFileDialog_setLabelText(QFileDialog* self, int label, struct miqt_string text);
struct miqt_string QFileDialog_labelText(const QFileDialog* self, int label);
void QFileDialog_setSupportedSchemes(QFileDialog* self, struct miqt_array /* of struct miqt_string */  schemes);
struct miqt_array /* of struct miqt_string */  QFileDialog_supportedSchemes(const QFileDialog* self);
void QFileDialog_setProxyModel(QFileDialog* self, QAbstractProxyModel* model);
QAbstractProxyModel* QFileDialog_proxyModel(const QFileDialog* self);
void QFileDialog_setOption(QFileDialog* self, int option);
bool QFileDialog_testOption(const QFileDialog* self, int option);
void QFileDialog_setOptions(QFileDialog* self, int options);
int QFileDialog_options(const QFileDialog* self);
void QFileDialog_setVisible(QFileDialog* self, bool visible);
void QFileDialog_fileSelected(QFileDialog* self, struct miqt_string file);
void QFileDialog_connect_fileSelected(QFileDialog* self, intptr_t slot);
void QFileDialog_filesSelected(QFileDialog* self, struct miqt_array /* of struct miqt_string */  files);
void QFileDialog_connect_filesSelected(QFileDialog* self, intptr_t slot);
void QFileDialog_currentChanged(QFileDialog* self, struct miqt_string path);
void QFileDialog_connect_currentChanged(QFileDialog* self, intptr_t slot);
void QFileDialog_directoryEntered(QFileDialog* self, struct miqt_string directory);
void QFileDialog_connect_directoryEntered(QFileDialog* self, intptr_t slot);
void QFileDialog_urlSelected(QFileDialog* self, QUrl* url);
void QFileDialog_connect_urlSelected(QFileDialog* self, intptr_t slot);
void QFileDialog_urlsSelected(QFileDialog* self, struct miqt_array /* of QUrl* */  urls);
void QFileDialog_connect_urlsSelected(QFileDialog* self, intptr_t slot);
void QFileDialog_currentUrlChanged(QFileDialog* self, QUrl* url);
void QFileDialog_connect_currentUrlChanged(QFileDialog* self, intptr_t slot);
void QFileDialog_directoryUrlEntered(QFileDialog* self, QUrl* directory);
void QFileDialog_connect_directoryUrlEntered(QFileDialog* self, intptr_t slot);
void QFileDialog_filterSelected(QFileDialog* self, struct miqt_string filter);
void QFileDialog_connect_filterSelected(QFileDialog* self, intptr_t slot);
struct miqt_string QFileDialog_getOpenFileName();
QUrl* QFileDialog_getOpenFileUrl();
struct miqt_string QFileDialog_getSaveFileName();
QUrl* QFileDialog_getSaveFileUrl();
struct miqt_string QFileDialog_getExistingDirectory();
QUrl* QFileDialog_getExistingDirectoryUrl();
struct miqt_array /* of struct miqt_string */  QFileDialog_getOpenFileNames();
struct miqt_array /* of QUrl* */  QFileDialog_getOpenFileUrls();
void QFileDialog_saveFileContent(struct miqt_string fileContent, struct miqt_string fileNameHint);
void QFileDialog_done(QFileDialog* self, int result);
void QFileDialog_accept(QFileDialog* self);
void QFileDialog_changeEvent(QFileDialog* self, QEvent* e);
struct miqt_string QFileDialog_tr2(const char* s, const char* c);
struct miqt_string QFileDialog_tr3(const char* s, const char* c, int n);
void QFileDialog_setOption2(QFileDialog* self, int option, bool on);
struct miqt_string QFileDialog_getOpenFileNameWithParent(QWidget* parent);
struct miqt_string QFileDialog_getOpenFileName2(QWidget* parent, struct miqt_string caption);
struct miqt_string QFileDialog_getOpenFileName3(QWidget* parent, struct miqt_string caption, struct miqt_string dir);
struct miqt_string QFileDialog_getOpenFileName4(QWidget* parent, struct miqt_string caption, struct miqt_string dir, struct miqt_string filter);
QUrl* QFileDialog_getOpenFileUrlWithParent(QWidget* parent);
QUrl* QFileDialog_getOpenFileUrl2(QWidget* parent, struct miqt_string caption);
QUrl* QFileDialog_getOpenFileUrl3(QWidget* parent, struct miqt_string caption, QUrl* dir);
QUrl* QFileDialog_getOpenFileUrl4(QWidget* parent, struct miqt_string caption, QUrl* dir, struct miqt_string filter);
struct miqt_string QFileDialog_getSaveFileNameWithParent(QWidget* parent);
struct miqt_string QFileDialog_getSaveFileName2(QWidget* parent, struct miqt_string caption);
struct miqt_string QFileDialog_getSaveFileName3(QWidget* parent, struct miqt_string caption, struct miqt_string dir);
struct miqt_string QFileDialog_getSaveFileName4(QWidget* parent, struct miqt_string caption, struct miqt_string dir, struct miqt_string filter);
QUrl* QFileDialog_getSaveFileUrlWithParent(QWidget* parent);
QUrl* QFileDialog_getSaveFileUrl2(QWidget* parent, struct miqt_string caption);
QUrl* QFileDialog_getSaveFileUrl3(QWidget* parent, struct miqt_string caption, QUrl* dir);
QUrl* QFileDialog_getSaveFileUrl4(QWidget* parent, struct miqt_string caption, QUrl* dir, struct miqt_string filter);
struct miqt_string QFileDialog_getExistingDirectoryWithParent(QWidget* parent);
struct miqt_string QFileDialog_getExistingDirectory2(QWidget* parent, struct miqt_string caption);
struct miqt_string QFileDialog_getExistingDirectory3(QWidget* parent, struct miqt_string caption, struct miqt_string dir);
struct miqt_string QFileDialog_getExistingDirectory4(QWidget* parent, struct miqt_string caption, struct miqt_string dir, int options);
QUrl* QFileDialog_getExistingDirectoryUrlWithParent(QWidget* parent);
QUrl* QFileDialog_getExistingDirectoryUrl2(QWidget* parent, struct miqt_string caption);
QUrl* QFileDialog_getExistingDirectoryUrl3(QWidget* parent, struct miqt_string caption, QUrl* dir);
QUrl* QFileDialog_getExistingDirectoryUrl4(QWidget* parent, struct miqt_string caption, QUrl* dir, int options);
QUrl* QFileDialog_getExistingDirectoryUrl5(QWidget* parent, struct miqt_string caption, QUrl* dir, int options, struct miqt_array /* of struct miqt_string */  supportedSchemes);
struct miqt_array /* of struct miqt_string */  QFileDialog_getOpenFileNamesWithParent(QWidget* parent);
struct miqt_array /* of struct miqt_string */  QFileDialog_getOpenFileNames2(QWidget* parent, struct miqt_string caption);
struct miqt_array /* of struct miqt_string */  QFileDialog_getOpenFileNames3(QWidget* parent, struct miqt_string caption, struct miqt_string dir);
struct miqt_array /* of struct miqt_string */  QFileDialog_getOpenFileNames4(QWidget* parent, struct miqt_string caption, struct miqt_string dir, struct miqt_string filter);
struct miqt_array /* of QUrl* */  QFileDialog_getOpenFileUrlsWithParent(QWidget* parent);
struct miqt_array /* of QUrl* */  QFileDialog_getOpenFileUrls2(QWidget* parent, struct miqt_string caption);
struct miqt_array /* of QUrl* */  QFileDialog_getOpenFileUrls3(QWidget* parent, struct miqt_string caption, QUrl* dir);
struct miqt_array /* of QUrl* */  QFileDialog_getOpenFileUrls4(QWidget* parent, struct miqt_string caption, QUrl* dir, struct miqt_string filter);

bool QFileDialog_override_virtual_setVisible(void* self, intptr_t slot);
void QFileDialog_virtualbase_setVisible(void* self, bool visible);
bool QFileDialog_override_virtual_done(void* self, intptr_t slot);
void QFileDialog_virtualbase_done(void* self, int result);
bool QFileDialog_override_virtual_accept(void* self, intptr_t slot);
void QFileDialog_virtualbase_accept(void* self);
bool QFileDialog_override_virtual_changeEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_changeEvent(void* self, QEvent* e);
bool QFileDialog_override_virtual_sizeHint(void* self, intptr_t slot);
QSize* QFileDialog_virtualbase_sizeHint(const void* self);
bool QFileDialog_override_virtual_minimumSizeHint(void* self, intptr_t slot);
QSize* QFileDialog_virtualbase_minimumSizeHint(const void* self);
bool QFileDialog_override_virtual_open(void* self, intptr_t slot);
void QFileDialog_virtualbase_open(void* self);
bool QFileDialog_override_virtual_exec(void* self, intptr_t slot);
int QFileDialog_virtualbase_exec(void* self);
bool QFileDialog_override_virtual_reject(void* self, intptr_t slot);
void QFileDialog_virtualbase_reject(void* self);
bool QFileDialog_override_virtual_keyPressEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_keyPressEvent(void* self, QKeyEvent* param1);
bool QFileDialog_override_virtual_closeEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_closeEvent(void* self, QCloseEvent* param1);
bool QFileDialog_override_virtual_showEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_showEvent(void* self, QShowEvent* param1);
bool QFileDialog_override_virtual_resizeEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_resizeEvent(void* self, QResizeEvent* param1);
bool QFileDialog_override_virtual_contextMenuEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_contextMenuEvent(void* self, QContextMenuEvent* param1);
bool QFileDialog_override_virtual_eventFilter(void* self, intptr_t slot);
bool QFileDialog_virtualbase_eventFilter(void* self, QObject* param1, QEvent* param2);
bool QFileDialog_override_virtual_devType(void* self, intptr_t slot);
int QFileDialog_virtualbase_devType(const void* self);
bool QFileDialog_override_virtual_heightForWidth(void* self, intptr_t slot);
int QFileDialog_virtualbase_heightForWidth(const void* self, int param1);
bool QFileDialog_override_virtual_hasHeightForWidth(void* self, intptr_t slot);
bool QFileDialog_virtualbase_hasHeightForWidth(const void* self);
bool QFileDialog_override_virtual_paintEngine(void* self, intptr_t slot);
QPaintEngine* QFileDialog_virtualbase_paintEngine(const void* self);
bool QFileDialog_override_virtual_event(void* self, intptr_t slot);
bool QFileDialog_virtualbase_event(void* self, QEvent* event);
bool QFileDialog_override_virtual_mousePressEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_mousePressEvent(void* self, QMouseEvent* event);
bool QFileDialog_override_virtual_mouseReleaseEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_mouseReleaseEvent(void* self, QMouseEvent* event);
bool QFileDialog_override_virtual_mouseDoubleClickEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_mouseDoubleClickEvent(void* self, QMouseEvent* event);
bool QFileDialog_override_virtual_mouseMoveEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_mouseMoveEvent(void* self, QMouseEvent* event);
bool QFileDialog_override_virtual_wheelEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_wheelEvent(void* self, QWheelEvent* event);
bool QFileDialog_override_virtual_keyReleaseEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_keyReleaseEvent(void* self, QKeyEvent* event);
bool QFileDialog_override_virtual_focusInEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_focusInEvent(void* self, QFocusEvent* event);
bool QFileDialog_override_virtual_focusOutEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_focusOutEvent(void* self, QFocusEvent* event);
bool QFileDialog_override_virtual_enterEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_enterEvent(void* self, QEnterEvent* event);
bool QFileDialog_override_virtual_leaveEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_leaveEvent(void* self, QEvent* event);
bool QFileDialog_override_virtual_paintEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_paintEvent(void* self, QPaintEvent* event);
bool QFileDialog_override_virtual_moveEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_moveEvent(void* self, QMoveEvent* event);
bool QFileDialog_override_virtual_tabletEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_tabletEvent(void* self, QTabletEvent* event);
bool QFileDialog_override_virtual_actionEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_actionEvent(void* self, QActionEvent* event);
bool QFileDialog_override_virtual_dragEnterEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_dragEnterEvent(void* self, QDragEnterEvent* event);
bool QFileDialog_override_virtual_dragMoveEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_dragMoveEvent(void* self, QDragMoveEvent* event);
bool QFileDialog_override_virtual_dragLeaveEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_dragLeaveEvent(void* self, QDragLeaveEvent* event);
bool QFileDialog_override_virtual_dropEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_dropEvent(void* self, QDropEvent* event);
bool QFileDialog_override_virtual_hideEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_hideEvent(void* self, QHideEvent* event);
bool QFileDialog_override_virtual_nativeEvent(void* self, intptr_t slot);
bool QFileDialog_virtualbase_nativeEvent(void* self, struct miqt_string eventType, void* message, intptr_t* result);
bool QFileDialog_override_virtual_metric(void* self, intptr_t slot);
int QFileDialog_virtualbase_metric(const void* self, int param1);
bool QFileDialog_override_virtual_initPainter(void* self, intptr_t slot);
void QFileDialog_virtualbase_initPainter(const void* self, QPainter* painter);
bool QFileDialog_override_virtual_redirected(void* self, intptr_t slot);
QPaintDevice* QFileDialog_virtualbase_redirected(const void* self, QPoint* offset);
bool QFileDialog_override_virtual_sharedPainter(void* self, intptr_t slot);
QPainter* QFileDialog_virtualbase_sharedPainter(const void* self);
bool QFileDialog_override_virtual_inputMethodEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_inputMethodEvent(void* self, QInputMethodEvent* param1);
bool QFileDialog_override_virtual_inputMethodQuery(void* self, intptr_t slot);
QVariant* QFileDialog_virtualbase_inputMethodQuery(const void* self, int param1);
bool QFileDialog_override_virtual_focusNextPrevChild(void* self, intptr_t slot);
bool QFileDialog_virtualbase_focusNextPrevChild(void* self, bool next);
bool QFileDialog_override_virtual_timerEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_timerEvent(void* self, QTimerEvent* event);
bool QFileDialog_override_virtual_childEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_childEvent(void* self, QChildEvent* event);
bool QFileDialog_override_virtual_customEvent(void* self, intptr_t slot);
void QFileDialog_virtualbase_customEvent(void* self, QEvent* event);
bool QFileDialog_override_virtual_connectNotify(void* self, intptr_t slot);
void QFileDialog_virtualbase_connectNotify(void* self, QMetaMethod* signal);
bool QFileDialog_override_virtual_disconnectNotify(void* self, intptr_t slot);
void QFileDialog_virtualbase_disconnectNotify(void* self, QMetaMethod* signal);

void QFileDialog_protectedbase_adjustPosition(bool* _dynamic_cast_ok, void* self, QWidget* param1);
void QFileDialog_protectedbase_updateMicroFocus(bool* _dynamic_cast_ok, void* self);
void QFileDialog_protectedbase_create(bool* _dynamic_cast_ok, void* self);
void QFileDialog_protectedbase_destroy(bool* _dynamic_cast_ok, void* self);
bool QFileDialog_protectedbase_focusNextChild(bool* _dynamic_cast_ok, void* self);
bool QFileDialog_protectedbase_focusPreviousChild(bool* _dynamic_cast_ok, void* self);
QObject* QFileDialog_protectedbase_sender(bool* _dynamic_cast_ok, const void* self);
int QFileDialog_protectedbase_senderSignalIndex(bool* _dynamic_cast_ok, const void* self);
int QFileDialog_protectedbase_receivers(bool* _dynamic_cast_ok, const void* self, const char* signal);
bool QFileDialog_protectedbase_isSignalConnected(bool* _dynamic_cast_ok, const void* self, QMetaMethod* signal);

void QFileDialog_delete(QFileDialog* self);

#ifdef __cplusplus
} /* extern C */
#endif

#endif
