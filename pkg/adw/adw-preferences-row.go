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
	GTypePreferencesRow = coreglib.Type(C.adw_preferences_row_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePreferencesRow, F: marshalPreferencesRow},
	})
}

// PreferencesRowOverrides contains methods that are overridable.
type PreferencesRowOverrides struct {
}

func defaultPreferencesRowOverrides(v *PreferencesRow) PreferencesRowOverrides {
	return PreferencesRowOverrides{}
}

// PreferencesRow: gtk.ListBoxRow used to present preferences.
//
// The AdwPreferencesRow widget has a title that preferenceswindow will use to
// let the user look for a preference. It doesn't present the title in any way
// and lets you present the preference as you please.
//
// actionrow and its derivatives are convenient to use as preference rows as
// they take care of presenting the preference's title while letting you compose
// the inputs of the preference around it.
type PreferencesRow struct {
	_ [0]func() // equal guard
	gtk.ListBoxRow
}

var (
	_ gtk.Widgetter     = (*PreferencesRow)(nil)
	_ coreglib.Objector = (*PreferencesRow)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*PreferencesRow, *PreferencesRowClass, PreferencesRowOverrides](
		GTypePreferencesRow,
		initPreferencesRowClass,
		wrapPreferencesRow,
		defaultPreferencesRowOverrides,
	)
}

func initPreferencesRowClass(gclass unsafe.Pointer, overrides PreferencesRowOverrides, classInitFunc func(*PreferencesRowClass)) {
	if classInitFunc != nil {
		class := (*PreferencesRowClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapPreferencesRow(obj *coreglib.Object) *PreferencesRow {
	return &PreferencesRow{
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
	}
}

func marshalPreferencesRow(p uintptr) (interface{}, error) {
	return wrapPreferencesRow(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPreferencesRow creates a new AdwPreferencesRow.
//
// The function returns the following values:
//
//   - preferencesRow: newly created AdwPreferencesRow.
//
func NewPreferencesRow() *PreferencesRow {
	var _cret *C.GtkWidget // in

	_cret = C.adw_preferences_row_new()

	var _preferencesRow *PreferencesRow // out

	_preferencesRow = wrapPreferencesRow(coreglib.Take(unsafe.Pointer(_cret)))

	return _preferencesRow
}

// Title gets the title of the preference represented by self.
//
// The function returns the following values:
//
//   - utf8: title.
//
func (self *PreferencesRow) Title() string {
	var _arg0 *C.AdwPreferencesRow // out
	var _cret *C.char              // in

	_arg0 = (*C.AdwPreferencesRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_preferences_row_get_title(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// TitleSelectable gets whether the user can copy the title from the label.
//
// The function returns the following values:
//
//   - ok: whether the user can copy the title from the label.
//
func (self *PreferencesRow) TitleSelectable() bool {
	var _arg0 *C.AdwPreferencesRow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.AdwPreferencesRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_preferences_row_get_title_selectable(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UseMarkup gets whether to use Pango markup for the title label.
//
// The function returns the following values:
//
//   - ok: whether to use markup.
//
func (self *PreferencesRow) UseMarkup() bool {
	var _arg0 *C.AdwPreferencesRow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.AdwPreferencesRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_preferences_row_get_use_markup(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UseUnderline gets whether an embedded underline in the title indicates a
// mnemonic.
//
// The function returns the following values:
//
//   - ok: whether an embedded underline in the title indicates a mnemonic.
//
func (self *PreferencesRow) UseUnderline() bool {
	var _arg0 *C.AdwPreferencesRow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.AdwPreferencesRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_preferences_row_get_use_underline(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetTitle sets the title of the preference represented by self.
//
// The title is interpreted as Pango markup unless preferencesrow:use-markup is
// set to FALSE.
//
// The function takes the following parameters:
//
//   - title: title.
//
func (self *PreferencesRow) SetTitle(title string) {
	var _arg0 *C.AdwPreferencesRow // out
	var _arg1 *C.char              // out

	_arg0 = (*C.AdwPreferencesRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_preferences_row_set_title(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(title)
}

// SetTitleSelectable sets whether the user can copy the title from the label
//
// See also gtk.Label:selectable.
//
// The function takes the following parameters:
//
//   - titleSelectable: TRUE if the user can copy the title from the label.
//
func (self *PreferencesRow) SetTitleSelectable(titleSelectable bool) {
	var _arg0 *C.AdwPreferencesRow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.AdwPreferencesRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if titleSelectable {
		_arg1 = C.TRUE
	}

	C.adw_preferences_row_set_title_selectable(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(titleSelectable)
}

// SetUseMarkup sets whether to use Pango markup for the title label.
//
// Subclasses may also use it for other labels, such as subtitle.
//
// See also pango.ParseMarkup().
//
// The function takes the following parameters:
//
//   - useMarkup: whether to use markup.
//
func (self *PreferencesRow) SetUseMarkup(useMarkup bool) {
	var _arg0 *C.AdwPreferencesRow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.AdwPreferencesRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if useMarkup {
		_arg1 = C.TRUE
	}

	C.adw_preferences_row_set_use_markup(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(useMarkup)
}

// SetUseUnderline sets whether an embedded underline in the title indicates a
// mnemonic.
//
// The function takes the following parameters:
//
//   - useUnderline: TRUE if underlines in the text indicate mnemonics.
//
func (self *PreferencesRow) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.AdwPreferencesRow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.AdwPreferencesRow)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.adw_preferences_row_set_use_underline(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(useUnderline)
}

// PreferencesRowClass: instance of this type is always passed by reference.
type PreferencesRowClass struct {
	*preferencesRowClass
}

// preferencesRowClass is the struct that's finalized.
type preferencesRowClass struct {
	native *C.AdwPreferencesRowClass
}

// ParentClass: parent class.
func (p *PreferencesRowClass) ParentClass() *gtk.ListBoxRowClass {
	valptr := &p.native.parent_class
	var _v *gtk.ListBoxRowClass // out
	_v = (*gtk.ListBoxRowClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
