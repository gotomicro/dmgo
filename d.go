/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm

import (
	"container/list"
	"io"
)

type Dm_build_283 struct {
	dm_build_284 *list.List
	dm_build_285 *dm_build_337
	dm_build_286 int
}

func Dm_build_287() *Dm_build_283 {
	return &Dm_build_283{
		dm_build_284: list.New(),
		dm_build_286: 0,
	}
}

func (dm_build_289 *Dm_build_283) Dm_build_288() int {
	return dm_build_289.dm_build_286
}

func (dm_build_291 *Dm_build_283) Dm_build_290(dm_build_292 *Dm_build_361, dm_build_293 int) int {
	var dm_build_294 = 0
	var dm_build_295 = 0
	for dm_build_294 < dm_build_293 && dm_build_291.dm_build_285 != nil {
		dm_build_295 = dm_build_291.dm_build_285.dm_build_345(dm_build_292, dm_build_293-dm_build_294)
		if dm_build_291.dm_build_285.dm_build_340 == 0 {
			dm_build_291.dm_build_327()
		}
		dm_build_294 += dm_build_295
		dm_build_291.dm_build_286 -= dm_build_295
	}
	return dm_build_294
}

func (dm_build_297 *Dm_build_283) Dm_build_296(dm_build_298 []byte, dm_build_299 int, dm_build_300 int) int {
	var dm_build_301 = 0
	var dm_build_302 = 0
	for dm_build_301 < dm_build_300 && dm_build_297.dm_build_285 != nil {
		dm_build_302 = dm_build_297.dm_build_285.dm_build_349(dm_build_298, dm_build_299, dm_build_300-dm_build_301)
		if dm_build_297.dm_build_285.dm_build_340 == 0 {
			dm_build_297.dm_build_327()
		}
		dm_build_301 += dm_build_302
		dm_build_297.dm_build_286 -= dm_build_302
		dm_build_299 += dm_build_302
	}
	return dm_build_301
}

func (dm_build_304 *Dm_build_283) Dm_build_303(dm_build_305 io.Writer, dm_build_306 int) int {
	var dm_build_307 = 0
	var dm_build_308 = 0
	for dm_build_307 < dm_build_306 && dm_build_304.dm_build_285 != nil {
		dm_build_308 = dm_build_304.dm_build_285.dm_build_354(dm_build_305, dm_build_306-dm_build_307)
		if dm_build_304.dm_build_285.dm_build_340 == 0 {
			dm_build_304.dm_build_327()
		}
		dm_build_307 += dm_build_308
		dm_build_304.dm_build_286 -= dm_build_308
	}
	return dm_build_307
}

func (dm_build_310 *Dm_build_283) Dm_build_309(dm_build_311 []byte, dm_build_312 int, dm_build_313 int) {
	if dm_build_313 == 0 {
		return
	}
	var dm_build_314 = dm_build_341(dm_build_311, dm_build_312, dm_build_313)
	if dm_build_310.dm_build_285 == nil {
		dm_build_310.dm_build_285 = dm_build_314
	} else {
		dm_build_310.dm_build_284.PushBack(dm_build_314)
	}
	dm_build_310.dm_build_286 += dm_build_313
}

func (dm_build_316 *Dm_build_283) dm_build_315(dm_build_317 int) byte {
	var dm_build_318 = dm_build_317
	var dm_build_319 = dm_build_316.dm_build_285
	for dm_build_318 > 0 && dm_build_319 != nil {
		if dm_build_319.dm_build_340 == 0 {
			continue
		}
		if dm_build_318 > dm_build_319.dm_build_340-1 {
			dm_build_318 -= dm_build_319.dm_build_340
			dm_build_319 = dm_build_316.dm_build_284.Front().Value.(*dm_build_337)
		} else {
			break
		}
	}
	return dm_build_319.dm_build_358(dm_build_318)
}
func (dm_build_321 *Dm_build_283) Dm_build_320(dm_build_322 *Dm_build_283) {
	if dm_build_322.dm_build_286 == 0 {
		return
	}
	var dm_build_323 = dm_build_322.dm_build_285
	for dm_build_323 != nil {
		dm_build_321.dm_build_324(dm_build_323)
		dm_build_322.dm_build_327()
		dm_build_323 = dm_build_322.dm_build_285
	}
	dm_build_322.dm_build_286 = 0
}
func (dm_build_325 *Dm_build_283) dm_build_324(dm_build_326 *dm_build_337) {
	if dm_build_326.dm_build_340 == 0 {
		return
	}
	if dm_build_325.dm_build_285 == nil {
		dm_build_325.dm_build_285 = dm_build_326
	} else {
		dm_build_325.dm_build_284.PushBack(dm_build_326)
	}
	dm_build_325.dm_build_286 += dm_build_326.dm_build_340
}

func (dm_build_328 *Dm_build_283) dm_build_327() {
	var dm_build_329 = dm_build_328.dm_build_284.Front()
	if dm_build_329 == nil {
		dm_build_328.dm_build_285 = nil
	} else {
		dm_build_328.dm_build_285 = dm_build_329.Value.(*dm_build_337)
		dm_build_328.dm_build_284.Remove(dm_build_329)
	}
}

func (dm_build_331 *Dm_build_283) Dm_build_330() []byte {
	var dm_build_332 = make([]byte, dm_build_331.dm_build_286)
	var dm_build_333 = dm_build_331.dm_build_285
	var dm_build_334 = 0
	var dm_build_335 = len(dm_build_332)
	var dm_build_336 = 0
	for dm_build_333 != nil {
		if dm_build_333.dm_build_340 > 0 {
			if dm_build_335 > dm_build_333.dm_build_340 {
				dm_build_336 = dm_build_333.dm_build_340
			} else {
				dm_build_336 = dm_build_335
			}
			copy(dm_build_332[dm_build_334:dm_build_334+dm_build_336], dm_build_333.dm_build_338[dm_build_333.dm_build_339:dm_build_333.dm_build_339+dm_build_336])
			dm_build_334 += dm_build_336
			dm_build_335 -= dm_build_336
		}
		if dm_build_331.dm_build_284.Front() == nil {
			dm_build_333 = nil
		} else {
			dm_build_333 = dm_build_331.dm_build_284.Front().Value.(*dm_build_337)
		}
	}
	return dm_build_332
}

type dm_build_337 struct {
	dm_build_338 []byte
	dm_build_339 int
	dm_build_340 int
}

func dm_build_341(dm_build_342 []byte, dm_build_343 int, dm_build_344 int) *dm_build_337 {
	return &dm_build_337{
		dm_build_342,
		dm_build_343,
		dm_build_344,
	}
}

func (dm_build_346 *dm_build_337) dm_build_345(dm_build_347 *Dm_build_361, dm_build_348 int) int {
	if dm_build_346.dm_build_340 <= dm_build_348 {
		dm_build_348 = dm_build_346.dm_build_340
	}
	dm_build_347.Dm_build_444(dm_build_346.dm_build_338[dm_build_346.dm_build_339 : dm_build_346.dm_build_339+dm_build_348])
	dm_build_346.dm_build_339 += dm_build_348
	dm_build_346.dm_build_340 -= dm_build_348
	return dm_build_348
}

func (dm_build_350 *dm_build_337) dm_build_349(dm_build_351 []byte, dm_build_352 int, dm_build_353 int) int {
	if dm_build_350.dm_build_340 <= dm_build_353 {
		dm_build_353 = dm_build_350.dm_build_340
	}
	copy(dm_build_351[dm_build_352:dm_build_352+dm_build_353], dm_build_350.dm_build_338[dm_build_350.dm_build_339:dm_build_350.dm_build_339+dm_build_353])
	dm_build_350.dm_build_339 += dm_build_353
	dm_build_350.dm_build_340 -= dm_build_353
	return dm_build_353
}

func (dm_build_355 *dm_build_337) dm_build_354(dm_build_356 io.Writer, dm_build_357 int) int {
	if dm_build_355.dm_build_340 <= dm_build_357 {
		dm_build_357 = dm_build_355.dm_build_340
	}
	dm_build_356.Write(dm_build_355.dm_build_338[dm_build_355.dm_build_339 : dm_build_355.dm_build_339+dm_build_357])
	dm_build_355.dm_build_339 += dm_build_357
	dm_build_355.dm_build_340 -= dm_build_357
	return dm_build_357
}
func (dm_build_359 *dm_build_337) dm_build_358(dm_build_360 int) byte {
	return dm_build_359.dm_build_338[dm_build_359.dm_build_339+dm_build_360]
}
