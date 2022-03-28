package helper

import (
	"context"
	"fmt"
	"testing"
)

// more test case see: https://olivere.github.io/elastic/
func Test_Elastic(t *testing.T) {
	ctx := context.Background()

	t.Run("ping", func(t *testing.T) {
		info, code, err := ESClient.Ping("http://172.16.254.219:39200").Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	})
}
