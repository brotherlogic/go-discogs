package discogsgo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type jsonUnmarshaller interface {
	Unmarshal([]byte, interface{}) error
}
type prodUnmarshaller struct{}

func (jsonUnmarshaller prodUnmarshaller) Unmarshal(inp []byte, v interface{}) error {
	return json.Unmarshal(inp, v)
}

type httpGetter interface {
     Get(url string) (*http.Response, error)
     }
type prodHTTPGetter struct{}
func (httpGetter prodHTTPGetter) Get(url string) (*http.Response, error) {
  return http.Get(url)
}

// DiscogsRetriever Main retriever type
type DiscogsRetriever struct {
	userAgent        string
	lastRetrieveTime int64
	userToken        string
	unmarshaller     jsonUnmarshaller
	getter		 httpGetter	 
}

// NewDiscogsRetriever Build a production retriever
func NewDiscogsRetriever() *DiscogsRetriever {
	return &DiscogsRetriever{unmarshaller: prodUnmarshaller{}, getter: prodHTTPGetter{}}
}

// Release a release in the discogs sense
type Release struct {
	ID    int
	Title string
}

// GetRelease returns a release from the discogs system
func (r *DiscogsRetriever) GetRelease(id int) (Release, error) {
	jsonString, _ := r.retrieve("/releases/" + strconv.Itoa(id))
	log.Printf("Returned %v", string(jsonString))
	var release Release
	err := r.unmarshaller.Unmarshal(jsonString, &release)

	if err != nil {
		return release, err
	}

	return release, nil
}


func (r *DiscogsRetriever) retrieve(path string) ([]byte, error) {
	urlv := "https://api.discogs.com/" + path
	response, err := r.getter.Get(urlv)

	if err != nil {
		return make([]byte, 0), err
	}
	
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		return body, nil
}
