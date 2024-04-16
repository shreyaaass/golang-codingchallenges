# JSON Parser


## Check if JSON is Valid
### Algorithm
- create a tokenizer(split by  whitespaces and \n)
- Check if it starts with a left bracket
- If it does, we start parsing the object
    - We then verify if the next token is either a right bracket or a name value pair
    - In case of a name value pair, we split by colon
    - The first part is the name, and the second part is the value
    - Name should be a string
    - Value could be a string, number, boolean, null, or an object or an array
    - If the value is an object or an array, we recursively call the function
    - Else we check for the validity of the types
- If it doesn't, we return false


##
- We check if there are blank spaces in front of the object
- If there are, we remove them
- Then we replace "\n" with "", and split by commas
- For each of the element we check if it is a valid name value pair or not