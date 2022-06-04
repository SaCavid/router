package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"regexp"
	"strings"
)

type Urls struct {
	Path  string
	Regex string
	Vars  map[string]string
}

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	var arr []Urls
	registeredPath := []string{"/path/no variable", "/path/{id}", "/path/{id}/{status}", "/path/{status}/articles/{year}"}

	for _, v := range registeredPath {
		url := Urls{
			Path:  v,
			Regex: "",
			Vars:  make(map[string]string),
		}

		pathArr := strings.Split(v, "/")
		pathArr = pathArr[1:]
		regexString := "^"
		for _, val := range pathArr {
			matched, _ := regexp.MatchString("{([0-9A-Za-z_-]+)}", val)
			if matched {
				regexString += "/([0-9A-Za-z_-]+)"
				val = strings.Replace(val, "{", "", -1)
				val = strings.Replace(val, "}", "", -1)
				url.Vars[val] = ""
			} else {
				regexString += "/" + val
			}
		}
		regexString += "$"
		url.Regex = regexString
		arr = append(arr, url)
	}

	realPath := []string{"/path/no variable", "/path/12", "/path/15/delete", "/path/confirmed/articles/2022"}
	for _, value := range realPath {
		log.Print("888888888888888888888888")
		for _, v := range arr {
			matched, _ := regexp.MatchString(v.Regex, value)
			if matched {
				pathArr := strings.Split(v.Path, "/")
				pathArr = pathArr[1:]
				realArr := strings.Split(value, "/")
				realArr = realArr[1:]
				for key, val := range pathArr {
					matched, _ := regexp.MatchString("{([0-9A-Za-z_-]+)}", val)
					if matched {
						val = strings.Replace(val, "{", "", -1)
						val = strings.Replace(val, "}", "", -1)
						if _, ok := v.Vars[val]; ok {
							v.Vars[val] = realArr[key]
						} else {
							log.Fatal().Msg("Key not exists. Router Error")
						}
						log.Print("---", val, " ", realArr[key])
					}
				}
			}
		}
	}

	for k := range arr {
		log.Print(arr[k])
	}

}
