package api

import (
	"net/http"
	"strconv"

	"github.com/nimilgp/URLcommentary/internal/config"
	"github.com/nimilgp/URLcommentary/internal/dblayer"
	"github.com/nimilgp/URLcommentary/internal/response"
)

func (s *APIServer) getNewestParentComments(w http.ResponseWriter, r *http.Request) {
	pageId, err := strconv.Atoi(r.PathValue("pageid"))
	if err != nil {
		s.logger.Info("invalid param", "pageid")
	}
	arg := dblayer.RetrieveNewestParentCommentsParams{
		Pageid: int32(pageId),
		Limit:  int32(config.Cfg.OffsetSize),
		Offset: 0,
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
