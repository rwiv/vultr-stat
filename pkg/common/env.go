package common

import (
	"fmt"
	"os"
)

type Env struct {
	VultrApiKey string
}

func ReadEnv() (*Env, error) {
	vultrApiKey := os.Getenv("VULTR_API_KEY")
	if vultrApiKey == "" {
		return nil, fmt.Errorf("VULTR_API_KEY is not set")
	}
	return &Env{
		VultrApiKey: vultrApiKey,
	}, nil
}
