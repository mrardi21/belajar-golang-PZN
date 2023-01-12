package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ardi")
	if result != "Hello Ardi" {
		// panic("Result is not 'Hello Ardi'")
		// t.Fail()
		t.Error("Result must be 'Hello Ardi'")
	}
	fmt.Println("TestHelloWorld Done")

}
func TestHelloWorld2(t *testing.T) {
	result := HelloWorld("Mardiansyah")
	if result != "Hello Mardiansyah" {
		// panic("Result is not 'Hello Mardiansyah'")
		// t.FailNow()
		t.Fatal("Result must be 'Hello Mardiansyah'")
	}
	fmt.Println("TestHelloWorld2 Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Ardi")
	assert.Equal(t, "Hello Ardi", result, "Result mush be 'Hello Ardi'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Ardi")
	require.Equal(t, "Hello Ardi", result, "Result mush be 'Hello Ardi'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run on Windows")
	}
	result := HelloWorld("Ardi")
	require.Equal(t, "Hello Ardi", result, "Result mush be 'Hello Ardi'")
}

func TestMain(m *testing.M) {
	// Before
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	// After
	fmt.Println("AFTER UNIT TEST")
}

func TestSubTest(t *testing.T) {
	t.Run("Ardi", func(t *testing.T) {
		result := HelloWorld("Ardi")
		assert.Equal(t, "Hello Ardi", result, "Result mush be 'Hello Ardi'")
	})
	t.Run("Mardiansyah", func(t *testing.T) {
		result := HelloWorld("Mardiansyah")
		assert.Equal(t, "Hello Mardiansyah", result, "Result mush be 'Hello Mardiansyah'")
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Ardi",
			request:  "Ardi",
			expected: "Hello Ardi",
		},
		{
			name:     "Mardiansyah",
			request:  "Mardiansyah",
			expected: "Hello Mardiansyah",
		},
		{
			name:     "Ganteng",
			request:  "Ganteng",
			expected: "Hello Ganteng",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

/* -------------------------------- Benchmark ------------------------------- */
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ardi")
	}
}
func BenchmarkHelloWorldMardiansyah(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Mardiansyah")
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Ardi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ardi")
		}
	})
	b.Run("Mardiansyah", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Mardiansyah")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	bencmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Ardi",
			request: "Ardi",
		},
		{
			name:    "Mardiansyah",
			request: "Mardiansyah",
		},
		{
			name:    "Ardi Mardiansyah",
			request: "Ardi Mardiansyah",
		},
	}

	for _, bencmark := range bencmarks {
		b.Run(bencmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bencmark.request)
			}
		})
	}
}
