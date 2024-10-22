package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mdafaardiansyah/forumista-backend/internal/configs"
	"github.com/mdafaardiansyah/forumista-backend/internal/handlers/memberships"
	membershipRepo "github.com/mdafaardiansyah/forumista-backend/internal/repository/memberships"
	"github.com/mdafaardiansyah/forumista-backend/pkg/internalsql"
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

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisialisasi ke Database", err)
	}

	_ = membershipRepo.NewRepository(db)

	membershipHandler := memberships.NewHandler(r)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
