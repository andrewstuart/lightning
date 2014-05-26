package lightning

import (
  "net/http"
  "fmt"
  "time"
  "io/ioutil"
  "encoding/json"
)

const API_ROOT = "https://api.spark.io/v1/"
const API_DEVICES = "devices/"

//TODO Make this an interface. For http, serial, local tcp, etc...
// type CoreTalker interface {
//   func Fn (string, ...string) (string, error)
//   func Var (string) (string, error)
//   func Event (string) (chan string, error)
// }

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

var client = new(http.Client)

func (c Core) do (req *http.Request) (*http.Response, error) {
  req.Header.Add("Authorization", "Bearer " + c.key)

  return client.Do(req)
}

type strArgs struct {
  str string
}

func (s strArgs) Write(b []byte) (int, error) {
  b = append(b, []byte(s.str)...)

  return len(b), nil
}

//Fn takes the name of the function to run and as many arguments as needed and will pass those arguments to the function
func (c *Core) Fn (fname, fargs string) (string, error) {
  req, err1 := http.NewRequest("POST", API_ROOT + API_DEVICES + c.Id + "/" + fname, nil)
  if err1 != nil {
    return "", err1
  }

  fmt.Println(req)

  args := strArgs{fargs}

  err0 := req.Write(args)
  if err0 != nil {
    return "", err0
  }

  resp, err2 := c.do(req)

  if err2 != nil {
    return "", err2
  }

  if str, err := ioutil.ReadAll(resp.Body); err != nil {
    return "", err
  } else {
    return string(str), nil
  }
}

func NewCore(id, key string) (Core, error) {
  req, err1 := http.NewRequest("GET", API_ROOT + API_DEVICES + id, nil)

  //TODO: Abstract errors
  if err1 != nil {
    return Core{}, err1
  }

  c := Core{
    Id: id,
    key: key,
  }

  resp, err2 := c.do(req)

  if err2 != nil {
    return Core{}, err2
  }

  dec := json.NewDecoder(resp.Body)
  dec.Decode(&c)

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
