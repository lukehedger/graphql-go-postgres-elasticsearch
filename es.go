package api

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

type ES struct {
	*elasticsearch.Client
}

func OpenES() (*ES, error) {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}

	req := esapi.IndexRequest{
		Index:      "people",
		DocumentID: "1",
		Body:       strings.NewReader(`{"id": "1", "name": "Luke"}`),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.IsError() {
		return nil, err
	}
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to elasticsearch!")

	return &ES{es}, nil
}
