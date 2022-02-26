package global

import (
	"github.com/sirupsen/logrus"
	"github.com/wytfy/gredir/config"
)

var (
	CONF   *config.Config
	LOGGER *logrus.Logger
)
