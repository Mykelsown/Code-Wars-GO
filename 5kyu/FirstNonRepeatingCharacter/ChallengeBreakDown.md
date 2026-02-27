# My Ways Of Breaking Challenges Down To Simple Problems
Here contains how i solved the problem by chopping it down to different small pieces and the methods I used to to solve this challenge.

I noticed that taking time to actually do this would help to set focus on a particular section, help saves time during build and leads to a more structured code.

## **Title:** *First Non-Repeating Character*
## **Question:** 

Write a function that takes a string input, and returns the first character that is not repeated anywhere in the string.

For example, if given the input "stress", the function should return 't', since the letter t only occurs once in the string, and occurs first in the string.

As an added challenge, upper- and lowercase characters are considered the same character, but the function should return the correct case for the initial character. For example, the input "sTreSS" should return "T".

If a string contains only repeating characters, return an empty string ("");

Note: despite its name in some languages, your function should handle any Unicode codepoint:

"@#@@*"    --> "#"
"かか何"   --> "何"
"🐐🦊🐐" --> "🦊"

## **Breaking Down:** 

- *Step-1:* After receiving the string, then i converted it to lowercase. why did i convert to lowercase? I did this for easy and uniform working with the srting.

- *Step-2:*  Now how to actually check for the non-repeating character. I will acheive this by looping the string, check char1 if it matches any other characters in the string, if yes, i break out of the loop then add it to a map, where the value will be true. So, i go to the char2 in the string, then i repeat the same process again. Eventually if a character with no match is found, i still add to the map, but the value will be set to false.

- *Step-3:* To avoid adding of a character that has already been added to the map, being added again, i set a condition that says don't add if character already exist in map. 

- *Step-4:* After all the getting a value in the map that has false, i will get that key then find the index from the string inputed(the one i already made lowercase). After that i will then, check the position of the character in the string, so as to be able to match it with the position on the original string inputed(The string original case). Then i return the character on that position.

- *Step-5:* If all the values in the map isn't false, return empty string.


