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
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
  w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
  w.WriteHeader(status)
  w.Write([]byte(response))
}

func RespondError(w http.ResponseWriter, status int, message string) {
  errorMesage := map[string]string{"error": message}
  RespondJSON(w, status, errorMesage)
}
