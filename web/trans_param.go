/*trans parametr in web*/


func hello(ctx *web.Context, val string) {

	for par := range ctx.Params {
		println(par)
		ctx.WriteString("Hello World " + par + " ")
	}

}

func main() {
	web.Get("/(.*)", hello)
	web.Run("0.0.0.0:9999")
}

http://localhost:9999/?a=1&b=2    ->     Hello World a Hello World b