package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func ParseTemplateFile(filename string, data interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		doHandlerInternal(w, r, filename, data)
	}
}

func doHandlerInternal(w http.ResponseWriter, r *http.Request, filename string, data interface{}) {
	t, err := parseTemplateFile(filename)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse file: %s. err: %q", filename, err), http.StatusInternalServerError)
	} else if err = t.Execute(w, data); err != nil {
		http.Error(w, fmt.Sprintf("unable to execute form template: %q", err), http.StatusInternalServerError)
	}
}

func parseTemplateFile(filename string) (*template.Template, error) {
	return template.ParseFiles(filename, "../../web/templates/fragments/header.html",
		"../../web/templates/fragments/navigator.html", "../../web/templates/fragments/footer.html")
}
