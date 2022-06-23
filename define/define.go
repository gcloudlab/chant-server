package define

import "os"

var MailPassword = os.Getenv("MailPassword")
var MongoPassword = os.Getenv("MongoPassword")
