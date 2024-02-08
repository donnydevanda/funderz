package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserId(userID int) ([]Campaign, error)
}

// give direct access to DB
type repository struct {
	db *gorm.DB
}

// function to be called on main.go, to make it accesible for routing
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUserID (userID int) ([]Campaign, error) {
	var campaigns []Campaign
	// Preload: to load data from associated tables
	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}