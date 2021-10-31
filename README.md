# Mechanical Receptionist
来客を超音波センサーで感知して、LINE に通知してくれるアプリケーション。  
ヘッドホンなどをしていると来客に気づけないことがあり、それを解消するために作成した。

## Usage
```
$ make firebase-emulator/start
$ make watch
```

## Requirements
- Hardware
  - Arduino UNO
  - Arduino Ethernet Shield 2
  - HY-SRF05
- Software
  - [cosmtrek/air](https://github.com/cosmtrek/air) - ホットリロードに使用している

## Architecture
### Hardware
<img src="https://user-images.githubusercontent.com/40758815/139569928-796fa62d-d6cb-4047-bda7-1de9753dac31.png" width=600 />

### Software
