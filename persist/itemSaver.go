package persist

import (
	"context"
	"errors"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"spider/engine"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)

	go func() {
		for {
			item := <-out
			log.Printf("item : %v", item)
			err := save(item)
			if err != nil {
				log.Printf("item saver err : %v item : %v", err, item)
			}
		}
	}()
	return out
}

func save(item engine.Item) (err error) {
	client, err := elastic.NewClient(
		// must set false in docker,维护集群状态的，docker不需要
		elastic.SetSniff(false),
	)

	if err != nil {
		return err
	}

	if item.Type == "" {
		return errors.New("item type empty")
	}

	indexService := client.Index().
		Index("dating_profile").
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
