package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

//Tools ----------------------------------------------------------
type Tools struct {
	getStrAry toolsGetStrAry
	getStr    toolsGetStr
	getInt    toolsGetInt
}

var tools Tools

//getInt
type toolsGetInt struct{}

//getInt.fromStr
func (*toolsGetInt) fromStr(fromStr string) (resInt int, err error) {
	resInt, err = strconv.Atoi(fromStr)
	if err != nil {
		logger.Printf(err.Error())
		return 0, err
	}
	return
}

//getStr
type toolsGetStr struct{}

//getStr.fromInt
func (*toolsGetStr) fromInt(fromInt int) (resStr string) {
	return strconv.Itoa(fromInt)
}

//getStr.regex
func (*toolsGetStr) regex(fromStr *string, regexStr string) (resStr string, err error) {
	rgx := regexp.MustCompile(regexStr)
	resStr = rgx.FindString(*fromStr)
	if resStr == "" {
		err = errors.New("error : getStr.regex rgx.FindString has no regex match")
		return "", err
	}
	return resStr, err
}

//getStr.between
func (*toolsGetStr) between(originalStr string, startStr string, endStr string) (resStr string, err error) {
	startIdx := strings.Index(originalStr, startStr)
	if startIdx == -1 {
		err = errors.New("error : function getStr.between startStr " + startStr + " is not in originalStr " + originalStr)
		logger.Println(err.Error())
		return "", err
	}
	startIdx += len(startStr)
	endIdx := strings.Index(originalStr[startIdx:], endStr)
	if endIdx == -1 {
		err = errors.New("error : function getStr.between endStr " + endStr + " is not in originalStr " + originalStr)
		logger.Println(err.Error())
		return "", err
	}
	return originalStr[startIdx : startIdx+endIdx], err
}

//getStr.after
func (*toolsGetStr) after(originalStr string, afterStr string) (resStr string, err error) {
	// Get substring after a string.
	afterStrEndIdx := strings.LastIndex(originalStr, afterStr)
	if afterStrEndIdx == -1 {
		err = errors.New("error : function getStr.after afterStr " + afterStr + " is not in originalStr " + originalStr)
		return "", err
	}
	adjustedIdx := afterStrEndIdx + len(afterStr)
	if adjustedIdx >= len(originalStr) {
		err = errors.New("error : function getStr.after afterStrEndIdx+len(afterStr)>=len(originalStr)")
		return "", err
	}
	return originalStr[adjustedIdx:len(originalStr)], err
}

//getStr.before
func (toolsGetStr) before(originalStr string, beforeStr string) (resStr string, err error) {
	pos := strings.Index(originalStr, beforeStr)
	if pos == -1 {
		err = errors.New("error : function getStr.before beforeStr " + beforeStr + " is not in originalStr " + originalStr)
		return "", err
	}
	return originalStr[0:pos], err
}

//getStrAry
type toolsGetStrAry struct{}

//getStrAry.regex
func (*toolsGetStrAry) regex(fromStr *string, regexStr string) (resStrAry []string, err error) {
	rgx := regexp.MustCompile(regexStr)
	resStrAry = rgx.FindAllString(*fromStr, -1)
	if len(resStrAry) == 0 {
		err = errors.New("error : getStrAry.regex rgx.FindAllString has no regex match")
		return nil, err
	}
	return resStrAry, err
}

//getStrAry.noDuplicate
func (*toolsGetStrAry) noDuplicate(strAry *[]string) (res []string) {
	noDuplicate := make(map[string]bool)
	for _, str := range *strAry {
		if _, wasInserted := noDuplicate[str]; !wasInserted {
			noDuplicate[str] = true
			res = append(res, str)
		}
	}
	return res
}

//getStrAry.combind
func (*toolsGetStrAry) combind(strAry1 []string, strAry2 []string) (res []string) {
	for _, str := range strAry1 {
		res = append(res, str)
	}
	for _, str := range strAry2 {
		res = append(res, str)
	}
	return res
}
