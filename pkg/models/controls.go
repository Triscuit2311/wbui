package models

import (
	"cmp"
	"log"
)

const (
	CheckboxControl int = iota
	IntSliderControl
	FloatSliderControl
	DisplayaValueControl
)

type Control interface {
	GetType() int
}

type SliderModel[T cmp.Ordered] struct {
	Min, Max, Val, Step T
	Label               string
}

func NewSliderModel[T int64 | float64](label string, val, min, max, step T) SliderModel[T] {
	return SliderModel[T]{
		Min: min, Max: max, Val: val, Step: step, Label: label,
	}
}

func (m SliderModel[T]) GetType() int {
	switch any(m.Val).(type) {
	case int64:
		return IntSliderControl
	case float64:
		return FloatSliderControl
	}
	log.Fatal("Invalid type for slidermodel")
	return -1
}

type CheckboxModel struct {
	Val   bool
	Label string
}

func NewCheckboxModel(label string, val bool) CheckboxModel {
	return CheckboxModel{
		Val: val, Label: label,
	}
}

func (m CheckboxModel) GetType() int {
	return CheckboxControl
}

type DisplayValueModel struct {
	Val   string
	Label string
}

func NewDisplayValueModel(label string, val string) DisplayValueModel {
	return DisplayValueModel{
		Val: val, Label: label,
	}
}

func (m DisplayValueModel) GetType() int {
	return DisplayaValueControl
}
