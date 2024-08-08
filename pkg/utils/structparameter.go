//结构为参数
//
//struct 中存在 kvname tag
//
//map中对应的tag值即为对应tag字段的值
//
//例如 struct {A int `kvname:"a"`}
//
// map[string]string{"a" 1}
//
//解析之后， A的值未 1
package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// 将 key-value 形式的键值对 解析到对象中
//
//kv 键值对
//
//bind 解析到对象，必须是指针类型
//
//结构化参数
//
//struct 中存在 kvname tag
//
//map中对应的tag值即为对应tag字段的值
//
//例如 struct {A int `kvname:"a"`}
//
// map[string]string{"a" 1}
//
//解析之后， A的值为 1
func KvToStructParameter(kv map[string]string, bind interface{}) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%+v\b", err)
		}
	}()
	if kv == nil {
		return errors.New("kv is nil")
	}

	refType := reflect.TypeOf(bind)
	if refType.Kind() != reflect.Pointer {
		return fmt.Errorf("bind should  pointer,but get %v", refType.Kind())
	}
	refValue := reflect.ValueOf(bind)
	if refValue.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("bind should be struct pointer,but get %v", refValue.Kind())
	}

	refElem := refValue.Elem()
	for i := 0; i < refElem.NumField(); i++ {
		fieldType := refElem.Type().Field(i)
		tag := fieldType.Tag.Get("kvname")
		if tag == "" {
			continue
		}
		value, ok := kv[tag]
		if !ok {
			continue
		}
		fieldValue := refElem.Field(i)
		if !fieldValue.CanSet() {
			continue
		}

		switch fieldType.Type.Kind() {
		case reflect.Bool:
			isTrue := value != "0" && strings.ToUpper(value) != "FALSE"
			fieldValue.SetBool(isTrue)
		case reflect.Int:
			svalue, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				return err
			}
			fieldValue.SetInt(svalue)
		case reflect.Int8:
			svalue, err := strconv.ParseInt(value, 10, 8)
			if err != nil {
				return err
			}
			fieldValue.SetInt(svalue)
		case reflect.Int16:
			svalue, err := strconv.ParseInt(value, 10, 16)
			if err != nil {
				return err
			}
			fieldValue.SetInt(svalue)
		case reflect.Int32:
			svalue, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				return err
			}
			fieldValue.SetInt(svalue)
		case reflect.Int64:
			svalue, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
			fieldValue.SetInt(svalue)
		case reflect.Uint:
			svalue, err := strconv.ParseUint(value, 10, 32)
			if err != nil {
				return err
			}
			fieldValue.SetUint(svalue)
		case reflect.Uint8:
			svalue, err := strconv.ParseUint(value, 10, 8)
			if err != nil {
				return err
			}
			fieldValue.SetUint(svalue)
		case reflect.Uint16:
			svalue, err := strconv.ParseUint(value, 10, 16)
			if err != nil {
				return err
			}
			fieldValue.SetUint(svalue)
		case reflect.Uint32:
			svalue, err := strconv.ParseUint(value, 10, 32)
			if err != nil {
				return err
			}
			fieldValue.SetUint(svalue)
		case reflect.Uint64:
			svalue, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return err
			}
			fieldValue.SetUint(svalue)
		case reflect.Float32:
			svalue, err := strconv.ParseFloat(value, 32)
			if err != nil {
				return err
			}
			fieldValue.SetFloat(svalue)
		case reflect.Float64:
			svalue, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			fieldValue.SetFloat(svalue)
		case reflect.String:
			fieldValue.SetString(value)
		default:
			return fmt.Errorf("unsupport type  %v in filed %s", fieldType.Type.Kind(), tag)

		}
	}
	return nil
}
