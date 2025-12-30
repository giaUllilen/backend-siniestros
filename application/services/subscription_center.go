package services

import (
	"is-public-api/application/models"
	"is-public-api/application/storage"
)

type subscriptionCenterService struct {
	subscriptionCenterRepository storage.ISubscriptionCenterRepository
}

func NewSubscriptionCenterService(repo storage.ISubscriptionCenterRepository) *subscriptionCenterService {
	return &subscriptionCenterService{subscriptionCenterRepository: repo}
}

func (service *subscriptionCenterService) CreateSubscriptionCenterOptions(txContext *models.TxContext, key, option, value, version string, description *string) (string, error) {
	id, err := service.subscriptionCenterRepository.CreateSubscriptionOptions(txContext, key, option, value, version, description)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (service *subscriptionCenterService) FindSubscriptionCenterOptionsByID(txContext *models.TxContext, key string) (string, error) {
	id, err := service.subscriptionCenterRepository.FindSubscriptionOptionsByID(txContext, key)
	if err != nil {
		return "", err
	}

	return id, nil
}
