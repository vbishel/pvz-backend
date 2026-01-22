package v1

import (
	"context"
	"orders-service/internal/domain/acceptance"
	"orders-service/internal/domain/category"
	"orders-service/internal/domain/city"
	"orders-service/internal/domain/pickuppoint"
	"orders-service/internal/domain/product"
	"orders-service/internal/domain/status"
)

type PupService interface {
	Create(ctx context.Context, cityID city.CityID, roleID int) (pickuppoint.PickupPointID, error)
	Find(ctx context.Context, id pickuppoint.PickupPointID) (pickuppoint.PickupPoint, error)
}

type ProductService interface {
	Create(ctx context.Context, name string, categoryID category.CategoryID) (product.ProductID, error)
}

type AcceptanceService interface {
	Create(ctx context.Context, pupID pickuppoint.PickupPointID) (acceptance.AcceptanceID, error)
	ChangeStatus(ctx context.Context, statusID status.StatusID) (error)
	AddProduct(ctx context.Context, productID product.ProductID) (error)
}

type CityService interface {
	Create(ctx context.Context, name string) (city.CityID, error)
}
