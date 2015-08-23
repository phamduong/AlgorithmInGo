package sort

import (
	"github.com/phamduong/algorithm/fixture"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	assert := assert.New(t)
	var n int = fixture.RandomOneInt()
	input := fixture.GenerateArray(n)
	expected := fixture.GetExpectedResult(input)

	SelectionSort(input, n)

	assert.Equal(expected, input, "Selection sort error")
}

func TestInsertionSort(t *testing.T) {
	assert := assert.New(t)
	var n int = fixture.RandomOneInt()
	input := fixture.GenerateArray(n)
	expected := fixture.GetExpectedResult(input)

	InsertionSort(input, n)

	assert.Equal(expected, input, "Insertion sort error")
}

func TestBinaryInsertionSort(t *testing.T) {
	assert := assert.New(t)
	var n int = fixture.RandomOneInt()
	input := fixture.GenerateArray(n)
	expected := fixture.GetExpectedResult(input)

	BinaryInsertionSort(input, n)

	assert.Equal(expected, input, "BinaryInsertion sort error")
}

func TestInterchangeSort(t *testing.T) {
	assert := assert.New(t)
	var n int = fixture.RandomOneInt()
	input := fixture.GenerateArray(n)
	expected := fixture.GetExpectedResult(input)

	InterchangeSort(input, n)

	assert.Equal(expected, input, "Interchange sort error")
}

func TestBubleSort(t *testing.T) {
	assert := assert.New(t)
	var n int = fixture.RandomOneInt()
	input := fixture.GenerateArray(n)
	expected := fixture.GetExpectedResult(input)

	BubbleSort(input, n)

	assert.Equal(expected, input, "Buble sort error")
}

func TestShakerSort(t *testing.T) {
	assert := assert.New(t)
	var n int = fixture.RandomOneInt()
	input := fixture.GenerateArray(n)
	expected := fixture.GetExpectedResult(input)

	ShakerSort(input, n)

	assert.Equal(expected, input, "Shaker sort error")
}

func TestHeapSort(t *testing.T) {
	assert := assert.New(t)
	var n int = fixture.RandomOneInt()
	input := fixture.GenerateArray(n)
	expected := fixture.GetExpectedResult(input)

	HeapSort(input, n)

	assert.Equal(expected, input, "Heap sort error")
}
