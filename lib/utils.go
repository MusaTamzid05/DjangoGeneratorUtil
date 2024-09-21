package lib

import (
    "strings"
    "os"
    "log"
    
)

func GetLastNameFromPath(path string) string {
    pathList := strings.Split(path, string(os.PathSeparator))
    return pathList[len(pathList) - 1]

}

func CopyFile(srcPath, dstDirPath string) error  {
    inFile, err := os.ReadFile(srcPath)

    if err != nil {
        return err
    }

    log.Println("Read ..")

    fileName := GetLastNameFromPath(srcPath)
    dstPath := dstDirPath + string(os.PathSeparator) + fileName

    f, err := os.Create(dstPath)

    if err != nil {
        return err
    }

    defer f.Close()

    err = os.WriteFile(dstPath, inFile, 0755)

    if err != nil {
        return err
    }

    return nil

}
