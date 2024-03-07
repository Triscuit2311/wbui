package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/a-h/templ"
	"log"
	"net/http"
	"strconv"
	"time"
	data "wbui/pkg/data"
	"wbui/pkg/models"
	ui "wbui/pkg/ui"
)

// TODO:
// write tests to ensure data values are updated corretly from ui controls
// design value display component
// something fun

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

func getValue(w http.ResponseWriter, r *http.Request) {
	id64, err := strconv.ParseInt(r.Header.Get("ID"), 10, 64)
	if err != nil {
		log.Fatalf("Bad parse on request ID in getValue (id:%v)", r.Header.Get("ID"))
		panic(err)
	}

	id := int(id64)
	t, ok := controlMap[id]
	if !ok {
		log.Fatalf("Bad ID in getValue (not in map) (id:%v)", r.Header.Get("ID"))
	}

	if t != models.DisplayaValueControl {
		log.Fatalf("Bad model type calling getValue(id:%v)", r.Header.Get("ID"))
	}

	m := data.GetKV[models.DisplayValueModel](&dataStore, id)
	w.Write([]byte(m.Val))
}

func updateValue(w http.ResponseWriter, r *http.Request) {
	id64, err := strconv.ParseInt(r.Header.Get("ID"), 10, 64)
	if err != nil {
		log.Fatalf("Bad parse on request ID in updateValue (id:%v)", r.Header.Get("ID"))
		panic(err)
	}

	id := int(id64)
	t, ok := controlMap[id]
	if !ok {
		log.Fatalf("Bad ID in updateValue (not in map) (id:%v)", r.Header.Get("ID"))
	}

	switch t {
	case models.IntSliderControl:
		m := data.GetKV[models.SliderModel[int64]](&dataStore, id)

		v, err := strconv.ParseInt(r.FormValue("control"), 10, 64)
		if err != nil {
			log.Fatalf("Error parsing value in updateValue (ID:%d)", id)
		}
		m.Val = v
		data.UpdateKV(&dataStore, id, m)
	case models.FloatSliderControl:
		m := data.GetKV[models.SliderModel[float64]](&dataStore, id)
		v, err := strconv.ParseFloat(r.FormValue("control"), 64)
		if err != nil {
			log.Fatalf("Error parsing value in updateValue (ID:%d)", id)
		}
		m.Val = v
		data.UpdateKV(&dataStore, id, m)
	case models.CheckboxControl:
		m := data.GetKV[models.CheckboxModel](&dataStore, id)
		v, err := strconv.ParseBool(r.Header.Get("VALUE"))
		if err != nil {
			log.Fatalf("Error parsing value in updateValue (ID:%d) (val:%s)", id, r.Header.Get("VALUE"))
		}
		m.Val = v
		data.UpdateKV(&dataStore, id, m)
		var buf bytes.Buffer
		ui.Checkbox(m, id).Render(context.Background(), &buf)
		w.Write(buf.Bytes())
	}

	// fmt.Printf("Update hit: ID:%v\n", id)
	// for k, v := range r.Form {
	// 	fmt.Printf("[%v:%v]\n", k, v)
	// }
	// for k, v := range r.Header {
	// 	fmt.Printf("[%v:%v]\n", k, v)
	// }

}

func AddSlider[T int64 | float64](label string, val, min, max, step T) {
	model := models.NewSliderModel[T](label, val, min, max, step)
	id := data.SetKV[models.SliderModel[T]](&dataStore, model)
	controlMap[id] = model.GetType()
}

func AddCheckbox(label string, val bool) {
	model := models.NewCheckboxModel(label, val)
	id := data.SetKV[models.CheckboxModel](&dataStore, model)
	controlMap[id] = model.GetType()
}

func AddDisplay(label string, val string) int {
	model := models.NewDisplayValueModel(label, val)
	id := data.SetKV[models.DisplayValueModel](&dataStore, model)
	controlMap[id] = model.GetType()
	return id
}

func initControls(w http.ResponseWriter, r *http.Request) {
	fmt.Println("initControls Called")
	var buf bytes.Buffer

	for id := 0; id <= len(controlMap); id++ {
		t, ok := controlMap[id]
		if !ok {
			continue
		}
		buf.Write(getControlHTML(id, t))
	}
	w.Write(buf.Bytes())
}

func getControlHTML(id, t int) []byte {
	var buf bytes.Buffer
	switch t {
	case models.IntSliderControl:
		m := data.GetKV[models.SliderModel[int64]](&dataStore, id)
		ui.IntSlider(m, id).Render(context.Background(), &buf)
	case models.FloatSliderControl:
		m := data.GetKV[models.SliderModel[float64]](&dataStore, id)
		ui.FloatSlider(m, id).Render(context.Background(), &buf)
	case models.CheckboxControl:
		m := data.GetKV[models.CheckboxModel](&dataStore, id)
		ui.Checkbox(m, id).Render(context.Background(), &buf)
	case models.DisplayaValueControl:
		m := data.GetKV[models.DisplayValueModel](&dataStore, id)
		ui.DisplayValue(m, id).Render(context.Background(), &buf)
	}
	return buf.Bytes()
}

func main() {

	AddSlider[int64]("i1", 10, 0, 100, 5)

	AddSlider[int64]("i2", 4, 0, 110, 1)

	AddSlider[int64]("i3", 1000, 0, 10000, 100)

	AddSlider[float64]("f1", 1.0, 0.0, 10.0, 0.5)

	AddSlider[float64]("f2", 1.0, 0.0, 10.0, 0.001)

	AddCheckbox("b1", true)

	AddCheckbox("b2", false)

	//	AddCheckbox("b4", true)

	displayTracker := 0
	displayId := AddDisplay("S1", fmt.Sprintf("%d", displayTracker))

	AddCheckbox("b3", true)

	comp := ui.Index()

	http.Handle("/", templ.Handler(comp))

	http.HandleFunc("/checkValues", checkValuesHandler)
	http.HandleFunc("/internal/updateValue", updateValue)
	http.HandleFunc("/internal/getValue", getValue)
	http.HandleFunc("/initControls", initControls)

	go http.ListenAndServe(":3000", nil)

	for {
		time.Sleep(time.Second * 5)

		//TODO: rm test
		displayTracker++

		m := data.GetKV[models.DisplayValueModel](&dataStore, displayId)
		m.Val = fmt.Sprintf("%d", displayTracker)
		data.UpdateKV(&dataStore, displayId, m)

		for id := 0; id <= len(controlMap); id++ {
			t, ok := controlMap[id]
			if !ok {
				continue
			}

			switch t {
			case models.IntSliderControl:
				m := data.GetKV[models.SliderModel[int64]](&dataStore, id)
				fmt.Printf("\tID:%d [v:%v]", id, m.Val)
			case models.FloatSliderControl:
				m := data.GetKV[models.SliderModel[float64]](&dataStore, id)
				fmt.Printf("\tID:%d [v:%v]", id, m.Val)
			case models.CheckboxControl:
				m := data.GetKV[models.CheckboxModel](&dataStore, id)
				fmt.Printf("\tID:%d [v:%v]", id, m.Val)

			}
		}
		fmt.Print("\n")
	}
}
