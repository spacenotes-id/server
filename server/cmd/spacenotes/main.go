package main

import (
	"github.com/tfkhdyt/SpaceNotes/server/internal/infrastructure/container"
	"github.com/tfkhdyt/SpaceNotes/server/internal/infrastructure/http"
)

func init() {
	container.InitDi()
}

func main() {
	http.StartFiberServer()
}
