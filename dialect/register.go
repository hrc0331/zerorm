package dialect

func init() {
	registerDialector("mysql", &mysql{})
}

func registerDialector(name string, dialector Dialector) {
	dialectorMap[name] = dialector
}
