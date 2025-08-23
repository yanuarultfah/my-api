//go:build integration

package service_test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"my-api/domain/users"
	"my-api/service"
	"my-api/utils/errors"
)

// NOTE:
// These are true integration tests. They call the real users domain package (my-api/domain/users).
// Make sure your test environment (DB, config, etc.) is set up via environment variables
// that the users package reads (e.g., DSN).
// Run with:  ginkgo -v -r --tags=integration --label-filter=integration
// or:        go test -tags=integration ./...

func TestUsersServiceIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "usersService Integration Suite")
}

var _ = Describe("usersService (integration)", Label("integration"), func() {
	var svc = service.UsersService

	BeforeEach(func() {
		// Basic sanity check to help avoid running on an unconfigured env.
		// Customize this according to your users package expectations.
		if os.Getenv("DSN") == "" && os.Getenv("DATABASE_URL") == "" {
			Skip("No DSN/DATABASE_URL set; skipping integration tests")
		}
		// svc = service.UsersService
	})

	When("CreateUser is called with a valid user", func() {
		It("should persist and return the user", func() {
			id := fmt.Sprintf("it-%d", rand.New(rand.NewSource(time.Now().UnixNano())).Int())
			u := users.User{
				UserId: id,
				Status: "active",
				// Add other required fields for your domain/users.Validate() here
				// e.g. Email, Name, etc.
			}

			created, err := svc.CreateUser(u)
			Expect(err).To(BeNil(), "CreateUser returned error: %v", err)
			Expect(created).NotTo(BeNil())
			Expect(created.UserId).To(Equal(id))

			By("fetching the same user with GetUser")
			got, err2 := svc.GetUser(id)
			Expect(err2).To(BeNil(), "GetUser returned error: %v", err2)
			Expect(got).NotTo(BeNil())
			Expect(got.UserId).To(Equal(id))
		})
	})

	When("CreateUser is called with invalid data", func() {
		It("should fail validation", func() {
			u := users.User{}
			created, err := svc.CreateUser(u)
			Expect(created).To(BeNil())
			Expect(err).NotTo(BeNil())
			// your users.Validate() should return *errors.RestErr; assert type
			var restErr *errors.RestErr
			Expect(err).To(BeAssignableToTypeOf(restErr))
			// Optional: Expect(restErr.Status).To(Equal(http.StatusBadRequest))
		})
	})

	When("SearchByStatus is called", func() {
		It("should return a non-empty list for known status", func() {
			// Make at least one user to ensure non-empty result.
			id := fmt.Sprintf("it-%d", rand.New(rand.NewSource(time.Now().UnixNano())).Int())
			u := users.User{UserId: id, Status: "active"}
			_, _ = svc.CreateUser(u) // ignore error if duplicate

			list, err := svc.SearchByStatus("active")
			Expect(err).To(BeNil(), "SearchByStatus returned error: %v", err)
			Expect(list).NotTo(BeNil())
			Expect(len(list)).To(BeNumerically(">=", 1))
		})
	})
})
