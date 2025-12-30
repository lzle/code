package utils

import (
	uuid "github.com/satori/go.uuid"
)

func NewUUid()(uuidString string) {
	UUid := uuid.NewV4()
	uuidString = UUid.String()
	return uuidString
}


