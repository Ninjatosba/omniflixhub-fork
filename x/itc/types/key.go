package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName               = "itc"
	StoreKey          string = ModuleName
	QuerierRoute      string = ModuleName
	RouterKey         string = ModuleName
	DefaultParamSpace        = ModuleName
)

var (
	PrefixCampaignId         = []byte{0x01}
	PrefixCampaignCreator    = []byte{0x02}
	PrefixCampaignsCount     = []byte{0x03}
	PrefixNextCampaignNumber = []byte{0x10}
	PrefixClaimByCampaignId  = []byte{0x11}
	PrefixClaimsByClaimer    = []byte{0x12}
	PrefixInactiveCampaign   = []byte{0x13}
	PrefixActiveCampaign     = []byte{0x14}
)

func KeyCampaignIdPrefix(id string) []byte {
	return append(PrefixCampaignId, []byte(id)...)
}

func KeyCampaignCreatorPrefix(creator sdk.AccAddress, id string) []byte {
	return append(append(PrefixCampaignCreator, creator.Bytes()...), []byte(id)...)
}

func KeyClaimPrefix(id uint64) []byte {
	return append(PrefixClaimByCampaignId, sdk.Uint64ToBigEndian(id)...)
}

func KeyInActiveCampaignPrefix(id uint64) []byte {
	return append(PrefixInactiveCampaign, sdk.Uint64ToBigEndian(id)...)
}

func KeyActiveCampaignPrefix(id uint64) []byte {
	return append(PrefixActiveCampaign, sdk.Uint64ToBigEndian(id)...)
}
