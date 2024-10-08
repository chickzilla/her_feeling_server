package utils

import (
	"errors"
	"os"
	"sync"
)

func GetEnv(wg *sync.WaitGroup, envChan chan<- string, errChan chan<- error) {
	defer wg.Done()

	/*err := godotenv.Load()
	if err != nil {
		errChan <- errors.New("error while loading .env file")
		return
	}*/

	textAIURL := os.Getenv("AI_TEXT_URL")
	if textAIURL == "" {
		errChan <- errors.New("cannot read AI_TEXT_URL in .env file")
		return
	}
	envChan <- textAIURL
}

func GetEnvNoCon(envName string) (envValue string, errs error) {

	/*godotenv.Load()
	if err != nil {
		return "", errors.New("error while loading .env file")
	}
	*/

	textURL := os.Getenv(envName)
	if textURL == "" {
		return "", errors.New("cannot read " + envName + " in .env file")
	}

	return textURL, nil

}
