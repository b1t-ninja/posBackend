# posBackend
This is the corrosponding backend for the pos system. 
It uses a Notion integration, to interact with the Database, as well as the Gin framework.

To speed up the semi slow NotionAPI query, I have implemented a cache, which will be updated either automatically after every request, or via an endpoint. <br>
This refresh endpoint is triggered by the frontend every time the user confirms the order, so that the data remains fresh. Also since I don't expect anyone updating their inventory multiple times a day, especially not system gastronomies that have huge orgaisational overhead, I suppose stale data isn't really a problem anyway. <br>
But better safe than sorry I always say.

The existing endpoints are:
1. GET localhost:8080/notionDB (get all entries from the DB)
2. GET localhost:8080/refreshCache (refresh the cache)

## Uage 
- Add .env file with your Notion Integration credentials
- Build the project via `go build`
- Run the server

## Example response to 1.
```json
[
    {
        "ingredients": [
            {
                "name": "peach",
                "quantity": 3
            },
            {
                "name": " coconut",
                "quantity": 10
            }
        ],
        "name": "Peachbeach Juice",
        "picture": "superlongimageurl",
        "price": 10,
        "size": "Large"
    },
    {
        "ingredients": [
            {
                "name": "peach",
                "quantity": 6
            },
            {
                "name": " coconut",
                "quantity": 4
            }
        ],
        "name": "Peachbeach Juice",
        "picture": "superlongimageurl",
        "price": 8,
        "size": "Medium"
    }
]
```
As a side note, if build for release, the response time thanks to the cache is quite impressive.
<img width="957" alt="Screenshot 2024-12-30 at 13 27 34" src="https://github.com/user-attachments/assets/dfff2e50-ba8a-4821-bad6-b00046a0e809" />
I was definetly surprised, especailly when comparing it with the time it takes to get that data from Notion <br>
<img width="959" alt="Screenshot 2024-12-30 at 13 29 08" src="https://github.com/user-attachments/assets/34f5ba1c-86a7-47cd-8426-0bb057c2389c" />
The downside of this cache is that more memory is consumed as it stays within memory even after the request. So if this was a large data set, we might run into a problem.
