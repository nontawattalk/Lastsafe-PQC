package gui

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"

    "github.com/nontawattalk/lastsafe-pqc/internal/orchestrator"
)

// Start launches the Fyne GUI for Lastsafe-PQC.
func Start(orch *orchestrator.Orchestrator) {
    a := app.New()
    w := a.NewWindow("Lastsafe-PQC")

    backupBtn := widget.NewButton("Backup Now", func() {
        // TODO: Gather user input for repository, password, and source directory
        _ = orch.Backup("./repo", "password", "/path/to/source")
    })

    restoreBtn := widget.NewButton("Restore", func() {
        // TODO: Provide snapshot ID and destination
        _ = orch.Restore("./repo", "password", "latest", "/restore/target")
    })

    syncBtn := widget.NewButton("Sync", func() {
        // TODO: Provide remote and local paths
        _ = orch.Sync("/path/to/source", "remote:backup")
    })

    content := container.NewVBox(
        widget.NewLabel("Lastsafe-PQC"),
        backupBtn,
        restoreBtn,
        syncBtn,
    )

    w.SetContent(content)
    w.Resize(fyne.NewSize(400, 200))
    w.ShowAndRun()
}
