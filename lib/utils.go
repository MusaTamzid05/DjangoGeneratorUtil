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

    log.Println("Creating stat")

    if _, err = os.Stat(dstDirPath); os.IsNotExist(err)  {
        err := os.Mkdir(dstDirPath, os.ModeDir|0775)

        if err !=  nil {
            log.Fatalln(dstDirPath, err)
        }
    }


    for _, file := range files  {
        if file.IsDir() {
            childDirPath := srcDirPath + string(os.PathSeparator) + file.Name()
            childDstDirPath := dstDirPath + string(os.PathSeparator) + file.Name()
            err = CopyDir(childDirPath, childDstDirPath)

            if err != nil {
                log.Fatalln(err)

            }

            continue
        }


        srcFilePath := srcDirPath + string(os.PathSeparator) + file.Name()
        err = CopyFile(srcFilePath, dstDirPath)

        if err != nil {
            return err
        }
    }

    return nil
}
