package usr

//Syncdata data to be stored
type Syncdata struct {
	S3path   string `bson:"s3path"`
	Username string `bson:"username"`
	Password string `bson:"password"`
}
