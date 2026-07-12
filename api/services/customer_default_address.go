package services

import (
	"context"
	"fmt"

	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbquery "github.com/SONEsee/go-echo/pkg/db-pkg/db-query"
	dbupdate "github.com/SONEsee/go-echo/pkg/db-pkg/db-update"
)

// SetCustomerDefaultAddressServices ຕັ້ງ default address ໃໝ່ໃຫ້ customer — validate ວ່າ address ນັ້ນເປັນຂອງ customer ນີ້ແທ້,
// sync is_default ຂອງທຸກ address (ເລືອກ=true, ອື່ນ=false) ພ້ອມກັບ customers.default_address_id ໃນ transaction ດຽວ
func SetCustomerDefaultAddressServices(ctx context.Context, customerID int64, addressID int64) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)

		customer, err := dbquery.GetCustomerByIDForUpdate(ctx, db, int(customerID))
		if err != nil {
			return err
		}

		addr, err := dbquery.GetCustomerAddressByIDForUpdate(ctx, db, addressID)
		if err != nil {
			return err
		}
		if int64(addr.CustomerID) != customerID {
			return fmt.Errorf("address %d does not belong to customer %d", addressID, customerID)
		}

		if err := dbupdate.ClearCustomerAddressDefaults(ctx, db, customer.ID, addressID); err != nil {
			return err
		}
		if err := dbupdate.SetCustomerAddressDefault(ctx, db, addressID, true); err != nil {
			return err
		}
		return dbupdate.SetCustomerDefaultAddressID(ctx, db, customer.ID, &addressID)
	})
}
