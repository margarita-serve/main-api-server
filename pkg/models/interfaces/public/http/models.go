package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddRoutes(e *echo.Echo) {
	e.POST("/models", Post)
}

// @Summary Create model
// @Description Create model's info
// @Accept json
// @Produce json
// @Param modelBody body PostModelsRequest true "name of the model"
// @Success 200 {object} PostModelsResponse
// @Router /project/{projectID}/models [post]
func Post(c echo.Context) error {

	//...

	post_response := new(PostModelsResponse)
	post_response.ID = "asdf"
	return c.JSONPretty(http.StatusOK, *post_response, "  ")
}

type PostModelsRequest struct {
	Name                string `json:"name"`
	Description         string `json:"description"`
	Language            string
	TargetType          string //Required [‘Binary’, ‘Regression’, ‘Multiclass’, ‘Anomaly’, ‘Transform’, ‘Unstructured’]
	PositiveClassLabel  string //이진분류시 참값 라벨
	NegativeClassLabel  string //이진분류시 거짓값 라벨
	ClassLabels         string //멀티클래스분류 일때 라벨
	PredictionThreshold string //이진분류 모델의 결과가 확률일 때 이진 범주에 매핑
}

type PostModelsResponse struct {
	ID string `json:"id"`
}
