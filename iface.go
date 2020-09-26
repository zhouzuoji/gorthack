package gorthack

type emptyInterface struct {
	typ  *TypeInfo
	dptr unsafePtr
}

// DestructureInterface extract type-info and data pointer from an interface{}
func DestructureInterface(v interface{}) (ti *TypeInfo, dptr unsafePtr) {
	iface := (*emptyInterface)(unsafePtr(&v))
	return iface.typ, iface.dptr
}
