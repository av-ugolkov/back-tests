package main

import (
	"fmt"
	"log"
	"net/smtp"

	"gorm.io/gorm"
)

type EmailGorm struct {
	gorm.Model
	From    string
	To      string
	Subject string
	Message string
}

type EmailRepository interface {
	Save(from string, to string, subject string, message string) error
}

type EmailDBRepository struct {
	db *gorm.DB
}

func NewEmailRepository(db *gorm.DB) EmailRepository {
	return &EmailDBRepository{
		db: db,
	}
}

func (r *EmailDBRepository) Save(from string, to string, subject string, message string) error {
	email := EmailGorm{
		From:    from,
		To:      to,
		Subject: subject,
		Message: message,
	}

	err := r.db.Create(&email).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

type EmailSender interface {
	Send(from string, to string, subject string, message string) error
}

type EmailSMTPSender struct {
	smtpHost     string
	smtpPassword string
	smtpPort     int
}

func NewEmailSender(smtpHost string, smtpPassword string, smtpPort int) EmailSender {
	return &EmailSMTPSender{
		smtpHost:     smtpHost,
		smtpPassword: smtpPassword,
		smtpPort:     smtpPort,
	}
}

func (s *EmailSMTPSender) Send(from string, to string, subject string, message string) error {
	auth := smtp.PlainAuth("", from, s.smtpPassword, s.smtpHost)

	server := fmt.Sprintf("%s:%d", s.smtpHost, s.smtpPort)

	err := smtp.SendMail(server, auth, from, []string{to}, []byte(message))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

type EmailService struct {
	repository EmailRepository
	sender     EmailSender
}

func NewEmailService(repository EmailRepository, sender EmailSender) *EmailService {
	return &EmailService{
		repository: repository,
		sender:     sender,
	}
}

func (s *EmailService) Send(from string, to string, subject string, message string) error {
	err := s.repository.Save(from, to, subject, message)
	if err != nil {
		return err
	}

	return s.sender.Send(from, to, subject, message)
}
