package main

import (
    "django_generator/lib"
    "log"
)


func main() {
    //lib.GenerateCSS(false)
    err := lib.CopyDir("assets/bootstrap", "static")

    log.Println(err)


}
