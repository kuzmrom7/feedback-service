FORMAT: 1A

# Feedback-service API
API for McDonald's reviews on DC with DC parser 

## GET /
+ Response 200 (text/plain)

        Server started!
        
        
## GET /reviews
This method return reviews

### Request

+ Query Parameters

    + page (number)
        + Default: `1`
        
    + sort (string)
        + rate - sorting by date 
        + rating - sorting by rating
        + Default: `0`
         
    + Answers (boolean)
        + true - for reviews only with a answers 
        
### Response

+ Response 200 (application/json)

    + Body

            {
                "status": true,
                "message": "success",
                "data": {
                        reviews: [
                            {
                                "id": "d16d95ac-6dfe-40c0-a376-a5cfa6658f8b",
                                "author": "Erik",
                                "answers": [],
                                "body": "Спасибо, все привезли быстро, еда была теплой)",
                                "orderHash": "8728b37cd456862699ee558c3cb14c393d8b20d5",
                                "rated": "2019-11-26T19:12:54+0300",
                                "rating": 5,
                                "created": "2019-11-26T16:23:33.53365Z",
                                "updated": "2019-11-26T16:23:33.53365Z"
                            },
                            {
                                "id": "808d89f2-42ce-4894-aefa-58e53ee8cc5b",
                                "author": "Елена",
                                "answers": [
                                {
                                "answer": "Елена, благодарим Вас за отзыв. Приносим наши искренние извинения за не полностью доставленный заказ. Мы приложим все усилия, чтобы подобной ситуации в дальнейшем не повторилось. Ждём Ваших новых заказов.",
                                "sourceId": "dc"
                                }
                                ],
                                "body": "Не положили соус(",
                                "orderHash": "5a6f9c413de54581ad1fec0a78e649f98a825b7b",
                                "rated": "2019-11-26T15:11:50+0300",
                                "rating": 3,
                                "created": "2019-11-26T16:23:33.53365Z",
                                "updated": "2019-11-26T16:23:33.53365Z"
                            },
                           ]
                        "totalPages": 61,
                        "currentPage": 1,
                        "nextPage": 2 // if nextPage = -1 it is last page
                     }
             }
    
