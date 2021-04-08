package employee

// Employee type check
func IsEmployee(t interface{}) bool {
	switch t.(type) {
	case Employee:
		return true
	default:
		return false
	}
}
