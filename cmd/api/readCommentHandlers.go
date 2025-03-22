package api

import (
	"net/http"
	"strconv"

	"github.com/nimilgp/URLcommentary/internal/config"
	"github.com/nimilgp/URLcommentary/internal/dblayer"
	"github.com/nimilgp/URLcommentary/internal/response"
)

func (s *APIServer) getNewestParentComments(w http.ResponseWriter, r *http.Request) {
	param1 := r.PathValue("pageid")
	pageId, err := strconv.Atoi(param1)
	if err != nil {
		s.logger.Info("invalid param", "pageid", param1)
	}

	param2 := r.PathValue("offset")
	offset, err := strconv.Atoi(param2)
	if err != nil {
		s.logger.Info("invalid param", "offset", param2)
	}
	arg := dblayer.RetrieveNewestParentCommentsParams{
		Pageid: int32(pageId),
		Limit:  int32(config.Cfg.OffsetSize),
		Offset: int32(offset) * int32(config.Cfg.OffsetSize),
	}
	comments, err := s.querier.RetrieveNewestParentComments(s.ctx, arg)
	if err != nil {
		s.logger.Warn("get newset comments failed")
	}
	err = response.JSON(w, http.StatusOK, envelope{"newest-parent-comments": comments})
	if err != nil {
		s.serverError(w, r, err)
	}
}

func (s *APIServer) getOldestParentComments(w http.ResponseWriter, r *http.Request) {
	param1 := r.PathValue("pageid")
	pageId, err := strconv.Atoi(param1)
	if err != nil {
		s.logger.Info("invalid param", "pageid", param1)
	}

	param2 := r.PathValue("offset")
	offset, err := strconv.Atoi(param2)
	if err != nil {
		s.logger.Info("invalid param", "offset", param2)
	}
	arg := dblayer.RetrieveOldestParentCommentsParams{
		Pageid: int32(pageId),
		Limit:  int32(config.Cfg.OffsetSize),
		Offset: int32(offset) * int32(config.Cfg.OffsetSize),
	}
	comments, err := s.querier.RetrieveOldestParentComments(s.ctx, arg)
	if err != nil {
		s.logger.Warn("get newset comments failed")
	}
	err = response.JSON(w, http.StatusOK, envelope{"oldest-parent-comments": comments})
	if err != nil {
		s.serverError(w, r, err)
	}
}
