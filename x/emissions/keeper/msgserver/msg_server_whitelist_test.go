package msgserver_test

import (
	"github.com/allora-network/allora-chain/x/emissions/types"
)

func (s *MsgServerTestSuite) TestAddWhitelistAdmin() {
	ctx := s.ctx
	require := s.Require()
	msgServer := s.msgServer

	adminAddr := s.addrsStr[0]
	newAdminAddr := nonAdminAccounts[0].String()

	// Verify that newAdminAddr is not a whitelist admin
	isWhitelistAdmin, err := s.emissionsKeeper.IsWhitelistAdmin(ctx, newAdminAddr)
	require.NoError(err, "IsWhitelistAdmin should not return an error")
	require.False(isWhitelistAdmin, "newAdminAddr should not be a whitelist admin")

	// Attempt to add newAdminAddr to whitelist by adminAddr
	msg := &types.AddToWhitelistAdminRequest{
		Sender:  adminAddr,
		Address: newAdminAddr,
	}

	_, err = msgServer.AddToWhitelistAdmin(ctx, msg)
	require.NoError(err, "Adding to whitelist admin should succeed")

	// Verify that newAdminAddr is now a whitelist admin
	isWhitelistAdmin, err = s.emissionsKeeper.IsWhitelistAdmin(ctx, newAdminAddr)
	require.NoError(err, "IsWhitelistAdmin should not return an error")
	require.True(isWhitelistAdmin, "newAdminAddr should be a whitelist admin")
}

func (s *MsgServerTestSuite) TestAddWhitelistAdminInvalidUnauthorized() {
	ctx := s.ctx
	require := s.Require()

	nonAdminAddr := nonAdminAccounts[0]
	targetAddr := s.addrsStr[1]

	// Attempt to add targetAddr to whitelist by nonAdminAddr
	msg := &types.AddToWhitelistAdminRequest{
		Sender:  nonAdminAddr.String(),
		Address: targetAddr,
	}

	_, err := s.msgServer.AddToWhitelistAdmin(ctx, msg)
	require.ErrorIs(err, types.ErrNotPermittedToUpdateWhitelistAdmins, "Should fail due to unauthorized access")
}

func (s *MsgServerTestSuite) TestRemoveWhitelistAdmin() {
	ctx := s.ctx
	require := s.Require()
	msgServer := s.msgServer

	adminAddr := s.addrsStr[0]
	adminToRemove := s.addrsStr[1]

	// Attempt to remove adminToRemove from the whitelist by adminAddr
	removeMsg := &types.RemoveFromWhitelistAdminRequest{
		Sender:  adminAddr,
		Address: adminToRemove,
	}
	_, err := msgServer.RemoveFromWhitelistAdmin(ctx, removeMsg)
	require.NoError(err, "Removing from whitelist admin should succeed")

	// Verify that adminToRemove is no longer a whitelist admin
	isWhitelistAdmin, err := s.emissionsKeeper.IsWhitelistAdmin(ctx, adminToRemove)
	require.NoError(err, "IsWhitelistAdmin check should not return an error")
	require.False(isWhitelistAdmin, "adminToRemove should not be a whitelist admin anymore")
}

func (s *MsgServerTestSuite) TestRemoveWhitelistAdminInvalidUnauthorized() {
	ctx := s.ctx
	require := s.Require()

	nonAdminAddr := nonAdminAccounts[0]

	// Attempt to remove an admin from whitelist by nonAdminAddr
	msg := &types.RemoveFromWhitelistAdminRequest{
		Sender:  nonAdminAddr.String(),
		Address: s.addrsStr[0],
	}

	_, err := s.msgServer.RemoveFromWhitelistAdmin(ctx, msg)
	require.ErrorIs(err, types.ErrNotPermittedToUpdateWhitelistAdmins, "Should fail due to unauthorized access")
}

func (s *MsgServerTestSuite) TestAddToGlobalWhitelist() {
	ctx := s.ctx
	require := s.Require()
	msgServer := s.msgServer

	adminAddr := s.addrsStr[0]
	targetAddr := s.addrsStr[1]

	// Add targetAddr to global whitelist by adminAddr
	msg := &types.AddToGlobalWhitelistRequest{
		Sender:  adminAddr,
		Address: targetAddr,
	}
	_, err := msgServer.AddToGlobalWhitelist(ctx, msg)
	require.NoError(err, "Adding to global whitelist should succeed")

	// Verify targetAddr is now in global whitelist
	isWhitelisted, err := s.emissionsKeeper.IsWhitelistGlobalActor(ctx, targetAddr)
	require.NoError(err, "IsWhitelistGlobalActor check should not return an error")
	require.True(isWhitelisted, "targetAddr should be in global whitelist")
}

func (s *MsgServerTestSuite) TestAddToGlobalWhitelistInvalidUnauthorized() {
	ctx := s.ctx
	require := s.Require()

	nonAdminAddr := nonAdminAccounts[0]
	targetAddr := s.addrsStr[1]

	// Attempt to add targetAddr to global whitelist by nonAdminAddr
	msg := &types.AddToGlobalWhitelistRequest{
		Sender:  nonAdminAddr.String(),
		Address: targetAddr,
	}

	_, err := s.msgServer.AddToGlobalWhitelist(ctx, msg)
	require.ErrorIs(err, types.ErrNotPermittedToUpdateGlobalWhitelist, "Should fail due to unauthorized access")
}

func (s *MsgServerTestSuite) TestRemoveFromGlobalWhitelist() {
	ctx := s.ctx
	require := s.Require()
	msgServer := s.msgServer

	adminAddr := s.addrsStr[0]
	targetAddr := s.addrsStr[1]

	// First add targetAddr to global whitelist
	addMsg := &types.AddToGlobalWhitelistRequest{
		Sender:  adminAddr,
		Address: targetAddr,
	}
	_, err := msgServer.AddToGlobalWhitelist(ctx, addMsg)
	require.NoError(err, "Adding to global whitelist should succeed")

	// Remove targetAddr from global whitelist
	removeMsg := &types.RemoveFromGlobalWhitelistRequest{
		Sender:  adminAddr,
		Address: targetAddr,
	}
	_, err = msgServer.RemoveFromGlobalWhitelist(ctx, removeMsg)
	require.NoError(err, "Removing from global whitelist should succeed")

	// Verify targetAddr is no longer in global whitelist
	isWhitelisted, err := s.emissionsKeeper.IsWhitelistGlobalActor(ctx, targetAddr)
	require.NoError(err, "IsWhitelistGlobalActor check should not return an error")
	require.False(isWhitelisted, "targetAddr should not be in global whitelist")
}

func (s *MsgServerTestSuite) TestRemoveFromGlobalWhitelistInvalidUnauthorized() {
	ctx := s.ctx
	require := s.Require()

	nonAdminAddr := nonAdminAccounts[0]
	targetAddr := s.addrsStr[1]

	// Attempt to remove targetAddr from global whitelist by nonAdminAddr
	msg := &types.RemoveFromGlobalWhitelistRequest{
		Sender:  nonAdminAddr.String(),
		Address: targetAddr,
	}

	_, err := s.msgServer.RemoveFromGlobalWhitelist(ctx, msg)
	require.ErrorIs(err, types.ErrNotPermittedToUpdateGlobalWhitelist, "Should fail due to unauthorized access")
}

func (s *MsgServerTestSuite) TestEnableTopicWhitelist() {
	ctx := s.ctx
	require := s.Require()
	msgServer := s.msgServer
	adminAddr := s.addrsStr[0]
	topicId := s.CreateOneTopic().Id

	// Enable whitelist for topic
	msg := &types.EnableTopicWhitelistRequest{
		Sender:  adminAddr,
		TopicId: topicId,
	}
	_, err := msgServer.EnableTopicWhitelist(ctx, msg)
	require.NoError(err, "Enabling topic whitelist should succeed")

	// Verify topic whitelist is enabled
	isEnabled, err := s.emissionsKeeper.IsTopicWhitelistEnabled(ctx, topicId)
	require.NoError(err, "IsTopicWhitelistEnabled check should not return an error")
	require.True(isEnabled, "Topic whitelist should be enabled")
}

func (s *MsgServerTestSuite) TestEnableTopicWhitelistInvalidUnauthorized() {
	ctx := s.ctx
	require := s.Require()
	nonAdminAddr := nonAdminAccounts[0]
	topicId := s.CreateOneTopic().Id

	// Attempt to enable whitelist for topic by nonAdminAddr
	msg := &types.EnableTopicWhitelistRequest{
		Sender:  nonAdminAddr.String(),
		TopicId: topicId,
	}

	_, err := s.msgServer.EnableTopicWhitelist(ctx, msg)
	require.ErrorIs(err, types.ErrNotPermittedToUpdateTopicWhitelist, "Should fail due to unauthorized access")
}

func (s *MsgServerTestSuite) TestDisableTopicWhitelist() {
	ctx := s.ctx
	require := s.Require()
	msgServer := s.msgServer
	adminAddr := s.addrsStr[0]
	topicId := s.CreateOneTopic().Id

	// First enable whitelist for topic
	enableMsg := &types.EnableTopicWhitelistRequest{
		Sender:  adminAddr,
		TopicId: topicId,
	}
	_, err := msgServer.EnableTopicWhitelist(ctx, enableMsg)
	require.NoError(err, "Enabling topic whitelist should succeed")

	// Disable whitelist for topic
	disableMsg := &types.DisableTopicWhitelistRequest{
		Sender:  adminAddr,
		TopicId: topicId,
	}
	_, err = msgServer.DisableTopicWhitelist(ctx, disableMsg)
	require.NoError(err, "Disabling topic whitelist should succeed")

	// Verify topic whitelist is disabled
	isEnabled, err := s.emissionsKeeper.IsTopicWhitelistEnabled(ctx, topicId)
	require.NoError(err, "IsTopicWhitelistEnabled check should not return an error")
	require.False(isEnabled, "Topic whitelist should be disabled")
}

func (s *MsgServerTestSuite) TestDisableTopicWhitelistInvalidUnauthorized() {
	ctx := s.ctx
	require := s.Require()
	nonAdminAddr := nonAdminAccounts[0]
	topicId := s.CreateOneTopic().Id

	// Attempt to disable whitelist for topic by nonAdminAddr
	msg := &types.DisableTopicWhitelistRequest{
		Sender:  nonAdminAddr.String(),
		TopicId: topicId,
	}

	_, err := s.msgServer.DisableTopicWhitelist(ctx, msg)
	require.ErrorIs(err, types.ErrNotPermittedToUpdateTopicWhitelist, "Should fail due to unauthorized access")
}
