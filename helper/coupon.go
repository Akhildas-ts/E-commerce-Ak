package helper

import (
	"ak/database"
	"fmt"
)

func GetCouponDiscountPrice(userID int, grandTotal float64) (float64, error) {

	var count int

	err := database.DB.Raw("select count(*) from  used_coupons where user_id = ? and used = false",userID).Scan(&count).Error

	if err != nil {
		return 0.0,err
	}

	if count <0{
		return 0.0,err
	}

	type CouponDetails struct{
		DiscountPercentage int 
		MinimumPrice float64
	}


	var coupD CouponDetails

	err = database.DB.Raw("select discount_percentage,minimum_price from coupons where id = (select coupon_id from used_coupons where user_id = ? and  used = false)",userID).Scan(&coupD).Error

	if err != nil {

		return 0.0,err
	}

	var totalPrice float64
	err = database.DB.Raw("select COALESCE(SUM(total_price), 0) from carts where user_id = ?", userID).Scan(&totalPrice).Error
	if err != nil {
		return 0.0, err
	}
	fmt.Println("coupd discount price",coupD.DiscountPercentage)
	fmt.Println("coupd  minimumprice",coupD.MinimumPrice)
	fmt.Println("total price",totalPrice)


	if totalPrice < coupD.MinimumPrice {

		return 0.0,err
	}

	return ((float64(coupD.DiscountPercentage) * totalPrice) / 100), nil


}