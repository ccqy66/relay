package utils

func MapForeach(base map[string]string, foreachFunc func(key string, value string)) {
	if base == nil || len(base) <= 0 {
		return
	}

	for key, value := range base{
		foreachFunc(key, value)
	}
}
