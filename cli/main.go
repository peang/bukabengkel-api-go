package main

import (
	"fmt"

	"github.com/peang/bukabengkel-api-go/src/cmd"
	"github.com/peang/bukabengkel-api-go/src/config"
	repository "github.com/peang/bukabengkel-api-go/src/repositories"
	file_service "github.com/peang/bukabengkel-api-go/src/services/file_services"
	"github.com/peang/bukabengkel-api-go/src/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "my-cli-app",
	Short: "A simple CLI application built with Go and Echo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from the root command!")
	},
}

func main() {
	// Load config
	configApp := config.LoadConfig()

	// Setup logger
	appLogger := utils.NewApiLogger(configApp)
	appLogger.InitLogger()

	// Setup Databse
	db := config.LoadDatabase(configApp)
	defer db.Close()

	// services
	fileService := file_service.NewAWSS3Service(configApp)

	// Repositories
	imageRepo := repository.NewImageRepository(db, fileService)
	// productRepo := repository.NewProductRepository(db, imageRepo)
	productDistributorRepo := repository.NewProductDistributorRepository(db, imageRepo)
	productCategoryDistributorRepo := repository.NewProductCategoryDistributorRepository(db)

	// Usecases
	// productUsecase := usecase.NewProductUsecase(productRepo)

	Register(productDistributorRepo, productCategoryDistributorRepo)
}

func Register(
	productDistributorRepo *repository.ProductDistributorRepository,
	productCategoryDistributorRepo *repository.ProductCategoryDistributorRepository,
) {
	rootCmd := &cobra.Command{
		Use:   "",
		Short: "Ping Command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Pong !")
		},
	}

	asian := cmd.NewSyncAsian(productDistributorRepo, productCategoryDistributorRepo)
	syncAsianCmd := &cobra.Command{
		Use:   "sync-asian",
		Short: "Sync Asian Products",
		Run:   asian.Execute,
	}

	rootCmd.AddCommand(syncAsianCmd)

	rootCmd.Execute()
}
