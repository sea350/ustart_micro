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
				"CustomURL":{
					"type":"text",
					"analyzer":"my_analyzer",
					"fields":{
						"raw":{
							"type":"keyword"
						}
					}
				},
				"Name":{
					"type": "text",
					"analyzer":"my_analyzer",
					"fields":{
						"raw":{
							"type":"keyword"
						}
					}
				},
				"LastName":{
					"type":"text",
					"analyzer":"my_analyzer",
					"fields":{
						"raw":{
							"type":"keyword"
						}
					}
				},
				"Tags":{
					"type":"text",
					"analyzer":"my_analyzer",
					"fields":{
						"raw":{
							"type":"keyword"
						}
					}
				},
				"School":{
					"type":"text",
					"analyzer":"my_analyzer",
					"fields":{
						"raw":{
							"type":"keyword"
						}
					}
				},
				
				"SkillsNeeded":{
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
