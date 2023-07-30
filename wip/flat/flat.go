package flat

import (
	"reflect"
	"strings"
)

func Expand(value map[string]interface{}) map[string]interface{} {
	return ExpandPrefixed(value, "")
}

func ExpandPrefixed(value map[string]interface{}, prefix string) map[string]interface{} {
	m := make(map[string]interface{})
	ExpandPrefixedToResult(value, prefix, m)
	return m
}

func ExpandPrefixedToResult(value map[string]interface{}, prefix string, result map[string]interface{}) {
	if prefix != "" {
		prefix += "."
	}
	for k, val := range value {
		if !strings.HasPrefix(k, prefix) {
			continue
		}

		key := k[len(prefix):]
		idx := strings.Index(key, ".")
		if idx != -1 {
			key = key[:idx]
		}
		if _, ok := result[key]; ok {
			continue
		}
		if idx == -1 {
			result[key] = val
			continue
		}

		// It contains a period, so it is a more complex structure
		result[key] = ExpandPrefixed(value, k[:len(prefix)+len(key)])
	}
}

func Flat(value interface{}) map[string]interface{} {
	return FlattenPrefixed(value, "")
}

func FlattenPrefixed(value interface{}, prefix string) map[string]interface{} {
	m := make(map[string]interface{})
	FlattenPrefixedToResult(value, prefix, m)
	return m
}

func FlattenPrefixedToResult(value interface{}, prefix string, m map[string]interface{}) {
	base := ""
	if prefix != "" {
		base = prefix + "."
	}

	original := reflect.ValueOf(value)
	kind := original.Kind()
	if kind == reflect.Ptr || kind == reflect.Interface {
		original = reflect.Indirect(original)
		kind = original.Kind()
	}
	t := original.Type()

	switch kind {
	// case reflect.Slice:
	// s := reflect.ValueOf(original)
	// for i := 0; i < s.Len(); i++ {
	// 	m[fmt.Sprintf("%v", i)] = FlattenPrefixedToResult(s.Index(i), base+fmt.Sprintf("%v", i), m)
	// }

	case reflect.Map:
		if t.Key().Kind() != reflect.String {
			break
		}
		for _, childKey := range original.MapKeys() {
			childValue := original.MapIndex(childKey)
			FlattenPrefixedToResult(childValue.Interface(), base+childKey.String(), m)
		}
	case reflect.Struct:
		for i := 0; i < original.NumField(); i += 1 {
			childValue := original.Field(i)
			childKey := t.Field(i).Name
			FlattenPrefixedToResult(childValue.Interface(), base+childKey, m)
		}
	default:
		if prefix != "" {
			m[prefix] = value
		}
	}
}
