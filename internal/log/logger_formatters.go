package log

import (
	"github.com/sirupsen/logrus"

	"github.com/anhnmt/drprune/internal/consts"
)

// Configure the logrus format to use "text" formatter
func textFormatter() *logrus.TextFormatter {
	return &logrus.TextFormatter{
		DisableColors:             true,
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		TimestampFormat:           consts.DefaultTimestampFormat,
		DisableLevelTruncation:    true,
	}
}

// Configure the logrus format to use "color" formatter
func colorFormatter() *logrus.TextFormatter {
	return &logrus.TextFormatter{
		DisableColors:             false,
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		TimestampFormat:           consts.DefaultTimestampFormat,
		DisableLevelTruncation:    true,
	}
}

// Configure the logrus format to use "json" formatter
func jsonFormatter(pretty bool) *logrus.JSONFormatter {
	return &logrus.JSONFormatter{
		TimestampFormat:  consts.DefaultTimestampFormat,
		DisableTimestamp: false,
		PrettyPrint:      pretty,
	}
}
