package utils

import (
	"github.com/google/uuid"
)

// Generate a unique ID
func UniqueID() string {
	return uuid.NewString()
}
