package utils

import (
	"context"
	"crypto/rand"
	"log"

	"github.com/google/uuid"
)

type contextKey string

const RequestIdKey contextKey = "request_id"

// ContextWithRequestId returns a new context with the request_id
func ContextWithRequestId(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, RequestIdKey, requestId)
}

// GetRequestId retrieves request_id from context
func GetRequestId(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if val, ok := ctx.Value(RequestIdKey).(string); ok {
		return val
	}
	return ""
}


var charstring = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomId(length int) string {

	bytes := make([]byte, length)
	randomBytes := make([]byte, length)
	if _, err := rand.Read(randomBytes); err != nil {
		log.Printf("Failed to generate random id: %v", err)
		return ""
	}

	for i, b := range randomBytes {
		bytes[i] = charstring[int(b)%len(charstring)]
	}

	return string(bytes)

}

func GenerateUuid() string {
	return uuid.NewString()
}
