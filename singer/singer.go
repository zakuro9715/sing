package singer

import (
  "appengine"
  "appengine/datastore"
  "appengine/user"
  "strings"
  "time"
)

type Singer struct {
  UserID       string
  Name         string
  Introduction string
  BirthDate    time.Time
}

func (s *Singer) Key(ct appengine.Context) *datastore.Key {
  return datastore.NewKey(ct, "Singer", s.UserID, 0, nil)
}

func Register(u *user.User, name string, introduction string, ct appengine.Context) (*datastore.Key, error) {
  s := &Singer{u.ID, name, introduction, time.Now()}
  k := s.Key(ct)
  return datastore.Put(ct, k, s)
}

func Get(k *datastore.Key, ct appengine.Context) (*Singer, error) {
  s := new(Singer)
  err := datastore.Get(ct, k, s)
  return s, err
}

func Find(uID string, ct appengine.Context) (*Singer, error) {
  k := datastore.NewKey(ct, "Singer", uID, 0, nil)
  return Get(k, ct)
}

func ExtractUserID(path string) string {
  splitedPath := strings.Split(path, "/")

  // /id/
  // index of id is 1.
  if len(splitedPath) < 2 {
    return ""
  }
  return splitedPath[1]
}
