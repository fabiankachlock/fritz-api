package transform

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	// TransformEmpty is the tag for empty fields
	// it allows the empty value of any type independent of the target type
	// this is especially used because the fritzbox api will sometimes return empty array for object types
	TransformExtendedEmpty = "extendedEmpty"
	TransformStringToInt   = "stringToInt"
	TransformToSlice       = "toSlice"
)

func transformExtendedEmpty(value interface{}, target reflect.Value) interface{} {
	if reflect.ValueOf(value).IsZero() {
		return reflect.Zero(target.Type()).Interface()
	}
	return value
}

func transformStringToInt(value interface{}, target reflect.Value) interface{} {
	if reflect.ValueOf(value).IsZero() {
		return 0
	}
	if reflect.ValueOf(value).Kind() == reflect.String {
		if intValue, err := strconv.Atoi(value.(string)); err == nil {
			return intValue
		}
	}
	return value
}

func transformToSlice(value interface{}, target reflect.Value) interface{} {
	if reflect.TypeOf(value).Kind() != reflect.Slice {
		slice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(value)), 0, 0)
		slice = reflect.Append(slice, reflect.ValueOf(value))
		return slice.Interface()
	}
	return value
}

var transformersMap = map[string]func(interface{}, reflect.Value) interface{}{
	TransformExtendedEmpty: transformExtendedEmpty,
	TransformStringToInt:   transformStringToInt,
	TransformToSlice:       transformToSlice,
}

func MapToStruct[T any](raw map[string]interface{}, target T) T {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr {
		targetValue = reflect.ValueOf(&target)
	}

	for key, value := range raw {
		field := findJsonField(targetValue, key)
		if !field.IsValid() || !field.CanSet() {
			// fmt.Println("field not found", key, value, field, reflect.ValueOf(target).Type())
			continue
		}

		transformers := findTransformer(targetValue, key)
		for _, transformerKey := range transformers {
			transformer, ok := transformersMap[transformerKey]
			if ok {
				value = transformer(value, field)
			}
		}

		switch valueType := value.(type) {
		case map[string]interface{}:
			if field.Kind() == reflect.Struct {
				// recurse
				field.Set(reflect.ValueOf(MapToStruct(valueType, field.Addr().Interface())).Elem())
			} else if field.Kind() == reflect.Map {
				// convert map
				// elementType := field.Type()
				keyType := field.Type().Key()
				elementType := field.Type().Elem()
				elementValue := reflect.New(elementType)
				mapValue := reflect.MakeMap(reflect.MapOf(keyType, elementType))

				switch elementType.Kind() {
				case reflect.Struct:
					for key, value := range valueType {
						mappedStruct := MapToStruct(value.(map[string]interface{}), elementValue.Interface())
						mapValue.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(mappedStruct).Elem())
					}
				default:
					for key, value := range valueType {
						mapValue.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
					}
				}

				field.Set(mapValue)
			}
		case []interface{}:
			// convert slice
			if field.Kind() == reflect.Slice {
				sliceLen := len(valueType)
				elementType := field.Type().Elem()
				slice := reflect.MakeSlice(reflect.SliceOf(elementType), 0, sliceLen)
				fmt.Println("slice", key, elementType.Kind(), field.Type().Elem().Kind())
				switch elementType.Kind() {
				case reflect.Struct:
					for i := 0; i < sliceLen; i++ {
						elementValue := reflect.New(elementType)
						mappedStruct := MapToStruct(valueType[i].(map[string]interface{}), elementValue.Interface())
						slice = reflect.Append(slice, reflect.ValueOf(mappedStruct).Elem())
					}
				default:
					for i := 0; i < sliceLen; i++ {
						slice = reflect.Append(slice, reflect.ValueOf(valueType[i]))
					}
				}

				field.Set(slice)
			}
		case float64:
			switch field.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				field.SetInt(int64(valueType))
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				field.SetUint(uint64(valueType))
			case reflect.Float32, reflect.Float64:
				field.SetFloat(valueType)
			}
		default:
			if reflect.TypeOf(value).AssignableTo(field.Type()) {
				field.Set(reflect.ValueOf(value))
			} else if reflect.ValueOf(value).CanConvert(field.Type()) {
				field.Set(reflect.ValueOf(value).Convert(field.Type()))
			}
		}
	}
	return target
}

func findJsonField(value reflect.Value, key string) reflect.Value {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() == reflect.Interface {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return reflect.Value{}
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := value.Type().Field(i)

		tagDefinition := fieldType.Tag.Get("json")
		if key == tagDefinition || (strings.HasPrefix(tagDefinition, key) && strings.Contains(tagDefinition, ",")) {
			return field
		}
	}

	return reflect.Value{}
}

func findTransformer(value reflect.Value, key string) []string {
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return []string{}
	}

	for i := 0; i < value.NumField(); i++ {
		fieldType := value.Type().Field(i)

		tagDefinition := fieldType.Tag.Get("json")
		if key == tagDefinition || (strings.HasPrefix(tagDefinition, key) && strings.Contains(tagDefinition, ",")) {
			return strings.Split(fieldType.Tag.Get("transform"), ",")
		}
	}

	return []string{}
}
