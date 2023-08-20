import boto3
import requests
import datetime
import json
from requests_aws4auth import AWS4Auth

region = 'us-east-1' 
service = 'es'
credentials = boto3.Session().get_credentials()
awsauth = AWS4Auth(credentials.access_key, credentials.secret_key, region, service, session_token=credentials.token)

host = 'https://search-ddb-to-es-vxbv3zpxppqxmlolen5ttlkgde.us-east-1.es.amazonaws.com' # the Amazon ES domain
index = 'lambda-index'
type = 'lambda-type'
url = host + '/' + index + '/' + type

headers = { "Content-Type": "application/json" }

def lambda_handler(event, context):
    count = 0
    try:
        for record in event['Records']:
            event = record['eventName']
            
            document = dict()
            document['eventTimeStamp'] = datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S.%f')
            document['eventType'] = event
            
            if event.lower() == 'insert':
                document['NewImage'] = record['dynamodb']['NewImage']
            if event.lower() == 'remove':
                document['OldImage'] = record['dynamodb']['OldImage']
            if event.lower() == 'modify':
                document['NewImage'] = record['dynamodb']['NewImage']
                document['OldImage'] = record['dynamodb']['OldImage']
            
            if len(document) > 2:
                print(json.dumps(document, default=str))
                r = requests.post(url, auth=awsauth, json=document, headers=headers)
                r.raise_for_status()
                count += 1
            
        print(f'{count} number of records processed.')
    except Exception as e:
        print(e)