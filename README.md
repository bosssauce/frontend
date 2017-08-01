# frontend

a Ponzu addon to create a web frontend for your CMS.

### To install:

from within your Ponzu project, run:
```bash
$ ponzu add github.com/bosssauce/frontend
```
Usage: 

```go
// content/song.go

package content

import (
    // import the frontend addon
    "github.com/bosssauce/frontend"
    ...
)

type Song struct {
    title   string  `json:"title"`
    ...
}

func init() {
    // add routes/handlers to the frontend Router, which embeds a *mux.Router
    frontend.Router.HandleFunc("/songs", func(res http.ResponseWriter, req *http.Request) {
        // GET /api/contents?type=Song
        ...
    })
}
```
