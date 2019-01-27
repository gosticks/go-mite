package mite

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// -------------------------------------------------------------
// ~ Private Methods
// -------------------------------------------------------------

func (m *Mite) getFromMite(suffix string, params map[string]string) (*http.Response, error) {

	client := &http.Client{}

	url := m.MitePathWithParam(suffix)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()

	// Add params
	for k, v := range params {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()

	m.l.Debug("Out -> GET: " + req.URL.String())
	req.Header.Add("User-Agent", m.AppName+"; mite-go/v0.5")
	req.Header.Add("X-MiteApiKey", m.ApiKey)
	// Set auth key header
	// req.Header.Set("X-MiteApiKey", m.ApiKey)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *Mite) postToMite(suffix string, data interface{}) (*http.Response, error) {
	return m.sendToMite(suffix, data, http.MethodPost)
}

func (m *Mite) patchAtMite(suffix string, data interface{}) (*http.Response, error) {
	return m.sendToMite(suffix, data, http.MethodPatch)
}

func (m *Mite) deleteFromMite(suffix string, data interface{}) (*http.Response, error) {
	return m.sendToMite(suffix, data, http.MethodDelete)
}

func (m *Mite) sendToMite(suffix string, data interface{}, method string) (*http.Response, error) {

	client := &http.Client{}

	url := m.MitePathWithParam(suffix)
	b, errMarshal := json.Marshal(data)
	if errMarshal != nil {
		return nil, errMarshal
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()

	// q.Add("api_key", m.ApiKey)

	req.URL.RawQuery = q.Encode()

	// m.l.Debug("Out -> POST: " + req.URL.String())

	// Set auth key header
	req.Header.Add("User-Agent", m.AppName+"; mite-go/v0.5")
	req.Header.Add("X-MiteApiKey", m.ApiKey)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *Mite) getAndDecodeFromSuffix(suffix string, target interface{}, params map[string]string) error {
	resp, errResp := m.getFromMite(suffix, params)
	if errResp != nil {
		return errResp
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Mite responded with error" + resp.Status + fmt.Sprint(resp.StatusCode))
	}

	// Unmarshal data
	err := json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		m.l.Error("Failed to decode", err)
		return err
	}
	return nil
}
