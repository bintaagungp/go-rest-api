package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Category struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Category  string    `gorm:"size:255;not null;unique" json:"category"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *Category) Prepare() {
	c.ID = 0
	c.Category = html.EscapeString(strings.TrimSpace(c.Category))
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

func (c *Category) Validate() error {
	if c.Category == "" {
		return errors.New("Required Category")
	}
	return nil
}

func (c *Category) Save(db *gorm.DB) (*Category, error) {
	err := db. /* Debug(). */ Model(&Category{}).Create(&c).Error
	if err != nil {
		return &Category{}, err
	}
	return c, nil
}

func (c *Category) Find(db *gorm.DB, id uint64) (*Category, error) {
	err := db. /* Debug(). */ Model(&Category{}).Where("id = ?", id).Take(&c).Error
	if err != nil {
		return &Category{}, err
	}
	return c, nil
}

func (c *Category) FindAll(db *gorm.DB) (*[]Category, error) {
	categories := []Category{}
	err := db. /* Debug(). */ Model(&Category{}).Limit(100).Find(&categories).Error
	if err != nil {
		return &[]Category{}, err
	}
	return &categories, nil
}

func (c *Category) Delete(db *gorm.DB, id uint32) (uint64, error) {
	db = db. /* Debug(). */ Where("id = ?", id).Delete(&Category{})
	if db.Error != nil {
		return 0, errors.New("Category not found")
	}
	return uint64(db.RowsAffected), nil
}

func (c *Category) Update(db *gorm.DB, id uint64) (uint64, error) {
	db = db. /* Debug(). */ Model(&Category{}).Where("id = ?", id).Updates(map[string]interface{}{"category": c.Category, "updated_at": time.Now()})
	err := db.Error
	if err != nil {
		return 0, err
	}
	return uint64(db.RowsAffected), nil
}
