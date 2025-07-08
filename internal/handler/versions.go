package handler

import (
	"net/http"
	"text/template"
)

type VersionData struct {
	Version string
	Pods    []string
	Percent float64
}

type PageData struct {
	Namespace string
	Versions  []VersionData
}

var templates = template.Must(template.ParseGlob("cmd/vermilion/templates/*.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	templates.ExecuteTemplate(w, "layout", data)
}

func VersionsHandler(w http.ResponseWriter, r *http.Request) {
	ns := r.URL.Query().Get("namespace")
	pods := mockFetch(ns)

	total := 0
	for _, v := range pods {
		total += len(v.Pods)
	}

	for i := range pods {
		pods[i].Percent = float64(len(pods[i].Pods)) / float64(total) * 100
	}

	data := PageData{
		Namespace: ns,
		Versions:  pods,
	}
	templates.ExecuteTemplate(w, "table", data)
}

func mockFetch(ns string) []VersionData {
	return []VersionData{
		{Version: "1.0.0", Pods: []string{"a", "b"}},
		{Version: "2.0.0", Pods: []string{"c", "d"}},
	}
}
