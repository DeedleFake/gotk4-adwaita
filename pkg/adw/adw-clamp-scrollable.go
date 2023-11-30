// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeClampScrollable = coreglib.Type(C.adw_clamp_scrollable_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeClampScrollable, F: marshalClampScrollable},
	})
}

// ClampScrollableOverrides contains methods that are overridable.
type ClampScrollableOverrides struct {
}

func defaultClampScrollableOverrides(v *ClampScrollable) ClampScrollableOverrides {
	return ClampScrollableOverrides{}
}

// ClampScrollable: scrollable clamp.
//
// AdwClampScrollable is a variant of clamp that implements the gtk.Scrollable
// interface.
//
// The primary use case for AdwClampScrollable is clamping gtk.ListView.
type ClampScrollable struct {
	_ [0]func() // equal guard
	gtk.Widget

	*coreglib.Object
	gtk.Orientable
	gtk.Scrollable
}

var (
	_ gtk.Widgetter     = (*ClampScrollable)(nil)
	_ coreglib.Objector = (*ClampScrollable)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ClampScrollable, *ClampScrollableClass, ClampScrollableOverrides](
		GTypeClampScrollable,
		initClampScrollableClass,
		wrapClampScrollable,
		defaultClampScrollableOverrides,
	)
}

func initClampScrollableClass(gclass unsafe.Pointer, overrides ClampScrollableOverrides, classInitFunc func(*ClampScrollableClass)) {
	if classInitFunc != nil {
		class := (*ClampScrollableClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapClampScrollable(obj *coreglib.Object) *ClampScrollable {
	return &ClampScrollable{
		Widget: gtk.Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: gtk.Accessible{
				Object: obj,
			},
			Buildable: gtk.Buildable{
				Object: obj,
			},
			ConstraintTarget: gtk.ConstraintTarget{
				Object: obj,
			},
		},
		Object: obj,
		Orientable: gtk.Orientable{
			Object: obj,
		},
		Scrollable: gtk.Scrollable{
			Object: obj,
		},
	}
}

func marshalClampScrollable(p uintptr) (interface{}, error) {
	return wrapClampScrollable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewClampScrollable creates a new AdwClampScrollable.
//
// The function returns the following values:
//
//   - clampScrollable: newly created AdwClampScrollable.
//
func NewClampScrollable() *ClampScrollable {
	var _cret *C.GtkWidget // in

	_cret = C.adw_clamp_scrollable_new()

	var _clampScrollable *ClampScrollable // out

	_clampScrollable = wrapClampScrollable(coreglib.Take(unsafe.Pointer(_cret)))

	return _clampScrollable
}

// Child gets the child widget of self.
//
// The function returns the following values:
//
//   - widget (optional): child widget of self.
//
func (self *ClampScrollable) Child() gtk.Widgetter {
	var _arg0 *C.AdwClampScrollable // out
	var _cret *C.GtkWidget          // in

	_arg0 = (*C.AdwClampScrollable)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_clamp_scrollable_get_child(_arg0)
	runtime.KeepAlive(self)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gtk.Widgetter)
				return ok
			})
			rv, ok := casted.(gtk.Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// MaximumSize gets the maximum size allocated to the child.
//
// The function returns the following values:
//
//   - gint: maximum size to allocate to the child.
//
func (self *ClampScrollable) MaximumSize() int {
	var _arg0 *C.AdwClampScrollable // out
	var _cret C.int                 // in

	_arg0 = (*C.AdwClampScrollable)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_clamp_scrollable_get_maximum_size(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// TighteningThreshold gets the size above which the child is clamped.
//
// The function returns the following values:
//
//   - gint: size above which the child is clamped.
//
func (self *ClampScrollable) TighteningThreshold() int {
	var _arg0 *C.AdwClampScrollable // out
	var _cret C.int                 // in

	_arg0 = (*C.AdwClampScrollable)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_clamp_scrollable_get_tightening_threshold(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Unit gets the length unit for maximum size and tightening threshold.
//
// The function returns the following values:
//
//   - lengthUnit: length unit.
//
func (self *ClampScrollable) Unit() LengthUnit {
	var _arg0 *C.AdwClampScrollable // out
	var _cret C.AdwLengthUnit       // in

	_arg0 = (*C.AdwClampScrollable)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_clamp_scrollable_get_unit(_arg0)
	runtime.KeepAlive(self)

	var _lengthUnit LengthUnit // out

	_lengthUnit = LengthUnit(_cret)

	return _lengthUnit
}

// SetChild sets the child widget of self.
//
// The function takes the following parameters:
//
//   - child (optional) widget.
//
func (self *ClampScrollable) SetChild(child gtk.Widgetter) {
	var _arg0 *C.AdwClampScrollable // out
	var _arg1 *C.GtkWidget          // out

	_arg0 = (*C.AdwClampScrollable)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	C.adw_clamp_scrollable_set_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetMaximumSize sets the maximum size allocated to the child.
//
// It is the width if the clamp is horizontal, or the height if it is vertical.
//
// The function takes the following parameters:
//
//   - maximumSize: maximum size.
//
func (self *ClampScrollable) SetMaximumSize(maximumSize int) {
	var _arg0 *C.AdwClampScrollable // out
	var _arg1 C.int                 // out

	_arg0 = (*C.AdwClampScrollable)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.int(maximumSize)

	C.adw_clamp_scrollable_set_maximum_size(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(maximumSize)
}

// SetTighteningThreshold sets the size above which the child is clamped.
//
// Starting from this size, the clamp will tighten its grip on the child, slowly
// allocating less and less of the available size up to the maximum allocated
// size. Below that threshold and below the maximum width, the child will be
// allocated all the available size.
//
// If the threshold is greater than the maximum size to allocate to the child,
// the child will be allocated all the width up to the maximum. If the threshold
// is lower than the minimum size to allocate to the child, that size will be
// used as the tightening threshold.
//
// Effectively, tightening the grip on the child before it reaches its maximum
// size makes transitions to and from the maximum size smoother when resizing.
//
// The function takes the following parameters:
//
//   - tighteningThreshold: tightening threshold.
//
func (self *ClampScrollable) SetTighteningThreshold(tighteningThreshold int) {
	var _arg0 *C.AdwClampScrollable // out
	var _arg1 C.int                 // out

	_arg0 = (*C.AdwClampScrollable)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.int(tighteningThreshold)

	C.adw_clamp_scrollable_set_tightening_threshold(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(tighteningThreshold)
}

// SetUnit sets the length unit for maximum size and tightening threshold.
//
// Allows the sizes to vary depending on the text scale factor.
//
// The function takes the following parameters:
//
//   - unit: length unit.
//
func (self *ClampScrollable) SetUnit(unit LengthUnit) {
	var _arg0 *C.AdwClampScrollable // out
	var _arg1 C.AdwLengthUnit       // out

	_arg0 = (*C.AdwClampScrollable)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.AdwLengthUnit(unit)

	C.adw_clamp_scrollable_set_unit(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(unit)
}

// ClampScrollableClass: instance of this type is always passed by reference.
type ClampScrollableClass struct {
	*clampScrollableClass
}

// clampScrollableClass is the struct that's finalized.
type clampScrollableClass struct {
	native *C.AdwClampScrollableClass
}

func (c *ClampScrollableClass) ParentClass() *gtk.WidgetClass {
	valptr := &c.native.parent_class
	var _v *gtk.WidgetClass // out
	_v = (*gtk.WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
