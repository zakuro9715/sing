package sing

import (
  "appengine"
  "appengine/user"
  "net/http"
  "singer"
  "strings"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    status := http.StatusMethodNotAllowed
    http.Error(w, http.StatusText(status), status)
    return
  }
  ct := appengine.NewContext(r)
  u := user.Current(ct)
  singer, err := singer.Find(u.ID, ct)
  if err != nil {
    status := http.StatusInternalServerError
    http.Error(w, http.StatusText(status), status)
    return
  }

  text := r.FormValue("text")
  if st := validateText(text, ct); st != "" {
    http.Error(w, st, http.StatusBadRequest)
    return
  }
  _, err = Add(singer.Key(ct), text, ct)
  if err != nil {
    status := http.StatusInternalServerError
    http.Error(w, http.StatusText(status), status)
    return
  }
  http.Redirect(w, r, "/", http.StatusSeeOther)
}

func validateText(text string, ct appengine.Context) string {
  if strings.Replace(text, " ", "", -1) == "" {
    return "Text is empty."
  }
  if len(text) > 256 {
    return "Text is too long."
  }
  sings, err := FindLatest(10, ct)
  if err != nil {
    return "Text is bad."
  }
  for _, s := range sings {
    if s.Text == text {
      return "Same text have arleady posted."
    }
  }
  return ""
}
