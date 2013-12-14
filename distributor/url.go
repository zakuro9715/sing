package distributor

import (
  "net/http"
  "sing"
  "singer"
)

var urlMap = map[string]http.HandlerFunc{
  "^/$":          sing.TimeLinePageHandler,
  "^/post/$":     sing.PostHandler,
  "^/w/$":        singer.RegisterPageHandler,
  "^/register/$": singer.RegisterPageHandler,
}
