/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"io"
	"math"
)

type Dm_build_358 struct {
	dm_build_359 []byte
	dm_build_360 int
}

func Dm_build_361(dm_build_362 int) *Dm_build_358 {
	return &Dm_build_358{make([]byte, 0, dm_build_362), 0}
}

func Dm_build_363(dm_build_364 []byte) *Dm_build_358 {
	return &Dm_build_358{dm_build_364, 0}
}

func (dm_build_366 *Dm_build_358) dm_build_365(dm_build_367 int) *Dm_build_358 {

	dm_build_368 := len(dm_build_366.dm_build_359)
	dm_build_369 := cap(dm_build_366.dm_build_359)

	if dm_build_368+dm_build_367 <= dm_build_369 {
		dm_build_366.dm_build_359 = dm_build_366.dm_build_359[:dm_build_368+dm_build_367]
	} else {
		remain := dm_build_367 + dm_build_368 - dm_build_369
		nbuf := make([]byte, dm_build_367+dm_build_368, 2*dm_build_369+remain)
		copy(nbuf, dm_build_366.dm_build_359)
		dm_build_366.dm_build_359 = nbuf
	}

	return dm_build_366
}

func (dm_build_371 *Dm_build_358) Dm_build_370() int {
	return len(dm_build_371.dm_build_359)
}

func (dm_build_373 *Dm_build_358) Dm_build_372(dm_build_374 int) *Dm_build_358 {
	for i := dm_build_374; i < len(dm_build_373.dm_build_359); i++ {
		dm_build_373.dm_build_359[i] = 0
	}
	dm_build_373.dm_build_359 = dm_build_373.dm_build_359[:dm_build_374]
	return dm_build_373
}

func (dm_build_376 *Dm_build_358) Dm_build_375(dm_build_377 int) *Dm_build_358 {
	dm_build_376.dm_build_360 = dm_build_377
	return dm_build_376
}

func (dm_build_379 *Dm_build_358) Dm_build_378() int {
	return dm_build_379.dm_build_360
}

func (dm_build_381 *Dm_build_358) Dm_build_380(dm_build_382 bool) int {
	return len(dm_build_381.dm_build_359) - dm_build_381.dm_build_360
}

func (dm_build_384 *Dm_build_358) Dm_build_383(dm_build_385 int, dm_build_386 bool, dm_build_387 bool) *Dm_build_358 {

	if dm_build_386 {
		if dm_build_387 {
			dm_build_384.dm_build_365(dm_build_385)
		} else {
			dm_build_384.dm_build_359 = dm_build_384.dm_build_359[:len(dm_build_384.dm_build_359)-dm_build_385]
		}
	} else {
		if dm_build_387 {
			dm_build_384.dm_build_360 += dm_build_385
		} else {
			dm_build_384.dm_build_360 -= dm_build_385
		}
	}

	return dm_build_384
}

func (dm_build_389 *Dm_build_358) Dm_build_388(dm_build_390 io.Reader, dm_build_391 int) (int, error) {
	dm_build_392 := len(dm_build_389.dm_build_359)
	dm_build_389.dm_build_365(dm_build_391)
	dm_build_393 := 0
	for dm_build_391 > 0 {
		n, err := dm_build_390.Read(dm_build_389.dm_build_359[dm_build_392+dm_build_393:])
		if n > 0 && err == io.EOF {
			dm_build_393 += n
			dm_build_389.dm_build_359 = dm_build_389.dm_build_359[:dm_build_392+dm_build_393]
			return dm_build_393, nil
		} else if n > 0 && err == nil {
			dm_build_391 -= n
			dm_build_393 += n
		} else if n == 0 && err != nil {
			return -1, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
		}
	}

	return dm_build_393, nil
}

func (dm_build_395 *Dm_build_358) Dm_build_394(dm_build_396 io.Writer) (*Dm_build_358, error) {
	if _, err := dm_build_396.Write(dm_build_395.dm_build_359); err != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
	}
	return dm_build_395, nil
}

func (dm_build_398 *Dm_build_358) Dm_build_397(dm_build_399 bool) int {
	dm_build_400 := len(dm_build_398.dm_build_359)
	dm_build_398.dm_build_365(1)

	if dm_build_399 {
		return copy(dm_build_398.dm_build_359[dm_build_400:], []byte{1})
	} else {
		return copy(dm_build_398.dm_build_359[dm_build_400:], []byte{0})
	}
}

func (dm_build_402 *Dm_build_358) Dm_build_401(dm_build_403 byte) int {
	dm_build_404 := len(dm_build_402.dm_build_359)
	dm_build_402.dm_build_365(1)

	return copy(dm_build_402.dm_build_359[dm_build_404:], Dm_build_1.Dm_build_179(dm_build_403))
}

func (dm_build_406 *Dm_build_358) Dm_build_405(dm_build_407 int16) int {
	dm_build_408 := len(dm_build_406.dm_build_359)
	dm_build_406.dm_build_365(2)

	return copy(dm_build_406.dm_build_359[dm_build_408:], Dm_build_1.Dm_build_182(dm_build_407))
}

func (dm_build_410 *Dm_build_358) Dm_build_409(dm_build_411 int32) int {
	dm_build_412 := len(dm_build_410.dm_build_359)
	dm_build_410.dm_build_365(4)

	return copy(dm_build_410.dm_build_359[dm_build_412:], Dm_build_1.Dm_build_185(dm_build_411))
}

func (dm_build_414 *Dm_build_358) Dm_build_413(dm_build_415 uint8) int {
	dm_build_416 := len(dm_build_414.dm_build_359)
	dm_build_414.dm_build_365(1)

	return copy(dm_build_414.dm_build_359[dm_build_416:], Dm_build_1.Dm_build_197(dm_build_415))
}

func (dm_build_418 *Dm_build_358) Dm_build_417(dm_build_419 uint16) int {
	dm_build_420 := len(dm_build_418.dm_build_359)
	dm_build_418.dm_build_365(2)

	return copy(dm_build_418.dm_build_359[dm_build_420:], Dm_build_1.Dm_build_200(dm_build_419))
}

func (dm_build_422 *Dm_build_358) Dm_build_421(dm_build_423 uint32) int {
	dm_build_424 := len(dm_build_422.dm_build_359)
	dm_build_422.dm_build_365(4)

	return copy(dm_build_422.dm_build_359[dm_build_424:], Dm_build_1.Dm_build_203(dm_build_423))
}

func (dm_build_426 *Dm_build_358) Dm_build_425(dm_build_427 uint64) int {
	dm_build_428 := len(dm_build_426.dm_build_359)
	dm_build_426.dm_build_365(8)

	return copy(dm_build_426.dm_build_359[dm_build_428:], Dm_build_1.Dm_build_206(dm_build_427))
}

func (dm_build_430 *Dm_build_358) Dm_build_429(dm_build_431 float32) int {
	dm_build_432 := len(dm_build_430.dm_build_359)
	dm_build_430.dm_build_365(4)

	return copy(dm_build_430.dm_build_359[dm_build_432:], Dm_build_1.Dm_build_203(math.Float32bits(dm_build_431)))
}

func (dm_build_434 *Dm_build_358) Dm_build_433(dm_build_435 float64) int {
	dm_build_436 := len(dm_build_434.dm_build_359)
	dm_build_434.dm_build_365(8)

	return copy(dm_build_434.dm_build_359[dm_build_436:], Dm_build_1.Dm_build_206(math.Float64bits(dm_build_435)))
}

func (dm_build_438 *Dm_build_358) Dm_build_437(dm_build_439 []byte) int {
	dm_build_440 := len(dm_build_438.dm_build_359)
	dm_build_438.dm_build_365(len(dm_build_439))
	return copy(dm_build_438.dm_build_359[dm_build_440:], dm_build_439)
}

func (dm_build_442 *Dm_build_358) Dm_build_441(dm_build_443 []byte) int {
	return dm_build_442.Dm_build_409(int32(len(dm_build_443))) + dm_build_442.Dm_build_437(dm_build_443)
}

func (dm_build_445 *Dm_build_358) Dm_build_444(dm_build_446 []byte) int {
	return dm_build_445.Dm_build_413(uint8(len(dm_build_446))) + dm_build_445.Dm_build_437(dm_build_446)
}

func (dm_build_448 *Dm_build_358) Dm_build_447(dm_build_449 []byte) int {
	return dm_build_448.Dm_build_417(uint16(len(dm_build_449))) + dm_build_448.Dm_build_437(dm_build_449)
}

func (dm_build_451 *Dm_build_358) Dm_build_450(dm_build_452 []byte) int {
	return dm_build_451.Dm_build_437(dm_build_452) + dm_build_451.Dm_build_401(0)
}

func (dm_build_454 *Dm_build_358) Dm_build_453(dm_build_455 string, dm_build_456 string, dm_build_457 *DmConnection) int {
	dm_build_458 := Dm_build_1.Dm_build_214(dm_build_455, dm_build_456, dm_build_457)
	return dm_build_454.Dm_build_441(dm_build_458)
}

func (dm_build_460 *Dm_build_358) Dm_build_459(dm_build_461 string, dm_build_462 string, dm_build_463 *DmConnection) int {
	dm_build_464 := Dm_build_1.Dm_build_214(dm_build_461, dm_build_462, dm_build_463)
	return dm_build_460.Dm_build_444(dm_build_464)
}

func (dm_build_466 *Dm_build_358) Dm_build_465(dm_build_467 string, dm_build_468 string, dm_build_469 *DmConnection) int {
	dm_build_470 := Dm_build_1.Dm_build_214(dm_build_467, dm_build_468, dm_build_469)
	return dm_build_466.Dm_build_447(dm_build_470)
}

func (dm_build_472 *Dm_build_358) Dm_build_471(dm_build_473 string, dm_build_474 string, dm_build_475 *DmConnection) int {
	dm_build_476 := Dm_build_1.Dm_build_214(dm_build_473, dm_build_474, dm_build_475)
	return dm_build_472.Dm_build_450(dm_build_476)
}

func (dm_build_478 *Dm_build_358) Dm_build_477() byte {
	dm_build_479 := Dm_build_1.Dm_build_94(dm_build_478.dm_build_359, dm_build_478.dm_build_360)
	dm_build_478.dm_build_360++
	return dm_build_479
}

func (dm_build_481 *Dm_build_358) Dm_build_480() int16 {
	dm_build_482 := Dm_build_1.Dm_build_98(dm_build_481.dm_build_359, dm_build_481.dm_build_360)
	dm_build_481.dm_build_360 += 2
	return dm_build_482
}

func (dm_build_484 *Dm_build_358) Dm_build_483() int32 {
	dm_build_485 := Dm_build_1.Dm_build_103(dm_build_484.dm_build_359, dm_build_484.dm_build_360)
	dm_build_484.dm_build_360 += 4
	return dm_build_485
}

func (dm_build_487 *Dm_build_358) Dm_build_486() int64 {
	dm_build_488 := Dm_build_1.Dm_build_108(dm_build_487.dm_build_359, dm_build_487.dm_build_360)
	dm_build_487.dm_build_360 += 8
	return dm_build_488
}

func (dm_build_490 *Dm_build_358) Dm_build_489() float32 {
	dm_build_491 := Dm_build_1.Dm_build_113(dm_build_490.dm_build_359, dm_build_490.dm_build_360)
	dm_build_490.dm_build_360 += 4
	return dm_build_491
}

func (dm_build_493 *Dm_build_358) Dm_build_492() float64 {
	dm_build_494 := Dm_build_1.Dm_build_117(dm_build_493.dm_build_359, dm_build_493.dm_build_360)
	dm_build_493.dm_build_360 += 8
	return dm_build_494
}

func (dm_build_496 *Dm_build_358) Dm_build_495() uint8 {
	dm_build_497 := Dm_build_1.Dm_build_121(dm_build_496.dm_build_359, dm_build_496.dm_build_360)
	dm_build_496.dm_build_360 += 1
	return dm_build_497
}

func (dm_build_499 *Dm_build_358) Dm_build_498() uint16 {
	dm_build_500 := Dm_build_1.Dm_build_125(dm_build_499.dm_build_359, dm_build_499.dm_build_360)
	dm_build_499.dm_build_360 += 2
	return dm_build_500
}

func (dm_build_502 *Dm_build_358) Dm_build_501() uint32 {
	dm_build_503 := Dm_build_1.Dm_build_130(dm_build_502.dm_build_359, dm_build_502.dm_build_360)
	dm_build_502.dm_build_360 += 4
	return dm_build_503
}

func (dm_build_505 *Dm_build_358) Dm_build_504(dm_build_506 int) []byte {
	dm_build_507 := Dm_build_1.Dm_build_152(dm_build_505.dm_build_359, dm_build_505.dm_build_360, dm_build_506)
	dm_build_505.dm_build_360 += dm_build_506
	return dm_build_507
}

func (dm_build_509 *Dm_build_358) Dm_build_508() []byte {
	return dm_build_509.Dm_build_504(int(dm_build_509.Dm_build_483()))
}

func (dm_build_511 *Dm_build_358) Dm_build_510() []byte {
	return dm_build_511.Dm_build_504(int(dm_build_511.Dm_build_477()))
}

func (dm_build_513 *Dm_build_358) Dm_build_512() []byte {
	return dm_build_513.Dm_build_504(int(dm_build_513.Dm_build_480()))
}

func (dm_build_515 *Dm_build_358) Dm_build_514(dm_build_516 int) []byte {
	return dm_build_515.Dm_build_504(dm_build_516)
}

func (dm_build_518 *Dm_build_358) Dm_build_517() []byte {
	dm_build_519 := 0
	for dm_build_518.Dm_build_477() != 0 {
		dm_build_519++
	}
	dm_build_518.Dm_build_383(dm_build_519, false, false)
	return dm_build_518.Dm_build_504(dm_build_519)
}

func (dm_build_521 *Dm_build_358) Dm_build_520(dm_build_522 int, dm_build_523 string, dm_build_524 *DmConnection) string {
	return Dm_build_1.Dm_build_251(dm_build_521.Dm_build_504(dm_build_522), dm_build_523, dm_build_524)
}

func (dm_build_526 *Dm_build_358) Dm_build_525(dm_build_527 string, dm_build_528 *DmConnection) string {
	return Dm_build_1.Dm_build_251(dm_build_526.Dm_build_508(), dm_build_527, dm_build_528)
}

func (dm_build_530 *Dm_build_358) Dm_build_529(dm_build_531 string, dm_build_532 *DmConnection) string {
	return Dm_build_1.Dm_build_251(dm_build_530.Dm_build_510(), dm_build_531, dm_build_532)
}

func (dm_build_534 *Dm_build_358) Dm_build_533(dm_build_535 string, dm_build_536 *DmConnection) string {
	return Dm_build_1.Dm_build_251(dm_build_534.Dm_build_512(), dm_build_535, dm_build_536)
}

func (dm_build_538 *Dm_build_358) Dm_build_537(dm_build_539 string, dm_build_540 *DmConnection) string {
	return Dm_build_1.Dm_build_251(dm_build_538.Dm_build_517(), dm_build_539, dm_build_540)
}

func (dm_build_542 *Dm_build_358) Dm_build_541(dm_build_543 int, dm_build_544 byte) int {
	return dm_build_542.Dm_build_577(dm_build_543, Dm_build_1.Dm_build_179(dm_build_544))
}

func (dm_build_546 *Dm_build_358) Dm_build_545(dm_build_547 int, dm_build_548 int16) int {
	return dm_build_546.Dm_build_577(dm_build_547, Dm_build_1.Dm_build_182(dm_build_548))
}

func (dm_build_550 *Dm_build_358) Dm_build_549(dm_build_551 int, dm_build_552 int32) int {
	return dm_build_550.Dm_build_577(dm_build_551, Dm_build_1.Dm_build_185(dm_build_552))
}

func (dm_build_554 *Dm_build_358) Dm_build_553(dm_build_555 int, dm_build_556 int64) int {
	return dm_build_554.Dm_build_577(dm_build_555, Dm_build_1.Dm_build_188(dm_build_556))
}

func (dm_build_558 *Dm_build_358) Dm_build_557(dm_build_559 int, dm_build_560 float32) int {
	return dm_build_558.Dm_build_577(dm_build_559, Dm_build_1.Dm_build_191(dm_build_560))
}

func (dm_build_562 *Dm_build_358) Dm_build_561(dm_build_563 int, dm_build_564 float64) int {
	return dm_build_562.Dm_build_577(dm_build_563, Dm_build_1.Dm_build_194(dm_build_564))
}

func (dm_build_566 *Dm_build_358) Dm_build_565(dm_build_567 int, dm_build_568 uint8) int {
	return dm_build_566.Dm_build_577(dm_build_567, Dm_build_1.Dm_build_197(dm_build_568))
}

func (dm_build_570 *Dm_build_358) Dm_build_569(dm_build_571 int, dm_build_572 uint16) int {
	return dm_build_570.Dm_build_577(dm_build_571, Dm_build_1.Dm_build_200(dm_build_572))
}

func (dm_build_574 *Dm_build_358) Dm_build_573(dm_build_575 int, dm_build_576 uint32) int {
	return dm_build_574.Dm_build_577(dm_build_575, Dm_build_1.Dm_build_203(dm_build_576))
}

func (dm_build_578 *Dm_build_358) Dm_build_577(dm_build_579 int, dm_build_580 []byte) int {
	return copy(dm_build_578.dm_build_359[dm_build_579:], dm_build_580)
}

func (dm_build_582 *Dm_build_358) Dm_build_581(dm_build_583 int, dm_build_584 []byte) int {
	return dm_build_582.Dm_build_549(dm_build_583, int32(len(dm_build_584))) + dm_build_582.Dm_build_577(dm_build_583+4, dm_build_584)
}

func (dm_build_586 *Dm_build_358) Dm_build_585(dm_build_587 int, dm_build_588 []byte) int {
	return dm_build_586.Dm_build_541(dm_build_587, byte(len(dm_build_588))) + dm_build_586.Dm_build_577(dm_build_587+1, dm_build_588)
}

func (dm_build_590 *Dm_build_358) Dm_build_589(dm_build_591 int, dm_build_592 []byte) int {
	return dm_build_590.Dm_build_545(dm_build_591, int16(len(dm_build_592))) + dm_build_590.Dm_build_577(dm_build_591+2, dm_build_592)
}

func (dm_build_594 *Dm_build_358) Dm_build_593(dm_build_595 int, dm_build_596 []byte) int {
	return dm_build_594.Dm_build_577(dm_build_595, dm_build_596) + dm_build_594.Dm_build_541(dm_build_595+len(dm_build_596), 0)
}

func (dm_build_598 *Dm_build_358) Dm_build_597(dm_build_599 int, dm_build_600 string, dm_build_601 string, dm_build_602 *DmConnection) int {
	return dm_build_598.Dm_build_581(dm_build_599, Dm_build_1.Dm_build_214(dm_build_600, dm_build_601, dm_build_602))
}

func (dm_build_604 *Dm_build_358) Dm_build_603(dm_build_605 int, dm_build_606 string, dm_build_607 string, dm_build_608 *DmConnection) int {
	return dm_build_604.Dm_build_585(dm_build_605, Dm_build_1.Dm_build_214(dm_build_606, dm_build_607, dm_build_608))
}

func (dm_build_610 *Dm_build_358) Dm_build_609(dm_build_611 int, dm_build_612 string, dm_build_613 string, dm_build_614 *DmConnection) int {
	return dm_build_610.Dm_build_589(dm_build_611, Dm_build_1.Dm_build_214(dm_build_612, dm_build_613, dm_build_614))
}

func (dm_build_616 *Dm_build_358) Dm_build_615(dm_build_617 int, dm_build_618 string, dm_build_619 string, dm_build_620 *DmConnection) int {
	return dm_build_616.Dm_build_593(dm_build_617, Dm_build_1.Dm_build_214(dm_build_618, dm_build_619, dm_build_620))
}

func (dm_build_622 *Dm_build_358) Dm_build_621(dm_build_623 int) byte {
	return Dm_build_1.Dm_build_219(dm_build_622.Dm_build_648(dm_build_623, 1))
}

func (dm_build_625 *Dm_build_358) Dm_build_624(dm_build_626 int) int16 {
	return Dm_build_1.Dm_build_222(dm_build_625.Dm_build_648(dm_build_626, 2))
}

func (dm_build_628 *Dm_build_358) Dm_build_627(dm_build_629 int) int32 {
	return Dm_build_1.Dm_build_225(dm_build_628.Dm_build_648(dm_build_629, 4))
}

func (dm_build_631 *Dm_build_358) Dm_build_630(dm_build_632 int) int64 {
	return Dm_build_1.Dm_build_228(dm_build_631.Dm_build_648(dm_build_632, 8))
}

func (dm_build_634 *Dm_build_358) Dm_build_633(dm_build_635 int) float32 {
	return Dm_build_1.Dm_build_231(dm_build_634.Dm_build_648(dm_build_635, 4))
}

func (dm_build_637 *Dm_build_358) Dm_build_636(dm_build_638 int) float64 {
	return Dm_build_1.Dm_build_234(dm_build_637.Dm_build_648(dm_build_638, 8))
}

func (dm_build_640 *Dm_build_358) Dm_build_639(dm_build_641 int) uint8 {
	return Dm_build_1.Dm_build_237(dm_build_640.Dm_build_648(dm_build_641, 1))
}

func (dm_build_643 *Dm_build_358) Dm_build_642(dm_build_644 int) uint16 {
	return Dm_build_1.Dm_build_240(dm_build_643.Dm_build_648(dm_build_644, 2))
}

func (dm_build_646 *Dm_build_358) Dm_build_645(dm_build_647 int) uint32 {
	return Dm_build_1.Dm_build_243(dm_build_646.Dm_build_648(dm_build_647, 4))
}

func (dm_build_649 *Dm_build_358) Dm_build_648(dm_build_650 int, dm_build_651 int) []byte {
	return dm_build_649.dm_build_359[dm_build_650 : dm_build_650+dm_build_651]
}

func (dm_build_653 *Dm_build_358) Dm_build_652(dm_build_654 int) []byte {
	dm_build_655 := dm_build_653.Dm_build_627(dm_build_654)
	return dm_build_653.Dm_build_648(dm_build_654+4, int(dm_build_655))
}

func (dm_build_657 *Dm_build_358) Dm_build_656(dm_build_658 int) []byte {
	dm_build_659 := dm_build_657.Dm_build_621(dm_build_658)
	return dm_build_657.Dm_build_648(dm_build_658+1, int(dm_build_659))
}

func (dm_build_661 *Dm_build_358) Dm_build_660(dm_build_662 int) []byte {
	dm_build_663 := dm_build_661.Dm_build_624(dm_build_662)
	return dm_build_661.Dm_build_648(dm_build_662+2, int(dm_build_663))
}

func (dm_build_665 *Dm_build_358) Dm_build_664(dm_build_666 int) []byte {
	dm_build_667 := 0
	for dm_build_665.Dm_build_621(dm_build_666) != 0 {
		dm_build_666++
		dm_build_667++
	}

	return dm_build_665.Dm_build_648(dm_build_666-dm_build_667, int(dm_build_667))
}

func (dm_build_669 *Dm_build_358) Dm_build_668(dm_build_670 int, dm_build_671 string, dm_build_672 *DmConnection) string {
	return Dm_build_1.Dm_build_251(dm_build_669.Dm_build_652(dm_build_670), dm_build_671, dm_build_672)
}

func (dm_build_674 *Dm_build_358) Dm_build_673(dm_build_675 int, dm_build_676 string, dm_build_677 *DmConnection) string {
	return Dm_build_1.Dm_build_251(dm_build_674.Dm_build_656(dm_build_675), dm_build_676, dm_build_677)
}

func (dm_build_679 *Dm_build_358) Dm_build_678(dm_build_680 int, dm_build_681 string, dm_build_682 *DmConnection) string {
	return Dm_build_1.Dm_build_251(dm_build_679.Dm_build_660(dm_build_680), dm_build_681, dm_build_682)
}

func (dm_build_684 *Dm_build_358) Dm_build_683(dm_build_685 int, dm_build_686 string, dm_build_687 *DmConnection) string {
	return Dm_build_1.Dm_build_251(dm_build_684.Dm_build_664(dm_build_685), dm_build_686, dm_build_687)
}
