package elasticstore

func mapping(indexName string) string {

	return `{
	"settings": {
		"analysis": {
			"analyzer": {
				"my_analyzer": {
					"type": "custom",
					"filter": [
						"lowercase"
					],
					"tokenizer": "whitespace"
				}
			}
		}
	},
	"mappings":{
		"properties":{
			"UUID":{
				"type": "text",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			},
			"ActivityID":{
				"type":"text",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			},
			"Seen":{
				"type":"text",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			},
			"TimeSeen":{
				"type":"text",
				"analyzer":"my_analyzer",
				"fields":{
					"raw":{
						"type":"keyword"
					}
				}
			}
		}
	}	
}
`
}
