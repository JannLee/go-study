# Go study Day-7
- 스터디 일시: 2023년 5월 24일 수요일 오후 5시 ~ 오후 7시
- 참석자(5명) 
    - 김덕현, 신윤호, 심우상, 유재욱, 정회형

## 스터디 범위
- 교재: Tucker의 Go 언어 프로그래밍
- 범위: 22장 자료구조 ~ 23장 에러 핸들링

## 스터디 진행
- 당일 랜덤 뽑기(원판 돌리기)
- 발표자 1명 : 김덕현
- 서기 1명 : 다같이
- 공지 1명 : 유재욱

## 22장 자료구조

v,ok : = m[3] 
맵에서 3번째 요소가 존재하는지는  v(value),ok(존재여부)를 받아서 존재여부의 true/false
여부로 확인할수 있다. 

이게 다른 언어와 차이가 있다 . javascript = undefined , python = null 을 리턴?



## 23장 에러 핸들링

Error 

에러를 리턴하더라도 defer file.close() 통해서 반드시 처리 해야할 부분 처리 
에러는 fmt.Errof() , errors package New로 생성 가능

Errorf를 통해서 발생한 에러에 커스텀한 메시지를 발생 시킬수 있고(래핑)
errors.As() 를 통해서 원본 에러메시지가 NumError타입인지(특정타입) 검사할수 있다. 
좀 특이함.

panic()

java System.exit(0)와 비슷하고 e.printStackTrace()처럼 콜 스택을 뿌려준다.
패닉 전파 복구 
 defer구문에 recover()를 사용함으로 panic 전파를 중단하고 프로그램을 종료시키지 않을수 있다.
 

recover()

recover는 panic 객체를 리턴해준다.

쿠버네티스 패닉 예제

func (proxy *ProxyHttpServer) handleHttps(w http.ResponseWriter, r *http.Request) {
	
    ctx := &ProxyCtx{Req: r, Session: atomic.AddInt64(&proxy.sess, 1), proxy: proxy}

	hij, ok := w.(http.Hijacker)
	if !ok {
		panic("httpserver does not support hijacking")
	}



## 다음 스터디
- 2023년 6월 7일 수요일

## 회고
- 김덕현: golang에서 에러를 바라보는 관점이 신기했습니다.
- 신윤호: 재미있었습니다. !!! 
- 심우상: golang만의 특이한 점을 공부할 수 있어서 좋았습니다~
- 유재욱:
- 정회형: try-catch가 아니라 panic 후 recover로 살린다는 개념이 헷갈리긴 합니다.
