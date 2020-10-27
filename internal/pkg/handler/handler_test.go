package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/iot-for-tillgenglighet/messaging-golang/pkg/messaging"
	"github.com/iot-for-tillgenglighet/ngsi-ld-golang/pkg/datamodels/fiware"
	ngsi "github.com/iot-for-tillgenglighet/ngsi-ld-golang/pkg/ngsi-ld"
)

func TestMain(m *testing.M) {
	log.SetFormatter(&log.JSONFormatter{})
	os.Exit(m.Run())
}

func TestThatPatchWaterTempDevicePublishesOnTheMessageQueue(t *testing.T) {
	m := msgMock{}

	jsonBytes, _ := json.Marshal(createDevicePatchWithValue("sk-elt-temp-02", "t%3D12"))
	req, _ := http.NewRequest("PATCH", createURL("/ngsi-ld/v1/entities/urn:ngsi-ld:Device:sk-elt-temp-02/attrs/"), bytes.NewBuffer(jsonBytes))
	w := httptest.NewRecorder()

	ctxreg := createContextRegistry(&m, nil)
	ngsi.NewUpdateEntityAttributesHandler(ctxreg).ServeHTTP(w, req)

	if m.PublishCount != 1 {
		t.Error("Wrong publish count: ", m.PublishCount, "!=", 1)
	}
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

type msgMock struct {
	PublishCount uint32
}

func (m *msgMock) PublishOnTopic(message messaging.TopicMessage) error {
	m.PublishCount++
	return nil
}
