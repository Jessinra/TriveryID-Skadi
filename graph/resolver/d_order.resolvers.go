package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gitlab.com/trivery-id/skadi/graph/generated"
	"gitlab.com/trivery-id/skadi/graph/model"
	"gitlab.com/trivery-id/skadi/utils/ptr"
)

func (r *orderResolver) Requester(ctx context.Context, obj *model.Order) (*model.User, error) {
	user, err := userService.GetUser(ctx, obj.RequesterID)
	if err != nil {
		return nil, err
	}

	return model.NewUser(user), nil
}

func (r *orderResolver) Shopper(ctx context.Context, obj *model.Order) (*model.User, error) {
	if id := obj.ShopperID; id == nil || id == ptr.Uint64(0) {
		return nil, nil
	}

	user, err := userService.GetUser(ctx, *obj.ShopperID)
	if err != nil {
		return nil, err
	}

	return model.NewUser(user), nil
}

func (r *orderResolver) Product(ctx context.Context, obj *model.Order) (*model.Product, error) {
	product, err := productService.GetProduct(ctx, obj.ProductID)
	if err != nil {
		return nil, err
	}

	return model.NewProduct(product), nil
}

// Order returns generated.OrderResolver implementation.
func (r *Resolver) Order() generated.OrderResolver { return &orderResolver{r} }

type orderResolver struct{ *Resolver }
