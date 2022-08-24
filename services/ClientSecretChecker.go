package services

import "arvan-credit-service/configs/auth"

type ClientSecretChecker struct{}

func (c ClientSecretChecker) GetSecret(token string) (uint, bool) {
	clientId, ok := auth.Secrets()["client_secret_tokens"].(map[string]uint)[token]
	return clientId, ok
}
