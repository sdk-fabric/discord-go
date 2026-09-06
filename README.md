
# discord-go

This [SDK](https://github.com/sdk-fabric/discord-go) is managed by the [SDK Fabric](https://sdk-fabric.org/) project, a global infrastructure to
automatically generate SDKs for every API.

You can find more information about this SDK at [TypeHub](https://typehub.cloud/):
https://app.typehub.cloud/d/sdkfabric/discord

## Usage

```go
import (
	"github.com/sdk-fabric/discord-go/sdk"
)

var client, _ = sdk.Build("[access_token]");

// Get a channel by ID.
response, err := client.Channel().Get("channel_id")

// Update a channel's settings.
response, err := client.Channel().Update("channel_id", ChannelUpdate{})

// Delete a channel, or close a private message.
response, err := client.Channel().Delete("channel_id")

// Returns all pinned messages in the channel as an array of message objects.
response, err := client.Channel().Getpins("channel_id")

// Create a new invite object for the channel.
response, err := client.Channel().Createinvite("channel_id", ChannelInvite{})

// Retrieves the messages in a channel.
response, err := client.Message().Getall("channel_id", "around", "before", "after", 1)

// Retrieves a specific message in the channel.
response, err := client.Message().Get("channel_id", "message_id")

// Post a message to a guild text or DM channel.
response, err := client.Message().Create("channel_id", Message{})

// Edit a previously sent message.
response, err := client.Message().Update("channel_id", "message_id", Message{})

// Delete a message.
response, err := client.Message().Remove("channel_id", "message_id")

// Crosspost a message in an Announcement Channel to following channels.
response, err := client.Message().Crosspost("channel_id", "message_id")

response, err := client.Message().Getreactionsbyemoji("channel_id", "message_id", "emoji", 1, "after", 1)

response, err := client.Message().Deleteallreactions("channel_id", "message_id")

// Returns the user object of the requester's account.
response, err := client.User().Getcurrent()

// Returns a user object for a given user ID.
response, err := client.User().Get("user_id")
```
