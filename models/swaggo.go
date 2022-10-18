package models

import "time"

type RegisterRequest struct {
	Age      uint   `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type RegisterResponse struct {
	Age      uint   `json:"age"`
	Email    string `json:"email"`
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UpdateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdateUserResponse struct {
	Id uint `json:"id"`
	UpdateUserRequest
	Age       uint      `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteResponse struct {
	Message string `json:"message"`
}

type CreatePhotoRequest struct {
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url"`
}

type CreatePhotoResponse struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	Photo_url string    `json:"photo_url"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetPhotoResponse struct {
	CreatePhotoResponse
	UpdatedAt time.Time `json:"updated_at"`
	UserP
}

type UpdatePhotoRequest struct {
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url"`
}

type UpdatePhotoResponse struct {
	Id uint `json:"id"`
	UpdatePhotoRequest
	UserId    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCommentRequest struct {
	Message  string `json:"message"`
	Photo_id uint   `json:"photo_id"`
}

type CreateCommentResponse struct {
	Id uint `json:"id"`
	CreateCommentRequest
	UserId    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetCommentResponse struct {
	CreateCommentRequest
	CreatedAt time.Time `json:"created_at"`
	User      UserC
	Photo     PhotoC
}

type UpdateCommentRequest struct {
	DeleteResponse
}

type UpdateCommentResponse struct {
	Id uint `json:"id"`
	UpdateCommentRequest
	UserId    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateSocialMediaRequest struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

type CreateSocialMediaResponse struct {
	Id uint `json:"id"`
	CreateSocialMediaRequest
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetSocialMediaResponseItem struct {
	CreateSocialMediaRequest
	UpdatedAt time.Time `json:"updated_at"`
	User      UserS
}

type GetSocialMediaResponse struct {
	SocialMedias []GetSocialMediaResponseItem `json:"social_medias"`
}

type UpdateSocialMediaRequest struct {
	CreateSocialMediaRequest
}

type UpdateSocialMediaResponse struct {
	Id uint `json:"id"`
	UpdateSocialMediaRequest
	UserId    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
