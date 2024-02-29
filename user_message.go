package main

import "github.com/FPNL/google-bm-api-golang/businessmessages/v1"

// UserMessage Received from user
// https://developers.google.com/business-communications/business-messages/reference/rest/v1/UserMessage
type UserMessage struct {
	// ClientToken, Secret didn't contained in UserMessage. But it's convenience for json.Unmarshal when Agent verify webhook callback.
	// https://developers.google.com/business-communications/business-messages/guides/how-to/agents#webhook
	ClientToken string `json:"clientToken"`
	Secret      string `json:"secret"`

	RequestId      string `json:"requestId"`
	ConversationId string `json:"conversationId"`
	CustomAgentId  string `json:"customAgentId"`
	Agent          string `json:"agent"`
	Context        struct {
		EntryPoint string `json:"entryPoint"`
		UserInfo   struct {
			DisplayName      string `json:"displayName"`
			UserDeviceLocale string `json:"userDeviceLocale"`
		} `json:"userInfo"`
		Widget struct {
			URL           string `json:"url"`
			WidgetContext string `json:"widgetContext"`
		} `json:"widget"`
		ResolvedLocale string `json:"resolvedLocale"`
		CustomContext  string `json:"customContext"`
		PlaceId        string `json:"placeId"`
		NearPlaceId    string `json:"nearPlaceId"`
	} `json:"context"`
	SendTime           string `json:"sendTime"`
	DialogflowResponse struct {
		QueryText   string `json:"queryText"`
		FaqResponse struct {
			UserQuestion string `json:"userQuestion"`
			Answers      []struct {
				FaqQuestion          string  `json:"faqQuestion"`
				FaqAnswer            string  `json:"faqAnswer"`
				MatchConfidenceLevel string  `json:"matchConfidenceLevel"`
				MatchConfidence      float32 `json:"matchConfidence"`
			} `json:"answers"`
		} `json:"faqResponse"`
		IntentResponses []struct {
			IntentName                string  `json:"intentName"`
			IntentDisplayName         string  `json:"intentDisplayName"`
			IntentDetectionConfidence float32 `json:"intentDetectionConfidence"`
			fulfillmentMessages       []struct {
				error struct {
					Code    int                      `json:"code"`
					Message string                   `json:"message"`
					Details []map[string]interface{} `json:"details"`
				}
				Text             string `json:"text"`
				JsonPayload      string `json:"jsonPayload"`
				LiveAgentHandoff struct {
					Metadata map[string]interface{} `json:"metadata"`
				} `json:"liveAgentHandoff"`
			}
		}
		AutoResponded         bool `json:"autoResponded"`
		AutoRespondedMessages []struct {
			Message        businessmessages.BusinessMessagesMessage `json:"message"`
			ResponseSource string                                   `json:"responseSource"`
		} `json:"autoRespondedMessages"`
	} `json:"dialogflowResponse"`
	Name    string `json:"name"`
	Message struct {
		MessageId  string `json:"messageId"`
		Name       string `json:"name"`
		Text       string `json:"text"`
		CreateTime string `json:"createTime"`
	} `json:"message"`
	SuggestionResponse struct {
		Message      string `json:"message"`
		PostbackData string `json:"postbackData"`
		CreateTime   string `json:"createTime"`
		Text         string `json:"text"`
		Type         string `json:"type"`
	} `json:"suggestionResponse"`
	UserStatus struct {
		CreateTime         string `json:"createTime"`
		IsTyping           bool   `json:"isTyping"`
		RequestedLiveAgent bool   `json:"requestedLiveAgent"`
	} `json:"userStatus"`
}
