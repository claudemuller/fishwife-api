package routes

import (
	"encoding/json"
	"net/http"

	"github.com/claudemuller/fishwife/internal/pkg"
)

func Health(w http.ResponseWriter, _ *http.Request, _ pkg.AppState) {
	res := pkg.Response{
		Message: "Service is Healthy",
	}

	resJSON, err := json.Marshal(res)
	if err != nil {
		pkg.SendInternalServerError(w)

		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(resJSON)
}
