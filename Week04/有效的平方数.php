<?php

class Solution
{

    /**
     * @param Integer $num
     * @return Boolean
     */
    function isPerfectSquare($num)
    {
        if ($num == 0) {
            return false;
        }
        $min = 0;
        $max = $num;
        while ($min <= $max) {
            $mid = $min + ($max - $min) >> 1;
            if ($mid * $mid == $num) {
                return true;
            } elseif ($mid * $mid < $num) {
                $min = $mid + 1;
            } else {
                $max = $mid - 1;
            }
        }
        return false;
    }
}