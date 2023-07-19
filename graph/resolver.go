package graph

import (
	"github.com/vgekko/ani-go/internal/usecase"
	"golang.org/x/exp/slog"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	Uc *usecase.UseCase
	Log *slog.Logger
}
