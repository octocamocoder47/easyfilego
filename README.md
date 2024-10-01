# easyfilego
This is a package for GoLang to make file handling simpler.

There are couple of modes in this to open a file.

```
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
```
There is a Open function in this which has 2 parameters
1. File path (String)
2. File mode which will be like easyfilego.ModeRead

To read from file use Read() and return ouput in a String format

To write to a file use Write(content String)
this will also take care if file is not created it will create it.

To close the file use Close()

To read line by line use ReadLines(). This function will return array of strings as output.
