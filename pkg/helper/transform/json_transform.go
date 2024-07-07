package transform

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	// TransformEmpty is the tag for empty fields
	// it allows the empty value of any type independent of the target type
	// this is especially used because the fritzbox api will sometimes return empty array for object types
	TransformExtendedEmpty = "extendedEmpty"
)

func transformExtendedEmpty(value interface{}, target reflect.Value) interface{} {
	isEmpty := reflect.ValueOf(value).IsZero()
	// switch valueType := value.(type) {
	// case map[string]interface{}:
	// 	isEmpty = len(valueType) == 0
	// case []interface{}:
	// 	isEmpty = len(valueType) == 0
	// default:
	// 	isEmpty = value == reflect.Zero(reflect.TypeOf(value)).Interface()
	// }

	if isEmpty {
		return reflect.Zero(target.Type()).Interface()
	}
	return value
}

var transformersMap = map[string]func(interface{}, reflect.Value) interface{}{
	TransformExtendedEmpty: transformExtendedEmpty,
}

func MapToStruct[T any](raw map[string]interface{}, target T) T {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr {
		targetValue = reflect.ValueOf(&target)
	}

	for key, value := range raw {
		field := findJsonField(targetValue, key)
		if !field.IsValid() || !field.CanSet() {
			fmt.Println("field not found", key, value, field, reflect.ValueOf(target).Type())
			continue
		}

		transformers := findTransformer(targetValue, key)
		transformExtendedEmpty(value, field)
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
				elementValue := reflect.New(elementType)
				slice := reflect.MakeSlice(reflect.SliceOf(elementType), 0, sliceLen)

				switch elementType.Kind() {
				case reflect.Struct:
					for i := 0; i < sliceLen; i++ {
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
		default:
			// try to set field
			if reflect.TypeOf(value).AssignableTo(field.Type()) {
				field.Set(reflect.ValueOf(value))
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
