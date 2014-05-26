package lightning

import (
  "net/http"
  "os"
  "log"
  "date"
  "io"
)

const root = "https://api.spark.io/v1/"
const devices = "devices/"

type Variable struct {
  Name string `json:"name"`
  Type string `json:"type"`
}

type Core struct {
  Id string `json:"id"`
  key string
  Name string `json:"name"`
  LastApp string `json:"lastApp"`
  LastHeard date.Date `json:"lastHeard"`
  Connected bool
  Functions [4]string `json:"functions"`
  Variables []Variable `json:"variables"`
}

var client = &http.Client{}

func NewCore(id, key string) (Core, error) {
  req, err := http.NewRequest("GET", root + devices + id, nil)

  if err != nil {
    log.Fatal(err)
  }

  req.Header.Add("Authorization", "Bearer " + key)

  resp, err2 := client.Do(req)
  if err2 != nil {
    log.Fatal(err)
  }

  io.Copy(os.Stdout, resp.Body)

  return Core{id, key}, nil
}

type Collection struct {
  Key string
}

func NewCollection(key string) (Collection, error) {
  c := Collection {
    Key: key,
  }

  return c, nil
}
