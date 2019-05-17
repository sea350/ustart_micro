package elasticstore

var mapping = `
{
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
			"USER":{
				"properties":{
					"Username":{
						"type":"text",
						"analyzer":"my_analyzer",
						"fields":{
							"raw":{
								"type":"keyword"
							}
						}
					},
					"FirstName":{
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
					"University":{
						"type":"text",
						"analyzer":"my_analyzer",
						"fields":{
							"raw":{
								"type":"keyword"
							}
						}
					},
					"AcademicRecord": [
						"School":{
							"type":"text",
							"analyzer":"my_analyzer",
							"fields":{
								"raw":{
									"type":"keyword"
								}
							}
						},
						"Majors":{
							"type":"text",
							"analyzer":"my_analyzer",
							"fields":{
								"raw":{
									"type":"keyword"
								}
							}
						},
						"Minors":{
							"type":"text",
							"analyzer":"my_analyzer",
							"fields":{
								"raw":{
									"type":"keyword"
								}
							}
						},
						"Graduation":{
							"type":"text",
							"analyzer":"my_analyzer",
							"fields":{
								"raw":{
									"type":"keyword"
								}
							}
						},
						"EducationLevel":{
							"type":"text",
							"analyzer":"my_analyzer",
							"fields":{
								"raw":{
									"type":"keyword"
								}
							}
						}
					]
				}
			}
		}
	}
}
`
