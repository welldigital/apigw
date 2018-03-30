package apigw

import "github.com/aws/aws-lambda-go/events"

// PagedResponse is the return type for lists.
type PagedResponse struct {
	StartIndex int         `json:"startIndex"`
	PageSize   int         `json:"pageSize"`
	NextIndex  *int        `json:"nextIndex,omitempty"`
	Data       interface{} `json:"data"`
}

// Paged returns paged data.
func Paged(startIndex, pageSize, nextIndex int, data interface{}, status int) (r events.APIGatewayProxyResponse, err error) {
	l := PagedResponse{
		StartIndex: startIndex,
		PageSize:   pageSize,
		Data:       data,
	}
	if nextIndex > 0 {
		l.NextIndex = &nextIndex
	}
	return JSON(l, status)
}
