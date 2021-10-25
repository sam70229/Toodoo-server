`GET`
> /api/v1/todos

Parameters

Name		|Type    |Description                                              
----------|:------:|---------------------------------------------------------
owner	  | string |owner of the todos                                                         
start_date| Bigint |start of the range date (This compares to todos expire date) only accept int of unix timestamp.
end_date  | Bigint |end of the range date (This compares to todos expire data) only accept int of unix timestamp.     
sort	  | string |Can be `expired date`, `completed`<p> default: `priority`



`GET`
> /api/v1/todos/{uid}

Parameters

Name	  |Type    |Description                                              
----------|:------:|---------------------------------------------------------
uid 	  | string |uid of todo



`POST`
> /api/v1/todo/add

Parameters

Name		   |  Type  |Description                                     
---------------|:------:|------------------------------------------------
title          | string |**Required**. The name of todo.                 
category       | string |**Required**. The category of todo.             
createdDate    | bigint |**Required**. The created date of todo.         
expiredDate    | bigint |The expired date of todo.                       
priority       |  int   |**Required**. The priority of the todo. 1 - High, 2 - Medium, 3 - Low, 4 - None.
remind         | string |Remind time of todos
completed      |  bool  |Todo complete or not
recentlyDeleted|  bool  |Recently deleted todo or not. if recently deleted it will not show in main page.
desc           | string |Detail description of todo.

