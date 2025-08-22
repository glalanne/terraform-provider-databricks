package xpprovider

import (
	"context"
	"github.com/databricks/terraform-provider-databricks/internal/providers/sdkv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GetProvider(_ context.Context) (*schema.Provider, error) {
	return sdkv2.DatabricksProvider(), nil
}
