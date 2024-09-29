package easyfilego

import (
    "bufio"
    "io"
    "os"
)

// Open a file based on the mode provided
func Open(filePath string, mode string) (*File, error) {
    var file *os.File
    var err error

    switch mode {
    case ModeRead: // Open in read-only mode
        file, err = os.Open(filePath)
    case ModeWrite: // Create or truncate the file for writing
        file, err = os.Create(filePath)
    case ModeAppend: // Open in append mode
        file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    case ModeReadWrite: // Open in read-write mode
        file, err = os.OpenFile(filePath, os.O_RDWR, 0644)
    case ModeWritePlus: // Create or overwrite and open in read-write mode
        file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
    case ModeReadBinary: // Open in read-only binary mode
        file, err = os.OpenFile(filePath, os.O_RDONLY, 0644)
    case ModeWriteBinary: // Create or truncate in binary write mode
        file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
    case ModeReadWriteBin: // Open in read-write binary mode
        file, err = os.OpenFile(filePath, os.O_RDWR, 0644)
    case ModeWriteAppend: // Open in read-write append mode
        file, err = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
    default:
        return nil, ErrInvalidMode
    }

    if err != nil {
        return nil, err
    }

    return &File{
        file:   file,
        mode:   mode,
        buffer: bufio.NewReader(file),
    }, nil
}

// Read the entire content of a file into a string (Buffered)
func (f *File) Read() (string, error) {
    content, err := io.ReadAll(f.buffer)
    if err != nil {
        return "", err
    }
    return string(content), nil
}

// Write string content to a file (Buffered)
func (f *File) Write(content string) error {
    if f.mode == ModeRead || f.mode == ModeReadBinary {
        return ErrWriteOnReadOnly
    }

    writer := bufio.NewWriter(f.file)
    _, err := writer.WriteString(content)
    if err != nil {
        return err
    }
    return writer.Flush()
}

// Close the file
func (f *File) Close() error {
    return f.file.Close()
}

// ReadLines reads the file line-by-line into a string array
func (f *File) ReadLines() ([]string, error) {
    var lines []string
    scanner := bufio.NewScanner(f.buffer)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        return nil, err
    }
    return lines, nil
}
