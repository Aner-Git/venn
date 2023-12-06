package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestParseSortDefault(t *testing.T) {
	got := parseSortDefault()
	assert.NotNil(t, got)
}

func TestParseSortFilterErrors(t *testing.T) {
	got := parseSortDefault()
	assert.NotNil(t, got)

	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request, _ = http.NewRequest("GET", "?sort=[foo]", nil)

	got(ctx)
	if rec.Code != 400 {
		t.Fail()
	}
}
