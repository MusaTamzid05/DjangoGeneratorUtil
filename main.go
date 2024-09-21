package main

import (
    "django_generator/lib"
)


func main() {
    //lib.GenerateCSS(false)
    lib.CopyDir("assets/bootstrap", "static")


}
