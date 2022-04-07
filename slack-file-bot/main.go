package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	//SET ENV

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3362450863105-3354027307109-XLXhksMjivx9w2bRLszp3Wg7")
	os.Setenv("CHANNEL_ID", "C03AND8S6M7")

	//Set Parameters To Upload a file
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"certificate Dad.pdf"} //File or Files to be uploaded

	//loop for in case of multiple files to upload
	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		fmt.Printf("Name: %s , URL: %s\n", file.Name, file.URL)
	}

}
