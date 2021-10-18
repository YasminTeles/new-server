package context

import "context"

const ContextKeyRequestID string = "requestID"

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
