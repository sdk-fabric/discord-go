
// Client automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

import (
    
    
    
    
    "github.com/apioo/sdkgen-go"
    
    
    
    
)

type Client struct {
    internal *sdkgen.ClientAbstract
}

func (client *Client) Channel() *ChannelTag {
    return NewChannelTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *Client) Message() *MessageTag {
    return NewMessageTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *Client) User() *UserTag {
    return NewUserTag(client.internal.HttpClient, client.internal.Parser)
}




func NewClient(baseUrl string, credentials sdkgen.CredentialsInterface) (*Client, error) {
	var client, err = sdkgen.NewClient(baseUrl, credentials)
    if err != nil {
        return &Client{}, err
    }

	return &Client{
		internal: client,
	}, nil
}

func Build(token string) (*Client, error) {
    var credentials = sdkgen.HttpBearer{Token: token}

    return NewClient("https://discord.com/api/v10", credentials)
}

func BuildAnonymous() (*Client, error) {
    var credentials = sdkgen.Anonymous{}

    return NewClient("https://discord.com/api/v10", credentials)
}
