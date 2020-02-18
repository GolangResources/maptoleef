package maptoleef

import (
	"fmt"
	"sort"
	"strconv"
)

type LEEF struct {
	Manufacter	string
	ProductName	string
	ProductVersion	string
	EventID		string
	Delimiter	string
}

func (l *LEEF) MapToLEEF(message map[string]interface{}) string {
	if l.Delimiter == "" {
		l.Delimiter = "0x09"
	}
	leefString := fmt.Sprintf("LEEF:2.0|%v|%v|%v|%v|%v|", l.Manufacter, l.ProductName, l.ProductVersion, l.EventID, l.Delimiter)
	keys := make([]string, 0, len(message))
	for key := range message {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _ , key := range keys {
		value := message[key]
		switch t := value.(type) {
			case string:
				leefString += fmt.Sprintf("%s=%s\t", key, value)
			case int:
				leefString += fmt.Sprintf("%s=%s\t", key, value)
			case float64:
				leefString += fmt.Sprintf("%s=%s\t", key, value)
			case uint32:
				leefString += fmt.Sprintf("%s=%s\t", key, strconv.FormatUint(uint64(value.(uint32)), 10))
			case uint64:
				leefString += fmt.Sprintf("%s=%s\t", key, strconv.FormatUint(value.(uint64), 10))
			default:
				leefString += fmt.Sprintf("%s=%s\t", key, value)
				_ = t
		}
	}
	return leefString
}
