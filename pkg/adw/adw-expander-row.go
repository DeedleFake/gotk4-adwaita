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
	GTypeExpanderRow = coreglib.Type(C.adw_expander_row_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeExpanderRow, F: marshalExpanderRow},
	})
}

// ExpanderRowOverrides contains methods that are overridable.
type ExpanderRowOverrides struct {
}

func defaultExpanderRowOverrides(v *ExpanderRow) ExpanderRowOverrides {
	return ExpanderRowOverrides{}
}

// ExpanderRow: gtk.ListBoxRow used to reveal widgets.
//
// <picture> <source srcset="expander-row-dark.png"
// media="(prefers-color-scheme: dark)"> <img src="expander-row.png"
// alt="expander-row"> </picture>
//
// The AdwExpanderRow widget allows the user to reveal or hide widgets below it.
// It also allows the user to enable the expansion of the row, allowing to
// disable all that the row contains.
//
// # AdwExpanderRow as GtkBuildable
//
// The AdwExpanderRow implementation of the gtk.Buildable interface supports
// adding a child as an suffix widget by specifying “suffix” as the “type”
// attribute of a <child> element.
//
// It also supports adding it as a prefix widget by specifying “prefix” as the
// “type” attribute of a <child> element.
//
// # CSS nodes
//
// AdwExpanderRow has a main CSS node with name row and the .expander style
// class. It has the .empty style class when it contains no children.
//
// It contains the subnodes row.header for its main embedded row, list.nested
// for the list it can expand, and image.expander-row-arrow for its arrow.
type ExpanderRow struct {
	_ [0]func() // equal guard
	PreferencesRow
}

var (
	_ gtk.Widgetter     = (*ExpanderRow)(nil)
	_ coreglib.Objector = (*ExpanderRow)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ExpanderRow, *ExpanderRowClass, ExpanderRowOverrides](
		GTypeExpanderRow,
		initExpanderRowClass,
		wrapExpanderRow,
		defaultExpanderRowOverrides,
	)
}

func initExpanderRowClass(gclass unsafe.Pointer, overrides ExpanderRowOverrides, classInitFunc func(*ExpanderRowClass)) {
	if classInitFunc != nil {
		class := (*ExpanderRowClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapExpanderRow(obj *coreglib.Object) *ExpanderRow {
	return &ExpanderRow{
		PreferencesRow: PreferencesRow{
			ListBoxRow: gtk.ListBoxRow{
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
				Actionable: gtk.Actionable{
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
		},
	}
}

func marshalExpanderRow(p uintptr) (interface{}, error) {
	return wrapExpanderRow(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewExpanderRow creates a new AdwExpanderRow.
//
// The function returns the following values:
//
//   - expanderRow: newly created AdwExpanderRow.
//
func NewExpanderRow() *ExpanderRow {
	var _cret *C.GtkWidget // in

	_cret = C.adw_expander_row_new()

	var _expanderRow *ExpanderRow // out

	_expanderRow = wrapExpanderRow(coreglib.Take(unsafe.Pointer(_cret)))

	return _expanderRow
}

// AddAction adds an action widget to self.
//
// Deprecated: Use expanderrow.AddSuffix to add a suffix.
//
// The function takes the following parameters:
//
//   - widget: widget.
//
func (self *ExpanderRow) AddAction(widget gtk.Widgetter) {
	var _arg0 *C.AdwExpanderRow // out
	var _arg1 *C.GtkWidget      // out

	_arg0 = (*C.AdwExpanderRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	C.adw_expander_row_add_action(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}

// AddPrefix adds a prefix widget to self.
//
// The function takes the following parameters:
//
//   - widget: widget.
//
func (self *ExpanderRow) AddPrefix(widget gtk.Widgetter) {
	var _arg0 *C.AdwExpanderRow // out
	var _arg1 *C.GtkWidget      // out

	_arg0 = (*C.AdwExpanderRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	C.adw_expander_row_add_prefix(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}

// AddRow adds a widget to self.
//
// The widget will appear in the expanding list below self.
//
// The function takes the following parameters:
//
//   - child: widget.
//
func (self *ExpanderRow) AddRow(child gtk.Widgetter) {
	var _arg0 *C.AdwExpanderRow // out
	var _arg1 *C.GtkWidget      // out

	_arg0 = (*C.AdwExpanderRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.adw_expander_row_add_row(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// AddSuffix adds an suffix widget to self.
//
// The function takes the following parameters:
//
//   - widget: widget.
//
func (self *ExpanderRow) AddSuffix(widget gtk.Widgetter) {
	var _arg0 *C.AdwExpanderRow // out
	var _arg1 *C.GtkWidget      // out

	_arg0 = (*C.AdwExpanderRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(widget).Native()))

	C.adw_expander_row_add_suffix(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(widget)
}

// EnableExpansion gets whether the expansion of self is enabled.
//
// The function returns the following values:
//
//   - ok: whether the expansion of self is enabled.
//
func (self *ExpanderRow) EnableExpansion() bool {
	var _arg0 *C.AdwExpanderRow // out
	var _cret C.gboolean        // in

	_arg0 = (*C.AdwExpanderRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_expander_row_get_enable_expansion(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Expanded gets whether self is expanded.
//
// The function returns the following values:
//
//   - ok: whether self is expanded.
//
func (self *ExpanderRow) Expanded() bool {
	var _arg0 *C.AdwExpanderRow // out
	var _cret C.gboolean        // in

	_arg0 = (*C.AdwExpanderRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_expander_row_get_expanded(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IconName gets the icon name for self.
//
// Deprecated: Use expanderrow.AddPrefix to add an icon.
//
// The function returns the following values:
//
//   - utf8 (optional): icon name for self.
//
func (self *ExpanderRow) IconName() string {
	var _arg0 *C.AdwExpanderRow // out
	var _cret *C.char           // in

	_arg0 = (*C.AdwExpanderRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_expander_row_get_icon_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// ShowEnableSwitch gets whether the switch enabling the expansion of self is
// visible.
//
// The function returns the following values:
//
//   - ok: whether the switch enabling the expansion is visible.
//
func (self *ExpanderRow) ShowEnableSwitch() bool {
	var _arg0 *C.AdwExpanderRow // out
	var _cret C.gboolean        // in

	_arg0 = (*C.AdwExpanderRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_expander_row_get_show_enable_switch(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Subtitle gets the subtitle for self.
//
// The function returns the following values:
//
//   - utf8: subtitle for self.
//
func (self *ExpanderRow) Subtitle() string {
	var _arg0 *C.AdwExpanderRow // out
	var _cret *C.char           // in

	_arg0 = (*C.AdwExpanderRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_expander_row_get_subtitle(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SubtitleLines gets the number of lines at the end of which the subtitle label
// will be ellipsized.
//
// The function returns the following values:
//
//   - ok: number of lines at the end of which the subtitle label will be
//     ellipsized.
//
func (self *ExpanderRow) SubtitleLines() bool {
	var _arg0 *C.AdwExpanderRow // out
	var _cret C.gboolean        // in

	_arg0 = (*C.AdwExpanderRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_expander_row_get_subtitle_lines(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TitleLines gets the number of lines at the end of which the title label will
// be ellipsized.
//
// The function returns the following values:
//
//   - ok: number of lines at the end of which the title label will be
//     ellipsized.
//
func (self *ExpanderRow) TitleLines() bool {
	var _arg0 *C.AdwExpanderRow // out
	var _cret C.gboolean        // in

	_arg0 = (*C.AdwExpanderRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_expander_row_get_title_lines(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
func (self *ExpanderRow) Remove(child gtk.Widgetter) {
	var _arg0 *C.AdwExpanderRow // out
	var _arg1 *C.GtkWidget      // out

	_arg0 = (*C.AdwExpanderRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.adw_expander_row_remove(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetEnableExpansion sets whether the expansion of self is enabled.
//
// The function takes the following parameters:
//
//   - enableExpansion: whether to enable the expansion.
//
func (self *ExpanderRow) SetEnableExpansion(enableExpansion bool) {
	var _arg0 *C.AdwExpanderRow // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.AdwExpanderRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if enableExpansion {
		_arg1 = C.TRUE
	}

	C.adw_expander_row_set_enable_expansion(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(enableExpansion)
}

// SetExpanded sets whether self is expanded.
//
// The function takes the following parameters:
//
//   - expanded: whether to expand the row.
//
func (self *ExpanderRow) SetExpanded(expanded bool) {
	var _arg0 *C.AdwExpanderRow // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.AdwExpanderRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if expanded {
		_arg1 = C.TRUE
	}

	C.adw_expander_row_set_expanded(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(expanded)
}

// SetIconName sets the icon name for self.
//
// Deprecated: Use expanderrow.AddPrefix to add an icon.
//
// The function takes the following parameters:
//
//   - iconName (optional): icon name.
//
func (self *ExpanderRow) SetIconName(iconName string) {
	var _arg0 *C.AdwExpanderRow // out
	var _arg1 *C.char           // out

	_arg0 = (*C.AdwExpanderRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if iconName != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(iconName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.adw_expander_row_set_icon_name(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(iconName)
}

// SetShowEnableSwitch sets whether the switch enabling the expansion of self is
// visible.
//
// The function takes the following parameters:
//
//   - showEnableSwitch: whether to show the switch enabling the expansion.
//
func (self *ExpanderRow) SetShowEnableSwitch(showEnableSwitch bool) {
	var _arg0 *C.AdwExpanderRow // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.AdwExpanderRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if showEnableSwitch {
		_arg1 = C.TRUE
	}

	C.adw_expander_row_set_show_enable_switch(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(showEnableSwitch)
}

// SetSubtitle sets the subtitle for self.
//
// The subtitle is interpreted as Pango markup unless preferencesrow:use-markup
// is set to FALSE.
//
// The function takes the following parameters:
//
//   - subtitle: subtitle.
//
func (self *ExpanderRow) SetSubtitle(subtitle string) {
	var _arg0 *C.AdwExpanderRow // out
	var _arg1 *C.char           // out

	_arg0 = (*C.AdwExpanderRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(subtitle)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_expander_row_set_subtitle(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(subtitle)
}

// SetSubtitleLines sets the number of lines at the end of which the subtitle
// label will be ellipsized.
//
// If the value is 0, the number of lines won't be limited.
//
// The function takes the following parameters:
//
//   - subtitleLines: number of lines at the end of which the subtitle label
//     will be ellipsized.
//
func (self *ExpanderRow) SetSubtitleLines(subtitleLines int) {
	var _arg0 *C.AdwExpanderRow // out
	var _arg1 C.int             // out

	_arg0 = (*C.AdwExpanderRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.int(subtitleLines)

	C.adw_expander_row_set_subtitle_lines(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(subtitleLines)
}

// SetTitleLines sets the number of lines at the end of which the title label
// will be ellipsized.
//
// If the value is 0, the number of lines won't be limited.
//
// The function takes the following parameters:
//
//   - titleLines: number of lines at the end of which the title label will be
//     ellipsized.
//
func (self *ExpanderRow) SetTitleLines(titleLines int) {
	var _arg0 *C.AdwExpanderRow // out
	var _arg1 C.int             // out

	_arg0 = (*C.AdwExpanderRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.int(titleLines)

	C.adw_expander_row_set_title_lines(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(titleLines)
}

// ExpanderRowClass: instance of this type is always passed by reference.
type ExpanderRowClass struct {
	*expanderRowClass
}

// expanderRowClass is the struct that's finalized.
type expanderRowClass struct {
	native *C.AdwExpanderRowClass
}

// ParentClass: parent class.
func (e *ExpanderRowClass) ParentClass() *PreferencesRowClass {
	valptr := &e.native.parent_class
	var _v *PreferencesRowClass // out
	_v = (*PreferencesRowClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
