package main

import (
	"context"
	"log"
	"net/http"
	"os"

	os "github.com/shadowscatcher/shodan/search"
	"github.com/shadowscatcher/shodan/search/ssl_versions"
)

func main() {
	nginxSearch := search.Params{
		Page: 1,
		Query: search.Query{
			Product: "nginx",
			ASN:     "AS14618",
			SSLOpts: search.SSLOpts{
				Cert: search.CertOptions{
					Expired: true,
				},
				Version: ssl_versions.TLSv1_2,
			},
		},
	}

	client, _ := shodan.GetClient(os.Getenv("SHODAN_API_KEY"), http.DefaultClient, true)
	ctx := context.Background()
	result, err := client.Search(ctx, nginxSearch)
	if err != nil {
		log.Fatal(err)
	}

	for _, match := range result.Matches {

		if match.HTTP != nil && match.HTTP.Favicon != nil {
			//newQuery := search.Query{HTTP: search.HTTP{Favicon: search.Favicon{Hash: match.HTTP.Favicon.Hash}}}
		}
	}

	nginxSearch.Page++
	nginxSearch.Query.Port = 443
	result, err = client.Search(ctx, nginxSearch)
	if err != nil {
		log.Fatal(err)
	}
}
