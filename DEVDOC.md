# offTime REST API

## To-do

- [ ] v4 UUIDs
- [x] Image serving/uploading

## design-doc

### Features

### Paths

- [x] GET       /users
- [x] GET       /users/{username}
- [x] PUT       /users/{username}
- [x] PATCH     /users/{username}
- [x] DELETE    /users/{username}
- [x] PUT       /users/{username}/picture
- [ ] GET       /users/{username}/usageHistory 
- [ ] POST      /users/{username}/usageHistory 
- [ ] DELETE    /users/{username}/usageHistory 
- [ ] GET       /rooms
- [ ] GET       /rooms/{roomID}

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