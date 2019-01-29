package hik_api

import (
	"fmt"
	"strconv"
	"strings"
)

func interfaceToString(src interface{}) string {
	str, _ := src.(string)
	return str
}

func interfaceToBool(src interface{}) bool {
	strBool := interfaceToString(src)
	return strings.ToLower(strBool) == "true"
}

func interfaceToInt(src interface{}) int {
	strNumber := interfaceToString(src)
	number, _ := strconv.Atoi(strNumber)
	return number
}

func interfaceToFloat64(src interface{}) float64 {
	strNumber := interfaceToString(src)
	number, _ := strconv.ParseFloat(strNumber, 64)
	return number
}

func intToString(src int) string {
	strInt := strconv.Itoa(src)
	return strInt
}

func float64ToString(src float64) string {
	strFloat64 := fmt.Sprint(src)
	return strFloat64
}

func boolToString(src bool) string {
	if src {
		return "true"
	}

	return "false"
}
