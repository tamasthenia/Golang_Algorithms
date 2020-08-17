package main

import "strings"

/*
Given a valid (IPv4) IP address, return a defanged version of that IP address.

A defanged IP address replaces every period "." with "[.]".

Example 1:

Input: address = "1.1.1.1"
Output: "1[.]1[.]1[.]1"
*/

/*func defangIPaddr(address string) string {


    arrOfString := strings.Split(address, ".")

    ans := ""

    for i, v := range arrOfString{

        if i == len(arrOfString)-1{
            ans = ans + v
        } else {
            ans = ans + v + "[.]"
        }
    }

    return ans
}*/

func defangIPaddr(address string) string {

	return strings.Replace(address, ".", "[.]", -1)

}
