# offTime REST API

## To-do

- [x] v4 UUIDs
- [x] Image serving/uploading
- [x] websockets

## design-doc

### Features

### REST Routes

- [x] GET       /users
- [x] GET       /users/{username}
- [x] PUT       /users/{username}
- [x] PATCH     /users/{username}
- [x] DELETE    /users/{username}
- [x] PUT       /users/{username}/picture
- [ ] GET       /users/{username}/usageHistory 
- [ ] POST      /users/{username}/usageHistory 
- [ ] DELETE    /users/{username}/usageHistory 
- [x] GET    /users/{username}/roomHistory 
- [x] GET       /rooms
- [x] GET       /rooms/{roomID}

### WebSocket events

- [x] createRoom
- [x] joinRoom
- [x] updateRoomUsage

### CRUDs

- [x] GetAllUsers
- [x] GetUser
- [x] CreateUser
- [x] UpdateUser
- [x] DeleteUser

- [ ] GetAllUsageHistoryOfUser
- [ ] CreateUsageHistoryForUser
- [ ] DeleteAllUsageHistoryForUser

- [x] GetAllRooms
- [x] GetRoom
- [x] CreateRoom
- [x] UpdateRoom

## dev-log