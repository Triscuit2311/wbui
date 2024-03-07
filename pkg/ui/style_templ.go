// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package ui

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func style() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style type=\"text/css\">\n\n\nbody {\n    font-family: 'JetBrains Mono', 'Fira Code', 'Consolas', monospace;\n    margin: 0;\n    padding: 0;\n    background-color: #ffffff;\n}\n\n.ui-title {\n    text-align: center;\n    color: #333;\n    font-size: 28px;\n    margin: 40px 0;\n}\n\n.controls-block {\n    padding: 20px;\n    display: flex;\n    flex-direction: column;\n    align-items: center;\n    background-color: #f7f7f7;\n    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);\n    border-radius: 8px;\n    margin: 20px;\n}\n\n.controls-block > div {\n    width: 80%;\n    display: flex;\n    justify-content: center;\n    margin-bottom: 20px;\n}\n\n.intslidercontainer, .floatslidercontainer, .checkboxcontainer, .displayvaluecontainer {\n    display: flex;\n    align-items: center;\n    justify-content: space-between;\n    margin-bottom: 20px;\n    padding: 10px;\n    max-width: 90%;\n    background-color: #ffffff;\n    border: 1px solid #e0e0e0;\n    border-radius: 4px;\n}\n\n.label-text {\n    flex: 1;\n    font-size: 16px;\n    color: #333;\n}\n\ninput[type=\"range\"], input[type=\"checkbox\"] {\n    flex: 2;\n    cursor: pointer;\n}\n\ninput[type=\"range\"] {\n    -webkit-appearance: none;\n    height: 8px;\n    border-radius: 4px;\n    background: #ddd;\n    outline: none;\n}\n\ninput[type=\"range\"]::-webkit-slider-thumb {\n    -webkit-appearance: none;\n    appearance: none;\n    width: 20px;\n    height: 20px;\n    border-radius: 50%;\n    background: #007bff;\n    cursor: pointer;\n}\n\ninput[type=\"checkbox\"] {\n    -webkit-appearance: none;\n    appearance: none;\n    max-width: 18px;\n    height: 18px;\n    background-color: #f0f0f0;\n    border: 2px solid #ddd;\n    border-radius: 4px;\n    cursor: pointer;\n}\n\ninput[type=\"checkbox\"]:checked {\n    background-color: #007bff;\n    border: 2px solid #007bff;\n}\n\n.output-text {\n    flex: 1;\n    text-align: center;\n    font-size: 16px;\n    color: #333;\n}\n\nfooter {\n    text-align: center;\n    padding: 20px 0;\n    background-color: #f7f7f7;\n    position: fixed;\n    bottom: 0;\n    width: 100%;\n}\n\nfooter p {\n    margin: 0;\n    color: #333;\n    font-size: 16px;\n}\n\n\n</style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}