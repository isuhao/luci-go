// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package grpcmon

import (
	"testing"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/luci/luci-go/common/clock"
	"github.com/luci/luci-go/common/clock/testclock"
	"github.com/luci/luci-go/common/tsmon/distribution"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUnaryClientInterceptor(t *testing.T) {
	Convey("Captures count and duration", t, func() {
		c, memStore := testContext()

		// Fake call that runs for 500 ms.
		invoker := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
			clock.Get(ctx).(testclock.TestClock).Add(500 * time.Millisecond)
			return nil
		}

		// Run the call with the interceptor.
		NewUnaryClientInterceptor(nil)(c, "/service/method", nil, nil, nil, invoker)

		count, err := memStore.Get(c, grpcClientCount, time.Time{}, []interface{}{"/service/method", 0})
		So(err, ShouldBeNil)
		So(count, ShouldEqual, 1)

		duration, err := memStore.Get(c, grpcClientDuration, time.Time{}, []interface{}{"/service/method", 0})
		So(err, ShouldBeNil)
		So(duration.(*distribution.Distribution).Sum(), ShouldEqual, 500)
	})
}
