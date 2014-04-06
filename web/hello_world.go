/*simple web hello world*/


import (
    "github.com/hoisie/web"
)

func hello(val string) string { return "hello " + val } 

func main() {
    web.Get("/(.*)", hello)
    web.Run("0.0.0.0:9999")
}


 http://localhost:9999/world -> hello world

