package model

type Message struct {
	Type     string
	FileName string
}

type File struct {
	Filename string   `bson:"_id"`
	Header   string   `bson:"header"`
	Data     []string `bson:"data"`
}
