//A library for interfacing with the Spark cloud api
package lightning

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

const API_ROOT = "https://api.spark.io/v1/"
const API_DEVICES = "devices/"

type Core struct {
	Id        string            `json:"id"`
	Name      string            `json:"name"`
	LastApp   string            `json:"last_app"`
	LastHeard time.Time         `json:"last_heard"`
	Connected bool              `json:"connected"`
	Functions [4]string         `json:"functions"`
	Variables map[string]string `json:"variables"`
	key       string
}

var client = &http.Client{}

func (c *Core) do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+c.key)
	req.Header.Add("Content-Type", "application/json")

	return client.Do(req)
}

//fnResponse is the internal type for decoding the response
type fnResponse struct {
	ReturnValue interface{} `json:"return_value"`
}

//Fn takes the name of the function to run and as many arguments as needed and will pass those arguments to the function
func (c *Core) Fn(fname string, fargs ...string) (interface{}, error) {

	s := "[\"" + strings.Join(fargs, "\",\"") + "\"]"

	args := strings.NewReader(s)

	req, err1 := http.NewRequest("POST", API_ROOT+API_DEVICES+c.Id+"/"+fname, args)
	if err1 != nil {
		return "", err1
	}

	resp, err2 := c.do(req)
	if err2 != nil {
		return "", err2
	}

	dec := json.NewDecoder(resp.Body)
	var r fnResponse

	err := dec.Decode(&r)
	return r.ReturnValue, err
}

func NewCore(id, key string) (Core, error) {
	req, err := http.NewRequest("GET", API_ROOT+API_DEVICES+id, nil)

	//TODO: Abstract errors
	if err != nil {
		return Core{}, err
	}

	c := Core{
		Id:  id,
		key: key,
	}

	resp, reqErr := c.do(req)

	if reqErr != nil {
		return Core{}, reqErr
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
