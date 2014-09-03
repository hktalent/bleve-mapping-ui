//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package simple

import (
	"github.com/blevesearch/bleve/registry"
	"github.com/blevesearch/bleve/search/highlight"
)

const Name = "simple"

const defaultFragmentSize = 200

type Fragmenter struct {
	fragmentSize int
}

func NewFragmenter(fragmentSize int) *Fragmenter {
	return &Fragmenter{
		fragmentSize: fragmentSize,
	}
}

func (s *Fragmenter) Fragment(orig []byte, ot highlight.TermLocations) []*highlight.Fragment {
	rv := make([]*highlight.Fragment, 0)

	maxbegin := 0
	for currTermIndex, termLocation := range ot {
		// start with with this
		// it should be the highest scoring fragment with this term first
		start := termLocation.Start
		end := start + s.fragmentSize
		if end > len(orig) {
			end = len(orig)
			// we hit end, so push back as far as we can without crossing maxbegin
			extra := s.fragmentSize - (end - start)
			if start-extra >= maxbegin {
				start -= extra
			} else {
				start = maxbegin
			}
		}
		// however, we'd rather have the tokens centered more in the frag
		// lets try to do that as best we can, without affecting the score
		// find the end of the last term in this fragment
		minend := end
		for _, innerTermLocation := range ot[currTermIndex:] {
			if innerTermLocation.End > end {
				break
			}
			minend = innerTermLocation.End
		}

		// find the smaller of the two rooms to move
		roomToMove := end - minend
		if start-maxbegin < roomToMove {
			roomToMove = start - maxbegin
		}

		offset := roomToMove / 2
		rv = append(rv, &highlight.Fragment{Orig: orig, Start: start - offset, End: end - offset})
		// set maxbegin to the end of the current term location
		// so that next one won't back up to include it
		maxbegin = termLocation.End

	}
	if len(ot) == 0 {
		// if there were no terms to highlight
		// produce a single fragment from the beginning
		start := 0
		end := start + s.fragmentSize
		if end > len(orig) {
			end = len(orig)
		}
		rv = append(rv, &highlight.Fragment{Orig: orig, Start: start, End: end})
	}

	return rv
}

func Constructor(config map[string]interface{}, cache *registry.Cache) (highlight.Fragmenter, error) {
	size := defaultFragmentSize
	sizeVal, ok := config["size"].(float64)
	if ok {
		size = int(sizeVal)
	}
	return NewFragmenter(size), nil
}

func init() {
	registry.RegisterFragmenter(Name, Constructor)
}
