`GET`
> /api/v1/todos

Parameters

Name		|Type    |Description                                              
----------|:------:|---------------------------------------------------------
owner		| string |owner of the todos                                                         
start_date| Bigint |This will return todos expire date greater than the value
end_date  | Bigint |todos end date                                           
sort		| string |Can be `expired date`, `completed`<p> default: `priority`



`GET`
> /api/v1/todos/{uid}

Parameters

Name		|Type    |Description                                              
----------|:------:|---------------------------------------------------------
uid 		|  uuid  |uuid of todo



`POST`
> /api/v1/todo/add

Parameters

Name		     |  Type  |Description                                     
---------------|:------:|------------------------------------------------
title          | string |**Required**. The name of todo.                 
category       | string |**Required**. The category of todo.             
createdDate    | bigint |**Required**. The created date of todo.         
expiredDate    | bigint |The expired date of todo.                       
priority       |  int   |**Required**. The priority of the todo. 1 - High, 2 - Medium, 3 - Low, 4 - None.
remind         | string |Remind time of todos
completed      |  bool  |Todo complete or not
recentlyDeleted|  bool  |Recently deleted todo or not. if recently deleted it will not show in main page.
desc           | string |


var createdDate: Double
    var expiredDate: Double?
    var priority: Int
    var remind: String?
    var desc: String?
    var subtasks: [TDSubtask]?
    var completed: Bool
    var recentlyDelete: Bool