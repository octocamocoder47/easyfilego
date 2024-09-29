package easyfilego

import "errors"

var (
    ErrFileNotFound     = errors.New("file not found")
    ErrInvalidMode      = errors.New("invalid file mode")
    ErrWriteOnReadOnly  = errors.New("attempted to write to read-only file")
    ErrFileAlreadyExists = errors.New("file already exists")
)

