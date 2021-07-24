// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: libadwaita-1
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.adw_view_stack_get_type()), F: marshalViewStacker},
		{T: externglib.Type(C.adw_view_stack_page_get_type()), F: marshalViewStackPager},
	})
}

// ViewStack: view container for adw.ViewSwitcher.
//
// AdwViewStack is a container which only shows one page at a time. It is
// typically used to hold an application's main views.
//
// It doesn't provide a way to transition between pages. Instead, a separate
// widget such as adw.ViewSwitcher can be used with AdwViewStack to provide this
// functionality.
//
// AdwViewStack pages can have a title, an icon, an attention request, and a
// numbered badge that adw.ViewSwitcher will use to let users identify which
// page is which. Set them using the adw.ViewStackPage:title,
// adw.ViewStackPage:icon-name, adw.ViewStackPage:needs-attention, and
// adw.ViewStackPage:badge-number properties.
//
// Transitions between views are animated by crossfading. These animations
// respect the gtk.Settings:gtk-enable-animations setting.
//
// AdwViewStack maintains a adw.ViewStackPage object for each added child, which
// holds additional per-child properties. You obtain the adw.ViewStackPage for a
// child with adw.ViewStack.GetPage() and you can obtain a gtk.SelectionModel
// containing all the pages with adw.ViewStack.GetPages().
//
//
// AdwViewStack as GtkBuildable
//
// To set child-specific properties in a .ui file, create adw.ViewStackPage
// objects explicitly, and set the child widget as a property on it:
//
//      <object class="AdwViewStack" id="stack">
//        <child>
//          <object class="AdwViewStackPage">
//            <property name="name">overview</property>
//            <property name="title">Overview</property>
//            <property name="child">
//              <object class="AdwStatusPage">
//                <property name="title">Welcome!</property>
//              </object>
//            </property>
//          </object>
//        </child>
//
//
//
// CSS nodes
//
// AdwViewStack has a single CSS node named stack.
type ViewStack struct {
	gtk.Widget
}

func wrapViewStack(obj *externglib.Object) *ViewStack {
	return &ViewStack{
		Widget: gtk.Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: gtk.Accessible{
				Object: obj,
			},
			Buildable: gtk.Buildable{
				Object: obj,
			},
			ConstraintTarget: gtk.ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalViewStacker(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapViewStack(obj), nil
}

// NewViewStack creates a new AdwViewStack.
func NewViewStack() *ViewStack {
	var _cret *C.GtkWidget // in

	_cret = C.adw_view_stack_new()

	var _viewStack *ViewStack // out

	_viewStack = wrapViewStack(externglib.Take(unsafe.Pointer(_cret)))

	return _viewStack
}

// Add adds a child to self.
func (self *ViewStack) Add(child gtk.Widgetter) *ViewStackPage {
	var _arg0 *C.AdwViewStack     // out
	var _arg1 *C.GtkWidget        // out
	var _cret *C.AdwViewStackPage // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.adw_view_stack_add(_arg0, _arg1)

	var _viewStackPage *ViewStackPage // out

	_viewStackPage = wrapViewStackPage(externglib.Take(unsafe.Pointer(_cret)))

	return _viewStackPage
}

// AddNamed adds a child to self.
//
// The child is identified by the name.
func (self *ViewStack) AddNamed(child gtk.Widgetter, name string) *ViewStackPage {
	var _arg0 *C.AdwViewStack     // out
	var _arg1 *C.GtkWidget        // out
	var _arg2 *C.char             // out
	var _cret *C.AdwViewStackPage // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if name != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.adw_view_stack_add_named(_arg0, _arg1, _arg2)

	var _viewStackPage *ViewStackPage // out

	_viewStackPage = wrapViewStackPage(externglib.Take(unsafe.Pointer(_cret)))

	return _viewStackPage
}

// AddTitled adds a child to self.
//
// The child is identified by the name. The title will be used by
// adw.ViewSwitcher to represent child, so it should be short.
func (self *ViewStack) AddTitled(child gtk.Widgetter, name string, title string) *ViewStackPage {
	var _arg0 *C.AdwViewStack     // out
	var _arg1 *C.GtkWidget        // out
	var _arg2 *C.char             // out
	var _arg3 *C.char             // out
	var _cret *C.AdwViewStackPage // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if name != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.adw_view_stack_add_titled(_arg0, _arg1, _arg2, _arg3)

	var _viewStackPage *ViewStackPage // out

	_viewStackPage = wrapViewStackPage(externglib.Take(unsafe.Pointer(_cret)))

	return _viewStackPage
}

// ChildByName finds the child with name in self.
func (self *ViewStack) ChildByName(name string) gtk.Widgetter {
	var _arg0 *C.AdwViewStack // out
	var _arg1 *C.char         // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.adw_view_stack_get_child_by_name(_arg0, _arg1)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gtk.Widgetter)
	}

	return _widget
}

// Hhomogeneous gets whether self is horizontally homogeneous.
func (self *ViewStack) Hhomogeneous() bool {
	var _arg0 *C.AdwViewStack // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_stack_get_hhomogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InterpolateSize gets whether self will interpolate its size when changing the
// visible child.
func (self *ViewStack) InterpolateSize() bool {
	var _arg0 *C.AdwViewStack // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_stack_get_interpolate_size(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Page gets the adw.ViewStackPage object for child.
func (self *ViewStack) Page(child gtk.Widgetter) *ViewStackPage {
	var _arg0 *C.AdwViewStack     // out
	var _arg1 *C.GtkWidget        // out
	var _cret *C.AdwViewStackPage // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.adw_view_stack_get_page(_arg0, _arg1)

	var _viewStackPage *ViewStackPage // out

	_viewStackPage = wrapViewStackPage(externglib.Take(unsafe.Pointer(_cret)))

	return _viewStackPage
}

// Pages returns a GListModel that contains the pages of the stack.
//
// This can be used to keep an up-to-date view. The model also implements
// gtk.SelectionModel and can be used to track and change the visible page.
func (self *ViewStack) Pages() gtk.SelectionModeller {
	var _arg0 *C.AdwViewStack      // out
	var _cret *C.GtkSelectionModel // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_stack_get_pages(_arg0)

	var _selectionModel gtk.SelectionModeller // out

	_selectionModel = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(gtk.SelectionModeller)

	return _selectionModel
}

// TransitionRunning gets whether the self is currently in a transition from one
// page to another.
func (self *ViewStack) TransitionRunning() bool {
	var _arg0 *C.AdwViewStack // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_stack_get_transition_running(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Vhomogeneous gets whether self is vertically homogeneous.
func (self *ViewStack) Vhomogeneous() bool {
	var _arg0 *C.AdwViewStack // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_stack_get_vhomogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// VisibleChild gets the currently visible child of self, .
func (self *ViewStack) VisibleChild() gtk.Widgetter {
	var _arg0 *C.AdwViewStack // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_stack_get_visible_child(_arg0)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gtk.Widgetter)
	}

	return _widget
}

// VisibleChildName returns the name of the currently visible child of self.
func (self *ViewStack) VisibleChildName() string {
	var _arg0 *C.AdwViewStack // out
	var _cret *C.char         // in

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_stack_get_visible_child_name(_arg0)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Remove removes a child widget from self.
func (self *ViewStack) Remove(child gtk.Widgetter) {
	var _arg0 *C.AdwViewStack // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.adw_view_stack_remove(_arg0, _arg1)
}

// SetHhomogeneous sets self to be horizontally homogeneous or not.
func (self *ViewStack) SetHhomogeneous(hhomogeneous bool) {
	var _arg0 *C.AdwViewStack // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(self.Native()))
	if hhomogeneous {
		_arg1 = C.TRUE
	}

	C.adw_view_stack_set_hhomogeneous(_arg0, _arg1)
}

// SetInterpolateSize sets whether self will interpolate its size when changing
// the visible child.
func (self *ViewStack) SetInterpolateSize(interpolateSize bool) {
	var _arg0 *C.AdwViewStack // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(self.Native()))
	if interpolateSize {
		_arg1 = C.TRUE
	}

	C.adw_view_stack_set_interpolate_size(_arg0, _arg1)
}

// SetVhomogeneous sets self to be vertically homogeneous or not.
func (self *ViewStack) SetVhomogeneous(vhomogeneous bool) {
	var _arg0 *C.AdwViewStack // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(self.Native()))
	if vhomogeneous {
		_arg1 = C.TRUE
	}

	C.adw_view_stack_set_vhomogeneous(_arg0, _arg1)
}

// SetVisibleChild makes child the visible child of self.
func (self *ViewStack) SetVisibleChild(child gtk.Widgetter) {
	var _arg0 *C.AdwViewStack // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.adw_view_stack_set_visible_child(_arg0, _arg1)
}

// SetVisibleChildName makes the child with name visible.
func (self *ViewStack) SetVisibleChildName(name string) {
	var _arg0 *C.AdwViewStack // out
	var _arg1 *C.char         // out

	_arg0 = (*C.AdwViewStack)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_view_stack_set_visible_child_name(_arg0, _arg1)
}

// ViewStackPage: auxiliary class used by adw.ViewStack.
type ViewStackPage struct {
	*externglib.Object
}

func wrapViewStackPage(obj *externglib.Object) *ViewStackPage {
	return &ViewStackPage{
		Object: obj,
	}
}

func marshalViewStackPager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapViewStackPage(obj), nil
}

// BadgeNumber gets the badge number for this page.
func (self *ViewStackPage) BadgeNumber() uint {
	var _arg0 *C.AdwViewStackPage // out
	var _cret C.guint             // in

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_stack_page_get_badge_number(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Child gets the stack child to which self belongs.
func (self *ViewStackPage) Child() gtk.Widgetter {
	var _arg0 *C.AdwViewStackPage // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_stack_page_get_child(_arg0)

	var _widget gtk.Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gtk.Widgetter)

	return _widget
}

// IconName gets the icon name of the page.
func (self *ViewStackPage) IconName() string {
	var _arg0 *C.AdwViewStackPage // out
	var _cret *C.char             // in

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_stack_page_get_icon_name(_arg0)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Name gets the name of the page.
func (self *ViewStackPage) Name() string {
	var _arg0 *C.AdwViewStackPage // out
	var _cret *C.char             // in

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_stack_page_get_name(_arg0)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// NeedsAttention gets whether the page is marked as “needs attention”.
func (self *ViewStackPage) NeedsAttention() bool {
	var _arg0 *C.AdwViewStackPage // out
	var _cret C.gboolean          // in

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_stack_page_get_needs_attention(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Title gets the page title.
func (self *ViewStackPage) Title() string {
	var _arg0 *C.AdwViewStackPage // out
	var _cret *C.char             // in

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_stack_page_get_title(_arg0)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// UseUnderline gets whether underlines in the page title indicate mnemonics.
func (self *ViewStackPage) UseUnderline() bool {
	var _arg0 *C.AdwViewStackPage // out
	var _cret C.gboolean          // in

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_stack_page_get_use_underline(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Visible gets whether self is visible in its AdwViewStack.
//
// This is independent from the gtk.Widget:visible property of its widget.
func (self *ViewStackPage) Visible() bool {
	var _arg0 *C.AdwViewStackPage // out
	var _cret C.gboolean          // in

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_view_stack_page_get_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetBadgeNumber sets the badge number for this page.
func (self *ViewStackPage) SetBadgeNumber(badgeNumber uint) {
	var _arg0 *C.AdwViewStackPage // out
	var _arg1 C.guint             // out

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(badgeNumber)

	C.adw_view_stack_page_set_badge_number(_arg0, _arg1)
}

// SetIconName sets the icon name of the page.
func (self *ViewStackPage) SetIconName(iconName string) {
	var _arg0 *C.AdwViewStackPage // out
	var _arg1 *C.char             // out

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(self.Native()))
	if iconName != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_view_stack_page_set_icon_name(_arg0, _arg1)
}

// SetName sets the name of the page.
func (self *ViewStackPage) SetName(name string) {
	var _arg0 *C.AdwViewStackPage // out
	var _arg1 *C.char             // out

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(self.Native()))
	if name != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_view_stack_page_set_name(_arg0, _arg1)
}

// SetNeedsAttention sets whether the page is marked as “needs attention”.
func (self *ViewStackPage) SetNeedsAttention(needsAttention bool) {
	var _arg0 *C.AdwViewStackPage // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(self.Native()))
	if needsAttention {
		_arg1 = C.TRUE
	}

	C.adw_view_stack_page_set_needs_attention(_arg0, _arg1)
}

// SetTitle sets the page title.
func (self *ViewStackPage) SetTitle(title string) {
	var _arg0 *C.AdwViewStackPage // out
	var _arg1 *C.char             // out

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(self.Native()))
	if title != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_view_stack_page_set_title(_arg0, _arg1)
}

// SetUseUnderline sets whether underlines in the page title indicate mnemonics.
func (self *ViewStackPage) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.AdwViewStackPage // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(self.Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.adw_view_stack_page_set_use_underline(_arg0, _arg1)
}

// SetVisible sets whether page is visible in its AdwViewStack.
func (self *ViewStackPage) SetVisible(visible bool) {
	var _arg0 *C.AdwViewStackPage // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.AdwViewStackPage)(unsafe.Pointer(self.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.adw_view_stack_page_set_visible(_arg0, _arg1)
}
