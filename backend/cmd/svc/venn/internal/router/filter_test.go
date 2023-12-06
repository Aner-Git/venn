package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestParseFilterDefault(t *testing.T) {
	got := parseFilterDefault()
	assert.NotNil(t, got)
}

func TestFilterNullNoFilterQuery(t *testing.T) {
	got := parseFilter([]string{})
	assert.NotNil(t, got)
	ctx := &gin.Context{}
	got(ctx)
	_, exist := ctx.Get(filterText)
	assert.False(t, exist)
}

func TestParseQueryFilter(t *testing.T) {
	got := parseFilter([]string{"fooFilter"})
	assert.NotNil(t, got)

	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request, _ = http.NewRequest("GET", "?fooFilter=[0]", nil)

	got(ctx)
	filterMap := ctx.GetStringMapString(filterText)
	assert.NotNil(t, filterMap)
	filterValueGot, exist := filterMap["fooFilter"]
	assert.True(t, exist)
	assert.Equal(t, "[0]", filterValueGot)
}
