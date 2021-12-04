package main

import (
	"testing"
)

func getTestValues() []string {
	return []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010"}
}

func TestMostCommon(t *testing.T) {
	gotPos1 := mostCommon(getTestValues(), 1)
	if gotPos1 != "1" {
		t.Errorf("Testing mostCommon1 got(%s) and wanted %d", gotPos1, 1)
	}
	gotPos2 := mostCommon(getTestValues(), 2)
	if gotPos2 != "0" {
		t.Errorf("Testing mostCommon2 got(%s) and wanted %d", gotPos2, 0)
	}
	gotPos3 := mostCommon(getTestValues(), 3)
	if gotPos3 != "1" {
		t.Errorf("Testing mostCommon3 got(%s) and wanted %d", gotPos3, 1)
	}
	gotPos4 := mostCommon(getTestValues(), 4)
	if gotPos4 != "1" {
		t.Errorf("Testing mostCommon4 got(%s) and wanted %d", gotPos4, 1)
	}
	gotPos5 := mostCommon(getTestValues(), 5)
	if gotPos5 != "0" {
		t.Errorf("Testing mostCommon5 got(%s) and wanted %d", gotPos5, 0)
	}
}

func TestMostCommonWithTie(t *testing.T) {
	got := mostCommon([]string {"1", "0"}, 1)

	if got != "1" {
		t.Errorf("Testing mostCommonWithTie got(%s) and wanted %s", got, "1")
	}
}

func TestBitStringToInt(t *testing.T) {
	got := bitStringToInt("10110")
	if got != 22 {
		t.Errorf("Testing bitStringToInt got(%d) and wanted %d", got, 22)
	}
}

func TestBitStringToIntWithError(t *testing.T) {
	got := bitStringToInt("22222")
	if got != -1 {
		t.Errorf("Testing bitStringToInt got(%d) and wanted %d", got, -1)
	}
}

func TestCalcGamaRate(t *testing.T) {
	got := calcGamaRate(getTestValues())

	if got != 22 {
		t.Errorf("Testing calcGamaRate got(%d) and wanted %d", got, 22)
	}
}

func TestCalcEpsilonRate(t *testing.T) {
	got := calcEpsilonRate(getTestValues())

	if got != 9 {
		t.Errorf("Testing calcGamaRate got(%d) and wanted %d", got, 9)
	}
}

func TestCalcPoserConsumption(t *testing.T) {
	got := calcPoserConsumption(getTestValues())

	if got != 198 {
		t.Errorf("Testing calcGamaRate got(%d) and wanted %d", got, 198)
	}
}

func TestFilterBitByPosition(t *testing.T) {
	got := filterBitByPosition(getTestValues(), 1, "1")

	if len(got) != 7 {
		t.Errorf("Testing filterBitByPosition got list with len(%d) and wanted %d", len(got), 7)
	}
}

func TestCalcOxygenGeneratorRating(t *testing.T) {
	got := calcOxygenGeneratorRating(getTestValues())

	if got != 23 {
		t.Errorf("Testing calcOxygenGeneratorRating got(%d) and wanted %d", got, 23)
	}
}

func TestCalcCo2ScrubberRating(t *testing.T) {
	got := calcCo2ScrubberRating(getTestValues())

	if got != 10 {
		t.Errorf("Testing calcOxygenGeneratorRating got(%d) and wanted %d", got, 10)
	}
}

func TestLifeSupportRating(t *testing.T) {
	got := calcLifeSupportRating(getTestValues())

	if got != 230 {
		t.Errorf("Testing calcOxygenGeneratorRating got(%d) and wanted %d", got, 230)
	}
}
