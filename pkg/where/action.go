package where

import (
	"gorm.io/gorm"
	"strings"
)

var whereMap = map[string]map[string]string{
	"eq":      {"action": "Common", "char": "="},
	"neq":     {"action": "Common", "char": "!= "},
	"gt":      {"action": "Common", "char": " > "},
	"egt":     {"action": "Common", "char": " >= "},
	"lt":      {"action": "Common", "char": " < "},
	"elt":     {"action": "Common", "char": " <= "},
	"null":    {"action": "ContainNull", "char": "is null "},
	"notnull": {"action": "ContainNull", "char": " is not null "},
	"like":    {"action": "Like", "char": "  "},
	"in":      {"action": "ContainIn", "char": " in "},
	"notin":   {"action": "ContainIn", "char": " not in "},
	"or":      {"action": "ContainIn", "char": ""},
}

type Action struct {
}

func (a *Action) Common(char string) Hooks {
	return func(option *HooksOption) *gorm.DB {
		return option.DB.Where(option.Field+" "+char+" ? ", option.Value.Interface())
	}
}
func (a *Action) ContainOr() Hooks {
	return func(option *HooksOption) *gorm.DB {
		return option.DB.Or(option.Value.Interface())
	}
}
func (a *Action) ContainNull(char string) Hooks {
	return func(option *HooksOption) *gorm.DB {
		return option.DB.Where(option.Field + " " + char)
	}
}
func (a *Action) ContainIn(char string) Hooks {
	return func(option *HooksOption) *gorm.DB {
		split := strings.Split(option.Value.String(), ",")
		return option.DB.Where(option.Field+" "+char+" (?) ", split)
	}
}

func (a *Action) Like() Hooks {
	return func(option *HooksOption) *gorm.DB {
		return option.DB.Where(option.Field+" like ?", "%"+option.Value.String()+"%")
	}
}

func (a *Action) Init() map[string]Hooks {
	m := make(map[string]Hooks)
	for key, val := range whereMap {
		switch val["action"] {
		case "Common":
			m[key] = a.Common(val["char"])
		case "ContainNull":
			m[key] = a.ContainNull(val["char"])
		case "Like":
			m[key] = a.Like()
		case "ContainIn":
			m[key] = a.ContainIn(val["char"])
		case "ContainOr":
			m[key] = a.ContainOr()
		}
	}
	return m
}

/*
*@Author Administrator
*@Date 22/4/2021 10:06
*@desc
 */
func NewAction() map[string]Hooks {
	return new(Action).Init()
}
