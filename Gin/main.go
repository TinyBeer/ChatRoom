package main

import (
	"ChatRoom/Gin/dao"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Config() {
	pathName, _ := os.Executable()
	workDir := filepath.Dir(pathName)
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {

	Config()

	dao.InitDao()

	r := gin.Default()

	r = Router(r)

	r.Run(":" + viper.GetString("server.port"))
}

func RandomString(n int) string {
	var letters = []byte("asdfghjklqwertyuiopzxcvbnm")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
