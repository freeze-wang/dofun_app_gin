package handlers

import (
	"dofun/config"
	"dofun/pkg/ginutils/mail"
	"dofun/pkg/ginutils/router"
	"path"

	userModel "dofun/app/models/user"

	"github.com/flosch/pongo2"
)

// SendMail 发送邮件
func SendMail(mailTo []string, subject string, templateName string, tplData map[string]interface{}) error {
	filename := path.Join(config.AppConfig.ViewsPath, templateName)
	template := pongo2.Must(pongo2.FromCache(filename))

	body, err := template.Execute(tplData)
	if err != nil {
		return err
	}

	mail := &mail.Mail{
		Driver:   config.MailConfig.Driver,
		Host:     config.MailConfig.Host,
		Port:     config.MailConfig.Port,
		User:     config.MailConfig.User,
		Password: config.MailConfig.Password,
		FromName: config.MailConfig.FromName,
		MailTo:   mailTo,
		Subject:  subject,
		Body:     body,
	}

	return mail.Send()
}

// SendVerifyEmail 发送激活用户的邮件
func SendVerifyEmail(u *userModel.User) error {
	subject := "感谢注册 Weibo 应用！请确认你的邮箱。"
	tpl := "mail/verify.html"
	verifyURL := router.G("verification.verify", "token", u.LastToken)

	return SendMail([]string{u.Email}, subject, tpl, map[string]interface{}{"URL": verifyURL})
}

