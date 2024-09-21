package lib


import (
    "log"
    "os"
)



func GenerateCSS(bootstrap bool) {
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



    if bootstrap == false {
        err = CopyFile("assets/scripts.js", "static" + string(os.PathSeparator) + "js")

        if err !=  nil {
            log.Fatalln(err)
        }


        err = CopyFile("assets/styles.css", "static" + string(os.PathSeparator) + "css")


        if err !=  nil {
            log.Fatalln(err)
        }


        return

    }





}


