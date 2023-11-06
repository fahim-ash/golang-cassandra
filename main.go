package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"os"
	"time"
)

type CassandraSettings struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Keyspace string `yaml:"keyspace"`
}

type ProductMasterPrices struct {
	ID                          uuid.UUID  `json:"id"`
	ProductID                   *int       `json:"product_id,omitempty"`
	ProductCode                 *string    `json:"product_code,omitempty"`
	Barcode                     *string    `json:"barcode,omitempty"`
	ProductIdentifier           *string    `json:"product_identifier,omitempty"`
	BaseTP                      *float64   `json:"base_tp,omitempty"`
	BaseMRP                     *float64   `json:"base_mrp,omitempty"`
	BranchID                    *int       `json:"branch_id,omitempty"`
	VendorID                    *int       `json:"vendor_id,omitempty"`
	BranchBaseTP                *float64   `json:"branch_base_tp,omitempty"`
	BranchBaseMRP               *float64   `json:"branch_base_mrp,omitempty"`
	PromotionID                 *int       `json:"promotion_id,omitempty"`
	PromotionStartDate          *time.Time `json:"promotion_start_date,omitempty"`
	PromotionEndDate            *time.Time `json:"promotion_end_date,omitempty"`
	PromotionPercentage         *float64   `json:"promotion_percentage,omitempty"`
	PromotionAmount             *float64   `json:"promotion_amount,omitempty"`
	DiscountPercentage          *float64   `json:"discount_percentage,omitempty"`
	DiscountAmount              *float64   `json:"discount_amount,omitempty"`
	EffectiveTP                 *float64   `json:"effective_tp,omitempty"`
	EffectiveMRP                *float64   `json:"effective_mrp,omitempty"`
	MaxMRP                      *float64   `json:"max_mrp,omitempty"`
	MinMRP                      *float64   `json:"min_mrp,omitempty"`
	CustomerCategoryID          *int       `json:"customer_category_id,omitempty"`
	MembershipCardID            *int       `json:"membership_card_id,omitempty"`
	BatchProductPriceID         *int       `json:"batch_product_price_id,omitempty"`
	EntryTime                   *time.Time `json:"entry_time,omitempty"`
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

func LoadCassandraSettings() (*CassandraSettings, error) {
	// Read all the contents of the settings file.
	data, err := os.ReadFile("settings.yml")
	if err != nil {
		return nil, err
	}

	// Decode the YAML data into a CassandraSettings struct.
	var settings CassandraSettings
	err = yaml.Unmarshal(data, &settings)
	if err != nil {
		return nil, err
	}

	// Return the CassandraSettings struct.
	return &settings, nil
}

func ConnectToCassandra(settings *CassandraSettings) (*gocql.Session, error) {
	cluster := gocql.NewCluster(settings.Host)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: settings.Username,
		Password: settings.Password,
	}

	cluster.Keyspace = settings.Keyspace

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return session, nil
}

var session *gocql.Session

func main() {
	var err error
	// Initialize Cassandra settings
	settings, err := LoadCassandraSettings()
	if err != nil {
		log.Fatal(err)
	}

	// Establish a connection to Cassandra
	session, err = ConnectToCassandra(settings)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	router := mux.NewRouter()

	router.HandleFunc("/get_price_data/", GetPriceData).Methods("GET")
	router.HandleFunc("/post_price_data/", PostPriceData).Methods("POST")

	log.Fatal(http.ListenAndServe("0.0.0.0:8888", router))
}

func GetPriceData(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	// Access individual query parameters by name
	branchId := queryParams.Get("branch_id")

	query := "SELECT product_id, product_code, barcode, product_identifier, base_tp, base_mrp, branch_id, vendor_id, branch_base_tp," +
		"branch_base_mrp, promotion_id, promotion_start_date, promotion_end_date, promotion_percentage, promotion_amount, discount_percentage, discount_amount, effective_tp," +
		"effective_mrp, max_mrp, min_mrp, customer_category_id, membership_card_id, batch_product_price_id, entry_time, variant_1, variant_2, variant_3, variant_4," +
		"variant_5, variant_6, promotional_discount_3_percent, promotional_discount_3, promotional_discount_4_percent, promotional_discount_4, promotional_discount_5_percent, " +
		"promotional_discount_5 FROM product_master_prices"

	if branchId != "" {
		query += " where branch_id = " + branchId
		query += " ALLOW FILTERING"
	}
	iter := session.Query(query).Iter()

	var allProductPrices []ProductMasterPrices
	var priceData ProductMasterPrices

	scanParams := []interface{}{&priceData.ProductID, &priceData.ProductCode, &priceData.Barcode, &priceData.ProductIdentifier,
		&priceData.BaseTP, &priceData.BaseMRP, &priceData.BranchID, &priceData.VendorID,
		&priceData.BranchBaseTP, &priceData.BranchBaseMRP, &priceData.PromotionID, &priceData.PromotionStartDate,
		&priceData.PromotionEndDate, &priceData.PromotionPercentage, &priceData.PromotionAmount, &priceData.DiscountPercentage, &priceData.DiscountAmount,
		&priceData.EffectiveTP, &priceData.EffectiveMRP, &priceData.MaxMRP, &priceData.MinMRP, &priceData.CustomerCategoryID, &priceData.MembershipCardID,
		&priceData.BatchProductPriceID, &priceData.EntryTime, &priceData.Variant1, &priceData.Variant2, &priceData.Variant3, &priceData.Variant4,
		&priceData.Variant5, &priceData.Variant6, &priceData.PromotionalDiscount3Percent, &priceData.PromotionalDiscount3, &priceData.PromotionalDiscount4Percent,
		&priceData.PromotionalDiscount4, &priceData.PromotionalDiscount5Percent, &priceData.PromotionalDiscount5}

	for iter.Scan(scanParams...) {
		allProductPrices = append(allProductPrices, priceData)
	}
	fmt.Println(ProductMasterPrices{}, "hello bro are you looking for me?")

	err := json.NewEncoder(w).Encode(allProductPrices)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func PostPriceData(w http.ResponseWriter, r *http.Request) {
	var priceData ProductMasterPrices

	err := json.NewDecoder(r.Body).Decode(&priceData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO product_master_prices (product_id, product_code, barcode, product_identifier, base_tp, base_mrp, branch_id, vendor_id, " +
		"branch_base_tp, branch_base_mrp, promotion_id, promotion_start_date, promotion_end_date, promotion_percentage, promotion_amount, " +
		"discount_percentage, discount_amount, effective_tp, effective_mrp, max_mrp, min_mrp, customer_category_id, membership_card_id, batch_product_price_id, " +
		"entry_time, variant_1, variant_2, variant_3, variant_4, variant_5, variant_6, promotional_discount_3_percent, promotional_discount_3, " +
		"promotional_discount_4_percent, promotional_discount_4, promotional_discount_5_percent, promotional_discount_5) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	err = session.Query(query, gocql.TimeUUID(), priceData.ProductID, priceData.ProductCode, priceData.Barcode, priceData.ProductIdentifier,
		priceData.BaseTP, priceData.BaseMRP, priceData.BranchID, priceData.VendorID, priceData.BranchBaseTP, priceData.BranchBaseMRP,
		priceData.PromotionID, priceData.PromotionStartDate, priceData.PromotionEndDate, priceData.PromotionPercentage, priceData.PromotionAmount,
		priceData.DiscountPercentage, priceData.DiscountAmount, priceData.EffectiveTP, priceData.EffectiveMRP, priceData.MaxMRP,
		priceData.MinMRP, priceData.CustomerCategoryID, priceData.MembershipCardID, priceData.BatchProductPriceID, priceData.EntryTime,
		priceData.Variant1, priceData.Variant2, priceData.Variant3, priceData.Variant4, priceData.Variant5, priceData.Variant6,
		priceData.PromotionalDiscount3Percent, priceData.PromotionalDiscount3, priceData.PromotionalDiscount4Percent,
		priceData.PromotionalDiscount4, priceData.PromotionalDiscount5Percent, priceData.PromotionalDiscount5).Exec()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(priceData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
