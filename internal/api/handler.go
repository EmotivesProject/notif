package api

import (
	"net/http"

	"github.com/TomBowyerResearchProject/common/response"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	response.MessageResponseJSON(w, http.StatusOK, response.Message{Message: "Health ok"})
}
