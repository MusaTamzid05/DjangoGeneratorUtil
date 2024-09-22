package main

import (
    "flag"
    "django_generator/lib"
    "log"
)


func main() {
    cssPtr := flag.Bool("css", false , "Generate css")
    bootstrapPtr := flag.Bool("bootstrap", false , "Bootstrap generator")
    templatePtr := flag.Bool("template", false , "Template generator")
    viewNamePtr := flag.String("view", "", "The name of the view")

    flag.Parse()

    if *cssPtr {
        lib.GenerateCSS(*bootstrapPtr)

        return
    }

    if *templatePtr {
        lib.GenerateTemplate()
        return
    }

    if *viewNamePtr != "" {
        lib.GenerateView(*viewNamePtr)
    } else {
        log.Println("View name cannot be empty")

    }




    //lib.GenerateCSS(true)
    //lib.GenerateTemplate()
    //lib.GenerateView("city")


}
