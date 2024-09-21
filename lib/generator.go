package lib


import (
    "log"
    "os"
    "strings"
)



func GenerateCSS(bootstrap bool) {


    if bootstrap == false {

        dirName := "static"

        err := os.Mkdir(dirName, os.ModeDir|0775)

        if err !=  nil {
            log.Fatalln(dirName, err)
        }

        log.Println("Creating vanilla static files")

        dirName = "static" + string(os.PathSeparator) + "css"
        err = os.Mkdir(dirName, os.ModeDir|0775)

        if err !=  nil {
            log.Fatalln(dirName, err)
        }

        dirName = "static" + string(os.PathSeparator) + "js"
        err = os.Mkdir(dirName, os.ModeDir|0775)

        if err !=  nil {
            log.Fatalln(dirName, err)
        }

        err = CopyFile("assets/scripts.js", "static" + string(os.PathSeparator) + "js")

        if err !=  nil {
            log.Fatalln(err)
        }


        err = CopyFile("assets/styles.css", "static" + string(os.PathSeparator) + "css")


        if err !=  nil {
            log.Fatalln(err)
        }



    } else {
        CopyDir("assets/bootstrap", "static")

    }
}


func GenerateTemplate() {
    templateName := "templates"
    err := os.Mkdir(templateName, os.ModeDir|0775)

    if err !=  nil {
        log.Fatalln(templateName, err)
    }

    currentPath, err :=  os.Getwd()

    if err !=  nil {
        log.Fatalln(err)
    }

    currentDirName := GetLastNameFromPath(currentPath)


    dstDirPath := currentPath + string(os.PathSeparator) +  templateName +  string(os.PathSeparator) +  currentDirName


    err = os.Mkdir(dstDirPath, os.ModeDir|0775)

    if err !=  nil {
        log.Fatalln(err)
    }

    baseTemplatePath := "assets" + string(os.PathSeparator) + "base.html"
    err = CopyFile(baseTemplatePath, dstDirPath)

}

func GenerateView(viewName string) {

    currentDirPath, err := os.Getwd()

    if err != nil {
        log.Fatalln(err)

    }

    currentViewPath :=  currentDirPath + string(os.PathSeparator) + "views.py"

    if FileExists(currentViewPath) == false {
        log.Fatalln(currentViewPath, " does not exists")
    }

    assetPath := GetAssetPath()
    srcViewPath := assetPath + string(os.PathSeparator) + "views.py"

    srcData, err := os.ReadFile(srcViewPath)

    if err != nil {
        log.Fatalln(err)

    }

    srcString := string(srcData)
    targetString := strings.Replace(srcString, "target", viewName, -1)


    file, err := os.OpenFile(currentViewPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

    if err != nil {
        log.Fatalln(err)
    }

    _, err = file.WriteString(targetString)


    if err != nil {
        log.Fatalln(err)
    }


}
