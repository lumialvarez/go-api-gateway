package httpclient

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func HttpPassThrough(urlBase string, ctx *gin.Context) error {
	path := ctx.Param("proxyPath")
	client := &http.Client{}
	request, err := http.NewRequest(ctx.Request.Method, urlBase+path, ctx.Request.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	//fmt.Printf("%s", body)
	for name, values := range response.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Println(name, value)
			ctx.Header(name, value)
		}
	}
	ctx.Data(http.StatusOK, response.Header.Get("Content-Type"), body)
	return nil
}
