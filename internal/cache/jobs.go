package cache

import (
	"encoding/json"
	"hastane-takip/internal/models"
	trait "hastane-takip/internal/trait"
	"hastane-takip/internal/utils"
)

func GetJobCategoriesCache() ([]models.JobCategory, error) {
	data, err := utils.GetKey("job_categories")
	if err == nil && data != "" {
		var categories []models.JobCategory
		err = json.Unmarshal([]byte(data), &categories)
		if err != nil {
			return nil, err
		}
		return categories, nil
	}

	staffHandler := trait.NewStaffHandler()
	categories, err := staffHandler.GetJobCategoriesFromDB()
	if err != nil {
		return nil, err
	}
	dataBytes, err := json.Marshal(categories)
	if err != nil {
		return nil, err
	}
	dataString := string(dataBytes)
	err = utils.SetKey("job_categories", dataString)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
