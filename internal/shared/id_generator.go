package shared

import "github.com/google/uuid"

func NewIDGenerator() string {
	return uuid.New().String()
}
