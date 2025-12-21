package constants

// Base folder for the entire application
const BaseFolder = "hypershop"

// Role-based folders
const (
	FolderAdmin      = BaseFolder + "/admin"
	FolderSellers    = BaseFolder + "/sellers"
	FolderStores     = BaseFolder + "/stores"
	FolderWarehouses = BaseFolder + "/warehouses"
	FolderRiders     = BaseFolder + "/riders"
	FolderConsumers  = BaseFolder + "/consumers"
)

// Entity-based folders
const (
	FolderProducts   = BaseFolder + "/products"
	FolderCategories = BaseFolder + "/categories"
	FolderBrands     = BaseFolder + "/brands"
	FolderOrders     = BaseFolder + "/orders"
	FolderBanners    = BaseFolder + "/banners"
)

// Asset type folders
const (
	FolderProfilePics = "/profiles"
	FolderDocuments   = "/documents"
	FolderGallery     = "/gallery"
	FolderThumbnails  = "/thumbnails"
)
