// Package job은 오래 걸리는 작업의 id, 진행 단계, 멱등 키, 취소를 다룬다.
//
// 작업 상태는 휘발성이다. 진실은 Kotlin의 GSMSV DB에 있고, Go가 재시작해
// 작업을 잃으면 Kotlin이 실제 상태를 다시 확인해 복구한다.
package job
