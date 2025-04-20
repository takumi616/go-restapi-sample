package util

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

func WriteResponse(ctx context.Context, w http.ResponseWriter, statusCode int, body any) {
	if err := json.NewEncoder(w).Encode(body); err != nil {
		slog.ErrorContext(
			ctx, "failed to write an response",
			"err", err,
		)
	}
}
