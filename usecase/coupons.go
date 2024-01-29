package usecase

import (
	"ak/models"
	"ak/repository"
	"errors"
	"fmt"
)

func AddCoupon(coupon models.AddCoupon) (string, error) {

	if coupon.DiscountPercentage < 0 {
		return "",errors.New("discount price is less than zero")
	}

	if coupon.MinimumPrice < 0 {
		return "",errors.New("minum price is less than zero")
	}

	couponExist, err := repository.CouponExist(coupon.Coupon)

	if err != nil {

		return "", err
	}

	if couponExist {

		alreadyExist, err := repository.CouponRevalidateIfExpired(coupon.Coupon)

		if err != nil {

			return "", err
		}

		if alreadyExist {

			fmt.Println("error from alredyexist", err)

			return "coupon have alredy exist !!", err

		}

		return "made the coupon valid", nil
	}

	err = repository.AddCoupon(coupon)

	if err != nil {
		return "", err
	}

	return "successfully added the coupon", nil

}

func GetCoupon() ([]models.Coupon,error) {

	return repository.GetCoupon()
}

func ExpireCoupon(couponID int)error {

	couponExist	,err := repository.ExistCoupon(couponID)

	if err != nil {

		return err
	}

	if couponExist{
		err = repository.CouponAlreadyExpired(couponID)

		if err != nil{

			return err
		}

		return nil
	}

	return errors.New("coupon does not exist")


}

func AddCategoryOffer(categoryOffer models.CategoryOfferReceiver) error {

	return repository.AddCategoryOffer(categoryOffer)

}
