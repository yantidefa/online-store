package paymentservice

import (
	"online-store/models"
	cartrepository "online-store/repositories/cart_repository"
	paymentrepository "online-store/repositories/payment_repository"
	"online-store/utilities"
)

func CreatePayment(request models.Payment) (*models.Payment, error) {
	request.CreatedAt = utilities.Times()

	getCart, err := cartrepository.GetCart("", request.ProductId.String(), request.UserId.String())
	if err != nil {
		return nil, err
	}

	if len(getCart) != 0 {
		for i := 0; i < len(getCart); i++ {
			if getCart[i].Qty < request.Qty {
				getCart[i].Qty -= request.Qty
				_, err := cartrepository.UpdateCartById(*getCart[i])
				if err != nil {
					return nil, err
				}
			} else if getCart[i].Qty == request.Qty {
				_, err = cartrepository.DeleteCartById(getCart[i].ID.String(), request.CreatedAt)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	createPayment, err := paymentrepository.CreatePayment(request)
	if err != nil {
		return createPayment, err
	}
	
	return createPayment, nil
}

func GetPayment(PaymentId, productId, userId string) ([]*models.Payment, error) {
	getPayment, err := paymentrepository.GetPayment(PaymentId, productId, userId)
	if err != nil {
		return getPayment, err
	}
	return getPayment, nil
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

func UpdatePaymentById(request models.Payment) (int64, error) {
	request.UpdatedAt = utilities.Times()
	deleteProduct, err := paymentrepository.UpdatePaymentById(request)
	if err != nil {
		return 0, err
	}
	return deleteProduct, nil
}
