// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package marketplacemetering provides a client for AWSMarketplace Metering.
package marketplacemetering

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opBatchMeterUsage = "BatchMeterUsage"

// BatchMeterUsageRequest generates a "aws/request.Request" representing the
// client's request for the BatchMeterUsage operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See BatchMeterUsage for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the BatchMeterUsage method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the BatchMeterUsageRequest method.
//    req, resp := client.BatchMeterUsageRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/BatchMeterUsage
func (c *MarketplaceMetering) BatchMeterUsageRequest(input *BatchMeterUsageInput) (req *request.Request, output *BatchMeterUsageOutput) {
	op := &request.Operation{
		Name:       opBatchMeterUsage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchMeterUsageInput{}
	}

	output = &BatchMeterUsageOutput{}
	req = c.newRequest(op, input, output)
	return
}

// BatchMeterUsage API operation for AWSMarketplace Metering.
//
// BatchMeterUsage is called from a SaaS application listed on the AWS Marketplace
// to post metering records for a set of customers.
//
// For identical requests, the API is idempotent; requests can be retried with
// the same records or a subset of the input records.
//
// Every request to BatchMeterUsage is for one product. If you need to meter
// usage for multiple products, you must make multiple calls to BatchMeterUsage.
//
// BatchMeterUsage can process up to 25 UsageRecords at a time.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWSMarketplace Metering's
// API operation BatchMeterUsage for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInternalServiceErrorException "InternalServiceErrorException"
//   An internal error has occurred. Retry your request. If the problem persists,
//   post a message with details on the AWS forums.
//
//   * ErrCodeInvalidProductCodeException "InvalidProductCodeException"
//   The product code passed does not match the product code used for publishing
//   the product.
//
//   * ErrCodeInvalidUsageDimensionException "InvalidUsageDimensionException"
//   The usage dimension does not match one of the UsageDimensions associated
//   with products.
//
//   * ErrCodeInvalidCustomerIdentifierException "InvalidCustomerIdentifierException"
//   You have metered usage for a CustomerIdentifier that does not exist.
//
//   * ErrCodeTimestampOutOfBoundsException "TimestampOutOfBoundsException"
//   The timestamp value passed in the meterUsage() is out of allowed range.
//
//   * ErrCodeThrottlingException "ThrottlingException"
//   The calls to the MeterUsage API are throttled.
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/BatchMeterUsage
func (c *MarketplaceMetering) BatchMeterUsage(input *BatchMeterUsageInput) (*BatchMeterUsageOutput, error) {
	req, out := c.BatchMeterUsageRequest(input)
	err := req.Send()
	return out, err
}

const opMeterUsage = "MeterUsage"

// MeterUsageRequest generates a "aws/request.Request" representing the
// client's request for the MeterUsage operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See MeterUsage for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the MeterUsage method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the MeterUsageRequest method.
//    req, resp := client.MeterUsageRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/MeterUsage
func (c *MarketplaceMetering) MeterUsageRequest(input *MeterUsageInput) (req *request.Request, output *MeterUsageOutput) {
	op := &request.Operation{
		Name:       opMeterUsage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &MeterUsageInput{}
	}

	output = &MeterUsageOutput{}
	req = c.newRequest(op, input, output)
	return
}

// MeterUsage API operation for AWSMarketplace Metering.
//
// API to emit metering records. For identical requests, the API is idempotent.
// It simply returns the metering record ID.
//
// MeterUsage is authenticated on the buyer's AWS account, generally when running
// from an EC2 instance on the AWS Marketplace.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWSMarketplace Metering's
// API operation MeterUsage for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInternalServiceErrorException "InternalServiceErrorException"
//   An internal error has occurred. Retry your request. If the problem persists,
//   post a message with details on the AWS forums.
//
//   * ErrCodeInvalidProductCodeException "InvalidProductCodeException"
//   The product code passed does not match the product code used for publishing
//   the product.
//
//   * ErrCodeInvalidUsageDimensionException "InvalidUsageDimensionException"
//   The usage dimension does not match one of the UsageDimensions associated
//   with products.
//
//   * ErrCodeInvalidEndpointRegionException "InvalidEndpointRegionException"
//   The endpoint being called is in a region different from your EC2 instance.
//   The region of the Metering service endpoint and the region of the EC2 instance
//   must match.
//
//   * ErrCodeTimestampOutOfBoundsException "TimestampOutOfBoundsException"
//   The timestamp value passed in the meterUsage() is out of allowed range.
//
//   * ErrCodeDuplicateRequestException "DuplicateRequestException"
//   A metering record has already been emitted by the same EC2 instance for the
//   given {usageDimension, timestamp} with a different usageQuantity.
//
//   * ErrCodeThrottlingException "ThrottlingException"
//   The calls to the MeterUsage API are throttled.
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/MeterUsage
func (c *MarketplaceMetering) MeterUsage(input *MeterUsageInput) (*MeterUsageOutput, error) {
	req, out := c.MeterUsageRequest(input)
	err := req.Send()
	return out, err
}

const opResolveCustomer = "ResolveCustomer"

// ResolveCustomerRequest generates a "aws/request.Request" representing the
// client's request for the ResolveCustomer operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See ResolveCustomer for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the ResolveCustomer method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the ResolveCustomerRequest method.
//    req, resp := client.ResolveCustomerRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/ResolveCustomer
func (c *MarketplaceMetering) ResolveCustomerRequest(input *ResolveCustomerInput) (req *request.Request, output *ResolveCustomerOutput) {
	op := &request.Operation{
		Name:       opResolveCustomer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResolveCustomerInput{}
	}

	output = &ResolveCustomerOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ResolveCustomer API operation for AWSMarketplace Metering.
//
// ResolveCustomer is called by a SaaS application during the registration process.
// When a buyer visits your website during the registration process, the buyer
// submits a registration token through their browser. The registration token
// is resolved through this API to obtain a CustomerIdentifier and product code.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWSMarketplace Metering's
// API operation ResolveCustomer for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInvalidTokenException "InvalidTokenException"

//
//   * ErrCodeExpiredTokenException "ExpiredTokenException"
//   The submitted registration token has expired. This can happen if the buyer's
//   browser takes too long to redirect to your page, the buyer has resubmitted
//   the registration token, or your application has held on to the registration
//   token for too long. Your SaaS registration website should redeem this token
//   as soon as it is submitted by the buyer's browser.
//
//   * ErrCodeThrottlingException "ThrottlingException"
//   The calls to the MeterUsage API are throttled.
//
//   * ErrCodeInternalServiceErrorException "InternalServiceErrorException"
//   An internal error has occurred. Retry your request. If the problem persists,
//   post a message with details on the AWS forums.
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/ResolveCustomer
func (c *MarketplaceMetering) ResolveCustomer(input *ResolveCustomerInput) (*ResolveCustomerOutput, error) {
	req, out := c.ResolveCustomerRequest(input)
	err := req.Send()
	return out, err
}

// A BatchMeterUsageRequest contains UsageRecords, which indicate quantities
// of usage within your application.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/BatchMeterUsageRequest
type BatchMeterUsageInput struct {
	_ struct{} `type:"structure"`

	// Product code is used to uniquely identify a product in AWS Marketplace. The
	// product code should be the same as the one used during the publishing of
	// a new product.
	//
	// ProductCode is a required field
	ProductCode *string `min:"1" type:"string" required:"true"`

	// The set of UsageRecords to submit. BatchMeterUsage accepts up to 25 UsageRecords
	// at a time.
	//
	// UsageRecords is a required field
	UsageRecords []*UsageRecord `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchMeterUsageInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s BatchMeterUsageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchMeterUsageInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "BatchMeterUsageInput"}
	if s.ProductCode == nil {
		invalidParams.Add(request.NewErrParamRequired("ProductCode"))
	}
	if s.ProductCode != nil && len(*s.ProductCode) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("ProductCode", 1))
	}
	if s.UsageRecords == nil {
		invalidParams.Add(request.NewErrParamRequired("UsageRecords"))
	}
	if s.UsageRecords != nil {
		for i, v := range s.UsageRecords {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "UsageRecords", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetProductCode sets the ProductCode field's value.
func (s *BatchMeterUsageInput) SetProductCode(v string) *BatchMeterUsageInput {
	s.ProductCode = &v
	return s
}

// SetUsageRecords sets the UsageRecords field's value.
func (s *BatchMeterUsageInput) SetUsageRecords(v []*UsageRecord) *BatchMeterUsageInput {
	s.UsageRecords = v
	return s
}

// Contains the UsageRecords processed by BatchMeterUsage and any records that
// have failed due to transient error.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/BatchMeterUsageResult
type BatchMeterUsageOutput struct {
	_ struct{} `type:"structure"`

	// Contains all UsageRecords processed by BatchMeterUsage. These records were
	// either honored by AWS Marketplace Metering Service or were invalid.
	Results []*UsageRecordResult `type:"list"`

	// Contains all UsageRecords that were not processed by BatchMeterUsage. This
	// is a list of UsageRecords. You can retry the failed request by making another
	// BatchMeterUsage call with this list as input in the BatchMeterUsageRequest.
	UnprocessedRecords []*UsageRecord `type:"list"`
}

// String returns the string representation
func (s BatchMeterUsageOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s BatchMeterUsageOutput) GoString() string {
	return s.String()
}

// SetResults sets the Results field's value.
func (s *BatchMeterUsageOutput) SetResults(v []*UsageRecordResult) *BatchMeterUsageOutput {
	s.Results = v
	return s
}

// SetUnprocessedRecords sets the UnprocessedRecords field's value.
func (s *BatchMeterUsageOutput) SetUnprocessedRecords(v []*UsageRecord) *BatchMeterUsageOutput {
	s.UnprocessedRecords = v
	return s
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/MeterUsageRequest
type MeterUsageInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the permissions required for the action, but does
	// not make the request. If you have the permissions, the request returns DryRunOperation;
	// otherwise, it returns UnauthorizedException.
	//
	// DryRun is a required field
	DryRun *bool `type:"boolean" required:"true"`

	// Product code is used to uniquely identify a product in AWS Marketplace. The
	// product code should be the same as the one used during the publishing of
	// a new product.
	//
	// ProductCode is a required field
	ProductCode *string `min:"1" type:"string" required:"true"`

	// Timestamp of the hour, recorded in UTC. The seconds and milliseconds portions
	// of the timestamp will be ignored.
	//
	// Timestamp is a required field
	Timestamp *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`

	// It will be one of the fcp dimension name provided during the publishing of
	// the product.
	//
	// UsageDimension is a required field
	UsageDimension *string `min:"1" type:"string" required:"true"`

	// Consumption value for the hour.
	//
	// UsageQuantity is a required field
	UsageQuantity *int64 `type:"integer" required:"true"`
}

// String returns the string representation
func (s MeterUsageInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s MeterUsageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MeterUsageInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "MeterUsageInput"}
	if s.DryRun == nil {
		invalidParams.Add(request.NewErrParamRequired("DryRun"))
	}
	if s.ProductCode == nil {
		invalidParams.Add(request.NewErrParamRequired("ProductCode"))
	}
	if s.ProductCode != nil && len(*s.ProductCode) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("ProductCode", 1))
	}
	if s.Timestamp == nil {
		invalidParams.Add(request.NewErrParamRequired("Timestamp"))
	}
	if s.UsageDimension == nil {
		invalidParams.Add(request.NewErrParamRequired("UsageDimension"))
	}
	if s.UsageDimension != nil && len(*s.UsageDimension) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("UsageDimension", 1))
	}
	if s.UsageQuantity == nil {
		invalidParams.Add(request.NewErrParamRequired("UsageQuantity"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDryRun sets the DryRun field's value.
func (s *MeterUsageInput) SetDryRun(v bool) *MeterUsageInput {
	s.DryRun = &v
	return s
}

// SetProductCode sets the ProductCode field's value.
func (s *MeterUsageInput) SetProductCode(v string) *MeterUsageInput {
	s.ProductCode = &v
	return s
}

// SetTimestamp sets the Timestamp field's value.
func (s *MeterUsageInput) SetTimestamp(v time.Time) *MeterUsageInput {
	s.Timestamp = &v
	return s
}

// SetUsageDimension sets the UsageDimension field's value.
func (s *MeterUsageInput) SetUsageDimension(v string) *MeterUsageInput {
	s.UsageDimension = &v
	return s
}

// SetUsageQuantity sets the UsageQuantity field's value.
func (s *MeterUsageInput) SetUsageQuantity(v int64) *MeterUsageInput {
	s.UsageQuantity = &v
	return s
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/MeterUsageResult
type MeterUsageOutput struct {
	_ struct{} `type:"structure"`

	MeteringRecordId *string `type:"string"`
}

// String returns the string representation
func (s MeterUsageOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s MeterUsageOutput) GoString() string {
	return s.String()
}

// SetMeteringRecordId sets the MeteringRecordId field's value.
func (s *MeterUsageOutput) SetMeteringRecordId(v string) *MeterUsageOutput {
	s.MeteringRecordId = &v
	return s
}

// Contains input to the ResolveCustomer operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/ResolveCustomerRequest
type ResolveCustomerInput struct {
	_ struct{} `type:"structure"`

	// When a buyer visits your website during the registration process, the buyer
	// submits a registration token through the browser. The registration token
	// is resolved to obtain a CustomerIdentifier and product code.
	//
	// RegistrationToken is a required field
	RegistrationToken *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ResolveCustomerInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ResolveCustomerInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResolveCustomerInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ResolveCustomerInput"}
	if s.RegistrationToken == nil {
		invalidParams.Add(request.NewErrParamRequired("RegistrationToken"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetRegistrationToken sets the RegistrationToken field's value.
func (s *ResolveCustomerInput) SetRegistrationToken(v string) *ResolveCustomerInput {
	s.RegistrationToken = &v
	return s
}

// The result of the ResolveCustomer operation. Contains the CustomerIdentifier
// and product code.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/ResolveCustomerResult
type ResolveCustomerOutput struct {
	_ struct{} `type:"structure"`

	// The CustomerIdentifier is used to identify an individual customer in your
	// application. Calls to BatchMeterUsage require CustomerIdentifiers for each
	// UsageRecord.
	CustomerIdentifier *string `min:"1" type:"string"`

	// The product code is returned to confirm that the buyer is registering for
	// your product. Subsequent BatchMeterUsage calls should be made using this
	// product code.
	ProductCode *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ResolveCustomerOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ResolveCustomerOutput) GoString() string {
	return s.String()
}

// SetCustomerIdentifier sets the CustomerIdentifier field's value.
func (s *ResolveCustomerOutput) SetCustomerIdentifier(v string) *ResolveCustomerOutput {
	s.CustomerIdentifier = &v
	return s
}

// SetProductCode sets the ProductCode field's value.
func (s *ResolveCustomerOutput) SetProductCode(v string) *ResolveCustomerOutput {
	s.ProductCode = &v
	return s
}

// A UsageRecord indicates a quantity of usage for a given product, customer,
// dimension and time.
//
// Multiple requests with the same UsageRecords as input will be deduplicated
// to prevent double charges.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/UsageRecord
type UsageRecord struct {
	_ struct{} `type:"structure"`

	// The CustomerIdentifier is obtained through the ResolveCustomer operation
	// and represents an individual buyer in your application.
	//
	// CustomerIdentifier is a required field
	CustomerIdentifier *string `min:"1" type:"string" required:"true"`

	// During the process of registering a product on AWS Marketplace, up to eight
	// dimensions are specified. These represent different units of value in your
	// application.
	//
	// Dimension is a required field
	Dimension *string `min:"1" type:"string" required:"true"`

	// The quantity of usage consumed by the customer for the given dimension and
	// time.
	//
	// Quantity is a required field
	Quantity *int64 `type:"integer" required:"true"`

	// Timestamp of the hour, recorded in UTC. The seconds and milliseconds portions
	// of the timestamp will be ignored.
	//
	// Your application can meter usage for up to one hour in the past.
	//
	// Timestamp is a required field
	Timestamp *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`
}

// String returns the string representation
func (s UsageRecord) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UsageRecord) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UsageRecord) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UsageRecord"}
	if s.CustomerIdentifier == nil {
		invalidParams.Add(request.NewErrParamRequired("CustomerIdentifier"))
	}
	if s.CustomerIdentifier != nil && len(*s.CustomerIdentifier) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("CustomerIdentifier", 1))
	}
	if s.Dimension == nil {
		invalidParams.Add(request.NewErrParamRequired("Dimension"))
	}
	if s.Dimension != nil && len(*s.Dimension) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Dimension", 1))
	}
	if s.Quantity == nil {
		invalidParams.Add(request.NewErrParamRequired("Quantity"))
	}
	if s.Timestamp == nil {
		invalidParams.Add(request.NewErrParamRequired("Timestamp"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCustomerIdentifier sets the CustomerIdentifier field's value.
func (s *UsageRecord) SetCustomerIdentifier(v string) *UsageRecord {
	s.CustomerIdentifier = &v
	return s
}

// SetDimension sets the Dimension field's value.
func (s *UsageRecord) SetDimension(v string) *UsageRecord {
	s.Dimension = &v
	return s
}

// SetQuantity sets the Quantity field's value.
func (s *UsageRecord) SetQuantity(v int64) *UsageRecord {
	s.Quantity = &v
	return s
}

// SetTimestamp sets the Timestamp field's value.
func (s *UsageRecord) SetTimestamp(v time.Time) *UsageRecord {
	s.Timestamp = &v
	return s
}

// A UsageRecordResult indicates the status of a given UsageRecord processed
// by BatchMeterUsage.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14/UsageRecordResult
type UsageRecordResult struct {
	_ struct{} `type:"structure"`

	// The MeteringRecordId is a unique identifier for this metering event.
	MeteringRecordId *string `type:"string"`

	// The UsageRecordResult Status indicates the status of an individual UsageRecord
	// processed by BatchMeterUsage.
	//
	//    * Success- The UsageRecord was accepted and honored by BatchMeterUsage.
	//
	//    * CustomerNotSubscribed- The CustomerIdentifier specified is not subscribed
	//    to your product. The UsageRecord was not honored. Future UsageRecords
	//    for this customer will fail until the customer subscribes to your product.
	//
	//    * DuplicateRecord- Indicates that the UsageRecord was invalid and not
	//    honored. A previously metered UsageRecord had the same customer, dimension,
	//    and time, but a different quantity.
	Status *string `type:"string" enum:"UsageRecordResultStatus"`

	// The UsageRecord that was part of the BatchMeterUsage request.
	UsageRecord *UsageRecord `type:"structure"`
}

// String returns the string representation
func (s UsageRecordResult) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UsageRecordResult) GoString() string {
	return s.String()
}

// SetMeteringRecordId sets the MeteringRecordId field's value.
func (s *UsageRecordResult) SetMeteringRecordId(v string) *UsageRecordResult {
	s.MeteringRecordId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *UsageRecordResult) SetStatus(v string) *UsageRecordResult {
	s.Status = &v
	return s
}

// SetUsageRecord sets the UsageRecord field's value.
func (s *UsageRecordResult) SetUsageRecord(v *UsageRecord) *UsageRecordResult {
	s.UsageRecord = v
	return s
}

const (
	// UsageRecordResultStatusSuccess is a UsageRecordResultStatus enum value
	UsageRecordResultStatusSuccess = "Success"

	// UsageRecordResultStatusCustomerNotSubscribed is a UsageRecordResultStatus enum value
	UsageRecordResultStatusCustomerNotSubscribed = "CustomerNotSubscribed"

	// UsageRecordResultStatusDuplicateRecord is a UsageRecordResultStatus enum value
	UsageRecordResultStatusDuplicateRecord = "DuplicateRecord"
)
