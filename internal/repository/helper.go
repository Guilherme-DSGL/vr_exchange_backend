package repository

import (
	"encoding/base64"
	"time"

	"github.com/Guilherme-DSGL/vr_exchange_backend/utils"
)

func DecodeCursor(encodedTime string) (time.Time, error) {
	byt, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, err
	}

	timeString := string(byt)
	t, err := time.Parse(utils.DateTimeFormat, timeString)

	return t, err
}

func EncodeCursor(t time.Time) string {
	timeString := t.Format(utils.DateTimeFormat)

	return base64.StdEncoding.EncodeToString([]byte(timeString))
}
