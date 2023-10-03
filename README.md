#### ssdata


```
data, err := Get("https://laof/fdate/ok")

if err != nil {
    return
}

mapdata := PingAll(data, true)
```
