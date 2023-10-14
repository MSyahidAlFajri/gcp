package gcf

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/petapedia/peda"
)

func init() {
	functions.HTTP("PetaPedia", Geojson)
}

func Geojson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://msyahidalfajri.github.io")
	fmt.Fprintf(w, peda.GCFHandler("MONGOULBI", "geojson", "bandaaceh"))
}