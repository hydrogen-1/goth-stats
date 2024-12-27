package templates

import (
	_ "embed"
	"html/template"
	"net/http"
)

//go:embed stats.html
var stats string

func RenderStats(w http.ResponseWriter, total int64, last24h int64) {
	data := map[string]interface{}{
		"Total":   total,
		"Last24h": last24h,
	}
	template, err := template.New("stats").Parse(stats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}
	err = template.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
	}
}
