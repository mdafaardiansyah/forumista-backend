package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mdafaardiansyah/forumista-backend/internal/configs"
	"github.com/mdafaardiansyah/forumista-backend/internal/handlers/memberships"
	"github.com/mdafaardiansyah/forumista-backend/internal/handlers/posts"
	membershipRepo "github.com/mdafaardiansyah/forumista-backend/internal/repository/memberships"
	postRepo "github.com/mdafaardiansyah/forumista-backend/internal/repository/posts"
	membershipSvc "github.com/mdafaardiansyah/forumista-backend/internal/service/memberships"
	postSvc "github.com/mdafaardiansyah/forumista-backend/internal/service/posts"
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

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	postService := postSvc.NewService(cfg, postRepo)

	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	postHandler := posts.NewHandler(r, postService)
	postHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
