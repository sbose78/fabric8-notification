package template

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	authapi "github.com/fabric8-services/fabric8-notification/auth/api"
	witapi "github.com/fabric8-services/fabric8-notification/wit/api"
)

func formatDate(date interface{}) string {
	format := "02 January 2006"
	if d, ok := date.(*time.Time); ok {
		return d.Format(format)
	}
	if d, ok := date.(string); ok {
		p, err := time.Parse(time.RFC3339, d)
		if err != nil {
			return "Unknown"
		}
		return p.Format(format)
	}
	return "Unknown"
}

func lower(data interface{}) string {
	if data == nil {
		return ""
	}

	return strings.ToLower(resolveString(data))
}

func sizeImage(data interface{}, size int) string {
	url := resolveString(data)
	if strings.Contains(url, "?") {
		return fmt.Sprintf("%v&s=%v", url, size)
	}
	return fmt.Sprintf("%v?s=%v", url, size)
}

func resolveString(data interface{}) string {
	if data == nil {
		return ""
	}
	value := reflect.ValueOf(data)
	if value.Type().Kind() == reflect.Ptr {
		return fmt.Sprint(value.Elem())
	}
	return fmt.Sprint(data)
}

func detailURL(webURL string, spaceOwner *authapi.User, space witapi.SpaceSingle, workitem witapi.WorkItemSingle) string {
	return fmt.Sprintf(
		"%v/%v/%v/plan/detail/%v",
		webURL,
		*spaceOwner.Data.Attributes.Username,
		*space.Data.Attributes.Name,
		workitem.Data.Attributes["system.number"])
}

func areaPath(parentPath, name string) string {
	p := parentPath
	if p == "/" {
		p = ""
	}
	return fmt.Sprintf("%v/%v", lower(p), lower(name))
}
