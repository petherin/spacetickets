package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/petherin/spacetickets/internal/domains/bookings"
)

func TestServer_Get(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
		want           string
		wantErr        error
		wantStatusCode int
	}{
		{
			name:           "1. Successfully returns all bookings",
			req:            httptest.NewRequest(http.MethodGet, "/api/v1/bookings", nil),
			want:           `[{"id":"uuid-1","first_name":"Ian","last_name":"Thomson","gender":"Male","birthday":"2000-01-02T00:00:00Z","launch_pad_id":"4079f070-3e58-4e61-8af7-05c8de8e1fbf","destination_id":"fbd40165-03c7-47a5-be72-c79f81ebbf67","launch_date":"2011-01-02T00:00:00Z","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]`,
			wantErr:        nil,
			wantStatusCode: 200,
		},
		{
			name:           "2. Errors when getting all bookings, returns 500 and error message",
			req:            httptest.NewRequest(http.MethodGet, "/api/v1/bookings", nil),
			want:           ``,
			wantErr:        fmt.Errorf("oops"),
			wantStatusCode: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handlers := NewBookingHandlers(bookerMock{ForceError: tt.wantErr}, nil, "")
			mux := &http.ServeMux{}
			mux.HandleFunc("GET /api/v1/bookings", handlers.Get)

			w := httptest.NewRecorder()

			mux.ServeHTTP(w, tt.req)

			res := w.Result()
			if status := res.StatusCode; status != tt.wantStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantStatusCode)
			}

			body := w.Body.String()
			if !strings.Contains(body, tt.want) {
				t.Errorf("handler returned unexpected body: got %v want %v", body, tt.want)
			}
		})
	}
}

func TestServer_Post(t *testing.T) {
	tests := []struct {
		name           string
		req            *http.Request
		client         *http.Client
		want           string
		wantStatusCode int
	}{
		{
			name: "1. Successfully creates a booking",
			req: httptest.NewRequest(http.MethodPost, "/api/v1/booking", strings.NewReader(`{
  "first_name": "Ian",
  "last_name": "Thomson",
  "gender": "Male",
  "birthday": "2000-04-12",
  "launch_pad_id": "4079f070-3e58-4e61-8af7-05c8de8e1fbf",
  "destination_id": "fbd40165-03c7-47a5-be72-c79f81ebbf67",
  "launch_date": "2022-10-05"
}`)),
			client: NewTestClient(func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(bytes.NewBufferString(`{
					"totalDocs": 0
}`)),
					Header: make(http.Header),
				}
			}),
			want:           `{"id":"uuid-1","first_name":"Ian","last_name":"Thomson","gender":"Male","birthday":"2000-01-02T00:00:00Z","launch_pad_id":"uuid-2","destination_id":"uuid-3","launch_date":"2011-01-02T00:00:00Z","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`,
			wantStatusCode: 200,
		},
		{
			name: "2. Booking cancelled, clash with SpaceX launch creates a booking",
			req: httptest.NewRequest(http.MethodPost, "/api/v1/booking", strings.NewReader(`{
  "first_name": "Ian",
  "last_name": "Thomson",
  "gender": "Male",
  "birthday": "2000-04-12",
  "launch_pad_id": "4079f070-3e58-4e61-8af7-05c8de8e1fbf",
  "destination_id": "fbd40165-03c7-47a5-be72-c79f81ebbf67",
  "launch_date": "2022-10-05"
}`)),
			client: NewTestClient(func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(bytes.NewBufferString(`{
					"totalDocs": 1
}`)),
					Header: make(http.Header),
				}
			}),
			want:           `{"Status": "Flight cancelled, overlaps with SpaceX launch"}`,
			wantStatusCode: 200,
		},
		{
			name: "3. Booking cancelled, flight to destination not running from launchpad today",
			req: httptest.NewRequest(http.MethodPost, "/api/v1/booking", strings.NewReader(`{
  "first_name": "Ian",
  "last_name": "Thomson",
  "gender": "Male",
  "birthday": "2000-04-12",
  "launch_pad_id": "false",
  "destination_id": "fbd40165-03c7-47a5-be72-c79f81ebbf67",
  "launch_date": "2022-10-05"
}`)),
			client: NewTestClient(func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(bytes.NewBufferString(`{
					"totalDocs": 0
}`)),
					Header: make(http.Header),
				}
			}),
			want:           `{"Status": "Flight cancelled, this launchpad does not fly to the destination on the requested day"}`,
			wantStatusCode: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handlers := NewBookingHandlers(bookerMock{}, tt.client, "")
			mux := &http.ServeMux{}
			mux.HandleFunc("POST /api/v1/booking", handlers.Post)

			w := httptest.NewRecorder()

			mux.ServeHTTP(w, tt.req)

			res := w.Result()
			if status := res.StatusCode; status != tt.wantStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantStatusCode)
			}

			body := w.Body.String()
			if !strings.Contains(body, tt.want) {
				t.Errorf("handler returned unexpected body: got %v want %v", body, tt.want)
			}
		})
	}
}

func TestServer_Delete(t *testing.T) {
	tests := []struct {
		name           string
		req            *http.Request
		want           string
		wantStatusCode int
	}{
		{
			name:           "1. Successfully deletes a booking",
			req:            httptest.NewRequest(http.MethodDelete, "/api/v1/booking/uuid-1", nil),
			want:           `{"Status": "Record deleted"}`,
			wantStatusCode: 200,
		},
		{
			name:           "2. ID not found in database, error returned",
			req:            httptest.NewRequest(http.MethodDelete, "/api/v1/booking/zero", nil),
			want:           `{"Status": "ID not recognised"}`,
			wantStatusCode: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handlers := NewBookingHandlers(bookerMock{}, nil, "")
			mux := &http.ServeMux{}
			mux.HandleFunc("DELETE /api/v1/booking/{id}", handlers.Delete)

			w := httptest.NewRecorder()

			mux.ServeHTTP(w, tt.req)

			res := w.Result()
			if status := res.StatusCode; status != tt.wantStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantStatusCode)
			}

			body := w.Body.String()
			if !strings.Contains(body, tt.want) {
				t.Errorf("handler returned unexpected body: got %v want %v", body, tt.want)
			}
		})
	}
}

type bookerMock struct {
	ForceError error
}

func (b bookerMock) GetAll() ([]bookings.Booking, error) {
	if b.ForceError != nil {
		return nil, b.ForceError
	}

	birthday, _ := time.Parse(time.DateOnly, "2000-01-02")
	launchDate, _ := time.Parse(time.DateOnly, "2011-01-02")

	return []bookings.Booking{
		{
			Id: "uuid-1",
			Customer: bookings.Customer{
				FirstName: "Ian",
				LastName:  "Thomson",
				Birthday:  birthday,
				Gender:    "Male",
			},
			LaunchPadId:   "4079f070-3e58-4e61-8af7-05c8de8e1fbf",
			DestinationId: "fbd40165-03c7-47a5-be72-c79f81ebbf67",
			LaunchDate:    launchDate,
		},
	}, nil
}

func (b bookerMock) Create(booking bookings.Booking) (*bookings.Booking, error) {
	birthday, _ := time.Parse(time.DateOnly, "2000-01-02")
	launchDate, _ := time.Parse(time.DateOnly, "2011-01-02")

	return &bookings.Booking{
		Id: "uuid-1",
		Customer: bookings.Customer{
			FirstName: "Ian",
			LastName:  "Thomson",
			Birthday:  birthday,
			Gender:    "Male",
		},
		LaunchPadId:   "uuid-2",
		DestinationId: "uuid-3",
		LaunchDate:    launchDate,
	}, nil
}

func (b bookerMock) Delete(bookingId string) (int64, error) {
	if bookingId == "zero" {
		return 0, nil
	}
	return 1, nil
}

func (b bookerMock) GetLaunchPad(id string) (*bookings.LaunchPad, error) {
	return &bookings.LaunchPad{
		Id:                "uuid-1",
		FullName:          "Cape Canaveral",
		SpaceXLaunchPadId: "123",
	}, nil
}

func (b bookerMock) IsLaunchScheduleValid(launchPadId, dayOfWeek, destinationId string) (bool, error) {
	if launchPadId == "false" {
		return false, nil
	}
	return true, nil
}

// RoundTripFunc defines a function for a RoundTrip
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip is a custom RoundTrip assigned to an http.client used in tests so we can control the response
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}
