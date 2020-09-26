package gorthack

import (
	"reflect"
)

// CastFunc cast different function type
// src: function type
// dst: pointer to function type
//go:inline
func CastFunc(src, dst interface{}) bool {
	return castFunc(reflect.ValueOf(src), dst)
}

// MethodToPureFunc cast a method to a function which requires an receiver
func MethodToPureFunc(self, dst interface{}, methodName string) bool {
	m, ok := reflect.TypeOf(self).MethodByName(methodName)
	return ok && castFunc(m.Func, dst)
}

// CastMethod cast a method to a function which do not requires an receiver
func CastMethod(self, dst interface{}, methodName string) bool {
	return castFunc(reflect.ValueOf(self).MethodByName(methodName), dst)
}

func castFunc(src reflect.Value, dst interface{}) bool {
	t1 := src.Type()
	t2 := reflect.TypeOf(dst)
	if t1.Kind() != reflect.Func || t2.Kind() != reflect.Ptr {
		return false
	}
	t2 = t2.Elem()
	if !isCompatibleFunc(t1, t2) {
		return false
	}
	p := reflect.New(t1)
	p.Elem().Set(src)
	_, addr := DestructureInterface(dst)
	*Ptr2Ptr(unsafePtr(addr)) = *Ptr2Ptr(unsafePtr(p.Pointer()))
	return true
}
