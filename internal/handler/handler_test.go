package handler

import (
	"testing"

	"github.com/Futturi/vktest/internal/service"
	"github.com/magiconair/properties/assert"
)

func TestHandler_NewHandl(t *testing.T) {
	services := &service.Service{}
	handler := NewHandl(services)
	assert.Equal(t, handler, &Handl{services})
}
