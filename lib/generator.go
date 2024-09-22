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

    // update the views.py
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

    defer file.Close()

    _, err = file.WriteString(targetString)


    if err != nil {
        log.Fatalln(err)
    }


    // update the tempalte dir
    currentDirName := GetLastNameFromPath(currentDirPath)
    templateDirPath := currentDirPath + string(os.PathSeparator) + "templates" + string(os.PathSeparator) + currentDirName

    _, err = os.Stat(templateDirPath)

    if os.IsNotExist(err) {
        log.Fatalln("No template exists , Generate the template first")
    }

    srcTemplateData, err := os.ReadFile(GetAssetPath() + string(os.PathSeparator) + "child.html")

    if err != nil {
        log.Fatalln(err)

    }

    childTemplatePath :=  templateDirPath + string(os.PathSeparator) + viewName + ".html"

    err = os.WriteFile(
        childTemplatePath,
        srcTemplateData, 
        0775)

    if err != nil {
        log.Fatalln(err)

    }

    // update the urls


    srcUrlPath := currentDirPath + string(os.PathSeparator) + "urls.py"
    srcUrlBytes , err := os.ReadFile(srcUrlPath)

    if err != nil {
        log.Fatalln(err)
    }

    srcUrlLines := strings.Split(string(srcUrlBytes), "\n")

    urlPatternStartIndex := -1

    for index, line := range srcUrlLines {
        if strings.Contains(line,"urlpatterns") {
            urlPatternStartIndex = index
            break
        }
    }

    urlPatternEndIndex := -1

    for index, line := range srcUrlLines[urlPatternStartIndex:len(srcUrlLines)] {
        if strings.Contains(line, "]") {
            urlPatternEndIndex = urlPatternStartIndex + index
            break
        }
    }

    newString := "\tpath(\"" + viewName + "/\"," + viewName + "),"
    srcUrlLines[urlPatternEndIndex] = newString
    srcUrlLines = append(srcUrlLines, "\t]")

    temp := "from .views import " + viewName + ""
    srcUrlLines = append([]string {temp}, srcUrlLines...)


    newUrlFileStr := ""

    for _, line := range srcUrlLines {
        newUrlFileStr += line + "\n"
    }

    err = os.Remove(srcUrlPath)

    if err != nil {
        log.Fatalln(err)
    }

    newURLFile, err := os.Create(srcUrlPath)


    if err != nil {
        log.Fatalln(err)
    }

    _, err = newURLFile.WriteString(newUrlFileStr)


    if err != nil {
        log.Fatalln(err)
    }

}




















