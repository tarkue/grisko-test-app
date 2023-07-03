# grisko-test-app

# localhost:8080/products
## GET 
Если не указывать ничего -- будет выдавать весь список, если указывать ключ-значение, то будет отбирать по ним

<code>localhost:8080/products</code>
<code>localhost:8080/products?name=foo</code>

Возвращает список из Product
## POST
### Создаёт товар.
При учёте что есть поле name, остальные поля можно и не заполнять

<code>localhost:8080/products?name=foo&info=bar</code>

Возвращает идентификатор созданного товара
## DELETE
### Удаляет по id

<code>localhost:8080/products?id=1</code>

Возвращает удалённый объект Product
## PUT
### Изменяет по id  
Чтобы описать изменения необходимо прописать их как JSON-Объект в update

<code>localhost:8080/products?id=1&update={"name":"name", "info":"info"}</code>

Возвращает объект Product


# Product

### Имеет поля
* ### ID    - идентификатор
* ### Name  - название товара
* ### Info  - описание товара
* ### Img   - фотография
* ### Price - цена
