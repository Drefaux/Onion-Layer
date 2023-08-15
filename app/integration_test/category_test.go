package integration__test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestCategoryScenario(t *testing.T) {
	ctx := context.Background()
	// create
	createCategoryRes, err := createCategory(ctx)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(createCategoryRes)

	// get
	getAllCategoryRes, err := getAllCategories(ctx)
	if err != nil {
		t.Fatal(err)
	}

	for _, item := range getAllCategoryRes {
		t.Logf("catgory: %+v\n", item)
	}
}

type CreateCategoryResponse struct {
	ID        int    `json:"id"`
	CompanyID int    `json:"companyID"`
	Name      string `json:"name"`
}

func createCategory(ctx context.Context) (*CreateCategoryResponse, error) {
	baseURL := "http://localhost" + testServerPort + "/api-onion-layer/v2/categories"
	msg := map[string]interface{}{
		"id":        1,
		"companyID": 1,
		"name":      "categoryX",
	}
	b, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	buf := new(bytes.Buffer)
	buf.Write(b)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL, buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	var parsedRes CreateCategoryResponse
	if err := json.Unmarshal(body, &parsedRes); err != nil {
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	return &parsedRes, nil
}

type GetAllCategoriesResponse struct {
	ID        int    `json:"id"`
	CompanyID int    `json:"companyID"`
	Name      string `json:"name"`
}

func getAllCategories(ctx context.Context) ([]*GetAllCategoriesResponse, error) {
	baseURL := "http://localhost" + testServerPort + "/api-onion-layer/v2/categories/category"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get all categories by company id: %w", err)
	}

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	var parsedRes []GetAllCategoriesResponse
	if err := json.Unmarshal(body, &parsedRes); err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	result := make([]*GetAllCategoriesResponse, len(parsedRes))
	for i := range parsedRes {
		result[i] = &parsedRes[i]
	}

	return result, nil

}
