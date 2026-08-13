/*
 * @lc app=leetcode id=273 lang=golang
 *
 * [273] Integer to English Words
 */

// @lc code=start
func numberToWords(num int) string {
    if num == 0 {
        return "Zero"
    }
    res := []string{}

    k := [][]string{
        {
            "", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine",
            "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen",
            "Seventeen", "Eighteen", "Nineteen", "Twenty",
        },
        {
            "","Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety",
        },
        {
            "", "Hundred",
        },
        {
            
                "",
                "Thousand",
                "Million",
                "Billion",
                // "Trillion",
                // "Quadrillion",
                // "Quintillion",
                // "Sextillion",
                // "Septillion",
                // "Octillion",
                // "Nonillion",
                // "Decillion",
        },
    }

    
    i := 0
    for num > 0 {
        tmp := num % 1000
        num /= 1000

        if tmp > 0 {
            res = append(res, k[3][i])
        }

        i++

        v := tmp % 100
        if v > 20 {
            if v % 10 != 0 {
                res = append(res, k[0][v%10])
            }
            res = append(res, k[1][v/10])
        } else if v > 0 {
            res = append(res, k[0][v])
        }

        if tmp > 99 {
            res = append(res, k[2][1])
            res = append(res, k[0][tmp/100])
        }
    }
    
    rev := 0
    for rev < len(res) / 2 {
        res[rev], res[len(res) - rev - 1] = res[len(res) - rev - 1], res[rev]
        rev ++
    }

    r := strings.Join(res, " ")
    r = strings.Trim(r, " ")
    return r
}
// @lc code=end

