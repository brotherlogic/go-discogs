package godiscogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"time"
)

type jsonUnmarshaller interface {
	Unmarshal([]byte, interface{}) error
}
type prodUnmarshaller struct{}

func (jsonUnmarshaller prodUnmarshaller) Unmarshal(inp []byte, v interface{}) error {
	err := json.Unmarshal(inp, v)
	return err
}

var httpCount int

// GetHTTPGetCount The number of http gets performed
func GetHTTPGetCount() int {
	return httpCount
}

type httpGetter interface {
	Get(url string) (*http.Response, error)
	Post(url string, data string) (*http.Response, error)
	Put(url string, data string) (*http.Response, error)
	Delete(url string, data string) (*http.Response, error)
}

// DiscogsRetriever Main retriever type
type DiscogsRetriever struct {
	userAgent        string
	lastRetrieveTime int64
	userToken        string
	unmarshaller     jsonUnmarshaller
	getter           httpGetter
	getSleep         int
	logger           func(string)
}

// NewDiscogsRetriever Build a production retriever
func NewDiscogsRetriever(token string, logger func(string)) *DiscogsRetriever {
	return &DiscogsRetriever{unmarshaller: prodUnmarshaller{}, getter: prodHTTPGetter{}, userToken: token, getSleep: 2000, lastRetrieveTime: time.Now().Unix(), logger: logger}
}

var lastTimeRetrieved time.Time

// Urls list of urls in pagination
type Urls struct {
	Next string
}

// Pagination the pagination structure
type Pagination struct {
	Pages int
	Page  int
	Urls  Urls
}

// CollectionResponse returned from discogs
type CollectionResponse struct {
	Pagination Pagination
	Releases   []*CollectionRelease
}

// CollectionRelease returned from collection pull
type CollectionRelease struct {
	ID         int `json:"id"`
	FolderID   int `json:"folder_id"`
	InstanceID int `json:"instance_id"`
	Rating     int `json:"rating"`
	Notes      []*Note
}

// WantlistResponse returned from discogs
type WantlistResponse struct {
	Pagination Pagination
	Wants      []*Release
}

// Version a version of a master release
type Version struct {
	Released string
}

// VersionsResponse returned from discogs
type VersionsResponse struct {
	Pagination Pagination
	Versions   []Version
}

// Pricing the single price
type Pricing struct {
	Currency string
	Value    float32
}

// PriceResponse response from get sale details
type PriceResponse struct {
	Price  Pricing
	Status string
}

// SellResponse response from selling a record
type SellResponse struct {
	ListingID int `json:"listing_id"`
}

// GetRateLimit returns the rate limit
func (r *DiscogsRetriever) GetRateLimit() int {
	_, headers, _ := r.retrieve("/releases/249504?token=" + r.userToken)
	val, _ := strconv.Atoi(headers.Get("X-Discogs-Ratelimit"))
	return val
}

func (r *DiscogsRetriever) processCollectionRelease(re *CollectionRelease) *Release {
	release := &Release{}
	for _, note := range re.Notes {
		// Media condition
		if note.FieldId == 1 {
			release.RecordCondition = note.Value
		}

		// Sleeve condition
		if note.FieldId == 2 {
			release.SleeveCondition = note.Value
		}
	}

	release.Id = int32(re.ID)
	release.InstanceId = int32(re.InstanceID)
	release.FolderId = int32(re.FolderID)
	release.Rating = int32(re.Rating)

	return release
}

// GetCollection gets all the releases in the users collection
func (r *DiscogsRetriever) GetCollection() []*Release {
	jsonString, _, _ := r.retrieve("/users/brotherlogic/collection/folders/0/releases?per_page=100&token=" + r.userToken)

	var releases []*CollectionRelease
	response := &CollectionResponse{}
	r.unmarshaller.Unmarshal(jsonString, response)

	releases = append(releases, response.Releases...)
	end := response.Pagination.Pages == response.Pagination.Page

	r.Log(fmt.Sprintf("FOUND %v PAGES", response.Pagination.Pages))

	for !end {
		jsonString, _, _ = r.retrieve(response.Pagination.Urls.Next[23:])
		newResponse := &CollectionResponse{}
		r.unmarshaller.Unmarshal(jsonString, &newResponse)

		releases = append(releases, newResponse.Releases...)
		end = newResponse.Pagination.Pages == newResponse.Pagination.Page
		response = newResponse
	}

	procReleases := []*Release{}
	for _, re := range releases {
		procReleases = append(procReleases, r.processCollectionRelease(re))
	}

	return procReleases
}

// InventoryResponse returned from discogs
type InventoryResponse struct {
	Pagination Pagination
	Listings   []*InventoryEntry
}

// BasicRelease is the release returned from the inventory pull
type BasicRelease struct {
	ID int
}

// InventoryEntry returned from invetory pull
type InventoryEntry struct {
	Price   Pricing
	ID      int `json:"id"`
	Release BasicRelease
}

// GetInventory gets all the releases that are for sale
func (r *DiscogsRetriever) GetInventory() ([]*ForSale, error) {
	jsonString, _, err := r.retrieve("/users/brotherlogic/inventory?status=For%20Sale&per_page=100&token=" + r.userToken)
	if err != nil {
		return []*ForSale{}, err
	}

	var items []*InventoryEntry
	response := &InventoryResponse{}
	r.unmarshaller.Unmarshal(jsonString, response)

	items = append(items, response.Listings...)
	end := response.Pagination.Pages == response.Pagination.Page

	for !end {
		jsonString, _, err = r.retrieve(response.Pagination.Urls.Next[23:])
		if err != nil {
			return []*ForSale{}, err
		}
		newResponse := &InventoryResponse{}
		r.unmarshaller.Unmarshal(jsonString, &newResponse)

		items = append(items, newResponse.Listings...)
		end = newResponse.Pagination.Pages == newResponse.Pagination.Page
		response = newResponse
	}

	sales := []*ForSale{}
	for _, re := range items {
		sales = append(sales, &ForSale{Id: int32(re.Release.ID), SaleId: int32(re.ID), SalePrice: int32(math.Round(float64(re.Price.Value * 100)))})
	}

	return sales, nil
}

// GetSalePrice gets the sale price for a release
func (r *DiscogsRetriever) GetSalePrice(releaseID int) (float32, error) {
	jsonString, _, err := r.retrieve("/marketplace/price_suggestions/" + strconv.Itoa(releaseID) + "?token=" + r.userToken)
	var resp map[string]Pricing
	r.unmarshaller.Unmarshal(jsonString, &resp)
	value := resp["Mint (M)"].Value

	if value == 0 && err == nil {
		return float32(100), err
	}
	return value, err
}

// SellRecord sells a given release
func (r *DiscogsRetriever) SellRecord(releaseID int, price float32, state string, condition, sleeve string) int {
	data := "{\"release_id\":" + strconv.Itoa(releaseID) + ", \"condition\":\"" + condition + "\", \"sleeve_condition\":\"" + sleeve + "\", \"price\":" + strconv.FormatFloat(float64(price), 'g', -1, 32) + ", \"status\":\"" + state + "\",\"weight\":\"auto\"}"
	databack, _ := r.post("/marketplace/listings?token="+r.userToken, data)
	var resp SellResponse
	r.unmarshaller.Unmarshal([]byte(databack), &resp)
	return resp.ListingID
}

// GetCurrentSalePrice gets the current sale price
func (r *DiscogsRetriever) GetCurrentSalePrice(saleID int) float32 {
	jsonString, _, _ := r.retrieve("/marketplace/listings/" + strconv.Itoa(saleID) + "?curr_abbr=USD&token=" + r.userToken)
	var resp PriceResponse
	r.unmarshaller.Unmarshal(jsonString, &resp)
	return resp.Price.Value
}

// GetCurrentSaleState gets the current sale state
func (r *DiscogsRetriever) GetCurrentSaleState(saleID int) SaleState {
	jsonString, _, _ := r.retrieve("/marketplace/listings/" + strconv.Itoa(saleID) + "?curr_abbr=USD&token=" + r.userToken)
	var resp PriceResponse
	r.unmarshaller.Unmarshal(jsonString, &resp)

	if resp.Status == "For Sale" {
		return SaleState_FOR_SALE
	} else if resp.Status == "Sold" || resp.Status == "Draft" || resp.Status == "Deleted" {
		return SaleState_SOLD
	}

	r.Log(fmt.Sprintf("Unknown sale status: %v", resp.Status))
	return SaleState_NOT_FOR_SALE
}

// UpdateSalePrice updates the sale price
func (r *DiscogsRetriever) UpdateSalePrice(saleID int, releaseID int, condition, sleeve string, price float32) error {
	data := "{\"release_id\":" + strconv.Itoa(releaseID) + ", \"condition\":\"" + condition + "\", \"sleeve_condition\":\"" + sleeve + "\", \"price\":" + strconv.FormatFloat(float64(price), 'g', -1, 32) + ", \"status\":\"For Sale\"}"
	_, err := r.post("/marketplace/listings/"+strconv.Itoa(saleID)+"?curr_abr=USD&token="+r.userToken, data)
	return err
}

// RemoveFromSale removes the listing from sale
func (r *DiscogsRetriever) RemoveFromSale(saleID int, releaseID int) error {
	data := "{\"release_id\":" + strconv.Itoa(releaseID) + ", \"condition\":\"Near Mint (NM or M-)\", \"price\":5.00, \"status\":\"Draft\"}"
	_, err := r.post("/marketplace/listings/"+strconv.Itoa(saleID)+"?curr_abr=USD&token="+r.userToken, data)
	return err
}

// AddToWantlist adds a record to the wantlist
func (r *DiscogsRetriever) AddToWantlist(releaseID int) {
	r.put("/users/brotherlogic/wants/"+strconv.Itoa(releaseID)+"?token="+r.userToken, "")
}

// RemoveFromWantlist adds a record to the wantlist
func (r *DiscogsRetriever) RemoveFromWantlist(releaseID int) {
	r.delete("/users/brotherlogic/wants/"+strconv.Itoa(releaseID)+"?token="+r.userToken, "")
}

//AddToFolderResponse the response back from an add request
type AddToFolderResponse struct {
	InstanceID  int `json:"instance_id"`
	ResourceURL string
	Simple      int
}

// MoveToFolder Moves the given release to the new folder
func (r *DiscogsRetriever) MoveToFolder(folderID int, releaseID int, instanceID int, newFolderID int) string {
	val, _ := r.post("/users/brotherlogic/collection/folders/"+strconv.Itoa(folderID)+"/releases/"+strconv.Itoa(releaseID)+"/instances/"+strconv.Itoa(instanceID)+"?token="+r.userToken, "{\"folder_id\": "+strconv.Itoa(newFolderID)+"}")
	return val
}

// DeleteInstance removes a record from the collection
func (r *DiscogsRetriever) DeleteInstance(folderID int, releaseID int, instanceID int) string {
	return r.delete("/users/brotherlogic/collection/folders/"+strconv.Itoa(folderID)+"/releases/"+strconv.Itoa(releaseID)+"/instances/"+strconv.Itoa(instanceID)+"?token="+r.userToken, "")
}

// FoldersResponse returned from discogs
type FoldersResponse struct {
	Pagination Pagination
	Folders    []Folder
}

// GetFolders gets all the folders for a given user
func (r *DiscogsRetriever) GetFolders() []Folder {
	jsonString, _, _ := r.retrieve("/users/brotherlogic/collection/folders?token=" + r.userToken)

	var folders []Folder
	var response FoldersResponse
	r.unmarshaller.Unmarshal(jsonString, &response)

	folders = append(folders, response.Folders...)

	return folders
}

func (r *DiscogsRetriever) throttle() {
	//Sleep here
	diff := time.Now().Sub(lastTimeRetrieved)
	if diff < time.Duration(r.getSleep)*time.Millisecond {
		time.Sleep(time.Duration(r.getSleep)*time.Millisecond - diff)
	}
	lastTimeRetrieved = time.Now()
}

func (r *DiscogsRetriever) retrieve(path string) ([]byte, http.Header, error) {
	urlv := "https://api.discogs.com/" + path

	r.throttle()
	response, err := r.getter.Get(urlv)
	if err != nil {
		return make([]byte, 0), make(http.Header), err
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 && response.StatusCode != 201 && response.StatusCode != 204 {
		r.Log(fmt.Sprintf("RETR (%v) %v -> %v", path, response.StatusCode, string(body)))
	}

	lastTimeRetrieved = time.Now()

	return body, response.Header, nil
}
