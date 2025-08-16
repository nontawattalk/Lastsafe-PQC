package restic

import (
    "fmt"
    "os/exec"
    "strings"
)

// Backup runs a restic backup for the given source directory into the repository.
func Backup(repo, password, src string) error {
    cmd := exec.Command("restic", "-r", repo, "backup", src)
    // supply password via environment variable or other secure method
    cmd.Env = append(cmd.Env, fmt.Sprintf("RESTIC_PASSWORD=%s", password))
    out, err := cmd.CombinedOutput()
    fmt.Println(string(out))
    return err
}

// Restore runs a restic restore of a given snapshot ID into the target directory.
func Restore(repo, password, snapshotID, target string) error {
    cmd := exec.Command("restic", "-r", repo, "restore", snapshotID, "--target", target)
    cmd.Env = append(cmd.Env, fmt.Sprintf("RESTIC_PASSWORD=%s", password))
    out, err := cmd.CombinedOutput()
    fmt.Println(string(out))
    return err
}
