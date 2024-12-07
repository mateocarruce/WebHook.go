package main

import (
    "fmt"
    "log"
    "net/http"
)

// Handler de Webhook
func webhookHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Mateo Carrasco - WebHook recibido")
}

func main() {
    http.HandleFunc("/webhook", webhookHandler) // Ruta WebHook
    fmt.Println("Servidor escuchando en el puerto 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil)) // Arrancar el servidor en puerto 8080
}
