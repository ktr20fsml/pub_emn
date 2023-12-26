package usecase_test

import (
	domainGeneral "api/domain/model/general"
	domainMachine "api/domain/model/machine"
	mock_repository "api/domain/repository/mock"
	"api/usecase"
	"context"
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func Test_NewMachineUsecase(t *testing.T) {
	transactionRepository := &mock_repository.MockTransactionRepository{}
	machineRepository := &mock_repository.MockMachineRepository{}
	generalRepository := &mock_repository.MockGeneralRepository{}
	machineUsecase := usecase.NewMachineUsecase(transactionRepository, machineRepository, generalRepository)
	if machineUsecase == nil {
		t.Errorf("MachineUsecase.NewMachineUsecase: should return \"machine\" usecase, but got: nil")
	}
}

func Test_FindMachineByID(t *testing.T) {
	tests := []struct {
		name       string
		repository *mock_repository.MockMachineRepository
		arg        domainMachine.MachineID
		want       *domainMachine.Machine
		isErr      bool
		err        error
	}{
		{
			name: "Successfully",
			repository: &mock_repository.MockMachineRepository{
				MockFindMachineByID: func(id domainMachine.MachineID) (*domainMachine.Machine, error) {
					machine := &domainMachine.Machine{
						ID:   "0001",
						Name: "test machine 1",
					}
					return machine, nil
				},
			},
			arg: domainMachine.MachineID("0001"),
			want: &domainMachine.Machine{
				ID:   "0001",
				Name: "test machine 1",
			},
			isErr: false,
			err:   nil,
		},
		{
			name: "Error",
			repository: &mock_repository.MockMachineRepository{
				MockFindMachineByID: func(id domainMachine.MachineID) (*domainMachine.Machine, error) {
					return nil, fmt.Errorf("FAILED TO FIND THE MACHINE DATA.")
				},
			},
			arg:   domainMachine.MachineID("XXXX"),
			want:  nil,
			isErr: true,
			err:   fmt.Errorf("FAILED TO FIND THE MACHINE DATA."),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			machineRepository := &mock_repository.MockMachineRepository{}
			machineRepository.MockFindMachineByID = tt.repository.MockFindMachineByID
			transactionRepository := &mock_repository.MockTransactionRepository{}
			generalRepository := &mock_repository.MockGeneralRepository{}
			machineUsecase := usecase.NewMachineUsecase(transactionRepository, machineRepository, generalRepository)

			got, err := machineUsecase.FindMachineByID(tt.arg)
			if tt.isErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("MachineRepository.FindMachineByID: error = %v, got err = %v", err, tt.err)
				}
			}
			if (err != nil) != tt.isErr {
				t.Errorf("MachineRepository.FindMachineByID: error = %v, isErr = %t", err, tt.isErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MachineRepository.FindMachineByID: got = %v, tt.want = %v", got, tt.want)
			}
		})
	}
}

func Test_FindAllMachines(t *testing.T) {
	tests := []struct {
		name       string
		repository *mock_repository.MockMachineRepository
		want       []*domainMachine.Machine
		isErr      bool
		err        error
	}{
		{
			name: "Successfully",
			repository: &mock_repository.MockMachineRepository{
				MockFindAllMachines: func() ([]*domainMachine.Machine, error) {
					machines := make([]*domainMachine.Machine, 3)
					for i := range machines {
						machines[i] = &domainMachine.Machine{
							ID:   domainMachine.MachineID(fmt.Sprintf("%04d", i+1)),
							Name: "test machine " + strconv.Itoa(i+1),
						}
					}

					return machines, nil
				},
			},
			want: []*domainMachine.Machine{
				{
					ID:   "0001",
					Name: "test machine 1",
				},
				{
					ID:   "0002",
					Name: "test machine 2",
				},
				{
					ID:   "0003",
					Name: "test machine 3",
				},
			},
			isErr: false,
			err:   nil,
		},
		{
			name: "Error",
			repository: &mock_repository.MockMachineRepository{
				MockFindAllMachines: func() ([]*domainMachine.Machine, error) {
					return nil, fmt.Errorf("FAILED TO FIND THE MACHINES DATA.")
				},
			},
			want:  nil,
			isErr: true,
			err:   fmt.Errorf("FAILED TO FIND THE MACHINES DATA."),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			machineRepository := &mock_repository.MockMachineRepository{}
			machineRepository.MockFindAllMachines = tt.repository.MockFindAllMachines
			transactionRepository := &mock_repository.MockTransactionRepository{}
			generalRepository := &mock_repository.MockGeneralRepository{}
			machineUsecase := usecase.NewMachineUsecase(transactionRepository, machineRepository, generalRepository)

			got, err := machineUsecase.FindAllMachines()
			if tt.isErr {
				if err.Error() != tt.err.Error() {
					t.Errorf("MachineRepository.MockFindAllMachines: error = %v, got err = %v", err, tt.err)
				}
			}
			if (err != nil) != tt.isErr {
				t.Errorf("MachineRepository.FindAllMachines: error = %v, isErr = %t", err, tt.isErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MachineRepository.FindAllMachines: got = %v, tt.want = %v", got, tt.want)
			}
		})
	}
}

func Test_CreateMachine(t *testing.T) {
	type repositories struct {
		transactionRepo *mock_repository.MockTransactionRepository
		machineRepo     *mock_repository.MockMachineRepository
		generalRepo     *mock_repository.MockGeneralRepository
	}
	type arguments struct {
		machine *domainMachine.Machine
	}
	tests := []struct {
		name  string
		repos repositories
		args  arguments
		want  error
		isErr bool
	}{
		{
			name: "Successfully",
			repos: repositories{
				transactionRepo: &mock_repository.MockTransactionRepository{
					MockExecWithTx: func(ctx context.Context, f func(context.Context) (interface{}, error)) (interface{}, error) {
						return nil, nil
					},
				},
				machineRepo: &mock_repository.MockMachineRepository{
					MockCreateMachine: func(ctx context.Context, machine *domainMachine.Machine) error {
						return nil
					},
				},
				generalRepo: &mock_repository.MockGeneralRepository{
					MockCreateTableInformation: func(ctx context.Context, tableInfo *domainGeneral.TableInformation) error {
						return nil
					},
				},
			},
			args: arguments{
				machine: &domainMachine.Machine{
					ID:   "0001",
					Name: "test machine 1",
				},
			},
			want:  nil,
			isErr: false,
		},
		{
			name: "Error due to context object does not exist.",
			repos: repositories{
				transactionRepo: &mock_repository.MockTransactionRepository{
					MockExecWithTx: func(ctx context.Context, f func(context.Context) (interface{}, error)) (interface{}, error) {
						return nil, fmt.Errorf("FAILED TO EXECUTE TRANSACTION.")
					},
				},
				machineRepo: &mock_repository.MockMachineRepository{
					MockCreateMachine: func(ctx context.Context, machine *domainMachine.Machine) error {
						return nil
					},
				},
				generalRepo: &mock_repository.MockGeneralRepository{
					MockCreateTableInformation: func(ctx context.Context, tableInfo *domainGeneral.TableInformation) error {
						return nil
					},
				},
			},
			args: arguments{
				machine: &domainMachine.Machine{
					ID:   "0001",
					Name: "test machine 1",
				},
			},
			want:  fmt.Errorf("FAILED TO EXECUTE TRANSACTION."),
			isErr: true,
		},
		{
			name: "Error due to \"CreateMachine\" not executed successfully.",
			repos: repositories{
				transactionRepo: &mock_repository.MockTransactionRepository{
					MockExecWithTx: func(ctx context.Context, f func(context.Context) (interface{}, error)) (interface{}, error) {
						return nil, fmt.Errorf("FAILED TO CREATE MACHINE.")
					},
				},
				machineRepo: &mock_repository.MockMachineRepository{
					MockCreateMachine: func(ctx context.Context, machine *domainMachine.Machine) error {
						return fmt.Errorf("FAILED TO CREATE MACHINE.")
					},
				},
				generalRepo: &mock_repository.MockGeneralRepository{
					MockCreateTableInformation: func(ctx context.Context, tableInfo *domainGeneral.TableInformation) error {
						return nil
					},
				},
			},
			args: arguments{
				machine: &domainMachine.Machine{
					ID:   "0001",
					Name: "test machine 1",
				},
			},
			want:  fmt.Errorf("FAILED TO CREATE MACHINE."),
			isErr: true,
		},
		{
			name: "Error due to \"CreateTableInformation\" not executed successfully.",
			repos: repositories{
				transactionRepo: &mock_repository.MockTransactionRepository{
					MockExecWithTx: func(ctx context.Context, f func(context.Context) (interface{}, error)) (interface{}, error) {
						return nil, fmt.Errorf("FAILED TO CREATE TABLE INFORMATION.")
					},
				},
				machineRepo: &mock_repository.MockMachineRepository{
					MockCreateMachine: func(ctx context.Context, machine *domainMachine.Machine) error {
						return nil
					},
				},
				generalRepo: &mock_repository.MockGeneralRepository{
					MockCreateTableInformation: func(ctx context.Context, tableInfo *domainGeneral.TableInformation) error {
						return fmt.Errorf("FAILED TO CREATE TABLE INFORMATION.")
					},
				},
			},
			args: arguments{
				machine: &domainMachine.Machine{
					ID:   "0001",
					Name: "test machine 1",
				},
			},
			want:  fmt.Errorf("FAILED TO CREATE TABLE INFORMATION."),
			isErr: true,
		},
		{
			name: "Error due to unsuitable argument being passed.",
			repos: repositories{
				transactionRepo: &mock_repository.MockTransactionRepository{
					MockExecWithTx: func(ctx context.Context, f func(context.Context) (interface{}, error)) (interface{}, error) {
						return nil, fmt.Errorf("FAILED TO CREATE MACHINE.")
					},
				},
				machineRepo: &mock_repository.MockMachineRepository{
					MockCreateMachine: func(ctx context.Context, machine *domainMachine.Machine) error {
						return fmt.Errorf("FAILED TO CREATE MACHINE.")
					},
				},
				generalRepo: &mock_repository.MockGeneralRepository{
					MockCreateTableInformation: func(ctx context.Context, tableInfo *domainGeneral.TableInformation) error {
						return nil
					},
				},
			},
			args: arguments{
				machine: nil,
			},
			want:  fmt.Errorf("FAILED TO CREATE MACHINE."),
			isErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRepository := &mock_repository.MockTransactionRepository{
				MockExecWithTx: tt.repos.transactionRepo.ExecWtihTx,
			}
			machineRepository := &mock_repository.MockMachineRepository{
				MockCreateMachine: tt.repos.machineRepo.CreateMachine,
			}
			generalRepository := &mock_repository.MockGeneralRepository{
				MockCreateTableInformation: tt.repos.generalRepo.CreateTableInformation,
			}
			machineUsecase := usecase.NewMachineUsecase(transactionRepository, machineRepository, generalRepository)

			err := machineUsecase.CreateMachine(tt.args.machine)
			if !tt.isErr {
				if err != nil {
					t.Errorf("FAILED TO CREATE MACHINE: %s", err)
				}
			}
			if (err != nil) != tt.isErr {
				t.Errorf("MachineUsecase.CreateMachine: error = %v, tt.want = %v", err, tt.want)
			}
		})
	}
}

func Test_UpdateMachine(t *testing.T) {
	type repositories struct {
		transactionRepo *mock_repository.MockTransactionRepository
		machineRepo     *mock_repository.MockMachineRepository
		generalRepo     *mock_repository.MockGeneralRepository
	}
	type arguments struct {
		machine *domainMachine.Machine
	}
	tests := []struct {
		name  string
		repos repositories
		args  arguments
		want  error
		isErr bool
	}{
		{
			name: "Successfully",
			repos: repositories{
				transactionRepo: &mock_repository.MockTransactionRepository{
					MockExecWithTx: func(ctx context.Context, f func(context.Context) (interface{}, error)) (interface{}, error) {
						return nil, nil
					},
				},
				machineRepo: &mock_repository.MockMachineRepository{
					MockUpdateMachine: func(ctx context.Context, machine *domainMachine.Machine) error {
						return nil
					},
				},
				generalRepo: &mock_repository.MockGeneralRepository{
					MockUpdateTableInformation: func(ctx context.Context, tableInfo *domainGeneral.TableInformation) error {
						return nil
					},
				},
			},
			args: arguments{
				machine: &domainMachine.Machine{
					ID:   "0001",
					Name: "test machine 1",
				},
			},
			want:  nil,
			isErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRepository := &mock_repository.MockTransactionRepository{
				MockExecWithTx: tt.repos.transactionRepo.ExecWtihTx,
			}
			machineRepository := &mock_repository.MockMachineRepository{
				MockUpdateMachine: tt.repos.machineRepo.UpdateMachine,
			}
			generalRepository := &mock_repository.MockGeneralRepository{
				MockUpdateTableInformation: tt.repos.generalRepo.UpdateTableInformation,
			}
			machineUsecase := usecase.NewMachineUsecase(transactionRepository, machineRepository, generalRepository)

			err := machineUsecase.CreateMachine(tt.args.machine)
			if !tt.isErr {
				if err != nil {
					t.Errorf("FAILED TO CREATE MACHINE: %s", err)
				}
			}
			if (err != nil) != tt.isErr {
				t.Errorf("MachineUsecase.CreateMachine: error = %v, tt.want = %v", err, tt.want)
			}
		})
	}
}
