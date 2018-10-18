package company

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
)

type ElasticFindAllCompanies struct {
}

func (e *ElasticFindAllCompanies) All() ([]*CompanyDto, int64) {

	elastic_url := "http://192.168.99.100:9200"
	options := elastic.SetURL(elastic_url)
	client, err := elastic.NewClient(options, elastic.SetSniff(false))
	defer client.Stop()

	if err != nil {
		panic(err)
	}

	searchResult, err := client.
		Search().
		Index("training").
		Type("test_type").
		Query(elastic.NewMatchAllQuery()).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	totalHits := searchResult.Hits.TotalHits
	if totalHits <= 0 {
		var notResults []*CompanyDto
		return notResults, 0
	}

	var l []*CompanyDto
	for _, hit := range searchResult.Hits.Hits {
		var c *CompanyDto
		err := json.Unmarshal(*hit.Source, &c)
		if err != nil {
			panic(err)
		}

		l = append(l, c)
	}

	return l, totalHits
}
