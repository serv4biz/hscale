package locals

import (
	"fmt"
	"strings"

	"github.com/SERV4BIZ/gfp/files"
	"github.com/SERV4BIZ/gfp/handler"
	"github.com/SERV4BIZ/gfp/jsons"
	"github.com/SERV4BIZ/hscale/config/utility"
)

// ListDatabase is listing database all
func ListDatabase() (*jsons.JSONArray, error) {
	jsaList := jsons.JSONArrayFactory()
	pathdir := fmt.Sprint(utility.GetAppDir(), utility.DS, "databases")
	filelist, err := files.ScanDir(pathdir)
	if handler.Error(err) {
		return nil, err
	}

	for _, fileName := range filelist {
		if strings.HasSuffix(fileName, ".json") {
			jsaList.PutString(strings.TrimSpace(strings.ToLower(strings.TrimSuffix(fileName, ".json"))))
		}
	}
	return jsaList, err
}