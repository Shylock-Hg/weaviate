//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package entities

const (
	// StrategyReplace allows for idem-potent PUT where the latest takes presence
	StrategyReplace       = "replace"
	StrategySetCollection = "setcollection"
	StrategyMapCollection = "mapcollection"
	StrategyRoaringSet    = "roaringset"
	StrategyInverted      = "inverted"
)

type SegmentStrategy uint16

const (
	SegmentStrategyReplace SegmentStrategy = iota
	SegmentStrategySetCollection
	SegmentStrategyMapCollection
	SegmentStrategyRoaringSet
	SegmentStrategyInverted
)

func SegmentStrategyFromString(in string) SegmentStrategy {
	switch in {
	case StrategyReplace:
		return SegmentStrategyReplace
	case StrategySetCollection:
		return SegmentStrategySetCollection
	case StrategyMapCollection:
		return SegmentStrategyMapCollection
	case StrategyRoaringSet:
		return SegmentStrategyRoaringSet
	case StrategyInverted:
		return SegmentStrategyInverted
	default:
		panic("unsupported strategy")
	}
}
