package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/max0nT/pr-assign/internal/controllers"
	"github.com/max0nT/pr-assign/internal/entities"
	"github.com/stretchr/testify/require"
)

const (
	host = "http://0.0.0.0:8080"
)

func MakeRequest(
	ctx context.Context,
	path string,
	method string,
	body io.Reader,
) (res *http.Response, err error) {
	req, err := http.NewRequestWithContext(ctx, method, host+path, body)

	if err != nil {
		return
	}

	res, err = http.DefaultClient.Do(req)
	return
}

var InactiveUser entities.User = entities.User{
	Id:       "uid5",
	Username: "username5",
	IsActive: false,
}

var ChangeUserActiveStatusData entities.UserChangeActive = entities.UserChangeActive{
	Id:       InactiveUser.Id,
	IsActive: !InactiveUser.IsActive,
}
var TeamAddData entities.TeamCreate = entities.TeamCreate{
	Name: "TeamName1",
	Users: []entities.User{
		{
			Id:       "uid1",
			Username: "username1",
			IsActive: true,
		},
		{
			Id:       "uid2",
			Username: "username2",
			IsActive: true,
		},
		{
			Id:       "uid3",
			Username: "username3",
			IsActive: true,
		},
		{
			Id:       "uid4",
			Username: "username4",
			IsActive: true,
		},
		InactiveUser,
	},
}
var PrOpenData entities.PrCreate = entities.PrCreate{
	PrId:      "pr1",
	PrName:    "The first Pr name",
	CreatedBy: "uid1",
}
var PrMergeData entities.PrMerge = entities.PrMerge{
	PrId: "pr1",
}

func TestTeamAdd(t *testing.T) {
	var body io.Reader
	rawBody, err := json.Marshal(TeamAddData)
	if err != nil {
		log.Fatalf(
			"Integration tests: Failed to parse team data before create: %s",
			err.Error(),
		)
		return
	}
	body = bytes.NewReader(rawBody)
	path := "/api/v1/team/add/"
	response, err := MakeRequest(
		context.Background(),
		path,
		http.MethodPost,
		body,
	)

	if err != nil {
		log.Fatalf(
			"Integration tests: Failed to make request to POST '/api/v1/team/add/`: %s",
			err.Error(),
		)
		return
	}

	require.Equal(t, response.StatusCode, http.StatusCreated)

	// Check There are team data
	getTeamData := "/api/v1/team/?team_name=" + TeamAddData.Name
	response, err = MakeRequest(
		context.Background(),
		getTeamData,
		http.MethodGet,
		body,
	)

	if err != nil {
		log.Fatalf(
			"Integration tests: Failed to make request to POST '%s`: %s",
			getTeamData,
			err.Error(),
		)
		return
	}

	require.Equal(t, response.StatusCode, http.StatusOK)

	var responseData entities.TeamRead
	if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	require.Equal(t, responseData.Name, TeamAddData.Name)
	require.Equal(t, len(responseData.Users), len(TeamAddData.Users))
}

func TestExistingTeamAdd(t *testing.T) {
	var body io.Reader
	rawBody, err := json.Marshal(TeamAddData)
	if err != nil {
		log.Fatalf(
			"Integration tests: Failed to parse team data before create: %s",
			err.Error(),
		)
		return
	}

	body = bytes.NewReader(rawBody)
	path := "/api/v1/team/add/"
	response, err := MakeRequest(
		context.Background(),
		path,
		http.MethodPost,
		body,
	)

	if err != nil {
		log.Fatalf(
			"Integration tests: Failed to make request to POST '/api/v1/team/add/`: %s",
			err.Error(),
		)
		return
	}
	require.Equal(t, response.StatusCode, http.StatusBadRequest)

	var ErrorData controllers.ErrorMessage
	if err := json.NewDecoder(response.Body).Decode(&ErrorData); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	require.Equal(
		t,
		ErrorData.Message,
		fmt.Sprintf(
			"Team with %s already exists",
			TeamAddData.Name,
		),
	)

}

func TestPrOpen(t *testing.T) {
	var body io.Reader
	rawBody, err := json.Marshal(PrOpenData)
	if err != nil {
		log.Fatalf(
			"Integration tests: Failed to parse team data before create: %s",
			err.Error(),
		)
		return
	}

	body = bytes.NewReader(rawBody)
	path := "/api/v1/pr/open/"
	response, err := MakeRequest(
		context.Background(),
		path,
		http.MethodPost,
		body,
	)
	if err != nil {
		log.Fatalf(
			"Integration tests: Failed to make request to POST '%s`: %s",
			path,
			err.Error(),
		)

	}

	require.Equal(t, response.StatusCode, http.StatusCreated)

	var responseData entities.PrRead
	if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	require.Equal(t, responseData.PrName, PrOpenData.PrName)
	require.Equal(t, len(responseData.Reviewers), 2)

	for _, reviewer := range responseData.Reviewers {
		require.Equal(t, reviewer.IsActive, true)
	}

}

func TestExistingPrOpen(t *testing.T) {
	var body io.Reader
	rawBody, err := json.Marshal(PrOpenData)
	if err != nil {
		log.Fatalf(
			"Integration tests: Failed to parse team data before create: %s",
			err.Error(),
		)
		return
	}

	body = bytes.NewReader(rawBody)
	path := "/api/v1/pr/open/"
	response, err := MakeRequest(
		context.Background(),
		path,
		http.MethodPost,
		body,
	)

	if err != nil {
		log.Fatalf(
			"Integration tests: Failed to make request to POST '%s`: %s",
			path,
			err.Error(),
		)
		return
	}
	require.Equal(t, response.StatusCode, http.StatusBadRequest)

	var ErrorData controllers.ErrorMessage
	if err := json.NewDecoder(response.Body).Decode(&ErrorData); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	require.Equal(
		t,
		ErrorData.Message,
		fmt.Sprintf(
			"Pr with id %s already exist",
			PrOpenData.PrId,
		),
	)

}

func TestMergePr(t *testing.T) {
	var body io.Reader
	rawBody, err := json.Marshal(PrOpenData)
	if err != nil {
		log.Fatalf(
			"Integration tests: Failed to parse team data before create: %s",
			err.Error(),
		)
		return
	}

	body = bytes.NewReader(rawBody)
	path := "/api/v1/pr/merge/"
	response, err := MakeRequest(
		context.Background(),
		path,
		http.MethodPatch,
		body,
	)
	if err != nil {
		log.Fatalf(
			"Integration tests: Failed to make request to PATCH '%s`: %s",
			path,
			err.Error(),
		)

	}

	require.Equal(t, response.StatusCode, http.StatusOK)

	var responseData entities.PrSimple
	if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	require.Equal(t, responseData.PrId, PrMergeData.PrId)
	require.Equal(t, responseData.IsMerged, true)

}

func TestChangeUserActive(t *testing.T) {
	var body io.Reader
	rawBody, err := json.Marshal(ChangeUserActiveStatusData)
	if err != nil {
		log.Fatalf(
			"Integration tests: Failed to parse team data before create: %s",
			err.Error(),
		)
		return
	}

	body = bytes.NewReader(rawBody)
	path := "/api/v1/user/change-status-active/"
	response, err := MakeRequest(
		context.Background(),
		path,
		http.MethodPatch,
		body,
	)
	if err != nil {
		log.Fatalf(
			"Integration tests: Failed to make request to PATCH '%s`: %s",
			path,
			err.Error(),
		)

	}

	require.Equal(t, response.StatusCode, http.StatusOK)

	var responseData entities.User
	if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	require.Equal(t, responseData.Id, InactiveUser.Id)
	require.Equal(t, responseData.IsActive, !InactiveUser.IsActive)
}
