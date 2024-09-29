package easyfilego

import (
    "os"
    "io"
)

// Exists checks if a file exists
func Exists(filePath string) bool {
    _, err := os.Stat(filePath)
    return !os.IsNotExist(err)
}

// Remove deletes a file
func Remove(filePath string) error {
    return os.Remove(filePath)
}

// Copy copies a file from src to dest using os and io methods
func Copy(src, dest string) error {
    sourceFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer sourceFile.Close()

    destFile, err := os.Create(dest)
    if err != nil {
        return err
    }
    defer destFile.Close()

    _, err = io.Copy(destFile, sourceFile)
    if err != nil {
        return err
    }
    return nil
}

// Move renames (moves) a file from src to dest
func Move(src, dest string) error {
    return os.Rename(src, dest)
}
