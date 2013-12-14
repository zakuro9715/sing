package sing

import (
  "appengine"
  "appengine/datastore"
  "singer"
  "time"
)

type Sing struct {
  SingerKey *datastore.Key
  Text      string `datastore:",noindex"`
  Date      time.Time
}

func (s *Sing) Singer(ct appengine.Context) (*singer.Singer, error) {
  return singer.Get(s.SingerKey, ct)
}

func Add(sk *datastore.Key, text string, ct appengine.Context) (*datastore.Key, error) {
  sing := &Sing{sk, text, time.Now()}
  k := datastore.NewIncompleteKey(ct, "Sing", nil)
  return datastore.Put(ct, k, sing)
}

func Get(k *datastore.Key, ct appengine.Context) (*Sing, error) {
  sing := new(Sing)
  err := datastore.Get(ct, k, sing)
  return sing, err
}

func FindLatest(limit int, ct appengine.Context) ([]Sing, error) {
  var sings []Sing
  q := datastore.NewQuery("Sing").
    Order("-Date").
    Limit(limit)
  _, err := q.GetAll(ct, &sings)
  return sings, err
}
