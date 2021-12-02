package dialect

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"unsafe"
)

func TestGetDataTypeOf(t *testing.T) {
	dialector := dialectorMap["mysql"]

	assert.Equal(t, "mysql", dialector.GetName(), "should equal")

	var testBool bool
	assert.Equal(t, "boolean", dialector.GetDataTypeOf(testBool), "should equal")

	var testInt8 int8
	assert.Equal(t, "tinyint", dialector.GetDataTypeOf(testInt8), "should equal")

	var testUint8 uint8
	assert.Equal(t, "tinyint unsigned", dialector.GetDataTypeOf(testUint8), "should equal")

	var testInt16 int16
	assert.Equal(t, "smallint", dialector.GetDataTypeOf(testInt16), "should equal")

	var testUint16 uint16
	assert.Equal(t, "smallint unsigned", dialector.GetDataTypeOf(testUint16), "should equal")

	var testInt32 int32
	assert.Equal(t, "integer", dialector.GetDataTypeOf(testInt32), "should equal")

	var testUint32 uint32
	assert.Equal(t, "integer unsigned", dialector.GetDataTypeOf(testUint32), "should equal")

	var testInt64 int64
	assert.Equal(t, "bigint", dialector.GetDataTypeOf(testInt64), "should equal")

	var testUint64 uint64
	assert.Equal(t, "bigint unsigned", dialector.GetDataTypeOf(testUint64), "should equal")

	var testInt int
	if size := unsafe.Sizeof(testInt); size == 2 {
		assert.Equal(t, "smallint", dialector.GetDataTypeOf(testInt), "should equal")
	} else if size == 4 {
		assert.Equal(t, "integer", dialector.GetDataTypeOf(testInt), "should equal")
	} else if size == 8 {
		assert.Equal(t, "bigint", dialector.GetDataTypeOf(testInt), "should equal")
	}

	var testUint uint
	if size := unsafe.Sizeof(testUint); size == 2 {
		assert.Equal(t, "smallint unsigned", dialector.GetDataTypeOf(testUint), "should equal")
	} else if size == 4 {
		assert.Equal(t, "integer unsigned", dialector.GetDataTypeOf(testUint), "should equal")
	} else if size == 8 {
		assert.Equal(t, "bigint unsigned", dialector.GetDataTypeOf(testUint), "should equal")
	}

	var testFloat32 float32
	assert.Equal(t, "float", dialector.GetDataTypeOf(testFloat32), "should equal")

	var testFloat64 float64
	assert.Equal(t, "double", dialector.GetDataTypeOf(testFloat64), "should equal")

	var testString string
	assert.Equal(t, "text", dialector.GetDataTypeOf(testString), "should equal")

	var testBytes []byte
	assert.Equal(t, "blob", dialector.GetDataTypeOf(testBytes), "should equal")

	var testTime time.Time
	assert.Equal(t, "datetime", dialector.GetDataTypeOf(testTime), "should equal")
}