package utils

import (
	"fmt"

	"github.com/coderkamlesh/hypershop_go/internal/constants"
)

// GetProductFolder returns folder path for product images
// Example: hypershop/products/seller_123
func GetProductFolder(sellerID string) string {
	return fmt.Sprintf("%s/%s", constants.FolderProducts, sellerID)
}

// GetUserFolder returns folder based on role and user ID
// Example: hypershop/sellers/user_456/profiles
func GetUserFolder(role, userID, assetType string) string {
	var baseFolder string

	switch role {
	case "ADMIN", "CATALOG_ADMIN":
		baseFolder = constants.FolderAdmin
	case "SELLER":
		baseFolder = constants.FolderSellers
	case "STORE_MANAGER":
		baseFolder = constants.FolderStores
	case "WAREHOUSE_MANAGER":
		baseFolder = constants.FolderWarehouses
	case "RIDER":
		baseFolder = constants.FolderRiders
	case "CONSUMER":
		baseFolder = constants.FolderConsumers
	default:
		baseFolder = constants.BaseFolder + "/general"
	}

	return fmt.Sprintf("%s/%s%s", baseFolder, userID, assetType)
}

// GetCategoryFolder returns folder for category images
func GetCategoryFolder(categoryID string) string {
	return fmt.Sprintf("%s/%s", constants.FolderCategories, categoryID)
}

// GetOrderFolder returns folder for order-related images
func GetOrderFolder(orderID string) string {
	return fmt.Sprintf("%s/%s", constants.FolderOrders, orderID)
}

// GetStoreFolder returns folder for store-specific images
func GetStoreFolder(storeID, assetType string) string {
	return fmt.Sprintf("%s/%s%s", constants.FolderStores, storeID, assetType)
}

// GetWarehouseFolder returns folder for warehouse-specific images
func GetWarehouseFolder(warehouseID, assetType string) string {
	return fmt.Sprintf("%s/%s%s", constants.FolderWarehouses, warehouseID, assetType)
}
