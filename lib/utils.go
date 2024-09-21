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

    fileName := GetLastNameFromPath(srcPath)
    dstPath := dstDirPath + string(os.PathSeparator) + fileName


    err = os.WriteFile(dstPath, inFile, 0755)

    if err != nil {
        return err
    }

    return nil

}


func CopyDir(srcDirPath, dstDirPath string) error  {
    files, err := os.ReadDir(srcDirPath)

    if err != nil {
        return err
    }

    for _, file := range files  {
        log.Println(file.Name())
    }

    return nil
}
