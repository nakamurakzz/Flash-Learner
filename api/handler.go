package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"

	"github.com/labstack/echo/v4"
	"github.com/nakamurakzz/flash-leaner/api/oapi"
)

type Server struct {
}

func NewServer() Server {
	return Server{}
}

var _ oapi.ServerInterface = Server{}

// List all sentence groups
// (GET /sentence-groups)
func (s Server) GetSentenceGroups(ctx echo.Context, params oapi.GetSentenceGroupsParams) error {
	res := map[string]interface{}{
		"pagination": oapi.PaginationResponse{
			PageToken: uuid.New(),
		},
		"sentence_groups": []oapi.SentenceGroup{
			{
				Id:              uuid.New(),
				Title:           "Group 1",
				Description:     "Description 1",
				TotalReputation: 2,
				IsActive:        true,
				CreatedAt:       time.Now(),
				UpdatedAt:       time.Now(),
			},
			{
				Id:              uuid.New(),
				Title:           "Group 2",
				Description:     "Description 2",
				TotalReputation: 3,
				IsActive:        true,
				CreatedAt:       time.Now(),
				UpdatedAt:       time.Now(),
			},
		},
	}
	return ctx.JSON(http.StatusOK, res)
}

// Create a new sentence group
// (POST /sentence-groups)
func (s Server) PostSentenceGroups(ctx echo.Context) error {
	return ctx.JSON(http.StatusCreated, nil)
}

// List all sentences in a group
// (GET /sentence-groups/{groupId}/sentences)
func (s Server) GetSentenceGroupsGroupIdSentences(ctx echo.Context, groupId openapi_types.UUID, params oapi.GetSentenceGroupsGroupIdSentencesParams) error {
	return ctx.JSON(http.StatusOK, "ok")
}

// Create a new sentence in a group
// (POST /sentence-groups/{groupId}/sentences)
func (s Server) PostSentenceGroupsGroupIdSentences(ctx echo.Context, groupId openapi_types.UUID) error {
	return ctx.JSON(http.StatusCreated, nil)
}

// Delete a sentence by ID
// (DELETE /sentence-groups/{groupId}/sentences/{sentenceId})
func (s Server) DeleteSentenceGroupsGroupIdSentencesSentenceId(ctx echo.Context, groupId openapi_types.UUID, sentenceId openapi_types.UUID) error {
	return ctx.JSON(http.StatusNoContent, nil)
}

// Get a sentence by ID
// (GET /sentence-groups/{groupId}/sentences/{sentenceId})
func (s Server) GetSentenceGroupsGroupIdSentencesSentenceId(ctx echo.Context, groupId openapi_types.UUID, sentenceId openapi_types.UUID) error {
	return ctx.JSON(http.StatusOK, "ok")
}

// Update a sentence by ID
// (PUT /sentence-groups/{groupId}/sentences/{sentenceId})
func (s Server) PutSentenceGroupsGroupIdSentencesSentenceId(ctx echo.Context, groupId openapi_types.UUID, sentenceId openapi_types.UUID) error {
	return ctx.JSON(http.StatusOK, "ok")
}

// Increment the reputation of a sentence
// (POST /sentence-groups/{groupId}/sentences/{sentenceId}/reputation:increment)
func (s Server) PostSentenceGroupsGroupIdSentencesSentenceIdReputationIncrement(ctx echo.Context, groupId openapi_types.UUID, sentenceId openapi_types.UUID) error {
	return ctx.JSON(http.StatusCreated, "ok")
}

// Delete a sentence group by ID
// (DELETE /sentence-groups/{id})
func (s Server) DeleteSentenceGroupsId(ctx echo.Context, id openapi_types.UUID) error {
	return ctx.JSON(http.StatusNoContent, nil)
}

// Get a sentence group by ID
// (GET /sentence-groups/{id})
func (s Server) GetSentenceGroupsId(ctx echo.Context, id openapi_types.UUID) error {
	return ctx.JSON(http.StatusOK, "ok")
}

// Update a sentence group by ID
// (PUT /sentence-groups/{id})
func (s Server) PutSentenceGroupsId(ctx echo.Context, id openapi_types.UUID) error {
	return ctx.JSON(http.StatusOK, "ok")
}

// Generate sentences based on prompt and tone by LLM
// (POST /sentences:generate)
func (s Server) PostSentencesGenerate(ctx echo.Context) error {
	return ctx.JSON(http.StatusCreated, nil)
}
