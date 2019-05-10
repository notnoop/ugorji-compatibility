package main

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

const generateNilSlicesAndMaps = true

func fakeStruct(a interface{}) error {

	reflectType := reflect.TypeOf(a)

	if reflectType.Kind() != reflect.Ptr {
		return errors.New("not pointer")
	}

	if reflect.ValueOf(a).IsNil() {
		return fmt.Errorf("nil pointer")
	}

	rval := reflect.ValueOf(a)

	finalValue, err := getValue(a, true)
	if err != nil {
		return err
	}

	rval.Elem().Set(finalValue.Elem().Convert(reflectType.Elem()))
	return nil
}

func getValue(a interface{}, top bool) (reflect.Value, error) {
	t := reflect.TypeOf(a)
	if t == nil {
		return reflect.Value{}, fmt.Errorf("interface{} not allowed")
	}
	k := t.Kind()

	switch k {
	case reflect.Ptr:
		if !top && rand.Intn(2) > 0 {
			return reflect.Zero(t), nil
		}

		v := reflect.New(t.Elem())
		var val reflect.Value
		var err error
		if a != reflect.Zero(reflect.TypeOf(a)).Interface() {
			val, err = getValue(reflect.ValueOf(a).Elem().Interface(), false)
			if err != nil {
				return reflect.Value{}, err
			}
		} else {
			val, err = getValue(v.Elem().Interface(), false)
			if err != nil {
				return reflect.Value{}, err
			}
		}
		v.Elem().Set(val.Convert(t.Elem()))
		return v, nil
	case reflect.Struct:

		switch t.String() {
		case "time.Time":
			ft := time.Now().Add(time.Duration(rand.Int63()))
			ft = ft.Round(time.Millisecond)
			return reflect.ValueOf(ft), nil
		default:
			v := reflect.New(t).Elem()
			for i := 0; i < v.NumField(); i++ {
				if !v.Field(i).CanSet() {
					continue // to avoid panic to set on unexported field in struct
				}

				if t.PkgPath() == "github.com/hashicorp/nomad/nomad/structs" &&
					t.Name() == "WriteMeta" &&
					t.Field(i).Name == "Index" {
					continue
				}
				if t.PkgPath() == "github.com/hashicorp/nomad/nomad/structs" &&
					t.Name() == "QueryOptions" &&
					t.Field(i).Name == "Prefix" {
					continue
				}

				val, err := getValue(v.Field(i).Interface(), false)
				if err != nil {
					return reflect.Value{}, err
				}
				val = val.Convert(v.Field(i).Type())
				v.Field(i).Set(val)

			}
			return v, nil
		}

	case reflect.String:
		res := randomString()
		return reflect.ValueOf(res), nil
	case reflect.Array, reflect.Slice:
		len := randomSliceAndMapSize()
		if generateNilSlicesAndMaps && len == 0 && rand.Intn(2) == 0 {
			return reflect.Zero(t), nil
		}
		v := reflect.MakeSlice(t, len, len)
		for i := 0; i < v.Len(); i++ {
			val, err := getValue(v.Index(i).Interface(), false)
			if err != nil {
				return reflect.Value{}, err
			}
			val = val.Convert(t.Elem())
			v.Index(i).Set(val)
		}
		return v, nil
	case reflect.Int:
		return reflect.ValueOf(randomInteger()), nil
	case reflect.Int8:
		return reflect.ValueOf(int8(randomInteger())), nil
	case reflect.Int16:
		return reflect.ValueOf(int16(randomInteger())), nil
	case reflect.Int32:
		return reflect.ValueOf(int32(randomInteger())), nil
	case reflect.Int64:
		return reflect.ValueOf(int64(randomInteger())), nil
	case reflect.Float32:
		return reflect.ValueOf(rand.Float32()), nil
	case reflect.Float64:
		return reflect.ValueOf(rand.Float64()), nil
	case reflect.Bool:
		val := rand.Intn(2) > 0
		return reflect.ValueOf(val), nil

	case reflect.Uint:
		return reflect.ValueOf(uint(randomInteger())), nil

	case reflect.Uint8:
		return reflect.ValueOf(uint8(randomInteger())), nil

	case reflect.Uint16:
		return reflect.ValueOf(uint16(randomInteger())), nil

	case reflect.Uint32:
		return reflect.ValueOf(uint32(randomInteger())), nil

	case reflect.Uint64:
		return reflect.ValueOf(uint64(randomInteger())), nil

	case reflect.Map:
		len := randomSliceAndMapSize()
		if generateNilSlicesAndMaps && len == 0 && rand.Intn(2) == 0 {
			return reflect.Zero(t), nil
		}
		v := reflect.MakeMap(t)
		for i := 0; i < len; i++ {
			keyInstance := reflect.New(t.Key()).Elem().Interface()
			key, err := getValue(keyInstance, false)
			if err != nil {
				return reflect.Value{}, err
			}
			key = key.Convert(t.Key())

			valueInstance := reflect.New(t.Elem()).Elem().Interface()
			val, err := getValue(valueInstance, false)
			if err != nil {
				return reflect.Value{}, err
			}
			val = val.Convert(t.Elem())
			v.SetMapIndex(key, val)
		}
		return v, nil
	default:
		err := fmt.Errorf("no support for kind %+v", t)
		return reflect.Value{}, err
	}

}

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func randomString() string {
	n := rand.Intn(10)
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func randomInteger() int {
	return rand.Intn(100)
}

func randomSliceAndMapSize() int {
	return rand.Intn(2)
}

func randomElementFromSliceString(s []string) string {
	return s[rand.Int()%len(s)]
}
