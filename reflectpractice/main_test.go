package main

import (
	"reflect"
	"testing"
)

type user struct {
	UserId string
	Name string
}

func TestReflectStructPrt(t *testing.T){
	var (
		model *user
		st reflect.Type
		elem reflect.Value
	)

	st = reflect.TypeOf(model)
	t.Log("reflect.TypeOf", st.Kind().String())
	st = st.Elem()
	t.Log("reflect.TypeOf.Elem", st.Kind().String())
	elem = reflect.New(st)

	t.Log("reflect.New", elem.Kind().String())
	t.Log("reflect.New.Elem", elem.Elem().Kind().String())

	model = elem.Interface().(*user)
	elem = elem.Elem()
	elem.FieldByName("UserId").SetString("Cool123")
	elem.FieldByName("Name").SetString("Vincennt")
	t.Log("model model.Name", model, model.Name)
}
