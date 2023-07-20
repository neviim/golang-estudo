package usecase

import "github.com/neviim/golang-estudo/internal/entity"

// dados que vai entrar no sistema
type OrderInput struct {
	ID    string
	Price float64
	Tax   float64
}

// dados que vai sair do sistema
type OrderOutput struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

// SOLID - "D" - Dependency Inversion | este é o caso de uso
type CalculateFinalPrice struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCalculateFinalPrice(OrderRepository entity.OrderRepositoryInterface) *CalculateFinalPrice {
	return &CalculateFinalPrice{
		OrderRepository: OrderRepository,
	}
}

func (c *CalculateFinalPrice) Execute(input OrderInput) (*OrderOutput, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}
	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}
	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}
	return &OrderOutput{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
