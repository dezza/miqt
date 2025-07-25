#include <QAbstractAxis>
#include <QChildEvent>
#include <QEvent>
#include <QLogValueAxis>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTimerEvent>
#include <qlogvalueaxis.h>
#include "gen_qlogvalueaxis.h"

#ifdef __cplusplus
extern "C" {
#endif

void miqt_exec_callback_QLogValueAxis_minChanged(intptr_t, double);
void miqt_exec_callback_QLogValueAxis_maxChanged(intptr_t, double);
void miqt_exec_callback_QLogValueAxis_rangeChanged(intptr_t, double, double);
void miqt_exec_callback_QLogValueAxis_labelFormatChanged(intptr_t, struct miqt_string);
void miqt_exec_callback_QLogValueAxis_baseChanged(intptr_t, double);
void miqt_exec_callback_QLogValueAxis_tickCountChanged(intptr_t, int);
void miqt_exec_callback_QLogValueAxis_minorTickCountChanged(intptr_t, int);
int miqt_exec_callback_QLogValueAxis_type(const QLogValueAxis*, intptr_t);
bool miqt_exec_callback_QLogValueAxis_event(QLogValueAxis*, intptr_t, QEvent*);
bool miqt_exec_callback_QLogValueAxis_eventFilter(QLogValueAxis*, intptr_t, QObject*, QEvent*);
void miqt_exec_callback_QLogValueAxis_timerEvent(QLogValueAxis*, intptr_t, QTimerEvent*);
void miqt_exec_callback_QLogValueAxis_childEvent(QLogValueAxis*, intptr_t, QChildEvent*);
void miqt_exec_callback_QLogValueAxis_customEvent(QLogValueAxis*, intptr_t, QEvent*);
void miqt_exec_callback_QLogValueAxis_connectNotify(QLogValueAxis*, intptr_t, QMetaMethod*);
void miqt_exec_callback_QLogValueAxis_disconnectNotify(QLogValueAxis*, intptr_t, QMetaMethod*);
#ifdef __cplusplus
} /* extern C */
#endif

class MiqtVirtualQLogValueAxis final : public QLogValueAxis {
public:

	MiqtVirtualQLogValueAxis(): QLogValueAxis() {}
	MiqtVirtualQLogValueAxis(QObject* parent): QLogValueAxis(parent) {}

	virtual ~MiqtVirtualQLogValueAxis() override = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__type = 0;

	// Subclass to allow providing a Go implementation
	virtual QAbstractAxis::AxisType type() const override {
		if (handle__type == 0) {
			return QLogValueAxis::type();
		}

		int callback_return_value = miqt_exec_callback_QLogValueAxis_type(this, handle__type);
		return static_cast<QAbstractAxis::AxisType>(callback_return_value);
	}

	friend int QLogValueAxis_virtualbase_type(const void* self);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__event = 0;

	// Subclass to allow providing a Go implementation
	virtual bool event(QEvent* event) override {
		if (handle__event == 0) {
			return QLogValueAxis::event(event);
		}

		QEvent* sigval1 = event;
		bool callback_return_value = miqt_exec_callback_QLogValueAxis_event(this, handle__event, sigval1);
		return callback_return_value;
	}

	friend bool QLogValueAxis_virtualbase_event(void* self, QEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__eventFilter = 0;

	// Subclass to allow providing a Go implementation
	virtual bool eventFilter(QObject* watched, QEvent* event) override {
		if (handle__eventFilter == 0) {
			return QLogValueAxis::eventFilter(watched, event);
		}

		QObject* sigval1 = watched;
		QEvent* sigval2 = event;
		bool callback_return_value = miqt_exec_callback_QLogValueAxis_eventFilter(this, handle__eventFilter, sigval1, sigval2);
		return callback_return_value;
	}

	friend bool QLogValueAxis_virtualbase_eventFilter(void* self, QObject* watched, QEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__timerEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void timerEvent(QTimerEvent* event) override {
		if (handle__timerEvent == 0) {
			QLogValueAxis::timerEvent(event);
			return;
		}

		QTimerEvent* sigval1 = event;
		miqt_exec_callback_QLogValueAxis_timerEvent(this, handle__timerEvent, sigval1);

	}

	friend void QLogValueAxis_virtualbase_timerEvent(void* self, QTimerEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__childEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void childEvent(QChildEvent* event) override {
		if (handle__childEvent == 0) {
			QLogValueAxis::childEvent(event);
			return;
		}

		QChildEvent* sigval1 = event;
		miqt_exec_callback_QLogValueAxis_childEvent(this, handle__childEvent, sigval1);

	}

	friend void QLogValueAxis_virtualbase_childEvent(void* self, QChildEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__customEvent = 0;

	// Subclass to allow providing a Go implementation
	virtual void customEvent(QEvent* event) override {
		if (handle__customEvent == 0) {
			QLogValueAxis::customEvent(event);
			return;
		}

		QEvent* sigval1 = event;
		miqt_exec_callback_QLogValueAxis_customEvent(this, handle__customEvent, sigval1);

	}

	friend void QLogValueAxis_virtualbase_customEvent(void* self, QEvent* event);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__connectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void connectNotify(const QMetaMethod& signal) override {
		if (handle__connectNotify == 0) {
			QLogValueAxis::connectNotify(signal);
			return;
		}

		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);
		miqt_exec_callback_QLogValueAxis_connectNotify(this, handle__connectNotify, sigval1);

	}

	friend void QLogValueAxis_virtualbase_connectNotify(void* self, QMetaMethod* signal);

	// cgo.Handle value for overwritten implementation
	intptr_t handle__disconnectNotify = 0;

	// Subclass to allow providing a Go implementation
	virtual void disconnectNotify(const QMetaMethod& signal) override {
		if (handle__disconnectNotify == 0) {
			QLogValueAxis::disconnectNotify(signal);
			return;
		}

		const QMetaMethod& signal_ret = signal;
		// Cast returned reference into pointer
		QMetaMethod* sigval1 = const_cast<QMetaMethod*>(&signal_ret);
		miqt_exec_callback_QLogValueAxis_disconnectNotify(this, handle__disconnectNotify, sigval1);

	}

	friend void QLogValueAxis_virtualbase_disconnectNotify(void* self, QMetaMethod* signal);

	// Wrappers to allow calling protected methods:
	friend QObject* QLogValueAxis_protectedbase_sender(bool* _dynamic_cast_ok, const void* self);
	friend int QLogValueAxis_protectedbase_senderSignalIndex(bool* _dynamic_cast_ok, const void* self);
	friend int QLogValueAxis_protectedbase_receivers(bool* _dynamic_cast_ok, const void* self, const char* signal);
	friend bool QLogValueAxis_protectedbase_isSignalConnected(bool* _dynamic_cast_ok, const void* self, QMetaMethod* signal);
};

QLogValueAxis* QLogValueAxis_new() {
	return new (std::nothrow) MiqtVirtualQLogValueAxis();
}

QLogValueAxis* QLogValueAxis_new2(QObject* parent) {
	return new (std::nothrow) MiqtVirtualQLogValueAxis(parent);
}

void QLogValueAxis_virtbase(QLogValueAxis* src, QAbstractAxis** outptr_QAbstractAxis) {
	*outptr_QAbstractAxis = static_cast<QAbstractAxis*>(src);
}

QMetaObject* QLogValueAxis_metaObject(const QLogValueAxis* self) {
	return (QMetaObject*) self->metaObject();
}

void* QLogValueAxis_metacast(QLogValueAxis* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QLogValueAxis_tr(const char* s) {
	QString _ret = QLogValueAxis::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QLogValueAxis_type(const QLogValueAxis* self) {
	QAbstractAxis::AxisType _ret = self->type();
	return static_cast<int>(_ret);
}

void QLogValueAxis_setMin(QLogValueAxis* self, double min) {
	self->setMin(static_cast<qreal>(min));
}

double QLogValueAxis_min(const QLogValueAxis* self) {
	qreal _ret = self->min();
	return static_cast<double>(_ret);
}

void QLogValueAxis_setMax(QLogValueAxis* self, double max) {
	self->setMax(static_cast<qreal>(max));
}

double QLogValueAxis_max(const QLogValueAxis* self) {
	qreal _ret = self->max();
	return static_cast<double>(_ret);
}

void QLogValueAxis_setRange(QLogValueAxis* self, double min, double max) {
	self->setRange(static_cast<qreal>(min), static_cast<qreal>(max));
}

void QLogValueAxis_setLabelFormat(QLogValueAxis* self, struct miqt_string format) {
	QString format_QString = QString::fromUtf8(format.data, format.len);
	self->setLabelFormat(format_QString);
}

struct miqt_string QLogValueAxis_labelFormat(const QLogValueAxis* self) {
	QString _ret = self->labelFormat();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QLogValueAxis_setBase(QLogValueAxis* self, double base) {
	self->setBase(static_cast<qreal>(base));
}

double QLogValueAxis_base(const QLogValueAxis* self) {
	qreal _ret = self->base();
	return static_cast<double>(_ret);
}

int QLogValueAxis_tickCount(const QLogValueAxis* self) {
	return self->tickCount();
}

void QLogValueAxis_setMinorTickCount(QLogValueAxis* self, int minorTickCount) {
	self->setMinorTickCount(static_cast<int>(minorTickCount));
}

int QLogValueAxis_minorTickCount(const QLogValueAxis* self) {
	return self->minorTickCount();
}

void QLogValueAxis_minChanged(QLogValueAxis* self, double min) {
	self->minChanged(static_cast<qreal>(min));
}

void QLogValueAxis_connect_minChanged(QLogValueAxis* self, intptr_t slot) {
	QLogValueAxis::connect(self, static_cast<void (QLogValueAxis::*)(qreal)>(&QLogValueAxis::minChanged), self, [=](qreal min) {
		qreal min_ret = min;
		double sigval1 = static_cast<double>(min_ret);
		miqt_exec_callback_QLogValueAxis_minChanged(slot, sigval1);
	});
}

void QLogValueAxis_maxChanged(QLogValueAxis* self, double max) {
	self->maxChanged(static_cast<qreal>(max));
}

void QLogValueAxis_connect_maxChanged(QLogValueAxis* self, intptr_t slot) {
	QLogValueAxis::connect(self, static_cast<void (QLogValueAxis::*)(qreal)>(&QLogValueAxis::maxChanged), self, [=](qreal max) {
		qreal max_ret = max;
		double sigval1 = static_cast<double>(max_ret);
		miqt_exec_callback_QLogValueAxis_maxChanged(slot, sigval1);
	});
}

void QLogValueAxis_rangeChanged(QLogValueAxis* self, double min, double max) {
	self->rangeChanged(static_cast<qreal>(min), static_cast<qreal>(max));
}

void QLogValueAxis_connect_rangeChanged(QLogValueAxis* self, intptr_t slot) {
	QLogValueAxis::connect(self, static_cast<void (QLogValueAxis::*)(qreal, qreal)>(&QLogValueAxis::rangeChanged), self, [=](qreal min, qreal max) {
		qreal min_ret = min;
		double sigval1 = static_cast<double>(min_ret);
		qreal max_ret = max;
		double sigval2 = static_cast<double>(max_ret);
		miqt_exec_callback_QLogValueAxis_rangeChanged(slot, sigval1, sigval2);
	});
}

void QLogValueAxis_labelFormatChanged(QLogValueAxis* self, struct miqt_string format) {
	QString format_QString = QString::fromUtf8(format.data, format.len);
	self->labelFormatChanged(format_QString);
}

void QLogValueAxis_connect_labelFormatChanged(QLogValueAxis* self, intptr_t slot) {
	QLogValueAxis::connect(self, static_cast<void (QLogValueAxis::*)(const QString&)>(&QLogValueAxis::labelFormatChanged), self, [=](const QString& format) {
		const QString format_ret = format;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray format_b = format_ret.toUtf8();
		struct miqt_string format_ms;
		format_ms.len = format_b.length();
		format_ms.data = static_cast<char*>(malloc(format_ms.len));
		memcpy(format_ms.data, format_b.data(), format_ms.len);
		struct miqt_string sigval1 = format_ms;
		miqt_exec_callback_QLogValueAxis_labelFormatChanged(slot, sigval1);
	});
}

void QLogValueAxis_baseChanged(QLogValueAxis* self, double base) {
	self->baseChanged(static_cast<qreal>(base));
}

void QLogValueAxis_connect_baseChanged(QLogValueAxis* self, intptr_t slot) {
	QLogValueAxis::connect(self, static_cast<void (QLogValueAxis::*)(qreal)>(&QLogValueAxis::baseChanged), self, [=](qreal base) {
		qreal base_ret = base;
		double sigval1 = static_cast<double>(base_ret);
		miqt_exec_callback_QLogValueAxis_baseChanged(slot, sigval1);
	});
}

void QLogValueAxis_tickCountChanged(QLogValueAxis* self, int tickCount) {
	self->tickCountChanged(static_cast<int>(tickCount));
}

void QLogValueAxis_connect_tickCountChanged(QLogValueAxis* self, intptr_t slot) {
	QLogValueAxis::connect(self, static_cast<void (QLogValueAxis::*)(int)>(&QLogValueAxis::tickCountChanged), self, [=](int tickCount) {
		int sigval1 = tickCount;
		miqt_exec_callback_QLogValueAxis_tickCountChanged(slot, sigval1);
	});
}

void QLogValueAxis_minorTickCountChanged(QLogValueAxis* self, int minorTickCount) {
	self->minorTickCountChanged(static_cast<int>(minorTickCount));
}

void QLogValueAxis_connect_minorTickCountChanged(QLogValueAxis* self, intptr_t slot) {
	QLogValueAxis::connect(self, static_cast<void (QLogValueAxis::*)(int)>(&QLogValueAxis::minorTickCountChanged), self, [=](int minorTickCount) {
		int sigval1 = minorTickCount;
		miqt_exec_callback_QLogValueAxis_minorTickCountChanged(slot, sigval1);
	});
}

struct miqt_string QLogValueAxis_tr2(const char* s, const char* c) {
	QString _ret = QLogValueAxis::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLogValueAxis_tr3(const char* s, const char* c, int n) {
	QString _ret = QLogValueAxis::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QLogValueAxis_override_virtual_type(void* self, intptr_t slot) {
	MiqtVirtualQLogValueAxis* self_cast = dynamic_cast<MiqtVirtualQLogValueAxis*>( (QLogValueAxis*)(self) );
	if (self_cast == nullptr) {
		return false;
	}

	self_cast->handle__type = slot;
	return true;
}

int QLogValueAxis_virtualbase_type(const void* self) {
	MiqtVirtualQLogValueAxis::AxisType _ret = static_cast<const MiqtVirtualQLogValueAxis*>(self)->QLogValueAxis::type();
	return static_cast<int>(_ret);
}

bool QLogValueAxis_override_virtual_event(void* self, intptr_t slot) {
	MiqtVirtualQLogValueAxis* self_cast = dynamic_cast<MiqtVirtualQLogValueAxis*>( (QLogValueAxis*)(self) );
	if (self_cast == nullptr) {
		return false;
	}

	self_cast->handle__event = slot;
	return true;
}

bool QLogValueAxis_virtualbase_event(void* self, QEvent* event) {
	return static_cast<MiqtVirtualQLogValueAxis*>(self)->QLogValueAxis::event(event);
}

bool QLogValueAxis_override_virtual_eventFilter(void* self, intptr_t slot) {
	MiqtVirtualQLogValueAxis* self_cast = dynamic_cast<MiqtVirtualQLogValueAxis*>( (QLogValueAxis*)(self) );
	if (self_cast == nullptr) {
		return false;
	}

	self_cast->handle__eventFilter = slot;
	return true;
}

bool QLogValueAxis_virtualbase_eventFilter(void* self, QObject* watched, QEvent* event) {
	return static_cast<MiqtVirtualQLogValueAxis*>(self)->QLogValueAxis::eventFilter(watched, event);
}

bool QLogValueAxis_override_virtual_timerEvent(void* self, intptr_t slot) {
	MiqtVirtualQLogValueAxis* self_cast = dynamic_cast<MiqtVirtualQLogValueAxis*>( (QLogValueAxis*)(self) );
	if (self_cast == nullptr) {
		return false;
	}

	self_cast->handle__timerEvent = slot;
	return true;
}

void QLogValueAxis_virtualbase_timerEvent(void* self, QTimerEvent* event) {
	static_cast<MiqtVirtualQLogValueAxis*>(self)->QLogValueAxis::timerEvent(event);
}

bool QLogValueAxis_override_virtual_childEvent(void* self, intptr_t slot) {
	MiqtVirtualQLogValueAxis* self_cast = dynamic_cast<MiqtVirtualQLogValueAxis*>( (QLogValueAxis*)(self) );
	if (self_cast == nullptr) {
		return false;
	}

	self_cast->handle__childEvent = slot;
	return true;
}

void QLogValueAxis_virtualbase_childEvent(void* self, QChildEvent* event) {
	static_cast<MiqtVirtualQLogValueAxis*>(self)->QLogValueAxis::childEvent(event);
}

bool QLogValueAxis_override_virtual_customEvent(void* self, intptr_t slot) {
	MiqtVirtualQLogValueAxis* self_cast = dynamic_cast<MiqtVirtualQLogValueAxis*>( (QLogValueAxis*)(self) );
	if (self_cast == nullptr) {
		return false;
	}

	self_cast->handle__customEvent = slot;
	return true;
}

void QLogValueAxis_virtualbase_customEvent(void* self, QEvent* event) {
	static_cast<MiqtVirtualQLogValueAxis*>(self)->QLogValueAxis::customEvent(event);
}

bool QLogValueAxis_override_virtual_connectNotify(void* self, intptr_t slot) {
	MiqtVirtualQLogValueAxis* self_cast = dynamic_cast<MiqtVirtualQLogValueAxis*>( (QLogValueAxis*)(self) );
	if (self_cast == nullptr) {
		return false;
	}

	self_cast->handle__connectNotify = slot;
	return true;
}

void QLogValueAxis_virtualbase_connectNotify(void* self, QMetaMethod* signal) {
	static_cast<MiqtVirtualQLogValueAxis*>(self)->QLogValueAxis::connectNotify(*signal);
}

bool QLogValueAxis_override_virtual_disconnectNotify(void* self, intptr_t slot) {
	MiqtVirtualQLogValueAxis* self_cast = dynamic_cast<MiqtVirtualQLogValueAxis*>( (QLogValueAxis*)(self) );
	if (self_cast == nullptr) {
		return false;
	}

	self_cast->handle__disconnectNotify = slot;
	return true;
}

void QLogValueAxis_virtualbase_disconnectNotify(void* self, QMetaMethod* signal) {
	static_cast<MiqtVirtualQLogValueAxis*>(self)->QLogValueAxis::disconnectNotify(*signal);
}

QObject* QLogValueAxis_protectedbase_sender(bool* _dynamic_cast_ok, const void* self) {
	MiqtVirtualQLogValueAxis* self_cast = dynamic_cast<MiqtVirtualQLogValueAxis*>( (QLogValueAxis*)(self) );
	if (self_cast == nullptr) {
		*_dynamic_cast_ok = false;
		return nullptr;
	}

	*_dynamic_cast_ok = true;
	return self_cast->sender();
}

int QLogValueAxis_protectedbase_senderSignalIndex(bool* _dynamic_cast_ok, const void* self) {
	MiqtVirtualQLogValueAxis* self_cast = dynamic_cast<MiqtVirtualQLogValueAxis*>( (QLogValueAxis*)(self) );
	if (self_cast == nullptr) {
		*_dynamic_cast_ok = false;
		return 0;
	}

	*_dynamic_cast_ok = true;
	return self_cast->senderSignalIndex();
}

int QLogValueAxis_protectedbase_receivers(bool* _dynamic_cast_ok, const void* self, const char* signal) {
	MiqtVirtualQLogValueAxis* self_cast = dynamic_cast<MiqtVirtualQLogValueAxis*>( (QLogValueAxis*)(self) );
	if (self_cast == nullptr) {
		*_dynamic_cast_ok = false;
		return 0;
	}

	*_dynamic_cast_ok = true;
	return self_cast->receivers(signal);
}

bool QLogValueAxis_protectedbase_isSignalConnected(bool* _dynamic_cast_ok, const void* self, QMetaMethod* signal) {
	MiqtVirtualQLogValueAxis* self_cast = dynamic_cast<MiqtVirtualQLogValueAxis*>( (QLogValueAxis*)(self) );
	if (self_cast == nullptr) {
		*_dynamic_cast_ok = false;
		return false;
	}

	*_dynamic_cast_ok = true;
	return self_cast->isSignalConnected(*signal);
}

void QLogValueAxis_delete(QLogValueAxis* self) {
	delete self;
}

