package handler

import (
	"net/http"
)

func (s *Server) GetRankingsTop(w http.ResponseWriter, req *http.Request, limit int64) {
	data, err := s.s.GetGlobalRankings(req.Context(), limit)
	if err != nil {
		http.Error(w, "Failed to fetch global rankings", http.StatusInternalServerError)
		return
	}

	var res []VideoRank
	for _, d := range data {
		res = append(res, convertToVideoRank(d))
	}

	writeJSONResponse(w, &RankingsResponse{Rankings: &res})
}

func (s *Server) GetRankingsTopUserID(w http.ResponseWriter, req *http.Request, limit int64, userID string) {
	data, err := s.s.GetUserRankings(req.Context(), limit, userID)
	if err != nil {
		http.Error(w, "Failed to fetch user rankings", http.StatusInternalServerError)
		return
	}

	var res []VideoRank
	for _, d := range data {
		res = append(res, convertToVideoRank(d))
	}

	writeJSONResponse(w, &RankingsResponse{Rankings: &res})
}
