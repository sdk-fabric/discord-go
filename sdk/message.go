
// message automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


type Message struct {
    Content string `json:"content"`
    Nonce string `json:"nonce"`
    Tts bool `json:"tts"`
    Embeds []MessageEmbed `json:"embeds"`
    AllowedMentions *MessageAllowedMentions `json:"allowed_mentions"`
    MessageReference string `json:"message_reference"`
    Flags int `json:"flags"`
    EnforceNonce bool `json:"enforce_nonce"`
}

