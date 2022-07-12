package services

import (
	"net/url"
	"strconv"

	"github.com/pkg/errors"
	app "github.com/sbward/authn"
	"github.com/sbward/authn/models"
	"github.com/sbward/authn/tokens/resets"
	"github.com/sirupsen/logrus"
)

func PasswordResetSender(cfg *app.Config, account *models.Account, logger logrus.FieldLogger) error {
	if account == nil || account.Locked {
		return nil
	}

	reset, err := resets.New(cfg, account.ID, account.PasswordChangedAt)
	if err != nil {
		return errors.Wrap(err, "New Reset")
	}
	resetStr, err := reset.Sign(cfg.ResetSigningKey)
	if err != nil {
		return errors.Wrap(err, "Sign")
	}

	err = WebhookSender(cfg.AppPasswordResetURL, &url.Values{
		"account_id": []string{strconv.Itoa(account.ID)},
		"token":      []string{resetStr},
	}, timeSensitiveDelivery)
	if err != nil {
		return errors.Wrap(err, "Webhook")
	}

	logger.WithField("accountID", account.ID).Info("sent password reset token")

	return nil
}
