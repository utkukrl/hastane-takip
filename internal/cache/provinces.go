package cache

import (
	"encoding/json"
	"hastane-takip/internal/models"
	"hastane-takip/internal/trait"
	"hastane-takip/internal/utils"
)

func GetProvincesCache() ([]models.Province, error) {
	data, err := utils.GetKey("provinces")
	if err == nil && data != "" {
		var provinces []models.Province
		err = json.Unmarshal([]byte(data), &provinces)
		if err != nil {
			return nil, err
		}
		return provinces, nil
	}
	hospitalHandler := trait.NewHospitalHandler()
	provinces, err := hospitalHandler.GetProvincesFromDB()
	if err != nil {
		return nil, err
	}
	dataBytes, err := json.Marshal(provinces)
	if err != nil {
		return nil, err
	}

	dataString := string(dataBytes)

	err = utils.SetKey("provinces", dataString)
	if err != nil {
		return nil, err
	}
	return provinces, nil
}
