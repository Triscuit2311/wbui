package main

import (
	"fmt"
	"github.com/a-h/templ"
	"net/http"
	model "wbui/pkg/models"
	ui "wbui/pkg/ui"
)

func main() {

	store := model.NewStore(10)

	var v1 string = "ll"
	var v2 int = 123
	var v3 float64 = 12.3456
	var v4 bool = false

	k1 := model.SetKV(&store, v1)
	k2 := model.SetKV(&store, v2)
	k3 := model.SetKV(&store, v3)
	k4 := model.SetKV(&store, v4)

	fmt.Printf("v1->k1->val: %v\n", model.GetKV[string](&store, k1))
	fmt.Printf("v2->k2->val: %v\n", model.GetKV[int](&store, k2))
	fmt.Printf("v3->k3->val: %v\n", model.GetKV[float64](&store, k3))
	fmt.Printf("v4->k4->val: %v\n", model.GetKV[bool](&store, k4))

	return

	comp := ui.Index()

	http.Handle("/", templ.Handler(comp))

	http.ListenAndServe(":3000", nil)
}
