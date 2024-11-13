package job

import (
	"app/config"
	"app/utils"
)

type emailJob struct {
	emailChan chan config.EmailJob_MessPayload
	smtp      utils.SmtpUtils
}

type EmailJob interface {
	handle()
	PushJob(data config.EmailJob_MessPayload)
}

func (j *emailJob) handle() {
	for q := range j.emailChan {
		go func(data config.EmailJob_MessPayload) {
			j.smtp.SendEmail(data.Content, data.Email)
		}(q)
	}
}

func (j *emailJob) PushJob(data config.EmailJob_MessPayload) {
	j.emailChan <- data
}

func NewEmailJob() EmailJob {
	return &emailJob{
		emailChan: config.GetEmailChan(),
		smtp:      utils.NewSmtpUtils(),
	}
}
