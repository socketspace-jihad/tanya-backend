package middlewares

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func GetStudentClassIDFromContext(r *http.Request) (uint, error) {
	schoolClassID, ok := r.Context().Value(ContextKey("school_class_id")).(uint)
	if !ok {
		return 0, errors.New("school_class_id not found")
	}
	return schoolClassID, nil
}

func StudentClassMiddlewares(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		studentClassID := r.URL.Query().Get("school_class_id")
		if studentClassID == "" {
			http.Error(w, fmt.Errorf("this endpoint should be contains school_class_id query params").Error(), http.StatusBadRequest)
		}
		studentClassIDParsed, err := strconv.Atoi(studentClassID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), ContextKey("school_class_id"), uint(studentClassIDParsed))
		h(w, r.WithContext(ctx))
	}
}
