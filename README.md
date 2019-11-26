# Тестовое задание Golang

## Задача

##### Необходимо разработать приложение, которое выполняет две задачи:
1) Стянуть все отзывы с ​https://www.delivery-club.ru/srv/Mcdonalds_msk/feedbacks ​и записать в любую СУБД на выбор, необходимо записывать текст отзыва, его, рейтинг, дату/время создания и данные пользователя, который оставил отзыв, а также не стоит забывать про то, что у отзывов есть ответы.

2) Отдавать JSON для гипотетического Reactjs-приложения с возможностью фильтрации и сортировки отзывов по рейтингу или дате, предусмотреть пагинатор с автоподгрузкой.


Приложение может использовать любую библиотеку для парсинга, каким образом тянуть данные отдается на вкус выполняющего тестовое задание, но обязательно стоит помнить о том, что чем дешевле парсинг, тем лучше.

При повторном запуске парсинга необходимо добавлять новые отзывы, а также обновлять старые. Важно: обновлять только то, что действительно изменилось, не стоит забывать про колонки created_at и updated_at.

Обязательно подумать про стратегии кеширования API-роута, по возможности что-то реализовать и обязательно приготовиться к рассказу, как сделать кеширование лучше для кейса: есть 500 одновременных запросов с разными параметрами фильтрации и сервер с 512 Mb RAM.

API необходимо задокументировать с помощью ​https://apiblueprint.org/.​


## Docs

- [API Docs](./docs/api.md)

- [How to run](./docs/instructions.md)