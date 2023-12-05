package api

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSuccesOrAbortNoError(t *testing.T) {
	got := successOrAbort(&gin.Context{}, 500, nil)
	assert.True(t, got)
}

func TestSuccesOrAbortError(t *testing.T) {
	rec := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(rec)
	successOrAbort(ctx, 500, errors.New("err"))

	if rec.Code != 500 {
		t.Fail()
	}
}
