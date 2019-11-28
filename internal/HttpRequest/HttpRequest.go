package HttpRequest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Post(urlStr string, params map[string]string)  (map[string]interface{}, error) {

	client := &http.Client{}
	query := url.Values{}
	for k := range params {
		query.Set(k, params[k])
	}
	obody := strings.NewReader(query.Encode())

	urlStr = "http://127.0.0.1:6010"+urlStr

	req, err := http.NewRequest("POST", urlStr, obody)
	if err!= nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()



	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}


	if data, ok := response["data"].(map[string]interface{}); ok {
		return data,nil
	}
	return nil, nil
}
