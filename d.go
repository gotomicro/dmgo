/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"container/list"
	"io"
)

type Dm_build_280 struct {
	dm_build_281 *list.List
	dm_build_282 *dm_build_334
	dm_build_283 int
}

func Dm_build_284() *Dm_build_280 {
	return &Dm_build_280{
		dm_build_281: list.New(),
		dm_build_283: 0,
	}
}

func (dm_build_286 *Dm_build_280) Dm_build_285() int {
	return dm_build_286.dm_build_283
}

func (dm_build_288 *Dm_build_280) Dm_build_287(dm_build_289 *Dm_build_358, dm_build_290 int) int {
	var dm_build_291 = 0
	var dm_build_292 = 0
	for dm_build_291 < dm_build_290 && dm_build_288.dm_build_282 != nil {
		dm_build_292 = dm_build_288.dm_build_282.dm_build_342(dm_build_289, dm_build_290-dm_build_291)
		if dm_build_288.dm_build_282.dm_build_337 == 0 {
			dm_build_288.dm_build_324()
		}
		dm_build_291 += dm_build_292
		dm_build_288.dm_build_283 -= dm_build_292
	}
	return dm_build_291
}

func (dm_build_294 *Dm_build_280) Dm_build_293(dm_build_295 []byte, dm_build_296 int, dm_build_297 int) int {
	var dm_build_298 = 0
	var dm_build_299 = 0
	for dm_build_298 < dm_build_297 && dm_build_294.dm_build_282 != nil {
		dm_build_299 = dm_build_294.dm_build_282.dm_build_346(dm_build_295, dm_build_296, dm_build_297-dm_build_298)
		if dm_build_294.dm_build_282.dm_build_337 == 0 {
			dm_build_294.dm_build_324()
		}
		dm_build_298 += dm_build_299
		dm_build_294.dm_build_283 -= dm_build_299
		dm_build_296 += dm_build_299
	}
	return dm_build_298
}

func (dm_build_301 *Dm_build_280) Dm_build_300(dm_build_302 io.Writer, dm_build_303 int) int {
	var dm_build_304 = 0
	var dm_build_305 = 0
	for dm_build_304 < dm_build_303 && dm_build_301.dm_build_282 != nil {
		dm_build_305 = dm_build_301.dm_build_282.dm_build_351(dm_build_302, dm_build_303-dm_build_304)
		if dm_build_301.dm_build_282.dm_build_337 == 0 {
			dm_build_301.dm_build_324()
		}
		dm_build_304 += dm_build_305
		dm_build_301.dm_build_283 -= dm_build_305
	}
	return dm_build_304
}

func (dm_build_307 *Dm_build_280) Dm_build_306(dm_build_308 []byte, dm_build_309 int, dm_build_310 int) {
	if dm_build_310 == 0 {
		return
	}
	var dm_build_311 = dm_build_338(dm_build_308, dm_build_309, dm_build_310)
	if dm_build_307.dm_build_282 == nil {
		dm_build_307.dm_build_282 = dm_build_311
	} else {
		dm_build_307.dm_build_281.PushBack(dm_build_311)
	}
	dm_build_307.dm_build_283 += dm_build_310
}

func (dm_build_313 *Dm_build_280) dm_build_312(dm_build_314 int) byte {
	var dm_build_315 = dm_build_314
	var dm_build_316 = dm_build_313.dm_build_282
	for dm_build_315 > 0 && dm_build_316 != nil {
		if dm_build_316.dm_build_337 == 0 {
			continue
		}
		if dm_build_315 > dm_build_316.dm_build_337-1 {
			dm_build_315 -= dm_build_316.dm_build_337
			dm_build_316 = dm_build_313.dm_build_281.Front().Value.(*dm_build_334)
		} else {
			break
		}
	}
	return dm_build_316.dm_build_355(dm_build_315)
}
func (dm_build_318 *Dm_build_280) Dm_build_317(dm_build_319 *Dm_build_280) {
	if dm_build_319.dm_build_283 == 0 {
		return
	}
	var dm_build_320 = dm_build_319.dm_build_282
	for dm_build_320 != nil {
		dm_build_318.dm_build_321(dm_build_320)
		dm_build_319.dm_build_324()
		dm_build_320 = dm_build_319.dm_build_282
	}
	dm_build_319.dm_build_283 = 0
}
func (dm_build_322 *Dm_build_280) dm_build_321(dm_build_323 *dm_build_334) {
	if dm_build_323.dm_build_337 == 0 {
		return
	}
	if dm_build_322.dm_build_282 == nil {
		dm_build_322.dm_build_282 = dm_build_323
	} else {
		dm_build_322.dm_build_281.PushBack(dm_build_323)
	}
	dm_build_322.dm_build_283 += dm_build_323.dm_build_337
}

func (dm_build_325 *Dm_build_280) dm_build_324() {
	var dm_build_326 = dm_build_325.dm_build_281.Front()
	if dm_build_326 == nil {
		dm_build_325.dm_build_282 = nil
	} else {
		dm_build_325.dm_build_282 = dm_build_326.Value.(*dm_build_334)
		dm_build_325.dm_build_281.Remove(dm_build_326)
	}
}

func (dm_build_328 *Dm_build_280) Dm_build_327() []byte {
	var dm_build_329 = make([]byte, dm_build_328.dm_build_283)
	var dm_build_330 = dm_build_328.dm_build_282
	var dm_build_331 = 0
	var dm_build_332 = len(dm_build_329)
	var dm_build_333 = 0
	for dm_build_330 != nil {
		if dm_build_330.dm_build_337 > 0 {
			if dm_build_332 > dm_build_330.dm_build_337 {
				dm_build_333 = dm_build_330.dm_build_337
			} else {
				dm_build_333 = dm_build_332
			}
			copy(dm_build_329[dm_build_331:dm_build_331+dm_build_333], dm_build_330.dm_build_335[dm_build_330.dm_build_336:dm_build_330.dm_build_336+dm_build_333])
			dm_build_331 += dm_build_333
			dm_build_332 -= dm_build_333
		}
		if dm_build_328.dm_build_281.Front() == nil {
			dm_build_330 = nil
		} else {
			dm_build_330 = dm_build_328.dm_build_281.Front().Value.(*dm_build_334)
		}
	}
	return dm_build_329
}

type dm_build_334 struct {
	dm_build_335 []byte
	dm_build_336 int
	dm_build_337 int
}

func dm_build_338(dm_build_339 []byte, dm_build_340 int, dm_build_341 int) *dm_build_334 {
	return &dm_build_334{
		dm_build_339,
		dm_build_340,
		dm_build_341,
	}
}

func (dm_build_343 *dm_build_334) dm_build_342(dm_build_344 *Dm_build_358, dm_build_345 int) int {
	if dm_build_343.dm_build_337 <= dm_build_345 {
		dm_build_345 = dm_build_343.dm_build_337
	}
	dm_build_344.Dm_build_437(dm_build_343.dm_build_335[dm_build_343.dm_build_336 : dm_build_343.dm_build_336+dm_build_345])
	dm_build_343.dm_build_336 += dm_build_345
	dm_build_343.dm_build_337 -= dm_build_345
	return dm_build_345
}

func (dm_build_347 *dm_build_334) dm_build_346(dm_build_348 []byte, dm_build_349 int, dm_build_350 int) int {
	if dm_build_347.dm_build_337 <= dm_build_350 {
		dm_build_350 = dm_build_347.dm_build_337
	}
	copy(dm_build_348[dm_build_349:dm_build_349+dm_build_350], dm_build_347.dm_build_335[dm_build_347.dm_build_336:dm_build_347.dm_build_336+dm_build_350])
	dm_build_347.dm_build_336 += dm_build_350
	dm_build_347.dm_build_337 -= dm_build_350
	return dm_build_350
}

func (dm_build_352 *dm_build_334) dm_build_351(dm_build_353 io.Writer, dm_build_354 int) int {
	if dm_build_352.dm_build_337 <= dm_build_354 {
		dm_build_354 = dm_build_352.dm_build_337
	}
	dm_build_353.Write(dm_build_352.dm_build_335[dm_build_352.dm_build_336 : dm_build_352.dm_build_336+dm_build_354])
	dm_build_352.dm_build_336 += dm_build_354
	dm_build_352.dm_build_337 -= dm_build_354
	return dm_build_354
}
func (dm_build_356 *dm_build_334) dm_build_355(dm_build_357 int) byte {
	return dm_build_356.dm_build_335[dm_build_356.dm_build_336+dm_build_357]
}
