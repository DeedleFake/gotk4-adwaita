// Code generated by girgen. DO NOT EDIT.

package adw

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// #include <stdlib.h>
// #include <adwaita.h>
import "C"

// TabButtonClass: instance of this type is always passed by reference.
type TabButtonClass struct {
	*tabButtonClass
}

// tabButtonClass is the struct that's finalized.
type tabButtonClass struct {
	native *C.AdwTabButtonClass
}

func (t *TabButtonClass) ParentClass() *gtk.WidgetClass {
	valptr := &t.native.parent_class
	var _v *gtk.WidgetClass // out
	_v = (*gtk.WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
