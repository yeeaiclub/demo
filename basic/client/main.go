package main

import (
	"fmt"
	"net/http"

	"github.com/yeeaiclub/a2a-go/sdk/client"
	"github.com/yeeaiclub/a2a-go/sdk/types"
)

func main() {
	httpClient := &http.Client{}
	r := client.NewA2ACardResolver(httpClient, "http://localhost:8080", client.WithAgentCardPath("card"))
	card, err := r.GetAgentCard()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(card)

	newClient := client.NewClient(httpClient, "http://localhost:8080/api")
	resp, err := newClient.SendMessage(types.MessageSendParam{
		Message: &types.Message{
			TaskID: "1",
			Role:   types.User,
			Parts: []types.Part{
				&types.TextPart{Kind: "text", Text: "hello, world"},
			},
		},
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.Error != nil {
		fmt.Println(resp.Error.Message)
		return
	}

	task, err := types.MapTo[types.Task](resp.Result)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(task.Id)
	fmt.Println(task.ContextId)
}
