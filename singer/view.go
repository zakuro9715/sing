package singer

import (
  "appengine"
  "html/template"
  "net/http"
)

var viewPageTemplate = template.Must(template.New("view").Parse(viewPageTemplateHtml))

func ViewPageHandler(w http.ResponseWriter, r *http.Request) {

}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
  ct := appengine.NewContext(r)
  uID := ExtractUserID(r.URL.Path)
  s, err := Find(uID, ct)
  if err != nil {
    http.NotFound(w, r)
    return
  }
  viewPageTemplate.Execute(w, s)
}
