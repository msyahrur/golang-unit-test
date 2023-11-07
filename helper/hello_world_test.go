package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Syahrur",
			request:  "Syahrur",
			expected: "Hello Syahrur",
		},
		{
			name:     "Roma",
			request:  "Roma",
			expected: "Hello Roma",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Result must be "+test.expected)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Syahrur", func(t *testing.T) {
		result := HelloWorld("Syahrur")
		require.Equal(t, "Hello Syahrur", result, "Result must be 'Hello Syahrur'")

	})
	t.Run("Roma", func(t *testing.T) {
		result := HelloWorld("Roma")
		require.Equal(t, "Hello Roma", result, "Result must be 'Hello Syahrur'")
	})
}

// WAJIB TestMain
func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	fmt.Println("AFTER UNIT TEST")
}
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not use Windows")
	}
	//terskip
	result := HelloWorld("Syahrur")
	require.Equal(t, "Hello Syahrur", result, "Result must be 'Hello Syahrur'")

}

// Kalau tanpa library di saran menggunakan error dan fatal krn ada keterangan/informasi.
// SANGAT DI SARAN MENGGUNAKAN ASSERT DAN REQUIRE DARIPADA IF ELSE FAIL
func TestHelloWorldRequire(t *testing.T) {
	//Equal = nilai yg kita harapkan
	result := HelloWorld("Syahrur")
	//assert Equal menggunakan Fail Now jd TIDAK dilanjut ke bwah
	require.Equal(t, "Hello Syahrur", result, "Result must be 'Hello Syahrur'")
	fmt.Println("TestHelloWorldAssert with Assert Done ")
}
func TestHelloWorldAssert(t *testing.T) {
	//Equal = nilai yg kita harapkan
	result := HelloWorld("Syahrur")
	//assert Equal menggunakan Fail jd dilanjut ke bwah
	assert.Equal(t, "Hello Syahrur", result, "Result must be 'Hello Syahrur'")
	fmt.Println("TestHelloWorldAssert with Assert Done ")
}
func TestHelloWorldSyahrur(t *testing.T) {
	result := HelloWorld("Syahrur")
	if result != "Hello Syahrur" {
		// error
		//t.Fail() // lanjutkan ke bwah
		t.Error("Result must be 'Hello Syahrur'")
	}
	fmt.Println("TestHelloWorldSyahrur Done")
}
func TestHelloWorldRoma(t *testing.T) {
	result := HelloWorld("Roma")
	if result != "Hello Roma" {
		// error
		//t.FailNow() //tidak lanjutkan ke bwah
		t.Fatal("Result must be 'Hello Roma'") //tidak lanjutkan ke bwah
	}
	fmt.Println("TestHelloWorldRoma Done")
}
