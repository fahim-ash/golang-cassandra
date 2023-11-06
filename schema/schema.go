package schema

import (
	"time"
)

type AllProductsPrices struct {
	ID                          int        `json:"id"`
	ProductID                   *int       `json:"product_id,omitempty"`
	Code                        *string    `json:"code,omitempty"`
	Barcode                     *string    `json:"barcode,omitempty"`
	BaseTP                      float64    `json:"base_tp"`
	BaseSalePrice               float64    `json:"base_sale_price"`
	EffectiveSalePrice          float64    `json:"effective_sale_price"`
	ProductSaleDiscount         float64    `json:"product_sale_discount"`
	ProductSaleDiscount2        float64    `json:"product_sale_discount_2"`
	PromotionalDiscount         float64    `json:"promotional_discount"`
	PromotionalDiscount2        float64    `json:"promotional_discount_2"`
	PromotionID                 *int       `json:"promotion_id,omitempty"`
	BatchProductPriceID         *int       `json:"batch_product_price_id,omitempty"`
	CustomerCategoryID          *int       `json:"customer_category_id,omitempty"`
	MembershipCardID            *int       `json:"membership_card_id,omitempty"`
	BranchID                    *int       `json:"branch_id,omitempty"`
	EntryDate                   *time.Time `json:"entry_date,omitempty"`
	EffectiveDate               *time.Time `json:"effective_date,omitempty"`
	EndDate                     *time.Time `json:"end_date,omitempty"`
	StartDate                   *time.Time `json:"start_date,omitempty"`
	Inactive                    bool       `json:"inactive"`
	Deleted                     bool       `json:"deleted"`
	Updated                     *time.Time `json:"updated,omitempty"`
	CreatedBy                   *int       `json:"created_by,omitempty"`
	Created                     *time.Time `json:"created,omitempty"`
	UpdatedBy                   *int       `json:"updated_by,omitempty"`
	VendorID                    *int       `json:"vendor_id,omitempty"`
	MaxMRP                      float64    `json:"max_mrp"`
	MinMRP                      float64    `json:"min_mrp"`
	Variant1                    *int       `json:"variant_1,omitempty"`
	Variant2                    *int       `json:"variant_2,omitempty"`
	Variant3                    *int       `json:"variant_3,omitempty"`
	Variant4                    *int       `json:"variant_4,omitempty"`
	Variant5                    *int       `json:"variant_5,omitempty"`
	Variant6                    *int       `json:"variant_6,omitempty"`
	PromotionalDiscount3Percent *float64   `json:"promotional_discount_3_percent,omitempty"`
	PromotionalDiscount3        *float64   `json:"promotional_discount_3,omitempty"`
	PromotionalDiscount4Percent *float64   `json:"promotional_discount_4_percent,omitempty"`
	PromotionalDiscount4        *float64   `json:"promotional_discount_4,omitempty"`
	PromotionalDiscount5Percent *float64   `json:"promotional_discount_5_percent,omitempty"`
	PromotionalDiscount5        *float64   `json:"promotional_discount_5,omitempty"`
}
