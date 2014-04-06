import (
	"net/http"
)

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Statistics</title>
<body><h3>Statistics</h3>
<p>Computes basic statistics for a given list of numbers</p>`
	form = `<form action="/" method="POST">
<label for="numbers">Numbers (comma or space-separated):</label><br />
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form>`
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":9001", nil)

}

func handleIndex(w http.ResponseWriter, r *http.Request) {
r.ParseForm()
//fmt.Fprint(w, pageTop, form)  // так или как сл. строки
w.Write([]byte(pageTop))
w.Write([]byte(form))

}
