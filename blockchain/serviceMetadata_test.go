package blockchain

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/singnet/snet-daemon/config"
	"github.com/stretchr/testify/assert"
)

var testLicenseJsonData = "\n  \"licenses\": { \"tiers\": [{\n  \"type\": \"Tier\",\n  \"planName\": \"Tier AAA\",\n  \"grpcServiceName\": \"ServiceA\",\n  \"grpcMethodName\": \"MethodA\",\n  \"range\": [\n    {\n      \"high\": 100,\n      \"fixedPriceInCogs\": 1\n    },\n    {\n      \"high\": 200,\n      \"fixedPriceInCogs\": 200000\n    },\n    {\n      \"high\": 300,\n      \"fixedPriceInCogs\": 100000\n    }\n  ],\n  \"detailsUrl\": \"http://abc.org/licenses/Tier.html\",\n  \"isActive\": \"true/false\"\n},\n {\n  \"type\": \"Tier\",\n  \"planName\": \"Tier BBB Applicable for All service.methods\",\n  \"range\": [\n    {\n      \"high\": 100,\n      \"fixedPriceInCogs\": 1\n    },\n    {\n      \"high\": 200,\n      \"fixedPriceInCogs\": 200000\n    },\n    {\n      \"high\": 300,\n      \"fixedPriceInCogs\": 100000\n    }\n  ],\n  \"detailsUrl\": \"http://abc.org/licenses/Tier.html\",\n  \"isActive\": \"true/false\"\n}], " +
	"\"subscriptions\": {\n   \"subscription\": [\n  {\n    \"periodInDays\": 30,\n    \"creditsInAGI\": 120,\n    \"planName\": \"Monthly For ServiceA/MethodA\",\n    \"licenseCost\": 90,\n    \"grpcServiceName\": \"ServiceA\",\n    \"grpcMethodName\": \"MethodA\"\n  },\n  {\n    \"periodInDays\": 30,\n    \"creditsInAGI\": 123,\n    \"planName\": \"Monthly\",\n    \"licenseCost\": 93\n  },\n  {\n    \"periodInDays\": 120,\n    \"creditsInAGI\": 160,\n    \"licenseCost\": 120,\n    \"planName\": \"Quarterly\"\n  },\n  {\n    \"periodInDays\": 365,\n    \"creditsInAGI\": 430,\n    \"licenseCost\": 390,\n    \"planName\": \"Yearly\"\n  }\n],       \"type\": \"Subscription\",\n          \"detailsUrl\": \"http://abc.org/licenses/Subscription.html\",\n          \"isActive\": \"true/false\"\n        }\n      }"
var testJsonData = "{   \"version\": 1,   \"display_name\": \"Example1\",   \"encoding\": \"grpc\",   \"service_type\": \"grpc\",   \"payment_expiration_threshold\": 40320,   \"model_ipfs_hash\": \"Qmdiq8Hu6dYiwp712GtnbBxagyfYyvUY1HYqkH7iN76UCc\", " +
	"  \"mpe_address\": \"0x7E6366Fbe3bdfCE3C906667911FC5237Cc96BD08\",   \"groups\": [     {    \"free_calls\": 12,  \"free_call_signer_address\": \"0x7DF35C98f41F3Af0df1dc4c7F7D4C19a71Dd059F\",  \"endpoints\": [\"http://34.344.33.1:2379\",\"http://34.344.33.1:2389\"],       \"group_id\": \"88ybRIg2wAx55mqVsA6sB4S7WxPQHNKqa4BPu/bhj+U=\",\"group_name\": \"default_group\",  " + testLicenseJsonData + " ,  \"pricing\": [         {           \"price_model\": \"fixed_price\",           \"price_in_cogs\": 2         },          {         \"package_name\": \"example_service\",         \"price_model\": \"fixed_price_per_method\",         \"default\":true,         \"details\": [           {             \"service_name\": \"Calculator\",             \"method_pricing\": [               {                 \"method_name\": \"add\",                 \"price_in_cogs\": 2               },               {                 \"method_name\": \"sub\",                 \"price_in_cogs\": 1               },               {                 \"method_name\": \"div\",                 \"price_in_cogs\": 2               },               {                 \"method_name\": \"mul\",                 \"price_in_cogs\": 3               }             ]           },           {             \"service_name\": \"Calculator2\",             \"method_pricing\": [               {                 \"method_name\": \"add\",                 \"price_in_cogs\": 2               },               {                 \"method_name\": \"sub\",                 \"price_in_cogs\": 1               },               {                 \"method_name\": \"div\",                 \"price_in_cogs\": 3               },               {                 \"method_name\": \"mul\",                 \"price_in_cogs\": 2               }             ]           }         ]       }]     },     {       \"endpoints\": [\"http://97.344.33.1:2379\",\"http://67.344.33.1:2389\"],       \"group_id\": \"99ybRIg2wAx55mqVsA6sB4S7WxPQHNKqa4BPu/bhj+U=\",       \"pricing\": [         {         \"package_name\": \"example_service\",         \"price_model\": \"fixed_price_per_method\",         \"details\": [           {             \"service_name\": \"Calculator\",             \"method_pricing\": [               {                 \"method_name\": \"add\",                 \"price_in_cogs\": 2               },               {                 \"method_name\": \"sub\",                 \"price_in_cogs\": 1               },               {                 \"method_name\": \"div\",                 \"price_in_cogs\": 2               },               {                 \"method_name\": \"mul\",                 \"price_in_cogs\": 3               }             ]           },           {             \"service_name\": \"Calculator2\",             \"method_pricing\": [               {                 \"method_name\": \"add\",                 \"price_in_cogs\": 2               },               {                 \"method_name\": \"sub\",                 \"price_in_cogs\": 1               },               {                 \"method_name\": \"div\",                 \"price_in_cogs\": 3               },               {                 \"method_name\": \"mul\",                 \"price_in_cogs\": 2               }             ]           }         ]       }]     }   ] } "

func TestAllGetterMethods(t *testing.T) {
	fmt.Println(testJsonData)
	metaData, err := InitServiceMetaDataFromJson(testJsonData)
	assert.Equal(t, err, nil)

	assert.Equal(t, metaData.GetVersion(), 1)
	assert.Equal(t, metaData.GetDisplayName(), "Example1")
	assert.Equal(t, metaData.GetServiceType(), "grpc")
	assert.Equal(t, metaData.GetWireEncoding(), "grpc")
	assert.Nil(t, metaData.GetDefaultPricing().PriceInCogs)
	assert.Equal(t, metaData.GetDefaultPricing().PricingDetails[0].MethodPricing[0].PriceInCogs, big.NewInt(2))
	assert.Equal(t, metaData.GetMpeAddress(), common.HexToAddress("0x7E6366Fbe3bdfCE3C906667911FC5237Cc96BD08"))
	assert.Equal(t, metaData.FreeCallSignerAddress(), common.HexToAddress("0x7DF35C98f41F3Af0df1dc4c7F7D4C19a71Dd059F"))
	assert.True(t, metaData.IsFreeCallAllowed())
	assert.Equal(t, 12, metaData.GetFreeCallsAllowed())
	assert.Equal(t, metaData.GetLicenses().Subscriptions.Type, "Subscription")

}

func TestSubscription(t *testing.T) {
	fmt.Println(testJsonData)
	metaData, err := InitServiceMetaDataFromJson(testJsonData)
	assert.Equal(t, err, nil)
	assert.Equal(t, 12, metaData.GetFreeCallsAllowed())
	assert.Equal(t, metaData.GetLicenses().Subscriptions.Type, "Subscription")
	assert.Equal(t, len(metaData.GetLicenses().Subscriptions.Subscription), 4)
	assert.Equal(t, metaData.GetLicenses().Subscriptions.Subscription[0].PlanName, "Monthly For ServiceA/MethodA")
	assert.Equal(t, metaData.GetLicenses().Subscriptions.Subscription[0].GrpcMethodName, "MethodA")
	assert.Equal(t, metaData.GetLicenses().Subscriptions.Subscription[0].GrpcServiceName, "ServiceA")
}

func TestTiers(t *testing.T) {
	fmt.Println(testJsonData)
	metaData, err := InitServiceMetaDataFromJson(testJsonData)
	assert.Equal(t, err, nil)

	assert.Equal(t, metaData.GetLicenses().Tiers[0].Type, "Tier")
	assert.Equal(t, metaData.GetLicenses().Tiers[0].Range[0].High, 200)
}
func TestInitServiceMetaDataFromJson(t *testing.T) {
	//Parse Bad JSON
	_, err := InitServiceMetaDataFromJson(strings.Replace(testJsonData, "{", "", 1))
	if err != nil {
		assert.Equal(t, err.Error(), "invalid character ':' after top-level value")
	}

	//Parse Bad JSON
	_, err = InitServiceMetaDataFromJson(strings.Replace(testJsonData, "0x7DF35C98f41F3Af0df1dc4c7F7D4C19a71Dd059F", "", 1))
	if err != nil {
		assert.Equal(t, err.Error(), "MetaData does not have 'free_call_signer_address defined correctly")
	}
	_, err = InitServiceMetaDataFromJson(strings.Replace(testJsonData, "default_pricing", "dummy", 1))
	if err != nil {
		assert.Equal(t, err.Error(), "MetaData does not have the default pricing set ")
	}

}

func TestReadServiceMetaDataFromLocalFile(t *testing.T) {
	metadata, err := ReadServiceMetaDataFromLocalFile("../service_metadata.json")
	assert.Equal(t, err, nil)
	assert.Equal(t, metadata.Version, 1)
}

func Test_getServiceMetaDataUrifromRegistry(t *testing.T) {
	assert.Panics(t, func() { getServiceMetaDataUrifromRegistry() })
	config.Vip().Set(config.BlockChainNetworkSelected, "ropsten")
	config.Validate()
	assert.Panics(t, func() { getServiceMetaDataUrifromRegistry() })

}

func Test_setDefaultPricing(t *testing.T) {
	err := setDefaultPricing(&ServiceMetadata{})
	assert.NotNil(t, err)
	err = setDefaultPricing(&ServiceMetadata{Groups: []OrganizationGroup{{GroupName: "default_group"}}})
	assert.Equal(t, err.Error(), "MetaData does not have the default pricing set ")
}

func Test_setGroup(t *testing.T) {
	err := setGroup(&ServiceMetadata{})
	assert.Equal(t, err.Error(), "group name default_group in config is invalid, there was no group found with this name in the metadata")
}
