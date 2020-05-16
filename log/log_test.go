package log

import (
    "testing"
    "math/rand"
    "fmt"
    "io/ioutil"
    "os"
)

func fileContentEqual(file string, content string, t *testing.T) bool {
    b, err := ioutil.ReadFile(file)
    if nil != err {
        return false
    }
    t.Logf("file => %s", b)
    t.Logf("content => %s", content)
    return string(b) == content
}

func testRotate(rotateType RotateType, t *testing.T) {
    frand := fmt.Sprintf("/tmp/rotate_%d", rand.Intn(1000))
    logger := NewWriter(frand, rotateType)
    defer logger.Close()
    file := logger.GetCurrentFilename()
    s := "test\n"
    logger.Write([]byte(s))
    if fileContentEqual(file, s, t) {
        name := "ROTATE_BY_DATE"
        switch(rotateType){
            case ROTATE_BY_HOUR:
                name = "ROTATE_BY_HOUR"
                break
            case ROTATE_BY_MINUTE:
                name = "ROTATE_BY_MINUTE"
                break
        }
        t.Error(name + " failed")
    }
    os.Remove(file)
}

func TestRotate(t *testing.T) {
    testRotate(ROTATE_BY_DATE, t)
    testRotate(ROTATE_BY_HOUR, t)
    testRotate(ROTATE_BY_MINUTE, t)
}
