package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*

<xliff version="1.2">
 <file original="Graphic Example.psd"
  source-language="en-US" target-language="ja-JP"
  tool="Rainbow" datatype="photoshop">
  <body>
   <trans-unit id="1" maxbytes="14">
    <source xml:lang="en-US">Quetzal</source>
    <target xml:lang="ja-JP">Quetzal</target>
   </trans-unit>
   <trans-unit id="3" maxbytes="114">
    <source xml:lang="en-US">An application to manipulate and
     process XLIFF documents</source>
    <target xml:lang="ja-JP">XLIFF 文書を編集、または処理
     するアプリケーションです。</target>
   </trans-unit>
   <trans-unit id="4" maxbytes="36">
    <source xml:lang="en-US">XLIFF Data Manager</source>
    <target xml:lang="ja-JP">XLIFF データ・マネージャ</target>
   </trans-unit>
  </body>
 </file>
</xliff>

*/

type Xliff struct {
	XMLName xml.Name `xml:"xliff"`
	File    File     `xml:"file"`
}
type File struct {
	File xml.Name `xml:"file"`
	Body Body     `xml:"body"`
}

type Body struct {
	Body       xml.Name    `xml:"body"`
	TransUnits []TransUnit `xml:"trans-unit"`
}

type TransUnit struct {
	TransUnit xml.Name `xml:"trans-unit"`
	Source    string   `xml:"source"`
	Target    string   `xml:"target"`
	Id        string   `xml:"id,attr"`
}

func main() {
	// Open our xmlFile
	xmlFile, err := os.Open("a.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened xml file")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var xliff Xliff
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &xliff)

	fmt.Printf("%v", xliff)
	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(xliff.File.Body.TransUnits); i++ {
		fmt.Println("Target: " + xliff.File.Body.TransUnits[i].Target)
		fmt.Println("id: " + xliff.File.Body.TransUnits[i].Id)
		// fmt.Println("User Name: " + users.Users[i].Name)
		// fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}

	xliff.File.Body.TransUnits[0].Target = "kyle is great"
	filename := "new_messages.xml"
	file, _ := os.Create(filename)

	xmlWriter := io.Writer(file)

	encoder := xml.NewEncoder(xmlWriter)
	encoder.Indent(" ", "    ")
	if err := encoder.Encode(xliff); err != nil {
		fmt.Printf("error : %v\n", err)
	}
	encoder.Flush()
	fmt.Printf("Write operation to %s completed\n", filename)
}
