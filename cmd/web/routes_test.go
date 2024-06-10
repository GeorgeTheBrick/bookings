package main

import (
	"fmt"
	"testing"

	"github.com/GeorgeTheBrick/bookings/internal/config"
	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing: Test Passed

	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux, type is %T", v))
	}
}
