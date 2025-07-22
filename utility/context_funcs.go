package utility

import (
	"context"

	"github.com/Zillaforge/toolkits/tracer"
)

// MustGetContextRequestID ...
func MustGetContextRequestID(ctx context.Context) (id string) {
	requestID := ctx.Value(tracer.RequestID)
	if requestID != nil {
		rid, ok := requestID.(string)
		if ok {
			return rid
		}
	}
	return tracer.EmptyRequestID
}
