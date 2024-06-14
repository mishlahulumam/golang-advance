package slice_test

import (
	"golang-advance/session-5-validator/entity"
	"golang-advance/session-5-validator/repository/slice"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserRepository(t *testing.T) {
	repo := slice.NewUserRepository([]entity.User{})

	t.Run("CreateUser", func(t *testing.T) {
		newUser := entity.User{Name: "Kobe", Email: "kobe@example.com", Password: "P4ssw0rd"}
		createdUser := repo.CreateUser(&newUser)

		require.Equal(t, 0, createdUser.ID)
		require.Equal(t, "Kobe", createdUser.Name)
		require.Equal(t, "kobe@example.com", createdUser.Email)
		require.NotZero(t, createdUser.CreatedAt)
		require.NotZero(t, createdUser.UpdatedAt)
	})

	t.Run("GetUserByID", func(t *testing.T) {
		user, found := repo.GetUserByID(0)

		require.True(t, found)
		require.Equal(t, "Kobe", user.Name)

		_, notFound := repo.GetUserByID(99)
		require.False(t, notFound)
	})

	t.Run("UpdateUser", func(t *testing.T) {
		update := entity.User{Name: "Kobe Updated", Email: "kobe.updated@example.com", Password: "newpassword"}
		updatedUser, found := repo.UpdateUser(0, update)

		require.True(t, found)
		require.Equal(t, "Kobe Updated", updatedUser.Name)
		require.Equal(t, "kobe.updated@example.com", updatedUser.Email)

		_, notFound := repo.UpdateUser(99, update)
		require.False(t, notFound)
	})

	t.Run("DeleteUser", func(t *testing.T) {
		deleted := repo.DeleteUser(0)
		require.True(t, deleted)

		notDeleted := repo.DeleteUser(99)
		require.False(t, notDeleted)
	})

	t.Run("GetAllUsers", func(t *testing.T) {
		allUsers := repo.GetAllUsers()
		require.Empty(t, allUsers)
	})
}
