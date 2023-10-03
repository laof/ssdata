#### ssdata

install

```
go get -u github.com/laof/ssdata
```

```
data, err := Get("https://laof/fdate/ok")

if err != nil {
    return
}

mapdata := PingAll(data)
```
