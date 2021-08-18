package solution

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type handler struct {
	s Story
}

var defaultHandlerTemplate = `<!DOCTYPE html>
<html lang="en" dir="ltr">
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
      <p>{{.}}</p>
    {{end}}
    <ul>
      {{range .Options}}
        <li><a href="/{{.Arc}}">{{.Text}}</a></li>
      {{end}}
    </ul>
  </body>
</html>
`

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTemplate))
}

// TODO: we should returnn type http.Handler and not the handler type
// it's important for some documentation reasons which i don't understand at this
// moment. need to some back for this point later.
func NewHandler(s Story) http.Handler {
	return handler{s}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// now what we want to do is serve any chapter on demand
	// based on parsing the request path
	path := r.URL.Path
	if path == "" || path == "/" {
		path = "/intro"
	}
	// it will convert "/intro" to "intro"
	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := tpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Someting went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found", http.StatusNotFound)
}

// JSONStory ... exporting this func because it's likely that it will be used a lot at multiple places
// and putting it here because it's related to the "story". so all the types related to that and all the
// relevant functions should be placed here seperately.
func JSONStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}
