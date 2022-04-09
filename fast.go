package gost

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ddo/go-fast"
	"golang.org/x/sync/errgroup"
)

// FastProvider is used to implement Provider.
type FastProvider struct {
	f          *fast.Fast
	urls       []string
	savingMode bool
	doer       *http.Client
}

// NewFastProvider constructor for FastProvider.
func NewFastProvider() (*FastProvider, error) {
	fastCom := fast.New()

	err := fastCom.Init()
	if err != nil {
		return nil, err
	}

	urls, err := fastCom.GetUrls()
	if err != nil {
		return nil, err
	}

	r := &FastProvider{
		f:    fastCom,
		urls: urls,
	}

	r.doer = http.DefaultClient

	return r, nil
}

// Download implements Provider.
func (r *FastProvider) Download(ctx context.Context) (float64, error) {
	u := r.urls[0]
	eg := errgroup.Group{}

	// Warming up
	sTime := time.Now()
	for i := 0; i < 2; i++ {
		eg.Go(func() error {
			return dlWarmUp(ctx, r.doer, u)
		})
	}
	if err := eg.Wait(); err != nil {
		return 0, err
	}
	fTime := time.Now()

	// If the bandwidth is too large, the download sometimes finish earlier than the latency.
	// In this case, we ignore the the latency that is included server information.
	// This is not affected to the final result since this is a warm up test.
	timeToSpend := fTime.Sub(sTime.Add(time.Millisecond)).Seconds()
	if timeToSpend < 0 {
		timeToSpend = fTime.Sub(sTime).Seconds()
	}

	// 1.125MB for each request (750 * 750 * 2)
	wuSpeed := 1.125 * 8 * 2 / timeToSpend

	// Decide workload by warm up speed
	workload := 0
	weight := 0
	skip := false
	if r.savingMode {
		workload = 6
		weight = 3
	} else if 50.0 < wuSpeed {
		workload = 32
		weight = 6
	} else if 10.0 < wuSpeed {
		workload = 16
		weight = 4
	} else if 4.0 < wuSpeed {
		workload = 8
		weight = 4
	} else if 2.5 < wuSpeed {
		workload = 4
		weight = 4
	} else {
		skip = true
	}

	// Main speedtest
	dlSpeed := wuSpeed
	if !skip {
		sTime = time.Now()
		for i := 0; i < workload; i++ {
			eg.Go(func() error {
				return downloadRequest(ctx, r.doer, u)
			})
		}
		if err := eg.Wait(); err != nil {
			return 0, err
		}
		fTime = time.Now()

		reqMB := dlSizes[weight] * dlSizes[weight] * 2 / 1000 / 1000
		dlSpeed = float64(reqMB) * 8 * float64(workload) / fTime.Sub(sTime).Seconds()
	}

	return dlSpeed, nil
}

// Upload implements Provider.
func (r *FastProvider) Upload(ctx context.Context) (float64, error) {
	u := r.urls[0]

	// Warm up
	sTime := time.Now()
	eg := errgroup.Group{}
	for i := 0; i < 2; i++ {
		eg.Go(func() error {
			return ulWarmUp(ctx, r.doer, u)
		})
	}
	if err := eg.Wait(); err != nil {
		return 0, err
	}
	fTime := time.Now()
	// 1.0 MB for each request
	wuSpeed := 1.0 * 8 * 2 / fTime.Sub(sTime.Add(time.Microsecond)).Seconds()

	// Decide workload by warm up speed
	workload := 0
	weight := 0
	skip := false
	if r.savingMode {
		workload = 1
		weight = 7
	} else if 50.0 < wuSpeed {
		workload = 40
		weight = 9
	} else if 10.0 < wuSpeed {
		workload = 16
		weight = 9
	} else if 4.0 < wuSpeed {
		workload = 8
		weight = 9
	} else if 2.5 < wuSpeed {
		workload = 4
		weight = 5
	} else {
		skip = true
	}

	// Main speedtest
	ulSpeed := wuSpeed
	if !skip {
		sTime = time.Now()
		for i := 0; i < workload; i++ {
			eg.Go(func() error {
				return uploadRequest(ctx, r.doer, u, weight)
			})
		}
		if err := eg.Wait(); err != nil {
			return 0, err
		}
		fTime = time.Now()

		reqMB := float64(ulSizes[weight]) / 1000
		ulSpeed = reqMB * 8 * float64(workload) / fTime.Sub(sTime).Seconds()
	}

	return ulSpeed, nil
}

var dlSizes = [...]int{350, 500, 750, 1000, 1500, 2000, 2500, 3000, 3500, 4000}
var ulSizes = [...]int{100, 300, 500, 800, 1000, 1500, 2500, 3000, 3500, 4000} // kB

// uses as the same method in speedtest package
func ulWarmUp(ctx context.Context, doer *http.Client, ulURL string) error {
	size := ulSizes[4]
	v := url.Values{}
	v.Add("content", strings.Repeat("0123456789", size*100-51))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ulURL, strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := doer.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(ioutil.Discard, resp.Body)
	return err
}

// uses as the same method in speedtest package
func uploadRequest(ctx context.Context, doer *http.Client, ulURL string, w int) error {
	size := ulSizes[w]
	v := url.Values{}
	v.Add("content", strings.Repeat("0123456789", size*100-51))

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ulURL, strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := doer.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(ioutil.Discard, resp.Body)
	return err
}

// uses as the same method in speedtest package
func dlWarmUp(ctx context.Context, doer *http.Client, dlURL string) error {
	xdlURL := dlURL

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, xdlURL, nil)
	if err != nil {
		return err
	}

	resp, err := doer.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(ioutil.Discard, resp.Body)
	return err
}

// uses as the same method in speedtest package
func downloadRequest(ctx context.Context, doer *http.Client, dlURL string) error {
	xdlURL := dlURL

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, xdlURL, nil)
	if err != nil {
		return err
	}

	resp, err := doer.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(ioutil.Discard, resp.Body)
	return err
}
