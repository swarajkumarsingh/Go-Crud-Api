package utils

import (
	"context"
	"net/http"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

func StartLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, _ := uuid.NewV4()
		reqLog := map[string]string{
			"requestID": u.String(),
			"startTime": time.Now().Format("2006-01-02 15:04:05.000000000"),
			"endTime":   "",
		}
		ctx := context.WithValue(r.Context(), "reqLog", reqLog)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
