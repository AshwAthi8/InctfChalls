Matching100

In this challenge user is supposed to enter a string input and there are three checks that user needs to pass inorder to get flag.

There are two integer arrays that store random values from 0 to 255


First check compares the length of the string with 30 if the check fails then "Bad Length" message is printed out.


Before we move to second check the input string is divided into two character arrays on the basis of odd and even indexes



Second check is used to check the array containing odd indexes of input. It first converts the character value to integer and then searches the value of each index in the first random array and when found the index is noted(lets assume i is the index) and then value of second random array at the noted ith index is compared with the given char array and if they are equal 1 is added to a variable which was 0 initially.


Similarly Third check is used to check the values of array containing even indexes of input.


After the completion of check2 and check3 their return value is compared with 15 . If the comparison passes then flag is printed out otherwise "Wrong Input" message is printed out.


Correct Input:   F!N4lLY_u_G0t_Th3_@0rR3@T_F\4G


Flag :     inctf{F!N4lLY_u_G0t_Th3_@0rR3@T_F\4G}
