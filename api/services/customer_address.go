package services

import (
	"context"
	"fmt"

	"github.com/SONEsee/go-echo/api/schema/requestbody"
	dbpkg "github.com/SONEsee/go-echo/pkg/db-pkg"
	dbdelete "github.com/SONEsee/go-echo/pkg/db-pkg/db-delete"
	dbinserts "github.com/SONEsee/go-echo/pkg/db-pkg/db-inserts"
	dbquery "github.com/SONEsee/go-echo/pkg/db-pkg/db-query"
	dbschema "github.com/SONEsee/go-echo/pkg/db-pkg/db-schema"
	dbupdate "github.com/SONEsee/go-echo/pkg/db-pkg/db-update"
	"github.com/SONEsee/go-echo/pkg/pagination"
)

// CreateCustomerAddressServices ສ້າງ address; ຖ້າ customer ຍັງບໍ່ມີ default_address_id ມາກ່ອນ (ແມ່ນ address ທຳອິດ) ຈະຕັ້ງເປັນ default ໃຫ້ອັດຕະໂນມັດ
func CreateCustomerAddressServices(ctx context.Context, req requestbody.CustomerAddressRequestBody) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)

		customer, err := dbquery.GetCustomerByIDForUpdate(ctx, db, req.CustomerID)
		if err != nil {
			return err
		}

		addressID, err := dbinserts.CreateCustomerAddress(ctx, db, req)
		if err != nil {
			return err
		}

		if customer.DefaultAddressID == nil {
			if err := dbupdate.SetCustomerAddressDefault(ctx, db, addressID, true); err != nil {
				return err
			}
			if err := dbupdate.SetCustomerDefaultAddressID(ctx, db, req.CustomerID, &addressID); err != nil {
				return err
			}
		}
		return nil
	})
}

func GetDataCustomerAddressServices(ctx context.Context, id *int, customerID *int, page, pageSize int) ([]dbschema.CustomerAddressDBSchema, *pagination.PaginationResult, error) {
	var paginationParam *pagination.PaginationParams
	if page > 0 || pageSize > 0 {
		paginationParam = pagination.NewPaginationParams(page, pageSize)
	}
	items, result, err := dbquery.GetCustomerAddressDataQuery(ctx, id, customerID, paginationParam)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get data: %w", err)
	}
	return items, result, nil
}

func UpdateCustomerAddressServicesPatch(ctx context.Context, id int64, req requestbody.CustomerAddressPatchRequest) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)
		return dbupdate.UpdateCustomerAddressPatch(ctx, db, id, req)
	})
}

// DeleteCustomerAddressServices ບໍ່ອະນຸຍາດລົບ address ທີ່ເປັນ default ຢູ່ — ຕ້ອງປ່ຽນ default ໄປແຖວອື່ນກ່ອນ
func DeleteCustomerAddressServices(ctx context.Context, id int64) error {
	tx := dbpkg.GetTransactionManager()
	return tx.WithTransaction(ctx, func(ctx context.Context) error {
		db := dbpkg.GetDBFromContext(ctx)

		addr, err := dbquery.GetCustomerAddressByIDForUpdate(ctx, db, id)
		if err != nil {
			return err
		}
		if addr.IsDefault {
			return fmt.Errorf("cannot delete default address: ກະລຸນາປ່ຽນທີ່ຢູ່ default ໄປອັນອື່ນກ່ອນ")
		}
		return dbdelete.DeleteCustomerAddress(ctx, db, id)
	})
}
