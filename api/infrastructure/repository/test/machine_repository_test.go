package repository

import (
	domainLocation "api/domain/model/location"
	domainMachine "api/domain/model/machine"
	"api/infrastructure/database/sql"
	"api/infrastructure/repository"
	"fmt"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func Test_NewMachineRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("FAILED TO CREATE SQL MOCK: %s", err)
	}
	dbx := sqlx.NewDb(db, "sqlmock")
	defer db.Close()
	defer dbx.Close()

	repo := repository.NewMachineRepository(dbx)
	if repo == nil {
		t.Errorf("FAILED TO CREATE \"machine repository\".")
	}
}

func Test_FindMachineByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("FAILED TO CREATE SQL MOCK: %s", err)
	}
	dbx := sqlx.NewDb(db, "sqlmock")
	defer db.Close()
	defer dbx.Close()

	repo := repository.NewMachineRepository(dbx)

	tests := []struct {
		name  string
		arg   domainMachine.MachineID
		want  *domainMachine.Machine
		isErr bool
		err   error
	}{
		{
			name: "Successfully",
			arg:  "0001",
			want: &domainMachine.Machine{
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
			isErr: false,
			err:   nil,
		},
		{
			name:  "Error",
			arg:   "XXXX",
			want:  nil,
			isErr: true,
			err:   fmt.Errorf("FAILED TO FIND A MAHCINE BY ID."),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isErr {
				mock.ExpectQuery(regexp.QuoteMeta(sql.FindAllMachines)).WillReturnError(fmt.Errorf("FAILED TO FIND A MACHINE BY ID."))
			} else {
				rows := sqlmock.NewRows([]string{
					"mst_machine_id",
					"machine_name",
				}).
					AddRow("0001", "test machine 1")
				mock.ExpectQuery(regexp.QuoteMeta(sql.FindMachineByID)).WithArgs(tt.arg).WillReturnRows(rows)
			}

			got, err := repo.FindMachineByID(domainMachine.MachineID(tt.arg))
			if (err != nil) != tt.isErr {
				t.Errorf("FAILED TO EXECUTE sqlmock.FindAllMachines: RETURN ERROR: %s", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MachineRepository.FindMachineByID RETURNS %v, but want %v\n", got, tt.want)
			}
		})
	}
}

func Test_FindAllMachines(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("FAILED TO CREATE SQL MOCK: %s", err)
	}
	dbx := sqlx.NewDb(db, "sqlmock")
	defer db.Close()
	defer dbx.Close()

	repo := repository.NewMachineRepository(dbx)

	tests := []struct {
		name  string
		want  []*domainMachine.Machine
		isErr bool
		err   error
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
			isErr: false,
			err:   nil,
		},
		{
			name:  "Error",
			want:  nil,
			isErr: true,
			err:   fmt.Errorf("FAILED TO FIND ALL MACHINES"),
		},
	}

	for _, tt := range tests {
		t.Run("Successfully", func(t *testing.T) {
			if tt.isErr {
				mock.ExpectQuery(sql.FindAllMachines).WillReturnError(fmt.Errorf("FAILED TO FIND ALL MACHINES"))
			} else {
				rows := sqlmock.NewRows([]string{
					"mst_machine_id",
					"machine_name",
				}).
					AddRow("0001", "test machine 1")
				mock.ExpectQuery(sql.FindAllMachines).WillReturnRows(rows)
			}

			got, err := repo.FindAllMachines()
			if (err != nil) != tt.isErr {
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

func Test_CreateMachine(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("FAILED TO CREATE SQL MOCK: %s", err)
	}
	dbx := sqlx.NewDb(db, "sqlmock")
	defer db.Close()
	defer dbx.Close()

	repo := repository.NewMachineRepository(dbx)

	tests := []struct {
		name        string
		arg         *domainMachine.Machine
		isErr       bool
		err         error
		rowAffected int64
	}{
		{
			name: "Successfully",
			arg: &domainMachine.Machine{
				ID:                 "0001",
				Name:               "test machine1 1",
				FactoryID:          "0001",
				MakerID:            "0001",
				Remark:             "",
				TableInformationID: "XXXXXXXXXX",
			},
			isErr:       false,
			err:         nil,
			rowAffected: 1,
		},
		{
			name:        "Error",
			arg:         &domainMachine.Machine{},
			isErr:       true,
			err:         fmt.Errorf("FAILED TO CREATE MACHINE."),
			rowAffected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isErr {
				mock.ExpectExec("INSERT INTO emn.mst_machine").
					WithArgs(tt.arg.ID, tt.arg.Name, tt.arg.FactoryID, tt.arg.MakerID, tt.arg.Remark, tt.arg.TableInformationID).
					WillReturnError(fmt.Errorf("FAILED TO CREATE MACHINE."))
			} else {
				mock.ExpectExec("INSERT INTO emn.mst_machine").
					WithArgs(tt.arg.ID, tt.arg.Name, tt.arg.FactoryID, tt.arg.MakerID, tt.arg.Remark, tt.arg.TableInformationID).
					WillReturnResult(sqlmock.NewResult(1, tt.rowAffected))
			}

			err := repo.CreateMachine(tt.arg)
			if (err != nil) != tt.isErr {
				t.Errorf("FAILED TO TEST; MachineRepository.CreateMachine RETURN ERROR: %s", err.Error())
			}
		})
	}
}

func Test_UpdateMachine(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("FAILED TO CREATE SQL MOCK: %s", err)
	}
	dbx := sqlx.NewDb(db, "sqlmock")
	defer db.Close()
	defer dbx.Close()

	repo := repository.NewMachineRepository(dbx)

	tests := []struct {
		name        string
		arg         *domainMachine.Machine
		isErr       bool
		want        error
		rowAffected int64
	}{
		{
			name: "Successfully",
			arg: &domainMachine.Machine{
				ID:                 "0001",
				Name:               "test machine 1",
				FactoryID:          "0001",
				MakerID:            "0001",
				Remark:             "",
				StopUsing:          time.Now(),
				TableInformationID: "XXXXXXXXXX",
			},
			isErr:       false,
			want:        nil,
			rowAffected: 1,
		},
		{
			name:        "Error",
			arg:         &domainMachine.Machine{},
			isErr:       true,
			want:        fmt.Errorf("FAILED TO UPDATE THE MACHINE"),
			rowAffected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isErr {
				mock.ExpectExec("UPDATE emn.mst_machine").
					WithArgs(tt.arg.Name, tt.arg.FactoryID, tt.arg.MakerID, tt.arg.Remark, tt.arg.StopUsing, tt.arg.TableInformationID, tt.arg.ID).
					WillReturnError(fmt.Errorf("FAILED TO UPDATE THE MACHINE."))
			} else {
				mock.ExpectExec("UPDATE emn.mst_machine").
					WithArgs(tt.arg.Name, tt.arg.FactoryID, tt.arg.MakerID, tt.arg.Remark, tt.arg.StopUsing, tt.arg.TableInformationID, tt.arg.ID).
					WillReturnResult(sqlmock.NewResult(1, tt.rowAffected))
			}

			err := repo.UpdateMachine(tt.arg)
			if (err != nil) != tt.isErr {
				t.Errorf("FAILED TO TEST; MachineRepository.UpdateMachine RETURNS ERROR: %s", err)
			}
		})
	}
}

func Test_DeleteMachine(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("FAILED TO CREATE SQL MOCK: %s", err)
	}
	dbx := sqlx.NewDb(db, "sqlmock")
	defer db.Close()
	defer dbx.Close()

	repo := repository.NewMachineRepository(dbx)

	tests := []struct {
		name        string
		arg         *domainMachine.Machine
		isErr       bool
		want        error
		rowAffected int64
	}{
		{
			name: "Successfully",
			arg: &domainMachine.Machine{
				ID: "0001",
			},
			isErr:       false,
			want:        nil,
			rowAffected: 1,
		},
		{
			name:        "Error",
			arg:         &domainMachine.Machine{},
			isErr:       true,
			want:        fmt.Errorf("FAILED TO DELETE THE MACHINE"),
			rowAffected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isErr {
				mock.ExpectExec("UPDATE emn.mst_machine").
					WithArgs(tt.arg.ID).
					WillReturnError(fmt.Errorf("FAILED TO DELETE THE MACHINE."))
			} else {
				mock.ExpectExec("UPDATE emn.mst_machine").
					WithArgs(tt.arg.ID).
					WillReturnResult(sqlmock.NewResult(1, tt.rowAffected))
			}

			err := repo.StopUsingMachine(tt.arg)
			if (err != nil) != tt.isErr {
				t.Errorf("FAILED TO TEST; MachineRepository.StopUsingMachine RETURNS ERROR: %s", err)
			}
		})
	}
}
