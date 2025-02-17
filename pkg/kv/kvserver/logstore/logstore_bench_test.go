// Copyright 2022 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package logstore

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/cockroachdb/cockroach/pkg/kv/kvserver/raftentry"
	"github.com/cockroachdb/cockroach/pkg/kv/kvserver/raftlog"
	"github.com/cockroachdb/cockroach/pkg/settings/cluster"
	"github.com/cockroachdb/cockroach/pkg/storage"
	"github.com/cockroachdb/cockroach/pkg/util/humanizeutil"
	"github.com/cockroachdb/cockroach/pkg/util/log"
	"github.com/cockroachdb/cockroach/pkg/util/metric"
	"github.com/stretchr/testify/require"
	"go.etcd.io/raft/v3/raftpb"
)

type discardBatch struct {
	storage.Batch
}

func (b *discardBatch) Commit(bool) error {
	return nil
}

func BenchmarkLogStore_StoreEntries(b *testing.B) {
	defer log.Scope(b).Close(b)
	const kb = 1 << 10
	const mb = 1 << 20
	for _, bytes := range []int64{1 * kb, 256 * kb, 512 * kb, 1 * mb, 2 * mb} {
		b.Run(fmt.Sprintf("bytes=%s", humanizeutil.IBytes(bytes)), func(b *testing.B) {
			runBenchmarkLogStore_StoreEntries(b, bytes)
		})
	}
}

func runBenchmarkLogStore_StoreEntries(b *testing.B, bytes int64) {
	const tenMB = 10 * 1 << 20
	ec := raftentry.NewCache(tenMB)
	const rangeID = 1
	eng := storage.NewDefaultInMemForTesting()
	defer eng.Close()
	s := LogStore{
		RangeID:     rangeID,
		Engine:      eng,
		StateLoader: NewStateLoader(rangeID),
		EntryCache:  ec,
		Settings:    cluster.MakeTestingClusterSettings(),
		Metrics: Metrics{
			RaftLogCommitLatency: metric.NewHistogram(metric.Metadata{}, 10*time.Second, metric.IOLatencyBuckets),
		},
	}

	ctx := context.Background()
	rs := RaftState{
		LastTerm: 1,
		ByteSize: 0,
	}
	var ents []raftpb.Entry
	data := make([]byte, bytes)
	rand.New(rand.NewSource(0)).Read(data)
	ents = append(ents, raftpb.Entry{
		Term:  1,
		Index: 1,
		Type:  raftpb.EntryNormal,
		Data:  raftlog.EncodeRaftCommand(raftlog.EntryEncodingStandardWithoutAC, "deadbeef", data),
	})
	stats := &AppendStats{}

	b.ReportAllocs()
	b.ResetTimer()
	// Use a batch that ignores Commit() so that we can run as many iterations as
	// we like without building up large amounts of data in the Engine. This hides
	// some allocations that would occur down in pebble but that's fine, we're not
	// here to measure those.
	batch := &discardBatch{}
	for i := 0; i < b.N; i++ {
		batch.Batch = newStoreEntriesBatch(eng)
		rd := Ready{
			HardState: raftpb.HardState{},
			Entries:   ents,
			MustSync:  true,
		}
		var err error
		rs, err = s.storeEntriesAndCommitBatch(ctx, rs, rd, stats, batch)
		if err != nil {
			b.Fatal(err)
		}
		batch.Batch.Close()
		ents[0].Index++
	}
	require.EqualValues(b, b.N, rs.LastIndex)
}
