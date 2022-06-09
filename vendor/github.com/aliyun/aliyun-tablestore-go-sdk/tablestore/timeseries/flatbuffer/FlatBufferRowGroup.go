// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatbuffer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FlatBufferRowGroup struct {
	_tab flatbuffers.Table
}

func GetRootAsFlatBufferRowGroup(buf []byte, offset flatbuffers.UOffsetT) *FlatBufferRowGroup {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FlatBufferRowGroup{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsFlatBufferRowGroup(buf []byte, offset flatbuffers.UOffsetT) *FlatBufferRowGroup {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &FlatBufferRowGroup{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *FlatBufferRowGroup) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FlatBufferRowGroup) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FlatBufferRowGroup) MeasurementName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *FlatBufferRowGroup) FieldNames(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *FlatBufferRowGroup) FieldNamesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *FlatBufferRowGroup) FieldTypes(j int) DataType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return DataType(rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1)))
	}
	return 0
}

func (rcv *FlatBufferRowGroup) FieldTypesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *FlatBufferRowGroup) MutateFieldTypes(j int, n DataType) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), int8(n))
	}
	return false
}

func (rcv *FlatBufferRowGroup) Rows(obj *FlatBufferRowInGroup, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *FlatBufferRowGroup) RowsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func FlatBufferRowGroupStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func FlatBufferRowGroupAddMeasurementName(builder *flatbuffers.Builder, measurementName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(measurementName), 0)
}
func FlatBufferRowGroupAddFieldNames(builder *flatbuffers.Builder, fieldNames flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(fieldNames), 0)
}
func FlatBufferRowGroupStartFieldNamesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func FlatBufferRowGroupAddFieldTypes(builder *flatbuffers.Builder, fieldTypes flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(fieldTypes), 0)
}
func FlatBufferRowGroupStartFieldTypesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func FlatBufferRowGroupAddRows(builder *flatbuffers.Builder, rows flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(rows), 0)
}
func FlatBufferRowGroupStartRowsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func FlatBufferRowGroupEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}