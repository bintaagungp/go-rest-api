package modeltests

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/victorsteven/fullstack/api/models"
	"gopkg.in/go-playground/assert.v1"
)

func TestGetAllCategory(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	category, err := category.FindAll(server.DB)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(*category), 3)
}

func TestGetCategoryById(t *testing.T) {

	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	// Test case
	category := models.Category{
		ID:       1,
		Category: "Hello",
	}

	result, err := category.Find(server.DB, category.ID)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, result.ID, category.ID)
	assert.Equal(t, result.Category, category.Category)
}

func TestCreateCategory(t *testing.T) {
	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	category := models.Category{
		Category: "Yup",
	}

	result, err := category.Save(server.DB)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, result.Category, category.Category)
}

func TestUpdateCategory(t *testing.T) {

	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	category := models.Category{
		Category: "Dunia",
	}

	result, err := category.Update(server.DB, 2)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, result, uint64(1))
}

func TestDeleteCategory(t *testing.T) {

	var err error

	err = resetCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	err = seedCategoryTable()
	if err != nil {
		t.Fatal(err)
	}

	result, err := category.Delete(server.DB, 1)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, result, uint64(1))
}
