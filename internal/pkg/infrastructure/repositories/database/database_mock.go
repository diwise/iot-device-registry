// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package database

import (
	"github.com/diwise/iot-device-registry/internal/pkg/infrastructure/repositories/models"
	"github.com/diwise/ngsi-ld-golang/pkg/datamodels/fiware"
	"sync"
)

// Ensure, that DatastoreMock does implement Datastore.
// If this is not the case, regenerate this file with moq.
var _ Datastore = &DatastoreMock{}

// DatastoreMock is a mock implementation of Datastore.
//
// 	func TestSomethingThatUsesDatastore(t *testing.T) {
//
// 		// make and configure a mocked Datastore
// 		mockedDatastore := &DatastoreMock{
// 			CreateDeviceFunc: func(device *fiware.Device) (*models.Device, error) {
// 				panic("mock out the CreateDevice method")
// 			},
// 			CreateDeviceModelFunc: func(deviceModel *fiware.DeviceModel) (*models.DeviceModel, error) {
// 				panic("mock out the CreateDeviceModel method")
// 			},
// 			GetDeviceFromIDFunc: func(id string) (*models.Device, error) {
// 				panic("mock out the GetDeviceFromID method")
// 			},
// 			GetDeviceModelFromIDFunc: func(id string) (*models.DeviceModel, error) {
// 				panic("mock out the GetDeviceModelFromID method")
// 			},
// 			GetDeviceModelFromPrimaryKeyFunc: func(id uint) (*models.DeviceModel, error) {
// 				panic("mock out the GetDeviceModelFromPrimaryKey method")
// 			},
// 			GetDeviceModelsFunc: func() ([]models.DeviceModel, error) {
// 				panic("mock out the GetDeviceModels method")
// 			},
// 			GetDevicesFunc: func() ([]models.Device, error) {
// 				panic("mock out the GetDevices method")
// 			},
// 			UpdateDeviceLocationFunc: func(deviceID string, lat float64, lon float64) error {
// 				panic("mock out the UpdateDeviceLocation method")
// 			},
// 			UpdateDeviceStateFunc: func(deviceID string, state string) error {
// 				panic("mock out the UpdateDeviceState method")
// 			},
// 			UpdateDeviceValueFunc: func(deviceID string, value string) error {
// 				panic("mock out the UpdateDeviceValue method")
// 			},
// 		}
//
// 		// use mockedDatastore in code that requires Datastore
// 		// and then make assertions.
//
// 	}
type DatastoreMock struct {
	// CreateDeviceFunc mocks the CreateDevice method.
	CreateDeviceFunc func(device *fiware.Device) (*models.Device, error)

	// CreateDeviceModelFunc mocks the CreateDeviceModel method.
	CreateDeviceModelFunc func(deviceModel *fiware.DeviceModel) (*models.DeviceModel, error)

	// GetDeviceFromIDFunc mocks the GetDeviceFromID method.
	GetDeviceFromIDFunc func(id string) (*models.Device, error)

	// GetDeviceModelFromIDFunc mocks the GetDeviceModelFromID method.
	GetDeviceModelFromIDFunc func(id string) (*models.DeviceModel, error)

	// GetDeviceModelFromPrimaryKeyFunc mocks the GetDeviceModelFromPrimaryKey method.
	GetDeviceModelFromPrimaryKeyFunc func(id uint) (*models.DeviceModel, error)

	// GetDeviceModelsFunc mocks the GetDeviceModels method.
	GetDeviceModelsFunc func() ([]models.DeviceModel, error)

	// GetDevicesFunc mocks the GetDevices method.
	GetDevicesFunc func() ([]models.Device, error)

	// UpdateDeviceLocationFunc mocks the UpdateDeviceLocation method.
	UpdateDeviceLocationFunc func(deviceID string, lat float64, lon float64) error

	// UpdateDeviceStateFunc mocks the UpdateDeviceState method.
	UpdateDeviceStateFunc func(deviceID string, state string) error

	// UpdateDeviceValueFunc mocks the UpdateDeviceValue method.
	UpdateDeviceValueFunc func(deviceID string, value string) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateDevice holds details about calls to the CreateDevice method.
		CreateDevice []struct {
			// Device is the device argument value.
			Device *fiware.Device
		}
		// CreateDeviceModel holds details about calls to the CreateDeviceModel method.
		CreateDeviceModel []struct {
			// DeviceModel is the deviceModel argument value.
			DeviceModel *fiware.DeviceModel
		}
		// GetDeviceFromID holds details about calls to the GetDeviceFromID method.
		GetDeviceFromID []struct {
			// ID is the id argument value.
			ID string
		}
		// GetDeviceModelFromID holds details about calls to the GetDeviceModelFromID method.
		GetDeviceModelFromID []struct {
			// ID is the id argument value.
			ID string
		}
		// GetDeviceModelFromPrimaryKey holds details about calls to the GetDeviceModelFromPrimaryKey method.
		GetDeviceModelFromPrimaryKey []struct {
			// ID is the id argument value.
			ID uint
		}
		// GetDeviceModels holds details about calls to the GetDeviceModels method.
		GetDeviceModels []struct {
		}
		// GetDevices holds details about calls to the GetDevices method.
		GetDevices []struct {
		}
		// UpdateDeviceLocation holds details about calls to the UpdateDeviceLocation method.
		UpdateDeviceLocation []struct {
			// DeviceID is the deviceID argument value.
			DeviceID string
			// Lat is the lat argument value.
			Lat float64
			// Lon is the lon argument value.
			Lon float64
		}
		// UpdateDeviceState holds details about calls to the UpdateDeviceState method.
		UpdateDeviceState []struct {
			// DeviceID is the deviceID argument value.
			DeviceID string
			// State is the state argument value.
			State string
		}
		// UpdateDeviceValue holds details about calls to the UpdateDeviceValue method.
		UpdateDeviceValue []struct {
			// DeviceID is the deviceID argument value.
			DeviceID string
			// Value is the value argument value.
			Value string
		}
	}
	lockCreateDevice                 sync.RWMutex
	lockCreateDeviceModel            sync.RWMutex
	lockGetDeviceFromID              sync.RWMutex
	lockGetDeviceModelFromID         sync.RWMutex
	lockGetDeviceModelFromPrimaryKey sync.RWMutex
	lockGetDeviceModels              sync.RWMutex
	lockGetDevices                   sync.RWMutex
	lockUpdateDeviceLocation         sync.RWMutex
	lockUpdateDeviceState            sync.RWMutex
	lockUpdateDeviceValue            sync.RWMutex
}

// CreateDevice calls CreateDeviceFunc.
func (mock *DatastoreMock) CreateDevice(device *fiware.Device) (*models.Device, error) {
	if mock.CreateDeviceFunc == nil {
		panic("DatastoreMock.CreateDeviceFunc: method is nil but Datastore.CreateDevice was just called")
	}
	callInfo := struct {
		Device *fiware.Device
	}{
		Device: device,
	}
	mock.lockCreateDevice.Lock()
	mock.calls.CreateDevice = append(mock.calls.CreateDevice, callInfo)
	mock.lockCreateDevice.Unlock()
	return mock.CreateDeviceFunc(device)
}

// CreateDeviceCalls gets all the calls that were made to CreateDevice.
// Check the length with:
//     len(mockedDatastore.CreateDeviceCalls())
func (mock *DatastoreMock) CreateDeviceCalls() []struct {
	Device *fiware.Device
} {
	var calls []struct {
		Device *fiware.Device
	}
	mock.lockCreateDevice.RLock()
	calls = mock.calls.CreateDevice
	mock.lockCreateDevice.RUnlock()
	return calls
}

// CreateDeviceModel calls CreateDeviceModelFunc.
func (mock *DatastoreMock) CreateDeviceModel(deviceModel *fiware.DeviceModel) (*models.DeviceModel, error) {
	if mock.CreateDeviceModelFunc == nil {
		panic("DatastoreMock.CreateDeviceModelFunc: method is nil but Datastore.CreateDeviceModel was just called")
	}
	callInfo := struct {
		DeviceModel *fiware.DeviceModel
	}{
		DeviceModel: deviceModel,
	}
	mock.lockCreateDeviceModel.Lock()
	mock.calls.CreateDeviceModel = append(mock.calls.CreateDeviceModel, callInfo)
	mock.lockCreateDeviceModel.Unlock()
	return mock.CreateDeviceModelFunc(deviceModel)
}

// CreateDeviceModelCalls gets all the calls that were made to CreateDeviceModel.
// Check the length with:
//     len(mockedDatastore.CreateDeviceModelCalls())
func (mock *DatastoreMock) CreateDeviceModelCalls() []struct {
	DeviceModel *fiware.DeviceModel
} {
	var calls []struct {
		DeviceModel *fiware.DeviceModel
	}
	mock.lockCreateDeviceModel.RLock()
	calls = mock.calls.CreateDeviceModel
	mock.lockCreateDeviceModel.RUnlock()
	return calls
}

// GetDeviceFromID calls GetDeviceFromIDFunc.
func (mock *DatastoreMock) GetDeviceFromID(id string) (*models.Device, error) {
	if mock.GetDeviceFromIDFunc == nil {
		panic("DatastoreMock.GetDeviceFromIDFunc: method is nil but Datastore.GetDeviceFromID was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	mock.lockGetDeviceFromID.Lock()
	mock.calls.GetDeviceFromID = append(mock.calls.GetDeviceFromID, callInfo)
	mock.lockGetDeviceFromID.Unlock()
	return mock.GetDeviceFromIDFunc(id)
}

// GetDeviceFromIDCalls gets all the calls that were made to GetDeviceFromID.
// Check the length with:
//     len(mockedDatastore.GetDeviceFromIDCalls())
func (mock *DatastoreMock) GetDeviceFromIDCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockGetDeviceFromID.RLock()
	calls = mock.calls.GetDeviceFromID
	mock.lockGetDeviceFromID.RUnlock()
	return calls
}

// GetDeviceModelFromID calls GetDeviceModelFromIDFunc.
func (mock *DatastoreMock) GetDeviceModelFromID(id string) (*models.DeviceModel, error) {
	if mock.GetDeviceModelFromIDFunc == nil {
		panic("DatastoreMock.GetDeviceModelFromIDFunc: method is nil but Datastore.GetDeviceModelFromID was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	mock.lockGetDeviceModelFromID.Lock()
	mock.calls.GetDeviceModelFromID = append(mock.calls.GetDeviceModelFromID, callInfo)
	mock.lockGetDeviceModelFromID.Unlock()
	return mock.GetDeviceModelFromIDFunc(id)
}

// GetDeviceModelFromIDCalls gets all the calls that were made to GetDeviceModelFromID.
// Check the length with:
//     len(mockedDatastore.GetDeviceModelFromIDCalls())
func (mock *DatastoreMock) GetDeviceModelFromIDCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockGetDeviceModelFromID.RLock()
	calls = mock.calls.GetDeviceModelFromID
	mock.lockGetDeviceModelFromID.RUnlock()
	return calls
}

// GetDeviceModelFromPrimaryKey calls GetDeviceModelFromPrimaryKeyFunc.
func (mock *DatastoreMock) GetDeviceModelFromPrimaryKey(id uint) (*models.DeviceModel, error) {
	if mock.GetDeviceModelFromPrimaryKeyFunc == nil {
		panic("DatastoreMock.GetDeviceModelFromPrimaryKeyFunc: method is nil but Datastore.GetDeviceModelFromPrimaryKey was just called")
	}
	callInfo := struct {
		ID uint
	}{
		ID: id,
	}
	mock.lockGetDeviceModelFromPrimaryKey.Lock()
	mock.calls.GetDeviceModelFromPrimaryKey = append(mock.calls.GetDeviceModelFromPrimaryKey, callInfo)
	mock.lockGetDeviceModelFromPrimaryKey.Unlock()
	return mock.GetDeviceModelFromPrimaryKeyFunc(id)
}

// GetDeviceModelFromPrimaryKeyCalls gets all the calls that were made to GetDeviceModelFromPrimaryKey.
// Check the length with:
//     len(mockedDatastore.GetDeviceModelFromPrimaryKeyCalls())
func (mock *DatastoreMock) GetDeviceModelFromPrimaryKeyCalls() []struct {
	ID uint
} {
	var calls []struct {
		ID uint
	}
	mock.lockGetDeviceModelFromPrimaryKey.RLock()
	calls = mock.calls.GetDeviceModelFromPrimaryKey
	mock.lockGetDeviceModelFromPrimaryKey.RUnlock()
	return calls
}

// GetDeviceModels calls GetDeviceModelsFunc.
func (mock *DatastoreMock) GetDeviceModels() ([]models.DeviceModel, error) {
	if mock.GetDeviceModelsFunc == nil {
		panic("DatastoreMock.GetDeviceModelsFunc: method is nil but Datastore.GetDeviceModels was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetDeviceModels.Lock()
	mock.calls.GetDeviceModels = append(mock.calls.GetDeviceModels, callInfo)
	mock.lockGetDeviceModels.Unlock()
	return mock.GetDeviceModelsFunc()
}

// GetDeviceModelsCalls gets all the calls that were made to GetDeviceModels.
// Check the length with:
//     len(mockedDatastore.GetDeviceModelsCalls())
func (mock *DatastoreMock) GetDeviceModelsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetDeviceModels.RLock()
	calls = mock.calls.GetDeviceModels
	mock.lockGetDeviceModels.RUnlock()
	return calls
}

// GetDevices calls GetDevicesFunc.
func (mock *DatastoreMock) GetDevices() ([]models.Device, error) {
	if mock.GetDevicesFunc == nil {
		panic("DatastoreMock.GetDevicesFunc: method is nil but Datastore.GetDevices was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetDevices.Lock()
	mock.calls.GetDevices = append(mock.calls.GetDevices, callInfo)
	mock.lockGetDevices.Unlock()
	return mock.GetDevicesFunc()
}

// GetDevicesCalls gets all the calls that were made to GetDevices.
// Check the length with:
//     len(mockedDatastore.GetDevicesCalls())
func (mock *DatastoreMock) GetDevicesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetDevices.RLock()
	calls = mock.calls.GetDevices
	mock.lockGetDevices.RUnlock()
	return calls
}

// UpdateDeviceLocation calls UpdateDeviceLocationFunc.
func (mock *DatastoreMock) UpdateDeviceLocation(deviceID string, lat float64, lon float64) error {
	if mock.UpdateDeviceLocationFunc == nil {
		panic("DatastoreMock.UpdateDeviceLocationFunc: method is nil but Datastore.UpdateDeviceLocation was just called")
	}
	callInfo := struct {
		DeviceID string
		Lat      float64
		Lon      float64
	}{
		DeviceID: deviceID,
		Lat:      lat,
		Lon:      lon,
	}
	mock.lockUpdateDeviceLocation.Lock()
	mock.calls.UpdateDeviceLocation = append(mock.calls.UpdateDeviceLocation, callInfo)
	mock.lockUpdateDeviceLocation.Unlock()
	return mock.UpdateDeviceLocationFunc(deviceID, lat, lon)
}

// UpdateDeviceLocationCalls gets all the calls that were made to UpdateDeviceLocation.
// Check the length with:
//     len(mockedDatastore.UpdateDeviceLocationCalls())
func (mock *DatastoreMock) UpdateDeviceLocationCalls() []struct {
	DeviceID string
	Lat      float64
	Lon      float64
} {
	var calls []struct {
		DeviceID string
		Lat      float64
		Lon      float64
	}
	mock.lockUpdateDeviceLocation.RLock()
	calls = mock.calls.UpdateDeviceLocation
	mock.lockUpdateDeviceLocation.RUnlock()
	return calls
}

// UpdateDeviceState calls UpdateDeviceStateFunc.
func (mock *DatastoreMock) UpdateDeviceState(deviceID string, state string) error {
	if mock.UpdateDeviceStateFunc == nil {
		panic("DatastoreMock.UpdateDeviceStateFunc: method is nil but Datastore.UpdateDeviceState was just called")
	}
	callInfo := struct {
		DeviceID string
		State    string
	}{
		DeviceID: deviceID,
		State:    state,
	}
	mock.lockUpdateDeviceState.Lock()
	mock.calls.UpdateDeviceState = append(mock.calls.UpdateDeviceState, callInfo)
	mock.lockUpdateDeviceState.Unlock()
	return mock.UpdateDeviceStateFunc(deviceID, state)
}

// UpdateDeviceStateCalls gets all the calls that were made to UpdateDeviceState.
// Check the length with:
//     len(mockedDatastore.UpdateDeviceStateCalls())
func (mock *DatastoreMock) UpdateDeviceStateCalls() []struct {
	DeviceID string
	State    string
} {
	var calls []struct {
		DeviceID string
		State    string
	}
	mock.lockUpdateDeviceState.RLock()
	calls = mock.calls.UpdateDeviceState
	mock.lockUpdateDeviceState.RUnlock()
	return calls
}

// UpdateDeviceValue calls UpdateDeviceValueFunc.
func (mock *DatastoreMock) UpdateDeviceValue(deviceID string, value string) error {
	if mock.UpdateDeviceValueFunc == nil {
		panic("DatastoreMock.UpdateDeviceValueFunc: method is nil but Datastore.UpdateDeviceValue was just called")
	}
	callInfo := struct {
		DeviceID string
		Value    string
	}{
		DeviceID: deviceID,
		Value:    value,
	}
	mock.lockUpdateDeviceValue.Lock()
	mock.calls.UpdateDeviceValue = append(mock.calls.UpdateDeviceValue, callInfo)
	mock.lockUpdateDeviceValue.Unlock()
	return mock.UpdateDeviceValueFunc(deviceID, value)
}

// UpdateDeviceValueCalls gets all the calls that were made to UpdateDeviceValue.
// Check the length with:
//     len(mockedDatastore.UpdateDeviceValueCalls())
func (mock *DatastoreMock) UpdateDeviceValueCalls() []struct {
	DeviceID string
	Value    string
} {
	var calls []struct {
		DeviceID string
		Value    string
	}
	mock.lockUpdateDeviceValue.RLock()
	calls = mock.calls.UpdateDeviceValue
	mock.lockUpdateDeviceValue.RUnlock()
	return calls
}
