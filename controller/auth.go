package controller

import (
	"app/config"
	"app/constant"
	"app/dto/request"
	"app/dto/response"
	"app/job"
	"app/service"
	"app/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type authController struct {
	authService service.AuthService
	jwtUtils    utils.JwtUtils
	emailJob    job.EmailJob
	redis       *redis.Client
}

type AuthControlle interface {
	Login(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	AcceptCopde(w http.ResponseWriter, r *http.Request)
}

func (c *authController) Register(w http.ResponseWriter, r *http.Request) {
	var payload request.RegisterReq

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, r, err)
		return
	}

	uuid, err := uuid.NewV6()
	if err != nil {
		InternalServerError(w, r, err)
	}

	code := c.authService.CreateCode(constant.LENGTH_CODE)

	err = c.authService.SaveInfoRegsiter(uuid.String(), code, payload)
	if err != nil {
		InternalServerError(w, r, err)
	}

	expToken := time.Now().Add(time.Second * time.Duration(constant.EXP_INFO_REGISTER))
	infoToken := map[string]interface{}{
		"uuid": uuid,
		"exp":  expToken,
	}

	token, err := c.jwtUtils.JwtEncode(infoToken)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	c.emailJob.PushJob(config.EmailJob_MessPayload{
		Email:   payload.Email,
		Content: code,
	})

	res := Response{
		Data: response.RegisterRes{
			Token: token,
			Exp:   expToken,
		},
	}

	render.JSON(w, r, res)
}

func (c *authController) AcceptCopde(w http.ResponseWriter, r *http.Request) {
	var payload request.AcceptCode

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, r, err)
		return
	}

	mapInfoToken, err := c.jwtUtils.JwtDecode(c.jwtUtils.GetToken(r))
	if err != nil {
		HandleError(w, r, err, 404)
		return
	}

	uuid := fmt.Sprint(mapInfoToken["uuid"])
	val, err := c.redis.Get(r.Context(), uuid).Result()
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	var infoRegister service.SaveInfoRegisterPayload
	err = json.Unmarshal([]byte(val), &infoRegister)
	if err != nil {
		InternalServerError(w, r, err)
	}

	if payload.Code != infoRegister.Code {
		InternalServerError(w, r, errors.New("code wrong"))
	}

	err = c.authService.CreateProfile(uuid)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	res := Response{
		Data:    nil,
		Message: "OK",
		Status:  200,
		Error:   nil,
	}

	render.JSON(w, r, res)
}

func (c *authController) Login(w http.ResponseWriter, r *http.Request) {
	var payload request.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		BadRequest(w, r, err)
		return
	}

	profile, err := c.authService.CompareProfile(payload)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	mapInfo := map[string]interface{}{
		"email": profile.Email,
		"role":  profile.Role.Code,
		"id":    profile.ID,
	}

	accessToken, refreshToken, err := c.authService.CreateToken(mapInfo)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}
	profile.Password = ""

	res := Response{
		Data: response.LoginRes{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			Profile:      *profile,
		},
	}

	render.JSON(w, r, res)
}

func (c *authController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	token := c.jwtUtils.GetToken(r)
	mapToken, err := c.jwtUtils.JwtDecode(token)

	if err != nil {
		HandleError(w, r, err, 404)
		return
	}

	profileId := uint(mapToken["id"].(float64))
	profile, err := c.authService.GetProfile(profileId)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}

	mapInfo := map[string]interface{}{
		"email": profile.Email,
		"role":  profile.Role.Code,
		"id":    profile.ID,
	}

	accessToken, refreshToken, err := c.authService.CreateToken(mapInfo)
	if err != nil {
		InternalServerError(w, r, err)
		return
	}
	profile.Password = ""

	res := Response{
		Data: response.LoginRes{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			Profile:      *profile,
		},
	}

	render.JSON(w, r, res)
}

func NewAuthController() AuthControlle {
	return &authController{
		authService: service.NewAuthService(),
		jwtUtils:    utils.NewJwtUtils(),
		redis:       config.GetRedisClient(),
		emailJob:    job.NewEmailJob(),
	}
}
