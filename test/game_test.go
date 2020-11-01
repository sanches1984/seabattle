package test

import (
	"bytes"
	"encoding/json"
	"github.com/sanches1984/seabattle/internal/app/api"
	"github.com/sanches1984/seabattle/internal/app/model"
	"github.com/stretchr/testify/require"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGame(t *testing.T) {
	t.Run("Create matrix", func(t *testing.T) {
		w := httptest.NewRecorder()
		api.CreateMatrix(w, prepareRequest("POST", `{"range": 6}`))
		require.Equal(t, w.Code, 200)
	})

	t.Run("Create ships", func(t *testing.T) {
		w := httptest.NewRecorder()
		api.Ship(w, prepareRequest("POST", `{"Coordinates": "2A 2A,4D 5D"}`))
		require.Equal(t, w.Code, 200)
	})

	t.Run("Make shot - miss", func(t *testing.T) {
		w := httptest.NewRecorder()
		api.Shot(w, prepareRequest("POST", `{"coord": "1B"}`))
		require.Equal(t, w.Code, 200)

		resp, err := prepareShotResponse(w.Body)
		require.Nil(t, err)
		require.Equal(t, resp, &model.ShotResponse{
			Destroy: false,
			Knock:   false,
			End:     false,
		})
	})

	t.Run("Make shot - error", func(t *testing.T) {
		w := httptest.NewRecorder()
		api.Shot(w, prepareRequest("POST", `{"coord": "1B"}`))
		require.Equal(t, w.Code, 400)
	})

	t.Run("Make shot - knock", func(t *testing.T) {
		w := httptest.NewRecorder()
		api.Shot(w, prepareRequest("POST", `{"coord": "4D"}`))
		require.Equal(t, w.Code, 200)

		resp, err := prepareShotResponse(w.Body)
		require.Nil(t, err)
		require.Equal(t, resp, &model.ShotResponse{
			Destroy: false,
			Knock:   true,
			End:     false,
		})
	})

	t.Run("Make shot - destroy", func(t *testing.T) {
		w := httptest.NewRecorder()
		api.Shot(w, prepareRequest("POST", `{"coord": "2A"}`))
		require.Equal(t, w.Code, 200)

		resp, err := prepareShotResponse(w.Body)
		require.Nil(t, err)
		require.Equal(t, resp, &model.ShotResponse{
			Destroy: true,
			Knock:   true,
			End:     false,
		})
	})

	t.Run("Get state", func(t *testing.T) {
		w := httptest.NewRecorder()
		api.State(w, prepareRequest("GET", ""))
		require.Equal(t, w.Code, 200)

		resp, err := prepareStateResponse(w.Body)
		require.Nil(t, err)
		require.Equal(t, resp, &model.StateResponse{
			ShipCount: 2,
			Destroyed: 1,
			Knocked:   1,
			ShotCount: 3,
		})
	})

	t.Run("Make shot - destroy (game ended)", func(t *testing.T) {
		w := httptest.NewRecorder()
		api.Shot(w, prepareRequest("POST", `{"coord": "5D"}`))
		require.Equal(t, w.Code, 200)
		resp, err := prepareShotResponse(w.Body)
		require.Nil(t, err)
		require.Equal(t, resp, &model.ShotResponse{
			Destroy: true,
			Knock:   true,
			End:     true,
		})
	})

	t.Run("Make shot - game ended", func(t *testing.T) {
		w := httptest.NewRecorder()
		api.Shot(w, prepareRequest("POST", `{"coord": "3B"}`))
		require.Equal(t, w.Code, 400)
	})

	t.Run("Clear game", func(t *testing.T) {
		w := httptest.NewRecorder()
		api.Clear(w, prepareRequest("POST", ""))
		require.Equal(t, w.Code, 200)
	})

	t.Run("Get state - cleared", func(t *testing.T) {
		w := httptest.NewRecorder()
		api.State(w, prepareRequest("GET", ""))
		require.Equal(t, w.Code, 200)

		resp, err := prepareStateResponse(w.Body)
		require.Nil(t, err)
		require.Equal(t, resp, &model.StateResponse{
			ShipCount: 2,
			Destroyed: 0,
			Knocked:   0,
			ShotCount: 0,
		})
	})
}

func prepareRequest(method, request string) *http.Request {
	reader := bytes.NewReader([]byte(request))
	r, err := http.NewRequest(method, "", reader)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func prepareShotResponse(b *bytes.Buffer) (*model.ShotResponse, error) {
	var resp model.ShotResponse
	err := json.Unmarshal(b.Bytes(), &resp)
	return &resp, err
}

func prepareStateResponse(b *bytes.Buffer) (*model.StateResponse, error) {
	var resp model.StateResponse
	err := json.Unmarshal(b.Bytes(), &resp)
	return &resp, err
}
