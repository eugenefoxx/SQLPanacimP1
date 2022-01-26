package panacim

import (
	"database/sql"

	"github.com/eugenefoxx/SQLPanacimP1/pkg/logging"
)

type PanaCIMStorage struct {
	DB     *sql.DB
	logger *logging.Logger
}
