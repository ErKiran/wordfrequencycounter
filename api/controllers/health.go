package controllers

import "net/http"

// Health godoc
// @Summary Health Check
// @Description Check if API is working
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {string} string
// @Router /api/v1/health [get]
func (server *Server) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
