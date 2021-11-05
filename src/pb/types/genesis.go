package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AdminList:      []Admin{},
		SuperAdminList: []SuperAdmin{},
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in admin
	adminIndexMap := make(map[string]struct{})

	for _, elem := range gs.AdminList {
		index := string(AdminKey(elem.Address))
		if _, ok := adminIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for admin")
		}
		adminIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in superAdmin
	superAdminIdMap := make(map[uint64]bool)
	superAdminCount := gs.GetSuperAdminCount()
	for _, elem := range gs.SuperAdminList {
		if _, ok := superAdminIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for superAdmin")
		}
		if elem.Id >= superAdminCount {
			return fmt.Errorf("superAdmin id should be lower or equal than the last id")
		}
		superAdminIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
