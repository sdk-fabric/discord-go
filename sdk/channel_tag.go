
// ChannelTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

import (
    
    
    
    "github.com/apioo/sdkgen-go"
    
    "net/http"
    
    
)

type ChannelTag struct {
    internal *sdkgen.TagAbstract
}

func (client *ChannelTag) Message() *ChannelMessageTag {
    return NewChannelMessageTag(client.internal.HttpClient, client.internal.Parser)
}




func NewChannelTag(httpClient *http.Client, parser *sdkgen.Parser) *ChannelTag {
	return &ChannelTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
