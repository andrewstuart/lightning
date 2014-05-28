package lightning

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const API_ROOT = "https://api.spark.io/v1/"
const API_DEVICES = "devices/"

type Core struct {
	Id        string `json:"id"`
	key       string
	Name      string    `json:"name"`
	LastApp   string    `json:"last_app"`
	LastHeard time.Time `json:"last_heard"`
	Connected bool
	Functions [4]string         `json:"functions"`
	Variables map[string]string `json:"variables"`
}

var client = &http.Client{}

func (c *Core) do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+c.key)
	req.Header.Add("Content-Type", "application/json")

	return client.Do(req)
}

type fnArgs struct {
	Args []string `json:"args,omitempty"`
}

func (f fnArgs) Read(b []byte) (l int, e error) {
	if j, err := json.Marshal(f); err == nil {
		lj := len(j)
		l = copy(b, j[:lj])
		return
	} else {
		e = err
		return
	}
}

//Fn takes the name of the function to run and as many arguments as needed and will pass those arguments to the function
func (c *Core) Fn(fname string, fargs ...string) (string, error) {
	args := fnArgs{fargs}

	req, err1 := http.NewRequest("POST", API_ROOT+API_DEVICES+c.Id+"/"+fname, args)
	if err1 != nil {
		return "", err1
	}

	req.Write(os.Stdout)

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
	req, err1 := http.NewRequest("GET", API_ROOT+API_DEVICES+id, nil)

	//TODO: Abstract errors
	if err1 != nil {
		return Core{}, err1
	}

	c := Core{
		Id:  id,
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
	c := Collection{
		Key: key,
	}

	return c, nil
}
