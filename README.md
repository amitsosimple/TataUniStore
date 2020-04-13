REST APIs
Create REST API which has

POST endpoint to create products (one product at a time)

GET endpoint to fetch product details

PUT end point to update product details

All above APIs should do basic validation and return appropriate responses code with the responses

Create another REST API which has

POST endpoint to create products price (one product at a time)

GET endpoint to fetch product price (for a product)

Create aggregate API to return product and price details by making calls to above APIs. Make sure that APIs are not called sequentially.

Sample data
Product data {“product_id”:“”,“sellerId”:“”,“title”:“”,“manufacturer”:“”,“isLowQuantity”:false,“isSoldOut”:false,“isBackorder”:false,“metafields”:[{“key”:“Capacity”,“value”:“”},{“key”:“”,“value”:“”}],“requiresShipping”:true,“isVisible”:true,“publishedAt”:{“$date”:“2020-02-12T08:05:39.743Z”},“createdAt”:{“$date”:“2010-08-23T05:53:16.134Z”},“updatedAt”:{“$date”:“2019-08-23T05:53:16.134Z”},“workflow”:{“status”:“new”}}

Price data

{“product_id”:“1212323",“price”:{“range”:“4.00-5.00”,“min”:4.50,“max”:5.00}}

Other instructions
Basic unit tests should be written for all methods.

Logging should be in place with different logging levels

_________________________________________________________________________________________________

Solution Strategy:
1. Create two services one for ProductCatalog and one for Pricing.
2. As these services are tightly bound, So the backhand data should be same, this will help maintainability.
3. For the fast solution, let have a file based data base on local machine. 
4. As the problem statement tells we have one production display method, I am assuming that this will display all the information of the products, which include the complete catalog with the pricing information.
5. Other APIs are very much for the specific service data.
6. Top of these server, we have a agreegator service which call the server APIs and display the result on the console.
7. Following are the APIs and their responsibilities:
-> 
