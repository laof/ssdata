package ssdata

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type List struct {
	Name     string   `json:"name"`
	Nodes    []string `json:"-"`
	Datetime string   `json:"datetime"`
	Length   int      `json:"length"`
	Data     string   `json:"data"`
}

type Data struct {
	List   []List     `json:"list"`
	Decode [][]string `json:"decode"`
	Update string     `json:"update"`
}

func Get(url string) (Data, error) {
	data := Data{}
	res, err := http.Get(url)
	if err != nil {
		return data, err
	}
	defer res.Body.Close()
	str, err := io.ReadAll(res.Body)

	if err != nil {
		return data, err
	}

	err = json.Unmarshal(str, &data)

	if err != nil {
		return data, err
	}

	for i, item := range data.List {
		txt := item.Data

		if item.Data == "" {
			continue
		}

		for _, arr := range data.Decode {
			txt = strings.ReplaceAll(txt, arr[1], arr[0])
		}

		data.List[i].Name = reverseString(item.Name)
		data.List[i].Datetime = reverseString(item.Datetime)
		data.List[i].Data = txt
		data.List[i].Nodes = strings.Split(txt, ",")
	}
	return data, nil
}

func encoding(arr []List) []List {

	for i, node := range arr {
		if len(node.Nodes) > 0 {
			text := strings.Join(node.Nodes, ",")
			for _, arr := range codemap {
				text = strings.ReplaceAll(text, arr[0], arr[1])
			}
			arr[i].Data = text
			arr[i].Length = len(node.Nodes)
			arr[i].Nodes = []string{}
		}
		arr[i].Name = reverseString(node.Name)
		arr[i].Datetime = reverseString(node.Datetime)

	}
	return arr
}

func GetDataString(list []List) (string, error) {

	data := Data{
		List:   encoding(list),
		Decode: codemap,
		Update: now(),
	}

	str, err := json.Marshal(data)

	if err != nil {
		return "", err
	}

	return string(str), nil
}
