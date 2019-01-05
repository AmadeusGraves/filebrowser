package bolt

import (
	"github.com/filebrowser/filebrowser/v2/settings"
	"github.com/asdine/storm"
	"github.com/filebrowser/filebrowser/v2/auth"
	"github.com/filebrowser/filebrowser/v2/share"
	"github.com/filebrowser/filebrowser/v2/storage"
	"github.com/filebrowser/filebrowser/v2/users"
)

// NewStorage creates a storage.Storage based on Bolt DB.
func NewStorage(db *storm.DB) *storage.Storage {
	users := users.NewStorage(usersBackend{db: db})
	share := share.NewStorage(shareBackend{db: db})
	settings := settings.NewStorage(settingsBackend{ db: db})
	auth := auth.NewStorage(authBackend{db: db}, users)

	return &storage.Storage{
		Auth:  auth,
		Users: users,
		Share: share,
		Settings: settings,
	}
}
