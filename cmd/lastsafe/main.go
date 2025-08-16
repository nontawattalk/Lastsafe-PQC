package main

import (
    "github.com/nontawattalk/lastsafe-pqc/internal/orchestrator"
    "github.com/nontawattalk/lastsafe-pqc/internal/gui"
)

func main() {
    orch := orchestrator.New()
    gui.Start(orch)
}
