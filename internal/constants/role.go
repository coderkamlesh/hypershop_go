package constants

type Role string

const (
	RoleAdmin            Role = "ADMIN"
	RoleCatalogAdmin     Role = "CATALOG_ADMIN"
	RoleSeller           Role = "SELLER"
	RoleWarehouseManager Role = "WAREHOUSE_MANAGER"
	RoleStoreManager     Role = "STORE_MANAGER"
	RoleRider            Role = "RIDER"
	RoleConsumer         Role = "CONSUMER"
)

func (r Role) String() string {
	return string(r)
}

func (r Role) IsValid() bool {
	switch r {
	case RoleAdmin, RoleCatalogAdmin, RoleSeller, RoleWarehouseManager,
		RoleStoreManager, RoleRider, RoleConsumer:
		return true
	}
	return false
}
