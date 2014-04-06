/*simple revel start*/

для создания проекта в коммандной строке:
revel new myapp

для запуска приложения
revel run myapp

в браузере http://localhost:9000/

**************************
conf / routes
**************************
GET / App.Index


***************************
app.go
***************************
package controllers

import "github.com/robfig/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	name := "Alex"
	return c.Render(name)
}

func (c App) Hello(myName string) revel.Result {
	return c.Render(myName)
}


*******************************
Hello.html
*******************************
{{set . "title" "Home"}}
{{template "header.html" .}}

<h1>Hello {{.myName}}</h1>
<a href="/">Back to form</a>

{{template "footer.html" .}}

******************************
index.html
******************************
{{set . "title" "Home"}}
{{template "header.html" .}}

<header class="hero-unit" style="background-color:#A9F16C">
  <div class="container">
    <div class="row">
      <div class="hero-text">
        <h1>Hello {{.name}}</h1>
        <p></p>
      </div>
    </div>
  </div>
</header>

<div class="container">
  <div class="row">
    <div class="span6">
	  <form action="/App/Hello" method="GET">
	    <input type="text" name="myName" /><br/>
		<input type="submit" value="Say Hello!" />
	  </form>
      {{template "flash.html" .}}
    </div>
  </div>
</div>

{{template "footer.html" .}}

