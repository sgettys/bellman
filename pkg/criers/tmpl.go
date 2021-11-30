package criers

import (
	"bytes"
	"encoding/json"
	"text/template"

	"github.com/Masterminds/sprig"
)

func GetString(data []byte, text string) (string, error) {
	tmpl, err := template.New("template").Funcs(sprig.TxtFuncMap()).Parse(text)
	if err != nil {
		return "", nil
	}

	buf := new(bytes.Buffer)
	// TODO: Should we send event directly or more events?
	err = tmpl.Execute(buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func convertLayoutTemplate(layout map[string]interface{}, data []byte) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	for key, value := range layout {
		m, err := convertTemplate(value, data)
		if err != nil {
			return nil, err
		}
		result[key] = m
	}
	return result, nil
}

func convertTemplate(value interface{}, data []byte) (interface{}, error) {
	switch v := value.(type) {
	case string:
		rendered, err := GetString(data, v)
		if err != nil {
			return nil, err
		}

		return rendered, nil
	case map[interface{}]interface{}:
		strKeysMap := make(map[string]interface{})
		for k, v := range v {
			res, err := convertTemplate(v, data)
			if err != nil {
				return nil, err
			}
			// TODO: It's a bit dangerous
			strKeysMap[k.(string)] = res
		}
		return strKeysMap, nil
	case map[string]interface{}:
		strKeysMap := make(map[string]interface{})
		for k, v := range v {
			res, err := convertTemplate(v, data)
			if err != nil {
				return nil, err
			}
			strKeysMap[k] = res
		}
		return strKeysMap, nil
	case []interface{}:
		listConf := make([]interface{}, len(v))
		for i := range v {
			t, err := convertTemplate(v[i], data)
			if err != nil {
				return nil, err
			}
			listConf[i] = t
		}
		return listConf, nil
	}
	return nil, nil
}

func serializeEventWithLayout(layout map[string]interface{}, data []byte) ([]byte, error) {
	var toSend []byte
	if layout != nil {
		res, err := convertLayoutTemplate(layout, data)
		if err != nil {
			return nil, err
		}

		toSend, err = json.Marshal(res)
		if err != nil {
			return nil, err
		}
	}
	return toSend, nil
}
