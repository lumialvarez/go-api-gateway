package httpclient

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func HttpPassThrough(urlBase string, ctx *gin.Context) error {
	path := ctx.Param("proxyPath")
	client := &http.Client{}
	request, err := http.NewRequest(ctx.Request.Method, urlBase+path, ctx.Request.Body)
	for name, values := range ctx.Request.Header {
		// Loop over all values for the name.
		for _, value := range values {
			request.Header.Set(name, value)
		}
	}
	if err != nil {
		return err
	}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	//Ignoring error in case when response is nil
	defer func() {
		_ = response.Body.Close()
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	//fmt.Printf("%s", body)
	for name, values := range response.Header {
		// Loop over all values for the name.
		for _, value := range values {
			ctx.Header(name, value)
		}
	}
	ctx.Data(http.StatusOK, response.Header.Get("Content-Type"), body)
	return nil
}
