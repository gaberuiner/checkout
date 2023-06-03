package swagger

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ProductService struct {
	Token string
}

type RequestProduct struct {
	Token string `json:"token"`
	Sku   uint32 `json:"sku"`
}

type ResponseProduct struct {
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

func (p *ProductService) GetItems(sku uint32) (ResponseProduct, error) {
	request := RequestProduct{Token: p.Token, Sku: sku}
	reqBody, err := json.Marshal(request)
	if err != nil {
		return ResponseProduct{}, err
	}
	url := "http://route256.pavl.uk:8080/get_product"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return ResponseProduct{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ResponseProduct{}, err
	}
	defer resp.Body.Close()
	var productResponse ResponseProduct
	err = json.NewDecoder(resp.Body).Decode(&productResponse)
	if err != nil {
		return ResponseProduct{}, err
	}

	return productResponse, nil
}

func NewProductService(token string) *ProductService {
	return &ProductService{
		Token: token,
	}
}
