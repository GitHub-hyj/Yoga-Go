package yogaclient

import (
	"github.com/GitHub-hyj/Yoga-Go/internal/yogaconfig"
	"github.com/GitHub-hyj/Yoga-Go/internal/HttpRequest"
	"fmt"
)


func Login(urlStr string, params map[string]string) *yogaconfig.User {

	data, err := HttpRequest.Post(urlStr,params)
	if err != nil {
		panic(err)
	}
	fmt.Println("data=",data)

	return nil

}