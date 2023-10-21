package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// before
	fmt.Println("Sebelum unit test")

	m.Run()

	// after
	fmt.Println("Setelah unit test")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Semua")

	if result == "Hi Semua" {
		fmt.Println("TestHelloWorld done")
	}
}

func TestHelloWorldHanan(t *testing.T) {
	result := HelloWorld("Hanan")

	if result != "Hello Hanan" {
		// error
		t.Fail()
	}

	fmt.Println("TestHelloWorldHanan done")
}

func TestHelloWorldCathy(t *testing.T) {
	result := HelloWorld("Cathy")

	if result != "Hello Cathy" {
		// error
		t.FailNow()
	}

	fmt.Println("TestHelloWorldCathy done")
}

func TestHelloWorldHanan1(t *testing.T) {
	result := HelloWorld("Hanan")

	if result != "Hello Hanan" {
		// error
		t.Error("Harusnya Hi Hanan")
	}

	fmt.Println("TestHelloWorldHanan1 done")
}

func TestHelloWorldCathy2(t *testing.T) {
	result := HelloWorld("Cathy")

	if result != "Hello Cathy" {
		// error
		t.Fatal("Harusnya Hi Cathy")
	}

	fmt.Println("TestHelloWorldCathy2 done")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Cathy")
	assert.Equal(t, "Hi Cathy", result, "Harusnya Hi Cathy")
	fmt.Println("TestHelloWorldAssertion done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Cathy")
	require.Equal(t, "Hi Cathy", result, "Harusnya Hi Cathy")
	fmt.Println("TestHelloWorldAssertion done")
}

func TestHelloWorldSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on macOs")
	}

	result := HelloWorld("Cathy")
	require.Equal(t, "Hi Cathy", result, "Harusnya Hi Cathy")
}

func TestSubTest(t *testing.T) {
	t.Run("Hanan", func(t *testing.T) {
		result := HelloWorld("Hanan")
		require.Equal(t, "Hi Hanan", result, "Harusnya Hi Hanan")
	})

	t.Run("Cathy", func(t *testing.T) {
		result := HelloWorld("Cathy")
		require.Equal(t, "Hi Cathy", result, "Harusnya Hi Cathy")
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Hanan",
			request:  "Hanan",
			expected: "Hi Hanan",
		},
		{
			name:     "Cathy",
			request:  "Cathy",
			expected: "Hi Cathy",
		},
		{
			name:     "Cathy",
			request:  "Cathy",
			expected: "Hi Cathy",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Hanan")
	}
}

func BenchmarkHelloWorldCathy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Cathy")
	}
}

func BenchmarkSubTest(b *testing.B) {
	b.Run("Hanan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Hanan")
		}
	})

	b.Run("Cathy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Cathy")
		}
	})
}

func BenchmarkTableHelloWorld(b *testing.B) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Hanan",
			request:  "Hanan",
			expected: "Hi Hanan",
		},
		{
			name:     "Cathy",
			request:  "Cathy",
			expected: "Hi Cathy",
		},
		{
			name:     "Cathy",
			request:  "Cathy",
			expected: "Hi Cathy",
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(test.request)
			}
		})
	}
}
