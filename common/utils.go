package common

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type RespCode int
type Role int

const (
	SUCCEED RespCode = 0
	FAILED           = -1
)
const APP_VERSION = "0.2.1"
const EN_KEY1 = "5c072181d566c72ff169"
const EN_KEY2 = "a624d7bbb012c6490df2653f5f0f4037c4d407d7291e"

func GetRespResult(code int, message string, dataList interface{}, totalCount int) interface{} {
	return GetRespResultWithPage(code, message, dataList, 1, 1, totalCount)
}

func GetRespResultWithPage(code int, message string, dataList interface{}, pageIndex int, totalPages int, totalCount int) interface{} {
	return gin.H{
		"code":       code,
		"dataList":   dataList,
		"pageIndex":  pageIndex,
		"totalPages": totalPages,
		"totalCount": totalCount,
		"message":    message,
	}
}

func GetEnvValue(envKey string, defaultValue string) (envValue string) {
	envValue = os.Getenv(envKey)
	if envValue == "" {
		envValue = defaultValue
	}
	return
}

const (
	PLATFORM_ADMIN Role = iota + 1 // value --> 1
	NORMAL_ADMIN
	NORMAL_MEMBER
	GROUP_ADMIN
	GROUP_MASTER
	GROUP_DEVELOP
)

func DecryptString(encryptedString string) string {
	if len(encryptedString) == 0 {
		return encryptedString
	}

	key, _ := hex.DecodeString(EN_KEY1 + EN_KEY2)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error(err)
		return encryptedString
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Error(err)
		return encryptedString
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Error(err)
		return encryptedString
	}

	return fmt.Sprintf("%s", plaintext)
}
