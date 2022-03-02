# go를 개발하기 전에 이 과정을 실행 후 개발을 진행한다.

1. main.go 디렉터리에서 [go install]로 의존성 패키지를 설치 받는다.
2. [go mod tidy] 명령어를 사용하여 모듈을 갱신한다.

# Vscode로 작업할 경우

main.go와 함께 있는 go.mod 파일이 작업시 최상위 폴더로 인식되어야 합니다.
예) wscmanager/src/를 '폴더열기' 후 작업