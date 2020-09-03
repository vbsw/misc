/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package properties provides functions to read/write properties.
package properties

import (
	"github.com/vbsw/misc/parser/propertiesparser"
	"io"
	"io/ioutil"
	"os"
)

// ReadFile reads properties from file. The file must be in UTF-8.
func ReadFile(path string) (map[string]string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err == nil {
		return ReadBytes(bytes), err
	}
	return nil, err
}

// ReadBytes is the same as Read, but reads from byte array.
func ReadBytes(bytes []byte) map[string]string {
	var parser propertiesparser.LineParser
	var name string
	props := make(map[string]string)
	buffer := make([]byte, 32)
	for i := 0; i < len(bytes); {
		i = parser.ParseLine(bytes, i)
		if parser.LineType == propertiesparser.LTProperty {
			buffer = parser.PropertyName(bytes, buffer[:0])
			name = string(buffer)
			buffer = parser.PropertyValue(bytes, buffer[:0])
			props[name] = string(buffer)
			name = ""
		} else if parser.LineType == propertiesparser.LTPropertyNext {
			buffer = parser.PropertyName(bytes, buffer[:0])
			name = string(buffer)
			buffer = parser.PropertyValue(bytes, buffer[:0])
		} else if parser.LineType == propertiesparser.LTPropertyContNext {
			buffer = parser.PropertyValue(bytes, buffer)
		} else if parser.LineType == propertiesparser.LTPropertyCont {
			buffer = parser.PropertyValue(bytes, buffer)
			props[name] = string(buffer)
			name = ""
		} else if parser.LineType == propertiesparser.LTUnknownFormat && len(name) > 0 {
			props[name] = string(buffer)
			name = ""
		}
	}
	return props
}

// WriteFile writes key value pairs from props to path.
func WriteFile(path string, props map[string]string) error {
	if len(props) > 0 {
		out, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)

		if err == nil {
			defer out.Close()
			bytes := propsToBytes(props)
			_, err = out.Write(bytes)

			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}
