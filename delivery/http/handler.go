package http

import (
	"golang-project/config/dbiface"
)

type Handler struct {
	Col dbiface.CollectionAPI
}
