/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"io"
	"math"
)

type Dm_build_361 struct {
	dm_build_362 []byte
	dm_build_363 int
}

func Dm_build_364(dm_build_365 int) *Dm_build_361 {
	return &Dm_build_361{make([]byte, 0, dm_build_365), 0}
}

func Dm_build_366(dm_build_367 []byte) *Dm_build_361 {
	return &Dm_build_361{dm_build_367, 0}
}

func (dm_build_369 *Dm_build_361) dm_build_368(dm_build_370 int) *Dm_build_361 {

	dm_build_371 := len(dm_build_369.dm_build_362)
	dm_build_372 := cap(dm_build_369.dm_build_362)

	if dm_build_371+dm_build_370 <= dm_build_372 {
		dm_build_369.dm_build_362 = dm_build_369.dm_build_362[:dm_build_371+dm_build_370]
	} else {

		var calCap = int64(math.Max(float64(2*dm_build_372), float64(dm_build_370+dm_build_371)))

		nbuf := make([]byte, dm_build_370+dm_build_371, calCap)
		copy(nbuf, dm_build_369.dm_build_362)
		dm_build_369.dm_build_362 = nbuf
	}

	return dm_build_369
}

func (dm_build_374 *Dm_build_361) Dm_build_373() int {
	return len(dm_build_374.dm_build_362)
}

func (dm_build_376 *Dm_build_361) Dm_build_375(dm_build_377 int) *Dm_build_361 {
	for i := dm_build_377; i < len(dm_build_376.dm_build_362); i++ {
		dm_build_376.dm_build_362[i] = 0
	}
	dm_build_376.dm_build_362 = dm_build_376.dm_build_362[:dm_build_377]
	return dm_build_376
}

func (dm_build_379 *Dm_build_361) Dm_build_378(dm_build_380 int) *Dm_build_361 {
	dm_build_379.dm_build_363 = dm_build_380
	return dm_build_379
}

func (dm_build_382 *Dm_build_361) Dm_build_381() int {
	return dm_build_382.dm_build_363
}

func (dm_build_384 *Dm_build_361) Dm_build_383(dm_build_385 bool) int {
	return len(dm_build_384.dm_build_362) - dm_build_384.dm_build_363
}

func (dm_build_387 *Dm_build_361) Dm_build_386(dm_build_388 int, dm_build_389 bool, dm_build_390 bool) *Dm_build_361 {

	if dm_build_389 {
		if dm_build_390 {
			dm_build_387.dm_build_368(dm_build_388)
		} else {
			dm_build_387.dm_build_362 = dm_build_387.dm_build_362[:len(dm_build_387.dm_build_362)-dm_build_388]
		}
	} else {
		if dm_build_390 {
			dm_build_387.dm_build_363 += dm_build_388
		} else {
			dm_build_387.dm_build_363 -= dm_build_388
		}
	}

	return dm_build_387
}

func (dm_build_392 *Dm_build_361) Dm_build_391(dm_build_393 io.Reader, dm_build_394 int) (int, error) {
	dm_build_395 := len(dm_build_392.dm_build_362)
	dm_build_392.dm_build_368(dm_build_394)
	dm_build_396 := 0
	for dm_build_394 > 0 {
		n, err := dm_build_393.Read(dm_build_392.dm_build_362[dm_build_395+dm_build_396:])
		if n > 0 && err == io.EOF {
			dm_build_396 += n
			dm_build_392.dm_build_362 = dm_build_392.dm_build_362[:dm_build_395+dm_build_396]
			return dm_build_396, nil
		} else if n > 0 && err == nil {
			dm_build_394 -= n
			dm_build_396 += n
		} else if n == 0 && err != nil {
			return -1, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
		}
	}

	return dm_build_396, nil
}

func (dm_build_398 *Dm_build_361) Dm_build_397(dm_build_399 io.Writer) (*Dm_build_361, error) {
	if _, err := dm_build_399.Write(dm_build_398.dm_build_362); err != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
	}
	return dm_build_398, nil
}

func (dm_build_401 *Dm_build_361) Dm_build_400(dm_build_402 bool) int {
	dm_build_403 := len(dm_build_401.dm_build_362)
	dm_build_401.dm_build_368(1)

	if dm_build_402 {
		return copy(dm_build_401.dm_build_362[dm_build_403:], []byte{1})
	} else {
		return copy(dm_build_401.dm_build_362[dm_build_403:], []byte{0})
	}
}

func (dm_build_405 *Dm_build_361) Dm_build_404(dm_build_406 byte) int {
	dm_build_407 := len(dm_build_405.dm_build_362)
	dm_build_405.dm_build_368(1)

	return copy(dm_build_405.dm_build_362[dm_build_407:], Dm_build_1.Dm_build_179(dm_build_406))
}

func (dm_build_409 *Dm_build_361) Dm_build_408(dm_build_410 int8) int {
	dm_build_411 := len(dm_build_409.dm_build_362)
	dm_build_409.dm_build_368(1)

	return copy(dm_build_409.dm_build_362[dm_build_411:], Dm_build_1.Dm_build_182(dm_build_410))
}

func (dm_build_413 *Dm_build_361) Dm_build_412(dm_build_414 int16) int {
	dm_build_415 := len(dm_build_413.dm_build_362)
	dm_build_413.dm_build_368(2)

	return copy(dm_build_413.dm_build_362[dm_build_415:], Dm_build_1.Dm_build_185(dm_build_414))
}

func (dm_build_417 *Dm_build_361) Dm_build_416(dm_build_418 int32) int {
	dm_build_419 := len(dm_build_417.dm_build_362)
	dm_build_417.dm_build_368(4)

	return copy(dm_build_417.dm_build_362[dm_build_419:], Dm_build_1.Dm_build_188(dm_build_418))
}

func (dm_build_421 *Dm_build_361) Dm_build_420(dm_build_422 uint8) int {
	dm_build_423 := len(dm_build_421.dm_build_362)
	dm_build_421.dm_build_368(1)

	return copy(dm_build_421.dm_build_362[dm_build_423:], Dm_build_1.Dm_build_200(dm_build_422))
}

func (dm_build_425 *Dm_build_361) Dm_build_424(dm_build_426 uint16) int {
	dm_build_427 := len(dm_build_425.dm_build_362)
	dm_build_425.dm_build_368(2)

	return copy(dm_build_425.dm_build_362[dm_build_427:], Dm_build_1.Dm_build_203(dm_build_426))
}

func (dm_build_429 *Dm_build_361) Dm_build_428(dm_build_430 uint32) int {
	dm_build_431 := len(dm_build_429.dm_build_362)
	dm_build_429.dm_build_368(4)

	return copy(dm_build_429.dm_build_362[dm_build_431:], Dm_build_1.Dm_build_206(dm_build_430))
}

func (dm_build_433 *Dm_build_361) Dm_build_432(dm_build_434 uint64) int {
	dm_build_435 := len(dm_build_433.dm_build_362)
	dm_build_433.dm_build_368(8)

	return copy(dm_build_433.dm_build_362[dm_build_435:], Dm_build_1.Dm_build_209(dm_build_434))
}

func (dm_build_437 *Dm_build_361) Dm_build_436(dm_build_438 float32) int {
	dm_build_439 := len(dm_build_437.dm_build_362)
	dm_build_437.dm_build_368(4)

	return copy(dm_build_437.dm_build_362[dm_build_439:], Dm_build_1.Dm_build_206(math.Float32bits(dm_build_438)))
}

func (dm_build_441 *Dm_build_361) Dm_build_440(dm_build_442 float64) int {
	dm_build_443 := len(dm_build_441.dm_build_362)
	dm_build_441.dm_build_368(8)

	return copy(dm_build_441.dm_build_362[dm_build_443:], Dm_build_1.Dm_build_209(math.Float64bits(dm_build_442)))
}

func (dm_build_445 *Dm_build_361) Dm_build_444(dm_build_446 []byte) int {
	dm_build_447 := len(dm_build_445.dm_build_362)
	dm_build_445.dm_build_368(len(dm_build_446))
	return copy(dm_build_445.dm_build_362[dm_build_447:], dm_build_446)
}

func (dm_build_449 *Dm_build_361) Dm_build_448(dm_build_450 []byte) int {
	return dm_build_449.Dm_build_416(int32(len(dm_build_450))) + dm_build_449.Dm_build_444(dm_build_450)
}

func (dm_build_452 *Dm_build_361) Dm_build_451(dm_build_453 []byte) int {
	return dm_build_452.Dm_build_420(uint8(len(dm_build_453))) + dm_build_452.Dm_build_444(dm_build_453)
}

func (dm_build_455 *Dm_build_361) Dm_build_454(dm_build_456 []byte) int {
	return dm_build_455.Dm_build_424(uint16(len(dm_build_456))) + dm_build_455.Dm_build_444(dm_build_456)
}

func (dm_build_458 *Dm_build_361) Dm_build_457(dm_build_459 []byte) int {
	return dm_build_458.Dm_build_444(dm_build_459) + dm_build_458.Dm_build_404(0)
}

func (dm_build_461 *Dm_build_361) Dm_build_460(dm_build_462 string, dm_build_463 string, dm_build_464 *DmConnection) int {
	dm_build_465 := Dm_build_1.Dm_build_217(dm_build_462, dm_build_463, dm_build_464)
	return dm_build_461.Dm_build_448(dm_build_465)
}

func (dm_build_467 *Dm_build_361) Dm_build_466(dm_build_468 string, dm_build_469 string, dm_build_470 *DmConnection) int {
	dm_build_471 := Dm_build_1.Dm_build_217(dm_build_468, dm_build_469, dm_build_470)
	return dm_build_467.Dm_build_451(dm_build_471)
}

func (dm_build_473 *Dm_build_361) Dm_build_472(dm_build_474 string, dm_build_475 string, dm_build_476 *DmConnection) int {
	dm_build_477 := Dm_build_1.Dm_build_217(dm_build_474, dm_build_475, dm_build_476)
	return dm_build_473.Dm_build_454(dm_build_477)
}

func (dm_build_479 *Dm_build_361) Dm_build_478(dm_build_480 string, dm_build_481 string, dm_build_482 *DmConnection) int {
	dm_build_483 := Dm_build_1.Dm_build_217(dm_build_480, dm_build_481, dm_build_482)
	return dm_build_479.Dm_build_457(dm_build_483)
}

func (dm_build_485 *Dm_build_361) Dm_build_484() byte {
	dm_build_486 := Dm_build_1.Dm_build_94(dm_build_485.dm_build_362, dm_build_485.dm_build_363)
	dm_build_485.dm_build_363++
	return dm_build_486
}

func (dm_build_488 *Dm_build_361) Dm_build_487() int16 {
	dm_build_489 := Dm_build_1.Dm_build_98(dm_build_488.dm_build_362, dm_build_488.dm_build_363)
	dm_build_488.dm_build_363 += 2
	return dm_build_489
}

func (dm_build_491 *Dm_build_361) Dm_build_490() int32 {
	dm_build_492 := Dm_build_1.Dm_build_103(dm_build_491.dm_build_362, dm_build_491.dm_build_363)
	dm_build_491.dm_build_363 += 4
	return dm_build_492
}

func (dm_build_494 *Dm_build_361) Dm_build_493() int64 {
	dm_build_495 := Dm_build_1.Dm_build_108(dm_build_494.dm_build_362, dm_build_494.dm_build_363)
	dm_build_494.dm_build_363 += 8
	return dm_build_495
}

func (dm_build_497 *Dm_build_361) Dm_build_496() float32 {
	dm_build_498 := Dm_build_1.Dm_build_113(dm_build_497.dm_build_362, dm_build_497.dm_build_363)
	dm_build_497.dm_build_363 += 4
	return dm_build_498
}

func (dm_build_500 *Dm_build_361) Dm_build_499() float64 {
	dm_build_501 := Dm_build_1.Dm_build_117(dm_build_500.dm_build_362, dm_build_500.dm_build_363)
	dm_build_500.dm_build_363 += 8
	return dm_build_501
}

func (dm_build_503 *Dm_build_361) Dm_build_502() uint8 {
	dm_build_504 := Dm_build_1.Dm_build_121(dm_build_503.dm_build_362, dm_build_503.dm_build_363)
	dm_build_503.dm_build_363 += 1
	return dm_build_504
}

func (dm_build_506 *Dm_build_361) Dm_build_505() uint16 {
	dm_build_507 := Dm_build_1.Dm_build_125(dm_build_506.dm_build_362, dm_build_506.dm_build_363)
	dm_build_506.dm_build_363 += 2
	return dm_build_507
}

func (dm_build_509 *Dm_build_361) Dm_build_508() uint32 {
	dm_build_510 := Dm_build_1.Dm_build_130(dm_build_509.dm_build_362, dm_build_509.dm_build_363)
	dm_build_509.dm_build_363 += 4
	return dm_build_510
}

func (dm_build_512 *Dm_build_361) Dm_build_511(dm_build_513 int) []byte {
	dm_build_514 := Dm_build_1.Dm_build_152(dm_build_512.dm_build_362, dm_build_512.dm_build_363, dm_build_513)
	dm_build_512.dm_build_363 += dm_build_513
	return dm_build_514
}

func (dm_build_516 *Dm_build_361) Dm_build_515() []byte {
	return dm_build_516.Dm_build_511(int(dm_build_516.Dm_build_490()))
}

func (dm_build_518 *Dm_build_361) Dm_build_517() []byte {
	return dm_build_518.Dm_build_511(int(dm_build_518.Dm_build_484()))
}

func (dm_build_520 *Dm_build_361) Dm_build_519() []byte {
	return dm_build_520.Dm_build_511(int(dm_build_520.Dm_build_487()))
}

func (dm_build_522 *Dm_build_361) Dm_build_521(dm_build_523 int) []byte {
	return dm_build_522.Dm_build_511(dm_build_523)
}

func (dm_build_525 *Dm_build_361) Dm_build_524() []byte {
	dm_build_526 := 0
	for dm_build_525.Dm_build_484() != 0 {
		dm_build_526++
	}
	dm_build_525.Dm_build_386(dm_build_526, false, false)
	return dm_build_525.Dm_build_511(dm_build_526)
}

func (dm_build_528 *Dm_build_361) Dm_build_527(dm_build_529 int, dm_build_530 string, dm_build_531 *DmConnection) string {
	return Dm_build_1.Dm_build_254(dm_build_528.Dm_build_511(dm_build_529), dm_build_530, dm_build_531)
}

func (dm_build_533 *Dm_build_361) Dm_build_532(dm_build_534 string, dm_build_535 *DmConnection) string {
	return Dm_build_1.Dm_build_254(dm_build_533.Dm_build_515(), dm_build_534, dm_build_535)
}

func (dm_build_537 *Dm_build_361) Dm_build_536(dm_build_538 string, dm_build_539 *DmConnection) string {
	return Dm_build_1.Dm_build_254(dm_build_537.Dm_build_517(), dm_build_538, dm_build_539)
}

func (dm_build_541 *Dm_build_361) Dm_build_540(dm_build_542 string, dm_build_543 *DmConnection) string {
	return Dm_build_1.Dm_build_254(dm_build_541.Dm_build_519(), dm_build_542, dm_build_543)
}

func (dm_build_545 *Dm_build_361) Dm_build_544(dm_build_546 string, dm_build_547 *DmConnection) string {
	return Dm_build_1.Dm_build_254(dm_build_545.Dm_build_524(), dm_build_546, dm_build_547)
}

func (dm_build_549 *Dm_build_361) Dm_build_548(dm_build_550 int, dm_build_551 byte) int {
	return dm_build_549.Dm_build_584(dm_build_550, Dm_build_1.Dm_build_179(dm_build_551))
}

func (dm_build_553 *Dm_build_361) Dm_build_552(dm_build_554 int, dm_build_555 int16) int {
	return dm_build_553.Dm_build_584(dm_build_554, Dm_build_1.Dm_build_185(dm_build_555))
}

func (dm_build_557 *Dm_build_361) Dm_build_556(dm_build_558 int, dm_build_559 int32) int {
	return dm_build_557.Dm_build_584(dm_build_558, Dm_build_1.Dm_build_188(dm_build_559))
}

func (dm_build_561 *Dm_build_361) Dm_build_560(dm_build_562 int, dm_build_563 int64) int {
	return dm_build_561.Dm_build_584(dm_build_562, Dm_build_1.Dm_build_191(dm_build_563))
}

func (dm_build_565 *Dm_build_361) Dm_build_564(dm_build_566 int, dm_build_567 float32) int {
	return dm_build_565.Dm_build_584(dm_build_566, Dm_build_1.Dm_build_194(dm_build_567))
}

func (dm_build_569 *Dm_build_361) Dm_build_568(dm_build_570 int, dm_build_571 float64) int {
	return dm_build_569.Dm_build_584(dm_build_570, Dm_build_1.Dm_build_197(dm_build_571))
}

func (dm_build_573 *Dm_build_361) Dm_build_572(dm_build_574 int, dm_build_575 uint8) int {
	return dm_build_573.Dm_build_584(dm_build_574, Dm_build_1.Dm_build_200(dm_build_575))
}

func (dm_build_577 *Dm_build_361) Dm_build_576(dm_build_578 int, dm_build_579 uint16) int {
	return dm_build_577.Dm_build_584(dm_build_578, Dm_build_1.Dm_build_203(dm_build_579))
}

func (dm_build_581 *Dm_build_361) Dm_build_580(dm_build_582 int, dm_build_583 uint32) int {
	return dm_build_581.Dm_build_584(dm_build_582, Dm_build_1.Dm_build_206(dm_build_583))
}

func (dm_build_585 *Dm_build_361) Dm_build_584(dm_build_586 int, dm_build_587 []byte) int {
	return copy(dm_build_585.dm_build_362[dm_build_586:], dm_build_587)
}

func (dm_build_589 *Dm_build_361) Dm_build_588(dm_build_590 int, dm_build_591 []byte) int {
	return dm_build_589.Dm_build_556(dm_build_590, int32(len(dm_build_591))) + dm_build_589.Dm_build_584(dm_build_590+4, dm_build_591)
}

func (dm_build_593 *Dm_build_361) Dm_build_592(dm_build_594 int, dm_build_595 []byte) int {
	return dm_build_593.Dm_build_548(dm_build_594, byte(len(dm_build_595))) + dm_build_593.Dm_build_584(dm_build_594+1, dm_build_595)
}

func (dm_build_597 *Dm_build_361) Dm_build_596(dm_build_598 int, dm_build_599 []byte) int {
	return dm_build_597.Dm_build_552(dm_build_598, int16(len(dm_build_599))) + dm_build_597.Dm_build_584(dm_build_598+2, dm_build_599)
}

func (dm_build_601 *Dm_build_361) Dm_build_600(dm_build_602 int, dm_build_603 []byte) int {
	return dm_build_601.Dm_build_584(dm_build_602, dm_build_603) + dm_build_601.Dm_build_548(dm_build_602+len(dm_build_603), 0)
}

func (dm_build_605 *Dm_build_361) Dm_build_604(dm_build_606 int, dm_build_607 string, dm_build_608 string, dm_build_609 *DmConnection) int {
	return dm_build_605.Dm_build_588(dm_build_606, Dm_build_1.Dm_build_217(dm_build_607, dm_build_608, dm_build_609))
}

func (dm_build_611 *Dm_build_361) Dm_build_610(dm_build_612 int, dm_build_613 string, dm_build_614 string, dm_build_615 *DmConnection) int {
	return dm_build_611.Dm_build_592(dm_build_612, Dm_build_1.Dm_build_217(dm_build_613, dm_build_614, dm_build_615))
}

func (dm_build_617 *Dm_build_361) Dm_build_616(dm_build_618 int, dm_build_619 string, dm_build_620 string, dm_build_621 *DmConnection) int {
	return dm_build_617.Dm_build_596(dm_build_618, Dm_build_1.Dm_build_217(dm_build_619, dm_build_620, dm_build_621))
}

func (dm_build_623 *Dm_build_361) Dm_build_622(dm_build_624 int, dm_build_625 string, dm_build_626 string, dm_build_627 *DmConnection) int {
	return dm_build_623.Dm_build_600(dm_build_624, Dm_build_1.Dm_build_217(dm_build_625, dm_build_626, dm_build_627))
}

func (dm_build_629 *Dm_build_361) Dm_build_628(dm_build_630 int) byte {
	return Dm_build_1.Dm_build_222(dm_build_629.Dm_build_655(dm_build_630, 1))
}

func (dm_build_632 *Dm_build_361) Dm_build_631(dm_build_633 int) int16 {
	return Dm_build_1.Dm_build_225(dm_build_632.Dm_build_655(dm_build_633, 2))
}

func (dm_build_635 *Dm_build_361) Dm_build_634(dm_build_636 int) int32 {
	return Dm_build_1.Dm_build_228(dm_build_635.Dm_build_655(dm_build_636, 4))
}

func (dm_build_638 *Dm_build_361) Dm_build_637(dm_build_639 int) int64 {
	return Dm_build_1.Dm_build_231(dm_build_638.Dm_build_655(dm_build_639, 8))
}

func (dm_build_641 *Dm_build_361) Dm_build_640(dm_build_642 int) float32 {
	return Dm_build_1.Dm_build_234(dm_build_641.Dm_build_655(dm_build_642, 4))
}

func (dm_build_644 *Dm_build_361) Dm_build_643(dm_build_645 int) float64 {
	return Dm_build_1.Dm_build_237(dm_build_644.Dm_build_655(dm_build_645, 8))
}

func (dm_build_647 *Dm_build_361) Dm_build_646(dm_build_648 int) uint8 {
	return Dm_build_1.Dm_build_240(dm_build_647.Dm_build_655(dm_build_648, 1))
}

func (dm_build_650 *Dm_build_361) Dm_build_649(dm_build_651 int) uint16 {
	return Dm_build_1.Dm_build_243(dm_build_650.Dm_build_655(dm_build_651, 2))
}

func (dm_build_653 *Dm_build_361) Dm_build_652(dm_build_654 int) uint32 {
	return Dm_build_1.Dm_build_246(dm_build_653.Dm_build_655(dm_build_654, 4))
}

func (dm_build_656 *Dm_build_361) Dm_build_655(dm_build_657 int, dm_build_658 int) []byte {
	return dm_build_656.dm_build_362[dm_build_657 : dm_build_657+dm_build_658]
}

func (dm_build_660 *Dm_build_361) Dm_build_659(dm_build_661 int) []byte {
	dm_build_662 := dm_build_660.Dm_build_634(dm_build_661)
	return dm_build_660.Dm_build_655(dm_build_661+4, int(dm_build_662))
}

func (dm_build_664 *Dm_build_361) Dm_build_663(dm_build_665 int) []byte {
	dm_build_666 := dm_build_664.Dm_build_628(dm_build_665)
	return dm_build_664.Dm_build_655(dm_build_665+1, int(dm_build_666))
}

func (dm_build_668 *Dm_build_361) Dm_build_667(dm_build_669 int) []byte {
	dm_build_670 := dm_build_668.Dm_build_631(dm_build_669)
	return dm_build_668.Dm_build_655(dm_build_669+2, int(dm_build_670))
}

func (dm_build_672 *Dm_build_361) Dm_build_671(dm_build_673 int) []byte {
	dm_build_674 := 0
	for dm_build_672.Dm_build_628(dm_build_673) != 0 {
		dm_build_673++
		dm_build_674++
	}

	return dm_build_672.Dm_build_655(dm_build_673-dm_build_674, int(dm_build_674))
}

func (dm_build_676 *Dm_build_361) Dm_build_675(dm_build_677 int, dm_build_678 string, dm_build_679 *DmConnection) string {
	return Dm_build_1.Dm_build_254(dm_build_676.Dm_build_659(dm_build_677), dm_build_678, dm_build_679)
}

func (dm_build_681 *Dm_build_361) Dm_build_680(dm_build_682 int, dm_build_683 string, dm_build_684 *DmConnection) string {
	return Dm_build_1.Dm_build_254(dm_build_681.Dm_build_663(dm_build_682), dm_build_683, dm_build_684)
}

func (dm_build_686 *Dm_build_361) Dm_build_685(dm_build_687 int, dm_build_688 string, dm_build_689 *DmConnection) string {
	return Dm_build_1.Dm_build_254(dm_build_686.Dm_build_667(dm_build_687), dm_build_688, dm_build_689)
}

func (dm_build_691 *Dm_build_361) Dm_build_690(dm_build_692 int, dm_build_693 string, dm_build_694 *DmConnection) string {
	return Dm_build_1.Dm_build_254(dm_build_691.Dm_build_671(dm_build_692), dm_build_693, dm_build_694)
}
