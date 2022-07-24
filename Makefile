install_deps:
		go get ./...

console_reader:
		go run ./app/console-reader/main.go

run_preloaded_values:
		go run ./app/run-preloaded-values/main.go "0-10/2 2 2 2 2 /usr/bin/find" "0 0,12 1 */2 * /run/script" "15-20,*/2 0 1,15 * 1-5 /abc" "0 22 * * 1-5 /xyz" "23 0-20/2 * * * /lmn" "*/15 0 1,15 * 1-5 /pqr"