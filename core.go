package lightning

import (
  "net/http"
  "os"
  "log"
  "fmt"
  "time"
  "encoding/json"
)

const API_ROOT = "https://api.spark.io/v1/"
const API_DEVICES = "devices/"

type Variable struct {
  Name string `json:"name"`
  Type string `json:"type"`
}

type Core struct {
  Id string `json:"id"`
  key string
  Name string `json:"name"`
  LastApp string `json:"last_app"`
  LastHeard time.Time `json:"last_heard"`
  Connected bool
  Functions [4]string `json:"functions"`
  Variables map[string]string `json:"variables"`
}

var client = &http.Client{}

func NewCore(id, key string) (Core, error) {
  req, err := http.NewRequest("GET", API_ROOT + API_DEVICES + id, nil)

  if err != nil {
    log.Fatal(err)
  }

  req.Header.Add("Authorization", "Bearer " + key)

  resp, err2 := client.Do(req)
  if err2 != nil {
    log.Fatal(err)
  }

  c := Core{
    Id: id,
    key: key,
  }

  dec := json.NewDecoder(resp.Body)
  dec.Decode(&c)

  fmt.Println(c.LastHeard == time.Time{})

  enc := json.NewEncoder(os.Stdout)
  enc.Encode(c)

  return c, nil
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
