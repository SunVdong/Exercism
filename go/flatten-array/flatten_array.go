package flatten

func Flatten(nested interface{}) []interface{} {
	result := []interface{}{}
	var flattenHelper func(interface{})

	flattenHelper = func(input interface{}) {
		switch v := input.(type) {
		case []interface{}:
			for _, item := range v {
				flattenHelper(item)
			}
		case nil:
		default:
			result = append(result, v)
		}
	}

	flattenHelper(nested)
	return result
}
