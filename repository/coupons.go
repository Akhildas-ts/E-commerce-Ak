package repository

import (
	"ak/database"
	"ak/helper"
	"ak/models"
	errorss "ak/repository/errors"
	"errors"
	"time"

	"fmt"
)

func CouponExist(couponName string) (bool, error) {

	var count int

	err := database.DB.Raw("select count(*) from coupons where coupon =?", couponName).Scan(&count).Error
	if err != nil {
		
		return false, err
	}

	return count > 0, nil
}

func CouponRevalidateIfExpired(couponName string) (bool, error) {

	var isvalid bool
	err := database.DB.Raw("select validity from coupons where coupon = ?", couponName).Scan(&isvalid).Error

	if err != nil {

		return false, err
	}

	if isvalid {
		fmt.Println("is valid", err)
		return true, err

	}

	err = database.DB.Exec("update coupons set validity = true where coupon = ?", couponName).Error

	if err != nil {

		return false, err
	}

	return false, err
}

func AddCoupon(coupon models.AddCoupon) error {
	err := database.DB.Exec("insert into coupons (coupon,discount_percentage,minimum_price,validity)values(?,?,?,?)", coupon.Coupon, coupon.DiscountPercentage, coupon.MinimumPrice, coupon.Validity)

	if err != nil {

		fmt.Println("error from add coupon", err)
		return nil
	}

	return nil

}

func GetCoupon() ([]models.Coupon, error) {

	var coupons []models.Coupon

	err := database.DB.Raw("select id,coupon,discount_percentage,minimum_price,validity from coupons").Scan(&coupons).Error

	if err != nil {
		return []models.Coupon{}, err
	}

	return coupons, nil
}

func ExistCoupon(couponID int) (bool, error) {

	var count int
	err := database.DB.Raw("select count(*) from coupons where id = ?", couponID).Scan(&count).Error
	if err != nil {
		return false, errorss.ErrCouponAlreadyexist

	}

	return count > 0, nil
}

func CouponAlreadyExpired(couponID int) error {

	var valid bool

	err := database.DB.Raw("select validity from coupons where id=?", couponID).Scan(&valid).Error

	if err != nil {

		return err
	}

	if valid {

		err := database.DB.Exec("update coupons set validity = false where id=?", couponID).Error

		if err != nil {

			return err
		}

		return nil
	}

	return errors.New("alredy expired")
}

func CouponValidity(couponName string) (bool, error) {

	var validity bool

	err := database.DB.Raw("select validity from coupons where coupon =?", couponName).Scan(&validity).Error

	if err != nil {

		return false, err
	}

	return validity, nil
}

func GetCouponMinimumAmount(coupon string) (float64, error) {

	var MinDiscountPrice float64
	err := database.DB.Raw("select minimum_price from coupons where coupon = ?", coupon).Scan(&MinDiscountPrice).Error
	if err != nil {
		return 0.0, err
	}
	return MinDiscountPrice, nil
}

func DidUserAlreadyUsedThisCoupon(coupon string, userid int) (bool, error) {

	var count int

	err := database.DB.Raw("select count(*) from used_coupons where coupon_id = (select id from coupons where coupon =?)and user_id=?", coupon, userid).Scan(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil

}

func UpdateUsedCoupon(coupon string, userID int) (bool, error) {

	var couponID uint
	err := database.DB.Raw("select id from coupons where coupon = ?", coupon).Scan(&couponID).Error
	if err != nil {
		return false, err
	}

	var count int
	// if a coupon have already been added, replace the order with current coupon and delete the existing coupon
	err = database.DB.Raw("select count(*) from used_coupons where user_id = ? and used = false", userID).Scan(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		err = database.DB.Exec("delete from used_coupons where user_id = ? and used = false", userID).Error
		if err != nil {
			return false, err
		}
	}

	err = database.DB.Exec("insert into used_coupons (coupon_id,user_id,used) values (?, ?, false)", couponID, userID).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
func UpdateCouponDetails(discountPrice float64,userId int) error {

	if discountPrice != 0.0 {

		err := database.DB.Exec(" update used_coupons set used = true where user_id = ?",userId).Error

		if err != nil {
			return err
		}
	}
	return nil
}

func AddCategoryOffer(categoryOffer models.CategoryOfferReceiver) error {

	if categoryOffer.OfferName == ""{
		return errors.New("offer name can't be  nil")
	}

	if categoryOffer.DiscountPercentage < 0 {
		return errors.New("discount percentage is less than zero")

	}

	if categoryOffer.OfferLimit <+ 0 {
		return errors.New("offer limit must be greater than zero ")
	}

	// check if the offer with the offer name already exist in the database
	var count int
	err := database.DB.Raw("select count(*) from category_offers where offer_name = ?", categoryOffer.OfferName).Scan(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("the offer already exists")
	}

	// if there is any other offer for this category delete that before adding this one
	count = 0
	err = database.DB.Raw("select count(*) from category_offers where category_id = ?", categoryOffer.CategoryID).Scan(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {

		err = database.DB.Exec("delete from category_offers where category_id = ?", categoryOffer.CategoryID).Error
		if err != nil {
			return err
		}
	}

	startDate := time.Now()
	endDate := time.Now().Add(time.Hour * 24 * 5)
	err = database.DB.Exec("INSERT INTO category_offers (category_id, offer_name, discount_percentage, start_date, end_date, offer_limit,offer_used) VALUES (?, ?, ?, ?, ?, ?, ?)", categoryOffer.CategoryID, categoryOffer.OfferName, categoryOffer.DiscountPercentage, startDate, endDate, categoryOffer.OfferLimit, 0).Error
	if err != nil {
		return err
	}

	return nil

}

func GetCouponDiscountPrice(UserID int, GrandTotal float64) (float64, error) {

	discountprice, err := helper.GetCouponDiscountPrice(UserID, GrandTotal)

	if err != nil {
		return 0.0, err
	}
	return discountprice, nil
}
