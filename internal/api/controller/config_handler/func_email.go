package config_handler

import (
	"fmt"
	"net/http"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/env"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
	"github.com/xinliangnote/go-gin-api/pkg/mail"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type emailRequest struct {
	Host string `form:"host"` // Mailbox server
	Port string `form:"port"` // Port
	User string `form:"user"` // sender's mailbox
	Pass string `form:"pass"` // sender password
	To   string `form:"to"`   // Recipient's email address, multiple use, split
}

type emailResponse struct {
	Email string `json:"email"` // email address
}

// Email modify mail configuration
// @Summary modify the mail configuration
// @Description modify the mail configuration
// @Tags API.config
// @Accept multipart/form-data
// @Produce json
// @Param host formData string true "mail server"
// @Param port formData string true "port"
// @Param user formData string true "Sender's mailbox"
// @Param pass formData string true "Sender password"
// @Param to formData string true "Recipient's email address, multiple use, split"
// @Success 200 {object} emailResponse
// @Failure 400 {object} code.Failure
// @Router /api/config/email [patch]
func (h *handler) Email() core.HandlerFunc {
	return func(c core.Context) {
		req := new(emailRequest)
		res := new(emailResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		options := &mail.Options{
			MailHost: req.Host,
			MailPort: cast.ToInt(req.Port),
			MailUser: req.User,
			MailPass: req.Pass,
			MailTo:   req.To,
			Subject:  fmt.Sprintf("%s[%s] Email alerter adjustment notification.", configs.ProjectName(), env.Active().Value()),
			Body:     fmt.Sprintf("%s[%s] You have been added as the system alarm notifier.", configs.ProjectName(), env.Active().Value()),
		}
		if err := mail.Send(options); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ConfigEmailError,
				"Mail Send error: "+err.Error()).WithErr(err),
			)
			return
		}

		viper.SetConfigName(env.Active().Value() + "_configs")
		viper.SetConfigType("toml")
		viper.AddConfigPath("./configs")

		viper.Set("mail.host", req.Host)
		viper.Set("mail.port", cast.ToInt(req.Port))
		viper.Set("mail.user", req.User)
		viper.Set("mail.pass", req.Pass)
		viper.Set("mail.to", req.To)

		err := viper.WriteConfig()
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ConfigEmailError,
				code.Text(code.ConfigEmailError)).WithErr(err),
			)
			return
		}

		res.Email = req.To
		c.Payload(res)
	}
}
