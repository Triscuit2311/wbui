package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	data "wbui/pkg/data"
	"wbui/pkg/models"

	//	"wbui/pkg/models"
	ui "wbui/pkg/ui"

	"github.com/a-h/templ"
)

// TODO:
// Store control IDS
// Store control DATA & Control Type
// Lookup data via ID
// determine type of control from stored control data
// render appropriate UI control to buffer
// ship it back with the same ID as originating request head

const (
	prealloc = 10
)

var (
	dataStore  = data.NewStore(prealloc)
	controlMap = make(map[int]int, prealloc) // Store control type for cross-ref
)

func checkValuesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r.Header.Get(\"ID\"): %v\n", r.Header.Get("ID"))
}

func AddSlider[T int64 | float64](label string, val, min, max, step T) {
	model := models.NewSliderModel[T](label, val, min, max, step)
	ID := data.SetKV[models.SliderModel[T]](&dataStore, model)
	controlMap[ID] = model.GetType()
}

func initControls(w http.ResponseWriter, r *http.Request) {
	fmt.Println("initControls Called")
	var buf bytes.Buffer

	for id := 0; id <= len(controlMap); id++ {
		t, ok := controlMap[id]
		if !ok {
			continue
		}
		switch t {
		case models.IntSliderControl:
			m := data.GetKV[models.SliderModel[int64]](&dataStore, id)
			ui.IntSlider(m, id).Render(context.Background(), &buf)
		case models.FloatSliderControl:
		case models.CheckboxControl:
		}
	}
	w.Write(buf.Bytes())
}

func main() {

	AddSlider[int64]("A", 10, 0, 100, 5)

	AddSlider[int64]("B", 4, 0, 10, 1)

	AddSlider[int64]("C", 1000, 0, 10000, 100)

	comp := ui.Index()

	http.Handle("/", templ.Handler(comp))

	http.HandleFunc("/checkValues", checkValuesHandler)
	http.HandleFunc("/initControls", initControls)

	http.ListenAndServe(":3000", nil)
}
