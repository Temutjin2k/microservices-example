package dto

func ReadIntDefault(val, defval int) int {
	if val == 0 {
		return defval
	}
	return val
}

func ReadStringDefault(val, defval string) string {
	if val == "" {
		return defval
	}
	return val
}
