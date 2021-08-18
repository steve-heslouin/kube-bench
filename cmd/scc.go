package cmd

import (
	"context"
	"fmt"

	securitycenter "cloud.google.com/go/securitycenter/apiv1"
	securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1"
)

// createFindingWithProperties demonstrates how to create a new security
// finding in CSCC that includes additional metadata via sourceProperties.
// sourceName is the full resource name of the source the finding should be
// associated with.
func writeFindingToScc(req []*securitycenterpb.CreateFindingRequest) error {
	// Instantiate a context and a security service client to make API calls.
	ctx := context.Background()
	client, err := securitycenter.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("securitycenter.NewClient: %v", err)
	}
	defer client.Close() // Closing the client safely cleans up background resources.

	for _, find := range req {
		finding, err := client.CreateFinding(ctx, find)
		if err != nil {
			return fmt.Errorf("CreateFinding: %v %s", err, finding.Name)
		}
	}
	
	return nil
}