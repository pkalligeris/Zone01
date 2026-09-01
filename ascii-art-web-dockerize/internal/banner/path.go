package banner

import (
	"errors"
	"fmt"
)

var ErrUnknownBanner = errors.New("unknown banner")

var supportedBanners = map[string]string{
	"standard":   "assets/banners/standard.txt",
	"shadow":     "assets/banners/shadow.txt",
	"thinkertoy": "assets/banners/thinkertoy.txt",
}

func GetBannerPath(name string) (string, error) {
	path, ok := supportedBanners[name]
	if !ok {
		return "", fmt.Errorf("%w: %s", ErrUnknownBanner, name)
	}
	return path, nil
}

