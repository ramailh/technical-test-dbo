package auth

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ramailh/technical-test-dbo/internal/common/consts"
	"github.com/ramailh/technical-test-dbo/internal/common/keys"
)

func (svc *authService) Logout(sessionID int64) error {
	err := svc.repo.DeleteLoginSession(sessionID)
	if err != nil {
		log.Println(err)
		return err
	}

	key := keys.CacheKeyGenerator(consts.CacheBlacklistSessionIDKey, fmt.Sprint(sessionID))
	svc.rds.Set(context.Background(), key, true, 24*time.Hour)

	return nil
}
