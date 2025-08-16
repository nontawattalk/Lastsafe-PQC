package rclone

import (
    "fmt"
    "os/exec"
)

// Sync synchronizes files from src to dst using rclone.
// Additional options can be passed via options slice.
func Sync(src, dst string, options ...string) error {
    args := []string{"sync", src, dst}
    args = append(args, options...)
    cmd := exec.Command("rclone", args...)
    out, err := cmd.CombinedOutput()
    fmt.Println(string(out))
    return err
}
