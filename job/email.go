package jobapp

import (
	"app/internal/connection"
	smptapp "app/pkg/smpt"
)

type emailJob struct {
	emailChan chan connection.EmailJob_MessPayload
}

type EmailJob interface {
	handle()
	PushJob(data connection.EmailJob_MessPayload)
}

func (j *emailJob) handle() {
	for q := range j.emailChan {
		go func(data connection.EmailJob_MessPayload) {
			smptapp.SendEmail(data.Content, data.Email)
		}(q)
	}
}

func (j *emailJob) PushJob(data connection.EmailJob_MessPayload) {
	j.emailChan <- data
}

func NewEmailJob() EmailJob {
	return &emailJob{
		emailChan: connection.GetEmailChan(),
	}
}
