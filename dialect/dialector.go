package dialect

var dialectorMap = map[string]Dialector{}

type Dialector interface {
	GetName() string
	GetDataTypeOf(interface{}) string
}