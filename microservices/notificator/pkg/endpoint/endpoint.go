package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/hemuku90/microservices/notificator/pkg/service"
)

// SendMailRequest collects the request parameters for the SendMail method.
type SendMailRequest struct {
	Email   string `json:"email"`
	Content string `json:"content"`
}

// SendMailResponse collects the response parameters for the SendMail method.
type SendMailResponse struct {
	E0 error `json:"e0"`
}

// MakeSendMailEndpoint returns an endpoint that invokes SendMail on the service.
func MakeSendMailEndpoint(s service.NotificatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendMailRequest)
		e0 := s.SendMail(ctx, req.Email, req.Content)
		return SendMailResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r SendMailResponse) Failed() error {
	return r.E0
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// SendMail implements Service. Primarily useful in a client.
func (e Endpoints) SendMail(ctx context.Context, email string, content string) (e0 error) {
	request := SendMailRequest{
		Content: content,
		Email:   email,
	}
	response, err := e.SendMailEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendMailResponse).E0
}
