package mutate

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/qiniu/logkit/utils/models"
)

func TestXmlTransformer(t *testing.T) {

	xtr := Xml{
		Key: "xml",
		New: "xml",
	}
	data1 := []Data{
		{
			"key1": "value1",
			"xml": `<?xml version="1.0" encoding="UTF-8"?>
		<note>
		 <to>Tove</to>
		 <from>Jani</from>
		 <heading>Reminder</heading>
		 <body>Don't forget me this weekend!</body>
		</note>`,
		},
	}
	newdata, err := xtr.Transform(data1)
	assert.NoError(t, err)
	expdata := []Data{
		{
			"key1": "value1",
			"xml": map[string]interface{}{
				"note": map[string]interface{}{
					"heading": "Reminder",
					"body":    "Don't forget me this weekend!",
					"to":      "Tove",
					"from":    "Jani",
				},
			},
		},
	}
	assert.Equal(t, expdata, newdata)

	xtr12 := Xml{
		Key: "xml",
	}
	data12 := []Data{
		{
			"key1": "value1",
			"xml": `<?xml version="1.0" encoding="UTF-8"?>
		<note>
		 <to>Tove</to>
		 <from>Jani</from>
		 <heading>Reminder</heading>
		 <body>Don't forget me this weekend!</body>
		</note>`,
		},
	}
	newdata12, err12 := xtr12.Transform(data12)
	assert.NoError(t, err12)
	expdata12 := []Data{
		{
			"key1": "value1",
			"xml": map[string]interface{}{
				"note": map[string]interface{}{
					"heading": "Reminder",
					"body":    "Don't forget me this weekend!",
					"to":      "Tove",
					"from":    "Jani",
				},
			},
		},
	}
	assert.Equal(t, expdata12, newdata12)

	xtr13 := Xml{
		Key: "xml",
		New: "xml2",
	}
	data13 := []Data{
		{
			"key1": "value1",
			"xml": `<?xml version="1.0" encoding="UTF-8"?>
		<note>
		 <to>Tove</to>
		 <from>Jani</from>
		 <heading>Reminder</heading>
		 <body>Don't forget me this weekend!</body>
		</note>`,
		},
	}
	newdata13, err13 := xtr13.Transform(data13)
	assert.NoError(t, err13)
	expdata13 := []Data{
		{
			"key1": "value1",
			"xml2": map[string]interface{}{
				"note": map[string]interface{}{
					"heading": "Reminder",
					"body":    "Don't forget me this weekend!",
					"to":      "Tove",
					"from":    "Jani",
				},
			},
			"xml": `<?xml version="1.0" encoding="UTF-8"?>
		<note>
		 <to>Tove</to>
		 <from>Jani</from>
		 <heading>Reminder</heading>
		 <body>Don't forget me this weekend!</body>
		</note>`,
		},
	}
	assert.Equal(t, expdata13, newdata13)

	xtr14 := Xml{
		Key:        "xml",
		New:        "xml2",
		DiscardKey: true,
	}
	data14 := []Data{
		{
			"key1": "value1",
			"xml": `<?xml version="1.0" encoding="UTF-8"?>
		<note>
		 <to>Tove</to>
		 <from>Jani</from>
		 <heading>Reminder</heading>
		 <body>Don't forget me this weekend!</body>
		</note>`,
		},
	}
	newdata14, err14 := xtr14.Transform(data14)
	assert.NoError(t, err14)
	expdata14 := []Data{
		{
			"key1": "value1",
			"xml2": map[string]interface{}{
				"note": map[string]interface{}{
					"heading": "Reminder",
					"body":    "Don't forget me this weekend!",
					"to":      "Tove",
					"from":    "Jani",
				},
			},
		},
	}
	assert.Equal(t, expdata14, newdata14)

	xtr3 := Xml{
		Key:    "xml",
		Expand: true,
	}
	data3 := []Data{
		{
			"key1": "value1",
			"xml": `<?xml version="1.0" encoding="UTF-8"?>
		<note>
		 <to>Tove</to>
		 <from>Jani</from>
		 <heading>Reminder</heading>
		 <body>Don't forget me this weekend!</body>
		</note>`,
		},
	}
	newdata3, err3 := xtr3.Transform(data3)
	assert.NoError(t, err3)
	expdata3 := []Data{
		{"body": "Don't forget me this weekend!", "key1": "value1", "to": "Tove", "from": "Jani", "heading": "Reminder", "xml": `<?xml version="1.0" encoding="UTF-8"?>
		<note>
		 <to>Tove</to>
		 <from>Jani</from>
		 <heading>Reminder</heading>
		 <body>Don't forget me this weekend!</body>
		</note>`},
	}
	assert.Equal(t, expdata3, newdata3)

	xtr4 := Xml{
		Key:    "xml.xml2",
		New:    "xml2",
		Expand: true,
	}
	data4 := []Data{
		{
			"key1": "value1",
			"xml": map[string]interface{}{"xml2": `<?xml version="1.0" encoding="UTF-8"?>
		<note>
		 <to>Tove</to>
		 <from>Jani</from>
		 <heading>Reminder</heading>
		 <body>Don't forget me this weekend!</body>
		</note>`},
		},
	}
	newdata4, err4 := xtr4.Transform(data4)
	assert.NoError(t, err4)
	expdata4 := []Data{
		{
			"key1": "value1",
			"xml": map[string]interface{}{"xml2": `<?xml version="1.0" encoding="UTF-8"?>
		<note>
		 <to>Tove</to>
		 <from>Jani</from>
		 <heading>Reminder</heading>
		 <body>Don't forget me this weekend!</body>
		</note>`},
			"xml2": map[string]interface{}{
				"body": "Don't forget me this weekend!",
				"to":   "Tove", "from": "Jani",
				"heading": "Reminder",
			},
		},
	}
	assert.Equal(t, expdata4, newdata4)

	xtr5 := Xml{
		Key:        "raw",
		Expand:     true,
		DiscardKey: true,
	}
	data5 := []Data{
		{
			"key1": "value1",
			"raw": `<?xml version="1.0" encoding="UTF-8"?>
	<note>
	<to>Tove</to>
	<from>Jani</from>
	<heading>Reminder</heading>
	<body>Don't forget me this weekend!</body>
	</note>`,
		},
		{
			"key1": "value1",
			"raw": `<?xml version="1.0" encoding="UTF-8"?>
	<note>
	<to>Tove</to>
	<from>Jani</from>
	<heading>Reminder</heading>
	<body>Don't forget me this weekend!</body>
	</note>`,
		},
	}
	newdata5, err5 := xtr5.Transform(data5)
	assert.NoError(t, err5)
	expdata5 := []Data{
		{
			"key1": "value1",
			"body": "Don't forget me this weekend!",
			"to":   "Tove", "from": "Jani",
			"heading": "Reminder",
		},
		{
			"key1": "value1",
			"body": "Don't forget me this weekend!",
			"to":   "Tove", "from": "Jani",
			"heading": "Reminder",
		},
	}
	assert.Equal(t, expdata5, newdata5)
}
