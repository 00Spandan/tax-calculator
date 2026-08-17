package taxclient

import (
	"context"
	"net/url"
	"strings"
	"time"

	taxv1 "github.com/00Spandan/financial-tools/gen/go/tax"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	grpcClient taxv1.TaxCalculatorClient
	conn       *grpc.ClientConn
}

type CalculateTaxParams struct {
	AnnualIncome float64
	TaxYear      string
	CountryCode  string
}

type CalculateTaxResult struct {
	TaxableIncome    float64
	TaxOwed          float64
	EffectiveTaxRate float64
}

func New(target string) (*Client, error) {
	connectionTarget := target
	var dialOptions []grpc.DialOption

	if strings.HasPrefix(target, "https://") || strings.HasPrefix(target, "http://") {
		parsedURL, err := url.Parse(target)
		if err != nil {
			return nil, err
		}
		connectionTarget = parsedURL.Host
		if parsedURL.Port() == "" {
			connectionTarget = parsedURL.Hostname() + ":443"
		}
		dialOptions = append(dialOptions, grpc.WithTransportCredentials(credentials.NewTLS(nil)))
	} else {
		dialOptions = append(dialOptions, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.NewClient(connectionTarget, dialOptions...)
	if err != nil {
		return nil, err
	}

	return &Client{
		grpcClient: taxv1.NewTaxCalculatorClient(conn),
		conn:       conn,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) CalculateTax(ctx context.Context, params CalculateTaxParams) (CalculateTaxResult, error) {
	requestContext, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	response, err := c.grpcClient.CalculateTax(requestContext, &taxv1.CalculateTaxRequest{
		AnnualIncome: params.AnnualIncome,
		TaxYear:      params.TaxYear,
		CountryCode:  params.CountryCode,
	})
	if err != nil {
		return CalculateTaxResult{}, err
	}

	return CalculateTaxResult{
		TaxableIncome:    response.GetTaxableIncome(),
		TaxOwed:          response.GetTaxOwed(),
		EffectiveTaxRate: response.GetEffectiveTaxRate(),
	}, nil
}
