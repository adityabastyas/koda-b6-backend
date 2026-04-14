package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"koda-b6-backend1/internal/lib"
	"koda-b6-backend1/internal/models"
	"koda-b6-backend1/internal/repository"
	"time"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	val, err := lib.RDB.Get(lib.Ctx, "products").Result()

	if err == nil {
		var cached []models.Product
		json.Unmarshal([]byte(val), &cached)
		fmt.Println("ambil dari redis 🔥")
		return cached, nil
	}

	// 2. kalau tidak ada → ambil DB
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	// 3. simpan ke redis
	jsonData, _ := json.Marshal(products)
	lib.RDB.Set(lib.Ctx, "products", jsonData, 5*time.Minute)

	fmt.Println("ambil dari DB + simpan redis")

	return products, nil
}

func (s *ProductService) GetByID(id int) (*models.Product, error) {
	if id <= 0 {
		return nil, errors.New("id tidak valid")
	}

	key := fmt.Sprintf("product:%d", id)

	val, err := lib.RDB.Get(lib.Ctx, key).Result()
	if err == nil {
		var product models.Product
		json.Unmarshal([]byte(val), &product)
		fmt.Println("ambil product by id dari redis 🔥")
		return &product, nil
	}

	product, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	jsonData, _ := json.Marshal(product)
	lib.RDB.Set(lib.Ctx, key, jsonData, 5*time.Minute)

	fmt.Println("ambil product by id dari DB + simpan redis")

	return product, nil
}

func (s *ProductService) Create(input models.ProductInput) error {
	if input.Name == "" {
		return errors.New("nama product tidak boleh kosong")
	}

	if input.Price <= 0 {
		return errors.New("harga product tidak valid")
	}

	lib.RDB.Del(lib.Ctx, "products")

	return s.repo.Create(input)
}

func (s *ProductService) Update(id int, input models.ProductInput) error {
	if id <= 0 {
		return errors.New("id tidak valid")
	}

	if input.Name == "" {
		return errors.New("nama product tidak boleh kosong")
	}

	if input.Price <= 0 {
		return errors.New("harga product tidak valid")
	}

	lib.RDB.Del(lib.Ctx, "products")

	return s.repo.Update(id, input)
}

func (s *ProductService) Delete(id int) error {
	if id <= 0 {
		return errors.New("id tidak valid")
	}

	key := fmt.Sprintf("product:%d", id)
	lib.RDB.Del(lib.Ctx, "products", key)

	return s.repo.Delete(id)
}
