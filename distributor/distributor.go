package distributor

import (
  "appengine"
  "appengine/user"
  "net/http"
  "regexp"
  "singer"
  "strings"
)

func init() {
  http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
  ct := appengine.NewContext(r)
  u := user.Current(ct)
  path := r.URL.Path
  if u == nil {
    redirectLoginPage(w, r, ct)
    return
  }

  // 無限リダイレクトになるのでpathが/register/のときはリダイレクトしない
  if !isRegistered(u, ct) && path != "/register/" {
    http.Redirect(w, r, "/register/", http.StatusSeeOther)
    return
  }
  // 登録済みならトップに飛ばす
  if isRegistered(u, ct) && path == "/register/" {
    http.Redirect(w, r, "/", http.StatusSeeOther)
    return
  }
  if !strings.HasSuffix(path, "/") {
    r.URL.Path += "/"
    http.Redirect(w, r, r.URL.String(), http.StatusSeeOther)
  }
  distribute(w, r)
}

func distribute(w http.ResponseWriter, r *http.Request) {
  path := r.URL.Path
  for pattern, handler := range urlMap {
    match, err := regexp.MatchString(pattern, path)
    if err != nil {
      status := http.StatusInternalServerError
      http.Error(w, http.StatusText(status), status)
      return
    }
    if match {
      handler(w, r)
      return
    }
  }
  http.NotFound(w, r)
}

func isRegistered(u *user.User, ct appengine.Context) bool {
  _, err := singer.Find(u.ID, ct)
  return err == nil
}

func redirectLoginPage(w http.ResponseWriter, r *http.Request, ct appengine.Context) {
  url, err := user.LoginURL(ct, r.URL.String())
  if err != nil {
    status := http.StatusInternalServerError
    http.Error(w, http.StatusText(status), status)
  }
  http.Redirect(w, r, url, http.StatusSeeOther)
}
