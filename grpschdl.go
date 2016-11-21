package hello

import (
  "fmt"
  "net/http"

  "appengine"
  "appengine/user"
)

func init() {
  http.HandleFunc("/", index)
  http.HandleFunc("/top", top)
}

func validateLogin(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  u := user.Current(c)
  if u == nil {
    url, err := user.LoginURL(c, r.URL.String())
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    w.Header().Set("location", url)
    w.WriteHeader(http.StatusFound)
    return
  }
}

func index(w http.ResponseWriter, r *http.Request) {
  validateLogin(w, r)
  http.Redirect(w, r, "/top", http.StatusMovedPermanently)
}

func top(w http.ResponseWriter, r *http.Request) {
  validateLogin(w, r)
  u := user.Current(appengine.NewContext(r))
  fmt.Fprintf(w, "Hello, %v!", u)
}
