package constant

import "wow-admin/utils/cast"

type Dictionary []map[string]interface{}

func (d Dictionary) RangeIntKeyValue() Dictionary  {
	dics := make(Dictionary, 0)
	for _, mp := range d {
		for k, v := range mp {
			dics = append(dics, map[string]interface{}{
				"value": cast.Int(k),
				"text": v,
				"select": false,
			})
		}
	}
	return dics
}

var Language = Dictionary{
	{"1": "Vue"},
	{"2": "React"},
	{"1": "angular"},
}