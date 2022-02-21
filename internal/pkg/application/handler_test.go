package application

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/matryer/is"

	"github.com/diwise/iot-device-registry/internal/pkg/infrastructure/repositories/database"
	"github.com/diwise/iot-device-registry/internal/pkg/infrastructure/repositories/models"
	"github.com/diwise/messaging-golang/pkg/messaging"
	"github.com/diwise/ngsi-ld-golang/pkg/datamodels/fiware"
	ngsi "github.com/diwise/ngsi-ld-golang/pkg/ngsi-ld"
	ngsitypes "github.com/diwise/ngsi-ld-golang/pkg/ngsi-ld/types"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestThatCreateEntityDoesNotAcceptUnknownBody(t *testing.T) {
	is := is.New(t)

	bodyContents := []byte("{\"json\":\"json\"}")
	req, _ := http.NewRequest("POST", createURL("/ngsi-ld/v1/entities"), bytes.NewBuffer(bodyContents))
	w := httptest.NewRecorder()

	ctxreg := createContextRegistry(log.Logger, nil, nil)
	ngsi.NewCreateEntityHandler(ctxreg).ServeHTTP(w, req)

	is.Equal(w.Code, http.StatusBadRequest) // create entity should return bad request
}

func TestThatCreateEntityStoresCorrectDevice(t *testing.T) {
	is := is.New(t)

	deviceID := fiware.DeviceIDPrefix + "deviceID"
	device := fiware.NewDevice(deviceID, "")
	device.RefDeviceModel, _ = fiware.NewDeviceModelRelationship(
		fiware.DeviceModelIDPrefix + "livboj",
	)
	jsonBytes, _ := json.Marshal(device)

	req, _ := http.NewRequest("POST", createURL("/ngsi-ld/v1/entities"), bytes.NewBuffer(jsonBytes))
	w := httptest.NewRecorder()

	db := &database.DatastoreMock{
		CreateDeviceFunc: func(device *fiware.Device) (*models.Device, error) {
			return &models.Device{}, nil
		},
	}

	ctxreg := createContextRegistry(log.Logger, nil, db)
	ngsi.NewCreateEntityHandler(ctxreg).ServeHTTP(w, req)

	is.Equal(len(db.CreateDeviceCalls()), 1)                // create count should be 1
	is.Equal(db.CreateDeviceCalls()[0].Device.ID, deviceID) // device id must match
}

func TestThatCreateEntityStoresCorrectDeviceModel(t *testing.T) {
	is := is.New(t)

	categories := []string{"sensor"}
	deviceModel := fiware.NewDeviceModel("badtemperatur", categories)
	deviceModel.ControlledProperty = ngsitypes.NewTextListProperty([]string{"temperature"})

	jsonBytes, _ := json.Marshal(deviceModel)

	req, _ := http.NewRequest("POST", createURL("/ngsi-ld/v1/entities"), bytes.NewBuffer(jsonBytes))
	w := httptest.NewRecorder()

	db := &database.DatastoreMock{
		CreateDeviceModelFunc: func(device *fiware.DeviceModel) (*models.DeviceModel, error) {
			return &models.DeviceModel{}, nil
		},
	}
	ctxreg := createContextRegistry(log.Logger, nil, db)
	ngsi.NewCreateEntityHandler(ctxreg).ServeHTTP(w, req)

	is.Equal(len(db.CreateDeviceModelCalls()), 1) // create count should be 1
}

func TestThatCreateEntityFailsOnUnknownEntity(t *testing.T) {
	is := is.New(t)

	db := &database.DatastoreMock{
		CreateDeviceModelFunc: func(*fiware.DeviceModel) (*models.DeviceModel, error) {
			return nil, errors.New("create should fail")
		},
	}

	categories := []string{"sensor"}
	deviceModel := fiware.NewDeviceModel("badtemperatur", categories)
	deviceModel.ControlledProperty = ngsitypes.NewTextListProperty([]string{"temperature"})

	jsonBytes, _ := json.Marshal(deviceModel)

	req, _ := http.NewRequest("POST", createURL("/ngsi-ld/v1/entities"), bytes.NewBuffer(jsonBytes))
	w := httptest.NewRecorder()

	ctxreg := createContextRegistry(log.Logger, nil, db)
	ngsi.NewCreateEntityHandler(ctxreg).ServeHTTP(w, req)

	is.Equal(w.Code, http.StatusBadRequest) // create entity should return bad request
}

func TestThatPatchWaterTempDevicePublishesOnTheMessageQueue(t *testing.T) {
	is := is.New(t)

	db := &database.DatastoreMock{
		GetDeviceFromIDFunc: func(string) (*models.Device, error) {
			return &models.Device{Latitude: 64, Longitude: 17}, nil
		},
		UpdateDeviceValueFunc: func(deviceID string, value string) error {
			return nil
		},
	}
	m := messaging.ContextMock{}

	jsonBytes, _ := json.Marshal(createDevicePatchWithValue("sk-elt-temp-02", "t%3D8"))

	req, _ := http.NewRequest("PATCH", createURL("/ngsi-ld/v1/entities/urn:ngsi-ld:Device:sk-elt-temp-02/attrs/"), bytes.NewBuffer(jsonBytes))
	w := httptest.NewRecorder()

	ctxreg := createContextRegistry(log.Logger, &m, db)
	ngsi.NewUpdateEntityAttributesHandler(ctxreg).ServeHTTP(w, req)

	is.Equal(len(m.SendCommandToCalls()), 1) // expected 1 command to have been sent
}

func TestThatPatchAmbientTempDevicePublishesOnTheMessageQueue(t *testing.T) {
	is := is.New(t)

	db := &database.DatastoreMock{
		GetDeviceFromIDFunc: func(string) (*models.Device, error) {
			return &models.Device{Latitude: 64, Longitude: 17}, nil
		},
		UpdateDeviceValueFunc: func(deviceID string, value string) error {
			return nil
		},
	}
	m := messaging.ContextMock{}

	jsonBytes, _ := json.Marshal(createDevicePatchWithValue("se:trafikverket:temp:SE_STA_VVIS2115", "t%3D12"))

	req, _ := http.NewRequest("PATCH", createURL("/ngsi-ld/v1/entities/urn:ngsi-ld:Device:se:trafikverket:temp:SE_STA_VVIS2115/attrs/"), bytes.NewBuffer(jsonBytes))
	w := httptest.NewRecorder()

	ctxreg := createContextRegistry(log.Logger, &m, db)
	ngsi.NewUpdateEntityAttributesHandler(ctxreg).ServeHTTP(w, req)

	is.Equal(len(m.SendCommandToCalls()), 1) // expected 1 command to have been sent
}

func TestThatPatchDeviceWithNoTempDoesNotPublishOnTheMessageQueue(t *testing.T) {
	is := is.New(t)

	db := &database.DatastoreMock{
		GetDeviceFromIDFunc: func(string) (*models.Device, error) {
			return &models.Device{Latitude: 64, Longitude: 17}, nil
		},
		UpdateDeviceValueFunc: func(deviceID string, value string) error {
			return nil
		},
	}
	m := messaging.ContextMock{}

	jsonBytes, _ := json.Marshal(createDevicePatchWithValue("se:trafikverket:temp:SE_STA_VVIS2115", "co2%3D12"))

	req, _ := http.NewRequest("PATCH", createURL("/ngsi-ld/v1/entities/urn:ngsi-ld:Device:se:trafikverket:temp:SE_STA_VVIS2115/attrs/"), bytes.NewBuffer(jsonBytes))
	w := httptest.NewRecorder()

	ctxreg := createContextRegistry(log.Logger, &m, db)
	ngsi.NewUpdateEntityAttributesHandler(ctxreg).ServeHTTP(w, req)

	is.Equal(len(m.SendCommandToCalls()), 0) // expected 0 commands to have been sent
}

func TestThatDeviceStateCanBeUpdated(t *testing.T) {
	is := is.New(t)

	db := &database.DatastoreMock{
		GetDeviceFromIDFunc: func(string) (*models.Device, error) {
			return &models.Device{Latitude: 64, Longitude: 17}, nil
		},
		UpdateDeviceStateFunc: func(deviceID, state string) error {
			return nil
		},
		UpdateDeviceValueFunc: func(deviceID string, value string) error {
			return nil
		},
	}
	m := messaging.ContextMock{}

	device := fiware.NewDevice("deviceId", "")
	device.DeviceState = ngsitypes.NewTextProperty("on")

	jsonBytes, _ := json.Marshal(device)

	req, _ := http.NewRequest("PATCH", createURL("/ngsi-ld/v1/entities/urn:ngsi-ld:Device:se:deviceId/attrs/"), bytes.NewBuffer(jsonBytes))
	w := httptest.NewRecorder()

	ctxreg := createContextRegistry(log.Logger, &m, db)
	ngsi.NewUpdateEntityAttributesHandler(ctxreg).ServeHTTP(w, req)

	is.Equal(db.UpdateDeviceValueCalls()[0].Value, "")

	is.Equal(len(db.UpdateDeviceStateCalls()), 1) // expected update device state to be called once
}

func TestRetrieveEntity(t *testing.T) {
	is := is.New(t)

	db := &database.DatastoreMock{
		GetDeviceModelFromIDFunc: func(string) (*models.DeviceModel, error) {
			return &models.DeviceModel{}, nil
		},
		GetDeviceFromIDFunc: func(string) (*models.Device, error) {
			return &models.Device{}, nil
		},
	}

	req, _ := http.NewRequest("GET", createURL("/ngsi-ld/v1/entities/urn:ngsi-ld:DeviceModel:sk-elt-temp-02"), nil)
	w := httptest.NewRecorder()

	ctxreg := createContextRegistry(log.Logger, nil, db)

	ngsi.NewRetrieveEntityHandler(ctxreg).ServeHTTP(w, req)

	is.Equal(w.Code, http.StatusOK) // request failed
}

func TestThatRetrieveEntityDoesNotReturnZeroLocations(t *testing.T) {
	is := is.New(t)

	db := &database.DatastoreMock{
		GetDeviceModelFromPrimaryKeyFunc: func(uint) (*models.DeviceModel, error) {
			return &models.DeviceModel{}, nil
		},
		GetDeviceFromIDFunc: func(string) (*models.Device, error) {
			return &models.Device{
				DeviceID: "urn:ngsi-ld:Device:sk-elt-temp-02",
				Value:    "t=12",
			}, nil
		},
	}

	req, _ := http.NewRequest("GET", createURL("/ngsi-ld/v1/entities/urn:ngsi-ld:Device:sk-elt-temp-02"), nil)
	w := httptest.NewRecorder()

	logger := log.Logger

	ctxreg := createContextRegistry(logger, nil, db)

	ngsi.NewRetrieveEntityHandler(ctxreg).ServeHTTP(w, req)

	is.Equal(w.Code, http.StatusOK) // request failed

	responseBytes, err := ioutil.ReadAll(w.Body)
	is.NoErr(err) // failed to read response body

	deviceStr, err := getDeviceAsString(responseBytes, logger)
	is.NoErr(err) // failed to get device as string

	is.Equal(deviceStr, zeroLocationDevice) // retrieve entity returned a zero location
}

const zeroLocationDevice string = `{"id":"urn:ngsi-ld:Device:sk-elt-temp-02","type":"Device","@context":["https://schema.lab.fiware.org/ld/context","https://uri.etsi.org/ngsi-ld/v1/ngsi-ld-core-context.jsonld"],"value":{"type":"Property","value":"t%3D12"},"refDeviceModel":{"type":"Relationship","object":"urn:ngsi-ld:DeviceModel:"}}`

// write unit test for retrieve entity where device is nil.

func getDeviceAsString(response []byte, log zerolog.Logger) (string, error) {
	device := fiware.Device{}

	err := json.Unmarshal(response, &device)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal device json")
		return "", err
	}

	deviceStr, err := json.Marshal(device)
	if err != nil {
		log.Error().Err(err).Msg("could not marshal device to json string")
		return "", err
	}

	return string(deviceStr), nil
}

func createDevicePatchWithValue(deviceid, value string) *fiware.Device {
	device := fiware.NewDevice(deviceid, value)
	return device
}

func createURL(path string, params ...string) string {
	url := "http://localhost:8080/ngsi-ld/v1" + path

	if len(params) > 0 {
		url = url + "?"

		for _, p := range params {
			url = url + p + "&"
		}

		url = strings.TrimSuffix(url, "&")
	}

	return url
}
