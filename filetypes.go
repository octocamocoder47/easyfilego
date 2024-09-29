package easyfilego

import (
    "os"
    "bufio"
)

// File modes similar to Python
const (
    ModeRead         = "r"
    ModeWrite        = "w"
    ModeAppend       = "a"
    ModeReadWrite    = "r+"
    ModeWritePlus    = "w+"
    ModeReadBinary   = "rb"
    ModeWriteBinary  = "wb"
    ModeReadWriteBin = "r+b"
    ModeWriteAppend  = "a+"
)

// File struct wrapping around os.File
type File struct {
    file   *os.File
    mode   string
    buffer *bufio.Reader
}
