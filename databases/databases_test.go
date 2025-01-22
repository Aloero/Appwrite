package databases

import (
	"fmt"
	"testing"

	// "github.com/Aloero/Appwrite/query"
	// "github.com/Aloero/Appwrite/models"
	"github.com/Aloero/Appwrite/client"
	// "github.com/mitchellh/mapstructure"
	// "github.com/Aloero/Appwrite/appwrite"

	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
    "math/rand"
)

func TestListDocuments(t *testing.T) {
	err := godotenv.Load("./../../.env")
	if err != nil {
		fmt.Println("Ошибка при загрузке .env файла")
		return
	}

	// userID, _ := strconv.ParseInt(os.Getenv("USERID"), 10, 64)

	databasesFast := NewDatabasesFast(os.Getenv("URL_ENDPOINT"), os.Getenv("PROJECT_ID"), os.Getenv("SECRET_API_KEY"), time.Second*60*10)
    results, err := databasesFast.ListDocuments(
		"62b87a794b6c790f658a",
		"62b87a794b6c73280cc7",
		[]string{},
		// []string{query.Equal("userid", 434)},
	)
	if err != nil {
		t.Errorf("Error with res: %v", err)
	}

	t.Logf("Number of documents: %d", len(results.Documents))
	
	// var userStatesList []UserStatesList
	// err = results.Decode(&userStatesList)
	// if err != nil {
	// 	t.Logf("err: %v", err)
	// }
	
	// // t.Logf("Number of UserStates: %d", userStatesList.UserStates[0].UserID)
	// t.Logf("UserStates: %d", len(userStatesList))
}

// func TestCreateDocuments(t *testing.T) {
// 	err := godotenv.Load("./../../.env")
// 	if err != nil {
// 		fmt.Println("Ошибка при загрузке .env файла")
// 		return
// 	}

// 	client := NewClient(
// 		WithProject(os.Getenv("PROJECT_ID")),
// 		WithKey(os.Getenv("SECRET_API_KEY")),
// 		WithEndpoint(os.Getenv("URL_ENDPOINT")),
// 	)

// 	databasesSlow := NewDatabases(client)

// 	var mapStruct map[string]interface{}
// 	err = mapstructure.Decode(UserState{UserID: 434, ChannelName: "656"}, &mapStruct)
// 	if err != nil {
// 		t.Logf("%v", err)
// 	}
	
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)
// 	go worker(databasesSlow, mapStruct)

// 	chainExit := make(chan struct{})
// 	<- chainExit
// }

func worker(databasesSlow *Databases, mapStruct map[string]interface{}) {
	idx := 0
	for {
		databasesSlow.CreateDocument(
			"62b87a794b6c790f658a",
			"62b87a794b6c73280cc7",
			Unique(),
			mapStruct,
		)

		idx += 1
		fmt.Println(idx)
	}
}



func Unique() string {
	timestamp := time.Now().UnixNano() / 1000

	choices := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	hexString := strconv.FormatInt(timestamp, 16)

	for i := 0; i < 7; i++ {
		hexString += choices[rand.Intn(len(choices))]
	}
	return hexString
}


func NewDatabases(clt client.Client) *Databases {
	return New(clt)
}

func WithEndpoint(endpoint string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Endpoint = endpoint
		return nil
	}
}

func WithKey(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Headers["X-Appwrite-Key"] = value
		return nil
	}
}

func WithProject(value string) client.ClientOption {
	return func(clt *client.Client) error {
		clt.Headers["X-Appwrite-Project"] = value
		return nil
	}
}

func NewClient(optionalSetters ...client.ClientOption) client.Client {
	return client.New(optionalSetters...)
}

type UserStatesList struct {
	// *models.DocumentList
	UserStates []UserState `json:"documents"`
}

type UserState struct {
	UserID   int64  `mapstructure:"userid" json:"userid"`
	Username string `mapstructure:"username" json:"username"`

	WaitingInputSub     bool `mapstructure:"waitinginputsub" json:"waitinginputsub"`
	WaitingInputChannel bool `mapstructure:"waitinginputchannel" json:"waitinginputchannel"`

	ChannelName              string  `mapstructure:"channelname" json:"channelname"`
	NumSubs                  int64   `mapstructure:"numsubs" json:"numsubs"`
	NumRubs                  float64 `mapstructure:"numrubs" json:"numrubs"`
	NumUSDT                  float64 `mapstructure:"numusdt" json:"numusdt"`
	NumRubs_intForDeclension int64   `mapstructure:"numrubs_intfordeclension" json:"numrubs_intfordeclension"`

	ADMINInputUSDT bool `mapstructure:"admininputusdt" json:"admininputusdt"`
	ADMINGetUSDT   bool `mapstructure:"admingetusdt" json:"admingetusdt"`

	DeclensionSubs int64 `mapstructure:"declensionsubs" json:"declensionsubs"`
	DeclensionRubs int64 `mapstructure:"declensionrubs" json:"declensionrubs"`

	LastMessageIsPhoto bool `mapstructure:"lastmessageisphoto" json:"lastmessageisphoto"`
}