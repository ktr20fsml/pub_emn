package handler_test

import (
	domainMachine "api/domain/model/machine"
	mock_service "api/domain/service/mock"
	"api/interface/adapter/handler"
	mock_usecase "api/usecase/mock"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Test_NewMachineHandler(t *testing.T) {
	usecase := &mock_usecase.MockMachineUsecase{}
	utility := &mock_service.MockUtilityService{}
	machineHandler := handler.NewMachineHandler(usecase, utility)
	if machineHandler == nil {
		t.Fatalf("FAILED TO TEST: MachineHandler.NewMachineHandler RETURNS nil.")
	}
}

func Test_GetMachineByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	type mocks struct {
		machineUsecase *mock_usecase.MockMachineUsecase
		utilityService *mock_service.MockUtilityService
	}
	tests := []struct {
		name   string
		mock   mocks
		arg    domainMachine.MachineID
		want   *domainMachine.Machine
		status int
		isErr  bool
		err    error
	}{
		{
			name: "Successfully",
			mock: mocks{
				machineUsecase: &mock_usecase.MockMachineUsecase{
					MockFindMachineByID: func(id domainMachine.MachineID) (*domainMachine.Machine, error) {
						machine := &domainMachine.Machine{
							ID:   domainMachine.MachineID("0001"),
							Name: "test machine 1",
						}

						return machine, nil
					},
				},
				utilityService: &mock_service.MockUtilityService{},
			},
			arg: domainMachine.MachineID("0001"),
			want: &domainMachine.Machine{
				ID:   "0001",
				Name: "test machine 1",
			},
			status: http.StatusOK,
			isErr:  false,
			err:    nil,
		},
		{
			name: "Error",
			mock: mocks{
				machineUsecase: &mock_usecase.MockMachineUsecase{
					MockFindMachineByID: func(id domainMachine.MachineID) (*domainMachine.Machine, error) {
						return nil, fmt.Errorf("FAILED TO FIND THE MACHINE DATA.")
					},
				},
				utilityService: &mock_service.MockUtilityService{},
			},
			arg:    domainMachine.MachineID("XXXX"),
			want:   nil,
			status: http.StatusBadRequest,
			isErr:  true,
			err:    fmt.Errorf("FAILED TO FIND THE MACHINE DATA."),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			machineUsecase := &mock_usecase.MockMachineUsecase{}
			machineUsecase.MockFindMachineByID = tt.mock.machineUsecase.FindMachineByID
			utilityService := tt.mock.utilityService
			machineHandler := handler.NewMachineHandler(machineUsecase, utilityService)

			g := gin.New()
			g.GET("/api/machine/:id", machineHandler.GetMachineByID)

			rec := httptest.NewRecorder()
			req, err := http.NewRequest("GET", fmt.Sprintf("/api/machine/%s", tt.arg), nil)
			if err != nil {
				t.Fatalf("FALED TO CREATE HTTP REQUEST: %s", err)
			}
			g.ServeHTTP(rec, req)

			if tt.status == http.StatusOK {
				machine := &domainMachine.Machine{}
				errJSON := json.Unmarshal(rec.Body.Bytes(), machine)
				if errJSON != nil {
					t.Errorf("FAILED TO UNMARSHAL JSON: %s", errJSON)
				}

				if !reflect.DeepEqual(machine, tt.want) {
					t.Errorf("GET \"api/machine/%s\" RESPONSE = %v, but want = %v", tt.arg, machine, tt.want)
				}
			}

			if tt.isErr != (tt.status == http.StatusBadRequest) {
				t.Errorf("ERROR = %#v, but StatusCode = %d", tt.err, tt.status)
			}
		})
	}
}

func Test_GetAllMachines(t *testing.T) {
	gin.SetMode(gin.TestMode)

	type mocks struct {
		machineUsecase *mock_usecase.MockMachineUsecase
		utilityService *mock_service.MockUtilityService
	}
	tests := []struct {
		name   string
		mock   mocks
		want   []*domainMachine.Machine
		status int
		isErr  bool
		err    error
	}{
		{
			name: "Successfully",
			mock: mocks{
				machineUsecase: &mock_usecase.MockMachineUsecase{
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
				utilityService: &mock_service.MockUtilityService{},
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
			status: http.StatusOK,
			isErr:  false,
			err:    nil,
		},
		{
			name: "Error due to non-existent ID search.",
			mock: mocks{
				machineUsecase: &mock_usecase.MockMachineUsecase{
					MockFindAllMachines: func() ([]*domainMachine.Machine, error) {
						return nil, fmt.Errorf("FAILED TO FIND ALL MACHINES DATA.")
					},
				},
				utilityService: &mock_service.MockUtilityService{},
			},
			want:   nil,
			status: http.StatusBadRequest,
			isErr:  true,
			err:    fmt.Errorf("FAILED TO FIND ALL MACHINES DATA."),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			machineUsecase := &mock_usecase.MockMachineUsecase{}
			machineUsecase.MockFindAllMachines = tt.mock.machineUsecase.FindAllMachines
			utilityService := tt.mock.utilityService
			machineHandler := handler.NewMachineHandler(machineUsecase, utilityService)

			g := gin.New()
			g.GET("/api/machine/all", machineHandler.GetAllMachines)

			rec := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "/api/machine/all", nil)
			if err != nil {
				t.Fatalf("FALED TO CREATE HTTP REQUEST: %s", err)
			}
			g.ServeHTTP(rec, req)

			if tt.status == http.StatusOK {
				machines := []*domainMachine.Machine{}
				if err != nil {
					t.Errorf("FAILED TO EXECUTE MockMachineUsecase.FindMachineByID: %s", err)
				}

				errJSON := json.Unmarshal(rec.Body.Bytes(), &machines)
				if errJSON != nil {
					t.Errorf("FAILED TO UNMARSHAL JSON: %s", errJSON)
				}

				for i, v := range tt.want {
					if !reflect.DeepEqual(machines[i], v) {
						t.Errorf("GET \"api/machine/all\" RESPONSE = %v, but want = %v", machines[i], v)
					}
				}
			}

			if tt.isErr != (tt.status == http.StatusBadRequest) {
				t.Errorf("ERROR = %#v, but StatusCode = %d", tt.err, tt.status)
			}
		})
	}
}

func Test_CreateMachine(t *testing.T) {
	const EXISTENT_ID string = "ALREADY EXISTS ID"
	type mocks struct {
		machineUsecase *mock_usecase.MockMachineUsecase
		utilityService *mock_service.MockUtilityService
	}
	type arguments struct {
		machine *domainMachine.Machine
	}
	type want struct {
		status int
		err    error
	}
	tests := []struct {
		name  string
		mock  mocks
		args  arguments
		want  want
		isErr bool
	}{
		{
			name: "Successfully",
			mock: mocks{
				machineUsecase: &mock_usecase.MockMachineUsecase{
					MockCreateMachine: func(machine *domainMachine.Machine) error {
						return nil
					},
				},
				utilityService: &mock_service.MockUtilityService{},
			},
			args: arguments{
				machine: &domainMachine.Machine{
					ID: domainMachine.MachineID("0001"),
				},
			},
			want: want{
				status: http.StatusOK,
				err:    nil,
			},
			isErr: false,
		},
		{
			name: "Error due to context does not exits.",
			mock: mocks{
				machineUsecase: &mock_usecase.MockMachineUsecase{
					MockCreateMachine: func(machine *domainMachine.Machine) error {
						return fmt.Errorf("FAILED TO GET THE CONTEXT OBJECT.")
					},
				},
				utilityService: &mock_service.MockUtilityService{},
			},
			args: arguments{
				machine: &domainMachine.Machine{
					ID: domainMachine.MachineID("0001"),
				},
			},
			want: want{
				status: http.StatusBadRequest,
				err:    fmt.Errorf("FAILED TO GET THE CONTEXT OBJECT."),
			},
			isErr: true,
		},
		{
			name: "Error due to machine's ID already exists.",
			mock: mocks{
				machineUsecase: &mock_usecase.MockMachineUsecase{
					MockCreateMachine: func(machine *domainMachine.Machine) error {
						existMachine := &domainMachine.Machine{
							ID: domainMachine.MachineID(EXISTENT_ID),
						}
						if machine.ID == existMachine.ID {
							return fmt.Errorf("THE MACHINE'S ID ALREADY EXISTS.")
						}
						return nil
					},
				},
				utilityService: &mock_service.MockUtilityService{},
			},
			args: arguments{
				machine: &domainMachine.Machine{
					ID: domainMachine.MachineID(EXISTENT_ID),
				},
			},
			want: want{
				status: http.StatusBadRequest,
				err:    fmt.Errorf("THE MACHINE'S ID ALREADY EXISTS."),
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			machineUsecase := &mock_usecase.MockMachineUsecase{}
			machineUsecase.MockCreateMachine = tt.mock.machineUsecase.CreateMachine
			utilityService := tt.mock.utilityService
			machineHandler := handler.NewMachineHandler(machineUsecase, utilityService)

			g := gin.New()
			store := cookie.NewStore([]byte("secret"))
			g.Use(sessions.Sessions("emn-session", store))
			g.POST("/api/machine", machineHandler.CreateMachine)

			w := httptest.NewRecorder()
			req, errNewReq := http.NewRequest(http.MethodPost, "/api/machine", nil)
			if errNewReq != nil {
				t.Fatal(errNewReq)
			}
			g.ServeHTTP(w, req)

			if tt.want.status == http.StatusOK {
				errCreate := tt.mock.machineUsecase.MockCreateMachine(tt.args.machine)
				if errCreate != nil {
					t.Error(errCreate)
				}

			}
			if tt.isErr != (tt.want.status == http.StatusBadRequest) {
				t.Errorf("ERROR = %#v, but StatusCode = %d", tt.want, tt.want.status)
			}
		})
	}
}

func Test_UpdateMachine(t *testing.T) {
	const (
		EXISTENT_ID     string = "EXISTENT ID"
		NON_EXISTENT_ID string = "NON EXISTENT ID"
	)
	type mocks struct {
		machineUsecase *mock_usecase.MockMachineUsecase
		utilityService *mock_service.MockUtilityService
	}
	type arguments struct {
		machine *domainMachine.Machine
	}
	type want struct {
		status int
		err    error
	}
	tests := []struct {
		name  string
		mock  mocks
		args  arguments
		want  want
		isErr bool
	}{
		{
			name: "Successfully",
			mock: mocks{
				machineUsecase: &mock_usecase.MockMachineUsecase{
					MockUpdateMachine: func(machine *domainMachine.Machine) error {
						return nil
					},
				},
				utilityService: &mock_service.MockUtilityService{},
			},
			args: arguments{
				machine: &domainMachine.Machine{
					ID: domainMachine.MachineID("0001"),
				},
			},
			want: want{
				status: http.StatusOK,
				err:    nil,
			},
			isErr: false,
		},
		{
			name: "Error due to context does not exits.",
			mock: mocks{
				machineUsecase: &mock_usecase.MockMachineUsecase{
					MockUpdateMachine: func(machine *domainMachine.Machine) error {
						return fmt.Errorf("FAILED TO GET THE CONTEXT OBJECT.")
					},
				},
				utilityService: &mock_service.MockUtilityService{},
			},
			args: arguments{
				machine: &domainMachine.Machine{
					ID: domainMachine.MachineID("0001"),
				},
			},
			want: want{
				status: http.StatusBadRequest,
				err:    fmt.Errorf("FAILED TO GET THE CONTEXT OBJECT."),
			},
			isErr: true,
		},
		{
			name: "Error due to machine's ID does not exists.",
			mock: mocks{
				machineUsecase: &mock_usecase.MockMachineUsecase{
					MockUpdateMachine: func(machine *domainMachine.Machine) error {
						existMachine := &domainMachine.Machine{
							ID: domainMachine.MachineID(EXISTENT_ID),
						}
						if machine.ID != existMachine.ID {
							return fmt.Errorf("THE MACHINE'S ID DOES NOT EXISTS.")
						}
						return nil
					},
				},
				utilityService: &mock_service.MockUtilityService{},
			},
			args: arguments{
				machine: &domainMachine.Machine{
					ID: domainMachine.MachineID(NON_EXISTENT_ID),
				},
			},
			want: want{
				status: http.StatusBadRequest,
				err:    fmt.Errorf("THE MACHINE'S ID ALREADY EXISTS."),
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			machineUsecase := &mock_usecase.MockMachineUsecase{}
			machineUsecase.MockUpdateMachine = tt.mock.machineUsecase.UpdateMachine
			utilityService := tt.mock.utilityService
			machineHandler := handler.NewMachineHandler(machineUsecase, utilityService)

			g := gin.New()
			store := cookie.NewStore([]byte("secret"))
			g.Use(sessions.Sessions("emn-session", store))
			g.POST("/api/machine", machineHandler.UpdateMachine)

			w := httptest.NewRecorder()
			req, errNewReq := http.NewRequest(http.MethodPut, "/api/machine", nil)
			if errNewReq != nil {
				t.Fatal(errNewReq)
			}
			g.ServeHTTP(w, req)

			if tt.want.status == http.StatusOK {
				errUsecase := tt.mock.machineUsecase.MockUpdateMachine(tt.args.machine)
				if errUsecase != nil {
					t.Error(errUsecase)
				}

			}
			if tt.isErr != (tt.want.status == http.StatusBadRequest) {
				t.Errorf("ERROR = %#v, but StatusCode = %d", tt.want, tt.want.status)
			}
		})
	}
}

func Test_DeleteMachine(t *testing.T) {
	const (
		EXISTENT_ID     string = "EXISTENT ID"
		NON_EXISTENT_ID string = "NON EXISTENT ID"
	)
	type mocks struct {
		machineUsecase *mock_usecase.MockMachineUsecase
		utilityService *mock_service.MockUtilityService
	}
	type arguments struct {
		machine *domainMachine.Machine
	}
	type want struct {
		status int
		err    error
	}
	tests := []struct {
		name  string
		mock  mocks
		args  arguments
		want  want
		isErr bool
	}{
		{
			name: "Successfully",
			mock: mocks{
				machineUsecase: &mock_usecase.MockMachineUsecase{
					MockStopUsingMachine: func(machine *domainMachine.Machine) error {
						return nil
					},
				},
				utilityService: &mock_service.MockUtilityService{},
			},
			args: arguments{
				machine: &domainMachine.Machine{
					ID: domainMachine.MachineID("0001"),
				},
			},
			want: want{
				status: http.StatusOK,
				err:    nil,
			},
			isErr: false,
		},
		{
			name: "Error due to context does not exits.",
			mock: mocks{
				&mock_usecase.MockMachineUsecase{
					MockUpdateMachine: func(machine *domainMachine.Machine) error {
						return fmt.Errorf("FAILED TO GET THE CONTEXT OBJECT.")
					},
				},
				&mock_service.MockUtilityService{},
			},
			args: arguments{
				machine: &domainMachine.Machine{
					ID: domainMachine.MachineID("0001"),
				},
			},
			want: want{
				status: http.StatusBadRequest,
				err:    fmt.Errorf("FAILED TO GET THE CONTEXT OBJECT."),
			},
			isErr: true,
		},
		{
			name: "Error due to machine's ID does not exists.",
			mock: mocks{
				machineUsecase: &mock_usecase.MockMachineUsecase{
					MockStopUsingMachine: func(machine *domainMachine.Machine) error {
						existMachine := &domainMachine.Machine{
							ID: domainMachine.MachineID(EXISTENT_ID),
						}
						if machine.ID != existMachine.ID {
							return fmt.Errorf("THE MACHINE'S ID DOES NOT EXISTS.")
						}
						return nil
					},
				},
				utilityService: &mock_service.MockUtilityService{},
			},
			args: arguments{
				machine: &domainMachine.Machine{
					ID: domainMachine.MachineID(NON_EXISTENT_ID),
				},
			},
			want: want{
				status: http.StatusBadRequest,
				err:    fmt.Errorf("THE MACHINE'S ID ALREADY EXISTS."),
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			machineUsecase := &mock_usecase.MockMachineUsecase{}
			machineUsecase.MockStopUsingMachine = tt.mock.machineUsecase.StopUsingMachine
			utilityService := tt.mock.utilityService
			machineHandler := handler.NewMachineHandler(machineUsecase, utilityService)

			g := gin.New()
			store := cookie.NewStore([]byte("secret"))
			g.Use(sessions.Sessions("emn-session", store))
			g.POST("/api/machine", machineHandler.StopUsingMachine)

			w := httptest.NewRecorder()
			req, errNewReq := http.NewRequest(http.MethodPut, "/api/machine", nil)
			if errNewReq != nil {
				t.Fatal(errNewReq)
			}
			g.ServeHTTP(w, req)

			if tt.want.status == http.StatusOK {
				errUsecase := tt.mock.machineUsecase.MockStopUsingMachine(tt.args.machine)
				if errUsecase != nil {
					t.Error(errUsecase)
				}

			}
			if tt.isErr != (tt.want.status == http.StatusBadRequest) {
				t.Errorf("ERROR = %#v, but StatusCode = %d", tt.want, tt.want.status)
			}
		})
	}
}
