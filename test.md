# **API Coverage Tracking**

This issue will track the progress of API coverage for the Roblox API in this project. Some of these endpoints may not be implemented as they could have special requirements (captchas etc.). Undocumented and hidden API endpoints have not yet been included.

Progress Legend:

- 🔲 Not Started
- 🟡 In Progress
- ✅ Completed

---

## **1. Users API**

**Base URL:** `https://users.roblox.com/`  
**Documentation:** [Users API Docs](https://users.roblox.com/docs)

| Endpoint                                    | Suggested Method Name                                                             | HTTP Method | Status |
| ------------------------------------------- | --------------------------------------------------------------------------------- | ----------- | ------ |
| `/v1/birthdate`                             | `GetBirthdate(ctx context.Context)`                                               | `GET`       | 🔲     |
| `/v1/birthdate`                             | `UpdateBirthdate(ctx context.Context, birthdate *BirthdateBuilder)`               | `POST`      | 🔲     |
| `/v1/description`                           | `GetDescription(ctx context.Context)`                                             | `GET`       | 🔲     |
| `/v1/description`                           | `UpdateDescription(ctx context.Context, desc *DescriptionBuilder)`                | `POST`      | 🔲     |
| `/v1/gender`                                | `GetGender(ctx context.Context)`                                                  | `GET`       | 🔲     |
| `/v1/gender`                                | `UpdateGender(ctx context.Context, gender *GenderBuilder)`                        | `POST`      | 🔲     |
| `/v1/display-names/validate`                | `ValidateDisplayName(ctx context.Context, b *ValidateDisplayNameBuilder)`         | `GET`       | 🔲     |
| `/v1/users/{userId}/display-names/validate` | `ValidateUserDisplayName(ctx context.Context, b *ValidateUserDisplayNameBuilder)` | `GET`       | 🔲     |
| `/v1/users/{userId}/display-names`          | `SetDisplayName(ctx context.Context, b *SetDisplayNameBuilder)`                   | `PATCH`     | 🔲     |
| `/v1/users/{userId}`                        | `GetUserByID(ctx context.Context, userID uint64)`                                 | `GET`       | ✅     |
| `/v1/users/authenticated`                   | `GetAuthUser(ctx context.Context)`                                                | `GET`       | ✅     |
| `/v1/users/authenticated/age-bracket`       | `GetAuthUserAgeBracket(ctx context.Context)`                                      | `GET`       | 🔲     |
| `/v1/users/authenticated/country-code`      | `GetAuthUserCountryCode(ctx context.Context)`                                     | `GET`       | 🔲     |
| `/v1/users/authenticated/roles`             | `GetAuthUserRoles(ctx context.Context)`                                           | `GET`       | 🔲     |
| `/v1/usernames/users`                       | `GetUsersByUsernames(ctx context.Context, b *UsersByUsernamesBuilder)`            | `POST`      | ✅     |
| `/v1/users`                                 | `GetUsersByIDs(ctx context.Context, b *UsersByIDsBuilder)`                        | `POST`      | ✅     |
| `/v1/users/{userId}/username-history`       | `GetUsernameHistory(ctx context.Context, b *UsernameHistoryBuilder)`              | `GET`       | ✅     |
| `/v1/users/search`                          | `SearchUsers(ctx context.Context, b *SearchUsersBuilder)`                         | `GET`       | ✅     |

---

## **2. Friends API**

**Base URL:** `https://friends.roblox.com/`  
**Documentation:** [Friends API Docs](https://friends.roblox.com/docs)

| Endpoint                                                    | Suggested Method Name                                                                   | HTTP Method | Status |
| ----------------------------------------------------------- | --------------------------------------------------------------------------------------- | ----------- | ------ |
| `/v1/metadata`                                              | `GetFriendsMetadata(ctx context.Context, b *FriendsMetadataBuilder)`                    | `GET`       | 🔲     |
| `/v1/my/friends/count`                                      | `GetMyFriendsCount(ctx context.Context)`                                                | `GET`       | 🔲     |
| `/v1/my/friends/requests`                                   | `GetMyFriendRequests(ctx context.Context, b *FriendRequestsBuilder)`                    | `GET`       | 🔲     |
| `/v1/user/friend-requests/count`                            | `GetFriendRequestsCount(ctx context.Context)`                                           | `GET`       | 🔲     |
| `/v1/users/{userId}/friends`                                | `GetUserFriends(ctx context.Context4, b *UserFriendsBuilder)`                           | `GET`       | 🔲     |
| `/v1/users/{userId}/friends/count`                          | `GetUserFriendsCount(ctx context.Context, userId uint64)`                               | `GET`       | 🔲     |
| `/v1/users/{userId}/friends/find`                           | `FindUserFriends(ctx context.Context, b *FindFriendsBuilder)`                           | `GET`       | 🔲     |
| `/v1/users/{userId}/friends/inactive`                       | `GetInactiveFriends(ctx context.Context, userId uint64)`                                | `GET`       | 🔲     |
| `/v1/users/{userId}/friends/online`                         | `GetOnlineFriends(ctx context.Context, b *OnlineFriendsBuilder)`                        | `GET`       | 🔲     |
| `/v1/users/{userId}/friends/search`                         | `SearchFriends(ctx context.Context, b *SearchFriendsBuilder)`                           | `GET`       | 🔲     |
| `/v1/users/{userId}/friends/statuses`                       | `GetFriendStatuses(ctx context.Context, b *FriendStatusesBuilder)`                      | `GET`       | 🔲     |
| `/v1/contacts/{targetContactId}/request-friendship`         | `RequestFriendshipByContact(ctx context.Context, targetContactId string)`               | `POST`      | 🔲     |
| `/v1/user/friend-requests/decline-all`                      | `DeclineAllFriendRequests(ctx context.Context)`                                         | `POST`      | 🔲     |
| `/v1/user/multiget-are-friends`                             | `CheckMultiAreFriends(ctx context.Context, b *MultiAreFriendsBuilder)`                  | `POST`      | 🔲     |
| `/v1/users/{requesterUserId}/accept-friend-request`         | `AcceptFriendRequest(ctx context.Context, requesterUserId uint64)`                      | `POST`      | 🔲     |
| `/v1/users/{requesterUserId}/decline-friend-request`        | `DeclineFriendRequest(ctx context.Context, requesterUserId uint64)`                     | `POST`      | 🔲     |
| `/v1/users/{senderUserId}/accept-friend-request-with-token` | `AcceptFriendRequestWithToken(ctx context.Context, b *AcceptFriendRequestTokenBuilder)` | `POST`      | 🔲     |
| `/v1/users/{targetUserId}/request-friendship`               | `RequestFriendship(ctx context.Context, b *FriendshipRequestBuilder)`                   | `POST`      | 🔲     |
| `/v1/users/{targetUserId}/unfriend`                         | `Unfriend(ctx context.Context, targetUserId uint64)`                                    | `POST`      | 🔲     |
| `/v1/users/{targetUserId}/followers`                        | `GetFollowers(ctx context.Context, b *FollowersBuilder)`                                | `GET`       | 🔲     |
| `/v1/users/{targetUserId}/followers/count`                  | `GetFollowersCount(ctx context.Context, targetUserId uint64)`                           | `GET`       | 🔲     |
| `/v1/users/{targetUserId}/followings`                       | `GetFollowings(ctx context.Context, b *FollowingsBuilder)`                              | `GET`       | 🔲     |
| `/v1/users/{targetUserId}/followings/count`                 | `GetFollowingsCount(ctx context.Context, targetUserId uint64)`                          | `GET`       | 🔲     |
| `/v1/user/following-exists`                                 | `CheckFollowingExists(ctx context.Context, b *FollowingExistsBuilder)`                  | `POST`      | 🔲     |
| `/v1/users/{targetUserId}/follow`                           | `FollowUser(ctx context.Context, b *FollowUserBuilder)`                                 | `POST`      | 🔲     |
| `/v1/users/{targetUserId}/unfollow`                         | `UnfollowUser(ctx context.Context, targetUserId uint64)`                                | `POST`      | 🔲     |

---

## **3. Games API**

**Base URL:** `https://games.roblox.com/`  
**Documentation:** [Games API Docs](https://games.roblox.com/docs)

| Endpoint                                               | Suggested Method Name                                                                       | HTTP Method | Status |
| ------------------------------------------------------ | ------------------------------------------------------------------------------------------- | ----------- | ------ |
| `/v1/games`                                            | `GetGamesDetails(ctx context.Context, b *GamesDetailsBuilder)`                              | `GET`       | 🔲     |
| `/v1/games/games-product-info`                         | `GetGamesProductInfo(ctx context.Context, b *GamesProductInfoBuilder)`                      | `GET`       | 🔲     |
| `/v1/games/list-spotlight`                             | `GetGamesSpotlight(ctx context.Context)`                                                    | `GET`       | 🔲     |
| `/v1/games/multiget-place-details`                     | `GetMultiplePlaceDetails(ctx context.Context, b *MultiplePlaceDetailsBuilder)`              | `GET`       | 🔲     |
| `/v1/games/multiget-playability-status`                | `GetMultiplePlayabilityStatus(ctx context.Context, b *MultiplePlayabilityStatusBuilder)`    | `GET`       | 🔲     |
| `/v1/games/recommendations/game/{universeId}`          | `GetGameRecommendations(ctx context.Context, b *GameRecommendationsBuilder)`                | `GET`       | 🔲     |
| `/v1/games/recommendations/algorithm/{algorithmName}`  | `GetAlgorithmRecommendations(ctx context.Context, b *AlgorithmRecommendationsBuilder)`      | `GET`       | 🔲     |
| `/v1/games/{placeId}/private-servers`                  | `GetPrivateServers(ctx context.Context, b *PrivateServersBuilder)`                          | `GET`       | 🔲     |
| `/v1/games/{placeId}/servers/{serverType}`             | `GetGameServers(ctx context.Context, b *GameServersBuilder)`                                | `GET`       | 🔲     |
| `/v1/games/{universeId}/favorites`                     | `GetGameFavoriteStatus(ctx context.Context, universeId uint64)`                             | `GET`       | 🔲     |
| `/v1/games/{universeId}/favorites`                     | `SetGameFavoriteStatus(ctx context.Context, b *SetFavoriteStatusBuilder)`                   | `POST`      | 🔲     |
| `/v1/games/{universeId}/favorites/count`               | `GetGameFavoritesCount(ctx context.Context, universeId uint64)`                             | `GET`       | 🔲     |
| `/v1/games/{universeId}/game-passes`                   | `GetGamePasses(ctx context.Context, b *GamePassesBuilder)`                                  | `GET`       | 🔲     |
| `/v1/games/{universeId}/votes/user`                    | `GetUserVoteStatus(ctx context.Context, universeId uint64)`                                 | `GET`       | 🔲     |
| `/v1/games/votes`                                      | `GetMultipleGameVotes(ctx context.Context, b *MultipleGameVotesBuilder)`                    | `GET`       | 🔲     |
| `/v1/games/{universeId}/user-votes`                    | `SetUserVote(ctx context.Context, b *SetUserVoteBuilder)`                                   | `PATCH`     | 🔲     |
| `/v1/private-servers`                                  | `GetPrivateServers(ctx context.Context, b *PrivateServersBuilder)`                          | `GET`       | 🔲     |
| `/v1/private-servers/enabled-in-universe/{universeId}` | `GetPrivateServersEnabled(ctx context.Context, universeId uint64)`                          | `GET`       | 🔲     |
| `/v1/private-servers/my-private-servers`               | `GetMyPrivateServers(ctx context.Context, b *MyPrivateServersBuilder)`                      | `GET`       | 🔲     |
| `/v1/vip-server/can-invite/{userId}`                   | `CanInviteToVipServer(ctx context.Context, userId uint64)`                                  | `GET`       | 🔲     |
| `/v1/vip-servers/{id}`                                 | `GetVipServerInfo(ctx context.Context, id uint64)`                                          | `GET`       | 🔲     |
| `/v1/vip-servers/{id}`                                 | `UpdateVipServer(ctx context.Context, b *UpdateVipServerBuilder)`                           | `PATCH`     | 🔲     |
| `/v1/games/vip-servers/{universeId}`                   | `CreateVipServer(ctx context.Context, b *CreateVipServerBuilder)`                           | `POST`      | 🔲     |
| `/v1/vip-servers/{id}/permissions`                     | `UpdateVipServerPermissions(ctx context.Context, b *UpdateVipServerPermissionsBuilder)`     | `PATCH`     | 🔲     |
| `/v1/vip-servers/{id}/subscription`                    | `UpdateVipServerSubscription(ctx context.Context, b *UpdateVipServerSubscriptionBuilder)`   | `PATCH`     | 🔲     |
| `/v1/vip-servers/{id}/voicesettings`                   | `UpdateVipServerVoiceSettings(ctx context.Context, b *UpdateVipServerVoiceSettingsBuilder)` | `PATCH`     | 🔲     |
| `/v2/games/{universeId}/media`                         | `GetGameMedia(ctx context.Context, universeId uint64)`                                      | `GET`       | 🔲     |
| `/v2/groups/{groupId}/games`                           | `GetGroupGames(ctx context.Context, b *GroupGamesV2Builder)`                                | `GET`       | 🔲     |
| `/v2/groups/{groupId}/gamesV2`                         | `GetGroupGamesAlt(ctx context.Context, b *GroupGamesV2AltBuilder)`                          | `GET`       | 🔲     |
| `/v2/users/{userId}/games`                             | `GetUserGames(ctx context.Context, b *UserGamesV2Builder)`                                  | `GET`       | 🔲     |

---

## **4. Inventory API**

**Base URL:** `https://inventory.roblox.com/`  
**Documentation:** [Inventory API Docs](https://inventory.roblox.com/docs)

| Endpoint                                                      | Suggested Method Name                                                                  | HTTP Method | Status |
| ------------------------------------------------------------- | -------------------------------------------------------------------------------------- | ----------- | ------ |
| `/v1/users/{userId}/assets/collectibles`                      | `GetUserCollectibles(ctx context.Context, b *CollectiblesBuilder)`                     | `GET`       | 🔲     |
| `/v1/users/{userId}/can-view-inventory`                       | `CanViewUserInventory(ctx context.Context, userId uint64)`                             | `GET`       | 🔲     |
| `/v1/users/{userId}/categories`                               | `GetUserCategories(ctx context.Context, userId uint64)`                                | `GET`       | 🔲     |
| `/v1/users/{userId}/categories/favorites`                     | `GetUserFavoriteCategories(ctx context.Context, userId uint64)`                        | `GET`       | 🔲     |
| `/v1/users/{userId}/items/{itemType}/{itemTargetId}`          | `GetUserItems(ctx context.Context, b *GetUserItemsBuilder)`                            | `GET`       | 🔲     |
| `/v1/users/{userId}/items/{itemType}/{itemTargetId}/is-owned` | `IsItemOwned(ctx context.Context, b *IsItemOwnedBuilder)`                              | `GET`       | 🔲     |
| `/v1/collections/items/{itemType}/{itemTargetId}`             | `RemoveFromCollection(ctx context.Context, b *RemoveFromCollectionBuilder)`            | `DELETE`    | 🔲     |
| `/v1/collections/items/{itemType}/{itemTargetId}`             | `AddToCollection(ctx context.Context, itemType int, itemTargetId uint64)`              | `POST`      | 🔲     |
| `/v2/assets/{assetId}/owners`                                 | `GetAssetOwners(ctx context.Context, b *AssetOwnersBuilder)`                           | `GET`       | 🔲     |
| `/v2/users/{userId}/inventory`                                | `GetUserInventory(ctx context.Context, b *UserInventoryBuilder)`                       | `GET`       | 🔲     |
| `/v2/users/{userId}/inventory/{assetTypeId}`                  | `GetUserInventoryByAssetType(ctx context.Context, b *UserInventoryByAssetTypeBuilder)` | `GET`       | 🔲     |
| `/v2/inventory/asset/{assetId}`                               | `DeleteUserAsset(ctx context.Context, assetId uint64)`                                 | `DELETE`    | 🔲     |
