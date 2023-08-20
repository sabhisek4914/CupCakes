# CupCakes

download the Data set CSV from here
The data in CSV is to be dumped into dynamo db using golang. 
Write a new job in golang to update 100 random rows that you have inserted in the previous step, with update_time column storing the actual time of update (create one if not exists) in dynamo. The new job should run every 5 minutes each time updating 100 records. You may use cloudwatch for configuring the cron. 
A Change data capture (CDC) from dynamo DB should be setup using dynamodb streams, with a golang lambda handler. This lambda should send the inserted, updated, deleted data to Elastic search. The change operation that occured on a record -- insert, update, delete should be a new column in ES. 
Create a reactjs app, use ReCharts to plot the graph below. The Recharts should use REST api hosted by AWS lambda (in golang). The AWS lambda should run an SQL on the Elastic search generating the response json for the Recharts. Given that we are updating the DB every 