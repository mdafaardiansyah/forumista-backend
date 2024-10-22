package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mdafaardiansyah/forumista-backend/internal/configs"
	"log"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs/"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Gagal inisialisasi Config", err)
	}
	cfg = configs.Get()
	log.Println("config", cfg)

	r.Run(cfg.Service.Port)
}
