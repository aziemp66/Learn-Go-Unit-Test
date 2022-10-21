package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Sebelum Unit Test")

	m.Run()

	fmt.Println("Setelah Unit Test")
}

func Test_HelloWorld(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko!" {
		// unit test failed
		// t.Fail()
		t.Error(
			"Text is Wrong Dumbass, Result Must be 'Hello Eko!'",
		) // t.fail but you can put arguments inside it
	}
	fmt.Println("Test_HelloWorld Done")
}

func Test_HelloWorldAzie(t *testing.T) {
	result := HelloWorld("Azie")
	if result != "Hello Azie!" {
		// t.FailNow()
		t.Fatal(
			"Text is Wrong you dickhead, Result Must be 'Hello Azie!'",
		) // t.FailNow but you can put arguments inside it
	}
	fmt.Println("Test_HelloWorldAzie Done")
}

func Test_HelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Azie")
	assert.Equal(t, "Hello Azie!", result, "Result Must be 'Hello Azie!'")
	fmt.Println("Dieksekusi")
}

func Test_HelloWorldRequire(t *testing.T) {
	result := HelloWorld("Melza")
	require.Equal(t, "Hello Melza!", result, "Result Must Be 'Hello Melza!'")
	fmt.Println("Tidak Dieksekusi")
}

func Test_Skip(t *testing.T) {
	Os := runtime.GOOS
	fmt.Println(Os)
	if Os == "linux" {
		t.Skip("Unit Tes Tidak Bisa Jalan di Linux")
	}

	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko!", result)
}

func Test_SubTest(t *testing.T) {
	t.Run("Melza", func(t *testing.T) {
		result := HelloWorld("Melza")
		require.Equal(t, "Hello Melza!", result)
	})
	t.Run("Pratama", func(t *testing.T) {
		result := HelloWorld("Pratama")
		require.Equal(t, "Hello Pratama!", result)
	})
}

func Test_HelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Eko",
			request:  "Eko",
			expected: "Hello Eko!",
		},
		{
			name:     "Melza",
			request:  "Melza",
			expected: "Hello Melza!",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Azie")
	}
}
