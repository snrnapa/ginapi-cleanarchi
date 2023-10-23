package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// envLoad env load
func EnvLoad() {
	const targetEnvName = "GO_ENV"
	// 環境未設定の場合はローカルを設定
	if "" == os.Getenv(targetEnvName) {
		_ = os.Setenv(targetEnvName, "local")
	}
	// err := godotenv.Load(fmt.Sprintf("./common/env/%s.env", os.Getenv(targetEnvName)))
	err := godotenv.Load(fmt.Sprintf("./common/env/%s.env", os.Getenv("GO_ENV")))

	if err != nil {
		log.Fatalf("Error loading env target env is %s", os.Getenv("GO_ENV"))
	}
}
