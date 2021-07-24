// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.adw_tab_page_get_type()), F: marshalTabPager},
		{T: externglib.Type(C.adw_tab_view_get_type()), F: marshalTabViewer},
	})
}

// TabPage: auxiliary class used by adw.TabView.
type TabPage struct {
	*externglib.Object
}

func wrapTabPage(obj *externglib.Object) *TabPage {
	return &TabPage{
		Object: obj,
	}
}

func marshalTabPager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTabPage(obj), nil
}

// Child gets the child of self.
func (self *TabPage) Child() gtk.Widgetter {
	var _arg0 *C.AdwTabPage // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.AdwTabPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_page_get_child(_arg0)

	var _widget gtk.Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gtk.Widgetter)

	return _widget
}

// Icon gets the icon of self.
func (self *TabPage) Icon() gio.Iconner {
	var _arg0 *C.AdwTabPage // out
	var _cret *C.GIcon      // in

	_arg0 = (*C.AdwTabPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_page_get_icon(_arg0)

	var _icon gio.Iconner // out

	if _cret != nil {
		_icon = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gio.Iconner)
	}

	return _icon
}

// IndicatorActivatable gets whether the indicator of self is activatable.
func (self *TabPage) IndicatorActivatable() bool {
	var _arg0 *C.AdwTabPage // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwTabPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_page_get_indicator_activatable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IndicatorIcon gets the indicator icon of self.
func (self *TabPage) IndicatorIcon() gio.Iconner {
	var _arg0 *C.AdwTabPage // out
	var _cret *C.GIcon      // in

	_arg0 = (*C.AdwTabPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_page_get_indicator_icon(_arg0)

	var _icon gio.Iconner // out

	if _cret != nil {
		_icon = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gio.Iconner)
	}

	return _icon
}

// Loading gets whether self is loading.
func (self *TabPage) Loading() bool {
	var _arg0 *C.AdwTabPage // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwTabPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_page_get_loading(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// NeedsAttention gets whether self needs attention.
func (self *TabPage) NeedsAttention() bool {
	var _arg0 *C.AdwTabPage // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwTabPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_page_get_needs_attention(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Parent gets the parent page of self.
func (self *TabPage) Parent() *TabPage {
	var _arg0 *C.AdwTabPage // out
	var _cret *C.AdwTabPage // in

	_arg0 = (*C.AdwTabPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_page_get_parent(_arg0)

	var _tabPage *TabPage // out

	if _cret != nil {
		_tabPage = wrapTabPage(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _tabPage
}

// Pinned gets whether self is pinned.
func (self *TabPage) Pinned() bool {
	var _arg0 *C.AdwTabPage // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwTabPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_page_get_pinned(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Selected gets whether self is selected.
func (self *TabPage) Selected() bool {
	var _arg0 *C.AdwTabPage // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwTabPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_page_get_selected(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Title gets the title of self.
func (self *TabPage) Title() string {
	var _arg0 *C.AdwTabPage // out
	var _cret *C.char       // in

	_arg0 = (*C.AdwTabPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_page_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Tooltip gets the tooltip of self.
func (self *TabPage) Tooltip() string {
	var _arg0 *C.AdwTabPage // out
	var _cret *C.char       // in

	_arg0 = (*C.AdwTabPage)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_page_get_tooltip(_arg0)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// SetIcon sets the icon of self.
func (self *TabPage) SetIcon(icon gio.Iconner) {
	var _arg0 *C.AdwTabPage // out
	var _arg1 *C.GIcon      // out

	_arg0 = (*C.AdwTabPage)(unsafe.Pointer(self.Native()))
	if icon != nil {
		_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))
	}

	C.adw_tab_page_set_icon(_arg0, _arg1)
}

// SetIndicatorActivatable sets whether the indicator of self is activatable.
func (self *TabPage) SetIndicatorActivatable(activatable bool) {
	var _arg0 *C.AdwTabPage // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.AdwTabPage)(unsafe.Pointer(self.Native()))
	if activatable {
		_arg1 = C.TRUE
	}

	C.adw_tab_page_set_indicator_activatable(_arg0, _arg1)
}

// SetIndicatorIcon sets the indicator icon of self.
func (self *TabPage) SetIndicatorIcon(indicatorIcon gio.Iconner) {
	var _arg0 *C.AdwTabPage // out
	var _arg1 *C.GIcon      // out

	_arg0 = (*C.AdwTabPage)(unsafe.Pointer(self.Native()))
	if indicatorIcon != nil {
		_arg1 = (*C.GIcon)(unsafe.Pointer(indicatorIcon.Native()))
	}

	C.adw_tab_page_set_indicator_icon(_arg0, _arg1)
}

// SetLoading sets wether self is loading.
func (self *TabPage) SetLoading(loading bool) {
	var _arg0 *C.AdwTabPage // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.AdwTabPage)(unsafe.Pointer(self.Native()))
	if loading {
		_arg1 = C.TRUE
	}

	C.adw_tab_page_set_loading(_arg0, _arg1)
}

// SetNeedsAttention sets whether self needs attention.
func (self *TabPage) SetNeedsAttention(needsAttention bool) {
	var _arg0 *C.AdwTabPage // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.AdwTabPage)(unsafe.Pointer(self.Native()))
	if needsAttention {
		_arg1 = C.TRUE
	}

	C.adw_tab_page_set_needs_attention(_arg0, _arg1)
}

// SetTitle sets the title of self.
func (self *TabPage) SetTitle(title string) {
	var _arg0 *C.AdwTabPage // out
	var _arg1 *C.char       // out

	_arg0 = (*C.AdwTabPage)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_tab_page_set_title(_arg0, _arg1)
}

// SetTooltip sets the tooltip of self.
func (self *TabPage) SetTooltip(tooltip string) {
	var _arg0 *C.AdwTabPage // out
	var _arg1 *C.char       // out

	_arg0 = (*C.AdwTabPage)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(tooltip)))
	defer C.free(unsafe.Pointer(_arg1))

	C.adw_tab_page_set_tooltip(_arg0, _arg1)
}

// TabView: dynamic tabbed container.
//
// AdwTabView is a container which shows one child at a time. While it provides
// keyboard shortcuts for switching between pages, it does not provide a visible
// tab bar and relies on external widgets for that, such as adw.TabBar.
//
// AdwTabView maintains a adw.TabPage object for each page, which holds
// additional per-page properties. You can obtain the AdwTabPage for a page with
// adw.TabView.GetPage(), and as the return value for adw.TabView.Append() and
// other functions for adding children.
//
// AdwTabView only aims to be useful for dynamic tabs in multi-window
// document-based applications, such as web browsers, file managers, text
// editors or terminals. It does not aim to replace gtk.Notebook for use cases
// such as tabbed dialogs.
//
// As such, it does not support disabling page reordering or detaching, or
// adding children via gtk.Builder.
//
//
// CSS nodes
//
// AdwTabView has a main CSS node with the name tabview.
type TabView struct {
	gtk.Widget
}

func wrapTabView(obj *externglib.Object) *TabView {
	return &TabView{
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

func marshalTabViewer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTabView(obj), nil
}

// NewTabView creates a new AdwTabView.
func NewTabView() *TabView {
	var _cret *C.AdwTabView // in

	_cret = C.adw_tab_view_new()

	var _tabView *TabView // out

	_tabView = wrapTabView(externglib.Take(unsafe.Pointer(_cret)))

	return _tabView
}

// AddPage adds child to self with parent as the parent.
//
// This function can be used to automatically position new pages, and to select
// the correct page when this page is closed while being selected (see
// adw.TabView.ClosePage()).
//
// If parent is NULL, this function is equivalent to adw.TabView.Append().
func (self *TabView) AddPage(child gtk.Widgetter, parent *TabPage) *TabPage {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 *C.AdwTabPage // out
	var _cret *C.AdwTabPage // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if parent != nil {
		_arg2 = (*C.AdwTabPage)(unsafe.Pointer(parent.Native()))
	}

	_cret = C.adw_tab_view_add_page(_arg0, _arg1, _arg2)

	var _tabPage *TabPage // out

	_tabPage = wrapTabPage(externglib.Take(unsafe.Pointer(_cret)))

	return _tabPage
}

// Append inserts child as the last non-pinned page.
func (self *TabView) Append(child gtk.Widgetter) *TabPage {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.GtkWidget  // out
	var _cret *C.AdwTabPage // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.adw_tab_view_append(_arg0, _arg1)

	var _tabPage *TabPage // out

	_tabPage = wrapTabPage(externglib.Take(unsafe.Pointer(_cret)))

	return _tabPage
}

// AppendPinned inserts child as the last pinned page.
func (self *TabView) AppendPinned(child gtk.Widgetter) *TabPage {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.GtkWidget  // out
	var _cret *C.AdwTabPage // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.adw_tab_view_append_pinned(_arg0, _arg1)

	var _tabPage *TabPage // out

	_tabPage = wrapTabPage(externglib.Take(unsafe.Pointer(_cret)))

	return _tabPage
}

// CloseOtherPages requests to close all pages other than page.
func (self *TabView) CloseOtherPages(page *TabPage) {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.AdwTabPage // out

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwTabPage)(unsafe.Pointer(page.Native()))

	C.adw_tab_view_close_other_pages(_arg0, _arg1)
}

// ClosePage requests to close page.
//
// Calling this function will result in the adw.TabView::close-page signal being
// emitted for page. Closing the page can then be confirmed or denied via
// adw.TabView.ClosePageFinish().
//
// If the page is waiting for a adw.TabView.ClosePageFinish() call, this
// function will do nothing.
//
// The default handler for adw.TabView::close-page will immediately confirm
// closing the page if it's non-pinned, or reject it if it's pinned. This
// behavior can be changed by registering your own handler for that signal.
//
// If page was selected, another page will be selected instead:
//
// If the adw.TabPage:parent value is NULL, the next page will be selected when
// possible, or if the page was already last, the previous page will be selected
// instead.
//
// If it's not NULL, the previous page will be selected if it's a descendant
// (possibly indirect) of the parent. If both the previous page and the parent
// are pinned, the parent will be selected instead.
func (self *TabView) ClosePage(page *TabPage) {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.AdwTabPage // out

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwTabPage)(unsafe.Pointer(page.Native()))

	C.adw_tab_view_close_page(_arg0, _arg1)
}

// ClosePageFinish completes a adw.TabView.ClosePage() call for page.
//
// If confirm is TRUE, page will be closed. If it's FALSE, it will be reverted
// to its previous state and adw.TabView.ClosePage() can be called for it again.
//
// This function should not be called unless a custom handler for
// adw.TabView::close-page is used.
func (self *TabView) ClosePageFinish(page *TabPage, confirm bool) {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.AdwTabPage // out
	var _arg2 C.gboolean    // out

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwTabPage)(unsafe.Pointer(page.Native()))
	if confirm {
		_arg2 = C.TRUE
	}

	C.adw_tab_view_close_page_finish(_arg0, _arg1, _arg2)
}

// ClosePagesAfter requests to close all pages after page.
func (self *TabView) ClosePagesAfter(page *TabPage) {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.AdwTabPage // out

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwTabPage)(unsafe.Pointer(page.Native()))

	C.adw_tab_view_close_pages_after(_arg0, _arg1)
}

// ClosePagesBefore requests to close all pages before page.
func (self *TabView) ClosePagesBefore(page *TabPage) {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.AdwTabPage // out

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwTabPage)(unsafe.Pointer(page.Native()))

	C.adw_tab_view_close_pages_before(_arg0, _arg1)
}

// DefaultIcon gets the default icon of self.
func (self *TabView) DefaultIcon() gio.Iconner {
	var _arg0 *C.AdwTabView // out
	var _cret *C.GIcon      // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_view_get_default_icon(_arg0)

	var _icon gio.Iconner // out

	_icon = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gio.Iconner)

	return _icon
}

// IsTransferringPage: whether a page is being transferred.
func (self *TabView) IsTransferringPage() bool {
	var _arg0 *C.AdwTabView // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_view_get_is_transferring_page(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MenuModel gets the tab context menu model for self.
func (self *TabView) MenuModel() gio.MenuModeller {
	var _arg0 *C.AdwTabView // out
	var _cret *C.GMenuModel // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_view_get_menu_model(_arg0)

	var _menuModel gio.MenuModeller // out

	if _cret != nil {
		_menuModel = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gio.MenuModeller)
	}

	return _menuModel
}

// NPages gets the number of pages in self.
func (self *TabView) NPages() int {
	var _arg0 *C.AdwTabView // out
	var _cret C.int         // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_view_get_n_pages(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NPinnedPages gets the number of pinned pages in self.
func (self *TabView) NPinnedPages() int {
	var _arg0 *C.AdwTabView // out
	var _cret C.int         // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_view_get_n_pinned_pages(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NthPage gets the adw.TabPage representing the child at position.
func (self *TabView) NthPage(position int) *TabPage {
	var _arg0 *C.AdwTabView // out
	var _arg1 C.int         // out
	var _cret *C.AdwTabPage // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = C.int(position)

	_cret = C.adw_tab_view_get_nth_page(_arg0, _arg1)

	var _tabPage *TabPage // out

	_tabPage = wrapTabPage(externglib.Take(unsafe.Pointer(_cret)))

	return _tabPage
}

// Page gets the adw.TabPage object representing child.
func (self *TabView) Page(child gtk.Widgetter) *TabPage {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.GtkWidget  // out
	var _cret *C.AdwTabPage // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.adw_tab_view_get_page(_arg0, _arg1)

	var _tabPage *TabPage // out

	_tabPage = wrapTabPage(externglib.Take(unsafe.Pointer(_cret)))

	return _tabPage
}

// PagePosition finds the position of page in self, starting from 0.
func (self *TabView) PagePosition(page *TabPage) int {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.AdwTabPage // out
	var _cret C.int         // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwTabPage)(unsafe.Pointer(page.Native()))

	_cret = C.adw_tab_view_get_page_position(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Pages returns a GListModel that contains the pages of self.
//
// This can be used to keep an up-to-date view. The model also implements
// gtk.SelectionModel and can be used to track and change the selected page.
func (self *TabView) Pages() gtk.SelectionModeller {
	var _arg0 *C.AdwTabView        // out
	var _cret *C.GtkSelectionModel // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_view_get_pages(_arg0)

	var _selectionModel gtk.SelectionModeller // out

	_selectionModel = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(gtk.SelectionModeller)

	return _selectionModel
}

// SelectedPage gets the currently selected page in self.
func (self *TabView) SelectedPage() *TabPage {
	var _arg0 *C.AdwTabView // out
	var _cret *C.AdwTabPage // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_view_get_selected_page(_arg0)

	var _tabPage *TabPage // out

	if _cret != nil {
		_tabPage = wrapTabPage(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _tabPage
}

// ShortcutWidget gets the shortcut widget for self.
func (self *TabView) ShortcutWidget() gtk.Widgetter {
	var _arg0 *C.AdwTabView // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_view_get_shortcut_widget(_arg0)

	var _widget gtk.Widgetter // out

	if _cret != nil {
		_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(gtk.Widgetter)
	}

	return _widget
}

// Insert inserts a non-pinned page at position.
//
// It's an error to try to insert a page before a pinned page, in that case
// adw.TabView.InsertPinned() should be used instead.
func (self *TabView) Insert(child gtk.Widgetter, position int) *TabPage {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.int         // out
	var _cret *C.AdwTabPage // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.int(position)

	_cret = C.adw_tab_view_insert(_arg0, _arg1, _arg2)

	var _tabPage *TabPage // out

	_tabPage = wrapTabPage(externglib.Take(unsafe.Pointer(_cret)))

	return _tabPage
}

// InsertPinned inserts a pinned page at position.
//
// It's an error to try to insert a pinned page after a non-pinned page, in that
// case adw.TabView.Insert() should be used instead.
func (self *TabView) InsertPinned(child gtk.Widgetter, position int) *TabPage {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.int         // out
	var _cret *C.AdwTabPage // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.int(position)

	_cret = C.adw_tab_view_insert_pinned(_arg0, _arg1, _arg2)

	var _tabPage *TabPage // out

	_tabPage = wrapTabPage(externglib.Take(unsafe.Pointer(_cret)))

	return _tabPage
}

// Prepend inserts child as the first non-pinned page.
func (self *TabView) Prepend(child gtk.Widgetter) *TabPage {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.GtkWidget  // out
	var _cret *C.AdwTabPage // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.adw_tab_view_prepend(_arg0, _arg1)

	var _tabPage *TabPage // out

	_tabPage = wrapTabPage(externglib.Take(unsafe.Pointer(_cret)))

	return _tabPage
}

// PrependPinned inserts child as the first pinned page.
func (self *TabView) PrependPinned(child gtk.Widgetter) *TabPage {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.GtkWidget  // out
	var _cret *C.AdwTabPage // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.adw_tab_view_prepend_pinned(_arg0, _arg1)

	var _tabPage *TabPage // out

	_tabPage = wrapTabPage(externglib.Take(unsafe.Pointer(_cret)))

	return _tabPage
}

// ReorderBackward reorders page to before its previous page if possible.
func (self *TabView) ReorderBackward(page *TabPage) bool {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.AdwTabPage // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwTabPage)(unsafe.Pointer(page.Native()))

	_cret = C.adw_tab_view_reorder_backward(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ReorderFirst reorders page to the first possible position.
func (self *TabView) ReorderFirst(page *TabPage) bool {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.AdwTabPage // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwTabPage)(unsafe.Pointer(page.Native()))

	_cret = C.adw_tab_view_reorder_first(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ReorderForward reorders page to after its next page if possible.
func (self *TabView) ReorderForward(page *TabPage) bool {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.AdwTabPage // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwTabPage)(unsafe.Pointer(page.Native()))

	_cret = C.adw_tab_view_reorder_forward(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ReorderLast reorders page to the last possible position.
func (self *TabView) ReorderLast(page *TabPage) bool {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.AdwTabPage // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwTabPage)(unsafe.Pointer(page.Native()))

	_cret = C.adw_tab_view_reorder_last(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ReorderPage reorders page to position.
//
// It's a programmer error to try to reorder a pinned page after a non-pinned
// one, or a non-pinned page before a pinned one.
func (self *TabView) ReorderPage(page *TabPage, position int) bool {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.AdwTabPage // out
	var _arg2 C.int         // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwTabPage)(unsafe.Pointer(page.Native()))
	_arg2 = C.int(position)

	_cret = C.adw_tab_view_reorder_page(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectNextPage selects the page after the currently selected page.
//
// If the last page was already selected, this function does nothing.
func (self *TabView) SelectNextPage() bool {
	var _arg0 *C.AdwTabView // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_view_select_next_page(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectPreviousPage selects the page before the currently selected page.
//
// If the first page was already selected, this function does nothing.
func (self *TabView) SelectPreviousPage() bool {
	var _arg0 *C.AdwTabView // out
	var _cret C.gboolean    // in

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))

	_cret = C.adw_tab_view_select_previous_page(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetDefaultIcon sets the default page icon for self.
func (self *TabView) SetDefaultIcon(defaultIcon gio.Iconner) {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.GIcon      // out

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer(defaultIcon.Native()))

	C.adw_tab_view_set_default_icon(_arg0, _arg1)
}

// SetMenuModel sets the tab context menu model for self.
func (self *TabView) SetMenuModel(menuModel gio.MenuModeller) {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.GMenuModel // out

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	if menuModel != nil {
		_arg1 = (*C.GMenuModel)(unsafe.Pointer(menuModel.Native()))
	}

	C.adw_tab_view_set_menu_model(_arg0, _arg1)
}

// SetPagePinned pins or unpins page.
//
// Pinned pages are guaranteed to be placed before all non-pinned pages; at any
// given moment the first adw.TabView:n-pinned-pages pages in self are
// guaranteed to be pinned.
//
// When a page is pinned or unpinned, it's automatically reordered: pinning a
// page moves it after other pinned pages; unpinning a page moves it before
// other non-pinned pages.
//
// Pinned pages can still be reordered between each other.
//
// adw.TabBar will display pinned pages in a compact form, never showing the
// title or close button, and only showing a single icon, selected in the
// following order:
//
// 1. adw.TabPage:indicator-icon 2. A spinner if adw.TabPage:loading is TRUE 3.
// adw.TabPage:icon 4. adw.TabView:default-icon
//
// Pinned pages cannot be closed by default, see adw.TabView::close-page for how
// to override that behavior.
//
// Changes the value of the adw.TabPage:pinned property.
func (self *TabView) SetPagePinned(page *TabPage, pinned bool) {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.AdwTabPage // out
	var _arg2 C.gboolean    // out

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwTabPage)(unsafe.Pointer(page.Native()))
	if pinned {
		_arg2 = C.TRUE
	}

	C.adw_tab_view_set_page_pinned(_arg0, _arg1, _arg2)
}

// SetSelectedPage sets the currently selected page in self.
func (self *TabView) SetSelectedPage(selectedPage *TabPage) {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.AdwTabPage // out

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwTabPage)(unsafe.Pointer(selectedPage.Native()))

	C.adw_tab_view_set_selected_page(_arg0, _arg1)
}

// SetShortcutWidget sets the shortcut widget for self.
func (self *TabView) SetShortcutWidget(widget gtk.Widgetter) {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}

	C.adw_tab_view_set_shortcut_widget(_arg0, _arg1)
}

// TransferPage transfers page from self to other_view.
//
// The page object will be reused.
//
// It's a programmer error to try to insert a pinned page after a non-pinned
// one, or a non-pinned page before a pinned one.
func (self *TabView) TransferPage(page *TabPage, otherView *TabView, position int) {
	var _arg0 *C.AdwTabView // out
	var _arg1 *C.AdwTabPage // out
	var _arg2 *C.AdwTabView // out
	var _arg3 C.int         // out

	_arg0 = (*C.AdwTabView)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.AdwTabPage)(unsafe.Pointer(page.Native()))
	_arg2 = (*C.AdwTabView)(unsafe.Pointer(otherView.Native()))
	_arg3 = C.int(position)

	C.adw_tab_view_transfer_page(_arg0, _arg1, _arg2, _arg3)
}
