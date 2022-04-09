package gost

import (
	"context"
	"fmt"

	"github.com/showwin/speedtest-go/speedtest"
)

// SpeedTestProvider is used to implement Provider.
type SpeedTestProvider struct {
	servers speedtest.Servers
}

// NewSpeedTestProvider constructor for SpeedTestProvider.
func NewSpeedTestProvider() (*SpeedTestProvider, error) {
	user, err := speedtest.FetchUserInfo()
	if err != nil {
		fmt.Println("Warning: Cannot fetch user information. http://www.speedtest.net/speedtest-config.php is temporarily unavailable.")
	}

	servers, err := speedtest.FetchServers(user)
	if err != nil {
		return nil, err
	}

	r := &SpeedTestProvider{
		servers: servers,
	}

	return r, nil
}

// Download implements Provider.
func (r *SpeedTestProvider) Download(ctx context.Context) (float64, error) {
	s := r.servers[0]
	if err := r.servers[0].DownloadTestContext(ctx, false); err != nil {
		return 0, err
	}

	return s.DLSpeed, nil
}

// Upload implements Provider.
func (r *SpeedTestProvider) Upload(ctx context.Context) (float64, error) {
	s := r.servers[0]
	if err := r.servers[0].UploadTestContext(ctx, false); err != nil {
		return 0, err
	}

	return s.ULSpeed, nil
}
