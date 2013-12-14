package singer

import (
  "appengine"
  "appengine/user"
  "fmt"
  "net/http"
  "strings"
)

func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    ct := appengine.NewContext(r)
    u := user.Current(ct)
    name := r.FormValue("name")
    intro := r.FormValue("introduction")

    if st := validate(name, intro); st != "" {
      http.Error(w, st, http.StatusBadRequest)
      return
    }

    _, err := Register(u, name, intro, ct)
    if err != nil {
      status := http.StatusInternalServerError
      http.Error(w, http.StatusText(status), status)
      return
    }
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
  } else {
    fmt.Fprint(w, registerPageHtml)
  }
}

func validate(name, text string) string {
  if strings.Replace(name, " ", "", -1) == "" {
    return "Name is empty."
  }
  if len(name) > 64 {
    return "Name is too long."
  }
  if len(text) > 128 {
    return "Introduction is too long."
  }
  return ""
}
