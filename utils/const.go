package utils

import "errors"

var (
	ErrExpiredToken = errors.New("Токений хугацаа дууссан байна")
	ErrInvalidToken = errors.New("Токен хандах эрхгүй байна")
)
