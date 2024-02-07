
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




func NewClient(baseUrl string, credentials sdkgen.CredentialsInterface) (*Client, error) {
	var client, err = sdkgen.NewClient(baseUrl, credentials)
    if err != nil {
        return &Client{}, err
    }

	return &Client{
		internal: client,
	}, nil
}

func Build(clientId string, clientSecret string, tokenStore sdkgen.TokenStoreInterface, scopes []string) (*Client, error) {
    var credentials = sdkgen.ClientCredentials{
        ClientId: clientId,
        ClientSecret, clientSecret,
        TokenUrl: "https://discord.com/api/oauth2/token",
        AuthorizationUrl: "https://discord.com/oauth2/authorize",
        TokenStore: tokenStore,
        Scopes: scopes,
    }

    client, err := NewClient("https://discord.com/api/v10", credentials)
    if err != nil {
        return &Client{}, err
    }

    return &Client {
        internal: client,
    }, nil
}

