package orchestrator

import (
    "github.com/nontawattalk/lastsafe-pqc/internal/pqc"
    "github.com/nontawattalk/lastsafe-pqc/internal/restic"
    "github.com/nontawattalk/lastsafe-pqc/internal/rclone"
)

type Orchestrator struct {
    pqcManager *pqc.PQCManager
}

func New() *Orchestrator {
    return &Orchestrator{
        pqcManager: pqc.NewPQC(),
    }
}

// Backup runs a restic backup for the given source directory into the repository.
func (o *Orchestrator) Backup(repo, password, src string) error {
    return restic.Backup(repo, password, src)
}

// Restore runs a restic restore of a given snapshot ID into the target directory.
func (o *Orchestrator) Restore(repo, password, snapshotID, target string) error {
    return restic.Restore(repo, password, snapshotID, target)
}

// Sync synchronizes a local path to a remote path using rclone.
func (o *Orchestrator) Sync(local, remote string) error {
    return rclone.Sync(local, remote)
}
