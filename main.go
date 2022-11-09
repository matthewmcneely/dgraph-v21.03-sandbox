package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

var (
	CIDStartID       = 700
	CIDGroups        = 50
	ProductsPerGroup = 5000
	BrandGroups      = []string{
		"GODEX",
		"AAAA",
		"BBBB",
		"CCCC",
		"DDDD",
		"EEEE",
		"FFFF",
		"GGGG",
		"HHHH",
	}
)

type Product struct {
	SKUID    string  `json:"sku_id_1,omitempty"`
	Title    string  `json:"title_1,omitempty"`
	Brand    string  `json:"brand_1,omitempty"`
	CID3     string  `json:"cid3_1,omitempty"`
	Price    float64 `json:"price_1,omitempty"`
	ClickCnt int64   `json:"click_cnt_1,omitempty"`
	Attrs    string  `json:"attrs_1,omitempty"`
}

// check logs fatal if err != nil.
func check(err error) {
	if err != nil {
		err = errors.Wrap(err, "")
		log.Fatalf("%+v", err)
	}
}

func randPriceInRange(min, max float64) float64 {
	return (max-min)*rand.Float64() + min
}

func createProducts(dg *dgo.Dgraph) {
	op := api.Operation{DropAll: true}
	check(dg.Alter(context.Background(), &op))

	for n := 0; n < CIDGroups; n++ {
		var products []Product
		for m := 0; m < ProductsPerGroup; m++ {
			product := Product{
				SKUID:    fmt.Sprintf("%d-%d", n, m),
				Title:    fmt.Sprintf("Product %d %d", n, m),
				Brand:    BrandGroups[rand.Intn(len(BrandGroups))],
				CID3:     fmt.Sprintf("%d", CIDStartID+n),
				Price:    randPriceInRange(10.0, 3000.0),
				ClickCnt: int64(rand.Intn(500)),
			}
			products = append(products, product)
		}
		data, err := json.Marshal(products)
		check(err)
		txn := dg.NewTxn()

		var mu api.Mutation
		mu.SetJson = data
		resp, err := txn.Mutate(context.Background(), &mu)
		check(err)
		if len(resp.Uids) != ProductsPerGroup {
			log.Fatalf("len of uids %d unexpected", len(resp.Uids))
		}
		check(txn.Commit(context.Background()))
	}
	log.Println("Added", CIDGroups*ProductsPerGroup, "products")
}

func main() {
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	check(err)
	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	createProducts(dg)
}
