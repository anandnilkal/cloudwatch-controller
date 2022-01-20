## Welcome to CloudWatch-Controller

CloudWatch controller supports following features
- CloudWatch Alarm Creation
- CloudWatch Alarm Updation
- CloudWatch Alarm Deletion
- Alarm Resource Tagging

# CRD examples

## Cloudwatch Alarm with Single Metric

```markdown
 apiVersion: cloudwatch.anandnilkal.io/v1alpha1
 kind: Alarms
 metadata:
   name: alarms-sample
 spec:
   name: alarms-sample
   operator: "GreaterThanThreshold"
   evaluationPeriod: 5
   alarmActions:
     - arn:aws:sns:us-east-1:123456789:alerts
   description: "test alarm creation"
   numDataPointsToAlarm: 5
   dimensions:
   - name: QueueName
     value: cloud-systems-Alert-Listener-Queue
   metricName: ApproximateNumberOfMessagesNotVisible
   statistics: Average
   namespace: AWS/SQS
   period: 60
   threshold: 10
   tags:
   - key: release
     value: test
```
<!-- Syntax highlighted code block

# Header 1
## Header 2
### Header 3

- Bulleted
- List

1. Numbered
2. List

**Bold** and _Italic_ and `Code` text

[Link](url) and ![Image](src)
 -->
## Cloudwatch Alarm with Math expression

```markdown
apiVersion: cloudwatch.anandnilkal.io/v1alpha1
kind: Alarms
metadata:
  name: samle-alarm-1
spec:
  # TODO(user): Add fields here
  name: samle-alarm-1
  operator: "GreaterThanThreshold"
  evaluationPeriod: 5
  numDataPointsToAlarm: 5
  alarmActions:
    - arn:aws:sns:us-east-1:123456789:alerts
  description: "test alarm creation"
  metrics:
  - id: e1
    expression: m1+m2
    label: e1
    returnData: true
  - id: m1
    returnData: false
    metricStat:
      period: 60
      stat: Average
      metric:
        dimensions:
        - name: StreamName
          value: cl-e2e-bssids-wips-stream
        metricName: BufferingTime
        namespace: KinesisProducerLibrary
  - id: m2
    returnData: false
    metricStat:
      period: 60
      stat: Average
      metric:
        dimensions:
        - name: StreamName
          value: cl-e2e-bssids-wips-stream
        metricName: BufferingTime
        namespace: KinesisProducerLibrary
  threshold: 10
```

# Installation

## kustomization

```console
cd config/default
kubectl apply -k ./
```
<!-- For more details see [Basic writing and formatting syntax](https://docs.github.com/en/github/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax).

### Jekyll Themes

Your Pages site will use the layout and styles from the Jekyll theme you have selected in your [repository settings](https://github.com/anandnilkal/cloudwatch-controller/settings/pages). The name of this theme is saved in the Jekyll `_config.yml` configuration file.

### Support or Contact

Having trouble with Pages? Check out our [documentation](https://docs.github.com/categories/github-pages-basics/) or [contact support](https://support.github.com/contact) and we’ll help you sort it out.
 -->
