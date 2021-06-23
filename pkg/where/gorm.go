package where

import (
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"reflect"
)

var Resiter []ResiterHooksOption

type HooksOption struct {
	Field       string
	Default     string
	Where       string
	StructField reflect.StructField
	Value       reflect.Value
	DB          *gorm.DB
}
type ResiterHooksOption struct {
	Hooks  Hooks
	Action HooksWhere
	Where  string
}
type Hooks func(option *HooksOption) *gorm.DB
type HooksWhere func(option *HooksOption) bool

var whereIn GromInterface

type Grom struct {
	Hooks      []Hooks
	HooksList  map[string]Hooks
	WhereList  []string
	HooksWhere map[string]HooksWhere
}
type GromInterface interface {
	ResiterHooks(where string, hk Hooks)
	ResiterAction(where string, hk HooksWhere)
	Where(db *gorm.DB, i interface{}) *gorm.DB
}

/*
*@Author Administrator
*@Date 21/4/2021 13:22
*@desc
 */
func (g *Grom) ResiterHooks(where string, hk Hooks) {
	g.HooksList[where] = hk
}
func (g *Grom) ResiterAction(where string, hk HooksWhere) {
	g.HooksWhere[where] = hk
}
func (g *Grom) SeparatorDefualt(option *HooksOption) bool {
	toString := cast.ToString(option.Value.Interface())
	if option.Default != toString {
		return true
	}
	return false
}

func (g *Grom) WhereAction(option *HooksOption) bool {
	if fn, ok := g.HooksWhere[option.Where]; ok {
		if fn(option) {
			return true
		}
	} else {
		if g.SeparatorDefualt(option) {
			return true
		}
	}
	return false
}
func (g *Grom) Where(db *gorm.DB, i interface{}) *gorm.DB {
	refType := reflect.TypeOf(i)
	refValue := reflect.ValueOf(i)
	for i := 0; i < refValue.NumField(); i++ {
		f := refType.Field(i)
		valueInterface := refValue.Field(i)
		if &valueInterface == nil {
			continue
		}
		if valueInterface.Kind() == reflect.Ptr {
			continue
		}
		if valueInterface.Kind() == reflect.Struct && valueInterface.Interface() != nil {
			db = g.Where(db, valueInterface.Interface())
		}

		value := valueInterface.Interface()
		if f.Tag == "" {
			continue
		}
		if value == nil {
			continue
		}
		field := f.Tag.Get("field")
		defaults := f.Tag.Get("default")
		where := f.Tag.Get("where")
		if field == "" || where == "" {
			continue
		}
		option := HooksOption{
			Field:       field,
			Default:     defaults,
			Where:       where,
			StructField: f,
			Value:       valueInterface,
			DB:          db,
		}
		if g.WhereAction(&option) {
			db = g.WhereSeparator(&option)
		}
	}
	return db
}

func (g *Grom) WhereSeparator(option *HooksOption) *gorm.DB {
	if fn, ok := g.HooksList[option.Where]; ok {
		option.DB = fn(option)
	}
	return option.DB
}

func NewGorm() GromInterface {
	if whereIn != nil {
		return whereIn
	}
	g := new(Grom)
	g.HooksWhere = map[string]HooksWhere{}
	g.HooksList = NewAction()
	options := Resiter
	if len(options) > 0 {
		for _, option := range options {
			if option.Hooks != nil {
				g.ResiterHooks(option.Where, option.Hooks)
			}
			if option.Action != nil {
				g.ResiterAction(option.Where, option.Action)
			}
		}
	}
	whereIn = g
	return g
}
func RegsterHooks(options ...ResiterHooksOption) {
	if options != nil {
		Resiter = append(Resiter, options...)
	}
}
