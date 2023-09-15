package models

import (
	"database/sql"
	"time"

	"github.com/marcosrgrave/mvc-crud-products/db"
	"github.com/marcosrgrave/mvc-crud-products/utils"
)

func ProductTable() *db.Table {
	return &db.Table{
		Name:    "tb_products",
		Columns: []string{"id", "name", "description", "price", "quantity", "created_at", "creator"},
	}
}

func ProductQueries() *db.Queries {
	return ProductTable().Queries()
}

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	CreatorID   string    `json:"creator_id"`
}

type ProductDTOInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	CreatorID   string  `json:"creator_id"`
}

type ProductDTOOutput struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
}

func ConvertProductDTOOutput(product Product) ProductDTOOutput {
	return ProductDTOOutput{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		CreatedAt:   product.CreatedAt,
	}
}

func ConvertManyProductDTOOutput(products []*Product) []ProductDTOOutput {
	var result []ProductDTOOutput
	for _, product := range products {
		result = append(result, ConvertProductDTOOutput(*product))
	}
	return result
}

func ConvertProductDTOInput(product *Product) *ProductDTOInput {
	return &ProductDTOInput{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		CreatorID:   product.CreatorID,
	}
}

func NewProduct(name string, description string, price float64, quantity int, creatorID string) *Product {
	return &Product{
		ID:          utils.UniqueID(),
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
		CreatedAt:   time.Now(),
		CreatorID:   creatorID,
	}
}

func ReadAllProducts(db *sql.DB) ([]*Product, error) {
	rows, err := db.Query(ProductQueries().SelectAll)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*Product

	for rows.Next() {
		var product Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Quantity,
			&product.CreatedAt,
			&product.CreatorID,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}

func ReadProduct(db *sql.DB, id string) (*Product, error) {
	// to avoid crashes from column changes, im using specific column names instead of *.
	stmt, err := db.Prepare(ProductQueries().SelectByID)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var product Product

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.CreatedAt, &product.CreatorID)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func CreateProduct(db *sql.DB, product *Product) error {
	// the statement below protects the database against SQL Injection, which is pretty cool.
	stmt, err := db.Prepare(ProductQueries().InsertByStruct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Description, product.Price, product.Quantity, product.CreatedAt, product.CreatorID)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare(ProductQueries().UpdateByStruct)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Description, product.Price, product.Quantity, product.CreatedAt, product.CreatorID, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare(ProductQueries().DeleteByID)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

// TODO:
// CRUD structs Validator: specific "columns" shouldn't have the standard init value. String shouldn't be "", Int shouldn't be 0, etc.

func (p *Product) Validate() bool {
	if p.ID == "" || p.Name == "" || p.Price == 0 || time.Time.IsZero(p.CreatedAt) || p.CreatorID == "" {
		return false
	}

	return false
}

// func GetStructValues(p Product) {
// 	val := reflect.ValueOf(p)
// 	typ := val.Type()

// 	for i := 0; i < val.NumField(); i++ {
// 		field := val.Field(i)
// 		fieldType := typ.Field(i)

// 		fmt.Printf("%s: %v\n", fieldType.Name, field.Interface())
// 	}
// }
