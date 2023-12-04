package main

import "github.com/Aner-Git/venn/backend/cmd/svc/venn/internal/router"

func main() {
	r := router.Setup()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
