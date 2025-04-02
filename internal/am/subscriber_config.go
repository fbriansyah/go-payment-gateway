package am

import "time"

type AckType int

const (
	AckTypeAuto AckType = iota
	AckTypeManual
)

var defaultAckWait = 30 * time.Second
var defaultMaxRedeliver = 5

type SubscriberConfig struct {
	msgFilter []string
	groupName string
	ackType AckType
	ackWait time.Duration
	maxRedeliver int
}

type SubscriberOption interface {
	configureSubscriberConfig(*SubscriberConfig)
}

func NewSubscriberConfig(options []SubscriberOption) SubscriberConfig {
	cfg := SubscriberConfig{
		msgFilter: []string{},
		groupName: "",
		ackType: AckTypeManual,
		ackWait: defaultAckWait,
		maxRedeliver: defaultMaxRedeliver,
	}

	for _, option := range options {
		option.configureSubscriberConfig(&cfg)
	}

	return cfg
}