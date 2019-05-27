package buildserver

import (
	"encoding/json"
	"net/http"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
)

func (s *Server) GetBuildPlan(build db.Build) http.Handler {
	hLog := s.logger.Session("get-build-plan")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		plan := build.PublicPlan()
		if plan == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err := json.NewEncoder(w).Encode(atc.PublicBuildPlan{
			Schema: build.Schema(),
			Plan:   plan,
		})
		if err != nil {
			hLog.Error("failed-to-encode-public-build-plan", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
