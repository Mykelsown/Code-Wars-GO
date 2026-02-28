# My Ways Of Breaking Challenges Down To Simple Problems
Here contains how i solved the problem by chopping it down to different small pieces and the methods I used to to solve this challenge.

I noticed that taking time to actually do this would help to set focus on a particular section, help saves time during build and leads to a more structured code.

## **Title:** *Not Very Secure*
## **Question:** 

In this example you have to validate if a user input string is alphanumeric. The given string is not nil/null/NULL/None, so you don't have to check that.

The string has the following conditions to be alphanumeric:

At least one character ("" is not valid)
Allowed characters are uppercase / lowercase latin letters and digits from 0 to 9
No whitespaces / underscore

## **Breaking Down:** 

- *Step-1:* check for empty string. That is, IF the length of the string is LESS THAN 1, the input isn't valid.

- *Step-2:* create a for range loop, so as to get the characters in the string individually as a rune for match up of the ASCII. 

- *Step-3:* Inside of the loop, then i will check with a condition if the characters in the string are matching the ASCII standard alphanumeric value, if YES, then the function returns true, if NO, it will return false.

