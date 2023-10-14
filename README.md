# AWS Lambda Function in Golang to Search Page in Confluence

- Trigger by AWS Chatbot in Slack channel
- Provide the query string with keyword
- This Lambda will get all labels from Confluence space
- This Lambda will send the query string and lables to Azure GPT
- Azure GPT will return the most relevant label
- This Lambda will Search the page in Confluence with the label
- (just for Hackathon DEMO purpose)
