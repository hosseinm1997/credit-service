package services

type IClientSecretChecker interface {
	GetSecret(token string) (uint, bool)
}
