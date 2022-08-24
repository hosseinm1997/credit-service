package auth

import "arvan-credit-service/types/structs"

func Secrets() map[string]any {
	return structs.Config{
		"client_secret_tokens": map[string]uint{
			// wallet service
			"XKP1Y7ktj27pnOJCxHbZZ32IqD5QjkWio5hfOZYx": 1,
		},
	}
}
