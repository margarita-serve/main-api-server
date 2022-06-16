package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddRoutes(e *echo.Echo) {
	e.POST("/modelPackages", Post)
	e.GET("/modelPackages", GetAll)
	e.GET("/modelPackages/{modelPackageId}", Get)
	e.PATCH("/modelPackages/{modelPackageId}", Patch)
	e.DELETE("/modelPackages/{modelPackageId}", Delete)
	e.GET("/modelPackages/locationTypeList", GetLocationTypeList)
	e.GET("/modelPackages/targetTypeList", GetTargetTypeList)
	e.GET("/modelPackages/runtimeFrameWrokList", GetRuntimeFrameWrokList)
}

type PostModelPackagesRequest struct {
	ProjectID            string               `json:"projectID" validate:"required" example:"ID-1234" extensions:"x-order=0 x-nullable=false"`         // 모델 패키지를 등록할 프로젝트 ID
	Name                 string               `json:"name" validate:"required" example:"ModelPackage Example" extensions:"x-order=1 x-nullable=false"` // 모델 패키지 명
	Description          string               `json:"description" example:"Description of this package" extensions:"x-order=2"`                        // 모델 패키지 설명
	Model                PostModelInfo        `json:"model" extensions:"x-order=3"`                                                                    // 패키지에 추가할 모델의 정보
	TargetType           string               `json:"targetType" extensions:"x-order=4" enums:"Binary,Regression" example:"‘Binary’"`                  // 모델의 예측타입 [‘Binary’, ‘Regression’, ‘Multiclass’, ‘Unstructured’]
	Features             string               `json:"features" extensions:"x-order=5" example:"feature_name1, feature_name2,....’"`
	TrainingDataset      PostTrainingDataset  `json:"trainingDataset"`      // 훈련 데이터셋
	HoldoutDataset       PostHoldoutDataset   `json:"holdoutDataset"`       // 홀드아웃 데이터셋
	PredictionTargetName string               `json:"predictionTargetName"` // 예측 타켓컬럼 명칭
	PositiveClassLabel   string               `json:"positiveClassLabel"`   // 모델의 예측 타입이 'Binary' 일 때 Positive클래스 명
	NegativeClassLabel   string               `json:"negativeClassLabel"`   // 모델의 예측 타입이 'Binary' 일 때 Negative클래스 명
	Threshold            float32              `json:"threshold"`            // 모델의 예측 타입이 'Binary' 일 때 Threshold
	Transformer          TransformerInfo      `json:"transformer"`          // 모델 입력 전처리기
	RuntimeFrameWork     PostRuntimeFrameWork `json:"runtimeFrameWork"`     // 서빙 런타임 프레임워크
}

type PostModelPackagesResponse struct {
	Message string `json:"message"` // 상세 메세지
}

type PostModelInfo struct {
	Name         string     `json:"name" example:"Model Example"` //모델명
	Version      string     `json:"version" example:"V1"`         //버전
	Description  string     `json:"description"`                  //설명
	LocationType string     `json:"locationType"`                 //모델파일이 저장되어 있는 스토리지의 종류 ['Local', 'BasicAuth', 'S3' ,'GS']
	URL          string     `json:"URL"`                          //스토리지 URL
	Credential   Credential `json:"credential"`                   //스토리지 접근 인증정보
}

type PostTrainingDataset struct {
	TargetName   string     `json:"targetName"`
	LocationType string     `json:"locationType"` //파일이 저장되어 있는 스토리지의 종류['Local', 'BasicAuth', 'S3' ,'GS']
	URL          string     `json:"URL"`          //스토리지 URL
	Credential   Credential `json:"credential"`   //스토리지 접근 인증정보
}

type PostHoldoutDataset struct {
	TargetName   string     `json:"targetName"`
	LocationType string     `json:"locationType"` //파일이 저장되어 있는 스토리지의 종류['Local', 'BasicAuth', 'S3' ,'GS']
	URL          string     `json:"URL"`          //스토리지 URL
	Credential   Credential `json:"credential"`   //스토리지 접근 인증정보
}

type TransformerInfo struct {
	LocationType string     `json:"locationType"` //파일이 저장되어 있는 스토리지의 종류 ['Local', 'BasicAuth', 'S3' ,'GS']
	URL          string     `json:"URL"`
	Credential   Credential `json:"credential"`
}

type Credential struct {
	BasicUserID        string //기본 유저인증. ID Location type이 'BasicAuth' 일때 사용
	BasicUserPassword  string //기본 유저인증 패스워드. ID Location type이 'BasicAuth' 일때 사용
	AWSaccesskeyID     string //AWS 엑세스 키 ID. ID Location type이 'S3' 일때 사용
	AWSsecretAccessKey string //AWS 시크릿 엑세스키. ID Location type이 'S3' 일때 사용
	GCPkey             string //GCP 키. ID Location type이 'GS' 일때 사용
}
type GetTargetType struct {
	Type string `json:"type" example:"[‘Binary’, ‘Regression’, ‘Multiclass’, ‘Unstructured’]"` // 모델의 예측종류 [‘Binary’, ‘Regression’, ‘Multiclass’, ‘Unstructured’]
}
type GetLocationType struct {
	Type string `json:"type" example:"['Local','BasicAuth','S3', 'GS']"` // ['Local', 'BasicAuth', 'S3' ,'GS']
}

type PostRuntimeFrameWork struct {
	Type    string `json:"type" example:"['Tensorflow','PyTorch','Scikit-learn','XGBoost','PMML','Spark''Lightgbm','PaddleServer','Triton']"` //런타임 프레임워크 ['Tensorflow','PyTorch','Scikit-learn','XGBoost','PMML','Spark''Lightgbm','PaddleServer','Triton']
	Version string `json:"version" example:"0.12"`                                                                                            //런타임 프레임워크 버전
}

type GetRuntimeFrameWork struct {
	Type string `json:"type" example:"['Tensorflow','PyTorch','Scikit-learn','XGBoost','PMML','Spark''Lightgbm','PaddleServer','Triton']"` //런타임 프레임워크 ['Tensorflow','PyTorch','Scikit-learn','XGBoost','PMML','Spark''Lightgbm','PaddleServer','Triton']
}

type PatchModelPackagesRequest struct {
	Name        string `json:"name" example:"Model Package Example" extensions:"x-order=0"`              // 모델 패키지 명
	Description string `json:"description" example:"Description of this package" extensions:"x-order=1"` // 모델 패키지 설명
}

type GetAllModelPackageResponse struct {
	Count       int    `json:"count" extensions:"x-order=0"`                                                                    //Optional. Number of items returned on this page.
	Next        string `json:"next" extensions:"x-order=1"`                                                                     //Required. URL pointing to the next page (if null, there is no next page). nullable: True format: uri
	Previous    string `json:"previous" extensions:"x-order=2"`                                                                 //Required. URL pointing to the previous page (if null, there is no previous page). nullable: True format: uri
	Name        string `json:"name" validate:"required" example:"ModelPackage Example" extensions:"x-order=1 x-nullable=false"` // 모델 패키지 명
	Description string `json:"description" validate:"required" example:"Description of this package" extensions:"x-order=2"`    // 모델 패키지 설명
	TargetType  string `json:"targetType" extensions:"x-order=4" enums:"Binary,Regression" example:"‘Binary’"`                  // 모델의 예측종류 [‘Binary’, ‘Regression’, ‘Multiclass’, ‘Unstructured’]
}

type GetModelPackageResponse struct {
	ProjectID            string             //프로젝트 ID
	ProjectName          string             //프로젝트 명
	Name                 string             `json:"name" validate:"required" example:"ModelPackage Example" extensions:"x-order=1 x-nullable=false"` // 모델 패키지 명
	Description          string             `json:"description" example:"Description of this package" extensions:"x-order=2"`                        // 모델 패키지 설명
	Model                GetModelInfo       `json:"model" extensions:"x-order=3"`                                                                    // 패키지에 추가할 모델의 정보
	TargetType           string             `json:"targetType" extensions:"x-order=4" enums:"Binary,Regression" example:"‘Binary’"`                  // 모델의 예측종류 [‘Binary’, ‘Regression’, ‘Multiclass’, ‘Unstructured’]
	Features             string             `json:"features" extensions:"x-order=5" example:"feature_name1, feature_name2,....’"`
	TrainingDataset      GetTrainingDataset `json:"trainingDataset"`      // 훈련 데이터셋
	HoldoutDataset       GetHoldoutDataset  `json:"holdoutDataset"`       // 홀드아웃 데이터셋
	PredictionTargetName string             `json:"predictionTargetName"` // 타켓컬럼 명칭
	PositiveClassLabel   string             `json:"positiveClassLabel"`   //
	NegativeClassLabel   string             `json:"negativeClassLabel"`   //
	Threshold            float32            `json:"threshold"`            //
	Transformer          TransformerInfo    `json:"transformer"`          // 모델 입력 전처리기
	RuntimeFrameWork     string             `json:"runtimeFrameWork"`     // 서빙 런타임 프레임워크
}

type GetModelInfo struct {
	ID           string     `json:"id" example:"Model Example"`   //모델ID
	Name         string     `json:"name" example:"Model Example"` //모델명
	Version      string     `json:"version" example:"V1"`         //버전
	Description  string     `json:"description"`                  //설명
	LocationType string     `json:"locationType"`                 //모델파일이 저장되어 있는 스토리지의 종류 ['Local', 'BasicAuth', 'S3' ,'GS']
	URL          string     `json:"URL"`                          //스토리지 URL
	Credential   Credential `json:"credential"`                   //스토리지 접근 인증정보
}

type GetTrainingDataset struct {
	TargetName   string     `json:"targetName"`
	LocationType string     `json:"locationType"` //파일이 저장되어 있는 스토리지의 종류['Local', 'BasicAuth', 'S3' ,'GS']
	URL          string     `json:"URL"`          //스토리지 URL
	Credential   Credential `json:"credential"`   //스토리지 접근 인증정보
}

type GetHoldoutDataset struct {
	TargetName   string     `json:"targetName"`
	LocationType string     `json:"locationType"` //파일이 저장되어 있는 스토리지의 종류['Local', 'BasicAuth', 'S3' ,'GS']
	URL          string     `json:"URL"`          //스토리지 URL
	Credential   Credential `json:"credential"`   //스토리지 접근 인증정보
}

// @Summary Create Package
// @Description 모델을 배포하기 위한 패키지 등록
// @Tags ModelPackage
// @Accept json
// @Produce json
// @Param body body PostModelPackagesRequest true "Info of Package"
// @Success 200 {object} PostModelPackagesResponse
// @Router /modelPackages [post]
func Post(c echo.Context) error {

	//...

	post_response := new(PostModelPackagesResponse)
	return c.JSONPretty(http.StatusOK, *post_response, "  ")
}

// @Summary 모델패키지 리스트 조회
// @Description
// @Tags ModelPackage
// @Accept json
// @Produce json
// @Param offset query integer false "Optional. This many results will be skipped. minimum: 0 default: 0"
// @Param limit query integer false "Optional. At most this many results are returned. minimum: 1 maximum: 1000 default: 20"
// @Success 200 {object} GetAllModelPackageResponse
// @Router /modelPackages [get]
func GetAll(c echo.Context) error {

	//...

	post_response := new(GetAllModelPackageResponse)
	return c.JSONPretty(http.StatusOK, *post_response, "  ")
}

// @Summary 모델패키지 조회
// @Description
// @Tags ModelPackage
// @Accept json
// @Produce json
// @Param modelPackageID path string true "Required. "
// @Success 200 {object} GetModelPackageResponse
// @Router /modelPackages/{id} [get]
func Get(c echo.Context) error {

	//...

	post_response := new(GetModelPackageResponse)
	return c.JSONPretty(http.StatusOK, *post_response, "  ")
}

// @Summary 모델패키지 아카이브
// @Description
// @Tags ModelPackage
// @Accept json
// @Produce json
// @Param modelPackageID path string true "Required. "
// @Success      200  {string}  string    "ok"
// @Router /modelPackages/{id} [delete]
func Delete(c echo.Context) error {
	//...

	post_response := new(PostModelPackagesResponse)
	return c.JSONPretty(http.StatusOK, *post_response, "  ")
}

// @Summary 모델패키지 수정
// @Description
// @Tags ModelPackage
// @Accept json
// @Produce json
// @Param modelPackageID query string true "Required. "
// @Param body body PatchModelPackagesRequest true "Archive Package"
// @Success      200  {string}  string    "ok"
// @Router /modelPackages/{id} [patch]
func Patch(c echo.Context) error {

	//...

	post_response := new(PostModelPackagesResponse)
	return c.JSONPretty(http.StatusOK, *post_response, "  ")
}

// @Summary Model's Location Type List Request
// @Description 모델 패키지 등록을 위한 모델파일 위치 종류 선택 정보
// @Tags ModelPackage
// @Accept json
// @Produce json
// @Success 200 {object} GetLocationType
// @Router /modelPackages/locationType [get]
func GetLocationTypeList(c echo.Context) error {

	//...

	post_response := new(GetLocationType)
	return c.JSONPretty(http.StatusOK, *post_response, "  ")
}

// @Summary Model's Target Type List Request
// @Description 모델 패키지 등록을 위한 모델의 예측 종류 리스트
// @Tags ModelPackage
// @Accept json
// @Produce json
// @Success 200 {object} GetTargetType
// @Router /modelPackages/targetType [get]
func GetTargetTypeList(c echo.Context) error {

	//...

	post_response := new(GetTargetType)
	return c.JSONPretty(http.StatusOK, *post_response, "  ")
}

// @Summary Model's Runtime FrameWork List List Request
// @Description 모델 패키지 등록을 위한 모델이 서빙 될 프레임워크 리스트 정보
// @Tags ModelPackage
// @Accept json
// @Produce json
// @Success 200 {object} GetRuntimeFrameWork
// @Router /modelPackages/runtimeFrameWorkList [get]
func GetRuntimeFrameWrokList(c echo.Context) error {

	//...

	post_response := new(GetRuntimeFrameWork)
	return c.JSONPretty(http.StatusOK, *post_response, "  ")
}
