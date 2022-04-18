package data

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestFindIndexByUserID(t *testing.T) {
	log := &logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		},
		Level: logrus.DebugLevel,
	}

	_ = NewUserDetail(*log)
	//log.Info(user.User)

	t.Run("User exist", func(t *testing.T) {
		res, err := FindIndexByUserID(1)
		if err != nil {
			log.Error(err)
		}
		got := res
		expected := 0
		assert.Equal(t, expected, got)
	})

	t.Run("User id does not exist", func(t *testing.T) {
		res, err := FindIndexByUserID(0)
		if err != nil {
			log.Error(err)
		}
		got := res
		expected := -1
		assert.Equal(t, expected, got)
	})
}
