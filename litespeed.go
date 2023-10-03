package ssdata

import (
	"context"
	"flag"
	"log"
	"strings"
	"time"

	"github.com/laof/lite-speed-test/web"
)

func profile(url string, filterBySuccess bool) []string {

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
	nodes := []string{}
	for _, node := range res {

		if (filterBySuccess && node.IsOk) || (!filterBySuccess && !node.IsOk) {
			nodes = append(nodes, node.Link)
		}

	}
	return nodes
}

// {name:[]}
func PingAll(data Data, filterBySuccess bool, max int) (res map[string][]string) {

	nodes := []string{}
	res = map[string][]string{}
	for _, item := range data.List {

		temp := item.Nodes
		if max > 0 && len(item.Nodes) > max {
			temp = item.Nodes[0:max]
		}
		nodes = append(nodes, temp...)
	}
	list := profile(strings.Join(nodes, "\n"), filterBySuccess)

	for _, item := range data.List {

		for _, link := range list {

			if includes(item.Nodes, link) {
				res[item.Name] = append(res[item.Name], link)
			}
		}

	}
	return
}
