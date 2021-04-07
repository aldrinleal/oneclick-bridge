# oneclick-bridge (ocb)

## What is it?

A small serverless application written in go that given a topic (created in the template), publishes a message into it from AWS IoT 1-Click.

## Sample

See the video at https://youtu.be/4LcWzgyEYCQ 

For the record, here's a sample 1-click event:

```json
{
  "deviceEvent": {
    "buttonClicked": {
      "additionalInfo": {
        "version": "1.8.0"
      },
      "clickType": "SINGLE",
      "reportedTime": "2021-04-07T03:49:05.968Z"
    }
  },
  "deviceInfo": {
    "attributes": {
      "deviceTemplateName": "clickRequest",
      "placementName": "sample-1click-placement",
      "projectName": "topicana",
      "projectRegion": "us-west-2"
    },
    "deviceId": "P5SJVQ2007JDEVICE",
    "remainingLife": 100,
    "type": "button"
  },
  "devicePayload": {
    "certificateId": "46d9631462ae5e47b847e23b94251188ad76cd6c318b6eb210e963f106a45088",
    "clickType": "SINGLE",
    "remainingLife": 100,
    "reportedTime": 1617767345968,
    "serialNumber": "P5SJVQ2007JDEVICE",
    "topic": "/Devices/Button/P5SJVQ",
    "version": "1.8.0"
  },
  "placementInfo": {
    "attributes": {
      "location": ""
    },
    "devices": {
      "oneClickRequest": ""
    },
    "placementName": "sample-1click-placement",
    "projectName": "topicana"
  }
}
```

## Deploying

(requires go and serverless framework):

```
$ make && sls deploy
```

## Tools:

* golang
* [yarn](https://yarnpkg.com/)
* [serverless](https://www.serverless.com/)
* [utern](https://github.com/knqyf263/utern)
* [jsonutil](https://github.com/bashtian/jsonutils) (used to convert json into struct)
* [qdsns](https://github.com/aldrinleal/qdsns)
* [ngrok](https://ngrok.io)

## Details

TBD
