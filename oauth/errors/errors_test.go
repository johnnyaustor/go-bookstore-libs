package errors

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestInternalServerError(t *testing.T) {
	restError := InternalServerError("this is message")
	assert.NotNil(t, restError)
	assert.EqualValues(t, "this is message", restError.Message)
	assert.EqualValues(t, http.StatusInternalServerError, restError.Status)
	assert.EqualValues(t, http.StatusText(http.StatusInternalServerError), restError.Error)
}

func TestNotFound(t *testing.T) {
	restError := NotFound("this is message")
	assert.NotNil(t, restError)
	assert.EqualValues(t, "this is message", restError.Message)
	assert.EqualValues(t, http.StatusNotFound, restError.Status)
	assert.EqualValues(t, http.StatusText(http.StatusNotFound), restError.Error)
}

func TestBadRequest(t *testing.T) {
	restError := BadRequest("this is message")
	assert.NotNil(t, restError)
	assert.EqualValues(t, "this is message", restError.Message)
	assert.EqualValues(t, http.StatusBadRequest, restError.Status)
	assert.EqualValues(t, http.StatusText(http.StatusBadRequest), restError.Error)
}
