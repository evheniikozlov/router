package router

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

const (
	intName, stringName, boolName, floatName, uintName, unusedName         = "int", "string", "bool", "float", "uint", "unused"
	intValue                                                       int64   = -123
	stringValue                                                            = "abc"
	boolValue                                                              = true
	floatValue                                                     float64 = 12.3
	uintValue                                                      uint64  = 123
)

func Test_NewParamsByRegexp(t *testing.T) {
	params := NewTestParams()
	if params[stringName] != stringValue {
		t.Errorf("%s is %s", stringName, params[stringName])
	}
	if params[intName] != strconv.FormatInt(intValue, 10) {
		t.Errorf("%s is %s", intName, params[intName])
	}
	if param, _ := strconv.ParseUint(params[uintName], 10, 64); param != uintValue {
		t.Errorf("%s is %d", uintName, param)
	}
	if param, _ := strconv.ParseBool(params[boolName]); param != boolValue {
		t.Errorf("%s is %t", boolName, param)
	}
	if param, _ := strconv.ParseFloat(params[floatName], 64); param != floatValue {
		t.Errorf("%s is %f", floatName, param)
	}
	if params[unusedName] != "" {
		t.Errorf("%s is %s", unusedName, params[unusedName])
	}
}

func Test_Params_GetString_ReturnCorrectParam_UsedName(t *testing.T) {
	params := NewTestParams()
	stringParam := params.GetString(stringName)
	if params.GetString(stringName) != stringValue {
		t.Errorf("%s is %s", stringName, stringParam)
	}
}

func Test_Params_GetString_ReturnEmptyString_UnusedName(t *testing.T) {
	params := NewTestParams()
	if params.GetString(unusedName) != "" {
		t.Error("Not an empty string, but the name is not used")
	}
}

func Test_Params_GetInt_ReturnCorrectParamAndNil_UsedName(t *testing.T) {
	params := NewTestParams()
	param, err := params.GetInt(intName)
	if param != intValue {
		t.Errorf("%s is %d", intName, param)
	}
	if err != nil {
		t.Errorf("error, %e", err)
	}
}

func Test_Params_GetInt_ReturnZeroAndError_UnusedName(t *testing.T) {
	params := NewTestParams()
	param, err := params.GetInt(unusedName)
	if param != 0 {
		t.Errorf("%s is not 0", unusedName)
	}
	if err == nil {
		t.Error("error is nil")
	}
}

func Test_Params_GetUint_ReturnCorrectParamAndNil_UsedName(t *testing.T) {
	params := NewTestParams()
	param, err := params.GetUint(uintName)
	if param != uintValue {
		t.Errorf("%s is %d", boolName, param)
	}
	if err != nil {
		t.Errorf("error, %e", err)
	}
}

func Test_Params_GetUint_ReturnZeroAndError_UnusedName(t *testing.T) {
	params := NewTestParams()
	param, err := params.GetUint(unusedName)
	if param != 0 {
		t.Errorf("%s is %d", unusedName, param)
	}
	if err == nil {
		t.Error("error is nil")
	}
}

func Test_Params_GetBool_ReturnCorrectParamAndNil_UsedName(t *testing.T) {
	params := NewTestParams()
	param, err := params.GetBool(boolName)
	if param != boolValue {
		t.Errorf("%s is %t", boolName, param)
	}
	if err != nil {
		t.Errorf("error, %e", err)
	}
}

func Test_Params_GetBool_ReturnsFalseAndError_UnusedName(t *testing.T) {
	params := NewTestParams()
	param, err := params.GetBool(unusedName)
	if param != false {
		t.Errorf("%s is %t", unusedName, param)
	}
	if err == nil {
		t.Error("error is nil")
	}
}

func Test_Params_GetFloat_ReturnCorrectParamAndNil_UsedName(t *testing.T) {
	params := NewTestParams()
	param, err := params.GetFloat(floatName)
	if param != floatValue {
		t.Errorf("%s is %f", boolName, param)
	}
	if err != nil {
		t.Errorf("error, %e", err)
	}
}

func Test_Params_GetFloat_ReturnZeroAndError_UnusedName(t *testing.T) {
	params := NewTestParams()
	param, err := params.GetFloat(unusedName)
	if param != 0 {
		t.Errorf("%s is %f", unusedName, param)
	}
	if err == nil {
		t.Error("error is nil")
	}
}

func NewTestParams() Params {
	return NewParamsByRegexp(fmt.Sprintf("%s/%d/%d/%t/%f", stringValue, intValue, uintValue, boolValue, floatValue), regexp.MustCompile(fmt.Sprintf("^(?P<%s>[a-zA-Z]*)/(?P<%s>[+\\-]?\\d*)/(?P<%s>\\d*)/(?P<%s>true|false)/(?P<%s>[+-]?([0-9]*[.])?[0-9]+)$", stringName, intName, uintName, boolName, floatName)))
}
