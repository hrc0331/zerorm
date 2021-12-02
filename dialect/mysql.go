package dialect

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"time"
)

// mysql dialector
type mysql struct {}

var _ Dialector = (*mysql)(nil)

func (s *mysql) GetName() string {
	return "mysql"
}

func (s *mysql) GetDataTypeOf(data interface{}) string {
	dataKind := reflect.Indirect(reflect.ValueOf(data)).Kind()
	switch dataKind {
	case reflect.Bool:
		return "boolean"
	case reflect.Int8, reflect.Uint8, reflect.Int16, reflect.Uint16, reflect.Int32, reflect.Uint32,
		reflect.Int64, reflect.Uint64, reflect.Int, reflect.Uint:
		return s.getDataTypeOfInt(data, dataKind)
	case reflect.Float32:
		return "float"
	case reflect.Float64:
		return "double"
	case reflect.String:
		return "text"
	case s.getBytesDataKind():
		return "blob"
	case s.getTimeDataKind():
		return "datetime"
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", reflect.TypeOf(data).Name(), dataKind.String()))
}

func (s *mysql) getDataTypeOfInt(data interface{}, dataKind reflect.Kind) (res string) {
	switch reflect.TypeOf(data).Size() {
	case 1:
		res = "tinyint"
	case 2:
		res = "smallint"
	case 4:
		res = "integer"
	case 8:
		res = "bigint"
	}
	if dataKind == reflect.Uint8 || dataKind == reflect.Uint16 || dataKind == reflect.Uint32 ||
		dataKind == reflect.Uint64 || dataKind == reflect.Uint {
		res += " unsigned"
	}

	return res
}

func (s *mysql) getBytesDataKind() reflect.Kind {
	var b []byte
	return reflect.TypeOf(b).Kind()
}

func (s *mysql) getTimeDataKind() reflect.Kind {
	var t time.Time
	return reflect.TypeOf(t).Kind()
}