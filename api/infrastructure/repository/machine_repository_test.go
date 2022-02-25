package repository

import (
	domainLocation "api/domain/model/location"
	domainMachine "api/domain/model/machine"
	"api/infrastructure/database/sql"
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func Test_NewMachineRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("FAILED TO CREATE SQL MOCK: %s", err)
	}
	dbx := sqlx.NewDb(db, "sqlmock")
	defer dbx.Close()

	repo := NewMachineRepository(dbx)
	if repo == nil {
		t.Errorf("FAILED TO CREATE \"machine repository\".")
	}
}

func Test_FindAllMachines(t *testing.T) {
	type wantErr struct {
		isErr bool
		err   error
	}
	tests := []struct {
		name    string
		want    []*domainMachine.Machine
		wantErr wantErr
	}{
		{
			name: "Successfully",
			want: []*domainMachine.Machine{
				{
					ID:   "0001",
					Name: "test machine 1",
					Factory: domainLocation.Factory{
						Company: domainLocation.Company{
							Address: domainLocation.Address{
								PhoneNumberList: []*domainLocation.PhoneNumberList{},
							},
						},
						Address: domainLocation.Address{
							PhoneNumberList: []*domainLocation.PhoneNumberList{},
						},
					},
					Maker: domainLocation.Company{
						Address: domainLocation.Address{
							PhoneNumberList: []*domainLocation.PhoneNumberList{},
						},
					},
				},
			},
			wantErr: wantErr{
				isErr: false,
				err:   nil,
			},
		},
		{
			name: "Error",
			want: nil,
			wantErr: wantErr{
				isErr: true,
				err:   fmt.Errorf("FAILED TO FIND ALL MACHINES"),
			},
		},
	}
	for _, tt := range tests {
		t.Run("Successfully", func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("FAILED TO CREATE SQL MOCK: %s", err)
			}
			dbx := sqlx.NewDb(db, "sqlmock")
			defer dbx.Close()

			if tt.wantErr.isErr {
				mock.ExpectQuery(sql.FindAllMachines).WillReturnError(fmt.Errorf("FAILED TO FIND ALL MACHINES"))
			}
			rows := sqlmock.NewRows([]string{
				"mst_machine_id",
				"machine_name",
			}).
				AddRow("0001", "test machine 1")
			mock.ExpectQuery(sql.FindAllMachines).WillReturnRows(rows)

			repo := NewMachineRepository(dbx)
			got, err := repo.FindAllMachines()
			if (err != nil) != tt.wantErr.isErr {
				t.Errorf("FAILED TO EXECUTE sqlmock.FindAllMachines: RETURN ERROR: %s", err)
			}
			for i := range got {
				if !reflect.DeepEqual(got[i], tt.want[i]) {
					t.Errorf("MachineRepository.FindAllMachines RETURNS %v, BUT WANT %v\n", got[i], tt.want[i])
				}
			}
		})

	}
}
