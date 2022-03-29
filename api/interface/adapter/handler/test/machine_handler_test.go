package handler_test

import (
	domainMachine "api/domain/model/machine"
	"api/interface/adapter/handler"
	mock_usecase "api/usecase/mock"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_NewMachineHandler(t *testing.T) {
	usecase := &mock_usecase.MockMachineUsecase{}
	machineHandler := handler.NewMachineHandler(usecase)
	if machineHandler == nil {
		t.Fatalf("FAILED TO TEST; MachineHandler.NewMachineHandler RETURNS nil.")
	}
}

func Test_FindMachineByID(t *testing.T) {
	mockMachineUsecase := &mock_usecase.MockMachineUsecase{
		MockFindMachineByID: func(id domainMachine.MachineID) (*domainMachine.Machine, error) {
			machine := &domainMachine.Machine{
				ID:   domainMachine.MachineID("0001"),
				Name: "test machine 1",
			}
			if id != machine.ID {
				return nil, fmt.Errorf("THE DATA FOR THAT ID(=%s) DOES NOT EXIST.", id)
			}
			return machine, nil
		},
	}

	tests := []struct {
		name               string
		testMachineUsecase *mock_usecase.MockMachineUsecase
		arg                domainMachine.MachineID
		want               *domainMachine.Machine
		status             int
		isErr              bool
		err                error
	}{
		{
			name:               "Successfully",
			testMachineUsecase: mockMachineUsecase,
			arg:                domainMachine.MachineID("0001"),
			want: &domainMachine.Machine{
				ID:   "0001",
				Name: "test machine 1",
			},
			status: http.StatusOK,
			isErr:  false,
			err:    nil,
		},
		{
			name:               "Error",
			testMachineUsecase: mockMachineUsecase,
			arg:                domainMachine.MachineID("XXXX"),
			want:               nil,
			status:             http.StatusBadRequest,
			isErr:              true,
			err:                fmt.Errorf("FAILED TO FIND THE MACHINE DATA."),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			machineUsecase := &mock_usecase.MockMachineUsecase{}
			machineUsecase.MockFindMachineByID = tt.testMachineUsecase.FindMachineByID
			machineHandler := handler.NewMachineHandler(machineUsecase)

			gin.SetMode(gin.TestMode)
			g := gin.New()
			g.GET("/api/machine/:id", machineHandler.GetMachineByID)

			rec := httptest.NewRecorder()
			req, err := http.NewRequest("GET", fmt.Sprintf("/api/machine/%s", tt.arg), nil)
			if err != nil {
				t.Fatalf("FALED TO CREATE HTTP REQUEST: %s", err)
			}
			g.ServeHTTP(rec, req)

			if tt.status == http.StatusOK {
				got, err := machineUsecase.FindMachineByID(tt.arg)
				if err != nil {
					t.Errorf("FAILED TO EXECUTE MockMachineUsecase.FindMachineByID: %s", err)
				}

				errJSON := json.Unmarshal(rec.Body.Bytes(), &got)
				if errJSON != nil {
					t.Errorf("FAILED TO UNMARSHAL JSON: %s", errJSON)
				}

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GET \"api/machine/%s\" RESPONSE = %v, but want = %v", tt.arg, got, tt.want)
				}
			}

			if tt.isErr != (tt.status == http.StatusBadRequest) {
				t.Errorf("ERROR = %#v, but StatusCode = %d", tt.err, tt.status)
			}
		})
	}
}

func Test_FindAllMachines(t *testing.T) {
	mockMachineUsecase := &mock_usecase.MockMachineUsecase{
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
	}

	tests := []struct {
		name               string
		testMachineUsecase *mock_usecase.MockMachineUsecase
		want               []*domainMachine.Machine
		status             int
		isErr              bool
		err                error
	}{
		{
			name:               "Successfully",
			testMachineUsecase: mockMachineUsecase,
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
			name:               "Error",
			testMachineUsecase: mockMachineUsecase,
			want:               nil,
			status:             http.StatusBadRequest,
			isErr:              true,
			err:                fmt.Errorf("FAILED TO FIND ALL MACHINES DATA."),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			machineUsecase := &mock_usecase.MockMachineUsecase{}
			machineUsecase.MockFindAllMachines = tt.testMachineUsecase.FindAllMachines
			machineHandler := handler.NewMachineHandler(machineUsecase)

			gin.SetMode(gin.TestMode)
			g := gin.New()
			g.GET("/api/machine/all", machineHandler.GetAllMachines)

			rec := httptest.NewRecorder()
			req, err := http.NewRequest("GET", fmt.Sprintf("/api/machine/all"), nil)
			if err != nil {
				t.Fatalf("FALED TO CREATE HTTP REQUEST: %s", err)
			}
			g.ServeHTTP(rec, req)

			if tt.status == http.StatusOK {
				got, err := machineUsecase.FindAllMachines()
				if err != nil {
					t.Errorf("FAILED TO EXECUTE MockMachineUsecase.FindMachineByID: %s", err)
				}

				errJSON := json.Unmarshal(rec.Body.Bytes(), &got)
				if errJSON != nil {
					t.Errorf("FAILED TO UNMARSHAL JSON: %s", errJSON)
				}

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GET \"api/machine/all\" RESPONSE = %v, but want = %v", got, tt.want)
				}
			}

			if tt.isErr != (tt.status == http.StatusBadRequest) {
				t.Errorf("ERROR = %#v, but StatusCode = %d", tt.err, tt.status)
			}
		})
	}
}

func Test_CreateMachine(t *testing.T) {
	const (
		EXISTENT_ID string = "ALREADY EXISTS ID"
		VALUE       string = "the value to something"
	)
	key := struct{}{}
	testCtx := context.Background()
	testCtx = context.WithValue(testCtx, &key, VALUE)

	mockMachineUsecase := &mock_usecase.MockMachineUsecase{
		MockCreateMachine: func(ctx context.Context, machine *domainMachine.Machine) error {
			existMachine := &domainMachine.Machine{
				ID: domainMachine.MachineID("ALREADY EXISTS ID"),
			}

			if ctx.Value(&key) == nil {
				return fmt.Errorf("FAILED TO GET THE CONTEXT OBJECT.")
			}

			if machine.ID == existMachine.ID {
				return fmt.Errorf("THE MACHINE'S ID ALREADY EXISTS.")
			}
			return nil
		},
	}

	type arguments struct {
		ctx     context.Context
		machine *domainMachine.Machine
	}
	tests := []struct {
		name               string
		testMachineUsecase *mock_usecase.MockMachineUsecase
		args               arguments
		want               error
		status             int
		isErr              bool
	}{
		{
			name:               "Successfully",
			testMachineUsecase: mockMachineUsecase,
			args: arguments{
				ctx: testCtx,
				machine: &domainMachine.Machine{
					ID: domainMachine.MachineID("0001"),
				},
			},
			want:   nil,
			status: http.StatusOK,
			isErr:  false,
		},
		{
			name:               "Error: Context does not exits.",
			testMachineUsecase: mockMachineUsecase,
			args: arguments{
				ctx: nil,
				machine: &domainMachine.Machine{
					ID: domainMachine.MachineID("0001"),
				},
			},
			want:   fmt.Errorf("FAILED TO GET THE CONTEXT OBJECT."),
			status: http.StatusBadRequest,
			isErr:  true,
		},
		{
			name:               "Error: Machine's ID already exists.",
			testMachineUsecase: mockMachineUsecase,
			args: arguments{
				ctx: testCtx,
				machine: &domainMachine.Machine{
					ID: domainMachine.MachineID(EXISTENT_ID),
				},
			},
			want:   fmt.Errorf("THE MACHINE'S ID ALREADY EXISTS."),
			status: http.StatusBadRequest,
			isErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			machineUsecase := &mock_usecase.MockMachineUsecase{}
			machineUsecase.MockCreateMachine = tt.testMachineUsecase.CreateMachine
			machineHandler := handler.NewMachineHandler(machineUsecase)

			gin.SetMode(gin.TestMode)
			g := gin.New()
			g.POST("/api/machine", machineHandler.CreateMachine)

			rec := httptest.NewRecorder()
			req, err := http.NewRequest("POST", fmt.Sprintf("/api/machine"), nil)
			if err != nil {
				t.Fatalf("FALED TO CREATE HTTP REQUEST: %s", err)
			}
			g.ServeHTTP(rec, req)

			if tt.status == http.StatusOK {
				err := machineUsecase.CreateMachine(tt.args.ctx, tt.args.machine)
				if err != nil {
					t.Errorf("FAILED TO EXECUTE MockMachineUsecase.CreateMachine: %s", err)
				}
			}

			if tt.isErr != (tt.status == http.StatusBadRequest) {
				t.Errorf("ERROR = %#v, but StatusCode = %d", tt.want, tt.status)
			}

		})
	}
}

func Test_UpdateMachine(t *testing.T) {
	const (
		VALUE string = "the value to something"
	)
	key := struct{}{}
	testCtx := context.Background()
	testCtx = context.WithValue(testCtx, &key, VALUE)

	mockMachineUsecase := &mock_usecase.MockMachineUsecase{
		MockUpdateMachine: func(ctx context.Context, machine *domainMachine.Machine) error {
			m := &domainMachine.Machine{
				ID: domainMachine.MachineID("0001"),
			}

			if ctx.Value(&key) == nil {
				return fmt.Errorf("FAILED TO GET THE CONTEXT OBJECT.")
			}

			if machine.ID != m.ID {
				return fmt.Errorf("THE MACHINE'S ID DOES NOT EXIST.")
			}

			return nil
		},
	}

	type arguments struct {
		ctx     context.Context
		machine *domainMachine.Machine
	}
	tests := []struct {
		name               string
		testMachineUsecase *mock_usecase.MockMachineUsecase
		args               arguments
		want               error
		status             int
		isErr              bool
	}{
		{
			name:               "Successfully",
			testMachineUsecase: mockMachineUsecase,
			args: arguments{
				ctx: testCtx,
				machine: &domainMachine.Machine{
					ID: domainMachine.MachineID("0001"),
				},
			},
			want:   nil,
			status: http.StatusOK,
			isErr:  false,
		},
		{
			name:               "Error: Context does not exits.",
			testMachineUsecase: mockMachineUsecase,
			args: arguments{
				ctx: nil,
				machine: &domainMachine.Machine{
					ID: domainMachine.MachineID("0001"),
				},
			},
			want:   fmt.Errorf("FAILED TO GET THE CONTEXT OBJECT."),
			status: http.StatusBadRequest,
			isErr:  true,
		},
		{
			name:               "Error: Machine's ID does not exist.",
			testMachineUsecase: mockMachineUsecase,
			args: arguments{
				ctx: testCtx,
				machine: &domainMachine.Machine{
					ID: domainMachine.MachineID("XXXX"),
				},
			},
			want:   fmt.Errorf("THE MACHINE'S ID DOES NOT EXIST."),
			status: http.StatusBadRequest,
			isErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			machineUsecase := &mock_usecase.MockMachineUsecase{}
			machineUsecase.MockUpdateMachine = tt.testMachineUsecase.UpdateMachine
			machineHandler := handler.NewMachineHandler(machineUsecase)

			gin.SetMode(gin.TestMode)
			g := gin.New()
			g.PUT("/api/machine", machineHandler.UpdateMachine)

			rec := httptest.NewRecorder()
			req, err := http.NewRequest("PUT", fmt.Sprintf("/api/machine"), nil)
			if err != nil {
				t.Fatalf("FALED TO CREATE HTTP REQUEST: %s", err)
			}
			g.ServeHTTP(rec, req)

			if tt.status == http.StatusOK {
				err := machineUsecase.UpdateMachine(tt.args.ctx, tt.args.machine)
				if err != nil {
					t.Errorf("FAILED TO EXECUTE MockMachineUsecase.CreateMachine: %s", err)
				}
			}

			if tt.isErr != (tt.status == http.StatusBadRequest) {
				t.Errorf("ERROR = %#v, but StatusCode = %d", tt.want, tt.status)
			}

		})
	}
}

func Test_DeleteMachine(t *testing.T) {
	const (
		VALUE string = "the value to something"
	)
	key := struct{}{}
	testCtx := context.Background()
	testCtx = context.WithValue(testCtx, &key, VALUE)

	mockMachineUsecase := &mock_usecase.MockMachineUsecase{
		MockStopUsingMachine: func(ctx context.Context, machine *domainMachine.Machine) error {
			m := &domainMachine.Machine{
				ID: domainMachine.MachineID("0001"),
			}

			if ctx.Value(&key) == nil {
				return fmt.Errorf("FAILED TO GET THE CONTEXT OBJECT.")
			}

			if machine.ID != m.ID {
				return fmt.Errorf("THE MACHINE'S ID DOES NOT EXIST.")
			}

			return nil
		},
	}

	type arguments struct {
		ctx     context.Context
		machine *domainMachine.Machine
	}
	tests := []struct {
		name               string
		testMachineUsecase *mock_usecase.MockMachineUsecase
		args               arguments
		want               error
		status             int
		isErr              bool
	}{
		{
			name:               "Successfully",
			testMachineUsecase: mockMachineUsecase,
			args: arguments{
				ctx: testCtx,
				machine: &domainMachine.Machine{
					ID: domainMachine.MachineID("0001"),
				},
			},
			want:   nil,
			status: http.StatusOK,
			isErr:  false,
		},
		{
			name:               "Error: Context does not exits.",
			testMachineUsecase: mockMachineUsecase,
			args: arguments{
				ctx: nil,
				machine: &domainMachine.Machine{
					ID: domainMachine.MachineID("0001"),
				},
			},
			want:   fmt.Errorf("FAILED TO GET THE CONTEXT OBJECT."),
			status: http.StatusBadRequest,
			isErr:  true,
		},
		{
			name:               "Error: Machine's ID does not exist.",
			testMachineUsecase: mockMachineUsecase,
			args: arguments{
				ctx: testCtx,
				machine: &domainMachine.Machine{
					ID: domainMachine.MachineID("XXXX"),
				},
			},
			want:   fmt.Errorf("THE MACHINE'S ID DOES NOT EXIST."),
			status: http.StatusBadRequest,
			isErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			machineUsecase := &mock_usecase.MockMachineUsecase{}
			machineUsecase.MockStopUsingMachine = tt.testMachineUsecase.StopUsingMachine
			machineHandler := handler.NewMachineHandler(machineUsecase)

			gin.SetMode(gin.TestMode)
			g := gin.New()
			g.PUT("/api/machine/stop", machineHandler.StopUsingMachine)

			rec := httptest.NewRecorder()
			req, err := http.NewRequest("PUT", fmt.Sprintf("/api/machine/stop"), nil)
			if err != nil {
				t.Fatalf("FALED TO CREATE HTTP REQUEST: %s", err)
			}
			g.ServeHTTP(rec, req)

			if tt.status == http.StatusOK {
				err := machineUsecase.StopUsingMachine(tt.args.ctx, tt.args.machine)
				if err != nil {
					t.Errorf("FAILED TO EXECUTE MockMachineUsecase.StopUsingMachine: %s", err)
				}
			}

			if tt.isErr != (tt.status == http.StatusBadRequest) {
				t.Errorf("ERROR = %#v, but StatusCode = %d", tt.want, tt.status)
			}

		})
	}
}
