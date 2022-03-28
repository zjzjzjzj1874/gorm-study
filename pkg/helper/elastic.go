package helper

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/olivere/elastic/v7"
)

var ESClient *elastic.Client

func init() {
	httpClient := &http.Client{}
	httpClient.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   10 * time.Second,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		IdleConnTimeout:       90 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	var options []elastic.ClientOptionFunc
	{
		options = append(options, elastic.SetHttpClient(httpClient))
		options = append(options, elastic.SetSniff(false))
		options = append(options, elastic.SetGzip(true))
		options = append(options, elastic.SetHealthcheckInterval(10*time.Second))
		options = append(options, elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)))
		//options = append(options, elastic.SetBasicAuth(userName, password))
		esEndpoint := "http://172.16.254.219:39200" // 本地mock的一个服务
		options = append(options, elastic.SetURL(esEndpoint))
	}

	var err error

	// 如果后面没有endpoint这些,默认的网关是:localhost:9200
	ESClient, err = elastic.NewClient(options...)
	if err != nil {
		panic(err)
	}
}
