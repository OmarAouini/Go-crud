package departments

// Employee type check
func IsEmployee(t interface{}) bool {
	switch t.(type) {
	case Department:
		return true
	default:
		return false
	}
}
