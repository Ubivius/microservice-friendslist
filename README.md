# microservice-friendslist
Friends list microservice for our online game framework.

**Get Friends by UserID**
----
  Returns all friend relationships of the specific user.

* **URL**

  /friends/:user_id

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
   `user_id=[integer]`

* **Data Params**

  None


**Get Invites By UserID**
----
  Resends all friend invitations for the specific user.

* **URL**

  /invites/:user_id

* **Method:**

  `GET`
  
*  **URL Params**

   **Required:**
 
   `user_id=[integer]`

* **Data Params**

  None


**Add new relationship**
----
  Add new relationship with specific data

* **URL**

  /relationships

* **Method:**

  `POST`
  
*  **URL Params**
 
  None

* **Data Params**

  **Required:**

  `user_1=[User]`
  `user_2=[User]`

  User
  `user_id=[integer]`
  `relationship_type=[integer]`


**Update relationship**
----
  Update relationship data

* **URL**

  /relationships

* **Method:**

  `PUT`
  
*  **URL Params**
 
  None

* **Data Params**

  **Required:**

  `id=[integer]`
  `user_1=[User]`
  `user_2=[User]`

  User
  `user_id=[integer]`
  `relationship_type=[integer]`


**Delete relationship**
----
  Delete relationship

* **URL**

  /relationships/:id

* **Method:**

  `DELETE`
  
*  **URL Params**
 
  **Required:**
 
   `id=[integer]`

* **Data Params**

  None
