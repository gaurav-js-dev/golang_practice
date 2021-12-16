package main

import (
	"flag"
	"fmt"
	adventure "golang_practice/gophercises/Adventure"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := flag.Int("port", 5500, "The Port to start Choose Your own adventure")
	filename := flag.String("file", "gopher.json", "The JSON file with the CYOA story")
	flag.Parse()
	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	story, err := adventure.JsonStory(f)

	if err != nil {
		panic(err)
	}

	// d := json.NewDecoder(f)
	// var story adventure.Story
	// if err := d.Decode(&story); err != nil {
	// 	panic(err)
	// }

	tpl := template.Must(template.New("").Parse(storyTemplate))
	h := adventure.NewHandler(story, adventure.WithTemplate(tpl), adventure.WithPathFunc(pathFn))

	mux := http.NewServeMux()
	mux.Handle("/story/", h)

	fmt.Printf("starting server at %d\n", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)

	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}

	return path[len("/story/"):]
}

var storyTemplate = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Choose your own adventure</title>
  </head>
  <body>
  <section class="page">
		<h1>{{.Title}}</h1>
		{{range .Paragraphs}}
		<p>{{.}}</p>
		{{end}}
		<ul>
		{{range .Options}}
		<li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
		{{end}}
		</ul>
	</section>
		<style>
		body {
			font-family: helvetica, arial;
		}
		h1 {
			text-align:center;
			position:relative;
		}
		.page {
			width: 80%;
			max-width: 500px;
			margin: auto;
			margin-top: 40px;
			margin-bottom: 40px;
			padding: 80px;
			background: #FFFCF6;
			border: 1px solid #eee;
			box-shadow: 0 10px 6px -6px #777;
		}
		ul {
			border-top: 1px dotted #ccc;
			padding: 10px 0 0 0;
			-webkit-padding-start: 0;
		}
		li {
			padding-top: 10px;
		}
		a,
		a:visited {
			text-decoration: none;
			color: #6295b5;
		}
		a:active,
		a:hover {
			color: #7792a2;
		}
		p {
			text-indent: 1em;
		}
		
		</style>
  </body>
</html>`
