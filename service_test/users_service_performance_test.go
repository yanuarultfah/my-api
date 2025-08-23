//go:build perf

package service_test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my-api/domain/users"
	"my-api/service"
)

// Performance tests WITHOUT gmeasure to avoid extra module hassles.
// Latency dicatat manual lalu dihitung P95.
//
// Jalankan:
//   ginkgo -v -r --tags=perf --label-filter=perf ./service_test
// atau:
//   go test -v -tags=perf ./service_test

func TestUsersServicePerformance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "usersService Performance Suite")
}

func p95(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	cp := make([]float64, len(values))
	copy(cp, values)
	sort.Float64s(cp)
	// nearest-rank method
	idx := int(0.95*float64(len(cp))+0.5) - 1
	if idx < 0 {
		idx = 0
	}
	if idx >= len(cp) {
		idx = len(cp) - 1
	}
	return cp[idx]
}

var _ = Describe("usersService (performance)", Label("perf"), func() {
	var svc = service.UsersService

	It("CreateUser P95 latency < 200ms (50 sampel)", func() {
		const samples = 50
		lat := make([]float64, 0, samples)

		for i := 0; i < samples; i++ {
			id := fmt.Sprintf("perf-%d-%d", i, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
			u := users.User{UserId: id, Status: "active", Email: fmt.Sprintf("user-%d@example.com", i)}

			start := time.Now()
			_, err := svc.CreateUser(u)
			Expect(err).To(BeNil())

			lat = append(lat, float64(time.Since(start).Milliseconds()))
		}

		p := p95(lat)
		By(fmt.Sprintf("CreateUser P95: %.2f ms", p))
		Expect(p).To(BeNumerically("<", 200.0), "P95 latency too high")
	})

	It("SearchByStatus P95 latency < 150ms (30 sampel, warm)", func() {
		// Warm up data
		for i := 0; i < 20; i++ {
			id := fmt.Sprintf("perf-warm-%d-%d", i, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
			u := users.User{UserId: id, Status: "active"}
			_, _ = svc.CreateUser(u)
		}

		const samples = 30
		lat := make([]float64, 0, samples)

		for i := 0; i < samples; i++ {
			start := time.Now()
			_, err := svc.SearchByStatus("active")
			Expect(err).To(BeNil())

			lat = append(lat, float64(time.Since(start).Milliseconds()))
		}

		p := p95(lat)
		By(fmt.Sprintf("SearchByStatus P95: %.2f ms", p))
		Expect(p).To(BeNumerically("<", 150.0), "P95 latency too high")
	})
})
