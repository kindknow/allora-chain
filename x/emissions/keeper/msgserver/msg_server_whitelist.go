package msgserver

import (
	"context"
	"time"

	"github.com/allora-network/allora-chain/x/emissions/metrics"
	"github.com/allora-network/allora-chain/x/emissions/types"
)

func (ms msgServer) AddToWhitelistAdmin(ctx context.Context, msg *types.AddToWhitelistAdminRequest) (_ *types.AddToWhitelistAdminResponse, err error) {
	defer metrics.RecordMetrics("AddToWhitelistAdmin", time.Now(), &err)

	// Validate the sender address
	err = ms.k.ValidateStringIsBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	// Validate the address
	err = ms.k.ValidateStringIsBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	// Check that sender is permitted to update global whitelists
	canUpdate, err := ms.k.CanUpdateGlobalWhitelists(ctx, msg.Sender)
	if err != nil {
		return nil, err
	} else if !canUpdate {
		return nil, types.ErrNotPermittedToUpdateWhitelistAdmins
	}

	// Add the address to the whitelist
	return &types.AddToWhitelistAdminResponse{}, ms.k.AddWhitelistAdmin(ctx, msg.Address)
}

func (ms msgServer) RemoveFromWhitelistAdmin(ctx context.Context, msg *types.RemoveFromWhitelistAdminRequest) (_ *types.RemoveFromWhitelistAdminResponse, err error) {
	defer metrics.RecordMetrics("RemoveFromWhitelistAdmin", time.Now(), &err)

	// Validate the sender address
	err = ms.k.ValidateStringIsBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	// Validate the address
	err = ms.k.ValidateStringIsBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	// Check that sender is permitted to update global whitelists
	canUpdate, err := ms.k.CanUpdateGlobalWhitelists(ctx, msg.Sender)
	if err != nil {
		return nil, err
	} else if !canUpdate {
		return nil, types.ErrNotPermittedToUpdateWhitelistAdmins
	}

	// Remove the address from the whitelist
	return &types.RemoveFromWhitelistAdminResponse{}, ms.k.RemoveWhitelistAdmin(ctx, msg.Address)
}

func (ms msgServer) AddToGlobalWhitelist(ctx context.Context, msg *types.AddToGlobalWhitelistRequest) (_ *types.AddToGlobalWhitelistResponse, err error) {
	defer metrics.RecordMetrics("AddToGlobalWhitelist", time.Now(), &err)

	// Validate the sender address
	err = ms.k.ValidateStringIsBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	// Validate the address
	err = ms.k.ValidateStringIsBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	// Check that sender is permitted to update global whitelists
	canUpdate, err := ms.k.CanUpdateGlobalWhitelists(ctx, msg.Sender)
	if err != nil {
		return nil, err
	} else if !canUpdate {
		return nil, types.ErrNotPermittedToUpdateGlobalWhitelist
	}

	// Add the address to the whitelist
	return &types.AddToGlobalWhitelistResponse{}, ms.k.AddToGlobalWhitelist(ctx, msg.Address)
}

func (ms msgServer) RemoveFromGlobalWhitelist(ctx context.Context, msg *types.RemoveFromGlobalWhitelistRequest) (_ *types.RemoveFromGlobalWhitelistResponse, err error) {
	defer metrics.RecordMetrics("RemoveFromGlobalWhitelist", time.Now(), &err)

	// Validate the sender address
	err = ms.k.ValidateStringIsBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	// Validate the address
	err = ms.k.ValidateStringIsBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	// Check that sender is permitted to update global whitelists
	canUpdate, err := ms.k.CanUpdateGlobalWhitelists(ctx, msg.Sender)
	if err != nil {
		return nil, err
	} else if !canUpdate {
		return nil, types.ErrNotPermittedToUpdateGlobalWhitelist
	}

	// Remove the address from the whitelist
	return &types.RemoveFromGlobalWhitelistResponse{}, ms.k.RemoveFromGlobalWhitelist(ctx, msg.Address)
}

func (ms msgServer) EnableTopicWorkerWhitelist(ctx context.Context, msg *types.EnableTopicWorkerWhitelistRequest) (_ *types.EnableTopicWorkerWhitelistResponse, err error) {
	defer metrics.RecordMetrics("EnableTopicWorkerWhitelist", time.Now(), &err)

	// Validate the address
	err = ms.k.ValidateStringIsBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	canUpdate, err := ms.k.CanUpdateTopicWhitelist(ctx, msg.TopicId, msg.Sender)
	if err != nil {
		return nil, err
	} else if !canUpdate {
		return nil, types.ErrNotPermittedToUpdateTopicWhitelist
	}

	return &types.EnableTopicWorkerWhitelistResponse{}, ms.k.EnableTopicWorkerWhitelist(ctx, msg.TopicId)
}

func (ms msgServer) DisableTopicWorkerWhitelist(ctx context.Context, msg *types.DisableTopicWorkerWhitelistRequest) (_ *types.DisableTopicWorkerWhitelistResponse, err error) {
	defer metrics.RecordMetrics("DisableTopicWorkerWhitelist", time.Now(), &err)

	// Validate the address
	err = ms.k.ValidateStringIsBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	canUpdate, err := ms.k.CanUpdateTopicWhitelist(ctx, msg.TopicId, msg.Sender)
	if err != nil {
		return nil, err
	} else if !canUpdate {
		return nil, types.ErrNotPermittedToUpdateTopicWhitelist
	}

	return &types.DisableTopicWorkerWhitelistResponse{}, ms.k.DisableTopicWorkerWhitelist(ctx, msg.TopicId)
}

func (ms msgServer) EnableTopicReputerWhitelist(ctx context.Context, msg *types.EnableTopicReputerWhitelistRequest) (_ *types.EnableTopicReputerWhitelistResponse, err error) {
	defer metrics.RecordMetrics("EnableTopicReputerWhitelist", time.Now(), &err)

	// Validate the address
	err = ms.k.ValidateStringIsBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	canUpdate, err := ms.k.CanUpdateTopicWhitelist(ctx, msg.TopicId, msg.Sender)
	if err != nil {
		return nil, err
	} else if !canUpdate {
		return nil, types.ErrNotPermittedToUpdateTopicWhitelist
	}

	return &types.EnableTopicReputerWhitelistResponse{}, ms.k.EnableTopicReputerWhitelist(ctx, msg.TopicId)
}

func (ms msgServer) DisableTopicReputerWhitelist(ctx context.Context, msg *types.DisableTopicReputerWhitelistRequest) (_ *types.DisableTopicReputerWhitelistResponse, err error) {
	defer metrics.RecordMetrics("DisableTopicReputerWhitelist", time.Now(), &err)

	// Validate the address
	err = ms.k.ValidateStringIsBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	canUpdate, err := ms.k.CanUpdateTopicWhitelist(ctx, msg.TopicId, msg.Sender)
	if err != nil {
		return nil, err
	} else if !canUpdate {
		return nil, types.ErrNotPermittedToUpdateTopicWhitelist
	}

	return &types.DisableTopicReputerWhitelistResponse{}, ms.k.DisableTopicReputerWhitelist(ctx, msg.TopicId)
}

func (ms msgServer) AddToTopicCreatorWhitelist(ctx context.Context, msg *types.AddToTopicCreatorWhitelistRequest) (_ *types.AddToTopicCreatorWhitelistResponse, err error) {
	defer metrics.RecordMetrics("AddToTopicCreatorWhitelist", time.Now(), &err)

	// Validate the address
	err = ms.k.ValidateStringIsBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	// Validate the address
	err = ms.k.ValidateStringIsBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	canUpdate, err := ms.k.CanUpdateTopicCreatorWhitelist(ctx, msg.Sender)
	if err != nil {
		return nil, err
	} else if !canUpdate {
		return nil, types.ErrNotPermittedToUpdateTopicCreatorWhitelist
	}

	return &types.AddToTopicCreatorWhitelistResponse{}, ms.k.AddToTopicCreatorWhitelist(ctx, msg.Address)
}

func (ms msgServer) RemoveFromTopicCreatorWhitelist(ctx context.Context, msg *types.RemoveFromTopicCreatorWhitelistRequest) (_ *types.RemoveFromTopicCreatorWhitelistResponse, err error) {
	defer metrics.RecordMetrics("RemoveFromTopicCreatorWhitelist", time.Now(), &err)

	// Validate the address
	err = ms.k.ValidateStringIsBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	// Validate the address
	err = ms.k.ValidateStringIsBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	canUpdate, err := ms.k.CanUpdateTopicCreatorWhitelist(ctx, msg.Sender)
	if err != nil {
		return nil, err
	} else if !canUpdate {
		return nil, types.ErrNotPermittedToUpdateTopicCreatorWhitelist
	}

	return &types.RemoveFromTopicCreatorWhitelistResponse{}, ms.k.RemoveFromTopicCreatorWhitelist(ctx, msg.Address)
}

func (ms msgServer) AddToTopicWorkerWhitelist(ctx context.Context, msg *types.AddToTopicWorkerWhitelistRequest) (_ *types.AddToTopicWorkerWhitelistResponse, err error) {
	defer metrics.RecordMetrics("AddToTopicWorkerWhitelist", time.Now(), &err)

	// Validate the sender
	err = ms.k.ValidateStringIsBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	// Validate the address
	err = ms.k.ValidateStringIsBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	canUpdate, err := ms.k.CanUpdateTopicWhitelist(ctx, msg.TopicId, msg.Sender)
	if err != nil {
		return nil, err
	} else if !canUpdate {
		return nil, types.ErrNotPermittedToUpdateTopicWorkerWhitelist
	}

	return &types.AddToTopicWorkerWhitelistResponse{}, ms.k.AddToTopicWorkerWhitelist(ctx, msg.TopicId, msg.Address)
}

func (ms msgServer) RemoveFromTopicWorkerWhitelist(ctx context.Context, msg *types.RemoveFromTopicWorkerWhitelistRequest) (_ *types.RemoveFromTopicWorkerWhitelistResponse, err error) {
	defer metrics.RecordMetrics("RemoveFromTopicWorkerWhitelist", time.Now(), &err)

	// Validate the sender
	err = ms.k.ValidateStringIsBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	// Validate the address
	err = ms.k.ValidateStringIsBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	canUpdate, err := ms.k.CanUpdateTopicWhitelist(ctx, msg.TopicId, msg.Sender)
	if err != nil {
		return nil, err
	} else if !canUpdate {
		return nil, types.ErrNotPermittedToUpdateTopicWorkerWhitelist
	}

	return &types.RemoveFromTopicWorkerWhitelistResponse{}, ms.k.RemoveFromTopicWorkerWhitelist(ctx, msg.TopicId, msg.Address)
}

func (ms msgServer) AddToTopicReputerWhitelist(ctx context.Context, msg *types.AddToTopicReputerWhitelistRequest) (_ *types.AddToTopicReputerWhitelistResponse, err error) {
	defer metrics.RecordMetrics("AddToTopicReputerWhitelist", time.Now(), &err)

	// Validate the sender
	err = ms.k.ValidateStringIsBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	// Validate the address
	err = ms.k.ValidateStringIsBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	canUpdate, err := ms.k.CanUpdateTopicWhitelist(ctx, msg.TopicId, msg.Sender)
	if err != nil {
		return nil, err
	} else if !canUpdate {
		return nil, types.ErrNotPermittedToUpdateTopicReputerWhitelist
	}

	return &types.AddToTopicReputerWhitelistResponse{}, ms.k.AddToTopicReputerWhitelist(ctx, msg.TopicId, msg.Address)
}

func (ms msgServer) RemoveFromTopicReputerWhitelist(ctx context.Context, msg *types.RemoveFromTopicReputerWhitelistRequest) (_ *types.RemoveFromTopicReputerWhitelistResponse, err error) {
	defer metrics.RecordMetrics("RemoveFromTopicReputerWhitelist", time.Now(), &err)

	// Validate the sender
	err = ms.k.ValidateStringIsBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	// Validate the address
	err = ms.k.ValidateStringIsBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	canUpdate, err := ms.k.CanUpdateTopicWhitelist(ctx, msg.TopicId, msg.Sender)
	if err != nil {
		return nil, err
	} else if !canUpdate {
		return nil, types.ErrNotPermittedToUpdateTopicReputerWhitelist
	}

	return &types.RemoveFromTopicReputerWhitelistResponse{}, ms.k.RemoveFromTopicReputerWhitelist(ctx, msg.TopicId, msg.Address)
}
