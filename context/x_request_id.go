package context

import "context"

type contextKey int

const ContextKeyRequestID contextKey = iota

func SetRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, ContextKeyRequestID, requestID)
}

func GetRequestID(ctx context.Context) string {
	requestID := ctx.Value(ContextKeyRequestID)

	if value, ok := requestID.(string); ok {
		return value
	}

	return ""
}
