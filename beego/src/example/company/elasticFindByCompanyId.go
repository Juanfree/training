package company

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	"os"
)

type ElasticFindByCompanyId struct {
}

func (e *ElasticFindByCompanyId) ByCompanyId(id int) ([]*CompanyDto, int64) {

	elasticUrl := os.Getenv("ELASTIC_URL")
	elasticUrl = "http://192.168.99.100:9200"
	options := elastic.SetURL(elasticUrl)
	client, err := elastic.NewClient(options, elastic.SetSniff(false))
	defer client.Stop()

	if err != nil {
		panic(err)
	}

	query := elastic.NewBoolQuery()
	musts := []elastic.Query{elastic.NewTermQuery("company_id", id)}
	query = query.Must(musts...)

	searchResult, err := client.
		Search().
		Index("training").
		Type("test_type").
		Query(query).
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
