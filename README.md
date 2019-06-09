# ZInductionResListing
GoLang backend and API for Restaurant Listing at Zomato

### CRUD Operations for API testing on Postman - 

- [POST/CREATE](/images/POST.png) - [CODE](https://github.com/noiceAndToit/ZInductionResListing/blob/master/repository/post/post_mysql.go#L77)
- [GET/FETCH](/images/GET.png) - [CODE](https://github.com/noiceAndToit/ZInductionResListing/blob/master/repository/post/post_mysql.go#L21)
- [DELETE](/images/DELETE.png) - [CODE](https://github.com/noiceAndToit/ZInductionResListing/blob/master/repository/post/post_mysql.go#L120)
- [DB Status](/images/DB.png)


## API ENDPOINTS

### All Posts
- Path : `/posts`
- Method: `GET`
- Response: `200`

### Create Post
- Path : `/posts`
- Method: `POST`
- Fields: `name, cusine, distance, address, cft, rating, timing`
- Response: `201`

### Details a Post
- Path : `/posts/{id}`
- Method: `GET`
- Response: `200`

### Update Post
- Path : `/posts/{id}`
- Method: `PUT`
- Fields: `name, cusine, distance, address, cft, rating, timing`
- Response: `200`

### Delete Post
- Path : `/posts/{id}`
- Method: `DELETE`
- Response: `204`
