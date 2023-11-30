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
	GTypeClamp = coreglib.Type(C.adw_clamp_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeClamp, F: marshalClamp},
	})
}

// ClampOverrides contains methods that are overridable.
type ClampOverrides struct {
}

func defaultClampOverrides(v *Clamp) ClampOverrides {
	return ClampOverrides{}
}

// Clamp: widget constraining its child to a given size.
//
// <picture> <source srcset="clamp-wide-dark.png" media="(prefers-color-scheme:
// dark)"> <img src="clamp-wide.png" alt="clamp-wide"> </picture> <picture>
// <source srcset="clamp-narrow-dark.png" media="(prefers-color-scheme: dark)">
// <img src="clamp-narrow.png" alt="clamp-narrow"> </picture>
//
// The AdwClamp widget constrains the size of the widget it contains to a given
// maximum size. It will constrain the width if it is horizontal, or the height
// if it is vertical. The expansion of the child from its minimum to its maximum
// size is eased out for a smooth transition.
//
// If the child requires more than the requested maximum size, it will be
// allocated the minimum size it can fit in instead.
//
// AdwClamp can scale with the text scale factor, use the clamplayout:unit
// property to enable that behavior.
//
// # CSS nodes
//
// AdwClamp has a single CSS node with name clamp.
type Clamp struct {
	_ [0]func() // equal guard
	gtk.Widget

	*coreglib.Object
	gtk.Orientable
}

var (
	_ gtk.Widgetter     = (*Clamp)(nil)
	_ coreglib.Objector = (*Clamp)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Clamp, *ClampClass, ClampOverrides](
		GTypeClamp,
		initClampClass,
		wrapClamp,
		defaultClampOverrides,
	)
}

func initClampClass(gclass unsafe.Pointer, overrides ClampOverrides, classInitFunc func(*ClampClass)) {
	if classInitFunc != nil {
		class := (*ClampClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapClamp(obj *coreglib.Object) *Clamp {
	return &Clamp{
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
	}
}

func marshalClamp(p uintptr) (interface{}, error) {
	return wrapClamp(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewClamp creates a new AdwClamp.
//
// The function returns the following values:
//
//   - clamp: newly created AdwClamp.
//
func NewClamp() *Clamp {
	var _cret *C.GtkWidget // in

	_cret = C.adw_clamp_new()

	var _clamp *Clamp // out

	_clamp = wrapClamp(coreglib.Take(unsafe.Pointer(_cret)))

	return _clamp
}

// Child gets the child widget of self.
//
// The function returns the following values:
//
//   - widget (optional): child widget of self.
//
func (self *Clamp) Child() gtk.Widgetter {
	var _arg0 *C.AdwClamp  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.AdwClamp)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_clamp_get_child(_arg0)
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
func (self *Clamp) MaximumSize() int {
	var _arg0 *C.AdwClamp // out
	var _cret C.int       // in

	_arg0 = (*C.AdwClamp)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_clamp_get_maximum_size(_arg0)
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
func (self *Clamp) TighteningThreshold() int {
	var _arg0 *C.AdwClamp // out
	var _cret C.int       // in

	_arg0 = (*C.AdwClamp)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_clamp_get_tightening_threshold(_arg0)
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
func (self *Clamp) Unit() LengthUnit {
	var _arg0 *C.AdwClamp     // out
	var _cret C.AdwLengthUnit // in

	_arg0 = (*C.AdwClamp)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_clamp_get_unit(_arg0)
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
func (self *Clamp) SetChild(child gtk.Widgetter) {
	var _arg0 *C.AdwClamp  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.AdwClamp)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	C.adw_clamp_set_child(_arg0, _arg1)
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
func (self *Clamp) SetMaximumSize(maximumSize int) {
	var _arg0 *C.AdwClamp // out
	var _arg1 C.int       // out

	_arg0 = (*C.AdwClamp)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.int(maximumSize)

	C.adw_clamp_set_maximum_size(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(maximumSize)
}

// SetTighteningThreshold sets the size above which the child is clamped.
//
// Starting from this size, the clamp will tighten its grip on the child,
// slowly allocating less and less of the available size up to the maximum
// allocated size. Below that threshold and below the maximum size, the child
// will be allocated all the available size.
//
// If the threshold is greater than the maximum size to allocate to the child,
// the child will be allocated all the size up to the maximum. If the threshold
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
func (self *Clamp) SetTighteningThreshold(tighteningThreshold int) {
	var _arg0 *C.AdwClamp // out
	var _arg1 C.int       // out

	_arg0 = (*C.AdwClamp)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.int(tighteningThreshold)

	C.adw_clamp_set_tightening_threshold(_arg0, _arg1)
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
func (self *Clamp) SetUnit(unit LengthUnit) {
	var _arg0 *C.AdwClamp     // out
	var _arg1 C.AdwLengthUnit // out

	_arg0 = (*C.AdwClamp)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.AdwLengthUnit(unit)

	C.adw_clamp_set_unit(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(unit)
}

// ClampClass: instance of this type is always passed by reference.
type ClampClass struct {
	*clampClass
}

// clampClass is the struct that's finalized.
type clampClass struct {
	native *C.AdwClampClass
}

func (c *ClampClass) ParentClass() *gtk.WidgetClass {
	valptr := &c.native.parent_class
	var _v *gtk.WidgetClass // out
	_v = (*gtk.WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
