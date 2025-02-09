package databases

import (
	// "encoding/json"
	"fmt"
	"testing"

	// "github.com/Aloero/Appwrite/query"
	// "github.com/Aloero/Appwrite/models"
	"github.com/Aloero/Appwrite/client"
	// "github.com/mitchellh/mapstructure"
	// "github.com/Aloero/Appwrite/appwrite"

	"github.com/joho/godotenv"
	"os"
	"math/rand"
	"strconv"
	"time"
	// "log"
)

func TestListDocuments(t *testing.T) {

	// a := `{"total":5000,"documents":[{"waitinginputsub":false,"numusdt":0,"declensionsubs":0,"waitinginputchannel":false,"numrubs":0,"userid":434,"username":"","channelname":"","numsubs":0,"numrubs_intfordeclension":0,"admininputusdt":false,"admingetusdt":false,"declensionrubs":0,"lastmessageisphoto":false,"$id":"62c4e54ce78e6ff21106","$createdAt":"2025-01-22T16:53:22.112+00:00","$updatedAt":"2025-01-22T16:53:22.112+00:00","$permissions":[],"$databaseId":"62b87a794b6c790f658a","$collectionId":"62b87a794b6c73280cc7"},{"waitinginputsub":false,"numusdt":0,"declensionsubs":0,"waitinginputchannel":false,"numrubs":0,"userid":434,"username":"","channelname":"","numsubs":0,"numrubs_intfordeclension":0,"admininputusdt":false,"admingetusdt":false,"declensionrubs":0,"lastmessageisphoto":false,"$id":"62c4e587a0a35c072a60","$createdAt":"2025-01-22T16:54:23.692+00:00","$updatedAt":"2025-01-22T16:54:23.692+00:00","$permissions":[],"$databaseId":"62b87a794b6c790f658a","$collectionId":"62b87a794b6c73280cc7"}]}`
	// b := `{"total":5000,"documents":[{"waitinginputsub":false,"numusdt":0,"declensionsubs":0,"waitinginputchannel":false,"numrubs":0,"userid":434,"username":"","channelname":"","numsubs":0,"numrubs_intfordeclension":0,"admininputusdt":false,"admingetusdt":false,"declensionrubs":0,"lastmessageisphoto":false,"$id":"62c4e54ce78e6ff21106","$createdAt":"2025-01-22T16:53:22.112+00:00","$updatedAt":"2025-01-22T16:53:22.112+00:00","$permissions":[],"$databaseId":"62b87a794b6c790f658a","$collectionId":"62b87a794b6c73280cc7"},{"waitinginputsub":false,"numusdt":0,"declensionsubs":0,"waitinginputchannel":false,"numrubs":0,"userid":434,"username":"","channelname":"","numsubs":0,"numrubs_intfordeclension":0,"admininputusdt":false,"admingetusdt":false,"declensionrubs":0,"lastmessageisphoto":false,"$id":"62c4e587a0a35c072a60","$createdAt":"2025-01-22T16:54:23.692+00:00","$updatedAt":"2025-01-22T16:54:23.692+00:00","$permissions":[],"$databaseId":"62b87a794b6c790f658a","$collectionId":"62b87a794b6c73280cc7"}]}`
	// bytesa := []byte(a)
	// bytesb := []byte(b)

	// type DocumentUnpackTest struct {
	// 	Total int64 `json:"total"`
	// 	Documents json.RawMessage `json:"documents"`
	// }

	// type TestStruct struct {
	// 	Total int64   `json:"total"`
	// 	Documents []interface{}   `json:"documents"`
	// }
	// type TestStruct struct {
	// 	Waitinginputsub bool   `json:"waitinginputsub"`
	// 	Numusdt int64   `json:"numusdt"`
	// 	Id string   `json:"id"`
	// }

	// var sa DocumentUnpackTest
	// err := json.Unmarshal(bytesa, &sa)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var sb DocumentUnpackTest
	// err = json.Unmarshal(bytesb, &sb)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// alls := sliseEndsStr(string(sa.Documents)) + `,` + sliseEndsStr(string(sb.Documents))
	// alls = addEndsStr(alls)

	// fmt.Println(alls)

	// var userState []UserState
	// err = json.Unmarshal([]byte(alls), &userState)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// s2 := string(s.Documents)
	// s2 = s2[1 : len(s2)-1]
	// fmt.Println(s2)

	// fmt.Println(userState)
	
	// var m []interface{}

	// m = append(m, s.Documents...)

	// fmt.Println(s)
	// fmt.Println(m[0].(map[string]interface{})["$collectionId"])

	err := godotenv.Load("./../../.env")
	if err != nil {
		fmt.Println("Ошибка при загрузке .env файла")
		return
	}

	// // userID, _ := strconv.ParseInt(os.Getenv("USERID"), 10, 64)

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

	// fmt.Println(string(results.Data)[len(string(results.Data))-1000:])
	
	var userStates []UserState
	err = results.Decode(&userStates)
	if err != nil {
		t.Logf("err: %v", err)
	}
	// fmt.Println(userStates)
	fmt.Println(results.Total)
	fmt.Println(results.Total)

	// var userStates []UserState
	// for i := 0; i < len(userStatesList); i++ {
	// 	for j := 0; j < len(userStatesList[i].UserStates); j++ {
	// 		userStates = append(userStates, userStatesList[i].UserStates[j])
	// 	}
	// }
}

// func sliseEndsStr(s string) string {
// 	return s[1 : len(s)-1]
// }

// func addEndsStr(s string) string {
// 	return "[" + s + "]"
// }

// type UserStatesList struct {
// 	UserStates []UserState `json:"documents"`

// }

// 	t.Logf("Number of UserStates: %d", userStatesList.UserStates[0].UserID)
// 	t.Logf("All documents: %d", len(results.Documents))
// 	t.Logf("UserStates: %d", len(userStates))
// }

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

type UserState struct {
	DocumentDescribe
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