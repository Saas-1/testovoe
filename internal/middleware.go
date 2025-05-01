package internal

import (
	"awesomeProject19/logger"
	"net/http"
)

func AllowMethod(h http.HandlerFunc, method string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Infof("Received request: %s %s\n", r.Method, r.URL)

		if r.Method != method {
			logger.Log.Warnf("Method not allowed: %s for URL: %s\n", r.Method, r.URL)
			http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
			return
		}

		logger.Log.Infof("Allowed method: %s for URL: %s\n", r.Method, r.URL)
		h(w, r)
	}
}
