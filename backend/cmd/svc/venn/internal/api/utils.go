package api

import (
	"fmt"

	"github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/common"
	"github.com/gin-gonic/gin"
)

func withMeta(data, meta any) common.MetaResponse {
	return common.MetaResponse{Data: data, Meta: meta}
}

func withPaginaion(page, pageSize, total int64) common.Pagination {
	return common.Pagination{Page: page, PageSize: pageSize, Total: total}
}

// getSort returns the  SortBy associated with the key as a string.
func getSort(ctx *gin.Context) (s *common.SortBy) {
	if val, ok := ctx.Get("sort"); ok && val != nil {
		s, _ = val.(*common.SortBy)
	}
	return
}

func getOrderClause(sb *common.SortBy, fields map[string]string, defaultSort string) string {
	if sb == nil {
		return defaultSort
	}
	if col, ok := fields[sb.ByName()]; ok {
		return fmt.Sprintf("%s %s", col, sb.GetOrder())
	}
	return defaultSort
}
