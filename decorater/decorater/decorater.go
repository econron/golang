package decorater

import (
	"reflect"
)

type Controller struct {
	actions map[string]reflect.Value
}

func NewController() *Controller {
	return &Controller{actions: make(map[string]reflect.Value)}
}

func(c *Controller) Action(name string, fn interface{}) {
	fnValue := reflect.ValueOf(fn)
	if fnValue.Kind() != reflect.Func {
		panic("Only functions can be registered")
	}
	c.actions[name] = fnValue
}

func (c *Controller) Execute(name string, params ...interface{}) []interface{} {
	fn, exists := c.actions[name]
	if !exists {
		return []interface{}{"Action not found"}
	}

	// パラメータをreflect.Valueに変換
	inputs := make([]reflect.Value, len(params))
	for i, param := range params {
		inputs[i] = reflect.ValueOf(param)
	}

	results := fn.Call(inputs)

	outputs := make([]interface{}, len(results))
	for i, result := range results {
		outputs[i] = result.Interface()
	}

	return outputs
}

