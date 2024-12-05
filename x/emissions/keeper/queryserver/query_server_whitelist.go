package queryserver

import (
	"context"
	"time"

	"cosmossdk.io/errors"
	"github.com/allora-network/allora-chain/x/emissions/metrics"

	"github.com/allora-network/allora-chain/x/emissions/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Params defines the handler for the Query/Params RPC method.
func (qs queryServer) IsWhitelistAdmin(ctx context.Context, req *types.IsWhitelistAdminRequest) (_ *types.IsWhitelistAdminResponse, err error) {
	defer metrics.RecordMetrics("IsWhitelistAdmin", time.Now(), &err)
	if err := qs.k.ValidateStringIsBech32(req.Address); err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid address: %s", err)
	}

	isAdmin, err := qs.k.IsWhitelistAdmin(ctx, req.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting whitelist admin")
	}

	return &types.IsWhitelistAdminResponse{IsAdmin: isAdmin}, nil
}

func (qs queryServer) IsTopicWorkerWhitelistEnabled(ctx context.Context, req *types.IsTopicWorkerWhitelistEnabledRequest) (_ *types.IsTopicWorkerWhitelistEnabledResponse, err error) {
	defer metrics.RecordMetrics("IsTopicWorkerWhitelistEnabled", time.Now(), &err)

	val, err := qs.k.IsTopicWorkerWhitelistEnabled(ctx, req.TopicId)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting topic worker whitelist enabled")
	}

	return &types.IsTopicWorkerWhitelistEnabledResponse{IsTopicWorkerWhitelistEnabled: val}, nil
}

func (qs queryServer) IsTopicReputerWhitelistEnabled(ctx context.Context, req *types.IsTopicReputerWhitelistEnabledRequest) (_ *types.IsTopicReputerWhitelistEnabledResponse, err error) {
	defer metrics.RecordMetrics("IsTopicReputerWhitelistEnabled", time.Now(), &err)

	val, err := qs.k.IsTopicReputerWhitelistEnabled(ctx, req.TopicId)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting topic reputer whitelist enabled")
	}

	return &types.IsTopicReputerWhitelistEnabledResponse{IsTopicReputerWhitelistEnabled: val}, nil
}

func (qs queryServer) IsWhitelistedTopicCreator(ctx context.Context, req *types.IsWhitelistedTopicCreatorRequest) (_ *types.IsWhitelistedTopicCreatorResponse, err error) {
	defer metrics.RecordMetrics("IsWhitelistedTopicCreator", time.Now(), &err)
	if err := qs.k.ValidateStringIsBech32(req.Address); err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid address: %s", err)
	}

	val, err := qs.k.IsWhitelistedTopicCreator(ctx, req.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting whitelisted topic creator")
	}

	return &types.IsWhitelistedTopicCreatorResponse{IsWhitelistedTopicCreator: val}, nil
}

func (qs queryServer) IsWhitelistedGlobalActor(ctx context.Context, req *types.IsWhitelistedGlobalActorRequest) (_ *types.IsWhitelistedGlobalActorResponse, err error) {
	defer metrics.RecordMetrics("IsWhitelistedGlobalActor", time.Now(), &err)
	if err := qs.k.ValidateStringIsBech32(req.Address); err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid address: %s", err)
	}

	val, err := qs.k.IsWhitelistedGlobalActor(ctx, req.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting whitelisted global actor")
	}

	return &types.IsWhitelistedGlobalActorResponse{IsWhitelistedGlobalActor: val}, nil
}

func (qs queryServer) IsWhitelistedTopicWorker(ctx context.Context, req *types.IsWhitelistedTopicWorkerRequest) (_ *types.IsWhitelistedTopicWorkerResponse, err error) {
	defer metrics.RecordMetrics("IsWhitelistedTopicWorker", time.Now(), &err)
	if err := qs.k.ValidateStringIsBech32(req.Address); err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid address: %s", err)
	}

	val, err := qs.k.IsWhitelistedTopicWorker(ctx, req.TopicId, req.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting whitelisted topic worker")
	}

	return &types.IsWhitelistedTopicWorkerResponse{IsWhitelistedTopicWorker: val}, nil
}

func (qs queryServer) IsWhitelistedTopicReputer(ctx context.Context, req *types.IsWhitelistedTopicReputerRequest) (_ *types.IsWhitelistedTopicReputerResponse, err error) {
	defer metrics.RecordMetrics("IsWhitelistedTopicReputer", time.Now(), &err)
	if err := qs.k.ValidateStringIsBech32(req.Address); err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid address: %s", err)
	}

	val, err := qs.k.IsWhitelistedTopicReputer(ctx, req.TopicId, req.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting whitelisted topic reputer")
	}

	return &types.IsWhitelistedTopicReputerResponse{IsWhitelistedTopicReputer: val}, nil
}

func (qs queryServer) CanUpdateGlobalWhitelists(ctx context.Context, req *types.CanUpdateGlobalWhitelistsRequest) (_ *types.CanUpdateGlobalWhitelistsResponse, err error) {
	defer metrics.RecordMetrics("CanUpdateGlobalWhitelists", time.Now(), &err)
	if err := qs.k.ValidateStringIsBech32(req.Address); err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid address: %s", err)
	}

	val, err := qs.k.CanUpdateGlobalWhitelists(ctx, req.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting can update global whitelists")
	}

	return &types.CanUpdateGlobalWhitelistsResponse{CanUpdateGlobalWhitelists: val}, nil
}

func (qs queryServer) CanUpdateParams(ctx context.Context, req *types.CanUpdateParamsRequest) (_ *types.CanUpdateParamsResponse, err error) {
	defer metrics.RecordMetrics("CanUpdateParams", time.Now(), &err)
	if err := qs.k.ValidateStringIsBech32(req.Address); err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid address: %s", err)
	}

	val, err := qs.k.CanUpdateParams(ctx, req.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting can update params")
	}

	return &types.CanUpdateParamsResponse{CanUpdateParams: val}, nil
}

func (qs queryServer) CanUpdateTopicWhitelist(ctx context.Context, req *types.CanUpdateTopicWhitelistRequest) (_ *types.CanUpdateTopicWhitelistResponse, err error) {
	defer metrics.RecordMetrics("CanUpdateTopicWhitelist", time.Now(), &err)
	if err := qs.k.ValidateStringIsBech32(req.Address); err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid address: %s", err)
	}

	val, err := qs.k.CanUpdateTopicWhitelist(ctx, req.TopicId, req.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting can update topic whitelist")
	}

	return &types.CanUpdateTopicWhitelistResponse{CanUpdateTopicWhitelist: val}, nil
}

func (qs queryServer) CanCreateTopic(ctx context.Context, req *types.CanCreateTopicRequest) (_ *types.CanCreateTopicResponse, err error) {
	defer metrics.RecordMetrics("CanCreateTopic", time.Now(), &err)
	if err := qs.k.ValidateStringIsBech32(req.Address); err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid address: %s", err)
	}

	val, err := qs.k.CanCreateTopic(ctx, req.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting can create topic")
	}

	return &types.CanCreateTopicResponse{CanCreateTopic: val}, nil
}

func (qs queryServer) CanSubmitWorkerPayload(ctx context.Context, req *types.CanSubmitWorkerPayloadRequest) (_ *types.CanSubmitWorkerPayloadResponse, err error) {
	defer metrics.RecordMetrics("CanSubmitWorkerPayload", time.Now(), &err)
	if err := qs.k.ValidateStringIsBech32(req.Address); err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid address: %s", err)
	}

	val, err := qs.k.CanSubmitWorkerPayload(ctx, req.TopicId, req.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting can submit worker payload")
	}

	return &types.CanSubmitWorkerPayloadResponse{CanSubmitWorkerPayload: val}, nil
}

func (qs queryServer) CanSubmitReputerPayload(ctx context.Context, req *types.CanSubmitReputerPayloadRequest) (_ *types.CanSubmitReputerPayloadResponse, err error) {
	defer metrics.RecordMetrics("CanSubmitReputerPayload", time.Now(), &err)
	if err := qs.k.ValidateStringIsBech32(req.Address); err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid address: %s", err)
	}

	val, err := qs.k.CanSubmitReputerPayload(ctx, req.TopicId, req.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting can submit reputer payload")
	}

	return &types.CanSubmitReputerPayloadResponse{CanSubmitReputerPayload: val}, nil
}
