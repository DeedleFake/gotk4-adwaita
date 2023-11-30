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
	GTypeApplicationWindow = coreglib.Type(C.adw_application_window_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeApplicationWindow, F: marshalApplicationWindow},
	})
}

// ApplicationWindowOverrides contains methods that are overridable.
type ApplicationWindowOverrides struct {
}

func defaultApplicationWindowOverrides(v *ApplicationWindow) ApplicationWindowOverrides {
	return ApplicationWindowOverrides{}
}

// ApplicationWindow: freeform application window.
//
// <picture> <source srcset="application-window-dark.png"
// media="(prefers-color-scheme: dark)"> <img src="application-window.png"
// alt="application-window"> </picture>
//
// AdwApplicationWindow is a gtk.ApplicationWindow subclass providing the same
// features as window.
//
// See window for details.
//
// Example of an AdwApplicationWindow UI definition:
//
//    <object class="AdwApplicationWindow">
//      <property name="content">
//        <object class="AdwToolbarView">
//          <child type="top">
//            <object class="AdwHeaderBar"/>
//          </child>
//          <property name="content">
//            <!-- ... -->
//          </property>
//        </object>
//      </property>
//    </object>
//
// Using gtk.Application:menubar is not supported and may result in visual
// glitches.
type ApplicationWindow struct {
	_ [0]func() // equal guard
	gtk.ApplicationWindow
}

var (
	_ coreglib.Objector = (*ApplicationWindow)(nil)
	_ gtk.Widgetter     = (*ApplicationWindow)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ApplicationWindow, *ApplicationWindowClass, ApplicationWindowOverrides](
		GTypeApplicationWindow,
		initApplicationWindowClass,
		wrapApplicationWindow,
		defaultApplicationWindowOverrides,
	)
}

func initApplicationWindowClass(gclass unsafe.Pointer, overrides ApplicationWindowOverrides, classInitFunc func(*ApplicationWindowClass)) {
	if classInitFunc != nil {
		class := (*ApplicationWindowClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapApplicationWindow(obj *coreglib.Object) *ApplicationWindow {
	return &ApplicationWindow{
		ApplicationWindow: gtk.ApplicationWindow{
			Window: gtk.Window{
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
				Root: gtk.Root{
					NativeSurface: gtk.NativeSurface{
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
					},
				},
				ShortcutManager: gtk.ShortcutManager{
					Object: obj,
				},
			},
			Object: obj,
			ActionGroup: gio.ActionGroup{
				Object: obj,
			},
			ActionMap: gio.ActionMap{
				Object: obj,
			},
		},
	}
}

func marshalApplicationWindow(p uintptr) (interface{}, error) {
	return wrapApplicationWindow(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewApplicationWindow creates a new AdwApplicationWindow for app.
//
// The function takes the following parameters:
//
//   - app: application instance.
//
// The function returns the following values:
//
//   - applicationWindow: newly created AdwApplicationWindow.
//
func NewApplicationWindow(app *gtk.Application) *ApplicationWindow {
	var _arg1 *C.GtkApplication // out
	var _cret *C.GtkWidget      // in

	_arg1 = (*C.GtkApplication)(unsafe.Pointer(coreglib.InternObject(app).Native()))

	_cret = C.adw_application_window_new(_arg1)
	runtime.KeepAlive(app)

	var _applicationWindow *ApplicationWindow // out

	_applicationWindow = wrapApplicationWindow(coreglib.Take(unsafe.Pointer(_cret)))

	return _applicationWindow
}

// AddBreakpoint adds breakpoint to self.
//
// The function takes the following parameters:
//
//   - breakpoint to add.
//
func (self *ApplicationWindow) AddBreakpoint(breakpoint *Breakpoint) {
	var _arg0 *C.AdwApplicationWindow // out
	var _arg1 *C.AdwBreakpoint        // out

	_arg0 = (*C.AdwApplicationWindow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.AdwBreakpoint)(unsafe.Pointer(coreglib.InternObject(breakpoint).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(breakpoint).Native()))

	C.adw_application_window_add_breakpoint(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(breakpoint)
}

// Content gets the content widget of self.
//
// This method should always be used instead of gtk.Window.GetChild().
//
// The function returns the following values:
//
//   - widget (optional): content widget of self.
//
func (self *ApplicationWindow) Content() gtk.Widgetter {
	var _arg0 *C.AdwApplicationWindow // out
	var _cret *C.GtkWidget            // in

	_arg0 = (*C.AdwApplicationWindow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_application_window_get_content(_arg0)
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

// CurrentBreakpoint gets the current breakpoint.
//
// The function returns the following values:
//
//   - breakpoint (optional): current breakpoint.
//
func (self *ApplicationWindow) CurrentBreakpoint() *Breakpoint {
	var _arg0 *C.AdwApplicationWindow // out
	var _cret *C.AdwBreakpoint        // in

	_arg0 = (*C.AdwApplicationWindow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_application_window_get_current_breakpoint(_arg0)
	runtime.KeepAlive(self)

	var _breakpoint *Breakpoint // out

	if _cret != nil {
		_breakpoint = wrapBreakpoint(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _breakpoint
}

// SetContent sets the content widget of self.
//
// This method should always be used instead of gtk.Window.SetChild().
//
// The function takes the following parameters:
//
//   - content (optional) widget.
//
func (self *ApplicationWindow) SetContent(content gtk.Widgetter) {
	var _arg0 *C.AdwApplicationWindow // out
	var _arg1 *C.GtkWidget            // out

	_arg0 = (*C.AdwApplicationWindow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if content != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(content).Native()))
	}

	C.adw_application_window_set_content(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(content)
}

// ApplicationWindowClass: instance of this type is always passed by reference.
type ApplicationWindowClass struct {
	*applicationWindowClass
}

// applicationWindowClass is the struct that's finalized.
type applicationWindowClass struct {
	native *C.AdwApplicationWindowClass
}

func (a *ApplicationWindowClass) ParentClass() *gtk.ApplicationWindowClass {
	valptr := &a.native.parent_class
	var _v *gtk.ApplicationWindowClass // out
	_v = (*gtk.ApplicationWindowClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
