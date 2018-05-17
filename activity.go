package cassandrainsertrecord

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/gocql/gocql"
)

// THIS IS ADDED
// log is the default package logger which we'll use to log
var log = logger.GetLogger("activity-CassandraInsertRecord")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {
	// Get the activity data from the context
	clusterIP := context.GetInput("ClusterIP").(string)
	keySpace := context.GetInput("Keyspace").(string)
	tableName := context.GetInput("TableName").(string)
	insertQuery := context.GetInput("InsertQuery").(string)

	// Use the log object to log the greeting
	log.Debugf("The Flogo engine says connect to [%s] with  [%s] with table [%s]", clusterIP, keySpace, tableName)
	log.Debugf("Flogo is about to insert [%s]  on cluster [%s]", insertQuery, clusterIP)
	
	fmt.Println("The Flogo engine says connect to "+clusterIP+" with  "+keySpace+" with table "+tableName)
	fmt.Println("Flogo is about to insert "+insertQuery+"  on cluster "+clusterIP)

	// Provide the cassandra cluster instance here.
	cluster := gocql.NewCluster(clusterIP)

	// gocql requires the keyspace to be provided before the session is created.
	// In future there might be provisions to do this later.
	cluster.Keyspace = keySpace
	
	session, err := cluster.CreateSession()
	log.Debugf("Session Created Sucessfully")
	fmt.Println("Session Created Sucessfully")

	if err != nil {
		log.Debugf("Could not connect to cassandra cluster : " , err)
		fmt.Println("Could not connect to cassandra cluster : " , err)
	}
	log.Debugf("Session : ", session)
	log.Debugf("Cluster : ", clusterIP)
	log.Debugf("Keyspace : ", keySpace)
	log.Debugf("Session Timeout : ", cluster.Timeout)
	log.Debugf("TableName : ", tableName)
	log.Debugf("Insert Query : ", insertQuery)

	log.Debugf("Next Step is Insert Query Execution")

	//fmt.Println("Session : ", session)
	//fmt.Println("Cluster : ", clusterIP)
	//fmt.Println("Keyspace : ", keySpace)
	//fmt.Println("Session Timeout : ", cluster.Timeout)
	//fmt.Println("TableName : ", tableName)
	//fmt.Println("Insert Query : ", insertQuery)

	fmt.Println("Next Step is Insert Query Execution")
	
	if err := session.Query(insertQuery).Exec(); err != nil {
		log.Debugf("Error In insert Query : " , err)
		fmt.Println("Error In insert Query : " , err)
	}
	

	// Set the result as part of the context
	context.SetOutput("result", "Record Inserted SuccessFully")
	
	fmt.Println("Record Inserted SuccessFully")
	// Signal to the Flogo engine that the activity is completed
	return true, nil
}
