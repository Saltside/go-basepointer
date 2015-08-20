package basepointer

import "testing"

func TestString(t *testing.T) {
	s := "banana"
	p := StringPtr(s)

	if p == nil {
		t.Fatal("StringPtr(s) = nil, want address")
	}

	if s != *p {
		t.Errorf("&s = %v, want %v", &s, p)
	}
}

func TestRune(t *testing.T) {
	s := rune('x')
	p := RunePtr(s)

	if p == nil {
		t.Fatal("RunePtr(s) = nil, want address")
	}

	if s != *p {
		t.Errorf("&s = %v, want %v", &s, p)
	}
}

func TestUintPtr(t *testing.T) {
	v := uint(1337)
	p := UintPtr(v)

	if p == nil {
		t.Fatal("address = nil, want real address")
	}

	if v != *p {
		t.Errorf("*p = %v, want %v", *p, v)
	}
}

func TestUint8Ptr(t *testing.T) {
	v := uint8(13)
	p := Uint8Ptr(v)

	if p == nil {
		t.Fatal("address = nil, want real address")
	}

	if v != *p {
		t.Errorf("*p = %v, want %v", *p, v)
	}
}

func TestUint16Ptr(t *testing.T) {
	v := uint16(1337)
	p := Uint16Ptr(v)

	if p == nil {
		t.Fatal("address = nil, want real address")
	}

	if v != *p {
		t.Errorf("*p = %v, want %v", *p, v)
	}
}

func TestUint32Ptr(t *testing.T) {
	v := uint32(1337)
	p := Uint32Ptr(v)

	if p == nil {
		t.Fatal("address = nil, want real address")
	}

	if v != *p {
		t.Errorf("*p = %v, want %v", *p, v)
	}
}

func TestUint64Ptr(t *testing.T) {
	v := uint64(1337)
	p := Uint64Ptr(v)

	if p == nil {
		t.Fatal("address = nil, want real address")
	}

	if v != *p {
		t.Errorf("*p = %v, want %v", *p, v)
	}
}

func TestIntPtr(t *testing.T) {
	v := int(1337)
	p := IntPtr(v)

	if p == nil {
		t.Fatal("address = nil, want real address")
	}

	if v != *p {
		t.Errorf("*p = %v, want %v", *p, v)
	}
}

func TestInt8Ptr(t *testing.T) {
	v := int8(13)
	p := Int8Ptr(v)

	if p == nil {
		t.Fatal("address = nil, want real address")
	}

	if v != *p {
		t.Errorf("*p = %v, want %v", *p, v)
	}
}

func TestInt16Ptr(t *testing.T) {
	v := int16(1337)
	p := Int16Ptr(v)

	if p == nil {
		t.Fatal("address = nil, want real address")
	}

	if v != *p {
		t.Errorf("*p = %v, want %v", *p, v)
	}
}

func TestInt32Ptr(t *testing.T) {
	v := int32(1337)
	p := Int32Ptr(v)

	if p == nil {
		t.Fatal("address = nil, want real address")
	}

	if v != *p {
		t.Errorf("*p = %v, want %v", *p, v)
	}
}

func TestInt64Ptr(t *testing.T) {
	v := int64(1337)
	p := Int64Ptr(v)

	if p == nil {
		t.Fatal("address = nil, want real address")
	}

	if v != *p {
		t.Errorf("*p = %v, want %v", *p, v)
	}
}

func TestFloat32Ptr(t *testing.T) {
	v := float32(1337)
	p := Float32Ptr(v)

	if p == nil {
		t.Fatal("address = nil, want real address")
	}

	if v != *p {
		t.Errorf("*p = %v, want %v", *p, v)
	}
}

func TestFloat64Ptr(t *testing.T) {
	v := float64(1337)
	p := Float64Ptr(v)

	if p == nil {
		t.Fatal("address = nil, want real address")
	}

	if v != *p {
		t.Errorf("*p = %v, want %v", *p, v)
	}
}

func TestBytePtr(t *testing.T) {
	v := byte(8)
	p := BytePtr(v)

	if p == nil {
		t.Fatal("address = nil, want real address")
	}

	if v != *p {
		t.Errorf("*p = %v, want %v", *p, v)
	}
}

func TestBoolPtr(t *testing.T) {
	v := bool(true)
	p := BoolPtr(v)

	if p == nil {
		t.Fatal("address = nil, want real address")
	}

	if v != *p {
		t.Errorf("*p = %v, want %v", *p, v)
	}
}

func TestComplex64Ptr(t *testing.T) {
	v := complex(float32(3.1415), float32(9264))
	p := Complex64Ptr(v)

	if p == nil {
		t.Fatal("address = nil, want real address")
	}

	if v != *p {
		t.Errorf("*p = %v, want %v", *p, v)
	}
}

func TestComplex128Ptr(t *testing.T) {
	v := complex(float64(3.1415), float64(9264))
	p := Complex128Ptr(v)

	if p == nil {
		t.Fatal("address = nil, want real address")
	}

	if v != *p {
		t.Errorf("*p = %v, want %v", *p, v)
	}
}
