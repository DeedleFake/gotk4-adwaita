// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
// extern void _gotk4_adw1_Breakpoint_ConnectUnapply(gpointer, guintptr);
// extern void _gotk4_adw1_Breakpoint_ConnectApply(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeBreakpointConditionLengthType = coreglib.Type(C.adw_breakpoint_condition_length_type_get_type())
	GTypeBreakpointConditionRatioType  = coreglib.Type(C.adw_breakpoint_condition_ratio_type_get_type())
	GTypeBreakpoint                    = coreglib.Type(C.adw_breakpoint_get_type())
	GTypeBreakpointCondition           = coreglib.Type(C.adw_breakpoint_condition_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeBreakpointConditionLengthType, F: marshalBreakpointConditionLengthType},
		coreglib.TypeMarshaler{T: GTypeBreakpointConditionRatioType, F: marshalBreakpointConditionRatioType},
		coreglib.TypeMarshaler{T: GTypeBreakpoint, F: marshalBreakpoint},
		coreglib.TypeMarshaler{T: GTypeBreakpointCondition, F: marshalBreakpointCondition},
	})
}

// BreakpointConditionLengthType describes length types for breakpointcondition.
//
// See breakpointcondition.NewLength.
//
// New values may be added to this enumeration over time.
type BreakpointConditionLengthType C.gint

const (
	// BreakpointConditionMinWidth: true if the width is greater than or equal
	// to the condition value.
	BreakpointConditionMinWidth BreakpointConditionLengthType = iota
	// BreakpointConditionMaxWidth: true if the width is less than or equal to
	// the condition value.
	BreakpointConditionMaxWidth
	// BreakpointConditionMinHeight: true if the height is greater than or equal
	// to the condition value.
	BreakpointConditionMinHeight
	// BreakpointConditionMaxHeight: true if the height is less than or equal to
	// the condition value.
	BreakpointConditionMaxHeight
)

func marshalBreakpointConditionLengthType(p uintptr) (interface{}, error) {
	return BreakpointConditionLengthType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for BreakpointConditionLengthType.
func (b BreakpointConditionLengthType) String() string {
	switch b {
	case BreakpointConditionMinWidth:
		return "MinWidth"
	case BreakpointConditionMaxWidth:
		return "MaxWidth"
	case BreakpointConditionMinHeight:
		return "MinHeight"
	case BreakpointConditionMaxHeight:
		return "MaxHeight"
	default:
		return fmt.Sprintf("BreakpointConditionLengthType(%d)", b)
	}
}

// BreakpointConditionRatioType describes ratio types for breakpointcondition.
//
// See breakpointcondition.NewRatio.
//
// New values may be added to this enumeration over time.
type BreakpointConditionRatioType C.gint

const (
	// BreakpointConditionMinAspectRatio: true if the aspect ratio is greater
	// than or equal to the condition value.
	BreakpointConditionMinAspectRatio BreakpointConditionRatioType = iota
	// BreakpointConditionMaxAspectRatio: true if the aspect ratio is less than
	// or equal to the condition value.
	BreakpointConditionMaxAspectRatio
)

func marshalBreakpointConditionRatioType(p uintptr) (interface{}, error) {
	return BreakpointConditionRatioType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for BreakpointConditionRatioType.
func (b BreakpointConditionRatioType) String() string {
	switch b {
	case BreakpointConditionMinAspectRatio:
		return "MinAspectRatio"
	case BreakpointConditionMaxAspectRatio:
		return "MaxAspectRatio"
	default:
		return fmt.Sprintf("BreakpointConditionRatioType(%d)", b)
	}
}

// BreakpointOverrides contains methods that are overridable.
type BreakpointOverrides struct {
}

func defaultBreakpointOverrides(v *Breakpoint) BreakpointOverrides {
	return BreakpointOverrides{}
}

// Breakpoint describes a breakpoint for window.
//
// Breakpoints are used to create adaptive UI, allowing to change the layout
// depending on available size.
//
// Breakpoint is a size threshold, specified by its condition, as well as one or
// more setters.
//
// Each setter has a target object, a property and a value. When a breakpoint is
// applied, each setter sets the target property on their target object to the
// specified value, and reset it back to the original value when it's unapplied.
//
// For more complicated scenarios, breakpoint::apply and breakpoint::unapply can
// be used instead.
//
// Breakpoints can be used within window, applicationwindow or breakpointbin.
//
// AdwBreakpoint as GtkBuildable:
//
// AdwBreakpoint supports specifying its condition via the <condition> element.
// The contents of the element must be a string in a format accepted by
// breakpointcondition.Parse().
//
// It also supports adding setters via the <setter> element. Each <setter>
// element must have the object attribute specifying the target object, and the
// property attribute specifying the property name. The contents of the element
// are used as the setter value.
//
// For G_TYPE_OBJECT and G_TYPE_BOXED derived properties, empty contents are
// treated as NULL.
//
// Setter values can be translated with the usual translatable, context and
// comments attributes.
//
// Example of an AdwBreakpoint UI definition:
//
//    <object class="AdwBreakpoint">
//      <condition>max-width: 400px</condition>
//      <setter object="button" property="visible">True</setter>
//      <setter object="box" property="orientation">vertical</setter>
//      <setter object="page" property="title" translatable="yes">Example</setter>
//    </object>.
type Breakpoint struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gtk.Buildable
}

var (
	_ coreglib.Objector = (*Breakpoint)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Breakpoint, *BreakpointClass, BreakpointOverrides](
		GTypeBreakpoint,
		initBreakpointClass,
		wrapBreakpoint,
		defaultBreakpointOverrides,
	)
}

func initBreakpointClass(gclass unsafe.Pointer, overrides BreakpointOverrides, classInitFunc func(*BreakpointClass)) {
	if classInitFunc != nil {
		class := (*BreakpointClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapBreakpoint(obj *coreglib.Object) *Breakpoint {
	return &Breakpoint{
		Object: obj,
		Buildable: gtk.Buildable{
			Object: obj,
		},
	}
}

func marshalBreakpoint(p uintptr) (interface{}, error) {
	return wrapBreakpoint(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectApply is emitted when the breakpoint is applied.
//
// This signal is emitted after the setters have been applied.
func (self *Breakpoint) ConnectApply(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "apply", false, unsafe.Pointer(C._gotk4_adw1_Breakpoint_ConnectApply), f)
}

// ConnectUnapply is emitted when the breakpoint is unapplied.
//
// This signal is emitted before resetting the setter values.
func (self *Breakpoint) ConnectUnapply(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "unapply", false, unsafe.Pointer(C._gotk4_adw1_Breakpoint_ConnectUnapply), f)
}

// NewBreakpoint creates a new AdwBreakpoint with condition.
//
// The function takes the following parameters:
//
//   - condition: condition.
//
// The function returns the following values:
//
//   - breakpoint: newly created AdwBreakpoint.
//
func NewBreakpoint(condition *BreakpointCondition) *Breakpoint {
	var _arg1 *C.AdwBreakpointCondition // out
	var _cret *C.AdwBreakpoint          // in

	_arg1 = (*C.AdwBreakpointCondition)(gextras.StructNative(unsafe.Pointer(condition)))
	runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(condition)), nil)

	_cret = C.adw_breakpoint_new(_arg1)
	runtime.KeepAlive(condition)

	var _breakpoint *Breakpoint // out

	_breakpoint = wrapBreakpoint(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _breakpoint
}

// AddSetter adds a setter to self.
//
// The setter will automatically set property on object to value when applying
// the breakpoint, and set it back to its original value upon unapplying it.
//
// Note that setting properties to their original values does not work for
// properties that have irreversible side effects. For example, changing
// gtk.Button:label while gtk.Button:icon-name is set will reset the icon.
// However, resetting the label will not set icon-name to its original value.
//
// Use the breakpoint::apply and breakpoint::unapply signals for those
// properties instead, as follows:
//
//    static void
//    breakpoint_apply_cb (MyWidget *self)
//    {
//      gtk_button_set_icon_name (self->button, "go-previous-symbolic");
//    }
//
//    static void
//    breakpoint_apply_cb (MyWidget *self)
//    {
//      gtk_button_set_label (self->button, _("_Back"));
//    }
//
//    // ...
//
//    g_signal_connect_swapped (breakpoint, "apply",
//                              G_CALLBACK (breakpoint_apply_cb), self);
//    g_signal_connect_swapped (breakpoint, "unapply",
//                              G_CALLBACK (breakpoint_unapply_cb), self);.
//
// The function takes the following parameters:
//
//   - object: target object.
//   - property: target property.
//   - value to set.
//
func (self *Breakpoint) AddSetter(object *coreglib.Object, property string, value *coreglib.Value) {
	var _arg0 *C.AdwBreakpoint // out
	var _arg1 *C.GObject       // out
	var _arg2 *C.char          // out
	var _arg3 *C.GValue        // out

	_arg0 = (*C.AdwBreakpoint)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GObject)(unsafe.Pointer(object.Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(property)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GValue)(unsafe.Pointer(value.Native()))

	C.adw_breakpoint_add_setter(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(self)
	runtime.KeepAlive(object)
	runtime.KeepAlive(property)
	runtime.KeepAlive(value)
}

// AddSetters adds n_setters setters to self.
//
// This is a convenience function for adding multiple setters at once.
//
// See breakpoint.AddSetter.
//
// This function is meant to be used by language bindings.
//
// The function takes the following parameters:
//
//   - objects: setter target object.
//   - names: setter target properties.
//   - values: setter values.
//
func (self *Breakpoint) AddSetters(objects []*coreglib.Object, names []string, values []*coreglib.Value) {
	var _arg0 *C.AdwBreakpoint // out
	var _arg2 **C.GObject      // out
	var _arg1 C.int
	var _arg3 **C.char   // out
	var _arg4 **C.GValue // out

	_arg0 = (*C.AdwBreakpoint)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (C.int)(len(objects))
	_arg2 = (**C.GObject)(C.calloc(C.size_t(len(objects)), C.size_t(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg2))
	{
		out := unsafe.Slice((**C.GObject)(_arg2), len(objects))
		for i := range objects {
			out[i] = (*C.GObject)(unsafe.Pointer(objects[i].Native()))
		}
	}
	_arg1 = (C.int)(len(names))
	_arg3 = (**C.char)(C.calloc(C.size_t(len(names)), C.size_t(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg3))
	{
		out := unsafe.Slice((**C.char)(_arg3), len(names))
		for i := range names {
			out[i] = (*C.char)(unsafe.Pointer(C.CString(names[i])))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}
	_arg1 = (C.int)(len(values))
	_arg4 = (**C.GValue)(C.calloc(C.size_t(len(values)), C.size_t(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg4))
	{
		out := unsafe.Slice((**C.GValue)(_arg4), len(values))
		for i := range values {
			out[i] = (*C.GValue)(unsafe.Pointer(values[i].Native()))
		}
	}

	C.adw_breakpoint_add_settersv(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(self)
	runtime.KeepAlive(objects)
	runtime.KeepAlive(names)
	runtime.KeepAlive(values)
}

// Condition gets the condition for self.
//
// The function returns the following values:
//
//   - breakpointCondition (optional): condition.
//
func (self *Breakpoint) Condition() *BreakpointCondition {
	var _arg0 *C.AdwBreakpoint          // out
	var _cret *C.AdwBreakpointCondition // in

	_arg0 = (*C.AdwBreakpoint)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_breakpoint_get_condition(_arg0)
	runtime.KeepAlive(self)

	var _breakpointCondition *BreakpointCondition // out

	if _cret != nil {
		_breakpointCondition = (*BreakpointCondition)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _breakpointCondition
}

// SetCondition sets the condition for self.
//
// The function takes the following parameters:
//
//   - condition (optional): new condition.
//
func (self *Breakpoint) SetCondition(condition *BreakpointCondition) {
	var _arg0 *C.AdwBreakpoint          // out
	var _arg1 *C.AdwBreakpointCondition // out

	_arg0 = (*C.AdwBreakpoint)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if condition != nil {
		_arg1 = (*C.AdwBreakpointCondition)(gextras.StructNative(unsafe.Pointer(condition)))
	}

	C.adw_breakpoint_set_condition(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(condition)
}

// BreakpointCondition describes condition for an breakpoint.
//
// An instance of this type is always passed by reference.
type BreakpointCondition struct {
	*breakpointCondition
}

// breakpointCondition is the struct that's finalized.
type breakpointCondition struct {
	native *C.AdwBreakpointCondition
}

func marshalBreakpointCondition(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &BreakpointCondition{&breakpointCondition{(*C.AdwBreakpointCondition)(b)}}, nil
}

// NewBreakpointConditionAnd constructs a struct BreakpointCondition.
func NewBreakpointConditionAnd(condition1 *BreakpointCondition, condition2 *BreakpointCondition) *BreakpointCondition {
	var _arg1 *C.AdwBreakpointCondition // out
	var _arg2 *C.AdwBreakpointCondition // out
	var _cret *C.AdwBreakpointCondition // in

	_arg1 = (*C.AdwBreakpointCondition)(gextras.StructNative(unsafe.Pointer(condition1)))
	runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(condition1)), nil)
	_arg2 = (*C.AdwBreakpointCondition)(gextras.StructNative(unsafe.Pointer(condition2)))
	runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(condition2)), nil)

	_cret = C.adw_breakpoint_condition_new_and(_arg1, _arg2)
	runtime.KeepAlive(condition1)
	runtime.KeepAlive(condition2)

	var _breakpointCondition *BreakpointCondition // out

	_breakpointCondition = (*BreakpointCondition)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_breakpointCondition)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.adw_breakpoint_condition_free((*C.AdwBreakpointCondition)(intern.C))
		},
	)

	return _breakpointCondition
}

// NewBreakpointConditionLength constructs a struct BreakpointCondition.
func NewBreakpointConditionLength(typ BreakpointConditionLengthType, value float64, unit LengthUnit) *BreakpointCondition {
	var _arg1 C.AdwBreakpointConditionLengthType // out
	var _arg2 C.double                           // out
	var _arg3 C.AdwLengthUnit                    // out
	var _cret *C.AdwBreakpointCondition          // in

	_arg1 = C.AdwBreakpointConditionLengthType(typ)
	_arg2 = C.double(value)
	_arg3 = C.AdwLengthUnit(unit)

	_cret = C.adw_breakpoint_condition_new_length(_arg1, _arg2, _arg3)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(value)
	runtime.KeepAlive(unit)

	var _breakpointCondition *BreakpointCondition // out

	_breakpointCondition = (*BreakpointCondition)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_breakpointCondition)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.adw_breakpoint_condition_free((*C.AdwBreakpointCondition)(intern.C))
		},
	)

	return _breakpointCondition
}

// NewBreakpointConditionOr constructs a struct BreakpointCondition.
func NewBreakpointConditionOr(condition1 *BreakpointCondition, condition2 *BreakpointCondition) *BreakpointCondition {
	var _arg1 *C.AdwBreakpointCondition // out
	var _arg2 *C.AdwBreakpointCondition // out
	var _cret *C.AdwBreakpointCondition // in

	_arg1 = (*C.AdwBreakpointCondition)(gextras.StructNative(unsafe.Pointer(condition1)))
	runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(condition1)), nil)
	_arg2 = (*C.AdwBreakpointCondition)(gextras.StructNative(unsafe.Pointer(condition2)))
	runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(condition2)), nil)

	_cret = C.adw_breakpoint_condition_new_or(_arg1, _arg2)
	runtime.KeepAlive(condition1)
	runtime.KeepAlive(condition2)

	var _breakpointCondition *BreakpointCondition // out

	_breakpointCondition = (*BreakpointCondition)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_breakpointCondition)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.adw_breakpoint_condition_free((*C.AdwBreakpointCondition)(intern.C))
		},
	)

	return _breakpointCondition
}

// NewBreakpointConditionRatio constructs a struct BreakpointCondition.
func NewBreakpointConditionRatio(typ BreakpointConditionRatioType, width int, height int) *BreakpointCondition {
	var _arg1 C.AdwBreakpointConditionRatioType // out
	var _arg2 C.int                             // out
	var _arg3 C.int                             // out
	var _cret *C.AdwBreakpointCondition         // in

	_arg1 = C.AdwBreakpointConditionRatioType(typ)
	_arg2 = C.int(width)
	_arg3 = C.int(height)

	_cret = C.adw_breakpoint_condition_new_ratio(_arg1, _arg2, _arg3)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _breakpointCondition *BreakpointCondition // out

	_breakpointCondition = (*BreakpointCondition)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_breakpointCondition)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.adw_breakpoint_condition_free((*C.AdwBreakpointCondition)(intern.C))
		},
	)

	return _breakpointCondition
}

// Copy copies self.
//
// The function returns the following values:
//
//   - breakpointCondition: copy of self.
//
func (self *BreakpointCondition) Copy() *BreakpointCondition {
	var _arg0 *C.AdwBreakpointCondition // out
	var _cret *C.AdwBreakpointCondition // in

	_arg0 = (*C.AdwBreakpointCondition)(gextras.StructNative(unsafe.Pointer(self)))

	_cret = C.adw_breakpoint_condition_copy(_arg0)
	runtime.KeepAlive(self)

	var _breakpointCondition *BreakpointCondition // out

	_breakpointCondition = (*BreakpointCondition)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_breakpointCondition)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.adw_breakpoint_condition_free((*C.AdwBreakpointCondition)(intern.C))
		},
	)

	return _breakpointCondition
}

// String returns a textual representation of self.
//
// The returned string can be parsed by breakpointcondition.Parse().
//
// The function returns the following values:
//
//   - utf8: newly allocated text string.
//
func (self *BreakpointCondition) String() string {
	var _arg0 *C.AdwBreakpointCondition // out
	var _cret *C.char                   // in

	_arg0 = (*C.AdwBreakpointCondition)(gextras.StructNative(unsafe.Pointer(self)))

	_cret = C.adw_breakpoint_condition_to_string(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// BreakpointConditionParse parses a condition from a string.
//
// Length conditions are specified as <type>: <value>[<unit>], where:
//
// - <type> can be min-width, max-width, min-height or max-height
//
// - <value> is a fractional number
//
// - <unit> can be px, pt or sp
//
// If the unit is omitted, px is assumed.
//
// See breakpointcondition.NewLength.
//
// Examples:
//
// - min-width: 500px
//
// - min-height: 400pt
//
// - max-width: 100sp
//
// - max-height: 500
//
// Ratio conditions are specified as <type>: <width>[/<height>], where:
//
// - <type> can be min-aspect-ratio or max-aspect-ratio
//
// - <width> and <height> are integer numbers
//
// See breakpointcondition.NewRatio.
//
// The ratio is represented as <width> divided by <height>.
//
// If <height> is omitted, it's assumed to be 1.
//
// Examples:
//
// - min-aspect-ratio: 4/3
//
// - max-aspect-ratio: 1
//
// The logical operators and, or can be used to compose a complex condition as
// follows:
//
// - <condition> and <condition>: the condition is true when both <condition>s
// are true, same as when using breakpointcondition.NewAnd
//
// - <condition> or <condition>: the condition is true when either of the
// <condition>s is true, same as when using breakpointcondition.NewOr
//
// Examples:
//
// - min-width: 400px and max-aspect-ratio: 4/3
//
// - max-width: 360sp or max-width: 360px
//
// Conditions can be further nested using parentheses, for example:
//
// - min-width: 400px and (max-aspect-ratio: 4/3 or max-height: 400px)
//
// If parentheses are omitted, the first operator takes priority.
//
// The function takes the following parameters:
//
//   - str: string specifying the condition.
//
// The function returns the following values:
//
//   - breakpointCondition: parsed condition.
//
func BreakpointConditionParse(str string) *BreakpointCondition {
	var _arg1 *C.char                   // out
	var _cret *C.AdwBreakpointCondition // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.adw_breakpoint_condition_parse(_arg1)
	runtime.KeepAlive(str)

	var _breakpointCondition *BreakpointCondition // out

	_breakpointCondition = (*BreakpointCondition)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_breakpointCondition)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.adw_breakpoint_condition_free((*C.AdwBreakpointCondition)(intern.C))
		},
	)

	return _breakpointCondition
}
