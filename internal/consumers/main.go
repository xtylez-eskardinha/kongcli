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

func ListConsumers(ctx context.Context, client *kong.Client) ([]*kong.Consumer, error) {
	return client.Consumers.ListAll(ctx)
}

func DeleteConsumer(ctx context.Context, client *kong.Client, consumer string) error {
	return client.Consumers.Delete(ctx, &consumer)
}

func UpdateConsumer(ctx context.Context, client *kong.Client, consumer *kong.Consumer) (*kong.Consumer, error) {
	return client.Consumers.Update(ctx, consumer)
}
func ListConsumersFiltered(ctx context.Context, client *kong.Client, tags []*string) ([]*kong.Consumer, error) {
	consumers, _, err := client.Consumers.List(ctx, &kong.ListOpt{Tags: tags})
	return consumers, err
}

// GetConsumer returns a consumer by name or custom id
func GetConsumer(ctx context.Context, client *kong.Client, consumer string, custom string) (*kong.Consumer, error) {
	if consumer != "" {
		return client.Consumers.Get(ctx, &consumer)
	}
	if custom != "" {
		return client.Consumers.GetByCustomID(ctx, &custom)
	}
	return nil, errors.New("No consumer or custom id given...")
}
