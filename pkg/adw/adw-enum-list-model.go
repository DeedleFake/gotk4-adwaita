// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <adwaita.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeEnumListItem  = coreglib.Type(C.adw_enum_list_item_get_type())
	GTypeEnumListModel = coreglib.Type(C.adw_enum_list_model_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeEnumListItem, F: marshalEnumListItem},
		coreglib.TypeMarshaler{T: GTypeEnumListModel, F: marshalEnumListModel},
	})
}

// EnumListItemOverrides contains methods that are overridable.
type EnumListItemOverrides struct {
}

func defaultEnumListItemOverrides(v *EnumListItem) EnumListItemOverrides {
	return EnumListItemOverrides{}
}

// EnumListItem: AdwEnumListItem is the type of items in a enumlistmodel.
type EnumListItem struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*EnumListItem)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*EnumListItem, *EnumListItemClass, EnumListItemOverrides](
		GTypeEnumListItem,
		initEnumListItemClass,
		wrapEnumListItem,
		defaultEnumListItemOverrides,
	)
}

func initEnumListItemClass(gclass unsafe.Pointer, overrides EnumListItemOverrides, classInitFunc func(*EnumListItemClass)) {
	if classInitFunc != nil {
		class := (*EnumListItemClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapEnumListItem(obj *coreglib.Object) *EnumListItem {
	return &EnumListItem{
		Object: obj,
	}
}

func marshalEnumListItem(p uintptr) (interface{}, error) {
	return wrapEnumListItem(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Name gets the enum value name.
//
// The function returns the following values:
//
//   - utf8: enum value name.
//
func (self *EnumListItem) Name() string {
	var _arg0 *C.AdwEnumListItem // out
	var _cret *C.char            // in

	_arg0 = (*C.AdwEnumListItem)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_enum_list_item_get_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Nick gets the enum value nick.
//
// The function returns the following values:
//
//   - utf8: enum value nick.
//
func (self *EnumListItem) Nick() string {
	var _arg0 *C.AdwEnumListItem // out
	var _cret *C.char            // in

	_arg0 = (*C.AdwEnumListItem)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_enum_list_item_get_nick(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Value gets the enum value.
//
// The function returns the following values:
//
//   - gint: enum value.
//
func (self *EnumListItem) Value() int {
	var _arg0 *C.AdwEnumListItem // out
	var _cret C.int              // in

	_arg0 = (*C.AdwEnumListItem)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_enum_list_item_get_value(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// EnumListModelOverrides contains methods that are overridable.
type EnumListModelOverrides struct {
}

func defaultEnumListModelOverrides(v *EnumListModel) EnumListModelOverrides {
	return EnumListModelOverrides{}
}

// EnumListModel: gio.ListModel representing values of a given enum.
//
// AdwEnumListModel contains objects of type enumlistitem.
type EnumListModel struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gio.ListModel
}

var (
	_ coreglib.Objector = (*EnumListModel)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*EnumListModel, *EnumListModelClass, EnumListModelOverrides](
		GTypeEnumListModel,
		initEnumListModelClass,
		wrapEnumListModel,
		defaultEnumListModelOverrides,
	)
}

func initEnumListModelClass(gclass unsafe.Pointer, overrides EnumListModelOverrides, classInitFunc func(*EnumListModelClass)) {
	if classInitFunc != nil {
		class := (*EnumListModelClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapEnumListModel(obj *coreglib.Object) *EnumListModel {
	return &EnumListModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalEnumListModel(p uintptr) (interface{}, error) {
	return wrapEnumListModel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewEnumListModel creates a new AdwEnumListModel for enum_type.
//
// The function takes the following parameters:
//
//   - enumType: type of the enum to construct the model from.
//
// The function returns the following values:
//
//   - enumListModel: newly created AdwEnumListModel.
//
func NewEnumListModel(enumType coreglib.Type) *EnumListModel {
	var _arg1 C.GType             // out
	var _cret *C.AdwEnumListModel // in

	_arg1 = C.GType(enumType)

	_cret = C.adw_enum_list_model_new(_arg1)
	runtime.KeepAlive(enumType)

	var _enumListModel *EnumListModel // out

	_enumListModel = wrapEnumListModel(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _enumListModel
}

// FindPosition finds the position of a given enum value in self.
//
// If the value is not found, GTK_INVALID_LIST_POSITION is returned.
//
// The function takes the following parameters:
//
//   - value: enum value.
//
// The function returns the following values:
//
func (self *EnumListModel) FindPosition(value int) uint {
	var _arg0 *C.AdwEnumListModel // out
	var _arg1 C.int               // out
	var _cret C.guint             // in

	_arg0 = (*C.AdwEnumListModel)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.int(value)

	_cret = C.adw_enum_list_model_find_position(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// EnumType gets the type of the enum represented by self.
//
// The function returns the following values:
//
//   - gType: enum type.
//
func (self *EnumListModel) EnumType() coreglib.Type {
	var _arg0 *C.AdwEnumListModel // out
	var _cret C.GType             // in

	_arg0 = (*C.AdwEnumListModel)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.adw_enum_list_model_get_enum_type(_arg0)
	runtime.KeepAlive(self)

	var _gType coreglib.Type // out

	_gType = coreglib.Type(_cret)

	return _gType
}

// EnumListItemClass: instance of this type is always passed by reference.
type EnumListItemClass struct {
	*enumListItemClass
}

// enumListItemClass is the struct that's finalized.
type enumListItemClass struct {
	native *C.AdwEnumListItemClass
}

// EnumListModelClass: instance of this type is always passed by reference.
type EnumListModelClass struct {
	*enumListModelClass
}

// enumListModelClass is the struct that's finalized.
type enumListModelClass struct {
	native *C.AdwEnumListModelClass
}
