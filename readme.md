# fake-cnu-server

충남대학교 사이버캠퍼스 이용 시 발생하는 불필요한 일부 요청을 로컬 서버로 우회하여 사용자 경험을 향상시킨다.

<video src="assets/demo.mp4" width="500px"></video>


## 사용법

준비물: Chrome 계열 브라우저

1. 실행 파일 다운로드: [Releases](https://github.com/wirekang/fake-cnu-server/releases)에서 다운받거나 직접 빌드

2. `config.json` 다운로드

3. `config.json`과 실행 파일을 한 폴더에 두고 실행파일 실행 (CLI에서만)

4. [ResourceOverride](https://chromewebstore.google.com/detail/resource-override/pkoacgokdfckfpndoffpifphamojphii) 설치 및 설정

| 설명                               | 단축 시간(대략) | From                                                          | To                                                    |
| ---------------------------------- | --------------- | ------------------------------------------------------------- | ----------------------------------------------------- |
| 이름, 학번 등 기본 정보, 쪽지 갯수 | 2초             | https://dcs-learning.cnu.ac.kr/api/v1/user/getUserInfo        | http://localhost:12312/api/v1/user/getUserInfo        |
| To-Do-List 갯수                    | 0.4초           | https://dcs-learning.cnu.ac.kr/api/v1/week/getStdTodoCnt      | http://localhost:12312/api/v1/week/getStdTodoCnt      |
| 공지사항 목록*                     | 0.4초           | https://dcs-learning.cnu.ac.kr/api/v1/board/getBoardItem/list | http://localhost:12312/api/v1/board/getBoardItem/list |


> 설정 사진

![alt text](assets/image1.png)

> 공지사항 목록은 아래를 뜻함

![alt text](assets/image2.png)


## 다이어그램

```mermaid

sequenceDiagram
    Browser->>CNU Server: 요청 A
    CNU Server->>Browser: 응답 A

    Browser->>CNU Server: 요청 B
    CNU Server->>CNU Server: (오래 걸림)
    CNU Server->>Browser: 응답 B

```

```mermaid

sequenceDiagram
    Browser->>CNU Server: 요청 A
    CNU Server->>Browser: 응답 A

    Browser->>Local Server: 요청 B
    Local Server->>Browser: 응답 B

```

## 주의 사항

- 본 프로그램 사용으로 발생하는 모든 문제는 사용자 본인 책임임
- 사물이 거울에 보이는 것보다 가까이 있음
