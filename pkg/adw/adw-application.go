// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeApplication = coreglib.Type(C.adw_application_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeApplication, F: marshalApplication},
	})
}

// ApplicationOverrides contains methods that are overridable.
type ApplicationOverrides struct {
}

func defaultApplicationOverrides(v *Application) ApplicationOverrides {
	return ApplicationOverrides{}
}

// Application: base class for Adwaita applications.
//
// AdwApplication handles library initialization by calling init in the default
// gio.Application::startup signal handler, in turn chaining up as required
// by gtk.Application. Therefore, any subclass of AdwApplication should always
// chain up its startup handler before using any Adwaita or GTK API.
//
// # Automatic Resources
//
// AdwApplication will automatically load stylesheets located in the
// application's resource base path (see gio.Application.SetResourceBasePath(),
// if they're present.
//
// They can be used to add custom styles to the application, as follows:
//
// - style.css contains styles that are always present.
//
// - style-dark.css contains styles only used when stylemanager:dark is TRUE.
//
// - style-hc.css contains styles used when the system high contrast preference
// is enabled.
//
// - style-hc-dark.css contains styles used when the system high contrast
// preference is enabled and stylemanager:dark is TRUE.
type Application struct {
	_ [0]func() // equal guard
	gtk.Application
}

var (
	_ coreglib.Objector = (*Application)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Application, *ApplicationClass, ApplicationOverrides](
		GTypeApplication,
		initApplicationClass,
		wrapApplication,
		defaultApplicationOverrides,
	)
}

func initApplicationClass(gclass unsafe.Pointer, overrides ApplicationOverrides, classInitFunc func(*ApplicationClass)) {
	if classInitFunc != nil {
		class := (*ApplicationClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapApplication(obj *coreglib.Object) *Application {
	return &Application{
		Application: gtk.Application{
			Application: gio.Application{
				Object: obj,
				ActionGroup: gio.ActionGroup{
					Object: obj,
				},
				ActionMap: gio.ActionMap{
					Object: obj,
				},
			},
		},
	}
}

func marshalApplication(p uintptr) (interface{}, error) {
	return wrapApplication(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewApplication creates a new AdwApplication.
//
// If application_id is not NULL, then it must be valid. See
// gio.Application().IDIsValid.
//
// If no application ID is given then some features (most notably application
// uniqueness) will be disabled.
//
// The function takes the following parameters:
//
//   - applicationId (optional): application ID.
//   - flags: application flags.
//
// The function returns the following values:
//
//   - application: newly created AdwApplication.
//
func NewApplication(applicationId string, flags gio.ApplicationFlags) *Application {
	var _arg1 *C.char             // out
	var _arg2 C.GApplicationFlags // out
	var _cret *C.AdwApplication   // in

	if applicationId != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(applicationId)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = C.GApplicationFlags(flags)

	_cret = C.adw_application_new(_arg1, _arg2)
	runtime.KeepAlive(applicationId)
	runtime.KeepAlive(flags)

	var _application *Application // out

	_application = wrapApplication(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _application
}

// StyleManager gets the style manager for self.
//
// This is a convenience property allowing to access AdwStyleManager through
// property bindings or expressions.
//
// The function returns the following values:
//
//   - styleManager: style manager.
//
func (self *Application) StyleManager() *StyleManager {
	var _arg0 *C.AdwApplication  // out
	var _cret *C.AdwStyleManager // in

	_arg0 = (*C.AdwApplication)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_application_get_style_manager(_arg0)
	runtime.KeepAlive(self)

	var _styleManager *StyleManager // out

	_styleManager = wrapStyleManager(coreglib.Take(unsafe.Pointer(_cret)))

	return _styleManager
}

// ApplicationClass: instance of this type is always passed by reference.
type ApplicationClass struct {
	*applicationClass
}

// applicationClass is the struct that's finalized.
type applicationClass struct {
	native *C.AdwApplicationClass
}

// ParentClass: parent class.
func (a *ApplicationClass) ParentClass() *gtk.ApplicationClass {
	valptr := &a.native.parent_class
	var _v *gtk.ApplicationClass // out
	_v = (*gtk.ApplicationClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
