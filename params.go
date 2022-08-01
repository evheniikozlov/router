package router

import (
	"regexp"
	"strconv"
)

type Params map[string]string

func NewParamsByRegexp(source string, sourceRegexp *regexp.Regexp) Params {
	params := make(map[string]string)
	matches := sourceRegexp.FindStringSubmatch(source)
	for index, name := range sourceRegexp.SubexpNames() {
		if index != 0 && name != "" {
			params[name] = matches[index]
		}
	}
	return params
}

func (params Params) GetString(name string) string {
	return params[name]
}

func (params Params) GetInt(name string) (int64, error) {
	return strconv.ParseInt(params[name], 10, 64)
}

func (params Params) GetUint(name string) (uint64, error) {
	return strconv.ParseUint(params[name], 10, 64)
}

func (params Params) GetBool(name string) (bool, error) {
	return strconv.ParseBool(params[name])
}

func (params Params) GetFloat(name string) (float64, error) {
	return strconv.ParseFloat(params[name], 64)
}
