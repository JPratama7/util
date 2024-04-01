package json

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteResponseBody_HappyPath(t *testing.T) {
	data := NewReturnData(200, true, "OK", "Test Data")
	recorder := httptest.NewRecorder()

	err := data.WriteResponseBody(recorder)

	assert.NoError(t, err)
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.Contains(t, recorder.Body.String(), "Test Data")
}

func TestWriteResponseBody_ErrorInMarshalling(t *testing.T) {
	data := NewReturnData(200, true, "OK", make(chan int))
	recorder := httptest.NewRecorder()

	err := data.WriteResponseBody(recorder)

	assert.Error(t, err)
}

func TestWriteResponseBody_WithDifferentStatusCodes(t *testing.T) {
	data := NewReturnData(404, false, "Not Found", 0)
	recorder := httptest.NewRecorder()

	err := data.WriteResponseBody(recorder)

	assert.NoError(t, err)
	assert.Equal(t, 404, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.Contains(t, recorder.Body.String(), "Not Found")
}
