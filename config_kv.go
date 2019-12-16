// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package config

import "strings"

func getInterfaceKeyValue(configs map[string]interface{}, key string) (vm interface{}, err error) {

	tokens := strings.Split(key, ".")
	vm = configs[tokens[0]]
	for i, t := range tokens {
		if i == 0 {
			continue
		}
		v, ok := vm.(map[interface{}]interface{})
		if !ok {
			return nil, ErrNotMap
		}
		vm = v[t]
	}
	if vm == nil {
		err = ErrValueNil
	}
	return
}

// setInterfaceKeyValue set key value into *configs
func setInterfaceKeyValue(configs *map[string]interface{}, key string, value interface{}) (err error) {
	tokens := strings.Split(key, ".")
	for i := len(tokens) - 1; i >= 0; i-- {
		if i == 0 {
			(*configs)[tokens[0]] = value
			return
		}
		v, _ := getInterfaceKeyValue(*configs, strings.Join(tokens[:i], "."))
		vm, ok := v.(map[interface{}]interface{})
		if !ok {
			value = map[interface{}]interface{}{tokens[i]: value}
			continue
		}
		vm[tokens[i]] = value
		value = vm
	}
	return
}

func getStringKeyValue(configs map[string]interface{}, key string) (vm interface{}, err error) {

	tokens := strings.Split(key, ".")
	vm = configs[tokens[0]]
	for i, t := range tokens {
		if i != 0 {
			v, ok := vm.(map[string]interface{})
			if !ok {
				return nil, ErrNotMap
			}
			vm = v[t]
		}
	}

	if vm == nil {
		err = ErrValueNil
	}

	return
}

// setStringKeyValue set key value into configs
func setStringKeyValue(configs *map[string]interface{}, key string, value interface{}) (err error) {
	tokens := strings.Split(key, ".")
	for i := len(tokens) - 1; i >= 0; i-- {
		if i == 0 {
			(*configs)[tokens[i]] = value
			return
		}
		v, _ := getStringKeyValue(*configs, strings.Join(tokens[:i], "."))
		vm, ok := v.(map[string]interface{})
		if !ok {
			value = map[string]interface{}{tokens[i]: value}
			continue
		}
		vm[tokens[i]] = value
		value = vm
	}
	return
}