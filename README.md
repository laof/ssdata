#### ssdata


```
data, err := Get("https://laof.github.io/get-nodes-test-app/json/data.json")

if err != nil {
    return
}

mapdata := PingAll(data)

```
