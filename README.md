Keywords marking. 

```
marking ./keywords.txt ./structure.json ./result.csv
```

# keywords.txt

List of keywords.

```
word word word
word word 
word word word word
```

# structure.json

JSON file with structure for marking.

```
{
	"columns": [
		{
			"name": "", // сolumn name
			"type": "", // column type: tree or tags
			"objects": [
				{
					"name": "", // object name
					"patterns": [], // patterns for searching ("patterns": [["pattern" AND "pattern_2" AND ...] OR ["pattern" AND "pattern_2" AND ...] OR ...])
					"sub": [] // child objects
				},
			]
		},
	]
}
```

# result.csv

The result in csv format.

___

See the folder "./example" for examples