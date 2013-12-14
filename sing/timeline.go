package sing

import (
  "appengine"
  "appengine/user"
  "html/template"
  "net/http"
  "singer"
)

type singData struct {
  Singer singer.Singer
  Text   string
}

type timeLine struct {
  Singer singer.Singer
  Sings  []singData
}

var timeLinePageTemplate = template.Must(template.New("timeline").Parse(timeLinePageTemplateHtml))

func TimeLinePageHandler(w http.ResponseWriter, r *http.Request) {
  ct := appengine.NewContext(r)
  u := user.Current(ct)
  tl, err := newTimeLine(u.ID, ct)
  if err != nil {
    status := http.StatusInternalServerError
    http.Error(w, http.StatusText(status), status)
    return
  }
  timeLinePageTemplate.Execute(w, tl)
}

func newTimeLine(uID string, ct appengine.Context) (*timeLine, error) {
  tl := new(timeLine)
  tmpsinger, err := singer.Find(uID, ct)
  tl.Singer = *tmpsinger
  if err != nil {
    return tl, err
  }
  sings, err := FindLatest(20, ct)
  if err != nil {
    return tl, err
  }

  tl.Sings = make([]singData, len(sings))
  for i, sing := range sings {
    s, err := sing.Singer(ct)
    if err != nil {
      return tl, err
    }
    tl.Sings[i] = singData{*s, sing.Text}
  }
  return tl, err
}
