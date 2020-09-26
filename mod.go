package gorthack

import "unsafe"

// TypeInfo is same as reflect.rtype,
type TypeInfo struct {
	size       uintptr
	ptrdata    uintptr // number of bytes in the type that can contain pointers
	hash       uint32  // hash of type; avoids computation in hash tables
	tflag      uint8   // extra type information flags
	align      uint8   // alignment of variable with this type
	fieldAlign uint8   // alignment of struct field with this type
	kind       uint8   // enumeration for C
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal     func(unsafe.Pointer, unsafe.Pointer) bool
	gcdata    *byte // garbage collection data
	str       int32 // string form
	ptrToThis int32 // type for pointer to this type, may be zero
}

type (
	unsafePtr = unsafe.Pointer

	// Ptr2Ptr is pointer to unsafe.Pointer
	Ptr2Ptr = *unsafePtr
)

//go:linkname typedmemmove runtime.typedmemmove
func typedmemmove(ti *TypeInfo, dst, src unsafePtr)

//go:linkname memclrNoHeapPointers runtime.memclrNoHeapPointers
func memclrNoHeapPointers(ptr unsafePtr, n uintptr)

//go:linkname memclrHasPointers runtime.memclrHasPointers
func memclrHasPointers(ptr unsafePtr, n uintptr)

//go:linkname memequal runtime.memequal
func memequal(x, y unsafePtr, size uintptr) bool

//go:linkname memhash runtime.memhash
func memhash(p unsafePtr, h uintptr, size uintptr) uintptr

// Copy copy a value of specified type
//go:inline
func Copy(ti *TypeInfo, dst, src unsafePtr) {
	typedmemmove(ti, dst, src)
}

// ZeroMemoryPOD fills POD memory with zeros
//go:inline
func ZeroMemoryPOD(ptr unsafePtr, n uintptr) {
	memclrNoHeapPointers(ptr, n)
}

// ZeroMemory fills non-POD memory with zeros
//go:inline
func ZeroMemory(ptr unsafePtr, n uintptr) {
	memclrHasPointers(ptr, n)
}
