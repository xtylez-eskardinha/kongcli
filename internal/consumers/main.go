package consumers

import (
	"context"
	"errors"

	"github.com/kong/go-kong/kong"
)

func AddConsumer(ctx context.Context, client *kong.Client, username string, customid string, tags []*string) (*kong.Consumer, error) {
	newConsumer := &kong.Consumer{
		Username: &username,
		CustomID: &customid,
		Tags:     tags,
	}

	createdConsumer, err := client.Consumers.Create(ctx, newConsumer)

	if err != nil {
		return nil, err
	}

	return createdConsumer, nil
}

// func ListAllConsumers(ctx context.Context) ([]*kong.Consumer, error) {
// 	return consumerService.ListAll(ctx)
// }

// func ListConsumers(ctx context.Context, size int, offset string, tags []*string, matchall bool) ([]*kong.Consumer, *kong.ListOpt, error) {
// 	opts := &kong.ListOpt{
// 		Size:         size,
// 		Offset:       offset,
// 		Tags:         tags,
// 		MatchAllTags: matchall,
// 	}
// 	return consumerService.List(ctx, opts)
// }

func GetConsumer(ctx context.Context, client *kong.Client, consumer string, custom string) (*kong.Consumer, error) {
	if consumer != "" {
		return client.Consumers.Get(ctx, &consumer)
	}
	if custom != "" {
		return client.Consumers.GetByCustomID(ctx, &custom)
	}
	return nil, errors.New("No consumer or custom id given...")
}
