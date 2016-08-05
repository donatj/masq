package main

import (
	"errors"
	"strconv"
)

var (
	errUnexpectedStringSuffix = errors.New("unexpected strIntSuffix string")
)

func strIntSuffixSplit(s string) (string, int, error) {
	ostr := ""
	osint := ""

	intStart := false
	for _, r := range s {
		isD := isDigit(r)

		if !isD && !intStart {

			ostr += string(r)
		} else if isD {
			intStart = true
			osint += string(r)
		} else {
			return "", 0, errUnexpectedStringSuffix
		}
	}

	oint, err := strconv.Atoi(osint)
	if err != nil {
		return "", 0, err
	}

	return ostr, oint, nil
}

func reverseStr(str string) (out string) {
	for _, s := range str {
		out = string(s) + out
	}

	return
}
