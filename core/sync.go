package core

import (
	"github.com/tiaguinho/esmsync/es"
	"reflect"
)

//sync data between mongo and elasticsearch
func sync(oplogs interface{}) int64 {
	length := reflect.ValueOf(oplogs).Len()

	var total int64
	for i := 0; i < length; i++ {
		esdata := es.Mapping(reflect.ValueOf(oplogs).Index(i).Interface())

		if len(esdata.Data) > 0 {
			elastic.Execute(esdata)

			total++
		}
	}

	return total
}