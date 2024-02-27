# Google Business Message API

# First of all

You should know your product in these potential situations
https://developers.google.com/business-communications/business-messages/guides/how-to/message/message-lifecycle

# Notice

## Google promises sending message within 7 days

Or until your webhook successfully receives the message.

## Not all devices accept all message styles

Not all messages can be sent to the user. Remember to provide fallback message for unsupported devices.
Fallback messages support `\n`, `\r\n`.

## HUMAN and BOT - Representatives
Conversation can be represented by `BOT`, then `HUMAN` can join any time.
BOT and HUMAN are called `Representatives`.

## Message Style

User can send 4 types of messages:
- Text
- Images
  - Google only keep 7 days
- Suggestions
- Authentication request

Representative can send 9 types of messages:
- Text
- Rich Text
- Images
- Suggested replies(!)
- Suggested actions(!)
- Authentication request suggestion
- Live agent request suggestion
- Rich cards
- Rich card carousels

## Events

- User request HUMAN
- Representative joined/left
- Typing 

# AI

Google provides [Dialogflow](https://cloud.google.com/dialogflow?hl=zh_tw)
