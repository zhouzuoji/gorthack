package gorthack

import (
	"reflect"
)

// MaybeCompatible checks if t1 and t2 maybe compatible
// unsafe.Pointer is compatible with any pointer type
func MaybeCompatible(t1, t2 reflect.Type) bool {
	if t1 == t2 {
		return true
	}

	if t1.Size() != t2.Size() || t1.Align() != t2.Align() || t1.FieldAlign() != t2.FieldAlign() {
		return false
	}

	k1, k2 := t1.Kind(), t2.Kind()

	ans := false
	switch k1 {
	case reflect.Array:
		ans = (k1 == k2) && t1.Len() == t2.Len() && MaybeCompatible(t1.Elem(), t2.Elem())
	case reflect.Chan:
		ans = (k1 == k2) && t1.ChanDir() == t2.ChanDir() && MaybeCompatible(t1.Elem(), t2.Elem())
	case reflect.Func:
		ans = isCompatibleFunc(t1, t2)
	case reflect.Interface:
		ans = t1 == t2
	case reflect.Map:
		ans = (k1 == k2) && MaybeCompatible(t1.Key(), t2.Key()) && MaybeCompatible(t1.Elem(), t2.Elem())
	case reflect.Ptr:
		ans = k2 == reflect.UnsafePointer || (k2 == reflect.Ptr && MaybeCompatible(t1.Elem(), t2.Elem()))
	case reflect.Slice:
		ans = (k1 == k2) && MaybeCompatible(t1.Elem(), t2.Elem())
	case reflect.Struct:
		ans = (k1 == k2) && isCompatibleStruct(t1, t2)
	case reflect.UnsafePointer:
		ans = k2 == reflect.UnsafePointer || k2 == reflect.Ptr
	default:
		ans = k1 == k2
	}
	return ans
}

func isCompatibleFunc(t1, t2 reflect.Type) bool {
	if t1.NumIn() != t2.NumIn() || t1.NumOut() != t2.NumOut() {
		return false
	}

	n := t1.NumIn()
	for i := 0; i < n; i++ {
		if !MaybeCompatible(t1.In(i), t2.In(i)) {
			return false
		}
	}

	n = t1.NumOut()
	for i := 0; i < n; i++ {
		if !MaybeCompatible(t1.Out(i), t2.Out(i)) {
			return false
		}
	}

	return true
}

func isCompatibleStruct(t1, t2 reflect.Type) bool {
	return false
}
