package main

import (
    "fmt"
    "testing"
    "reflect"
)

func TestPrefixSums(t *testing.T) {
    // No wrap
    res := PrefixSums([]int64{1,2,3,4}, 99999)

    if !reflect.DeepEqual(res, []int64{1,3,6,10}) {
        t.Error(fmt.Sprintf("Incorrect prefix array %v\n", res))
    }
    // Modulo
    res = PrefixSums([]int64{1,2,3,4}, 5)

    if !reflect.DeepEqual(res, []int64{1,3,1,0}) {
        t.Error(fmt.Sprintf("Incorrect prefix array %v\n", res))
    }

    // Test case 18
    res = PrefixSums([]int64{1,5,9}, 5)

    if !reflect.DeepEqual(res, []int64{1,1,0}) {
        t.Error(fmt.Sprintf("Incorrect prefix array: %v\n", res))
    }
}

func TestModSub(t *testing.T) {
    res := ModSub(int64(15), int64(6), 5)

    if res != 4 {
        t.Error(fmt.Sprintf("ModSub result is wrong, %v != 4\n", res))
    }

    res = ModSub(int64(0), int64(1), 5)
    
    if res != 4 {
        t.Error(fmt.Sprintf("ModSub result is wrong, %v != 4\n", res))
    }
}

func TestMaximumSum(t *testing.T) {
    res := maximumSum([]int64{3,3,9,9,5}, 7)

    if res != 6 {
        t.Error(fmt.Sprintf("sample input 0 result (%v) != 6", res))
    }

    res = maximumSum([]int64{1,2,3}, 2)

    if res != 1 {
        t.Error(fmt.Sprintf("test case 17 result (%v) != 1", res))
    }

    res = maximumSum([]int64{1,5,9}, 5)

    if res != 4 {
        t.Error(fmt.Sprintf("test case 18 result (%v) != 4", res))
    }
}

func TestCaseOne(t *testing.T) {
    res := maximumSum([]int64{
        846930887, 1681692778, 1714636916, 1957747794, 424238336, 719885387,
        1649760493, 596516650, 1189641422, 1025202363, 1350490028, 783368691,
        1102520060, 2044897764, 1967513927, 1365180541, 1540383427, 304089173,
        1303455737, 35005212, 521595369, 294702568, 1726956430, 336465783,
        861021531, 278722863, 233665124, 2145174068, 468703136, 1101513930,
        1801979803, 1315634023, 635723059, 1369133070, 1125898168, 1059961394,
        2089018457, 628175012, 1656478043, 1131176230, 1653377374, 859484422,
        1914544920, 608413785, 756898538, 1734575199, 1973594325, 149798316,
        2038664371, 1129566414}, 1804289384)

    if res != 1802192837 {
        t.Error(fmt.Sprintf("test case 1 result 1 (%v) != 1802192837", res))
    }
}
