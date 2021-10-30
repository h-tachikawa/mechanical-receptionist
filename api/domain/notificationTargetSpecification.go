package domain

import (
	"time"
)

type NotificationTargetSpecification struct {
	previous time.Time
	current  time.Time
}

func NewNotificationTargetSpecification(previous, current time.Time) NotificationTargetSpecification {
	return NotificationTargetSpecification{
		previous,
		current,
	}
}

func (d NotificationTargetSpecification) IsSatisfied() bool {
	duration := d.current.Sub(d.previous).Seconds()

	if duration > 60 {
		return true
	}
	return false
}
