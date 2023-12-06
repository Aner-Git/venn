# venn

## How to run
1) you must have docker on your system  - get it
2) clone the code and:
  * cd venn/frontend/webapp/ 
  * yarn install (or npm)
  * cd venn/local    (the dir has several utiltiy script for development)
  * ./compose-up.sh   (to bring the containers: you should have 4 container running: venn, webapp, mysql, adminer )
  **Note! mysql loads data on the first run and the backend container might not connect. Wait for log message to stop, Ctrl-c, and Rerun ./compose-up.sh**
  * navigate to http://localhost:9000/ 
  * ./compose-down.sh brings the system down




## API
The endpoint is /neighborhoods - full API description in the code

## Data Modeling / Code
The data model is simple sql: neighborhoods and  public_transport_availablities   
SQL was chosen as it easier to query and normalize   

Indices are added to max distance and age to support filtering   
While the data was noramlized for faster queries public transport data could have stayed in the neighborhoods table    

Frontend is written in React with large amount of modelarity. Backend in GO with simple logical structure and support w/ Gin Framework

## Error Handling
The server returns error when input query params are invalid or internal DB access issues   

## Testig
Unit tests are implemented on some of the backend files. Integration tests and DB tests need tobe added (time). 

## Performance
To not flood the server, paging was implemented.   
Cache is automatically performed on the frontend by react-query, so for example pagin 'back' without changing the query is fast.    
Table de normalization can be done to save joins

## Scalabilty/
The application is containerized with no state. I.e. scaling the number of containers is easy by running in a cluster w/  a LB    
Data is mostly read and so adding DB slaves for reading should expediet reads if needed. 



