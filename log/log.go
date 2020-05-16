package log

import (
    "io"
    //"fmt"
    "os"
    "sync"
    "time"
)

type RotateType int

const (
    ROTATE_BY_DATE RotateType = 0
    ROTATE_BY_HOUR RotateType = 1
    ROTATE_BY_MINUTE RotateType = 2
)

type RotateFileWriter struct {
    io.Writer
    currentFilename string
    baseFilename string
    rotateType RotateType
    fp *os.File
    lock sync.Mutex
}

func (logger *RotateFileWriter) Write(p []byte) (n int, err error) {
    logger.rotateFile()
    return logger.fp.Write(p)
}

func (logger *RotateFileWriter) rotateFile() {
    logger.lock.Lock()
    defer logger.lock.Unlock()
    t := time.Now()
    newFileName := ""
    basePath := logger.baseFilename
    if ROTATE_BY_MINUTE == logger.rotateType {
        newFileName = basePath + "_" + t.Format("0601021504") + ".log"
    } else if ROTATE_BY_HOUR == logger.rotateType {
        newFileName = basePath + "_" + t.Format("06010215") + ".log"
    } else {
        newFileName = basePath + "_" + t.Format("060102") + ".log"
    }
    if newFileName != logger.currentFilename {
        fp, err := os.OpenFile(newFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
        if nil != err {
            panic(err)
        }
        logger.currentFilename = newFileName
        logger.fp = fp
    }
}

func (logger *RotateFileWriter) Close() {
    logger.lock.Lock()
    defer logger.lock.Unlock()
    logger.fp.Close()
}

func (logger *RotateFileWriter) GetCurrentFilename() string {
    return logger.currentFilename
}

func NewWriter(baseFilename string, rotateType RotateType) *RotateFileWriter {
    logger := &RotateFileWriter {
        currentFilename: "",
        baseFilename: baseFilename,
        rotateType: rotateType,
    }
    return logger
}
