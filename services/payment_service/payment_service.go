package paymentservice

import (
	"errors"
	"online-store/models"
	cartrepository "online-store/repositories/cart_repository"
	categoryproductrepository "online-store/repositories/category_product_repository"
	paymentrepository "online-store/repositories/payment_repository"
	productrepository "online-store/repositories/product_repository"
	usersrepository "online-store/repositories/users_repository"
	"online-store/utilities"
)

func CreatePayment(request models.Payment) (*models.Payment, error) {
	request.CreatedAt = utilities.Times()

	products, err := productrepository.GetProduct(request.ProductId.String(), "", "")
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		request.Price = product.Price * request.Qty
		if product.Stock == 0 {
			return nil, errors.New("requested quantity exceeds available quantity in cart")
		}
	}

	cartItems, err := cartrepository.GetCart("", request.ProductId.String(), request.UserId.String())
	if err != nil {
		return nil, err
	}

	for _, cartItem := range cartItems {
		if request.Qty < cartItem.Qty {
			cartItem.Qty -= request.Qty
			if _, err := cartrepository.UpdateCartById(*cartItem); err != nil {
				return nil, err
			}
		} else if request.Qty == cartItem.Qty {
			if _, err := cartrepository.DeleteCartById(cartItem.ID.String(), request.CreatedAt); err != nil {
				return nil, err
			}
		}
	}

	createdPayment, err := paymentrepository.CreatePayment(request)
	if err != nil {
		return nil, err
	}

	requestUpdate := models.Product{
		ID:    request.ProductId,
		Stock: products[0].Stock - request.Qty,
	}
	_, err = productrepository.UpdateProductById(requestUpdate)
	if err != nil {
		return nil, err
	}

	return createdPayment, nil
}

func GetPayment(userId string) (*models.Checkout, error) {
	var checkout models.Checkout
	productPayment := []models.ProductPayment{}

	getPayment, err := paymentrepository.GetPayment("", "", userId)
	if err != nil {
		return nil, err
	}

	datauser, err := usersrepository.GetUser(userId, "")
	if err != nil {
		return nil, err
	}

	userPayment := models.UserPayment{
		UserName: datauser.Name,
		Email:    datauser.Email,
		Address:  datauser.Address,
		Phone:    datauser.Phone,
	}

	for py := 0; py < len(getPayment); py++ {
		products, err := productrepository.GetProduct(getPayment[py].ProductId.String(), "", "")
		if err != nil {
			return nil, err
		}

		for pd := 0; pd < len(products); pd++ {
			category, err := categoryproductrepository.GetCategoryProductById(products[pd].CategoryId.String())
			if err != nil {
				return nil, err
			}

			productPayment = append(productPayment, models.ProductPayment{
				ProductName:  products[pd].Name,
				CategoryName: category.CategoryName,
				Qty:          getPayment[py].Qty,
				Price:        getPayment[py].Price,
			})

		}

		checkout.TotalPrice += getPayment[py].Price
	}

	detailPayment := models.Checkout{
		User:       userPayment,
		Product:    productPayment,
		TotalPrice: checkout.TotalPrice,
	}

	return &detailPayment, nil
}

func DeletePaymentById(id string) (int64, error) {
	request := models.Payment{
		DeletedAt: utilities.Times(),
	}
	deletePayment, err := paymentrepository.DeletePaymentById(id, request.DeletedAt)
	if err != nil {
		return 0, err
	}
	return deletePayment, nil
}
