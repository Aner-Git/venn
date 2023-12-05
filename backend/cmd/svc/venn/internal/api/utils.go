package api

import "github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/common"

func withMeta(data, meta any) common.MetaResponse {
	return common.MetaResponse{Data: data, Meta: meta}
}

func withPaginaion(page, pageSize, total int64) common.Pagination {
	return common.Pagination{Page: page, PageSize: pageSize, Total: total}
}
