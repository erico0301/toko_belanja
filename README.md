# toko_belanja
Final Project 4 - Scalable Web Services with Golang - Hacktiv8

# Cara Penggunaan
URL Heroku :  https://toko-belanja-emf.herokuapp.com/

Endpoint List :

Users

POST/users/register

[POST]http://localhost:8080/users/register

  Request
  
   body:
    
   ![image](https://user-images.githubusercontent.com/92410169/150880933-0299e814-49df-4058-bfb1-2cd28ad7c3d8.png)

  Response:
   
   status 201
    
   data:
    
   ![image](https://user-images.githubusercontent.com/92410169/150882071-5c349c9f-73d3-4110-a2ae-114f541a0a21.png)
    

POST/users/login

[POST]http://localhost:8080/users/login

  Request
  
   body:
   
   ![image](https://user-images.githubusercontent.com/92410169/150882540-44ed34f3-e1de-4983-bf54-8f9e7cf8390a.png)

  Response:
   
   status 200
    
   data:
    
    

    
PATCH/users/topup

[PATCH]http://localhost:8080/users/topup

  Request
  
   headers:
   
   ![image](https://user-images.githubusercontent.com/92410169/150883387-ccb72e06-fedc-43cd-9fa2-81e367324a79.png)
  
   body:
   
   ![image](https://user-images.githubusercontent.com/92410169/150883544-71c2c041-5422-4f37-b29c-8ce2c44acacf.png)

  Response:
   
   status 200
    
   data:
    
   ![image](https://user-images.githubusercontent.com/92410169/150884248-9800150a-c6a5-4c54-b770-e0f48be4a4ac.png)



Categories

POST/categories

[POST]http://localhost:8080/categories

  Request
  
   headers:
   
  ![image](https://user-images.githubusercontent.com/92410169/150884726-67d43be7-6e4b-4c53-916b-cc72f9c11c1d.png) 
  
   body:
   
   ![image](https://user-images.githubusercontent.com/92410169/150885338-21d06610-9487-467e-b03d-07fca87b8a9d.png)
   
  Response:
   
   status 201
    
   data:
   
   ![image](https://user-images.githubusercontent.com/92410169/150886002-cd6cb9bc-d6dd-4f27-b41d-4f0eedfddf5e.png)


GET/categories

[GET]http://localhost:8080/categories

  Request
  
   headers:
   
  ![image](https://user-images.githubusercontent.com/92410169/150884726-67d43be7-6e4b-4c53-916b-cc72f9c11c1d.png) 
   
  Response:
   
   status 200
    
   data:
   
   ![image](https://user-images.githubusercontent.com/92410169/150887882-bc552b3a-ddd8-408f-b486-b56d20f8e5d0.png)
   
   ![image](https://user-images.githubusercontent.com/92410169/150888085-20259a21-bc13-4765-8c2e-da963367067f.png)
   
   
 PATCH/categories/:categoryId

 [PATCH]http://localhost:8080/categories/:categoryId

  Request
  
   headers:
   
  ![image](https://user-images.githubusercontent.com/92410169/150884726-67d43be7-6e4b-4c53-916b-cc72f9c11c1d.png) 
  
   params:
   
   ![image](https://user-images.githubusercontent.com/92410169/150888845-ea3540b1-0c06-4514-97f9-b5c5bcf30442.png)
  
   body:
   
   ![image](https://user-images.githubusercontent.com/92410169/150889106-5fff16c5-6632-4ade-813e-86f7e11180e0.png)
   
  Response:
   
   status 200
    
   data:
   
   ![image](https://user-images.githubusercontent.com/92410169/150889449-00657bd0-ed73-494e-bdbf-15790b0ba07f.png)
   

 DELETE/categories/:categoryId

 [DELETE]http://localhost:8080/categories/:categoryId

  Request
  
   headers:
   
  ![image](https://user-images.githubusercontent.com/92410169/150884726-67d43be7-6e4b-4c53-916b-cc72f9c11c1d.png) 
  
   params:
   
   ![image](https://user-images.githubusercontent.com/92410169/150888845-ea3540b1-0c06-4514-97f9-b5c5bcf30442.png)
  
  Response:
   
   status 200
    
   data:
   
   ![image](https://user-images.githubusercontent.com/92410169/150890303-0cbead6f-6e66-4aa3-aa45-23ba52931feb.png)



Products

POST/products

[POST]http://localhost:8080/products

  Request
  
   headers:
   
  ![image](https://user-images.githubusercontent.com/92410169/150884726-67d43be7-6e4b-4c53-916b-cc72f9c11c1d.png) 
  
   body:
   
   ![image](https://user-images.githubusercontent.com/92410169/150890715-a2f4adfe-b117-47e9-a899-d008a31f4cd4.png)
   
  Response:
   
   status 201
    
   data:
   
   ![image](https://user-images.githubusercontent.com/92410169/150891013-e5d7d5c3-c847-4304-993e-8e2dfd020f75.png)
   
   
GET/products

[GET]http://localhost:8080/products

  Request
  
   headers:
   
  ![image](https://user-images.githubusercontent.com/92410169/150884726-67d43be7-6e4b-4c53-916b-cc72f9c11c1d.png) 
   
  Response:
   
   status 200
    
   data:
   
   ![image](https://user-images.githubusercontent.com/92410169/150891573-8908dcdb-1e79-4319-b1ee-f417b4f42361.png)

 
 PUT/products/:productsId

 [PUT]http://localhost:8080/products/:productsId

  Request
  
   headers:
   
  ![image](https://user-images.githubusercontent.com/92410169/150884726-67d43be7-6e4b-4c53-916b-cc72f9c11c1d.png) 
  
   params:
   
   ![image](https://user-images.githubusercontent.com/92410169/150888845-ea3540b1-0c06-4514-97f9-b5c5bcf30442.png)
   
   body:
   
   ![image](https://user-images.githubusercontent.com/92410169/150892449-4c348a85-7352-4d37-963e-a055bec2583c.png)
  
  Response:
   
   status 200
    
   data:
   
   ![image](https://user-images.githubusercontent.com/92410169/150893013-c9a9433a-b4ce-4ef7-bf3f-e1d8960212c7.png)


 DELETE/products/:productsId

 [DELETE]http://localhost:8080/products/:productsId

  Request
  
   headers:
   
  ![image](https://user-images.githubusercontent.com/92410169/150884726-67d43be7-6e4b-4c53-916b-cc72f9c11c1d.png) 
  
   params:
   
   ![image](https://user-images.githubusercontent.com/92410169/150888845-ea3540b1-0c06-4514-97f9-b5c5bcf30442.png)
  
  Response:
   
   status 200
    
   data:
   
   ![image](https://user-images.githubusercontent.com/92410169/150893381-288c8699-f588-40d5-9aa7-2d0142f8a2c0.png)



TransactionHistories

POST/transactions

[POST]http://localhost:8080/transactions

  Request
  
   headers:
   
  ![image](https://user-images.githubusercontent.com/92410169/150884726-67d43be7-6e4b-4c53-916b-cc72f9c11c1d.png) 
  
   body:
   
   ![image](https://user-images.githubusercontent.com/92410169/150893816-d837d3ea-8951-49f9-b62b-df9232a24cde.png)
   
  Response:
   
   status 201
    
   data:
   
   ![image](https://user-images.githubusercontent.com/92410169/150894385-4652b7ba-3715-4ffd-8865-0e1cb019f29d.png)


GET/transactions/my-transactions

[GET]http://localhost:8080/transactions/my-transactions

  Request
  
   headers:
   
  ![image](https://user-images.githubusercontent.com/92410169/150884726-67d43be7-6e4b-4c53-916b-cc72f9c11c1d.png) 
   
  Response:
   
   status 200
    
   data:
   
   ![image](https://user-images.githubusercontent.com/92410169/150895333-fd03c910-3cfc-46fe-b034-31a4ff96efe8.png)

