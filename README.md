---
title: CassandraDB Insert Activity
---

# CassandraDB Insert Activity
This activity allows you to Insert a record to a particular table from the CassandraDB server

## Installation
### Flogo CLI
```bash
flogo install github.com/dhire05/cassandrainsertrecord
```

## Schema
Inputs and Outputs:

```json
{   
  "inputs":[
    {
      "name": "ClusterIP",
      "type": "string",
	  "required": true      
    },
	{
      "name": "Keyspace",
      "type": "string",
      "required": true
    },
	{
      "name": "TableName",
      "type": "string",
      "required": true
    },
	 {
      "name": "InsertQuery",
      "type": "string",
      "required": true
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "any"
    }
  ]
 }
```
## Settings
| Setting        | Required | Description |
|:---------------|:---------|:------------|
| ClusterIP      | True     | The CassandraDB cluster instance |         
| Keyspace       | True     | The name of the Keyspace
| TableName      | True     | The name of table to delete record
| InsertQuery    | True		| The insert query is to insert a record into the table|

## Example
The below example is to insert a record into CassandraDB

```json
{
  "id": "CassandraDB_1",
  "name": "CassandraDB connector",
  "description": "Insert record into CassandraDB",
  "activity": {
    "ref": "github.com/dhire05/cassandrainsertrecord",
    "input": {
      "ClusterIP": "127.0.0.1",
      "Keyspace": "sample",
      "TableName": "employee",
	  "InsertQuery": "INSERT INTO employee (empid , name, salary) VALUES (102, 'xyz', 5005)"          
    }
  }
}
```