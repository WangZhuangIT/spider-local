package persist

import (
	"context"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"spider/engine"
)

func ItemSaver(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)

	client, err := elastic.NewClient(
		// must set false in docker,维护集群状态的，docker不需要
		elastic.SetSniff(false),
	)

	if err != nil {
		return nil, err
	}

	go func() {
		for {
			item := <-out
			log.Printf("item : %v", item)
			err := save(item, client, index)
			if err != nil {
				log.Printf("item saver err : %v item : %v", err, item)
			}
		}
	}()
	return out, nil
}

func save(item engine.Item, client *elastic.Client, index string) (err error) {
	if item.Type == "" {
		return errors.New("item type empty")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.Do(context.Background())

	if err != nil {
		return err
	}
	return nil
}
