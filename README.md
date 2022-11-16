
# A hybrid distributed system implementation in GO



## Feature Description & details : 
> Gradebook
- Teacher portal showing a listing of an in-memory collection of studetns and their average mean grade score

- Student detials page showing the grades of possible domain types in the list (Homework, Quiz, Test)
- Ability to add a new grade for any student in any supoorted domain type 

## Implememantion Details


![Services](./imgs/components.png?raw=true "Services")
![Services Architecture](./imgs/architecture.png?raw=true "Services Architecture")



## Running commands

**Log Service**	
```bash
go run cmd/logservice/main.go 
```
**Registry Service**	
```bash
go run cmd/registryservice/main.go
```

**Grading Service**	
```bash
 go run cmd/gradingservice/main.go
```

**Teacher Portal Service**	
```bash
 go run cmd/teacherportal/main.go 
```


