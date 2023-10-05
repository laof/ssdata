package ssdata

import (
	"context"
	"flag"
	"log"
	"strings"
	"time"

	"github.com/laof/lite-speed-test/web"
)

type Results struct {
	Services     []string
	Failed       []string
	Success      []string
	FailedNodes  []string
	SuccessNodes []string
}

func profile(url string, rs Results, sm map[string][]string) Results {

	link := flag.String("link", url, "link to test")
	mode := flag.String("mode", "pingonly", "speed test mode")
	flag.Parse()
	// link := "vmess://aHR0cHM6Ly9naXRodWIuY29tL3h4ZjA5OC9MaXRlU3BlZWRUZXN0"
	if len(*link) < 1 {
		log.Fatal("link required")
	}
	opts := web.ProfileTestOptions{
		GroupName:     "Default",
		SpeedTestMode: *mode,        //  pingonly speedonly all
		PingMethod:    "googleping", // googleping
		SortMethod:    "rspeed",     // speed rspeed ping rping
		Concurrency:   2,
		TestMode:      2, // 2: ALLTEST 3: RETEST
		Subscription:  *link,
		Language:      "en", // en cn
		FontSize:      24,
		Theme:         "rainbow",
		Timeout:       10 * time.Second,
		OutputMode:    0, // 0: base64 1:file path 2: no pic 3: json 4: txt
	}
	ctx := context.Background()

	res, _ := web.TestContext(ctx, opts, &web.EmptyMessageWriter{})

	for _, node := range res {

		if node.IsOk {
			rs.SuccessNodes = append(rs.SuccessNodes, node.Link)
		} else {
			rs.FailedNodes = append(rs.FailedNodes, node.Link)
		}

		for key, value := range sm {

			if !includes(value, node.Link) {
				continue
			}

			if node.IsOk {

				if !includes(rs.Success, key) {
					rs.Success = append(rs.Success, key)
				}

			} else {
				if !includes(rs.Failed, key) {
					rs.Failed = append(rs.Failed, key)
				}
			}

		}

	}
	return rs
}

// {name:[]}
func PingAll(data Data, max int) Results {

	res := Results{}
	nodes := []string{}

	c := map[string][]string{}

	for _, item := range data.List {

		temp := item.Nodes
		if max > 0 && len(item.Nodes) > max {
			temp = item.Nodes[0:max]
		}

		c[item.Name] = temp

		res.Services = append(res.Services, item.Name)
		if len(temp) == 0 {
			res.Failed = append(res.Failed, item.Name)
		}

		nodes = append(nodes, temp...)
	}

	p := profile(strings.Join(nodes, "\n"), res, c)
	return p
}
