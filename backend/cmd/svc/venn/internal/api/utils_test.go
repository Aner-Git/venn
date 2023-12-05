package api

import (
	"testing"

	"github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUnknownError(t *testing.T) {
	want := common.Pagination{Page: 1, PageSize: 10, Total: 100}
	got := withPaginaion(1, 10, 100)
	assert.Equal(t, want, got)
}

func TestWithMeta(t *testing.T) {
	want := common.MetaResponse{Data: "foo", Meta: "bar"}
	got := withMeta("foo", "bar")
	assert.Equal(t, want, got)
}

func TestGetSort(t *testing.T) {
	ctx := &gin.Context{}
	sb := getSort(ctx)
	assert.Nil(t, sb)

	want := &common.SortBy{ID: "name", Desc: false}
	ctx.Set("sort", want)
	got := getSort(ctx)
	assert.Equal(t, want, got)
}

func TestGetOrderClause(t *testing.T) {
	defaultField := "foo"
	sortField := getOrderClause(nil, map[string]string{}, defaultField)
	assert.Equal(t, defaultField, sortField)

	sortField = getOrderClause(&common.SortBy{}, map[string]string{}, defaultField)
	assert.Equal(t, defaultField, sortField)

	sortField = getOrderClause(&common.SortBy{ID: "foo", Desc: true}, map[string]string{"foo": "bar"}, defaultField)
	assert.Equal(t, "bar DESC", sortField)

	sortField = getOrderClause(&common.SortBy{ID: "foo", Desc: false}, map[string]string{"foo": "bar"}, defaultField)
	assert.Equal(t, "bar ASC", sortField)
}
