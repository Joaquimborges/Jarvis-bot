package util

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestContainsValue(t *testing.T) {
	for _, tc := range []struct {
		name    string
		message string
		values  []string
		want    bool
	}{
		//given
		{
			name:    "should return true",
			message: "Lorem Ipsum is simply dummy text of the foo printing and typesetting bar industry",
			values:  []string{"foo", "bar"},
			want:    true,
		},
		{
			name:    "should return false",
			message: "Lorem Ipsum is simply dummy text of the printing and typesetting industry",
			values:  []string{"foo", "bar"},
			want:    false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			//when
			got := ContainsValue(tc.message, tc.values)
			//then
			require.Equal(t, got, tc.want)
		})
	}
}

func TestSliceEnvs(t *testing.T) {
	envs := "foo|bar"
	want := []string{"foo", "bar"}
	got := SliceEnvs(envs)
	require.NotNil(t, got)
	require.Equal(t, got, want)
}

func TestParseDate(t *testing.T) {
	given := time.Date(2024, time.January, 13, 12, 02, 23, 23, time.UTC)
	want := "13-January-2024"
	got := ParseDate(given)
	require.NotNil(t, got)
	require.Equal(t, got, want)
}

func TestGetNumberValueFromMessage(t *testing.T) {
	for _, tc := range []struct {
		message string
		want    int
	}{
		{
			message: "Lorem Ipsum is simply dummy text of the 10 printing and typesetting industry",
			want:    10,
		},
		{
			message: "Lorem Ipsum is simply dummy text of the foo printing and typesetting bar industry",
			want:    0,
		},
	} {
		t.Run("TestGetNumberValueFromMessage()", func(t *testing.T) {
			got := GetNumberValueFromMessage(tc.message)
			require.Equal(t, got, tc.want)
		})
	}
}

func TestBuildComparableTime(t *testing.T) {
	days := 7
	month := 0
	expectedNow := time.Now()
	expectedThen := expectedNow.AddDate(0, month, -days)

	now, then := BuildComparableTime(-days, month)
	require.WithinDuration(t, now, expectedNow, time.Second)
	require.WithinDuration(t, expectedThen, then, time.Second)
}

func TestCreateLocalTime(t *testing.T) {
	t.Run("should build with Sao_Paulo localtime", func(t *testing.T) {
		locale := "America/Sao_Paulo"
		location, err := time.LoadLocation(locale)
		require.NoError(t, err)

		expected := time.Now().In(location)
		got := CreateLocalTime(locale)
		require.NotNil(t, got)
		require.WithinDuration(t, got, expected, time.Second)
	})

	t.Run("should build with UTC", func(t *testing.T) {
		expected := time.Now().UTC()
		got := CreateLocalTime("foo")
		require.NotNil(t, got)
		require.WithinDuration(t, got, expected, time.Second)
	})
}
