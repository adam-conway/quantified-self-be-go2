package handler

import (
  "encoding/json"
  "net/http"
)

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
  response, err := json.Marshal(payload)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
  w.Write([]byte(response))
}

func RespondError(w http.ResponseWriter, status int, message string) {
  errorMesage := map[string]string{"error": message}
  RespondJSON(w, status, errorMesage)
}
