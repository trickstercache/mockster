/*
 * Copyright 2018 Comcast Cable Communications Management, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package prometheus

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQueryRangeHandler(t *testing.T) {

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://0/query_range?query=up&start=0&end=30&step=15", nil)
	queryRangeHandler(w, r)

	resp := w.Result()

	// it should return 200 OK
	if resp.StatusCode != 200 {
		t.Errorf("expected 200 got %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	const expected = `{"status":"success","data":{"resultType":"matrix","result":[{"metric":{"series_id":"0"},"values":[[0,"29"],[15,"81"],[30,"23"]]}]}}`

	if string(bodyBytes) != expected {
		t.Errorf("expected %s got %s", expected, bodyBytes)
	}

	// Test with a duration that includes a unit of measurement
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "http://0/query_range?query=up&start=0&end=30&step=15s", nil)
	queryRangeHandler(w, r)

	resp = w.Result()

	// it should return 200 OK
	if resp.StatusCode != 200 {
		t.Errorf("expected 200 got %d", resp.StatusCode)
	}

	bodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if string(bodyBytes) != expected {
		t.Errorf("expected %s got %s", expected, bodyBytes)
	}

}

func TestQueryRangeHandlerFloatTime(t *testing.T) {

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://0/query_range?query=up&start=0.000&end=30.456&step=15", nil)
	queryRangeHandler(w, r)

	resp := w.Result()

	// it should return 200 OK
	if resp.StatusCode != 200 {
		t.Errorf("expected 200 got %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	const expected = `{"status":"success","data":{"resultType":"matrix","result":[{"metric":{"series_id":"0"},"values":[[0,"29"],[15,"81"],[30,"23"]]}]}}`

	if string(bodyBytes) != expected {
		t.Errorf("expected %s got %s", expected, bodyBytes)
	}
}

func TestQueryRangeHandlerMissingParam(t *testing.T) {

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://0/query_range?q=up&start=0&end=30&step=15", nil)
	queryRangeHandler(w, r)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected %d got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestQueryRangeHandlerInvalidParam(t *testing.T) {

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://0/query_range?query=up&start=foo&end=30&step=15", nil)
	queryRangeHandler(w, r)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected %d got %d", http.StatusBadRequest, resp.StatusCode)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "http://0/query_range?query=up&start=0&end=foo&step=15", nil)
	queryRangeHandler(w, r)

	resp = w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected %d got %d", http.StatusBadRequest, resp.StatusCode)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "http://0/query_range?query=up&start=0&end=30&step=foo", nil)
	queryRangeHandler(w, r)

	resp = w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected %d got %d", http.StatusBadRequest, resp.StatusCode)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "http://0/query_range?query=up{status_code=400}&start=0&end=30&step=15", nil)
	queryRangeHandler(w, r)

	resp = w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected %d got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestQueryHandler(t *testing.T) {

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://0/query?query=up&time=0", nil)
	queryHandler(w, r)

	resp := w.Result()

	// it should return 200 OK
	if resp.StatusCode != 200 {
		t.Errorf("expected 200 got %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	const expected = `{"status":"success","data":{"resultType":"vector","result":[{"metric":{"series_id":"0"},"value":[0,"29"]}]}}`

	if string(bodyBytes) != expected {
		t.Errorf("expected %s got %s", expected, bodyBytes)
	}
}

func TestQueryHandlerFloatTime(t *testing.T) {

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://0/query?query=up&time=30.456", nil)
	queryHandler(w, r)

	resp := w.Result()

	// it should return 200 OK
	if resp.StatusCode != 200 {
		t.Errorf("expected 200 got %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	const expected = `{"status":"success","data":{"resultType":"vector","result":[{"metric":{"series_id":"0"},"value":[30,"23"]}]}}`

	if string(bodyBytes) != expected {
		t.Errorf("expected %s got %s", expected, bodyBytes)
	}
}

func TestQueryHandlerMissingParam(t *testing.T) {

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://0/query?q=up", nil)
	queryHandler(w, r)
	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected %d got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestQueryHandlerInvalidParam(t *testing.T) {

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://0/query?query=up&time=foo", nil)
	queryHandler(w, r)
	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected %d got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestParseTime(t *testing.T) {

	const time1 = "2006-01-02T15:04:05.999999999Z"
	_, err := parseTime(time1)
	if err != nil {
		t.Error(err)
	}

}

func TestParseDuration(t *testing.T) {

	// Test inferred seconds
	d, err := parseDuration("15")
	if err != nil {
		t.Error(err)
	}
	if d != 15 {
		t.Errorf("expected %d got %d", 15, d)
	}

	// Test unit of h
	d, err = parseDuration("1h")
	if err != nil {
		t.Error(err)
	}
	if d != 3600 {
		t.Errorf("expected %d got %d", 3600, d)
	}

	// Test invalid unit
	d, err = parseDuration("1x")
	if err == nil {
		t.Errorf("expected parseDuration error for input [%s] got [%d]", "1x", d)
	}

	// Test decimal
	d, err = parseDuration("1.3")
	if err == nil {
		t.Errorf("expected parseDuration error for input [%s] got [%d]", "1.3", d)
	}

	// Test empty
	d, err = parseDuration("")
	if err == nil {
		t.Errorf("expected parseDuration error for input [%s] got [%d]", "", d)
	}

	// Test Invalid
	d, err = parseDuration("1s1t")
	if err == nil {
		t.Errorf("expected parseDuration error for input [%s] got [%d]", "", d)
	}

	// Test Valid Units but No Value
	d, err = parseDuration("s")
	if err == nil {
		t.Errorf("expected parseDuration error for input [%s] got [%d]", "", d)
	}

	// Test Negative
	_, err = parseDuration("-1")
	if err != nil {
		t.Error(err)
	}

}
