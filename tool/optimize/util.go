package main

import (
	"go/types"
	"math/rand"
	"path"
	"strings"
)

func GetPkgPathFromTypeStr(str string) string {
	valueType := strings.TrimPrefix(str, "*")
	ext := path.Ext(valueType)
	if len(ext) > 0 {
		return strings.TrimSuffix(valueType, ext)
	}
	return ""
}

func GetPkgPathFromType(ty types.Type) string {
	return GetPkgPathFromTypeStr(ty.String())
}

func GetNameFromTypeStr(str string) string {
	idx := strings.Index(str, "[")
	if idx < 0 {
		return str
	}
	return strings.TrimPrefix(str[0:idx], "*")
}

func GetNameFromType(ty types.Type) string {
	return GetNameFromTypeStr(ty.String())
}

func ResetTypePkgNameStr(str string, pkgName string) string {
	pkgPath := GetPkgPathFromTypeStr(str)
	if len(pkgName) <= 0 {
		pkgPath += "."
	}
	return strings.Replace(str, pkgPath, pkgName, 1)
}

func GetRandString(letters []rune, n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetRandPkgName() string {
	startLetters := []rune("abcdefghijklmnopqrstuvwxyz")
	otherLetters := []rune("abcdefghijklmnopqrstuvwxyz1234567890_")
	return GetRandString(startLetters, 1) + GetRandString(otherLetters, 7)
}
